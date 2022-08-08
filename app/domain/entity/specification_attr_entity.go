package entity

import "time"

// WechatMallSpecificationAttrDO 商城-规格属性表
type WechatMallSpecificationAttrDO struct {
	ID         int       `gorm:"column:id;primary_key;AUTO_INCREMENT"`                  // 主键
	SpecID     int       `gorm:"column:spec_id;default:0;NOT NULL"`                     // 规格ID
	Value      string    `gorm:"column:value;NOT NULL"`                                 // 属性值
	Extend     string    `gorm:"column:extend;NOT NULL"`                                // 扩展
	Del        int       `gorm:"column:is_del;default:0;NOT NULL"`                      // 是否删除：0-否 1-是
	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 更新时间
}

func (m *WechatMallSpecificationAttrDO) TableName() string {
	return "wechat_mall_specification_attr"
}
