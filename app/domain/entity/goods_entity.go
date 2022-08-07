package entity

import "time"

// WechatMallGoodsDO 商城-商品表
type WechatMallGoodsDO struct {
	ID            int       `gorm:"column:id;primary_key;AUTO_INCREMENT"`                  // 主键
	BrandName     string    `gorm:"column:brand_name;NOT NULL"`                            // 品牌名称
	Title         string    `gorm:"column:title;NOT NULL"`                                 // 标题
	Price         string    `gorm:"column:price;default:0.00;NOT NULL"`                    // 价格
	DiscountPrice string    `gorm:"column:discount_price;default:0.00;NOT NULL"`           // 折扣
	CategoryID    int       `gorm:"column:category_id;default:0;NOT NULL"`                 // 分类ID
	Online        int       `gorm:"column:online;default:0;NOT NULL"`                      // 是否上架：0-下架 1-上架
	Picture       string    `gorm:"column:picture;NOT NULL"`                               // 主图
	BannerPicture string    `gorm:"column:banner_picture"`                                 // 轮播图
	DetailPicture string    `gorm:"column:detail_picture"`                                 // 详情图
	Tags          string    `gorm:"column:tags;NOT NULL"`                                  // 标签，示例：包邮$热门
	SaleNum       int       `gorm:"column:sale_num;default:0;NOT NULL"`                    // 商品销量
	Del           int       `gorm:"column:is_del;default:0;NOT NULL"`                      // 是否删除：0-否 1-是
	CreateTime    time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 创建时间
	UpdateTime    time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 更新时间
}

func (m *WechatMallGoodsDO) TableName() string {
	return "wechat_mall_goods"
}
