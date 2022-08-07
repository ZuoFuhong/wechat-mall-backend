package entity

import "time"

// WechatMallOrderDO 商城订单表
type WechatMallOrderDO struct {
	ID              int       `gorm:"column:id;primary_key;AUTO_INCREMENT"`                     // 主键
	OrderNo         string    `gorm:"column:order_no;NOT NULL"`                                 // 订单号
	UserID          int       `gorm:"column:user_id;default:0;NOT NULL"`                        // 用户ID
	PayAmount       string    `gorm:"column:pay_amount;default:0.00;NOT NULL"`                  // 订单金额（商品金额 + 运费 - 优惠金额）
	GoodsAmount     string    `gorm:"column:goods_amount;default:0.00;NOT NULL"`                // 商品小计金额
	DiscountAmount  string    `gorm:"column:discount_amount;default:0.00;NOT NULL"`             // 优惠金额
	DispatchAmount  string    `gorm:"column:dispatch_amount;default:0.00;NOT NULL"`             // 运费
	PayTime         time.Time `gorm:"column:pay_time;default:2006-01-02 15:04:05;NOT NULL"`     // 支付时间
	DeliverTime     time.Time `gorm:"column:deliver_time;default:2006-01-02 15:04:05;NOT NULL"` // 发货时间
	FinishTime      time.Time `gorm:"column:finish_time;default:2006-01-02 15:04:05;NOT NULL"`  // 成交时间
	Status          int       `gorm:"column:status;default:0;NOT NULL"`                         // 状态 -1 已取消 0-待付款 1-待发货 2-待收货 3-已完成 4-（待发货）退款申请 5-已退款
	AddressID       int       `gorm:"column:address_id;default:0;NOT NULL"`                     // 收货地址ID
	AddressSnapshot string    `gorm:"column:address_snapshot;NOT NULL"`                         // 收货地址快照
	WxappPrepayID   string    `gorm:"column:wxapp_prepay_id;NOT NULL"`                          // 微信预支付ID
	TransactionID   string    `gorm:"column:transaction_id;NOT NULL"`                           // 微信支付单号
	Remark          string    `gorm:"column:remark;NOT NULL"`                                   // 订单备注
	Del             int       `gorm:"column:is_del;default:0;NOT NULL"`                         // 是否删除：0-否 1-是
	CreateTime      time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL"`    // 创建时间
	UpdateTime      time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;NOT NULL"`    // 修改时间
}

func (m *WechatMallOrderDO) TableName() string {
	return "wechat_mall_order"
}

// OrderSaleData 报表数据-订单统计
type OrderSaleData struct {
	Time       string `gorm:"column:order_time" json:"time"`        // 下单时间
	OrderNum   int    `gorm:"column:order_num" json:"orderNum"`     // 订单数
	SaleAmount string `gorm:"column:sale_amount" json:"saleAmount"` // 销售额
}
