package service

import (
	"wechat-mall-web/env"
	"wechat-mall-web/model"
	"wechat-mall-web/store"
)

type WxappUser model.WxappUser
type CMSUser model.CMSUser

type Service struct {
	UserService    IUserService
	CMSUserService ICMSUserService
}

func NewService(conf *env.Conf) *Service {
	dbStore := store.NewMySQLStore(conf.Mysql.Username, conf.Mysql.Password, conf.Mysql.Addr)
	redisStore := store.NewRedisStore(conf.Redis.Addr, conf.Redis.Passwd, conf.Redis.Db)

	userService := NewUserService(dbStore, redisStore, conf)
	cmsUserService := NewCMSUserService(dbStore, redisStore)
	return &Service{UserService: userService, CMSUserService: cmsUserService}
}
