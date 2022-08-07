package interfaces

import "wechat-mall-backend/app/domain/entity"

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

type CMSCategoryReq struct {
	Id          int    `json:"id"`
	ParentId    int    `json:"parentId"`
	Name        string `json:"name" validate:"required"`
	Sort        int    `json:"sort"`
	Online      int    `json:"online"`
	Picture     string `json:"picture" validate:"required"`
	Description string `json:"description"`
}

type CMSGridCategoryReq struct {
	Id         int    `json:"id"`
	Name       string `json:"name" validate:"required"`
	CategoryId int    `json:"categoryId" validate:"required"`
	Picture    string `json:"picture" validate:"required"`
}

type CMSSpecificationReq struct {
	Id          int    `json:"id"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Unit        string `json:"unit" validate:"required"`
	Standard    int    `json:"standard"`
}

type CMSSpecificationAttrReq struct {
	Id     int    `json:"id"`
	SpecId int    `json:"specId" validate:"required"`
	Value  string `json:"value" validate:"required"`
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

type CMSChangePasswordReq struct {
	OldPassword string `json:"oldPassword" validate:"required"`
	NewPassword string `json:"newPassword" validate:"required"`
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

type WxappAuthUserInfoReq struct {
	NickName  string `json:"nickName" validate:"required"`
	AvatarUrl string `jsoN:"avatarUrl" validate:"required"`
	Gender    int    `json:"gender"`
	Country   string `json:"country"`
	Province  string `json:"province"`
	City      string `json:"city"`
}

type WxappAuthPhone struct {
	EncryptedData string `json:"encryptedData" validate:"required"`
	Iv            string `jsoN:"iv" validate:"required"`
}

type PortalTakeCouponReq struct {
	CouponId int `json:"couponId" validate:"required"`
}

type PortalCartGoodsReq struct {
	GoodsId int `json:"goodsId" validate:"required"` // 商品ID
	SkuId   int `json:"skuId" validate:"required"`   // skuId
	Num     int `json:"num" validate:"required"`     // 数量
}

type PortalEditCartReq struct {
	Id  int `json:"id"`  // 主键
	Num int `json:"num"` // 数量：-1 减一件 0 删除 1 加一件
}

type PortalAddressReq struct {
	Id          int    `json:"id"`          // 地址ID
	Contacts    string `json:"contacts"`    // 联系人
	Mobile      string `json:"mobile"`      // 手机号
	ProvinceId  string `json:"provinceId"`  // 省份编码
	CityId      string `json:"cityId"`      // 城市编码
	AreaId      string `json:"areaId"`      // 地区编码
	ProvinceStr string `json:"provinceStr"` // 省份
	CityStr     string `json:"cityStr"`     // 城市
	AreaStr     string `json:"areaStr"`     // 地区
	Address     string `json:"address"`     // 详细地址
	IsDefault   int    `json:"isDefault"`   // 默认收货地址：0-否 1-是
}

type OrderRefundApplyReq struct {
	OrderNo string `json:"orderNo" validate:"required"` // 订单号
	Reason  string `json:"reason" validate:"required"`  // 退款原因
}

type PortalCartPlaceOrderReq struct {
	AddressId      int                 `json:"addressId"`      // 收货地址ID
	CouponLogId    int                 `json:"couponLogId"`    // 优惠券记录ID
	DispatchAmount string              `json:"dispatchAmount"` // 运费
	ExpectAmount   string              `json:"expectAmount"`   // 预期支付金额
	GoodsList      []*entity.CartGoods `json:"goodsList"`      // 下单商品
}
