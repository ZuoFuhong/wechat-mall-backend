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
	router.Handle("/api/wxapp/auth-phone", chain.ThenFunc(portalHandler.AuthPhone)).Methods("POST")
	router.Handle("/api/wxapp/auth-info", chain.ThenFunc(portalHandler.AuthUserInfo)).Methods("POST")
	router.Handle("/api/home/banner", chain.ThenFunc(portalHandler.GetBannerList)).Methods("GET").Queries("page", "{page}").Queries("size", "{size}")
	router.Handle("/api/home/grid", chain.ThenFunc(portalHandler.GetGridCategoryList)).Methods("GET")
	router.Handle("/api/goods/list", chain.ThenFunc(portalHandler.GetGoodsList)).Methods("GET").Queries("k", "{k}").Queries("c", "{c}").Queries("o", "{o}").Queries("page", "{page}").Queries("size", "{size}")
	router.Handle("/api/goods/detail", chain.ThenFunc(portalHandler.GetGoodsDetail)).Methods("GET").Queries("id", "{id}")
	router.Handle("/api/cart/list", chain.ThenFunc(portalHandler.GetCartGoodsList)).Methods("GET").Queries("page", "{page}").Queries("size", "{size}")
	router.Handle("/api/cart/edit", chain.ThenFunc(portalHandler.EditCartGoods)).Methods("POST")
	router.Handle("/api/coupon/list", chain.ThenFunc(portalHandler.GetCouponList)).Methods("GET").Queries("page", "{page}").Queries("size", "{size}")
	router.Handle("/api/coupon/take", chain.ThenFunc(portalHandler.TakeCoupon)).Methods("POST")
	router.Handle("/api/user/coupon/list", chain.ThenFunc(portalHandler.GetUserCouponList)).Methods("GET").Queries("status", "{status}").Queries("page", "{page}").Queries("size", "{size}")
	router.Handle("/api/user/coupon", chain.ThenFunc(portalHandler.DoDeleteCouponLog)).Methods("DELETE").Queries("id", "{id}")
	router.Handle("/api/user/address/list", chain.ThenFunc(portalHandler.GetAddressList)).Methods("GET").Queries("page", "{page}").Queries("size", "{size}")
	router.Handle("/api/user/address/edit", chain.ThenFunc(portalHandler.EditAddress)).Methods("POST")
	router.Handle("/api/user/address", chain.ThenFunc(portalHandler.DoDeleteAddress)).Methods("DELETE").Queries("id", "{id}")
	router.Handle("/api/placeorder", chain.ThenFunc(portalHandler.PlaceOrder)).Methods("POST")
	router.Handle("/api/order/list", chain.ThenFunc(portalHandler.GetOrderList)).Methods("GET").Queries("status", "{status}").Queries("page", "{page}").Queries("size", "{size}")
	router.Handle("/api/order/detail", chain.ThenFunc(portalHandler.GetOrderDetail)).Methods("GET").Queries("id", "{id}")
	router.Handle("/wxpay/notify", chain.ThenFunc(portalHandler.WxPayNotify)).Methods("POST")
	router.Handle("/cms/user/login", chain.ThenFunc(cmsHandler.Login)).Methods("POST")
	router.Handle("/cms/user/refresh", chain.ThenFunc(cmsHandler.Refresh)).Methods("GET")
	router.Handle("/cms/user/info", chain.ThenFunc(cmsHandler.GetUserInfo)).Methods("GET")
	router.Handle("/cms/user/change_password", chain.ThenFunc(cmsHandler.DoChangePassword)).Methods("POST")
	router.Handle("/cms/admin/users", chain.ThenFunc(cmsHandler.GetUserList)).Methods("GET").Queries("page", "{page}").Queries("size", "{size}")
	router.Handle("/cms/admin/user", chain.ThenFunc(cmsHandler.DoEditUser)).Methods("POST")
	router.Handle("/cms/admin/user", chain.ThenFunc(cmsHandler.DoDeleteCMSUser)).Methods("DELETE").Queries("id", "{id}")
	router.Handle("/cms/admin/groups", chain.ThenFunc(cmsHandler.GetUserGroupList)).Methods("GET").Queries("page", "{page}").Queries("size", "{size}")
	router.Handle("/cms/admin/group", chain.ThenFunc(cmsHandler.DoEditUserGroup)).Methods("POST")
	router.Handle("/cms/admin/group", chain.ThenFunc(cmsHandler.DoDeleteUserGroup)).Methods("DELETE").Queries("id", "{id}")
	router.Handle("/cms/admin/authority", chain.ThenFunc(cmsHandler.GetModuleList)).Methods("GET")
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
	router.Handle("/cms/goods/spec", chain.ThenFunc(cmsHandler.GetGoodsSpecList)).Methods("GET").Queries("id", "{id}")
	router.Handle("/cms/sku/list", chain.ThenFunc(cmsHandler.GetSKUList)).Methods("GET").Queries("page", "{page}").Queries("size", "{size}")
	router.Handle("/cms/sku/{id:[0-9]+}", chain.ThenFunc(cmsHandler.GetSKU)).Methods("GET")
	router.Handle("/cms/sku/edit", chain.ThenFunc(cmsHandler.DoEditSKU)).Methods("POST")
	router.Handle("/cms/sku/{id:[0-9]+}", chain.ThenFunc(cmsHandler.DoDeleteSKU)).Methods("DELETE")
	router.Handle("/cms/coupon/list", chain.ThenFunc(cmsHandler.GetCouponList)).Methods("GET").Queries("page", "{page}").Queries("size", "{size}")
	router.Handle("/cms/coupon/{id:[0-9]+}", chain.ThenFunc(cmsHandler.GetCoupon)).Methods("GET")
	router.Handle("/cms/coupon/edit", chain.ThenFunc(cmsHandler.DoEditCoupon)).Methods("POST")
	router.Handle("/cms/coupon/{id:[0-9]+}", chain.ThenFunc(cmsHandler.DoDeleteCoupon)).Methods("DELETE")
}
