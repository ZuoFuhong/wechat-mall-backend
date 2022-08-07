package database

import (
	"context"
	"gorm.io/gorm"
	"wechat-mall-backend/app/domain/entity"
	"wechat-mall-backend/app/domain/repository"
	"wechat-mall-backend/pkg/log"
)

type CmsUserRepos struct {
	db *gorm.DB
}

func NewCmsUserRepos(db *gorm.DB) repository.ICmsUserRepos {
	return &CmsUserRepos{
		db: db,
	}
}

func (c *CmsUserRepos) GetCMSUserByUsername(ctx context.Context, uname string) (*entity.WechatMallCMSUserDO, error) {
	user := new(entity.WechatMallCMSUserDO)
	if err := c.db.Where("is_del = 0 AND username = ?", uname).Find(&user).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	return user, nil
}

func (c *CmsUserRepos) GetCMSUserByMobile(ctx context.Context, mobile string) (*entity.WechatMallCMSUserDO, error) {
	user := new(entity.WechatMallCMSUserDO)
	if err := c.db.Where("is_del = 0 AND mobile = ?", mobile).Find(&user).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	return user, nil
}

func (c *CmsUserRepos) GetCMSUserByEmail(ctx context.Context, email string) (*entity.WechatMallCMSUserDO, error) {
	user := new(entity.WechatMallCMSUserDO)
	if err := c.db.Where("is_del = 0 AND email = ?", email).Find(&user).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	return user, nil
}

func (c *CmsUserRepos) AddCMSUser(ctx context.Context, user *entity.WechatMallCMSUserDO) error {
	if err := c.db.Create(user).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Insert failed, err: %v", err)
		return err
	}
	return nil
}

func (c *CmsUserRepos) CountGroupUser(ctx context.Context, groupId int) (int, error) {
	empty := new(entity.WechatMallCMSUserDO)
	var total int64
	if err := c.db.Table(empty.TableName()).Where("is_del = 0 AND group_id = ?", groupId).Count(&total).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Count failed, err: %v", err)
		return 0, err
	}
	return int(total), nil
}

func (c *CmsUserRepos) QueryCMSUser(ctx context.Context, id int) (*entity.WechatMallCMSUserDO, error) {
	user := new(entity.WechatMallCMSUserDO)
	if err := c.db.Where("id = ?", id).Find(user).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	return user, nil
}

func (c *CmsUserRepos) UpdateCMSUserById(ctx context.Context, user *entity.WechatMallCMSUserDO) error {
	if err := c.db.Where("id = ?", user.ID).Updates(user).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Updates failed, err: %v", err)
		return err
	}
	return nil
}

func (c *CmsUserRepos) ListCMSUser(ctx context.Context, page, size int) ([]*entity.WechatMallCMSUserDO, error) {
	users := make([]*entity.WechatMallCMSUserDO, 0)
	if err := c.db.Where("is_del = 0 AND id != 1").Offset((page - 1) * size).Limit(size).Find(&users).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	return users, nil
}

func (c *CmsUserRepos) CountCMSUser(ctx context.Context) (int, error) {
	empty := new(entity.WechatMallCMSUserDO)
	var total int64
	if err := c.db.Table(empty.TableName()).Where("is_del = 0 AND id != 1").Find(&total).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Count failed, err: %v", err)
		return 0, err
	}
	return int(total), nil
}

func (c *CmsUserRepos) AddUserGroup(ctx context.Context, user *entity.WechatMallUserGroupDO) (int, error) {
	if err := c.db.Create(user).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Insert failed, err: %v", err)
		return 0, err
	}
	return user.ID, nil
}

func (c *CmsUserRepos) QueryUserGroupById(ctx context.Context, id int) (*entity.WechatMallUserGroupDO, error) {
	groupDO := new(entity.WechatMallUserGroupDO)
	if err := c.db.Where("id = ?", id).Find(groupDO).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	return groupDO, nil
}

func (c *CmsUserRepos) QueryUserGroupByName(ctx context.Context, gname string) (*entity.WechatMallUserGroupDO, error) {
	groupDO := new(entity.WechatMallUserGroupDO)
	if err := c.db.Where("is_del = 0 And name = ?", gname).Find(groupDO).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	return groupDO, nil
}

func (c *CmsUserRepos) QueryGroupList(ctx context.Context, page, size int) ([]*entity.WechatMallUserGroupDO, error) {
	groups := make([]*entity.WechatMallUserGroupDO, 0)
	if err := c.db.Where("is_del = 0").Offset((page - 1) * size).Limit(size).Find(&groups).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	return groups, nil
}

func (c *CmsUserRepos) CountUserCoupon(ctx context.Context) (int, error) {
	empty := new(entity.WechatMallUserGroupDO)
	var total int64
	if err := c.db.Table(empty.TableName()).Where("is_del = 0").Count(&total).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Count failed, err: %v", err)
		return 0, err
	}
	return int(total), nil
}

func (c *CmsUserRepos) UpdateGroupById(ctx context.Context, userDO *entity.WechatMallUserGroupDO) error {
	if err := c.db.Where("id = ?", userDO.ID).Updates(userDO).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Updates failed, err: %v", err)
		return err
	}
	return nil
}
