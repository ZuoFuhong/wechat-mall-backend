package entity

import "time"

// WechatMallModulePageDO CMS-模块页面
type WechatMallModulePageDO struct {
	ID          int       `gorm:"column:id;primary_key;AUTO_INCREMENT"`                  // 主键
	ModuleID    int       `gorm:"column:module_id;default:0;NOT NULL"`                   // 模块ID
	Name        string    `gorm:"column:name;NOT NULL"`                                  // 页面名称
	Description string    `gorm:"column:description;NOT NULL"`                           // 描述
	Del         int       `gorm:"column:is_del;default:0;NOT NULL"`                      // 是否删除：0-否 1-是
	CreateTime  time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 创建时间
	UpdateTime  time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 修改时间
}

func (m *WechatMallModulePageDO) TableName() string {
	return "wechat_mall_module_page"
}

type ModulePageAuth struct {
	Auth   string `json:"auth"`
	Module string `json:"module"`
}
