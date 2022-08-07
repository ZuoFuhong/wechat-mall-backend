package errcode

const (

	// ErrorInternalFaults 系统繁忙
	ErrorInternalFaults = 10004

	// ErrorWechatError 微信内部异常
	ErrorWechatError = 10007

	// ErrorTokenInvalid Token 无效或已过期
	ErrorTokenInvalid = 10008

	// ErrorRefreshTokenInvalid 刷新 Token 失败
	ErrorRefreshTokenInvalid = 10009

	// BadRequestParam 参数验证失败
	BadRequestParam = 10010

	// NotFoundBanner Banner不存在
	NotFoundBanner = 10011

	// NotFoundCategory 分类不存在
	NotFoundCategory = 10012

	// NotFoundGridCategory 宫格不存在
	NotFoundGridCategory = 10013

	// NotFoundSpecification 规格不存在
	NotFoundSpecification = 10014

	// NotFoundSpecificationAttr 规格属性不存在
	NotFoundSpecificationAttr = 10015

	// NotFoundGoods 商品不存在
	NotFoundGoods = 10016

	// NotFoundGoodsSku SKU不存在
	NotFoundGoodsSku = 10017

	// NotFoundCoupon 优惠券不存在
	NotFoundCoupon = 10019

	// NotFoundUserAddress 收货地址不存在
	NotFoundUserAddress = 10020

	// NotFoundOrderRecord 订单不存在
	NotFoundOrderRecord = 10021

	// NotFoundCartGoods 购物车商品不存在
	NotFoundCartGoods = 10022

	// NotFoundUserGroup 分组不存在
	NotFoundUserGroup = 10023

	// NotFoundCmsUser 用户不存在
	NotFoundCmsUser = 10024

	// NotFoundCmsPage 页面不存在
	NotFoundCmsPage = 10025

	// NotFoundUserRecord 用户不存在
	NotFoundUserRecord = 10026

	// NotFoundOrderRefund 退款单不存在
	NotFoundOrderRefund = 10027

	// NotAllowOperation 不允许操作
	NotAllowOperation = 10028
)
