package handler

import (
	"wechat-mall-backend/env"
	"wechat-mall-backend/service"
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
