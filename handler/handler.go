package handler

import (
	"wechat-mall-web/env"
	"wechat-mall-web/service"
)

type Handler struct {
	WxappHandler *WxappHandler
	CMSHandler   *CMSHandler
}

func NewHandler(conf *env.Conf) *Handler {
	s := service.NewService(conf)
	wxappHandler := NewWxappHandler(s)
	cmsHandler := NewCMSHandler(s)
	return &Handler{wxappHandler, cmsHandler}
}
