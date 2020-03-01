package handler

import (
	"encoding/json"
	"net/http"
	"wechat-mall-web/defs"
	"wechat-mall-web/service"
)

type CMSHandler struct {
	service *service.Service
}

func NewCMSHandler(service *service.Service) *CMSHandler {
	handler := &CMSHandler{}
	handler.service = service
	return handler
}

func (h *CMSHandler) Login(w http.ResponseWriter, r *http.Request) {
	loginReq := defs.CMSLoginReq{}
	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil {
		panic("无法解析请求！")
	}

	// todo: 验证密码

	resp := defs.CMSLoginResp{AccessToken: "JjzMJ2pWOTyT30Pf3fQUgNDsEaOvbtvL", RefreshToken: "JjzMJ2pWOTyT30Pf3fQUgNDsEaOvbtvL"}
	sendNormalResponse(w, resp)
}

func (h *CMSHandler) Register(w http.ResponseWriter, r *http.Request) {

}
