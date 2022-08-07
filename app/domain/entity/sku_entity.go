package entity

import "time"

// WechatMallSkuDO 商城-SKU表
type WechatMallSkuDO struct {
	ID         int       `gorm:"column:id;primary_key;AUTO_INCREMENT"`                  // 主键
	Title      string    `gorm:"column:title;NOT NULL"`                                 // 标题
	Price      string    `gorm:"column:price;default:0.00;NOT NULL"`                    // 价格
	Code       string    `gorm:"column:code;NOT NULL"`                                  // 编码
	Stock      int       `gorm:"column:stock;default:0;NOT NULL"`                       // 库存量
	GoodsID    int       `gorm:"column:goods_id;default:0;NOT NULL"`                    // 所属商品
	Online     int       `gorm:"column:online;default:0;NOT NULL"`                      // 是否上架: 0-下架 1-上架
	Picture    string    `gorm:"column:picture;NOT NULL"`                               // 图片
	Specs      string    `gorm:"column:specs;NOT NULL"`                                 // 规格属性
	Del        int       `gorm:"column:is_del;default:0;NOT NULL"`                      // 是否删除：0-否 1-是
	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 更新时间
}

func (m *WechatMallSkuDO) TableName() string {
	return "wechat_mall_sku"
}

type SkuSpecs struct {
	Key     string `json:"key"`
	Value   string `json:"value"`
	KeyId   int    `json:"keyId"`
	ValueId int    `json:"valueId"`
}
