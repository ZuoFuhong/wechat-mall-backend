package view

type PortalSkuVO struct {
	Id      int     `json:"id"`      // skuId
	Picture string  `json:"picture"` // 图片
	Title   string  `json:"title"`   // 标题
	Price   float64 `json:"price"`   // 价格
	Code    string  `json:"code"`    // 编码
	Stock   int     `json:"stock"`   // 库存量
	Specs   string  `json:"specs"`   // 多规格属性
}

type PortalGoodsInfo struct {
	Id            int             `json:"id"`            // 商品ID
	Title         string          `json:"title"`         // 标题
	Price         float64         `json:"price"`         // 价格
	Picture       string          `json:"picture"`       // 主图
	BannerPicture string          `json:"bannerPicture"` // 详情图
	DetailPicture string          `json:"detailPicture"` // 轮播图
	Tags          string          `json:"tags"`          // 标签
	Description   string          `json:"description"`   // 详情
	SkuList       []*PortalSkuVO  `json:"skuList"`       // sku列表
	SpecList      []*PortalSpecVO `json:"specList"`      // 规格列表
}

type PortalSpecVO struct {
	SpecId   int                 `json:"specId"`   // 规格ID
	Name     string              `json:"name"`     // 规格名称
	AttrList []*PortalSpecAttrVO `json:"attrList"` // 规格属性
}

type PortalSpecAttrVO struct {
	AttrId int    `json:"attrId"` // 属性ID
	Value  string `json:"value"`  // 属性名称
}

type CMSGoodsSpecVO struct {
	SpecId   int                       `json:"specId"`
	Name     string                    `json:"name" validate:"required"`
	AttrList []*CMSSpecificationAttrVO `json:"attrList"`
}

type CMSSpecificationAttrVO struct {
	Id     int    `json:"id"`
	SpecId int    `json:"specId"`
	Value  string `json:"value"`
	Extend string `json:"extend"`
}

type PortalGoodsListVO struct {
	Id       int     `json:"id"`       // 商品ID
	Title    string  `json:"title"`    // 标题
	Price    float64 `json:"price"`    // 价格
	Picture  string  `json:"picture"`  // 图片
	HumanNum int     `json:"humanNum"` // 购买人数
}

type PortalBrowseRecordVO struct {
	Id         int    `json:"id"`         // 记录ID
	GoodsId    int    `json:"goodsId"`    // 商品ID
	Picture    string `json:"picture"`    // 商品图片
	Title      string `json:"title"`      // 商品标题
	Price      string `json:"price"`      // 商品价格
	BrowseTime string `json:"browseTime"` // 浏览时间，格式：yyyy-MM-dd HH:mm:ss
}

type CMSSpecificationVO struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Unit        string `json:"unit"`
	Standard    int    `json:"standard"`
}

type CMSGoodsListVO struct {
	Id           int    `json:"id"`
	BrandName    string `json:"brandName"`
	Title        string `json:"title"`
	Price        string `json:"price"`
	CategoryId   int    `json:"categoryId"`
	CategoryName string `json:"categoryName"`
	Online       int    `json:"online"`
	Picture      string `json:"picture"`
}

type CMSGoodsVO struct {
	Id            int    `json:"id"`
	BrandName     string `json:"brandName"`
	Title         string `json:"title"`
	Price         string `json:"price"`
	DiscountPrice string `json:"discountPrice"`
	CategoryId    int    `json:"categoryId"`
	SubCategoryId int    `json:"subCategoryId"`
	CategoryName  string `json:"categoryName"`
	Online        int    `json:"online"`
	Picture       string `json:"picture"`
	BannerPicture string `json:"bannerPicture"`
	DetailPicture string `json:"detailPicture"`
	Tags          string `json:"tags"`
}

type CMSSkuListVO struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Price   string `json:"price"`
	Code    string `json:"code"`
	Stock   int    `json:"stock"`
	GoodsId int    `json:"goodsId"`
	Online  int    `json:"online"`
	Picture string `json:"picture"`
	Specs   string `json:"specs"`
}

type CMSSkuDetailVO struct {
	Id            int    `json:"id"`
	Title         string `json:"title"`
	Price         string `json:"price"`
	Code          string `json:"code"`
	Stock         int    `json:"stock"`
	GoodsId       int    `json:"goodsId"`
	CategoryId    int    `json:"categoryId"`
	SubCategoryId int    `json:"subCategoryId"`
	Online        int    `json:"online"`
	Picture       string `json:"picture"`
	Specs         string `json:"specs"`
}
