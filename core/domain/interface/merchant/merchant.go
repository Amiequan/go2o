/**
 * Copyright 2014 @ 56x.net.
 * name :
 * author : jarryliu
 * date : 2013-12-12 16:53
 * description :
 * history :
 */

package merchant

import (
	"github.com/ixre/go2o/core/domain/interface/merchant/shop"
	"github.com/ixre/go2o/core/domain/interface/merchant/user"
	"github.com/ixre/go2o/core/domain/interface/merchant/wholesaler"
)

const (
	KindAccountCharge           = 1
	KindAccountSettleOrder      = 2
	KindAccountPresent          = 3
	KindAccountTakePayment      = 4
	KindAccountTransferToMember = 5

	// KindＭachTakeOutToBankCard 商户提现
	KindＭachTakeOutToBankCard = 100
	// KindＭachTakOutRefund 商户提现失败返还给会员
	KindＭachTakOutRefund = 101
)

type (
	// IMerchant 商户接口
	//todo: 实现商户等级,商户的品牌
	IMerchant interface {
		// GetAggregateRootId 获取编号
		GetAggregateRootId() int64
		GetValue() Merchant
		// Complex 获取符合的商家信息
		Complex() *ComplexMerchant
		// SetValue 设置值
		SetValue(*Merchant) error
		// Stat 获取商户的状态,判断 过期时间、判断是否停用。
		// 过期时间通常按: 试合作期,即1个月, 后面每年延长一次。或者会员付费开通。
		Stat() error
		// SetEnabled 设置商户启用或停用
		SetEnabled(enabled bool) error
		// SelfSales 是否自营
		SelfSales() bool
		// Member 返回对应的会员编号
		Member() int64
		// Save 保存
		Save() (int64, error)
		// GetMajorHost 获取商户的域名
		GetMajorHost() string
		// BindMember 绑定会员号
		BindMember(memberId int) error
		// Account 获取商户账户
		Account() IAccount
		// EnableWholesale 启用批发
		EnableWholesale() error
		// Wholesaler 获取批发商实例
		Wholesaler() wholesaler.IWholesaler
		// UserManager 返回用户服务
		UserManager() user.IUserManager
		// ConfManager 返回设置服务
		ConfManager() IConfManager
		// SaleManager 销售服务
		SaleManager() ISaleManager
		// LevelManager 获取会员等级服务
		LevelManager() ILevelManager
		// KvManager 获取键值管理器
		KvManager() IKvManager
		// ProfileManager 企业资料服务
		ProfileManager() IProfileManager
		// ApiManager API服务
		ApiManager() IApiManager
		// ShopManager 商店服务
		ShopManager() shop.IShopManager
		// MemberKvManager 获取会员键值管理器
		MemberKvManager() IKvManager
		// 消息系统管理器
		//MssManager() mss.IMssManager
	}

	// IAccount 账户
	IAccount interface {
		// GetDomainId 获取领域对象编号
		GetDomainId() int64
		// GetValue 获取账户值
		GetValue() *Account
		// Save 保存
		Save() error
		// GetBalanceLog 根据编号获取余额变动信息
		GetBalanceLog(id int32) *BalanceLog
		// GetBalanceLogByOuterNo 根据号码获取余额变动信息
		GetBalanceLogByOuterNo(outerNo string) *BalanceLog
		// SaveBalanceLog 保存余额变动信息
		SaveBalanceLog(*BalanceLog) (int32, error)
		// SettleOrder 订单结账
		SettleOrder(orderNo string, amount int, tradeFee int, refundAmount int, remark string) error
		// TakePayment 支出
		TakePayment(outerNo string, amount int, csn int, remark string) error

		// 提现
		//todo:???

		// TransferToMember todo: 以下需要重构或移除
		// 转到会员账户
		TransferToMember(amount int) error

		// TransferToMember1 商户积分转会员积分
		TransferToMember1(amount float32) error

		// Present 赠送
		Present(amount int, remark string) error

		// Charge 充值
		Charge(kind int32, amount int, title, outerNo string,
			relateUser int64) error
	}

	IMerchantManager interface {
		// CreateSignUpToken 创建会员申请商户密钥
		CreateSignUpToken(memberId int64) string

		// GetMemberFromSignUpToken 根据商户申请密钥获取会员编号
		GetMemberFromSignUpToken(token string) int64

		// CommitSignUpInfo 提交商户注册信息
		CommitSignUpInfo(*MchSignUp) (int32, error)

		// ReviewMchSignUp 审核商户注册信息
		ReviewMchSignUp(id int32, pass bool, remark string) error

		// GetSignUpInfo 获取商户申请信息
		GetSignUpInfo(id int32) *MchSignUp

		// GetSignUpInfoByMemberId 获取会员申请的商户信息
		GetSignUpInfoByMemberId(memberId int64) *MchSignUp

		// GetMerchantByMemberId 获取会员关联的商户
		GetMerchantByMemberId(memberId int64) IMerchant

		// RemoveSignUp 删除会员的商户申请资料
		RemoveSignUp(memberId int64) error
	}

	// MchSignUp 商户申请信息
	MchSignUp struct {
		Id int32 `db:"id" pk:"yes" auth:"yes"`
		// 申请单号
		SignNo string `db:"sign_no"`
		// 会员编号
		MemberId int64 `db:"member_id"`
		// 用户名
		User string `db:"user"`
		// 密码
		Pwd string `db:"pwd"`
		// 盐
		Salt string `db:"salt"`
		// 商户名称号
		MchName string `db:"mch_name"`
		// 省
		Province int32 `db:"province"`
		// 市
		City int32 `db:"city"`
		// 区
		District int32 `db:"district"`
		// 详细地址
		Address string `db:"address"`
		// 店铺店铺
		ShopName string `db:"shop_name"`
		// 公司名称
		CompanyName string `db:"company_name"`
		// 营业执照编号
		CompanyNo string `db:"company_no"`
		// 法人
		PersonName string `db:"person_name"`
		// 法人身份证
		PersonId string `db:"person_id"`
		// 法人身份证
		PersonImage string `db:"person_image"`
		// 联系电话
		Phone string `db:"phone"`
		// 营业执照图片
		CompanyImage string `db:"company_image"`
		// 委托书
		AuthDoc string `db:"auth_doc"`
		// 备注
		Remark string `db:"remark"`
		// 提交时间
		SubmitTime int64 `db:"submit_time"`
		// 是否通过
		Reviewed int32 `db:"review_state"`
		// 更新时间
		UpdateTime int64 `db:"update_time"`
	}
	// 商户
	ComplexMerchant struct {
		Id int32
		// 关联的会员编号,作为结算账户
		MemberId int64
		// 用户
		Usr string
		// 密码
		Pwd string
		// 商户名称
		Name string
		// 是否自营
		SelfSales int32
		// 商户等级
		Level int32
		// 标志
		Logo string
		// 公司名称
		CompanyName string
		// 省
		Province int32
		// 市
		City int32
		// 区
		District int32
		// 是否启用
		Enabled int32
		// 过期时间
		ExpiresTime int64
		// 注册时间
		JoinTime int64
		// 更新时间
		UpdateTime int64
		// 登录时间
		LoginTime int64
		// 最后登录时间
		LastLoginTime int64
	}

	// 商户
	Merchant struct {
		Id int64 `db:"id" pk:"yes" auto:"yes"`
		// 会员编号
		MemberId int64 `db:"member_id"`
		// 登录用户
		LoginUser string `db:"login_user"`
		// 登录密码
		LoginPwd string `db:"login_pwd"`
		// 加密盐
		Salt string `db:"salt"`
		// 名称
		Name string `db:"name"`
		// 公司名称
		CompanyName string `db:"company_name"`
		// 是否自营
		SelfSales int16 `db:"self_sales"`
		// 商户等级
		Level int `db:"level"`
		// 标志
		Logo string `db:"logo"`
		// 省
		Province int `db:"province"`
		// 市
		City int `db:"city"`
		// 区
		District int `db:"district"`
		// 标志
		Flag int `db:"flag"`
		// 是否启用
		Enabled int `db:"enabled"`
		// 创建时间
		CreateTime int64 `db:"create_time"`
		// 过期时间
		ExpiresTime int64 `db:"expires_time"`
		// 更新时间
		UpdateTime int64 `db:"update_time"`
		// 登录时间
		LoginTime int64 `db:"login_time"`
		// 最后登录时间
		LastLoginTime int64 `db:"last_login_time"`
	}

	// 商户账户表
	Account struct {
		// 商户编号
		MchId int64 `db:"mch_id" pk:"yes"`
		// 余额
		Balance int64 `db:"balance"`
		// 冻结金额
		FreezeAmount int64 `db:"freeze_amount"`
		// 待入账金额
		AwaitAmount int64 `db:"await_amount"`
		// 平台赠送金额
		PresentAmount int64 `db:"present_amount"`
		// 累计销售总额
		SalesAmount int64 `db:"sales_amount"`
		// 累计退款金额
		RefundAmount int64 `db:"refund_amount"`
		// 已提取金额
		WithdrawAmount int64 `db:"take_amount"`
		// 线下销售金额
		OfflineSales int64 `db:"offline_sales"`
		// 更新时间
		UpdateTime int64 `db:"update_time"`
	}

	// 商户余额日志
	BalanceLog struct {
		// 编号
		Id int32 `db:"id" pk:"yes" auto:"yes"`
		// 商户编号
		MchId int64 `db:"mch_id"`
		// 日志类型
		Kind int `db:"kind"`
		// 标题
		Title string `db:"title"`
		// 外部订单号
		OuterNo string `db:"outer_no"`
		// 金额
		Amount int64 `db:"amount"`
		// 手续费
		CsnAmount int64 `db:"csn_amount"`
		// 状态
		State int `db:"state"`
		// 创建时间
		CreateTime int64 `db:"create_time"`
		// 更新时间
		UpdateTime int64 `db:"update_time"`
	}

	// 商户每日报表
	MchDayChart struct {
		// 编号
		Id int32 `db:"id" pk:"yes" auto:"yes"`
		// 商户编号
		MchId int64 `db:"mch_id"`
		// 新增订单数量
		OrderNumber int `db:"order_number"`
		// 订单额
		OrderAmount float32 `db:"order_amount"`
		// 购物会员数
		BuyerNumber int `db:"buyer_number"`
		// 支付单数量
		PaidNumber int `db:"paid_number"`
		// 支付总金额
		PaidAmount float32 `db:"paid_amount"`
		// 完成订单数
		CompleteOrders int `db:"complete_orders"`
		// 入帐金额
		InAmount float32 `db:"in_amount"`
		// 线下订单数量
		OfflineOrders int `db:"offline_orders"`
		// 线下订单金额
		OfflineAmount float32 `db:"offline_amount"`
		// 日期
		Date int64 `db:"date"`
		// 日期字符串
		DateStr string `db:"date_str"`
		// 更新时间
		UpdateTime int64 `db:"update_time"`
	}
)
