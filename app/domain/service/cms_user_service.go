package service

import (
	"context"
	"github.com/pkg/errors"
	"wechat-mall-backend/app/domain/entity"
	"wechat-mall-backend/app/domain/repository"
	"wechat-mall-backend/app/domain/view"
	"wechat-mall-backend/consts"
	"wechat-mall-backend/pkg/utils"
)

type ICMSUserService interface {
	CMSLoginValidate(ctx context.Context, username, password string) (*entity.WechatMallCMSUserDO, error)

	AddCMSUser(ctx context.Context, req *entity.WechatMallCMSUserDO) error

	UpdateCMSUser(ctx context.Context, userDO *entity.WechatMallCMSUserDO) error

	GetCMSUserList(ctx context.Context, page, size int) ([]*entity.WechatMallCMSUserDO, int, error)

	CountGroupUser(ctx context.Context, groupId int) (int, error)

	GetCMSUserById(ctx context.Context, id int) (*entity.WechatMallCMSUserDO, error)

	QueryUserGroupList(ctx context.Context, page, size int) ([]*entity.WechatMallUserGroupDO, int, error)

	QueryUserGroupById(ctx context.Context, id int) (*entity.WechatMallUserGroupDO, error)

	QueryUserGroupByName(ctx context.Context, name string) (*entity.WechatMallUserGroupDO, error)

	AddUserGroup(ctx context.Context, group *entity.WechatMallUserGroupDO) (int, error)

	UpdateUserGroup(ctx context.Context, group *entity.WechatMallUserGroupDO) error

	QueryGroupAuths(ctx context.Context, groupId int) ([]map[string][]*entity.ModulePageAuth, error)

	QueryGroupPages(ctx context.Context, groupId int) ([]int, error)

	RefreshGroupAuths(ctx context.Context, groupId int, auths []int) error

	GetModuleList(ctx context.Context) ([]*view.CMSModuleVO, error)
}

type CMSUserService struct {
	repos       repository.ICmsUserRepos
	moduleRepos repository.ICmsModuleRepos
}

func NewCMSUserService(repos repository.ICmsUserRepos, moduleRepos repository.ICmsModuleRepos) ICMSUserService {
	service := &CMSUserService{
		repos:       repos,
		moduleRepos: moduleRepos,
	}
	return service
}

func (s *CMSUserService) CMSLoginValidate(ctx context.Context, username, password string) (*entity.WechatMallCMSUserDO, error) {
	user, err := s.repos.GetCMSUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	if user.ID == 0 {
		return nil, errors.New("用户不存在")
	}
	encrpytStr := utils.Md5Encrpyt(password)
	if user.Password != encrpytStr {
		return nil, errors.New("密码错误")
	}
	return user, nil
}

func (s *CMSUserService) AddCMSUser(ctx context.Context, userDO *entity.WechatMallCMSUserDO) error {
	cmsUserDO, err := s.repos.GetCMSUserByUsername(ctx, userDO.Username)
	if err != nil {
		return err
	}
	if cmsUserDO.ID != 0 {
		return errors.New("用户名已注册")
	}
	if userDO.Email != "" {
		cmsUserDO, err = s.repos.GetCMSUserByEmail(ctx, userDO.Email)
		if err != nil {
			return err
		}
		if cmsUserDO.ID != 0 {
			return errors.New("邮箱已注册")
		}
	}
	if userDO.Mobile != "" {
		cmsUserDO, err := s.repos.GetCMSUserByMobile(ctx, userDO.Mobile)
		if err != nil {
			return err
		}
		if cmsUserDO.ID != 0 {
			return errors.New("手机号已注册")
		}
	}
	groupDO, err := s.repos.QueryUserGroupById(ctx, userDO.GroupID)
	if err != nil {
		return err
	}
	if groupDO.ID == consts.ZERO || groupDO.Del == consts.DELETE {
		return errors.New("not found user group record")
	}
	return s.repos.AddCMSUser(ctx, userDO)
}

func (s *CMSUserService) UpdateCMSUser(ctx context.Context, userDO *entity.WechatMallCMSUserDO) error {
	if userDO.Email != "" {
		cmsUserDO, err := s.repos.GetCMSUserByEmail(ctx, userDO.Email)
		if err != nil {
			return err
		}
		if cmsUserDO.ID != consts.ZERO && cmsUserDO.ID != userDO.ID {
			return errors.New("邮箱已注册")
		}
	}
	if userDO.Mobile != "" {
		cmsUserDO, err := s.repos.GetCMSUserByMobile(ctx, userDO.Mobile)
		if err != nil {
			return err
		}
		if cmsUserDO.ID != consts.ZERO && cmsUserDO.ID != userDO.ID {
			return errors.New("手机号已注册")
		}
	}
	if userDO.GroupID != consts.ZERO {
		groupDO, err := s.repos.QueryUserGroupById(ctx, userDO.GroupID)
		if err != nil {
			return err
		}
		if groupDO.ID == consts.ZERO || groupDO.Del == consts.DELETE {
			return errors.New("not found user group record")
		}
	}
	if err := s.repos.UpdateCMSUserById(ctx, userDO); err != nil {
		return err
	}
	return nil
}

func (s *CMSUserService) GetCMSUserList(ctx context.Context, page, size int) ([]*entity.WechatMallCMSUserDO, int, error) {
	userDOList, err := s.repos.ListCMSUser(ctx, page, size)
	if err != nil {
		return nil, 0, err
	}
	total, err := s.repos.CountCMSUser(ctx)
	if err != nil {
		return nil, 0, err
	}
	return userDOList, total, nil
}

func (s *CMSUserService) CountGroupUser(ctx context.Context, groupId int) (int, error) {
	return s.repos.CountGroupUser(ctx, groupId)
}

func (s *CMSUserService) GetCMSUserById(ctx context.Context, id int) (*entity.WechatMallCMSUserDO, error) {
	return s.repos.QueryCMSUser(ctx, id)
}

func (s *CMSUserService) QueryUserGroupList(ctx context.Context, page, size int) ([]*entity.WechatMallUserGroupDO, int, error) {
	groupList, err := s.repos.QueryGroupList(ctx, page, size)
	if err != nil {
		return nil, 0, err
	}
	total, err := s.repos.CountUserCoupon(ctx)
	if err != nil {
		return nil, 0, err
	}
	return groupList, total, nil
}

func (s *CMSUserService) QueryUserGroupById(ctx context.Context, id int) (*entity.WechatMallUserGroupDO, error) {
	return s.repos.QueryUserGroupById(ctx, id)
}

func (s *CMSUserService) QueryUserGroupByName(ctx context.Context, name string) (*entity.WechatMallUserGroupDO, error) {
	return s.repos.QueryUserGroupByName(ctx, name)
}

func (s *CMSUserService) AddUserGroup(ctx context.Context, group *entity.WechatMallUserGroupDO) (int, error) {
	return s.repos.AddUserGroup(ctx, group)
}

func (s *CMSUserService) UpdateUserGroup(ctx context.Context, group *entity.WechatMallUserGroupDO) error {
	return s.repos.UpdateGroupById(ctx, group)
}

func (s *CMSUserService) QueryGroupAuths(ctx context.Context, groupId int) ([]map[string][]*entity.ModulePageAuth, error) {
	permissionList, err := s.moduleRepos.ListGroupPagePermission(ctx, groupId)
	if err != nil {
		return nil, err
	}
	moduleMap := map[int][]int{}
	for _, v := range permissionList {
		pageDO, err := s.moduleRepos.QueryModulePageById(ctx, v.PageID)
		if err != nil {
			return nil, err
		}
		pageList := moduleMap[pageDO.ModuleID]
		moduleMap[pageDO.ModuleID] = append(pageList, v.PageID)
	}

	var auths []map[string][]*entity.ModulePageAuth
	for k, v := range moduleMap {
		moduleDO, err := s.moduleRepos.QueryModuleById(ctx, k)
		if err != nil {
			return nil, err
		}
		if moduleDO.ID == 0 || moduleDO.Del == consts.DELETE {
			continue
		}
		authArr := make([]*entity.ModulePageAuth, 0)
		for _, g := range v {
			pageDO, err := s.moduleRepos.QueryModulePageById(ctx, g)
			if err != nil {
				return nil, err
			}
			if pageDO.ID == 0 || pageDO.Del == consts.DELETE {
				continue
			}
			auth := &entity.ModulePageAuth{
				Module: moduleDO.Name,
				Auth:   pageDO.Name,
			}
			authArr = append(authArr, auth)
		}
		auth := map[string][]*entity.ModulePageAuth{}
		auth[moduleDO.Name] = authArr
		auths = append(auths, auth)
	}
	return auths, nil
}

func (s *CMSUserService) QueryGroupPages(ctx context.Context, groupId int) ([]int, error) {
	permissionList, err := s.moduleRepos.ListGroupPagePermission(ctx, groupId)
	if err != nil {
		return []int{}, err
	}
	var auths []int
	for _, v := range permissionList {
		auths = append(auths, v.PageID)
	}
	return auths, nil
}

func (s *CMSUserService) RefreshGroupAuths(ctx context.Context, groupId int, auths []int) error {
	if err := s.moduleRepos.RemoveGroupAllPagePermission(ctx, groupId); err != nil {
		return err
	}
	for _, v := range auths {
		pageDO, err := s.moduleRepos.QueryModulePageById(ctx, v)
		if err != nil {
			return err
		}
		if pageDO.ID == consts.ZERO || pageDO.Del == consts.DELETE {
			return errors.New("not found module record")
		}

		if err = s.moduleRepos.AddGroupPagePermission(ctx, v, groupId); err != nil {
			return err
		}
	}
	return nil
}

func (s *CMSUserService) GetModuleList(ctx context.Context) ([]*view.CMSModuleVO, error) {
	moduleList, err := s.moduleRepos.QueryModuleList(ctx)
	if err != nil {
		return nil, err
	}
	moduleVOList := make([]*view.CMSModuleVO, 0)
	for _, v := range moduleList {
		modulePageList, err := s.moduleRepos.ListModulePage(ctx, v.ID)
		if err != nil {
			return nil, err
		}
		voList := make([]*view.CMSModulePageVO, 0)
		for _, pv := range modulePageList {
			pageVO := &view.CMSModulePageVO{
				Id:          pv.ID,
				Name:        pv.Name,
				Description: pv.Description,
			}
			voList = append(voList, pageVO)
		}
		moduleVO := &view.CMSModuleVO{
			Id:          v.ID,
			Name:        v.Name,
			Description: v.Description,
			PageList:    voList,
		}
		moduleVOList = append(moduleVOList, moduleVO)
	}
	return moduleVOList, nil
}
