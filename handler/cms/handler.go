package cms

import "wechat-mall-backend/service"

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	handler := &Handler{}
	handler.service = service
	return handler
}
