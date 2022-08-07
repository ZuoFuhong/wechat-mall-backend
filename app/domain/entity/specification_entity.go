package entity

import "time"

// WechatMallSpecificationDO 商城-规格表
type WechatMallSpecificationDO struct {
	ID          int       `gorm:"column:id;primary_key;AUTO_INCREMENT"`                  // 主键
	Name        string    `gorm:"column:name;NOT NULL"`                                  // 规格名名称
	Description string    `gorm:"column:description;NOT NULL"`                           // 规格名描述
	Unit        string    `gorm:"column:unit;NOT NULL"`                                  // 单位
	Standard    int       `gorm:"column:standard;default:0;NOT NULL"`                    // 是否标准: 0-非标准 1-标准
	Del         int       `gorm:"column:is_del;default:0;NOT NULL"`                      // 是否删除：0-否 1-是
	CreateTime  time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 创建时间
	UpdateTime  time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 更新时间
}

func (m *WechatMallSpecificationDO) TableName() string {
	return "wechat_mall_specification"
}
