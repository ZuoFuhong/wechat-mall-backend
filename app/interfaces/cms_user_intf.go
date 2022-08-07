package interfaces

import (
	"encoding/json"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"wechat-mall-backend/app/domain/entity"
	"wechat-mall-backend/app/domain/view"
	"wechat-mall-backend/consts"
	"wechat-mall-backend/errcode"
	utils2 "wechat-mall-backend/pkg/utils"
)

// CmsUserLogin CMS-用户登录
func (m *MallHttpServiceImpl) CmsUserLogin(w http.ResponseWriter, r *http.Request) {
	loginReq := new(CMSLoginReq)
	if err := json.NewDecoder(r.Body).Decode(loginReq); err != nil {
		Error(w, errcode.BadRequestParam, "参数解析失败")
		return
	}
	if err := validator.New().Struct(loginReq); err != nil {
		Error(w, errcode.BadRequestParam, "缺少参数")
		return
	}
	cmsUser, err := m.cmsUserService.CMSLoginValidate(r.Context(), loginReq.Username, loginReq.Password)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	accessToken, _ := utils2.CreateToken(cmsUser.ID, consts.AccessTokenExpire)
	refreshToken, _ := utils2.CreateToken(cmsUser.ID, consts.RefreshTokenExpire)
	data := CMSTokenResp{AccessToken: accessToken, RefreshToken: refreshToken}
	Ok(w, data)
}

// Refresh CMS-刷新AccessToken
func (m *MallHttpServiceImpl) Refresh(w http.ResponseWriter, r *http.Request) {
	authorization := r.Header.Get("Authorization")
	if authorization == "" {
		Error(w, errcode.ErrorRefreshTokenInvalid, "刷新Token失败")
		return
	}
	tmpArr := strings.Split(authorization, " ")
	if len(tmpArr) != 2 {
		Error(w, errcode.ErrorRefreshTokenInvalid, "刷新Token失败")
		return
	}
	refreshToken := tmpArr[1]
	if !utils2.ValidateToken(refreshToken) {
		Error(w, errcode.ErrorRefreshTokenInvalid, "刷新Token失败")
		return
	}
	payload, err := utils2.ParseToken(refreshToken)
	if err != nil {
		Error(w, errcode.ErrorRefreshTokenInvalid, "刷新Token失败")
		return
	}
	newAccessToken, _ := utils2.CreateToken(payload.Uid, consts.AccessTokenExpire)
	newRefreshToken, _ := utils2.CreateToken(payload.Uid, consts.RefreshTokenExpire)

	data := CMSTokenResp{AccessToken: newAccessToken, RefreshToken: newRefreshToken}
	Ok(w, data)
}

// GetUserInfo CMS-查询用户信息及权限
func (m *MallHttpServiceImpl) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(consts.ContextKey).(int)
	userDO, err := m.cmsUserService.GetCMSUserById(r.Context(), userId)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	if userDO.ID == consts.ZERO || userDO.Del == consts.DELETE {
		Error(w, errcode.NotFoundCmsUser, "Not found cms user record")
		return
	}
	auths, err := m.cmsUserService.QueryGroupAuths(r.Context(), userDO.GroupID)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	groupDO, err := m.cmsUserService.QueryUserGroupById(r.Context(), userDO.GroupID)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	userVO := &view.CMSUserVO{
		Id:        userDO.ID,
		Username:  userDO.Username,
		Email:     userDO.Email,
		Mobile:    userDO.Mobile,
		Avatar:    userDO.Avatar,
		GroupId:   userDO.GroupID,
		GroupName: groupDO.Name,
	}
	data := make(map[string]interface{})
	data["user"] = userVO
	data["auths"] = auths
	Ok(w, data)
}

// DoChangePassword 登录用户-修改密码
func (m *MallHttpServiceImpl) DoChangePassword(w http.ResponseWriter, r *http.Request) {
	req := new(CMSChangePasswordReq)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		Error(w, errcode.BadRequestParam, "参数解析失败")
		return
	}
	if err := validator.New().Struct(req); err != nil {
		Error(w, errcode.BadRequestParam, "缺少参数")
		return
	}
	userId := r.Context().Value(consts.ContextKey).(int)
	userDO, err := m.cmsUserService.GetCMSUserById(r.Context(), userId)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	if userDO.ID == consts.ZERO || userDO.Del == consts.DELETE {
		Error(w, errcode.NotFoundCmsUser, "Not found cms user record")
		return
	}
	if utils2.Md5Encrpyt(req.OldPassword) != userDO.Password {
		Error(w, errcode.NotAllowOperation, "原始密码错误")
		return
	}
	userDO.Password = utils2.Md5Encrpyt(req.NewPassword)
	if err := m.cmsUserService.UpdateCMSUser(r.Context(), userDO); err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	Ok(w, "ok")
}

// GetUserList 查询-用户列表
func (m *MallHttpServiceImpl) GetUserList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])
	userList, total, err := m.cmsUserService.GetCMSUserList(r.Context(), page, size)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	userVOList := make([]*view.CMSUserVO, 0)
	for _, v := range userList {
		groupDO, err := m.cmsUserService.QueryUserGroupById(r.Context(), v.GroupID)
		if err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
		if groupDO.ID == consts.ZERO || groupDO.Del == consts.DELETE {
			Error(w, errcode.NotFoundUserGroup, "Not found user group record")
			return
		}
		userVO := &view.CMSUserVO{
			Id:        v.ID,
			Username:  v.Username,
			Email:     v.Email,
			Mobile:    v.Mobile,
			Avatar:    v.Avatar,
			GroupId:   v.GroupID,
			GroupName: groupDO.Name,
		}
		userVOList = append(userVOList, userVO)
	}
	data := make(map[string]interface{})
	data["list"] = userVOList
	data["total"] = total
	Ok(w, data)
}

// GetUser 查询-单个用户
func (m *MallHttpServiceImpl) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])
	userDO, err := m.cmsUserService.GetCMSUserById(r.Context(), userId)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	if userDO.ID == consts.ZERO || userDO.Del == consts.DELETE {
		Error(w, errcode.NotFoundCmsUser, "Not found cms user record")
		return
	}
	if userDO.ID == consts.ADMIN {
		Error(w, errcode.NotAllowOperation, "权限拒绝")
		return
	}
	groupDO, err := m.cmsUserService.QueryUserGroupById(r.Context(), userDO.GroupID)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	if groupDO.ID == consts.ZERO || groupDO.Del == consts.DELETE {
		Error(w, errcode.NotFoundUserGroup, "Not found user group record")
		return
	}
	userVO := &view.CMSUserVO{
		Id:        userDO.ID,
		Username:  userDO.Username,
		Email:     userDO.Email,
		Mobile:    userDO.Mobile,
		Avatar:    userDO.Avatar,
		GroupId:   userDO.GroupID,
		GroupName: groupDO.Name,
	}
	Ok(w, userVO)
}

// DoEditUser 新增/编辑-用户
func (m *MallHttpServiceImpl) DoEditUser(w http.ResponseWriter, r *http.Request) {
	req := new(CMSUserReq)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		Error(w, errcode.BadRequestParam, "参数解析失败")
		return
	}
	if err := validator.New().Struct(req); err != nil {
		Error(w, errcode.BadRequestParam, "缺少参数")
		return
	}
	matched, _ := regexp.MatchString("^[a-zA-Z0-9]{5,16}$", req.Username)
	if !matched {
		Error(w, errcode.BadRequestParam, "用户名不符合规范")
		return
	}
	matched, _ = regexp.MatchString("^1[358]\\d{9}$", req.Mobile)
	if !matched {
		Error(w, errcode.BadRequestParam, "请输入正确手机号")
		return
	}
	if req.Id == consts.ZERO {
		cmsUserDO := &entity.WechatMallCMSUserDO{}
		cmsUserDO.Username = req.Username
		cmsUserDO.Password = utils2.Md5Encrpyt(req.Mobile[6:])
		cmsUserDO.Email = req.Email
		cmsUserDO.Mobile = req.Mobile
		cmsUserDO.Avatar = req.Avatar
		cmsUserDO.GroupID = req.GroupId
		if err := m.cmsUserService.AddCMSUser(r.Context(), cmsUserDO); err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
	} else {
		cmsUserDO, err := m.cmsUserService.GetCMSUserById(r.Context(), req.Id)
		if err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
		if cmsUserDO.ID == consts.ZERO || cmsUserDO.Del == consts.DELETE {
			Error(w, errcode.NotFoundCmsUser, "Not found cms user record")
			return
		}
		if cmsUserDO.ID == consts.ADMIN {
			Error(w, errcode.NotAllowOperation, "权限拒绝")
			return
		}
		cmsUserDO.Avatar = req.Avatar
		cmsUserDO.Email = req.Email
		cmsUserDO.Mobile = req.Mobile
		cmsUserDO.GroupID = req.GroupId
		if err := m.cmsUserService.UpdateCMSUser(r.Context(), cmsUserDO); err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
	}
	Ok(w, "ok")
}

// DoResetCMSUserPassword 重置密码（supper权限）
func (m *MallHttpServiceImpl) DoResetCMSUserPassword(w http.ResponseWriter, r *http.Request) {
	req := new(CMSResetUserPasswdReq)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		Error(w, errcode.BadRequestParam, "参数解析失败")
		return
	}
	if err := validator.New().Struct(req); err != nil {
		Error(w, errcode.BadRequestParam, "缺少参数")
		return
	}
	userId := r.Context().Value(consts.ContextKey).(int)
	if userId != consts.ADMIN {
		Error(w, errcode.NotAllowOperation, "权限拒绝")
		return
	}
	userDO, err := m.cmsUserService.GetCMSUserById(r.Context(), req.UserId)
	if err != nil {
		Error(w, errcode.NotAllowOperation, "权限拒绝")
		return
	}
	if userDO.ID == consts.ZERO || userDO.Del == consts.DELETE {
		Error(w, errcode.NotFoundCmsUser, "Not found cms user record")
		return
	}
	matched, _ := regexp.MatchString("^[a-zA-Z0-9]{6,16}$", req.Password)
	if !matched {
		Error(w, errcode.BadRequestParam, "密码不符合规范")
		return
	}
	userDO.Password = utils2.Md5Encrpyt(req.Password)
	if err := m.cmsUserService.UpdateCMSUser(r.Context(), userDO); err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	Ok(w, "ok")
}

// DoDeleteCMSUser 删除-用户
func (m *MallHttpServiceImpl) DoDeleteCMSUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])
	cmsUserDO, err := m.cmsUserService.GetCMSUserById(r.Context(), userId)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	if cmsUserDO.ID == consts.ZERO || cmsUserDO.Del == consts.DELETE {
		Error(w, errcode.NotFoundCmsUser, "Not found cms user record")
		return
	}
	if cmsUserDO.ID == consts.ADMIN {
		Error(w, errcode.NotAllowOperation, "权限拒绝")
		return
	}
	cmsUserDO.Del = consts.DELETE
	if err := m.cmsUserService.UpdateCMSUser(r.Context(), cmsUserDO); err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	Ok(w, "ok")
}

// GetUserGroupList 查询-用户分组
func (m *MallHttpServiceImpl) GetUserGroupList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])

	groupList, total, err := m.cmsUserService.QueryUserGroupList(r.Context(), page, size)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	groupVOList := make([]*view.CMSUserGroupVO, 0)
	for _, v := range groupList {
		auths, err := m.cmsUserService.QueryGroupAuths(r.Context(), v.ID)
		if err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
		groupVO := &view.CMSUserGroupVO{
			Id:          v.ID,
			Name:        v.Name,
			Description: v.Description,
			Auths:       auths,
		}
		groupVOList = append(groupVOList, groupVO)
	}
	data := make(map[string]interface{})
	data["list"] = groupVOList
	data["total"] = total
	Ok(w, data)
}

// GetUserGroup 查询-单个分组
func (m *MallHttpServiceImpl) GetUserGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupId, _ := strconv.Atoi(vars["id"])
	groupDO, err := m.cmsUserService.QueryUserGroupById(r.Context(), groupId)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	if groupDO.ID == consts.ZERO || groupDO.Del == consts.DELETE {
		Error(w, errcode.NotFoundUserGroup, "系统繁忙")
		return
	}
	auths, err := m.cmsUserService.QueryGroupPages(r.Context(), groupId)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	groupVO := &view.CMSUserGroupVO{
		Id:          groupDO.ID,
		Name:        groupDO.Name,
		Description: groupDO.Description,
		Auths:       auths,
	}
	Ok(w, groupVO)
}

// DoEditUserGroup 新增/编辑 用户分组
func (m *MallHttpServiceImpl) DoEditUserGroup(w http.ResponseWriter, r *http.Request) {
	req := new(CMSUserGroupReq)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		Error(w, errcode.BadRequestParam, "参数解析失败")
		return
	}
	if err := validator.New().Struct(req); err != nil {
		Error(w, errcode.BadRequestParam, "缺少参数")
		return
	}
	if req.Id == consts.ZERO {
		groupDO, err := m.cmsUserService.QueryUserGroupByName(r.Context(), req.Name)
		if err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
		if groupDO.ID != consts.ZERO {
			Error(w, errcode.NotAllowOperation, "分组名已存在")
			return
		}
		groupDO.Name = req.Name
		groupDO.Description = req.Description
		groupId, err := m.cmsUserService.AddUserGroup(r.Context(), groupDO)
		if err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
		if err := m.cmsUserService.RefreshGroupAuths(r.Context(), groupId, req.Auths); err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
	} else {
		groupDO, err := m.cmsUserService.QueryUserGroupByName(r.Context(), req.Name)
		if err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
		if groupDO.ID != consts.ZERO && groupDO.ID != req.Id {
			Error(w, errcode.NotAllowOperation, "分组名已存在")
			return
		}
		groupDO, err = m.cmsUserService.QueryUserGroupById(r.Context(), req.Id)
		if err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
		if groupDO.ID == consts.ZERO || groupDO.Del == consts.DELETE {
			Error(w, errcode.NotFoundUserGroup, "Not found user group record")
			return
		}
		groupDO.Name = req.Name
		groupDO.Description = req.Description
		if err := m.cmsUserService.UpdateUserGroup(r.Context(), groupDO); err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
		if err := m.cmsUserService.RefreshGroupAuths(r.Context(), req.Id, req.Auths); err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
	}
	Ok(w, "ok")
}

// DoDeleteUserGroup 删除-用户分组
func (m *MallHttpServiceImpl) DoDeleteUserGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupId, _ := strconv.Atoi(vars["id"])
	groupDO, err := m.cmsUserService.QueryUserGroupById(r.Context(), groupId)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	if groupDO.ID == consts.ZERO || groupDO.Del == consts.DELETE {
		Error(w, errcode.NotFoundUserGroup, "Not found user group record")
		return
	}
	num, err := m.cmsUserService.CountGroupUser(r.Context(), groupId)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	if num > 0 {
		Error(w, errcode.NotAllowOperation, "分组中有用户，禁止删除")
		return
	}
	groupDO.Del = consts.DELETE
	if err := m.cmsUserService.UpdateUserGroup(r.Context(), groupDO); err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	Ok(w, "ok")
}

// GetModuleList 查询-权限模块
func (m *MallHttpServiceImpl) GetModuleList(w http.ResponseWriter, r *http.Request) {
	moduleList, err := m.cmsUserService.GetModuleList(r.Context())
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	Ok(w, moduleList)
}
