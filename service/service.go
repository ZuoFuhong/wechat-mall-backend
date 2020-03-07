package service

import (
	"wechat-mall-backend/env"
	"wechat-mall-backend/model"
)

type WxappUser model.WxappUser
type CMSUser model.CMSUser
type Banner model.Banner
type BannerItem model.BannerItem
type Category model.Category

type Service struct {
	UserService          IUserService
	CMSUserService       ICMSUserService
	BannerService        IBannerService
	CategoryService      ICategoryService
	GridCategoryService  IGridCategoryService
	SpecificationService ISpecificationService
	SPUService           ISPUService
	SKUService           ISKUService
	ActivityService      IActivityService
	CouponService        ICouponService
}

func NewService(conf *env.Conf) *Service {
	service := &Service{}
	service.UserService = NewUserService(conf)
	service.CMSUserService = NewCMSUserService()
	service.BannerService = NewBannerService()
	service.CategoryService = NewCategoryService()
	service.GridCategoryService = NewGridCategoryService()
	service.SpecificationService = NewSpecificationService()
	service.SPUService = NewSPUService()
	service.SKUService = NewSKUService()
	service.ActivityService = NewActivityService()
	service.CouponService = NewCouponService()
	return service
}
