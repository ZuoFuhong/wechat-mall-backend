package portal

import (
	"net/http"
	"wechat-mall-backend/defs"
	"wechat-mall-backend/errs"
)

// 小程序登录
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	if code == "" {
		panic(errs.NewParameterError("缺少code"))
	}
	resp, err := h.service.UserService.LoginCodeAuth(code)
	if err != nil {
		panic(err)
	}
	defs.SendNormalResponse(w, resp)
}

// 授权手机号
func (h *Handler) AuthPhone(w http.ResponseWriter, r *http.Request) {

}

// 授权用户信息
func (h *Handler) AuthUserInfo(w http.ResponseWriter, r *http.Request) {

}
