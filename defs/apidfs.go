package defs

const (
	AccessTokenExpire  = 2 * 3600
	RefreshTokenExpire = 30 * 24 * 3600
)

type WxappLoginResp struct {
	Token               string `json:"token" validate:"required"`
	ExpirationInMinutes int    `json:"expiration_in_minutes" validate:"required"`
}

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

type BannerVO struct {
	Id          int    `json:"id"`
	Picture     string `json:"picture"`
	Name        string `json:"name"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type BannerItemReq struct {
	Id       int    `json:"id"`
	BannerId int    `json:"banner_id" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Picture  string `json:"picture" validate:"required"`
	Keyword  string `json:"keyword"`
	Type     string `json:"type"`
}

type BannerItemVO struct {
	Id       int    `json:"id"`
	BannerId int    `json:"banner_id"`
	Name     string `json:"name"`
	Picture  string `json:"picture"`
	Keyword  string `json:"keyword"`
	Type     string `json:"type"`
}

type CategoryReq struct {
	Id          int    `json:"id"`
	ParentId    int    `json:"parent_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Sort        int    `json:"sort" validate:"required"`
	Online      int    `json:"online" validate:"required"`
	Picture     string `json:"picture" validate:"required"`
	Description string `json:"description"`
}

type CategoryVO struct {
	Id          int    `json:"id"`
	ParentId    int    `json:"parent_id"`
	Name        string `json:"name"`
	Sort        int    `json:"sort"`
	Online      int    `json:"online"`
	Picture     string `json:"picture"`
	Description string `json:"description"`
}

type GridCategoryReq struct {
	Id         int    `json:"id"`
	Title      string `json:"title" validate:"required"`
	Name       string `json:"name" validate:"required"`
	CategoryId int    `json:"category_id" validate:"required"`
	Picture    string `json:"picture" validate:"required"`
}

type GridCategoryVO struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Name       string `json:"name"`
	CategoryId int    `json:"category"`
	Picture    string `json:"picture"`
}

type SpecificationReq struct {
	Id          int    `json:"id"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Unit        string `json:"unit" validate:"required"`
	Standard    int    `json:"standard" validate:"required"`
}

type SpecificationVO struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Unit        string `json:"unit"`
	Standard    int    `json:"standard"`
}

type SpecificationAttrReq struct {
	Id     int    `json:"id"`
	SpecId int    `json:"spec_id" validate:"required"`
	Value  string `json:"value" validate:"required"`
	Extend string `json:"extend"`
}

type SpecificationAttrVO struct {
	Id     int    `json:"id"`
	SpecId int    `json:"spec_id"`
	Value  string `json:"value"`
	Extend string `json:"extend"`
}

type SPUReq struct {
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

type SPUVO struct {
	Id              int    `json:"id"`
	BrandName       string `json:"brand_name"`
	Title           string `json:"title"`
	SubTitle        string `json:"sub_title"`
	Price           string `json:"price"`
	DiscountPrice   string `json:"discount_price"`
	CategoryId      int    `json:"category_id"`
	DefaultSkuId    int    `json:"default_sku_id"`
	Online          int    `json:"online"`
	Picture         string `json:"picture"`
	ForThemePicture string `json:"for_theme_picture"`
	BannerPicture   string `json:"banner_picture"`
	DetailPicture   string `json:"detail_picture"`
	Tags            string `json:"tags"`
	SketchSpecId    int    `json:"sketch_spec_id"`
	Description     string `json:"description"`
}

type SKUReq struct {
	Id      int    `json:"id"`
	Title   string `json:"title" validate:"required"`
	Price   string `json:"price" validate:"required"`
	Code    string `json:"code" validate:"required"`
	Stock   int    `json:"stock" validate:"required"`
	SpuId   int    `json:"spu_id" validate:"required"`
	Online  int    `json:"online" validate:"required"`
	Picture string `json:"picture" validate:"required"`
	Specs   string `json:"specs" validate:"required"`
}

type SKUVO struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Price   string `json:"price"`
	Code    string `json:"code"`
	Stock   int    `json:"stock"`
	SpuId   int    `json:"spu_id"`
	Online  int    `json:"online"`
	Picture string `json:"picture"`
	Specs   string `json:"specs"`
}

type ActivityReq struct {
	Id                 int    `json:"id"`
	Title              string `json:"title" validate:"required"`
	Name               string `json:"name" validate:"required"`
	Remark             string `json:"remark" validate:"required"`
	Online             int    `json:"online" validate:"required"`
	StartTime          string `json:"start_time" validate:"required"`
	EndTime            string `json:"end_time" validate:"required"`
	Description        string `json:"description" validate:"required"`
	EntrancePicture    string `json:"entrance_picture" validate:"required"`
	InternalTopPicture string `json:"internal_top_picture" validate:"required"`
}

type ActivityVO struct {
	Id                 int    `json:"id"`
	Title              string `json:"title"`
	Name               string `json:"name"`
	Remark             string `json:"remark"`
	Online             int    `json:"online"`
	StartTime          string `json:"start_time"`
	EndTime            string `json:"end_time"`
	Description        string `json:"description"`
	EntrancePicture    string `json:"entrance_picture"`
	InternalTopPicture string `json:"internal_top_picture"`
}

type CouponReq struct {
	Id          int    `json:"id"`
	ActivityId  int    `json:"activity_id" validate:"required"`
	Title       string `json:"title" validate:"required"`
	FullMoney   string `json:"full_money" validate:"required"`
	Minus       string `json:"minus" validate:"required"`
	Rate        string `json:"rate" validate:"required"`
	Type        int    `json:"type" validate:"required"`
	StartTime   string `json:"start_time" validate:"required"`
	EndTime     string `json:"end_time" validate:"required"`
	Description string `json:"description"`
}

type CouponVO struct {
	Id          int    `json:"id"`
	ActivityId  int    `json:"activity_id"`
	Title       string `json:"title"`
	FullMoney   string `json:"full_money"`
	Minus       string `json:"minus"`
	Rate        string `json:"rate"`
	Type        int    `json:"type"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	Description string `json:"description"`
}
