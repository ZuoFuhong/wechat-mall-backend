package handler

import (
	"net/http"
	"wechat-mall-web/service"
)

type WxappHandler struct {
	service *service.Service
}

func NewWxappHandler(service *service.Service) *WxappHandler {
	handler := &WxappHandler{}
	handler.service = service
	return handler
}

func (wh *WxappHandler) Login(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	if code == "" {
		panic("缺少code")
	}
	resp, err := wh.service.UserService.LoginCodeAuth(code)
	if err != nil {
		panic(err)
	}

	sendNormalResponse(w, resp)
}
