package database

import (
	"context"
	"gorm.io/gorm"
	"time"
	"wechat-mall-backend/app/domain/entity"
	"wechat-mall-backend/app/domain/repository"
	"wechat-mall-backend/pkg/log"
)

type CmsModuleRepos struct {
	db *gorm.DB
}

func NewCmsModuleRepos(db *gorm.DB) repository.ICmsModuleRepos {
	return &CmsModuleRepos{
		db: db,
	}
}

func (c *CmsModuleRepos) QueryModuleList(ctx context.Context) ([]*entity.WechatMallModuleDO, error) {
	modules := make([]*entity.WechatMallModuleDO, 0)
	if err := c.db.Where("is_del = 0").Find(&modules).Find(&modules).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	return modules, nil
}

func (c *CmsModuleRepos) QueryModuleById(ctx context.Context, mid int) (*entity.WechatMallModuleDO, error) {
	module := new(entity.WechatMallModuleDO)
	if err := c.db.Where("id = ?", mid).Find(module).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	return module, nil
}

func (c *CmsModuleRepos) ListModulePage(ctx context.Context, mid int) ([]*entity.WechatMallModulePageDO, error) {
	pageList := make([]*entity.WechatMallModulePageDO, 0)
	if err := c.db.Where("is_del = 0 AND module_id = ?", mid).Find(&pageList).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	return pageList, nil
}

func (c *CmsModuleRepos) QueryModulePageById(ctx context.Context, pageId int) (*entity.WechatMallModulePageDO, error) {
	pageDO := new(entity.WechatMallModulePageDO)
	if err := c.db.Where("id = ?", pageId).Find(pageDO).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	return pageDO, nil
}

func (c *CmsModuleRepos) ListGroupPagePermission(ctx context.Context, groupId int) ([]*entity.WechatMallGroupPagePermission, error) {
	permissions := make([]*entity.WechatMallGroupPagePermission, 0)
	if err := c.db.Where("is_del = 0 AND group_id = ?", groupId).Find(&permissions).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	return permissions, nil
}

func (c *CmsModuleRepos) AddGroupPagePermission(ctx context.Context, pageId, groupId int) error {
	permission := &entity.WechatMallGroupPagePermission{
		PageID:     pageId,
		GroupID:    groupId,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	if err := c.db.Create(permission).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Create failed, err: %v", err)
		return err
	}
	return nil
}

func (c *CmsModuleRepos) RemoveGroupAllPagePermission(ctx context.Context, groupId int) error {
	empty := new(entity.WechatMallGroupPagePermission)
	if err := c.db.Table(empty.TableName()).Where("group_id = ?", groupId).Update("update_time", time.Now()).Update("is_del", 1).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Update failed, err: %v", err)
		return err
	}
	return nil
}
