package handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/prometheus/common/log"
	"net/http"
	"strings"
	"wechat-mall-backend/dbops"
	"wechat-mall-backend/defs"
	"wechat-mall-backend/errs"
	"wechat-mall-backend/service"
	"wechat-mall-backend/utils"
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
		panic(errs.ErrorRequestBodyParseFailed)
	}
	validate := validator.New()
	if err = validate.Struct(loginReq); err != nil {
		panic(errs.NewParameterError(err.Error()))
	}

	cmsUser, err := h.service.CMSUserService.CMSLoginValidate(loginReq.Username, loginReq.Password)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	accessToken, _ := utils.CreateToken(cmsUser.Id, defs.AccessTokenExpire)
	refreshToken, _ := utils.CreateToken(cmsUser.Id, defs.RefreshTokenExpire)

	resp := defs.CMSTokenResp{AccessToken: accessToken, RefreshToken: refreshToken}
	sendNormalResponse(w, resp)
}

func (h *CMSHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	authorization := r.Header.Get("Authorization")
	if authorization == "" {
		panic(errs.ErrorNotAuthUser)
	}
	tmpArr := strings.Split(authorization, " ")
	if len(tmpArr) != 2 {
		panic(errs.ErrorNotAuthUser)
	}
	refreshToken := tmpArr[1]
	if !utils.ValidateToken(refreshToken) {
		panic(errs.ErrorTokenInvalid)
	}
	payload, err := utils.ParseToken(refreshToken)
	if err != nil {
		panic(errs.ErrorTokenInvalid)
	}
	newAccessToken, _ := utils.CreateToken(payload.Uid, defs.AccessTokenExpire)
	newRefreshToken, _ := utils.CreateToken(payload.Uid, defs.RefreshTokenExpire)

	resp := defs.CMSTokenResp{AccessToken: newAccessToken, RefreshToken: newRefreshToken}
	sendNormalResponse(w, resp)
}

func (h *CMSHandler) Register(w http.ResponseWriter, r *http.Request) {
	registerReq := defs.CMSRegisterReq{}
	err := json.NewDecoder(r.Body).Decode(&registerReq)
	if err != nil {
		panic(errs.ErrorRequestBodyParseFailed)
	}
	validate := validator.New()
	if err = validate.Struct(registerReq); err != nil {
		panic(errs.NewParameterError(err.Error()))
	}

	err = h.service.CMSUserService.CMSUserRegister(&registerReq)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	code := utils.RandomStr(32)
	data, _ := json.Marshal(registerReq)
	err = dbops.SetStr(dbops.CMSCodePrefix+code, string(data), dbops.CMSCodeExpire)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	sendNormalResponse(w, fmt.Sprintf("已发送一封验证邮件至%s，请打开它进行验证！", registerReq.Email))
}

func (h *CMSHandler) RegisterActivate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]
	cacheData, err := dbops.GetStr(code)
	if err != nil {
		panic(errs.ErrorValidateCodeInvalid)
	}
	registerReq := defs.CMSRegisterReq{}
	_ = json.Unmarshal([]byte(cacheData), &registerReq)
	err = h.service.CMSUserService.AddCMSUser(registerReq.Username, registerReq.Password, registerReq.Email)
	if err != nil {
		log.Error(err)
		panic(err)
	}
}
