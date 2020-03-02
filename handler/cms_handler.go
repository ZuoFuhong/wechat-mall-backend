package handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/prometheus/common/log"
	"net/http"
	"wechat-mall-web/dbops"
	"wechat-mall-web/defs"
	"wechat-mall-web/service"
	"wechat-mall-web/utils"
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
	validate := validator.New()
	if err = validate.Struct(loginReq); err != nil {
		panic(err.Error())
	}

	cmsUser, err := h.service.CMSUserService.CMSLoginValidate(loginReq.Username, loginReq.Password)
	if err != nil {
		panic(err.Error())
	}
	accessToken, _ := utils.CreateToken(cmsUser.Id, 7200)
	refreshToken, _ := utils.CreateToken(cmsUser.Id, 7500)

	data, _ := json.Marshal(cmsUser)
	if err = dbops.SetStr(accessToken, string(data), 7200); err != nil {
		log.Error(err)
		panic("redis异常！")
	}

	resp := defs.CMSLoginResp{AccessToken: accessToken, RefreshToken: refreshToken}
	sendNormalResponse(w, resp)
}

func (h *CMSHandler) Register(w http.ResponseWriter, r *http.Request) {
	registerReq := defs.CMSRegisterReq{}
	err := json.NewDecoder(r.Body).Decode(&registerReq)
	if err != nil {
		panic("无法解析请求！")
	}
	validate := validator.New()
	if err = validate.Struct(registerReq); err != nil {
		panic(err.Error())
	}

	err = h.service.CMSUserService.CMSUserRegister(&registerReq)
	if err != nil {
		panic(err.Error())
	}
	code := utils.RandomStr(32)
	data, _ := json.Marshal(registerReq)
	err = dbops.SetStr(dbops.CMSCodePrefix+code, string(data), dbops.CMSCodeExpire)
	if err != nil {
		log.Error(err)
		panic("注册异常！")
	}
	sendNormalResponse(w, fmt.Sprintf("已发送一封验证邮件至%s，请打开它进行验证！", registerReq.Email))
}

func (h *CMSHandler) RegisterActivate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]
	cacheData, err := dbops.GetStr(code)
	if err != nil {
		log.Error(err)
		panic("验证链接失效！")
	}
	registerReq := defs.CMSRegisterReq{}
	_ = json.Unmarshal([]byte(cacheData), &registerReq)
	err = h.service.CMSUserService.AddCMSUser(registerReq.Username, registerReq.Password, registerReq.Email)
	if err != nil {
		log.Error(err)
		panic("验证异常，请稍后重试！")
	}
}
