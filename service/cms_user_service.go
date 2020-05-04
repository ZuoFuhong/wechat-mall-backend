package service

import (
	"wechat-mall-backend/dbops"
	"wechat-mall-backend/defs"
	"wechat-mall-backend/errs"
	"wechat-mall-backend/model"
	"wechat-mall-backend/utils"
)

type ICMSUserService interface {
	CMSLoginValidate(username, password string) *model.WechatMallCMSUserDO
	AddCMSUser(req *model.WechatMallCMSUserDO)
	UpdateCMSUser(userDO *model.WechatMallCMSUserDO)
	GetCMSUserList(page, size int) (*[]model.WechatMallCMSUserDO, int)
	CountGroupUser(groupId int) int
	GetCMSUserById(id int) *model.WechatMallCMSUserDO
	QueryUserGroupList(page, size int) (*[]model.WechatMallUserGroupDO, int)
	QueryUserGroupById(id int) *model.WechatMallUserGroupDO
	QueryUserGroupByName(name string) *model.WechatMallUserGroupDO
	AddUserGroup(group *model.WechatMallUserGroupDO) int
	UpdateUserGroup(group *model.WechatMallUserGroupDO)
	QueryGroupAuths(groupId int) *[]map[string][]defs.ModulePageAuth
	QueryGroupPages(groupId int) []int
	RefreshGroupAuths(groupId int, auths []int)
	GetModuleList() *[]defs.CMSModuleVO
}

type CMSUserService struct {
}

func NewCMSUserService() ICMSUserService {
	service := &CMSUserService{}
	return service
}

func (s *CMSUserService) CMSLoginValidate(username, password string) *model.WechatMallCMSUserDO {
	user, err := dbops.GetCMSUserByUsername(username)
	if err != nil {
		panic(err)
	}
	if user.Id == 0 {
		panic(errs.NewErrorCMSUser("用户不存在！"))
	}
	encrpytStr := utils.Md5Encrpyt(password)
	if user.Password != encrpytStr {
		panic(errs.NewErrorCMSUser("密码错误！"))
	}
	return user
}

func (s *CMSUserService) AddCMSUser(userDO *model.WechatMallCMSUserDO) {
	cmsUserDO, err := dbops.GetCMSUserByUsername(userDO.Username)
	if err != nil {
		panic(err)
	}
	if cmsUserDO.Id != 0 {
		panic(errs.NewErrorCMSUser("用户名已注册！"))
	}
	if userDO.Email != "" {
		cmsUserDO, err = dbops.GetCMSUserByEmail(userDO.Email)
		if err != nil {
			panic(err)
		}
		if cmsUserDO.Id != 0 {
			panic(errs.NewErrorCMSUser("邮箱已注册"))
		}
	}
	if userDO.Mobile != "" {
		cmsUserDO, err := dbops.GetCMSUserByMobile(userDO.Mobile)
		if err != nil {
			panic(err)
		}
		if cmsUserDO.Id != 0 {
			panic(errs.NewErrorCMSUser("手机号已注册"))
		}
	}
	groupDO, err := dbops.QueryUserGroupById(userDO.GroupId)
	if err != nil {
		panic(err)
	}
	if groupDO.Id == defs.ZERO || groupDO.Del == defs.DELETE {
		panic(errs.ErrorGroup)
	}
	err = dbops.AddCMSUser(userDO)
	if err != nil {
		panic(err)
	}
}

func (s *CMSUserService) UpdateCMSUser(userDO *model.WechatMallCMSUserDO) {
	if userDO.Email != "" {
		cmsUserDO, err := dbops.GetCMSUserByEmail(userDO.Email)
		if err != nil {
			panic(err)
		}
		if cmsUserDO.Id != defs.ZERO && cmsUserDO.Id != userDO.Id {
			panic(errs.NewErrorCMSUser("邮箱已注册！"))
		}
	}
	if userDO.Mobile != "" {
		cmsUserDO, err := dbops.GetCMSUserByMobile(userDO.Mobile)
		if err != nil {
			panic(err)
		}
		if cmsUserDO.Id != defs.ZERO && cmsUserDO.Id != userDO.Id {
			panic(errs.NewErrorCMSUser("手机号已注册"))
		}
	}
	if userDO.GroupId != defs.ZERO {
		groupDO, err := dbops.QueryUserGroupById(userDO.GroupId)
		if err != nil {
			panic(err)
		}
		if groupDO.Id == defs.ZERO || groupDO.Del == defs.DELETE {
			panic(errs.ErrorGroup)
		}
	}
	err := dbops.UpdateCMSUserById(userDO)
	if err != nil {
		panic(err)
	}
}

func (s *CMSUserService) GetCMSUserList(page, size int) (*[]model.WechatMallCMSUserDO, int) {
	userDOList, err := dbops.ListCMSUser(page, size)
	if err != nil {
		panic(err)
	}
	total, err := dbops.CountCMSUser()
	if err != nil {
		panic(err)
	}
	return userDOList, total
}

func (s *CMSUserService) CountGroupUser(groupId int) int {
	total, err := dbops.CountGroupUser(groupId)
	if err != nil {
		panic(err)
	}
	return total
}

func (s *CMSUserService) GetCMSUserById(id int) *model.WechatMallCMSUserDO {
	userDO, err := dbops.QueryCMSUser(id)
	if err != nil {
		panic(err)
	}
	return userDO
}

func (s *CMSUserService) QueryUserGroupList(page, size int) (*[]model.WechatMallUserGroupDO, int) {
	groupList, err := dbops.QueryGroupList(page, size)
	if err != nil {
		panic(err)
	}
	total, err := dbops.CountUserCoupon()
	if err != nil {
		panic(err)
	}
	return groupList, total
}

func (s *CMSUserService) QueryUserGroupById(id int) *model.WechatMallUserGroupDO {
	groupDO, err := dbops.QueryUserGroupById(id)
	if err != nil {
		panic(err)
	}
	return groupDO
}

func (s *CMSUserService) QueryUserGroupByName(name string) *model.WechatMallUserGroupDO {
	groupDO, err := dbops.QueryUserGroupByName(name)
	if err != nil {
		panic(err)
	}
	return groupDO
}

func (s *CMSUserService) AddUserGroup(group *model.WechatMallUserGroupDO) int {
	groupId, err := dbops.AddUserGroup(group)
	if err != nil {
		panic(err)
	}
	return int(groupId)
}

func (s *CMSUserService) UpdateUserGroup(group *model.WechatMallUserGroupDO) {
	err := dbops.UpdateGroupById(group)
	if err != nil {
		panic(err)
	}
}

func (s *CMSUserService) QueryGroupAuths(groupId int) *[]map[string][]defs.ModulePageAuth {
	permissionList, err := dbops.ListGroupPagePermission(groupId)
	if err != nil {
		panic(err)
	}
	moduleMap := map[int][]int{}
	for _, v := range *permissionList {
		pageDO, err := dbops.QueryModulePageById(v.PageId)
		if err != nil {
			panic(err)
		}
		pageList := moduleMap[pageDO.ModuleId]
		moduleMap[pageDO.ModuleId] = append(pageList, v.PageId)
	}

	auths := []map[string][]defs.ModulePageAuth{}
	for k, v := range moduleMap {
		moduleDO, err := dbops.QueryModuleById(k)
		if err != nil {
			panic(err)
		}
		if moduleDO.Id == 0 || moduleDO.Del == defs.DELETE {
			continue
		}
		authArr := []defs.ModulePageAuth{}
		for _, g := range v {
			pageDO, err := dbops.QueryModulePageById(g)
			if err != nil {
				panic(err)
			}
			if pageDO.Id == 0 || pageDO.Del == defs.DELETE {
				continue
			}
			auth := defs.ModulePageAuth{}
			auth.Module = moduleDO.Name
			auth.Auth = pageDO.Name
			authArr = append(authArr, auth)
		}
		auth := map[string][]defs.ModulePageAuth{}
		auth[moduleDO.Name] = authArr
		auths = append(auths, auth)
	}
	return &auths
}

func (s *CMSUserService) QueryGroupPages(groupId int) []int {
	permissionList, err := dbops.ListGroupPagePermission(groupId)
	if err != nil {
		panic(err)
	}
	auths := []int{}
	for _, v := range *permissionList {
		auths = append(auths, v.PageId)
	}
	return auths
}

func (s *CMSUserService) RefreshGroupAuths(groupId int, auths []int) {
	err := dbops.RemoveGroupAllPagePermission(groupId)
	if err != nil {
		panic(err)
	}
	for _, v := range auths {
		pageDO, err := dbops.QueryModulePageById(v)
		if err != nil {
			panic(err)
		}
		if pageDO.Id == defs.ZERO || pageDO.Del == defs.DELETE {
			panic(errs.ErrorModulePage)
		}
		err = dbops.AddGroupPagePermission(v, groupId)
		if err != nil {
			panic(err)
		}
	}
}

func (s *CMSUserService) GetModuleList() *[]defs.CMSModuleVO {
	moduleList, err := dbops.QueryModuleList()
	if err != nil {
		panic(err)
	}
	moduleVOList := []defs.CMSModuleVO{}
	for _, v := range *moduleList {
		modulePageList, err := dbops.ListModulePage(v.Id)
		if err != nil {
			panic(err)
		}
		pageVOList := []defs.CMSModulePageVO{}
		for _, pv := range *modulePageList {
			pageVO := defs.CMSModulePageVO{}
			pageVO.Id = pv.Id
			pageVO.Name = pv.Name
			pageVO.Description = pv.Description
			pageVOList = append(pageVOList, pageVO)
		}
		moduleVO := defs.CMSModuleVO{}
		moduleVO.Id = v.Id
		moduleVO.Name = v.Name
		moduleVO.Description = v.Description
		moduleVO.PageList = pageVOList
		moduleVOList = append(moduleVOList, moduleVO)
	}
	return &moduleVOList
}
