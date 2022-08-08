package entity

import "time"

// WechatMallModuleDO CMS-组成模块
type WechatMallModuleDO struct {
	ID          int       `gorm:"column:id;primary_key;AUTO_INCREMENT"`                  // 主键
	Name        string    `gorm:"column:name;NOT NULL"`                                  // 模块名称
	Description string    `gorm:"column:description;NOT NULL"`                           // 描述
	Del         int       `gorm:"column:is_del;default:0;NOT NULL"`                      // 是否删除：0-否 1-是
	CreateTime  time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 创建时间
	UpdateTime  time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 修改时间
}

func (m *WechatMallModuleDO) TableName() string {
	return "wechat_mall_module"
}
