package entity

import "time"

// WechatMallGroupPagePermission CMS-用户分组-页面权限表
type WechatMallGroupPagePermission struct {
	ID         int       `gorm:"column:id;primary_key;AUTO_INCREMENT"`                  // 主键
	GroupID    int       `gorm:"column:group_id;default:0;NOT NULL"`                    // 分组ID
	PageID     int       `gorm:"column:page_id;default:0;NOT NULL"`                     // 页面ID
	Del        int       `gorm:"column:is_del;default:0;NOT NULL"`                      // 是否删除：0-否 1-是
	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 修改时间
}

func (m *WechatMallGroupPagePermission) TableName() string {
	return "wechat_mall_group_page_permission"
}
