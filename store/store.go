package store

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"wechat-mall-web/model"
)

type ID model.ID
type WxappUser model.WxappUser
type CMSUser model.CMSUser

type MySQLStore struct {
	client *sql.DB
}

func NewMySQLStore(username, password, addr string) *MySQLStore {
	db, err := sql.Open("mysql", username+":"+password+"@tcp("+addr+")/wechat_mall")
	if err != nil {
		panic("Connect to mysql error：" + err.Error())
	}
	return &MySQLStore{client: db}
}
