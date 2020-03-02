package web

import (
	"github.com/gorilla/mux"
	"net/http"
	"wechat-mall-web/dbops"
	"wechat-mall-web/env"
)

type App struct {
	Conf   *env.Conf
	Router *mux.Router
}

func (app *App) Initialize() {
	conf := env.LoadConf()
	app.Conf = conf
	dbops.InitDbConn(conf.Mysql.Username, conf.Mysql.Password, conf.Mysql.Addr)
	dbops.InitRedisCli(conf.Redis.Addr, conf.Redis.Passwd, conf.Redis.Db)
	app.Router = NewRouter(app)
}

func (app *App) Run(addr string) {
	err := http.ListenAndServe(addr, app.Router)
	if err != nil {
		panic(err)
	}
}
