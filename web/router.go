package web

import (
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"wechat-mall-backend/handler"
)

func NewRouter(app *App) *mux.Router {
	router := mux.NewRouter()
	h := handler.NewHandler(app.Conf)
	registerHandler(router, h)
	return router
}

func registerHandler(router *mux.Router, handler *handler.Handler) {
	mw := &Middleware{}
	chain := alice.New(mw.LoggingHandler, mw.RecoverPanic)
	router.Handle("/api/wxapp/login", chain.ThenFunc(handler.WxappHandler.Login)).Methods("POST")
	router.Handle("/cms/login", chain.ThenFunc(handler.CMSHandler.Login)).Methods("POST")
	router.Handle("/cms/refresh", chain.ThenFunc(handler.CMSHandler.Refresh)).Methods("GET")
	router.Handle("/cms/register", chain.ThenFunc(handler.CMSHandler.Register)).Methods("POST")
	router.Handle("/cms/register/activate/{code:[0-9a-zA-Z]+}", chain.ThenFunc(handler.CMSHandler.RegisterActivate)).Methods("GET")
	router.Handle("/cms/banner/list", chain.ThenFunc(handler.CMSHandler.GetBannerList)).Methods("GET")
	router.Handle("/cms/banner/{id:[0-9]+}", chain.ThenFunc(handler.CMSHandler.GetBanner)).Methods("GET")
	router.Handle("/cms/banner/edit", chain.ThenFunc(handler.CMSHandler.DoEditBanner)).Methods("POST")
	router.Handle("/cms/banner/{id:[0-9]+}", chain.ThenFunc(handler.CMSHandler.DoDeleteBanner)).Methods("DELETE")
	router.Handle("/cms/banner_item/list/{id:[0-9]+}", chain.ThenFunc(handler.CMSHandler.GetBannerItemList)).Methods("GET")
	router.Handle("/cms/banner_item/{id:[0-9]+}", chain.ThenFunc(handler.CMSHandler.GetBannerItem)).Methods("GET")
	router.Handle("/cms/banner_item/edit", chain.ThenFunc(handler.CMSHandler.DoEditBannerItem)).Methods("POST")
	router.Handle("/cms/banner_item/{id:[0-9]+}", chain.ThenFunc(handler.CMSHandler.DoDeleteBannerItem)).Methods("DELETE")
	router.Handle("/cms/category/list", chain.ThenFunc(handler.CMSHandler.GetCategoryList)).Methods("GET")
	router.Handle("/cms/category/{id:[0-9]+}", chain.ThenFunc(handler.CMSHandler.GetCategoryById)).Methods("GET")
	router.Handle("/cms/category/edit", chain.ThenFunc(handler.CMSHandler.DoEditCategory)).Methods("POST")
	router.Handle("/cms/category/{id:[0-9]+}", chain.ThenFunc(handler.CMSHandler.DoDeleteCategory)).Methods("DELETE")
	router.Handle("/cms/grid_category/list", chain.ThenFunc(handler.CMSHandler.GetGridCategoryList)).Methods("GET")
	router.Handle("/cms/grid_category/{id:[0-9]+}", chain.ThenFunc(handler.CMSHandler.GetGridCategory)).Methods("GET")
	router.Handle("/cms/grid_category/edit", chain.ThenFunc(handler.CMSHandler.DoEditGridCategory)).Methods("POST")
	router.Handle("/cms/grid_category/{id:[0-9]+}", chain.ThenFunc(handler.CMSHandler.DoDeleteGridCategory)).Methods("DELETE")
	router.Handle("/cms/spec/list", chain.ThenFunc(handler.CMSHandler.GetSpecificationList)).Methods("GET")
	router.Handle("/cms/spec/{id:[0-9]+}", chain.ThenFunc(handler.CMSHandler.GetSpecification)).Methods("GET")
	router.Handle("/cms/spec/edit", chain.ThenFunc(handler.CMSHandler.DoEditSpecification)).Methods("POST")
	router.Handle("/cms/spec/{id:[0-9]+}", chain.ThenFunc(handler.CMSHandler.DoDeleteSpecification)).Methods("DELETE")
	router.Handle("/cms/spec/attr/list", chain.ThenFunc(handler.CMSHandler.GetSpecificationAttrList)).Methods("GET")
	router.Handle("/cms/spec/attr/{id:[0-9]+}", chain.ThenFunc(handler.CMSHandler.GetSpecificationAttr)).Methods("GET")
	router.Handle("/cms/spec/attr/edit", chain.ThenFunc(handler.CMSHandler.DoEditSpecificationAttr)).Methods("POST")
	router.Handle("/cms/spec/attr/{id:[0-9]+}", chain.ThenFunc(handler.CMSHandler.DoDeleteSpecificationAttr)).Methods("DELETE")
	router.Handle("/cms/spu/list", chain.ThenFunc(handler.CMSHandler.GetSPUList)).Methods("GET")
	router.Handle("/cms/spu/{id:[0-9]+}", chain.ThenFunc(handler.CMSHandler.GetSPU)).Methods("GET")
	router.Handle("/cms/spu/edit", chain.ThenFunc(handler.CMSHandler.DoEditSPU)).Methods("POST")
	router.Handle("/cms/spu/{id:[0-9]+}", chain.ThenFunc(handler.CMSHandler.DoDeleteSPU)).Methods("DELETE")
	router.Handle("/cms/sku/list", chain.ThenFunc(handler.CMSHandler.GetSKUList)).Methods("GET")
	router.Handle("/cms/sku/{id:[0-9]+}", chain.ThenFunc(handler.CMSHandler.GetSKU)).Methods("GET")
	router.Handle("/cms/sku/edit", chain.ThenFunc(handler.CMSHandler.DoEditSKU)).Methods("POST")
	router.Handle("/cms/sku/{id:[0-9]+}", chain.ThenFunc(handler.CMSHandler.DoDeleteSKU)).Methods("DELETE")
	router.Handle("/cms/activity/list", chain.ThenFunc(handler.CMSHandler.GetActivityList)).Methods("GET")
	router.Handle("/cms/activity/{id:[0-9]+}", chain.ThenFunc(handler.CMSHandler.GetActivity)).Methods("GET")
	router.Handle("/cms/activity/edit", chain.ThenFunc(handler.CMSHandler.DoEditActivity)).Methods("POST")
	router.Handle("/cms/activity/{id:[0-9]+}", chain.ThenFunc(handler.CMSHandler.DoDeleteSKU)).Methods("DELETE")
	router.Handle("/cms/coupon/list", chain.ThenFunc(handler.CMSHandler.GetCouponList)).Methods("GET")
	router.Handle("/cms/coupon/{id:[0-9]+}", chain.ThenFunc(handler.CMSHandler.GetCoupon)).Methods("GET")
	router.Handle("/cms/coupon/edit", chain.ThenFunc(handler.CMSHandler.DoEditCoupon)).Methods("POST")
	router.Handle("/cms/coupon/{id:[0-9]+}", chain.ThenFunc(handler.CMSHandler.DoDeleteCoupon)).Methods("DELETE")
}
