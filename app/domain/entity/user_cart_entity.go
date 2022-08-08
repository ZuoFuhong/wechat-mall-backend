package entity

import "time"

// WechatMallUserCartDO 商城-购物车表
type WechatMallUserCartDO struct {
	ID         int       `gorm:"column:id;primary_key;AUTO_INCREMENT"`                  // 主键
	UserID     int       `gorm:"column:user_id;default:0;NOT NULL"`                     // 用户ID
	GoodsID    int       `gorm:"column:goods_id;default:0;NOT NULL"`                    // 商品ID
	SkuID      int       `gorm:"column:sku_id;default:0;NOT NULL"`                      // sku ID
	Num        int       `gorm:"column:num;default:0;NOT NULL"`                         // 数量
	Del        int       `gorm:"column:is_del;default:0;NOT NULL"`                      // 是否删除：0-否 1-是
	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 修改时间
}

func (m *WechatMallUserCartDO) TableName() string {
	return "wechat_mall_user_cart"
}

type PortalCartGoods struct {
	Id      int     `json:"id"`      // 购物车ID
	GoodsId int     `json:"goodsId"` // 商品ID
	SkuId   int     `json:"skuId"`   // SKU ID
	Title   string  `json:"title"`   // 标题
	Price   float64 `json:"price"`   // SKU价格
	Picture string  `json:"picture"` // SKU图片
	Specs   string  `json:"specs"`   // specs值
	Num     int     `json:"num"`     // 数量
	Status  int     `json:"status"`  // 库存状态：0-正常 1-缺货 2-下架
}

type CartGoods struct {
	Num     int `json:"num"`     // 数量
	CartId  int `json:"cartId"`  // 购物车ID
	GoodsId int `json:"goodsId"` // 商品ID
	SkuId   int `json:"skuId"`   // skuID
}
