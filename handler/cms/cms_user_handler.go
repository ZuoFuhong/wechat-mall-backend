package cms

import (
	"encoding/json"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
	"wechat-mall-backend/defs"
	"wechat-mall-backend/errs"
	"wechat-mall-backend/model"
	"wechat-mall-backend/utils"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	loginReq := defs.CMSLoginReq{}
	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil {
		panic(errs.ErrorRequestBodyParseFailed)
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

func (h *Handler) Refresh(w http.ResponseWriter, r *http.Request) {
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
	defs.SendNormalResponse(w, resp)
}

func (h *Handler) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(defs.ContextKey).(int)
	userDO := h.service.CMSUserService.GetCMSUserById(userId)
	if userDO.Id == 0 {
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

func (h *Handler) DoChangePassword(w http.ResponseWriter, r *http.Request) {
	req := defs.CMSChangePasswordReq{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		panic(errs.ErrorParameterValidate)
	}
	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		panic(err.Error())
	}
	userId := r.Context().Value(defs.ContextKey).(int)
	userDO := h.service.CMSUserService.GetCMSUserById(userId)
	if userDO.Id == 0 {
		panic(errs.ErrorCMSUser)
	}
	if utils.Md5Encrpyt(req.OldPassword) != userDO.Password {
		panic(errs.NewErrorCMSUser("oldPassword mistake"))
	}
	userDO.Password = utils.Md5Encrpyt(req.NewPassword)
	h.service.CMSUserService.UpdateCMSUser(userDO)
	defs.SendNormalResponse(w, "ok")
}

func (h *Handler) GetUserList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])
	userList, total := h.service.CMSUserService.GetCMSUserList(page, size)
	userVOList := []defs.CMSUserVO{}
	for _, v := range *userList {
		userVO := defs.CMSUserVO{}
		userVO.Id = v.Id
		userVO.Username = v.Username
		userVO.Email = v.Email
		userVO.Mobile = v.Mobile
		userVO.Avatar = v.Avatar
		userVO.GroupId = v.GroupId
		userVOList = append(userVOList, userVO)
	}
	resp := make(map[string]interface{})
	resp["list"] = userVOList
	resp["total"] = total
	defs.SendNormalResponse(w, resp)
}

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
	if req.Id == 0 {
		cmsUserDO := model.WechatMallCMSUserDO{}
		cmsUserDO.Username = req.Username
		cmsUserDO.Password = utils.Md5Encrpyt(req.Password)
		cmsUserDO.Email = req.Email
		cmsUserDO.Mobile = req.Mobile
		cmsUserDO.Avatar = ""
		cmsUserDO.GroupId = req.GroupId
		h.service.CMSUserService.AddCMSUser(&cmsUserDO)
	} else {
		cmsUserDO := h.service.CMSUserService.GetCMSUserById(req.Id)
		if cmsUserDO.Id == 0 || cmsUserDO.Id == 1 {
			panic(errs.ErrorCMSUser)
		}
		cmsUserDO.Email = req.Email
		cmsUserDO.Mobile = req.Mobile
		cmsUserDO.GroupId = req.GroupId
		if req.Password != "" {
			cmsUserDO.Password = utils.Md5Encrpyt(req.Password)
		}
		h.service.CMSUserService.UpdateCMSUser(cmsUserDO)
	}
	defs.SendNormalResponse(w, "ok")
}

func (h *Handler) DoDeleteCMSUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])
	cmsUserDO := h.service.CMSUserService.GetCMSUserById(userId)
	if cmsUserDO.Id == 0 || cmsUserDO.Id == 1 {
		panic(errs.ErrorCMSUser)
	}
	cmsUserDO.Del = 1
	h.service.CMSUserService.UpdateCMSUser(cmsUserDO)
	defs.SendNormalResponse(w, "ok")
}

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
	if req.Id == 0 {
		groupDO := h.service.CMSUserService.QueryUserGroupByName(req.Name)
		if groupDO.Id != 0 {
			panic(errs.NewErrorGroup("The name already exists"))
		}
		groupDO.Name = req.Name
		groupDO.Description = req.Description
		groupId := h.service.CMSUserService.AddUserGroup(groupDO)
		h.service.CMSUserService.RefreshGroupAuths(groupId, req.Auths)
	} else {
		groupDO := h.service.CMSUserService.QueryUserGroupByName(req.Name)
		if groupDO.Id != 0 && groupDO.Id != req.Id {
			panic(errs.NewErrorGroup("The name already  exists"))
		}
		groupDO = h.service.CMSUserService.QueryUserGroupById(req.Id)
		if groupDO.Id == 0 {
			panic(errs.ErrorGroup)
		}
		groupDO.Name = req.Name
		groupDO.Description = req.Description
		h.service.CMSUserService.UpdateUserGroup(groupDO)
		h.service.CMSUserService.RefreshGroupAuths(req.Id, req.Auths)
	}
	defs.SendNormalResponse(w, "ok")
}

func (h *Handler) DoDeleteUserGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupId, _ := strconv.Atoi(vars["id"])
	groupDO := h.service.CMSUserService.QueryUserGroupById(groupId)
	if groupDO.Id == 0 {
		panic(errs.ErrorGroup)
	}
	num := h.service.CMSUserService.CountGroupUser(groupId)
	if num > 0 {
		panic(errs.NewErrorGroup("There are users in the group, Do not delete the group"))
	}
	groupDO.Del = 1
	h.service.CMSUserService.UpdateUserGroup(groupDO)
	defs.SendNormalResponse(w, "ok")
}

func (h *Handler) GetModuleList(w http.ResponseWriter, r *http.Request) {
	moduleList := h.service.CMSUserService.GetModuleList()
	defs.SendNormalResponse(w, moduleList)
}
