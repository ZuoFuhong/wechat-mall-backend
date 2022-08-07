package entity

import "time"

// WechatMallCouponLogDO 商城-优惠券领取记录表
type WechatMallCouponLogDO struct {
	ID         int       `gorm:"column:id;primary_key;AUTO_INCREMENT"`                  // 主键
	CouponID   int       `gorm:"column:coupon_id;default:0;NOT NULL"`                   // 优惠券ID
	UserID     int       `gorm:"column:user_id;default:0;NOT NULL"`                     // 用户ID
	UseTime    time.Time `gorm:"column:use_time"`                                       // 使用时间
	ExpireTime time.Time `gorm:"column:expire_time;NOT NULL"`                           // 过期时间
	Status     int       `gorm:"column:status;default:0;NOT NULL"`                      // 状态：0-未使用 1-已使用 2-已过期
	Code       string    `gorm:"column:code;NOT NULL"`                                  // 兑换码
	OrderNo    string    `gorm:"column:order_no;NOT NULL"`                              // 核销的订单号
	Del        int       `gorm:"column:is_del;default:0;NOT NULL"`                      // 是否删除：0-否 1-是
	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 更新时间
}

func (m *WechatMallCouponLogDO) TableName() string {
	return "wechat_mall_coupon_log"
}
