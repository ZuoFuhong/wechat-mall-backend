package defs

type WxappLoginResp struct {
	Token               string `json:"token" validate:"required"`
	ExpirationInMinutes int    `json:"expiration_in_minutes" validate:"required"`
}

type PortalBannerVO struct {
	Id      int    `json:"id"`
	Picture string `json:"picture"`
}

type PortalGridCategoryVO struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	CategoryId int    `json:"category"`
	Picture    string `json:"picture"`
}

type PortalCouponVO struct {
	Id          int    `json:"id"`          // 优惠券ID
	Title       string `json:"title"`       // 标题
	FullMoney   string `json:"fullMoney"`   // 满减额
	Minus       string `json:"minus"`       // 优惠额
	Rate        string `json:"rate"`        // 折扣
	Type        int    `json:"type"`        // 券类型：1-满减券 2-折扣券 3-无门槛券 4-满金额折扣券
	StartTime   string `json:"startTime"`   // 开始时间
	EndTime     string `json:"endTime"`     // 结束时间
	Description string `json:"description"` // 描述
	Status      int    `json:"status"`      // 领取状态：
}

type PortalTakeCouponReq struct {
	CouponId int `json:"couponId" validate:"required"`
}

type PortalUserCouponVO struct {
	CLogId      int    `json:"CLogId"`      // 领取记录ID
	CouponId    int    `json:"couponId"`    // 优惠券ID
	Title       string `json:"title"`       // 标题
	FullMoney   string `json:"fullMoney"`   // 满减额
	Minus       string `json:"minus"`       // 优惠额
	Rate        string `json:"rate"`        // 折扣
	Type        int    `json:"type"`        // 券类型：1-满减券 2-折扣券 3-无门槛券 4-满金额折扣券
	StartTime   string `json:"startTime"`   // 开始时间
	EndTime     string `json:"endTime"`     // 结束时间
	Description string `json:"description"` // 描述
}

type PortalGoodsListVO struct {
	Id            int    `json:"id"`            // 商品ID
	Title         string `json:"title"`         // 标题
	Price         string `json:"price"`         // 价格
	DiscountPrice string `json:"discountPrice"` // 折扣
	Picture       string `json:"picture"`       // 图片
	Tags          string `json:"tags"`          // 标签
	SaleNum       int    `json:"saleNum"`       // 销量
}

type PortalGoodsInfo struct {
	Id              int            `json:"id"`              // 商品ID
	BrandName       string         `json:"brandName"`       // 品牌
	Title           string         `json:"title"`           // 标题
	Price           string         `json:"price"`           // 价格
	DiscountPrice   string         `json:"discountPrice"`   // 折扣
	Picture         string         `json:"picture"`         // 主图
	BannerPicture   string         `json:"bannerPicture"`   // 详情图
	DetailPicture   string         `json:"detailPicture"`   // 轮播图
	Tags            string         `json:"tags"`            // 标签
	Description     string         `json:"description"`     // 详情
	MultiplePicture []string       `json:"multiplePicture"` // 多图
	SkuList         []PortalSkuVO  `json:"skuList"`         // sku列表
	SpecList        []PortalSpecVO `json:"specList"`        // 规格列表
}

type PortalSpecVO struct {
	SpecId   int                `json:"specId"`   // 规格ID
	Name     string             `json:"name"`     // 规格名称
	AttrList []PortalSpecAttrVO `json:"attrList"` // 规格属性
}

type PortalSpecAttrVO struct {
	AttrId int    `json:"attrId"` // 属性ID
	Value  string `json:"value"`  // 属性名称
}

type PortalSkuVO struct {
	Id    int    `json:"id"`    // skuId
	Title string `json:"title"` // 标题
	Price string `json:"price"` // 价格
	Code  string `json:"code"`  // 编码
	Stock int    `json:"stock"` // 库存量
	Specs string `json:"specs"` // 多规格属性
}

type PortalCartGoodsReq struct {
	GoodsId int `json:"goodsId" validate:"required"` // 商品ID
	SkuId   int `json:"skuId" validate:"required"`   // skuId
	Num     int `json:"num" validate:"required"`     // 数量
}

type PortalCartGoodsVO struct {
	GoodsId       int    `json:"goodsId"`       // 商品ID
	Title         string `json:"title"`         // 标题
	Price         string `json:"price"`         // 价格
	DiscountPrice string `json:"discountPrice"` // 折扣
	Picture       string `json:"picture"`       // 图片
	Tags          string `json:"tags"`          // 标签
	SkuId         int    `json:"skuId"`         // skuId
	Specs         string `json:"specs"`         // specs值
	Num           int    `json:"num"`           // 数量
	Status        int    `json:"status"`        // 库存状态：0-正常 1-缺货 2-下架
}

type PortalAddressVO struct {
	Id          int    `json:"id"`          // 地址ID
	Contacts    string `json:"contacts"`    // 联系人
	Mobile      string `json:"mobile"`      // 手机号
	ProvinceId  int    `json:"provinceId"`  // 省份编码
	CityId      int    `json:"cityId"`      // 城市编码
	AreaId      int    `json:"areaId"`      // 地区编码
	ProvinceStr string `json:"provinceStr"` // 省份
	CityStr     string `json:"cityStr"`     // 城市
	AreaStr     string `json:"areaStr"`     // 地区
	Address     string `json:"address"`     // 详细地址
	IsDefault   int    `json:"is_default"`  // 默认收货地址：0-否 1-是
}

type PortalAddressReq struct {
	Id          int    `json:"id"`          // 地址ID
	Contacts    string `json:"contacts"`    // 联系人
	Mobile      string `json:"mobile"`      // 手机号
	ProvinceId  int    `json:"provinceId"`  // 省份编码
	CityId      int    `json:"cityId"`      // 城市编码
	AreaId      int    `json:"areaId"`      // 地区编码
	ProvinceStr string `json:"provinceStr"` // 省份
	CityStr     string `json:"cityStr"`     // 城市
	AreaStr     string `json:"areaStr"`     // 地区
	Address     string `json:"address"`     // 详细地址
	IsDefault   int    `json:"is_default"`  // 默认收货地址：0-否 1-是
}

type PortalCartPlaceOrderReq struct {
	AddressId      int                `json:"addressId"`      // 收货地址ID
	CouponLogId    int                `json:"couponLogId"`    // 优惠券记录ID
	DispatchAmount string             `json:"dispatchAmount"` // 运费
	ExpectAmount   string             `json:"expectAmount"`   // 预期支付金额
	GoodsList      []PortalOrderGoods `json:"goodsList"`      // 下单商品
}

type PortalOrderGoods struct {
	GoodsId int `json:"goodsId"` // 用户ID
	SkuId   int `json:"skuId"`   // sku ID
	Num     int `json:"num"`     // 数量
	CartId  int `json:"cartId"`  // 购物车记录ID
}

type PortalOrderListVO struct {
	Id        int                  `json:"id"`        // 订单ID
	OrderNo   string               `json:"orderNo"`   // 订单号
	PlaceTime string               `json:"placeTime"` // 下单时间
	PayAmount string               `json:"payAmount"` // 支付金额
	Status    int                  `json:"status"`    // 订单状态 -1 已取消 0-待付款 1-待发货 2-待收货 3-已完成
	GoodsList []PortalOrderGoodsVO `json:"goodsList"`
}

type PortalOrderGoodsVO struct {
	GoodsId int    `json:"goodsId"` // 商品ID
	Title   string `json:"title"`   // 标题
	Price   string `json:"price"`   // 价格
	Picture string `json:"picture"` // 图片
	SkuId   int    `json:"skuId"`   // skuId
	Specs   string `json:"specs"`   // specs值
	Num     int    `json:"num"`     // 数量
}

type PortalOrderDetailVO struct {
	Id        int                  `json:"id"`        // 订单ID
	OrderNo   string               `json:"orderNo"`   // 订单号
	PlaceTime string               `json:"placeTime"` // 下单时间
	PayAmount string               `json:"payAmount"` // 支付金额
	Status    int                  `json:"status"`    // 订单状态 -1 已取消 0-待付款 1-待发货 2-待收货 3-已完成
	GoodsList []PortalOrderGoodsVO `json:"goodsList"`
	Address   AddressSnapshot      `json:"address"`
}
