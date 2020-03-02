package service

import (
	"wechat-mall-web/env"
	"wechat-mall-web/model"
)

type WxappUser model.WxappUser
type CMSUser model.CMSUser

type Service struct {
	UserService    IUserService
	CMSUserService ICMSUserService
}

func NewService(conf *env.Conf) *Service {
	userService := NewUserService(conf)
	cmsUserService := NewCMSUserService()
	return &Service{UserService: userService, CMSUserService: cmsUserService}
}
