package web

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"wechat-mall-backend/dbops"
	"wechat-mall-backend/dbops/rediscli"
	"wechat-mall-backend/env"
)

type App struct {
	Conf   *env.Conf
	Router *mux.Router
}

func (app *App) Initialize() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	conf := env.LoadConf()
	app.Conf = conf
	dbops.InitDbConn(conf.Mysql.Username, conf.Mysql.Password, conf.Mysql.Addr)
	rediscli.InitRedisCli(conf.Redis.Addr, conf.Redis.Passwd, conf.Redis.Db)
	app.Router = NewRouter(app)
}

func (app *App) Run() {
	addr := app.Conf.Http.Addr + ":" + app.Conf.Http.Port
	log.Print("Wechat-mall-backend runs on http://" + addr)
	err := http.ListenAndServe(addr, app.Router)
	if err != nil {
		panic(err)
	}
}
