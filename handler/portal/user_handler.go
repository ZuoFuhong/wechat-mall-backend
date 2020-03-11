package portal

import (
	"net/http"
	"wechat-mall-backend/defs"
	"wechat-mall-backend/errs"
)

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
