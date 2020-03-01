package web

import (
	"github.com/gorilla/mux"
	"net/http"
	"wechat-mall-web/env"
)

type App struct {
	Conf   *env.Conf
	Router *mux.Router
}

func (app *App) Initialize() {
	app.Conf = env.LoadConf()
	app.Router = NewRouter(app)
}

func (app *App) Run(addr string) {
	err := http.ListenAndServe(addr, app.Router)
	if err != nil {
		panic(err)
	}
}
