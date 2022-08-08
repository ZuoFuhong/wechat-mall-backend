package interfaces

import (
	"gorm.io/gorm"
	"net/http"
	"wechat-mall-backend/app/domain/service"
	"wechat-mall-backend/app/infra/database"
)

type MallHttpServiceImpl struct {
	addressService  service.IAddressService
	bannerService   service.IBannerService
	browseService   service.IBrowseRecordService
	cartService     service.ICartService
	categoryService service.ICategoryService
	cmsUserService  service.ICMSUserService
	couponService   service.ICouponService
	goodsService    service.IGoodsService
	gridService     service.IGridCategoryService
	orderService    service.IOrderService
	skuService      service.ISKUService
	specService     service.ISpecificationService
	userService     service.IUserService
}

func InitializeService(db *gorm.DB) *MallHttpServiceImpl {
	bannerRepos := database.NewBannerRepos(db)
	browseRepos := database.NewBrowseRepos(db)
	categoryRepos := database.NewCategoryRepos(db)
	moduleRepos := database.NewCmsModuleRepos(db)
	cmsUserRepos := database.NewCmsUserRepos(db)
	couponRepos := database.NewCouponRepos(db)
	goodsRepos := database.NewGoodsRepos(db)
	skuRepos := database.NewGoodsSkuRepos(db)
	gridRepos := database.NewGridCategoryRepos(db)
	orderRepos := database.NewOrderRepos(db)
	cartRepos := database.NewUserCart(db)
	userRepos := database.NewUserRepos(db)

	addressService := service.NewAddressService(userRepos)
	bannerService := service.NewBannerService(bannerRepos)
	browseService := service.NewBrowseRecordService(browseRepos)
	cartService := service.NewCartService(cartRepos, goodsRepos, skuRepos)
	categoryService := service.NewCategoryService(categoryRepos, goodsRepos)
	cmsUserService := service.NewCMSUserService(cmsUserRepos, moduleRepos)
	couponService := service.NewCouponService(couponRepos, categoryRepos)
	goodsService := service.NewGoodsService(goodsRepos, skuRepos, orderRepos)
	gridService := service.NewGridCategoryService(gridRepos)
	orderService := service.NewOrderService(orderRepos, goodsRepos, userRepos, cartRepos, skuRepos, couponRepos)
	skuService := service.NewSKUService(skuRepos)
	specService := service.NewSpecificationService(skuRepos)
	userService := service.NewUserService(userRepos)
	return &MallHttpServiceImpl{
		addressService:  addressService,
		bannerService:   bannerService,
		browseService:   browseService,
		cartService:     cartService,
		categoryService: categoryService,
		cmsUserService:  cmsUserService,
		couponService:   couponService,
		goodsService:    goodsService,
		gridService:     gridService,
		orderService:    orderService,
		skuService:      skuService,
		specService:     specService,
		userService:     userService,
	}
}

func (m *MallHttpServiceImpl) Ping(w http.ResponseWriter, r *http.Request) {
	Ok(w, "ok")
}
