package interfaces

import (
	"context"
	"encoding/json"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
	"wechat-mall-backend/app/domain/entity"
	"wechat-mall-backend/app/domain/view"
	"wechat-mall-backend/consts"
	"wechat-mall-backend/errcode"
	"wechat-mall-backend/pkg/log"
	"wechat-mall-backend/pkg/utils"
)

// Login 小程序-用户登录
func (m *MallHttpServiceImpl) Login(w http.ResponseWriter, r *http.Request) {
	code := mux.Vars(r)["code"]
	if code == "" {
		Error(w, errcode.BadRequestParam, "缺少 code")
		return
	}
	token, userId, err := m.userService.LoginCodeAuth(r.Context(), code)
	if err != nil {
		log.ErrorContextf(r.Context(), "call LoginCodeAuth failed, err: %v", err)
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	go m.recordVisitorRecod(r.Context(), userId, r)
	Ok(w, view.WxappLoginVO{Token: token})
}

// recordVisitorRecod 访客记录
func (m *MallHttpServiceImpl) recordVisitorRecod(ctx context.Context, userId int, r *http.Request) {
	userIP := utils.ReadUserIP(r)
	m.userService.DoAddVisitorRecord(ctx, userId, userIP)
}

// UserInfo 小程序-查询用户信息
func (m *MallHttpServiceImpl) UserInfo(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(consts.ContextKey).(int)
	userDO, err := m.userService.QueryUserInfo(r.Context(), userId)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	userVO := &view.WxappUserInfoVO{
		Uid:      userDO.ID,
		Nickname: userDO.Nickname,
		Avatar:   userDO.Avatar,
		Mobile:   utils.PhoneMark(userDO.Mobile),
	}
	Ok(w, userVO)
}

// AuthPhone 小程序-授权手机号
func (m *MallHttpServiceImpl) AuthPhone(w http.ResponseWriter, r *http.Request) {
	req := new(WxappAuthPhone)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		Error(w, errcode.BadRequestParam, "参数解析失败")
		return
	}
	if err := validator.New().Struct(req); err != nil {
		Error(w, errcode.BadRequestParam, "缺少参数")
		return
	}
	userId := r.Context().Value(consts.ContextKey).(int)
	authorization := r.Header.Get("Authorization")
	accessToken := strings.Split(authorization, " ")[1]
	if err := m.userService.DoWxUserPhoneSignature(r.Context(), userId, accessToken, req.EncryptedData, req.Iv); err != nil {
		log.ErrorContextf(r.Context(), "call DoWxUserPhoneSignature failed, err: %v", err)
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	Ok(w, "ok")
}

// AuthUserInfo 小程序-授权用户信息
func (m *MallHttpServiceImpl) AuthUserInfo(w http.ResponseWriter, r *http.Request) {
	req := new(WxappAuthUserInfoReq)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		Error(w, errcode.BadRequestParam, "参数解析失败")
		return
	}
	if err := validator.New().Struct(req); err != nil {
		Error(w, errcode.BadRequestParam, "缺少参数")
		return
	}
	userId := r.Context().Value(consts.ContextKey).(int)
	userDO := &entity.WechatMallUserDO{
		ID:       userId,
		Nickname: req.NickName,
		Avatar:   req.AvatarUrl,
		Gender:   req.Gender,
		Country:  req.Country,
		Province: req.Province,
		City:     req.City,
	}
	if err := m.userService.DoUserAuthInfo(r.Context(), userDO); err != nil {
		log.ErrorContextf(r.Context(), "call DoUserAuthInfo failed, err: %v", err)
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	Ok(w, "ok")
}

// UserBrowseHistory 用户-历史浏览
func (m *MallHttpServiceImpl) UserBrowseHistory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])
	userId := r.Context().Value(consts.ContextKey).(int)

	recordList, total, err := m.browseService.ListBrowseRecord(r.Context(), userId, page, size)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	data := make(map[string]interface{})
	data["list"] = recordList
	data["total"] = total
	Ok(w, data)
}
