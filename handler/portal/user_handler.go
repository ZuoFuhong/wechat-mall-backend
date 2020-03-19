package portal

import (
	"encoding/json"
	"github.com/go-playground/validator"
	"github.com/go-redis/redis"
	"net/http"
	"strings"
	"wechat-mall-backend/dbops/rediscli"
	"wechat-mall-backend/defs"
	"wechat-mall-backend/errs"
	"wechat-mall-backend/utils"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	if code == "" {
		panic(errs.NewParameterError("缺少code"))
	}
	token := h.service.UserService.LoginCodeAuth(code)
	// 访客记录
	payload, _ := utils.ParseToken(token)
	userIP := utils.ReadUserIP(r)
	h.service.UserService.DoAddVisitorRecord(payload.Uid, userIP)

	resp := make(map[string]interface{})
	resp["token"] = token
	defs.SendNormalResponse(w, resp)
}

func (h *Handler) AuthPhone(w http.ResponseWriter, r *http.Request) {
	req := defs.WxappAuthPhone{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		panic(err)
	}
	validate := validator.New()
	if err = validate.Struct(req); err != nil {
		panic(errs.NewParameterError(err.Error()))
	}
	userId := r.Context().Value(defs.ContextKey).(int)
	authorization := r.Header.Get("Authorization")
	accessToken := strings.Split(authorization, " ")[1]

	cacheData, err := rediscli.GetStr(defs.MiniappTokenPrefix + accessToken)
	if err == redis.Nil {
		panic(errs.ErrorTokenInvalid)
	}
	if err != nil {
		panic(err)
	}
	if cacheData == "" {
		panic(errs.ErrorTokenInvalid)
	}
	result := make(map[string]interface{})
	err = json.Unmarshal([]byte(cacheData), &result)
	if err != nil {
		panic(err)
	}
	h.service.UserService.DoWxUserPhoneSignature(userId, result["session_key"].(string), req.EncryptedData, req.Iv)
	defs.SendNormalResponse(w, "ok")
}

func (h *Handler) AuthUserInfo(w http.ResponseWriter, r *http.Request) {
	req := defs.WxappAuthUserInfoReq{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		panic(errs.ErrorParameterValidate)
	}
	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		panic(errs.NewParameterError(err.Error()))
	}
	userId := r.Context().Value(defs.ContextKey).(int)

	h.service.UserService.DoUserAuthInfo(userId, req)
	defs.SendNormalResponse(w, "ok")
}
