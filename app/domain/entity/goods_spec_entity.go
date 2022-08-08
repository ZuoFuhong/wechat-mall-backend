package entity

import "time"

// WechatMallGoodsSpecDO 商城-商品规格表
type WechatMallGoodsSpecDO struct {
	ID         int       `gorm:"column:id;primary_key;AUTO_INCREMENT"`                  // 主键
	GoodsID    int       `gorm:"column:goods_id;default:0;NOT NULL"`                    // 商品ID
	SpecID     int       `gorm:"column:spec_id;default:0;NOT NULL"`                     // 规格ID
	Del        int       `gorm:"column:is_del;default:0;NOT NULL"`                      // 是否删除：0-否 1-是
	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 更新时间
}

func (m *WechatMallGoodsSpecDO) TableName() string {
	return "wechat_mall_goods_spec"
}
