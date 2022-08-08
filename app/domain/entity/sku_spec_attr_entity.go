package entity

import "time"

// WechatMallSkuSpecAttrDO 商城-SKU关联的规格属性
type WechatMallSkuSpecAttrDO struct {
	ID         int       `gorm:"column:id;primary_key;AUTO_INCREMENT"`                  // 主键
	SkuID      int       `gorm:"column:sku_id;default:0;NOT NULL"`                      // sku表主键
	SpecID     int       `gorm:"column:spec_id;default:0;NOT NULL"`                     // 规格ID
	AttrID     int       `gorm:"column:attr_id;default:0;NOT NULL"`                     // 规格-属性ID
	Del        int       `gorm:"column:is_del;default:0;NOT NULL"`                      // 是否删除：0-否 1-是
	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 更新时间
}

func (m *WechatMallSkuSpecAttrDO) TableName() string {
	return "wechat_mall_sku_spec_attr"
}
