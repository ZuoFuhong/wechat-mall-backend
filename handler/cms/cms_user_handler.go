package cms

import (
	"encoding/json"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"wechat-mall-backend/defs"
	"wechat-mall-backend/errs"
	"wechat-mall-backend/model"
	"wechat-mall-backend/utils"
)

// CMS-用户登录
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	loginReq := defs.CMSLoginReq{}
	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil {
		panic(err)
	}
	validate := validator.New()
	if err = validate.Struct(loginReq); err != nil {
		panic(errs.NewParameterError(err.Error()))
	}

	cmsUser := h.service.CMSUserService.CMSLoginValidate(loginReq.Username, loginReq.Password)
	accessToken, _ := utils.CreateToken(cmsUser.Id, defs.AccessTokenExpire)
	refreshToken, _ := utils.CreateToken(cmsUser.Id, defs.RefreshTokenExpire)

	resp := defs.CMSTokenResp{AccessToken: accessToken, RefreshToken: refreshToken}
	defs.SendNormalResponse(w, resp)
}

// CMS-刷新AccessToken
func (h *Handler) Refresh(w http.ResponseWriter, r *http.Request) {
	authorization := r.Header.Get("Authorization")
	if authorization == "" {
		panic(errs.ErrorRefreshTokenInvalid)
	}
	tmpArr := strings.Split(authorization, " ")
	if len(tmpArr) != 2 {
		panic(errs.ErrorRefreshTokenInvalid)
	}
	refreshToken := tmpArr[1]
	if !utils.ValidateToken(refreshToken) {
		panic(errs.ErrorRefreshTokenInvalid)
	}
	payload, err := utils.ParseToken(refreshToken)
	if err != nil {
		panic(errs.ErrorRefreshTokenInvalid)
	}
	newAccessToken, _ := utils.CreateToken(payload.Uid, defs.AccessTokenExpire)
	newRefreshToken, _ := utils.CreateToken(payload.Uid, defs.RefreshTokenExpire)

	resp := defs.CMSTokenResp{AccessToken: newAccessToken, RefreshToken: newRefreshToken}
	defs.SendNormalResponse(w, resp)
}

// CMS-查询用户信息及权限
func (h *Handler) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(defs.ContextKey).(int)
	userDO := h.service.CMSUserService.GetCMSUserById(userId)
	if userDO.Id == defs.ZERO || userDO.Del == defs.DELETE {
		panic(errs.ErrorCMSUser)
	}
	auths := h.service.CMSUserService.QueryGroupAuths(userDO.GroupId)

	userVO := defs.CMSUserVO{}
	userVO.Id = userDO.Id
	userVO.Username = userDO.Username
	userVO.Email = userDO.Email
	userVO.Mobile = userDO.Mobile
	userVO.Avatar = userDO.Avatar
	userVO.GroupId = userDO.GroupId

	resp := make(map[string]interface{})
	resp["user"] = userVO
	resp["auths"] = auths
	defs.SendNormalResponse(w, resp)
}

// 登录用户-修改密码
func (h *Handler) DoChangePassword(w http.ResponseWriter, r *http.Request) {
	req := defs.CMSChangePasswordReq{}
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
	userDO := h.service.CMSUserService.GetCMSUserById(userId)
	if userDO.Id == defs.ZERO || userDO.Del == defs.DELETE {
		panic(errs.ErrorCMSUser)
	}
	if utils.Md5Encrpyt(req.OldPassword) != userDO.Password {
		panic(errs.NewErrorCMSUser("原始密码错误！"))
	}
	userDO.Password = utils.Md5Encrpyt(req.NewPassword)
	h.service.CMSUserService.UpdateCMSUser(userDO)
	defs.SendNormalResponse(w, "ok")
}

// 查询-用户列表
func (h *Handler) GetUserList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])
	userList, total := h.service.CMSUserService.GetCMSUserList(page, size)
	userVOList := []defs.CMSUserVO{}
	for _, v := range *userList {
		groupDO := h.service.CMSUserService.QueryUserGroupById(v.GroupId)
		if groupDO.Id == defs.ZERO || groupDO.Del == defs.DELETE {
			panic(errs.ErrorGroup)
		}
		userVO := defs.CMSUserVO{}
		userVO.Id = v.Id
		userVO.Username = v.Username
		userVO.Email = v.Email
		userVO.Mobile = v.Mobile
		userVO.Avatar = v.Avatar
		userVO.GroupId = v.GroupId
		userVO.GroupName = groupDO.Name
		userVOList = append(userVOList, userVO)
	}
	resp := make(map[string]interface{})
	resp["list"] = userVOList
	resp["total"] = total
	defs.SendNormalResponse(w, resp)
}

// 查询-单个用户
func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])
	userDO := h.service.CMSUserService.GetCMSUserById(userId)
	if userDO.Id == defs.ZERO || userDO.Del == defs.DELETE {
		panic(errs.ErrorCMSUser)
	}
	if userDO.Id == defs.ADMIN {
		panic(errs.NewErrorCMSUser("权限拒绝"))
	}
	groupDO := h.service.CMSUserService.QueryUserGroupById(userDO.GroupId)
	if groupDO.Id == defs.ZERO || groupDO.Del == defs.DELETE {
		panic(errs.ErrorGroup)
	}
	userVO := defs.CMSUserVO{}
	userVO.Id = userDO.Id
	userVO.Username = userDO.Username
	userVO.Email = userDO.Email
	userVO.Mobile = userDO.Mobile
	userVO.Avatar = userDO.Avatar
	userVO.GroupId = userDO.GroupId
	userVO.GroupName = groupDO.Name
	defs.SendNormalResponse(w, userVO)
}

// 新增/编辑-用户
func (h *Handler) DoEditUser(w http.ResponseWriter, r *http.Request) {
	req := &defs.CMSUserReq{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		panic(errs.ErrorParameterValidate)
	}
	validate := validator.New()
	if err = validate.Struct(req); err != nil {
		panic(errs.NewParameterError(err.Error()))
	}
	matched, _ := regexp.MatchString("^[a-zA-Z0-9]{6,16}$", req.Username)
	if !matched {
		panic(errs.NewParameterError("用户名不符合规范！"))
	}
	matched, _ = regexp.MatchString("^1[358]\\d{9}$", req.Mobile)
	if !matched {
		panic(errs.NewParameterError("请输入正确手机号！"))
	}
	if req.Id == defs.ZERO {
		cmsUserDO := model.WechatMallCMSUserDO{}
		cmsUserDO.Username = req.Username
		cmsUserDO.Password = utils.Md5Encrpyt(req.Mobile[6:])
		cmsUserDO.Email = req.Email
		cmsUserDO.Mobile = req.Mobile
		cmsUserDO.Avatar = req.Avatar
		cmsUserDO.GroupId = req.GroupId
		h.service.CMSUserService.AddCMSUser(&cmsUserDO)
	} else {
		cmsUserDO := h.service.CMSUserService.GetCMSUserById(req.Id)
		if cmsUserDO.Id == defs.ZERO || cmsUserDO.Del == defs.DELETE {
			panic(errs.ErrorCMSUser)
		}
		if cmsUserDO.Id == defs.ADMIN {
			panic(errs.NewErrorCMSUser("权限拒绝！"))
		}
		cmsUserDO.Avatar = req.Avatar
		cmsUserDO.Email = req.Email
		cmsUserDO.Mobile = req.Mobile
		cmsUserDO.GroupId = req.GroupId
		h.service.CMSUserService.UpdateCMSUser(cmsUserDO)
	}
	defs.SendNormalResponse(w, "ok")
}

// 重置密码（supper权限）
func (h *Handler) DoResetCMSUserPassword(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(defs.ContextKey).(int)
	if userId != defs.ADMIN {
		panic(errs.NewErrorCMSUser("权限拒绝"))
	}
	req := defs.CMSResetUserPasswdReq{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		panic(errs.ErrorParameterValidate)
	}
	validate := validator.New()
	if err = validate.Struct(req); err != nil {
		panic(errs.NewParameterError(err.Error()))
	}
	userDO := h.service.CMSUserService.GetCMSUserById(req.UserId)
	if userDO.Id == defs.ZERO || userDO.Del == defs.DELETE {
		panic(errs.ErrorCMSUser)
	}
	matched, _ := regexp.MatchString("^[a-zA-Z0-9]{6,16}$", req.Password)
	if !matched {
		panic(errs.NewParameterError("密码不符合规范！"))
	}
	userDO.Password = utils.Md5Encrpyt(req.Password)
	h.service.CMSUserService.UpdateCMSUser(userDO)
	defs.SendNormalResponse(w, "ok")
}

// 删除-用户
func (h *Handler) DoDeleteCMSUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])
	cmsUserDO := h.service.CMSUserService.GetCMSUserById(userId)
	if cmsUserDO.Id == defs.ZERO || cmsUserDO.Del == defs.DELETE {
		panic(errs.ErrorCMSUser)
	}
	if cmsUserDO.Id == defs.ADMIN {
		panic(errs.NewErrorCMSUser("权限拒绝"))
	}
	cmsUserDO.Del = defs.DELETE
	h.service.CMSUserService.UpdateCMSUser(cmsUserDO)
	defs.SendNormalResponse(w, "ok")
}

// 查询-用户分组
func (h *Handler) GetUserGroupList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])

	groupList, total := h.service.CMSUserService.QueryUserGroupList(page, size)
	groupVOList := []defs.CMSUserGroupVO{}
	for _, v := range *groupList {
		auths := h.service.CMSUserService.QueryGroupAuths(v.Id)
		groupVO := defs.CMSUserGroupVO{}
		groupVO.Id = v.Id
		groupVO.Name = v.Name
		groupVO.Description = v.Description
		groupVO.Auths = auths
		groupVOList = append(groupVOList, groupVO)
	}
	resp := make(map[string]interface{})
	resp["list"] = groupVOList
	resp["total"] = total
	defs.SendNormalResponse(w, resp)
}

// 查询-单个分组
func (h *Handler) GetUserGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupId, _ := strconv.Atoi(vars["id"])
	groupDO := h.service.CMSUserService.QueryUserGroupById(groupId)
	if groupDO.Id == defs.ZERO || groupDO.Del == defs.DELETE {
		panic(errs.ErrorGroup)
	}
	auths := h.service.CMSUserService.QueryGroupAuths(groupId)
	groupVO := defs.CMSUserGroupVO{}
	groupVO.Id = groupDO.Id
	groupVO.Name = groupDO.Name
	groupVO.Description = groupDO.Description
	groupVO.Auths = auths
	defs.SendNormalResponse(w, groupVO)
}

// 新增/编辑 用户分组
func (h *Handler) DoEditUserGroup(w http.ResponseWriter, r *http.Request) {
	req := defs.CMSUserGroupReq{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		panic(errs.ErrorParameterValidate)
	}
	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		panic(errs.NewParameterError(err.Error()))
	}
	if req.Id == defs.ZERO {
		groupDO := h.service.CMSUserService.QueryUserGroupByName(req.Name)
		if groupDO.Id != defs.ZERO {
			panic(errs.NewErrorGroup("分组名已存在！"))
		}
		groupDO.Name = req.Name
		groupDO.Description = req.Description
		groupId := h.service.CMSUserService.AddUserGroup(groupDO)
		h.service.CMSUserService.RefreshGroupAuths(groupId, req.Auths)
	} else {
		groupDO := h.service.CMSUserService.QueryUserGroupByName(req.Name)
		if groupDO.Id != defs.ZERO && groupDO.Id != req.Id {
			panic(errs.NewErrorGroup("分组名已存在！"))
		}
		groupDO = h.service.CMSUserService.QueryUserGroupById(req.Id)
		if groupDO.Id == defs.ZERO || groupDO.Del == defs.DELETE {
			panic(errs.ErrorGroup)
		}
		groupDO.Name = req.Name
		groupDO.Description = req.Description
		h.service.CMSUserService.UpdateUserGroup(groupDO)
		h.service.CMSUserService.RefreshGroupAuths(req.Id, req.Auths)
	}
	defs.SendNormalResponse(w, "ok")
}

// 删除-用户分组
func (h *Handler) DoDeleteUserGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupId, _ := strconv.Atoi(vars["id"])
	groupDO := h.service.CMSUserService.QueryUserGroupById(groupId)
	if groupDO.Id == defs.ZERO || groupDO.Del == defs.DELETE {
		panic(errs.ErrorGroup)
	}
	num := h.service.CMSUserService.CountGroupUser(groupId)
	if num > 0 {
		panic(errs.NewErrorGroup("分组中有用户，禁止删除！"))
	}
	groupDO.Del = defs.DELETE
	h.service.CMSUserService.UpdateUserGroup(groupDO)
	defs.SendNormalResponse(w, "ok")
}

// 查询-权限模块
func (h *Handler) GetModuleList(w http.ResponseWriter, r *http.Request) {
	moduleList := h.service.CMSUserService.GetModuleList()
	defs.SendNormalResponse(w, moduleList)
}
