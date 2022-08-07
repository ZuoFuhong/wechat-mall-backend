package entity

import "time"

// WechatMallGoodsBrowseRecord 商城-商品浏览记录
type WechatMallGoodsBrowseRecord struct {
	ID         int       `gorm:"column:id;primary_key;AUTO_INCREMENT"`                  // 主键
	UserID     int       `gorm:"column:user_id;default:0;NOT NULL"`                     // 用户ID
	GoodsID    int       `gorm:"column:goods_id;default:0;NOT NULL"`                    // 商品ID
	Picture    string    `gorm:"column:picture;NOT NULL"`                               // 商品图片
	Title      string    `gorm:"column:title;NOT NULL"`                                 // 商品名称
	Price      string    `gorm:"column:price;default:0.00;NOT NULL"`                    // 商品价格
	Del        int       `gorm:"column:is_del;default:0;NOT NULL"`                      // 是否删除：0-否 1-是
	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 修改时间
}

func (m *WechatMallGoodsBrowseRecord) TableName() string {
	return "wechat_mall_goods_browse_record"
}
