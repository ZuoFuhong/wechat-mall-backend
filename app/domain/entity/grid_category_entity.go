package entity

import "time"

// WechatMallGridCategoryDO 小程序-首页宫格表
type WechatMallGridCategoryDO struct {
	ID         int       `gorm:"column:id;primary_key;AUTO_INCREMENT"`                  // 主键
	Title      string    `gorm:"column:title;NOT NULL"`                                 // 宫格标题
	Name       string    `gorm:"column:name;NOT NULL"`                                  // 宫格名
	CategoryID int       `gorm:"column:category_id;default:0;NOT NULL"`                 // 顶级分类ID
	Picture    string    `gorm:"column:picture;NOT NULL"`                               // 图片地址
	Del        int       `gorm:"column:is_del;default:0;NOT NULL"`                      // 是否删除：0-否 1-是
	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 更新时间
}

func (m *WechatMallGridCategoryDO) TableName() string {
	return "wechat_mall_grid_category"
}
