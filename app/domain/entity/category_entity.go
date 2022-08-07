package entity

import "time"

// WechatMallCategoryDO 商城-分类表
type WechatMallCategoryDO struct {
	ID          int       `gorm:"column:id;primary_key;AUTO_INCREMENT"`                  // 主键
	ParentID    int       `gorm:"column:parent_id;default:0;NOT NULL"`                   // 父级分类ID
	Name        string    `gorm:"column:name;NOT NULL"`                                  // 分类名称
	Sort        int       `gorm:"column:sort;default:0;NOT NULL"`                        // 排序
	Online      int       `gorm:"column:online;default:0;NOT NULL"`                      // 是否上线：0-否 1-是
	Picture     string    `gorm:"column:picture;NOT NULL"`                               // 图片地址
	Description string    `gorm:"column:description;NOT NULL"`                           // 分类描述
	Del         int       `gorm:"column:is_del;default:0;NOT NULL"`                      // 是否删除：0-否 1-是
	CreateTime  time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 创建时间
	UpdateTime  time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 更新时间
}

func (m *WechatMallCategoryDO) TableName() string {
	return "wechat_mall_category"
}
