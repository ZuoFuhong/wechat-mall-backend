package entity

import "time"

// WechatMallVisitorRecord 商城-访客记录表
type WechatMallVisitorRecord struct {
	ID         int       `gorm:"column:id;primary_key;AUTO_INCREMENT"`                  // 主键
	UserID     int       `gorm:"column:user_id;default:0;NOT NULL"`                     // 平台用户ID
	IP         string    `gorm:"column:ip;NOT NULL"`                                    // 独立IP
	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 修改时间
}

func (m *WechatMallVisitorRecord) TableName() string {
	return "wechat_mall_visitor_record"
}
