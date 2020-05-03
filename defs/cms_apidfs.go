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
	Id           int    `json:"id"`
	Picture      string `json:"picture" validate:"required"`
	Name         string `json:"name" validate:"required"`
	BusinessType int    `json:"businessType"`
	BusinessId   int    `json:"businessId"`
	Status       int    `json:"status"`
}

type CMSGoodsBannerVO struct {
	Id            int    `json:"id"`
	Picture       string `json:"picture"`
	Name          string `json:"name"`
	GoodsId       int    `json:"goodsId"`
	CategoryId    int    `json:"categoryId"`
	SubCategoryId int    `json:"subCategoryId"`
	Status        int    `json:"status"`
}

type CMSBannerVO struct {
	Id      int    `json:"id"`
	Picture string `json:"picture"`
	Name    string `json:"name"`
	Status  int    `json:"status"`
}

type CMSCategoryReq struct {
	Id          int    `json:"id"`
	ParentId    int    `json:"parentId"`
	Name        string `json:"name" validate:"required"`
	Sort        int    `json:"sort"`
	Online      int    `json:"online"`
	Picture     string `json:"picture" validate:"required"`
	Description string `json:"description"`
}

type CMSCategoryVO struct {
	Id          int    `json:"id"`
	ParentId    int    `json:"parentId"`
	Name        string `json:"name"`
	Sort        int    `json:"sort"`
	Online      int    `json:"online"`
	Picture     string `json:"picture"`
	Description string `json:"description"`
}

type CMSGridCategoryReq struct {
	Id         int    `json:"id"`
	Name       string `json:"name" validate:"required"`
	CategoryId int    `json:"categoryId" validate:"required"`
	Picture    string `json:"picture" validate:"required"`
}

type CMSGridCategoryVO struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	CategoryId   int    `json:"categoryId"`
	CategoryName string `json:"categoryName"`
	Picture      string `json:"picture"`
}

type CMSGridCategoryDetailVO struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	CategoryId      int    `json:"categoryId"`
	SubCategoryId   int    `json:"subCategoryId"`
	SubCategoryName string `json:"subCategoryName"`
	Picture         string `json:"picture"`
}

type CMSSpecificationReq struct {
	Id          int    `json:"id"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Unit        string `json:"unit" validate:"required"`
	Standard    int    `json:"standard"`
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
	SpecId int    `json:"specId" validate:"required"`
	Value  string `json:"value" validate:"required"`
	Extend string `json:"extend"`
}

type CMSSpecificationAttrVO struct {
	Id     int    `json:"id"`
	SpecId int    `json:"specId"`
	Value  string `json:"value"`
	Extend string `json:"extend"`
}

type CMSGoodsReq struct {
	Id            int    `json:"id"`
	BrandName     string `json:"brandName"`
	Title         string `json:"title" validate:"required"`
	Price         string `json:"price" validate:"required"`
	DiscountPrice string `json:"discountPrice"`
	CategoryId    int    `json:"categoryId" validate:"required"`
	Online        int    `json:"online"`
	Picture       string `json:"picture" validate:"required"`
	BannerPicture string `json:"bannerPicture" validate:"required"`
	DetailPicture string `json:"detailPicture" validate:"required"`
	Tags          string `json:"tags"`
	SpecList      []int  `json:"specList"`
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

type CMSSKUReq struct {
	Id      int    `json:"id"`
	Title   string `json:"title" validate:"required"`
	Price   string `json:"price" validate:"required"`
	Code    string `json:"code"`
	Stock   int    `json:"stock"`
	GoodsId int    `json:"goodsId" validate:"required"`
	Online  int    `json:"online"`
	Picture string `json:"picture" validate:"required"`
	Specs   string `json:"specs" validate:"required"`
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

type CMSCouponReq struct {
	Id          int    `json:"id"`
	Title       string `json:"title" validate:"required"`
	FullMoney   string `json:"fullMoney" validate:"required"`
	Minus       string `json:"minus" validate:"required"`
	Rate        string `json:"rate" validate:"required"`
	Type        int    `json:"type" validate:"required"`
	GrantNum    int    `json:"grantNum" validate:"required"`
	LimitNum    int    `json:"limitNum" validate:"required"`
	StartTime   string `json:"startTime" validate:"required"`
	EndTime     string `json:"endTime" validate:"required"`
	Description string `json:"description"`
	Online      int    `json:"online"`
}

type CMSCouponVO struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	FullMoney   string `json:"fullMoney"`
	Minus       string `json:"minus"`
	Rate        string `json:"rate"`
	Type        int    `json:"type"`
	GrantNum    int    `json:"grantNum"`
	LimitNum    int    `json:"limitNum"`
	StartTime   string `json:"startTime"`
	EndTime     string `json:"endTime"`
	Description string `json:"description"`
	Online      int    `json:"online"`
}

type CMSUserVO struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Mobile    string `json:"mobile"`
	Avatar    string `json:"avatar"`
	GroupId   int    `json:"groupId"`
	GroupName string `json:"groupName"`
}

type CMSUserReq struct {
	Id       int    `json:"id"`
	Avatar   string `json:"avatar" validate:"required"`
	Username string `json:"username" validate:"required"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile" validate:"required"`
	GroupId  int    `json:"groupId" validate:"required"`
}

type CMSResetUserPasswdReq struct {
	UserId   int    `json:"userId" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type CMSUserGroupReq struct {
	Id          int    `json:"id"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Auths       []int  `json:"auths"`
}

type CMSUserGroupVO struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Auths       []int  `json:"auths"`
}

type CMSModuleVO struct {
	Id          int               `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	PageList    []CMSModulePageVO `json:"pageList"`
}

type CMSModulePageVO struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CMSChangePasswordReq struct {
	OldPassword string `json:"oldPassword" validate:"required"`
	NewPassword string `json:"newPassword" validate:"required"`
}

type CMSGoodsSpecVO struct {
	SpecId   int                      `json:"specId"`
	Name     string                   `json:"name" validate:"required"`
	AttrList []CMSSpecificationAttrVO `json:"attrList"`
}

type CMSMarketMetricsVO struct {
	VisitorNum    int `json:"visitorNum"`
	SellOutSKUNum int `json:"sellOutSKUNum"`
	WaitingOrder  int `json:"waitingOrder"`
	ActivistOrder int `json:"activistOrder"`
}

type CMSOrderInfoVO struct {
	OrderNo        string            `json:"orderNo"`
	PlaceTime      string            `json:"placeTime"`
	Address        string            `json:"address"`
	PayAmount      float64           `json:"payAmount"`
	GoodsAmount    float64           `json:"goodsAmount"`
	DiscountAmount float64           `json:"discountAmount"`
	DispatchAmount float64           `json:"dispatchAmount"`
	Status         int               `json:"status"`
	TransactionId  string            `json:"transactionId"`
	PayTime        string            `json:"payTime"`
	DeliverTime    string            `json:"deliverTime"`
	FinishTime     string            `json:"finishTime"`
	GoodsList      []CMSOrderGoodsVO `json:"goodsList"`
}

type CMSOrderGoodsVO struct {
	Picture string  `json:"picture"`
	Title   string  `json:"title"`
	Price   float64 `json:"price"`
	Specs   string  `json:"specs"`
	Num     int     `json:"num"`
}

type CMSModifyOrderStatusReq struct {
	OrderNo string `json:"orderNo" validate:"required"` // 订单号
	Otype   int    `json:"otype" validate:"required"`   // 操作方式：1-确认发货，2-确认收货，3-确认付款
}

type CMSModifyOrderRemarkReq struct {
	OrderNo string `json:"orderNo" validate:"required"` // 订单号
	Remark  string `json:"remark" validate:"required"`  // 备注
}

type CMSModifyOrderGoodsReq struct {
	OrderNo string `json:"orderNo" validate:"required"` // 订单号
	GoodsId int    `json:"goodsId" validate:"required"` // 关联订单商品ID
	Price   string `json:"price" validate:"required"`   // 价格
}
