package service

import (
	"wechat-mall-backend/env"
)

type Service struct {
	UserService          IUserService
	CMSUserService       ICMSUserService
	BannerService        IBannerService
	CategoryService      ICategoryService
	GridCategoryService  IGridCategoryService
	SpecificationService ISpecificationService
	GoodsService         IGoodsService
	SKUService           ISKUService
	CouponService        ICouponService
	CartService          ICartService
	AddressService       IAddressService
	OrderService         IOrderService
	BrowseRecordService  IBrowseRecordService
}

func NewService(conf *env.Conf) *Service {
	service := &Service{}
	service.UserService = NewUserService(conf)
	service.CMSUserService = NewCMSUserService()
	service.BannerService = NewBannerService()
	service.CategoryService = NewCategoryService()
	service.GridCategoryService = NewGridCategoryService()
	service.SpecificationService = NewSpecificationService()
	service.GoodsService = NewGoodsService()
	service.SKUService = NewSKUService()
	service.CouponService = NewCouponService()
	service.CartService = NewCartService()
	service.AddressService = NewAddressService()
	service.OrderService = NewOrderService()
	service.BrowseRecordService = NewBrowseRecordService()
	return service
}
