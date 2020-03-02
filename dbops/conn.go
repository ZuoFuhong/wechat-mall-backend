package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"wechat-mall-web/model"
)

type ID model.ID
type WxappUser model.WxappUser
type CMSUser model.CMSUser

var dbConn *sql.DB

func InitDbConn(username, password, addr string) {
	conn, err := sql.Open("mysql", username+":"+password+"@tcp("+addr+")/wechat_mall")
	if err != nil {
		panic("Connect to mysql error：" + err.Error())
	}
	dbConn = conn
}
