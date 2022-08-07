package entity

import "time"

// WechatMallOrderGoodsDO 商城订单-商品表
type WechatMallOrderGoodsDO struct {
	ID         int       `gorm:"column:id;primary_key;AUTO_INCREMENT"`                  // 主键
	OrderNo    string    `gorm:"column:order_no;NOT NULL"`                              // 订单号
	UserID     int       `gorm:"column:user_id;default:0;NOT NULL"`                     // 用户ID
	GoodsID    int       `gorm:"column:goods_id;default:0;NOT NULL"`                    // 商品ID
	SkuID      int       `gorm:"column:sku_id;default:0;NOT NULL"`                      // sku ID
	Picture    string    `gorm:"column:picture;NOT NULL"`                               // 商品图片
	Title      string    `gorm:"column:title;NOT NULL"`                                 // 商品标题
	Price      string    `gorm:"column:price;default:0.00;NOT NULL"`                    // 价格
	Specs      string    `gorm:"column:specs;NOT NULL"`                                 // sku规格属性
	Num        int       `gorm:"column:num;default:0;NOT NULL"`                         // 数量
	LockStatus int       `gorm:"column:lock_status;default:0;NOT NULL"`                 // 锁定状态：0-锁定 1-解锁
	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 修改时间
}

func (m *WechatMallOrderGoodsDO) TableName() string {
	return "wechat_mall_order_goods"
}
