package entity

import "time"

// WechatMallBannerDO 小程序Banner表
type WechatMallBannerDO struct {
	ID           int       `gorm:"column:id;primary_key;AUTO_INCREMENT"`                  // 主键
	Picture      string    `gorm:"column:picture;NOT NULL"`                               // 图片地址
	Name         string    `gorm:"column:name;NOT NULL"`                                  // 名称
	BusinessType int       `gorm:"column:business_type;default:0;NOT NULL"`               // 业务类型：1-商品
	BusinessID   int       `gorm:"column:business_id;default:0;NOT NULL"`                 // 业务主键
	Status       int       `gorm:"column:status;default:0;NOT NULL"`                      // 是否显示：0-否 1-是
	Del          int       `gorm:"column:is_del;default:0;NOT NULL"`                      // 是否删除：0-否 1-是
	CreateTime   time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 创建时间
	UpdateTime   time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 更新时间
}

func (m *WechatMallBannerDO) TableName() string {
	return "wechat_mall_banner"
}
