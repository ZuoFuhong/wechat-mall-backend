package web

import (
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"wechat-mall-web/handler"
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
	router.Handle("/api/cms/login", chain.ThenFunc(handler.CMSHandler.Login)).Methods("POST")
	router.Handle("/api/cms/register", chain.ThenFunc(handler.CMSHandler.Register)).Methods("POST")
	router.Handle("/cms/register/activate/{code:[0-9a-zA-Z]+}", chain.ThenFunc(handler.CMSHandler.RegisterActivate)).Methods("GET")
}
