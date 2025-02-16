package wallet

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ixre/go2o/core/domain/interface/wallet"
	"github.com/ixre/go2o/core/infrastructure/domain"
	"github.com/ixre/go2o/core/msq"
	"github.com/ixre/gof/algorithm"
	"github.com/ixre/gof/util"
	"strconv"
	"strings"
	"time"
)

var _ wallet.IWallet = new(WalletImpl)

func NewWallet(v *wallet.Wallet, repo wallet.IWalletRepo) wallet.IWallet {
	return &WalletImpl{
		_value: v,
		_repo:  repo,
	}
}

type WalletImpl struct {
	_value *wallet.Wallet
	_repo  wallet.IWalletRepo
}

func (w *WalletImpl) GetAggregateRootId() int64 {
	return w._value.Id
}

func (w *WalletImpl) Hash() string {
	if w._value.HashCode == "" {
		str := fmt.Sprintf("%d&%d*%d", w._value.UserId, w._value.WalletType, time.Now().Unix())
		hash := algorithm.DJBHash([]byte(str))
		w._value.HashCode = strconv.Itoa(hash)
	}
	return w._value.HashCode
}

func (w *WalletImpl) NodeId() int {
	if w._value.NodeId <= 0 {
		w._value.NodeId = int(w._value.UserId * int64(w._value.WalletType) % 10)
	}
	return w._value.NodeId
}

func (w *WalletImpl) Get() wallet.Wallet {
	return *w._value
}

func (w *WalletImpl) State() int {
	return int(w._value.State)
}

func (w *WalletImpl) getLog(logId int64) *wallet.WalletLog {
	return w._repo.GetLog(w.GetAggregateRootId(), logId)
}

func (w *WalletImpl) GetLog(logId int64) wallet.WalletLog {
	wl := w.getLog(logId)
	if wl != nil {
		return *wl
	}
	return wallet.WalletLog{}
}

func (w *WalletImpl) Save() (int64, error) {
	unix := time.Now().Unix()
	// 初始化
	if w.GetAggregateRootId() <= 0 {
		w.initWallet(unix)
	}
	err := w.checkWallet()
	// 保存
	if err == nil {
		w._value.UpdateTime = unix
		id, err2 := util.I64Err(w._repo.SaveWallet_(w._value))
		if err2 != nil {
			return id, err
		}
		w._value.Id = id
	}
	return w.GetAggregateRootId(), err
}
func (w *WalletImpl) initWallet(unix int64) {
	w._value.CreateTime = unix
	w._value.State = wallet.StatNormal
	w._value.WalletFlag = wallet.FlagCharge | wallet.FlagDiscount
	if w._value.WalletType <= 0 {
		w._value.WalletType = wallet.TPerson
	} else if w._value.WalletType != wallet.TMerchant &&
		w._value.WalletType != wallet.TPerson {
		panic("not support wallet type" + strconv.Itoa(w._value.WalletType))
	}
	if w._value.HashCode == "" {
		w.Hash()
	}
	if w._value.NodeId <= 0 {
		w.NodeId()
	}
}

// 检查钱包
func (w *WalletImpl) checkWallet() error {
	if w._value.UserId <= 0 {
		panic("incorrect wallet user id")
	}
	if flag := w._value.WalletFlag; flag <= 0 {
		panic("incorrect wallet flag:" + strconv.Itoa(flag))
	}
	if w._value.WalletName == "" || len(w._value.WalletName) > 40 {
		return wallet.ErrWalletName
	}
	// 判断是否存在
	match := w._repo.CheckWalletUserMatch(w._value.UserId,
		w._value.WalletType, w.GetAggregateRootId())
	if !match {
		return wallet.ErrSingletonWallet
	}
	return nil
}

// 检查数值、操作人员信息
func (w *WalletImpl) checkValueOpu(value int, checkOpu bool, operatorUid int, operatorName string) error {
	if value == 0 {
		return wallet.ErrAmountZero
	}
	if checkOpu && (operatorUid <= 0 || len(operatorName) == 0) {
		return wallet.ErrMissingOperator
	}
	return w.checkWalletState(w, false)
}

// 检查钱包状态
func (w *WalletImpl) checkWalletState(iw wallet.IWallet, target bool) error {
	if iw == nil {
		if target {
			return wallet.ErrNoSuchTargetWalletAccount
		}
		return wallet.ErrNoSuchWalletAccount
	}
	switch iw.State() {
	case wallet.StatNormal:
		return nil
	case wallet.StatDisabled:
		if target {
			return wallet.ErrTargetWalletAccountNotService
		}
		return wallet.ErrWalletDisabled
	case wallet.StatClosed:
		if target {
			return wallet.ErrTargetWalletAccountNotService
		}
		return wallet.ErrWalletClosed
	}
	panic("unknown wallet state")
}

// 创建钱包日志
func (w *WalletImpl) createWalletLog(kind int, value int, title string, operatorUid int,
	operatorName string) *wallet.WalletLog {
	unix := time.Now().Unix()
	return &wallet.WalletLog{
		WalletId:     w.GetAggregateRootId(),
		WalletUser:   w._value.UserName,
		Kind:         kind,
		Title:        strings.TrimSpace(title),
		OuterChan:    "",
		OuterNo:      "",
		Value:        int64(value),
		OperatorUid:  operatorUid,
		OperatorName: strings.TrimSpace(operatorName),
		Remark:       "",
		ReviewState:  wallet.ReviewPass,
		ReviewRemark: "",
		ReviewTime:   0,
		CreateTime:   unix,
		UpdateTime:   unix,
	}
}

// 保存钱包日志
func (w *WalletImpl) saveWalletLog(l *wallet.WalletLog) error {
	if l.Kind <= 0 {
		return errors.New("wallet log kind error")
	}
	if l.Value == 0 {
		return errors.New("incorrect value")
	}
	l.Title = strings.TrimSpace(l.Title)
	if l.Title == "" {
		return errors.New("wallet log title can't empty")
	}
	isUpdate := l.Id > 0
	id, err := util.I64Err(w._repo.SaveWalletLog_(l))
	if err == nil {
		l.Id = id
	}
	// 将钱包变动订阅到消息队列
	bytes, _ := json.Marshal(map[string]interface{}{
		"id":            id,
		"wallet_type":   w._value.WalletType,
		"user_id":       w._value.UserId,
		"update":        isUpdate,
		"amount":        l.Value,
		"procedure_fee": l.ProcedureFee,
		"balance":       l.Balance,
		"title":         l.Title,
		"outer_no":      l.OuterNo,
	})
	msq.Push(msq.WalletLogTopic, string(bytes))
	return err
}

// Adjust 调整余额，可能存在扣为负数的情况，需传入操作人员编号或操作人员名称
func (w *WalletImpl) Adjust(value int, title, outerNo string,
	remark string, operatorUid int, operatorName string) error {
	err := w.checkValueOpu(value, true, operatorUid, operatorName)
	if err == nil {
		w._value.AdjustAmount += value
		w._value.Balance += int64(value)
		l := w.createWalletLog(wallet.KAdjust, value, title, operatorUid, operatorName)
		l.OuterNo = outerNo
		l.Remark = remark
		l.Balance = w._value.Balance
		err = w.saveWalletLog(l)
		if err == nil {
			_, err = w.Save()
		}
	}
	return err
}

// Consume 消费
func (w *WalletImpl) Consume(amount int, title string, outerNo string, remark string) error {
	if amount > 0 {
		amount = -amount
	}
	if w._value.Balance < -int64(amount) {
		return wallet.ErrOutOfAmount
	}
	w._value.Balance += int64(amount)
	w._value.TotalPay += -amount
	l := w.createWalletLog(wallet.KConsume, amount, title, 0, "")
	l.OuterNo = outerNo
	l.Balance = w._value.Balance
	l.Remark = remark
	err := w.saveWalletLog(l)
	if err == nil {
		_, err = w.Save()
	}
	return err
}

// Discount 支付抵扣,must是否必须大于0
func (w *WalletImpl) Discount(value int, title, outerNo string, must bool) error {
	err := w.checkValueOpu(value, false, 0, "")
	if err == nil {
		if value > 0 {
			value = -value
		}
		if must && w._value.Balance < -int64(value) {
			return wallet.ErrOutOfAmount
		}
		w._value.Balance += int64(value)
		w._value.TotalPay += -value
		l := w.createWalletLog(wallet.KDiscount, value, title, 0, "")
		l.OuterNo = outerNo
		l.Balance = w._value.Balance
		err := w.saveWalletLog(l)
		if err == nil {
			_, err = w.Save()
		}
	}
	return err
}

func (w *WalletImpl) Freeze(data wallet.OperateData, operator wallet.Operator) (int, error) {
	err := w.checkValueOpu(data.Amount, false, operator.OperatorUid, operator.OperatorName)
	if err == nil {
		if data.Amount > 0 {
			data.Amount = -data.Amount
		}
		if w._value.Balance < -int64(data.Amount) {
			return 0, wallet.ErrOutOfAmount
		}
		w._value.Balance += int64(data.Amount)
		w._value.FreezeAmount += -data.Amount
		l := w.createWalletLog(wallet.KFreeze, data.Amount, data.Title, operator.OperatorUid, operator.OperatorName)
		l.OuterNo = data.OuterNo
		l.Balance = w._value.Balance
		err := w.saveWalletLog(l)
		if err == nil {
			_, err = w.Save()
		}
		return int(l.Id), err
	}
	return 0, err
}

func (w *WalletImpl) Unfreeze(value int, title, outerNo string, operatorUid int, operatorName string) error {
	err := w.checkValueOpu(value, false, operatorUid, operatorName)
	if err == nil {
		if value < 0 {
			value = -value
		}
		if w._value.FreezeAmount < value {
			return wallet.ErrOutOfAmount
		}
		w._value.Balance += int64(value)
		w._value.FreezeAmount += -value
		l := w.createWalletLog(wallet.KUnfreeze, value, title, operatorUid, operatorName)
		l.OuterNo = outerNo
		l.Balance = w._value.Balance
		err := w.saveWalletLog(l)
		if err == nil {
			_, err = w.Save()
		}
	}
	return err
}

func (w *WalletImpl) FreezeExpired(value int, remark string) error {
	if value == 0 {
		return wallet.ErrAmountZero
	}
	if value < 0 {
		value = -value
	}
	if w._value.FreezeAmount < value {
		return wallet.ErrOutOfAmount
	}
	w._value.FreezeAmount -= value
	w._value.ExpiredAmount += value
	l := w.createWalletLog(wallet.KExpired, -value, "过期失效", 0, "")
	err := w.saveWalletLog(l)
	if err == nil {
		_, err = w.Save()
	}
	return err
}

func (w *WalletImpl) CarryTo(d wallet.OperateData, freeze bool, procedureFee int) (int, error) {
	err := w.checkValueOpu(d.Amount, false, 0, "")
	if err == nil {
		if d.Amount < 0 {
			d.Amount = -d.Amount
		}
		if procedureFee < 0 {
			procedureFee = -procedureFee
		}
		k := wallet.KCarry
		if freeze {
			k = wallet.KFreeze
			w._value.FreezeAmount += d.Amount
		} else {
			w._value.Balance += int64(d.Amount)
		}
		// 保存日志
		l := w.createWalletLog(k, d.Amount, d.Title, 0, "")
		l.OuterNo = d.OuterNo
		l.ProcedureFee = -procedureFee
		l.ReviewState = wallet.ReviewPass
		l.ReviewTime = time.Now().Unix()
		l.Balance = w._value.Balance
		err = w.saveWalletLog(l)
		if err == nil {
			_, err = w.Save()
		}
		return int(l.Id), err
	}
	return 0, err
}

func (w *WalletImpl) Charge(value int, by int, title, outerNo string, remark string, operatorUid int, operatorName string) error {
	needOprUid := by == wallet.CServiceAgentCharge
	err := w.checkValueOpu(value, needOprUid, operatorUid, operatorName)
	if err == nil {
		if value < 0 {
			value = -value
		}
		var kind = by
		// 用户或客服充值、才会计入累计充值记录
		switch by {
		case wallet.CUserCharge, wallet.CServiceAgentCharge:
			kind = wallet.KCharge
			w._value.TotalCharge += int64(value)
		case wallet.CSystemCharge:
			kind = wallet.KCarry
		case wallet.CRefundCharge:
			kind = wallet.KPaymentOrderRefund
		default:
			if by < 20 {
				panic("wallet can't charge by internal defined kind")
			}
		}
		w._value.Balance += int64(value)
		// 保存日志
		l := w.createWalletLog(kind, value, title, operatorUid, operatorName)
		l.OuterNo = outerNo
		l.Remark = remark
		l.Balance = w._value.Balance
		err := w.saveWalletLog(l)
		if err == nil {
			_, err = w.Save()
		}
	}
	return err
}

func (w *WalletImpl) Refund(value int, kind int, title, outerNo string, operatorUid int, operatorName string) error {
	err := w.checkValueOpu(value, false, operatorUid, operatorName)
	if err == nil {
		if value < 0 {
			value = -value
		}
		if !(kind == wallet.KPaymentOrderRefund ||
			kind == wallet.KTransferRefund ||
			kind == wallet.KWithdrawRefund) {
			panic("not support refund kind")
		}
		switch kind {
		// 扣减总支付金额
		case wallet.KPaymentOrderRefund:
			w._value.TotalPay -= value
		}
		w._value.Balance += int64(value)
		// 保存日志
		l := w.createWalletLog(kind, value, title, operatorUid, operatorName)
		l.OuterNo = outerNo
		l.Balance = w._value.Balance
		err := w.saveWalletLog(l)
		if err == nil {
			_, err = w.Save()
		}
	}
	return err
}

func (w *WalletImpl) Transfer(toWalletId int64, value int, tradeFee int, title, toTitle, remark string) error {
	if value == 0 {
		return wallet.ErrAmountZero
	}
	if value < 0 {
		value = -value
	}
	if tradeFee < 0 {
		tradeFee = -tradeFee
	}
	var tw = w._repo.GetWallet(toWalletId)
	err := w.checkWalletState(tw, true)
	if err != nil {
		return err
	}
	// 验证金额
	if w._value.Balance < int64(value+tradeFee) {
		return wallet.ErrOutOfAmount
	}
	w._value.Balance -= int64(value + tradeFee)
	tradeNo := domain.NewTradeNo(8, int(w._value.UserId))
	l := w.createWalletLog(wallet.KTransferOut, -value, title, 0, "")
	l.ProcedureFee = -tradeFee
	l.OuterNo = tradeNo
	l.Remark = remark
	l.Balance = w._value.Balance
	err = w.saveWalletLog(l)
	if err == nil {
		if _, err = w.Save(); err == nil {
			err = tw.ReceiveTransfer(w.GetAggregateRootId(), value, tradeNo, toTitle, remark)
		}
	}
	return err
}

func (w *WalletImpl) ReceiveTransfer(fromWalletId int64, value int, tradeNo, title, remark string) error {
	if value == 0 {
		return wallet.ErrAmountZero
	}
	if value < 0 {
		value = -value
	}
	w._value.Balance += int64(value)
	l := w.createWalletLog(wallet.KTransferIn, value, title, 0, "")
	l.OuterNo = tradeNo
	l.Remark = remark
	l.Balance = w._value.Balance
	err := w.saveWalletLog(l)
	if err == nil {
		_, err = w.Save()
	}
	return err
}

// 申请提现,kind：提现方式,返回info_id,交易号 及错误,value为提现金额,tradeFee为手续费
func (w *WalletImpl) RequestWithdrawal(amount int, tradeFee int, kind int, title string,
	accountNo string, accountName string, bankName string) (int64, string, error) {
	if amount == 0 {
		return 0, "", wallet.ErrAmountZero
	}
	if amount < 0 {
		amount = -amount
	}
	if kind != wallet.KWithdrawToBankCard &&
		kind != wallet.KWithdrawToThirdPart && kind < 20 {
		return 0, "", wallet.ErrNotSupportTakeOutBusinessKind
	}
	// 判断是否暂停提现
	if wallet.TakeOutPause {
		return 0, "", wallet.ErrTakeOutPause
	}
	// 判断最低提现和最高提现
	if amount < wallet.MinTakeOutAmount {
		return 0, "", wallet.ErrLessThanMinTakeAmount
	}
	if amount > wallet.MaxTakeOutAmount {
		return 0, "", wallet.ErrMoreThanMinTakeAmount
	}
	// 余额是否不足
	if w._value.Balance < int64(amount+tradeFee) {
		return 0, "", wallet.ErrOutOfAmount
	}
	tradeNo := domain.NewTradeNo(8, int(w._value.UserId))
	w._value.Balance -= int64(amount)
	l := w.createWalletLog(kind, -(amount - tradeFee), title, 0, "")
	l.ProcedureFee = -tradeFee
	l.OuterNo = tradeNo
	l.ReviewState = wallet.ReviewAwaiting
	l.ReviewRemark = ""
	l.BankName = bankName
	l.AccountNo = accountNo
	l.AccountName = accountName
	l.Balance = w._value.Balance
	err := w.saveWalletLog(l)
	if err == nil {
		_, err = w.Save()
	}
	return l.Id, l.OuterNo, err
}

func (w *WalletImpl) ReviewWithdrawal(takeId int64, pass bool, remark string, operatorUid int, operatorName string) error {
	if err := w.checkValueOpu(1, true, operatorUid, operatorName); err != nil {
		return err
	}
	l := w.getLog(takeId)
	if l == nil {
		return wallet.ErrNoSuchTakeOutLog
	}
	if l.ReviewState != wallet.ReviewAwaiting {
		return wallet.ErrWithdrawState
	}
	l.ReviewTime = time.Now().Unix()
	if pass {
		l.ReviewState = wallet.ReviewPass
	} else {
		l.ReviewRemark = remark
		l.ReviewState = wallet.ReviewReject
		err := w.Refund(-(l.ProcedureFee + int(l.Value)), wallet.KWithdrawRefund, "提现退回",
			l.OuterNo, 0, "")
		if err != nil {
			return err
		}
	}
	l.OperatorUid = operatorUid
	l.OperatorName = operatorName
	l.UpdateTime = time.Now().Unix()
	return w.saveWalletLog(l)
}

func (w *WalletImpl) FinishWithdrawal(takeId int64, outerNo string) error {
	l := w.getLog(takeId)
	if l == nil {
		return wallet.ErrNoSuchTakeOutLog
	}
	if l.ReviewState != wallet.ReviewPass {
		return wallet.ErrWithdrawState
	}
	l.OuterNo = outerNo
	l.ReviewState = wallet.ReviewConfirm
	l.Remark = "转款凭证:" + outerNo
	return w.saveWalletLog(l)
}

func (w *WalletImpl) PagingLog(begin int, over int, opt map[string]string, sort string) (int, []*wallet.WalletLog) {
	where := bytes.NewBuffer(nil)
	// 添加业务类型筛选
	if kind, ok := opt["kind"]; ok {
		where.WriteString(" AND kind IN (")
		where.WriteString(kind)
		where.WriteString(")")
	}
	// 添加审核状态条件
	if reviewState, ok := opt["review_state"]; ok {
		where.WriteString(" AND review_state=")
		where.WriteString(reviewState)
	}
	// 添加金额
	minAmount, ok1 := opt["min_amount"]
	maxAmount, ok2 := opt["max_amount"]
	if ok1 && ok2 {
		where.WriteString(" AND value BETWEEN ")
		where.WriteString(minAmount)
		where.WriteString(" AND ")
		where.WriteString(maxAmount)
	} else {
		if ok1 {
			where.WriteString(" AND value >= ")
			where.WriteString(minAmount)
		} else {
			where.WriteString(" AND value <= ")
			where.WriteString(minAmount)
		}
	}
	// 添加时间
	beginTime, ok1 := opt["begin_time"]
	overTime, ok2 := opt["over_time"]
	if ok1 && ok2 {
		where.WriteString(" AND create_time BETWEEN ")
		where.WriteString(beginTime)
		where.WriteString(" AND ")
		where.WriteString(overTime)
	} else {
		if ok1 {
			where.WriteString(" AND create_time >= ")
			where.WriteString(beginTime)
		} else {
			where.WriteString(" AND create_time <= ")
			where.WriteString(overTime)
		}
	}
	// 添加操作人员筛选
	if operatorName, ok := opt["op_name"]; ok {
		where.WriteString(" AND op_name='")
		where.WriteString(operatorName)
		where.WriteString("'")
	}
	if operatorUid, ok := opt["op_uid"]; ok {
		where.WriteString(" AND op_uid='")
		where.WriteString(operatorUid)
	}
	return w._repo.PagingWalletLog(w.GetAggregateRootId(), w.NodeId(),
		begin, over, where.String(), sort)
}
