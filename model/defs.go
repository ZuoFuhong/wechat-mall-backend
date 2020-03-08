package model

type ID = int

type WxappUser struct {
	Id         ID
	Openid     string
	Nickname   string
	Avatar     string
	Mobile     string
	City       string
	CreateTime string
	UpdateTime string
}

type CMSUser struct {
	Id         ID
	Username   string
	Password   string
	Email      string
	Mobile     string
	Avatar     string
	CreateTime string
	UpdateTime string
}

type Banner struct {
	Id          ID
	Picture     string
	Name        string
	Title       string
	Description string
	Del         int
	CreateTime  string
	UpdateTime  string
}

type BannerItem struct {
	Id         ID
	BannerId   ID
	Name       string
	Picture    string
	Keyword    string
	Type       string
	Del        int
	CreateTime string
	UpdateTime string
}

type Category struct {
	Id          ID
	ParentId    int
	Name        string
	Sort        int
	Online      int
	Picture     string
	Description string
	Del         int
	CreateTime  string
	UpdateTime  string
}

type GridCategory struct {
	Id         ID
	Title      string
	Name       string
	CategoryId int
	Picture    string
	Del        int
	CreateTime string
	UpdateTime string
}

type Specification struct {
	Id          ID
	Name        string
	Description string
	Unit        string
	Standard    int
	Del         int
	CreateTime  string
	UpdateTime  string
}

type SpecificationAttr struct {
	Id         ID
	SpecId     ID
	Value      string
	Extend     string
	Del        int
	CreateTime string
	UpdateTime string
}

type SPU struct {
	Id              ID
	BrandName       string
	Title           string
	SubTitle        string
	Price           string
	DiscountPrice   string
	CategoryId      int
	DefaultSkuId    int
	Online          int
	Picture         string
	ForThemePicture string
	BannerPicture   string
	DetailPicture   string
	Tags            string
	SketchSpecId    int
	Description     string
	Del             int
	CreateTime      string
	UpdateTime      string
}

type SPUSpec struct {
	Id         ID
	SpuId      ID
	SpecId     int
	Del        int
	CreateTime string
	UpdateTime string
}

type SKU struct {
	Id         ID
	Title      string
	Price      string
	Code       string
	Stock      int
	SpuId      int
	Online     int
	Picture    string
	Specs      string
	Del        int
	CreateTime string
	UpdateTime string
}

type Activity struct {
	Id                 ID
	Title              string
	Name               string
	Remark             string
	Online             int
	StartTime          string
	EndTime            string
	Description        string
	EntrancePicture    string
	InternalTopPicture string
	Del                int
	CreateTime         string
	UpdateTime         string
}

type Coupon struct {
	Id          ID
	ActivityId  int
	Title       string
	FullMoney   string
	Minus       string
	Rate        string
	Type        int
	StartTime   string
	EndTime     string
	Description string
	Del         int
	CreateTime  string
	UpdateTime  string
}
