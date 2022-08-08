package entity

import "time"

// WechatMallCouponDO 商城-优惠券表
type WechatMallCouponDO struct {
	ID          int       `gorm:"column:id;primary_key;AUTO_INCREMENT"`                  // 主键
	Title       string    `gorm:"column:title;NOT NULL"`                                 // 标题
	FullMoney   string    `gorm:"column:full_money;default:0.00;NOT NULL"`               // 满减额
	Minus       string    `gorm:"column:minus;default:0.00;NOT NULL"`                    // 优惠额
	Rate        string    `gorm:"column:rate;default:0.00;NOT NULL"`                     // 折扣
	Type        int       `gorm:"column:type;default:0;NOT NULL"`                        // 券类型：1-满减券 2-折扣券 3-代金券 4-满金额折扣券
	GrantNum    int       `gorm:"column:grant_num;default:0;NOT NULL"`                   // 发券数量
	LimitNum    int       `gorm:"column:limit_num;default:0;NOT NULL"`                   // 单人限领
	StartTime   time.Time `gorm:"column:start_time;NOT NULL"`                            // 开始时间
	EndTime     time.Time `gorm:"column:end_time;NOT NULL"`                              // 结束时间
	Description string    `gorm:"column:description;NOT NULL"`                           // 描述
	Online      int       `gorm:"column:online;default:0;NOT NULL"`                      // 是否上架: 0-下架 1-上架
	Del         int       `gorm:"column:is_del;default:0;NOT NULL"`                      // 是否删除：0-否 1-是
	CreateTime  time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 创建时间
	UpdateTime  time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 更新时间
}

func (m *WechatMallCouponDO) TableName() string {
	return "wechat_mall_coupon"
}
