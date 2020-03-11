package web

import (
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"wechat-mall-backend/handler/cms"
	"wechat-mall-backend/handler/portal"
	"wechat-mall-backend/service"
)

func NewRouter(app *App) *mux.Router {
	router := mux.NewRouter()
	s := service.NewService(app.Conf)
	cmsHandler := cms.NewHandler(s)
	portalHandler := portal.NewHandler(s)

	registerHandler(router, cmsHandler, portalHandler)
	return router
}

func registerHandler(router *mux.Router, cmsHandler *cms.Handler, portalHandler *portal.Handler) {
	mw := &Middleware{}
	chain := alice.New(mw.LoggingHandler, mw.RecoverPanic, mw.ValidateAuthToken)
	router.Handle("/api/wxapp/login", chain.ThenFunc(portalHandler.Login)).Methods("POST")

	router.Handle("/cms/login", chain.ThenFunc(cmsHandler.Login)).Methods("POST")
	router.Handle("/cms/refresh", chain.ThenFunc(cmsHandler.Refresh)).Methods("GET")
	router.Handle("/cms/banner/list", chain.ThenFunc(cmsHandler.GetBannerList)).Methods("GET").Queries("page", "{page}").Queries("size", "{size}")
	router.Handle("/cms/banner/{id:[0-9]+}", chain.ThenFunc(cmsHandler.GetBanner)).Methods("GET")
	router.Handle("/cms/banner/edit", chain.ThenFunc(cmsHandler.DoEditBanner)).Methods("POST")
	router.Handle("/cms/banner/{id:[0-9]+}", chain.ThenFunc(cmsHandler.DoDeleteBanner)).Methods("DELETE")
	router.Handle("/cms/category/list", chain.ThenFunc(cmsHandler.GetCategoryList)).Methods("GET").Queries("page", "{page}").Queries("size", "{size}")
	router.Handle("/cms/category/{id:[0-9]+}", chain.ThenFunc(cmsHandler.GetCategoryById)).Methods("GET")
	router.Handle("/cms/category/edit", chain.ThenFunc(cmsHandler.DoEditCategory)).Methods("POST")
	router.Handle("/cms/category/{id:[0-9]+}", chain.ThenFunc(cmsHandler.DoDeleteCategory)).Methods("DELETE")
	router.Handle("/cms/grid_category/list", chain.ThenFunc(cmsHandler.GetGridCategoryList)).Methods("GET").Queries("page", "{page}").Queries("size", "{size}")
	router.Handle("/cms/grid_category/{id:[0-9]+}", chain.ThenFunc(cmsHandler.GetGridCategory)).Methods("GET")
	router.Handle("/cms/grid_category/edit", chain.ThenFunc(cmsHandler.DoEditGridCategory)).Methods("POST")
	router.Handle("/cms/grid_category/{id:[0-9]+}", chain.ThenFunc(cmsHandler.DoDeleteGridCategory)).Methods("DELETE")
	router.Handle("/cms/spec/list", chain.ThenFunc(cmsHandler.GetSpecificationList)).Methods("GET").Queries("page", "{page}").Queries("size", "{size}")
	router.Handle("/cms/spec/{id:[0-9]+}", chain.ThenFunc(cmsHandler.GetSpecification)).Methods("GET")
	router.Handle("/cms/spec/edit", chain.ThenFunc(cmsHandler.DoEditSpecification)).Methods("POST")
	router.Handle("/cms/spec/{id:[0-9]+}", chain.ThenFunc(cmsHandler.DoDeleteSpecification)).Methods("DELETE")
	router.Handle("/cms/spec/attr/list", chain.ThenFunc(cmsHandler.GetSpecificationAttrList)).Methods("GET").Queries("specId", "{specId}")
	router.Handle("/cms/spec/attr/{id:[0-9]+}", chain.ThenFunc(cmsHandler.GetSpecificationAttr)).Methods("GET")
	router.Handle("/cms/spec/attr/edit", chain.ThenFunc(cmsHandler.DoEditSpecificationAttr)).Methods("POST")
	router.Handle("/cms/spec/attr/{id:[0-9]+}", chain.ThenFunc(cmsHandler.DoDeleteSpecificationAttr)).Methods("DELETE")
	router.Handle("/cms/goods/list", chain.ThenFunc(cmsHandler.GetGoodsList)).Methods("GET").Queries("page", "{page}").Queries("size", "{size}")
	router.Handle("/cms/goods/{id:[0-9]+}", chain.ThenFunc(cmsHandler.GetGoods)).Methods("GET")
	router.Handle("/cms/goods/edit", chain.ThenFunc(cmsHandler.DoEditGoods)).Methods("POST")
	router.Handle("/cms/goods/{id:[0-9]+}", chain.ThenFunc(cmsHandler.DoDeleteGoods)).Methods("DELETE")
	router.Handle("/cms/goods/spec", chain.ThenFunc(cmsHandler.GetGoodsSpecList)).Methods("GET").Queries("goodsId", "{goodsId}")
	router.Handle("/cms/sku/list", chain.ThenFunc(cmsHandler.GetSKUList)).Methods("GET").Queries("page", "{page}").Queries("size", "{size}")
	router.Handle("/cms/sku/{id:[0-9]+}", chain.ThenFunc(cmsHandler.GetSKU)).Methods("GET")
	router.Handle("/cms/sku/edit", chain.ThenFunc(cmsHandler.DoEditSKU)).Methods("POST")
	router.Handle("/cms/sku/{id:[0-9]+}", chain.ThenFunc(cmsHandler.DoDeleteSKU)).Methods("DELETE")
	router.Handle("/cms/activity/list", chain.ThenFunc(cmsHandler.GetActivityList)).Methods("GET").Queries("page", "{page}").Queries("size", "{size}")
	router.Handle("/cms/activity/{id:[0-9]+}", chain.ThenFunc(cmsHandler.GetActivity)).Methods("GET")
	router.Handle("/cms/activity/edit", chain.ThenFunc(cmsHandler.DoEditActivity)).Methods("POST")
	router.Handle("/cms/activity/{id:[0-9]+}", chain.ThenFunc(cmsHandler.DoDeleteActivity)).Methods("DELETE")
	router.Handle("/cms/coupon/list", chain.ThenFunc(cmsHandler.GetCouponList)).Methods("GET").Queries("activityId", "{activityId}")
	router.Handle("/cms/coupon/{id:[0-9]+}", chain.ThenFunc(cmsHandler.GetCoupon)).Methods("GET")
	router.Handle("/cms/coupon/edit", chain.ThenFunc(cmsHandler.DoEditCoupon)).Methods("POST")
	router.Handle("/cms/coupon/{id:[0-9]+}", chain.ThenFunc(cmsHandler.DoDeleteCoupon)).Methods("DELETE")
}
