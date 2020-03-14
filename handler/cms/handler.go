package cms

import (
	"wechat-mall-backend/env"
	"wechat-mall-backend/service"
)

type Handler struct {
	conf    *env.Conf
	service *service.Service
}

func NewHandler(conf *env.Conf, service *service.Service) *Handler {
	handler := &Handler{}
	handler.conf = conf
	handler.service = service
	return handler
}
