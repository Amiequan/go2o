package order

import (
	"github.com/ixre/go2o/core/domain/interface/cart"
	"github.com/ixre/go2o/core/domain/interface/promotion"
)

type (
	// INormalOrder 普通订单
	INormalOrder interface {
		// RequireCart 读取购物车数据,用于预生成订单
		RequireCart(c cart.ICart) error
		// GetByVendor 根据运营商获取商品和运费信息,限未生成的订单
		GetByVendor() (items map[int][]*SubOrderItem, expressFee map[int]int64)
		// OnlinePaymentTradeFinish 在线支付交易完成
		OnlinePaymentTradeFinish() error
		// Submit 提交订单。如遇拆单,需均摊优惠抵扣金额到商品
		Submit() error

		//根据运营商拆单,返回拆单结果,及拆分的订单数组
		//BreakUpByVendor() ([]IOrder, error)

		// GetSubOrders 获取子订单列表
		GetSubOrders() []ISubOrder
		// ApplyCoupon 应用优惠券
		ApplyCoupon(coupon promotion.ICouponPromotion) error
		// GetCoupons 获取应用的优惠券
		GetCoupons() []promotion.ICouponPromotion
		// GetAvailableOrderPromotions 获取可用的促销,不包含优惠券
		GetAvailableOrderPromotions() []promotion.IPromotion
		// GetBestSavePromotion 获取最省的促销
		GetBestSavePromotion() (p promotion.IPromotion,
			saveFee float32, integral int)
		// GetPromotionBinds 获取促销绑定
		GetPromotionBinds() []*OrderPromotionBind
	}

	// ISubOrder 子订单(普通订单拆分)
	ISubOrder interface {
		// GetDomainId 获取领域对象编号
		GetDomainId() int64
		// GetValue 获取值对象
		GetValue() *NormalSubOrder
		// Complex 复合的订单信息
		Complex() *ComplexOrder
		// Items 获取商品项
		Items() []*SubOrderItem
		// PaymentFinishByOnlineTrade 在线支付交易完成
		PaymentFinishByOnlineTrade() error
		// AppendLog 记录订单日志
		AppendLog(logType LogType, system bool, message string) error
		// AddRemark 添加备注
		AddRemark(string)
		// Confirm 确认订单
		Confirm() error
		// PickUp 捡货(备货)
		PickUp() error
		// Ship 发货
		Ship(spId int32, spOrder string) error
		// BuyerReceived 已收货
		BuyerReceived() error
		// LogBytes 获取订单的日志
		LogBytes() []byte
		// Suspend 挂起
		Suspend(reason string) error
		// Cancel 取消订单/退款
		Cancel(reason string) error
		// Return 退回商品
		Return(snapshotId int64, quantity int32) error
		// RevertReturn 撤销退回商品
		RevertReturn(snapshotId int64, quantity int32) error
		// Decline 谢绝订单
		Decline(reason string) error
		// Submit 提交子订单
		Submit() (int64, error)
		// Destory 销毁订单
		Destory() error
	}

	// 子订单
	NormalSubOrder struct {
		// 编号
		Id int64 `db:"id" pk:"yes" auto:"yes"`
		// 订单号
		OrderNo string `db:"order_no"`
		// 订单编号
		OrderId int64 `db:"order_id"`
		// 购买人编号(冗余,便于商户处理数据)
		BuyerId int64 `db:"buyer_id"`
		// 运营商编号
		VendorId int64 `db:"vendor_id"`
		// 店铺编号
		ShopId int64 `db:"shop_id"`
		// 店铺名称
		ShopName string `db:"shop_name"`
		// 订单标题
		Subject string `db:"subject"`
		// 商品金额
		ItemAmount int64 `db:"item_amount"`
		// 优惠减免金额
		DiscountAmount int64 `db:"discount_amount"`
		// 运费
		ExpressFee int64 `db:"express_fee"`
		// 包装费用
		PackageFee int64 `db:"package_fee"`
		// 实际金额
		FinalAmount int64 `db:"final_amount"`
		// 是否挂起，如遇到无法自动进行的时挂起，来提示人工确认。
		IsSuspend int `db:"is_suspend"`
		// 顾客备注
		BuyerComment string `db:"buyer_comment"`
		// 系统备注
		Remark string `db:"remark" json:"remark"`
		// 订单状态
		Status int `db:"status" json:"status"`
		// 拆分状态: 0.默认 1:待拆分 2:无需拆分 3:已拆分
		BreakStatus int `db:"break_status"`
		// 下单时间
		CreateTime int64 `db:"create_time"`
		// 更新时间
		UpdateTime int64 `db:"update_time"`
		// 订单项
		Items []*SubOrderItem `db:"-"`
	}

	// SubOrderItem 订单商品项
	SubOrderItem struct {
		// 编号
		ID int64 `db:"id" pk:"yes" auto:"yes" json:"id"`
		// 订单编号
		OrderId int64 `db:"order_id"`
		// 卖家订单编号
		SellerOrderId int64 `db:"seller_order_id"`
		// 商品编号
		ItemId int64 `db:"item_id"`
		// 商品SKU编号
		SkuId int64 `db:"sku_id"`
		// 快照编号
		SnapshotId int64 `db:"snap_id"`
		// 数量
		Quantity int32 `db:"quantity"`
		// 退回数量(退货)
		ReturnQuantity int32 `db:"return_quantity"`
		// 金额
		Amount int64 `db:"amount"`
		// 最终金额, 可能会有优惠均摊抵扣的金额
		FinalAmount int64 `db:"final_amount"`
		// 是否发货
		IsShipped int32 `db:"is_shipped"`
		// 更新时间
		UpdateTime int64 `db:"update_time"`
		// 运营商编号
		VendorId int64 `db:"-"`
		// 商店编号
		ShopId int64 `db:"-"`
		// 重量,用于生成订单时存储数据
		Weight int32 `db:"-"`
		// 体积:毫升(ml)
		Bulk int32 `db:"-"`
		// 快递模板编号
		ExpressTplId int32 `db:"-"`
	}
)
