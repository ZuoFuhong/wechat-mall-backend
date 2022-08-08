package entity

import "time"

// WechatMallOrderRefund 商城-订单退款申请表
type WechatMallOrderRefund struct {
	ID           int       `gorm:"column:id;primary_key;AUTO_INCREMENT"`                    // 主键
	RefundNo     string    `gorm:"column:refund_no;NOT NULL"`                               // 退款编号
	UserID       int       `gorm:"column:user_id;default:0;NOT NULL"`                       // 平台用户ID
	OrderNo      string    `gorm:"column:order_no;NOT NULL"`                                // 订单号
	Reason       string    `gorm:"column:reason;NOT NULL"`                                  // 退款原因
	RefundAmount string    `gorm:"column:refund_amount;default:0.00;NOT NULL"`              // 退款金额
	Status       int       `gorm:"column:status;default:0;NOT NULL"`                        // 状态：0-退款申请 1-退款完成 2-撤销申请
	Del          int       `gorm:"column:is_del;default:0;NOT NULL"`                        // 是否删除：0-否 1-是
	RefundTime   time.Time `gorm:"column:refund_time;default:2006-01-02 15:04:05;NOT NULL"` // 退款时间
	CreateTime   time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL"`   // 创建时间
	UpdateTime   time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;NOT NULL"`   // 修改时间
}

func (m *WechatMallOrderRefund) TableName() string {
	return "wechat_mall_order_refund"
}
