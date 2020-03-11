package defs

type CMSLoginReq struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type CMSTokenResp struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type CMSRegisterReq struct {
	Username string `json:"username:" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}

type CMSBannerReq struct {
	Id          int    `json:"id"`
	Picture     string `json:"picture" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type CMSBannerVO struct {
	Id          int    `json:"id"`
	Picture     string `json:"picture"`
	Name        string `json:"name"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CMSCategoryReq struct {
	Id          int    `json:"id"`
	ParentId    int    `json:"parent_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Sort        int    `json:"sort" validate:"required"`
	Online      int    `json:"online" validate:"required"`
	Picture     string `json:"picture" validate:"required"`
	Description string `json:"description"`
}

type CMSCategoryVO struct {
	Id          int    `json:"id"`
	ParentId    int    `json:"parent_id"`
	Name        string `json:"name"`
	Sort        int    `json:"sort"`
	Online      int    `json:"online"`
	Picture     string `json:"picture"`
	Description string `json:"description"`
}

type CMSGridCategoryReq struct {
	Id         int    `json:"id"`
	Title      string `json:"title" validate:"required"`
	Name       string `json:"name" validate:"required"`
	CategoryId int    `json:"category_id" validate:"required"`
	Picture    string `json:"picture" validate:"required"`
}

type CMSGridCategoryVO struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Name       string `json:"name"`
	CategoryId int    `json:"category"`
	Picture    string `json:"picture"`
}

type CMSSpecificationReq struct {
	Id          int    `json:"id"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Unit        string `json:"unit" validate:"required"`
	Standard    int    `json:"standard" validate:"required"`
}

type CMSSpecificationVO struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Unit        string `json:"unit"`
	Standard    int    `json:"standard"`
}

type CMSSpecificationAttrReq struct {
	Id     int    `json:"id"`
	SpecId int    `json:"spec_id" validate:"required"`
	Value  string `json:"value" validate:"required"`
	Extend string `json:"extend"`
}

type CMSSpecificationAttrVO struct {
	Id     int    `json:"id"`
	SpecId int    `json:"spec_id"`
	Value  string `json:"value"`
	Extend string `json:"extend"`
}

type CMSGoodsReq struct {
	Id              int    `json:"id"`
	BrandName       string `json:"brand_name"`
	Title           string `json:"title" validate:"required"`
	SubTitle        string `json:"sub_title" validate:"required"`
	Price           string `json:"price" validate:"required"`
	DiscountPrice   string `json:"discount_price" validate:"required"`
	CategoryId      int    `json:"category_id" validate:"required"`
	DefaultSkuId    int    `json:"default_sku_id"`
	Online          int    `json:"online"`
	Picture         string `json:"picture" validate:"required"`
	ForThemePicture string `json:"for_theme_picture" validate:"required"`
	BannerPicture   string `json:"banner_picture" validate:"required"`
	DetailPicture   string `json:"detail_picture" validate:"required"`
	Tags            string `json:"tags" validate:"required"`
	SketchSpecId    int    `json:"sketch_spec_id" validate:"required"`
	Description     string `json:"description"`
	SpecList        string `json:"spec_list" validate:"required"`
}

type CMSGoodsVO struct {
	Id            int    `json:"id"`
	BrandName     string `json:"brand_name"`
	Title         string `json:"title"`
	SubTitle      string `json:"sub_title"`
	Price         string `json:"price"`
	DiscountPrice string `json:"discount_price"`
	CategoryId    int    `json:"category_id"`
	DefaultSkuId  int    `json:"default_sku_id"`
	Online        int    `json:"online"`
	Picture       string `json:"picture"`
	BannerPicture string `json:"banner_picture"`
	DetailPicture string `json:"detail_picture"`
	Tags          string `json:"tags"`
	SketchSpecId  int    `json:"sketch_spec_id"`
	Description   string `json:"description"`
}

type CMSSKUReq struct {
	Id      int    `json:"id"`
	Title   string `json:"title" validate:"required"`
	Price   string `json:"price" validate:"required"`
	Code    string `json:"code" validate:"required"`
	Stock   int    `json:"stock" validate:"required"`
	GoodsId int    `json:"goods_id" validate:"required"`
	Online  int    `json:"online" validate:"required"`
	Picture string `json:"picture" validate:"required"`
	Specs   string `json:"specs" validate:"required"`
}

type CMSSKUVO struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Price   string `json:"price"`
	Code    string `json:"code"`
	Stock   int    `json:"stock"`
	GoodsId int    `json:"goods_id"`
	Online  int    `json:"online"`
	Picture string `json:"picture"`
	Specs   string `json:"specs"`
}

type CMSCouponReq struct {
	Id          int    `json:"id"`
	Title       string `json:"title" validate:"required"`
	FullMoney   string `json:"full_money" validate:"required"`
	Minus       string `json:"minus" validate:"required"`
	Rate        string `json:"rate" validate:"required"`
	Type        int    `json:"type" validate:"required"`
	StartTime   string `json:"start_time" validate:"required"`
	EndTime     string `json:"end_time" validate:"required"`
	Description string `json:"description"`
	Online      int    `json:"online"`
}

type CMSCouponVO struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	FullMoney   string `json:"full_money"`
	Minus       string `json:"minus"`
	Rate        string `json:"rate"`
	Type        int    `json:"type"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	Description string `json:"description"`
	Online      int    `json:"online"`
}
