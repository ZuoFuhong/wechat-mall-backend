package database

import (
	"context"
	"gorm.io/gorm"
	"time"
	"wechat-mall-backend/app/domain/entity"
	"wechat-mall-backend/app/domain/repository"
	"wechat-mall-backend/pkg/log"
)

type CategoryRepos struct {
	db *gorm.DB
}

func NewCategoryRepos(db *gorm.DB) repository.ICategoryRepos {
	return &CategoryRepos{
		db: db,
	}
}

func (c *CategoryRepos) QueryCategoryList(ctx context.Context, pid, page, size int) ([]*entity.WechatMallCategoryDO, int, error) {
	categorys := make([]*entity.WechatMallCategoryDO, 0)
	if err := c.db.Where("is_del = 0 AND parent_id = ?", pid).Offset((page - 1) * size).Limit(size).Find(&categorys).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, 0, err
	}
	empty := new(entity.WechatMallCategoryDO)
	var total int64
	if err := c.db.Table(empty.TableName()).Where("is_del = 0 AND parent_id = ?", pid).Count(&total).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Count failed, err: %v", err)
		return nil, 0, err
	}
	return categorys, int(total), nil
}

func (c *CategoryRepos) QueryCategoryById(ctx context.Context, id int) (*entity.WechatMallCategoryDO, error) {
	categoryDO := new(entity.WechatMallCategoryDO)
	if err := c.db.Where("id = ?", id).Find(categoryDO).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	return categoryDO, nil
}

func (c *CategoryRepos) QueryCategoryByName(ctx context.Context, cname string) (*entity.WechatMallCategoryDO, error) {
	categoryDO := new(entity.WechatMallCategoryDO)
	if err := c.db.Where("is_del = 0 AND name = ?", cname).Find(categoryDO).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	return categoryDO, nil
}

func (c *CategoryRepos) AddCategory(ctx context.Context, categoryDO *entity.WechatMallCategoryDO) error {
	if err := c.db.Create(categoryDO).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Create failed, err: %v", err)
		return err
	}
	return nil
}

func (c *CategoryRepos) UpdateCategoryById(ctx context.Context, categoryDO *entity.WechatMallCategoryDO) error {
	if err := c.db.Where("id = ?", categoryDO.ID).Updates(categoryDO).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Updates failed, err: %v", err)
		return err
	}
	return nil
}

func (c *CategoryRepos) QuerySubCategoryByParentId(ctx context.Context, cid int) ([]int, error) {
	categorys := make([]*entity.WechatMallCategoryDO, 0)
	if err := c.db.Select("id").Where("is_del = 0 AND parent_id = ?", cid).Find(&categorys).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	cids := make([]int, 0)
	for _, cdo := range categorys {
		cids = append(cids, cdo.ID)
	}
	return cids, nil
}

func (c *CategoryRepos) UpdateSubCategoryOnline(ctx context.Context, categoryId, online int) error {
	categoryDO := &entity.WechatMallCategoryDO{
		ID:         categoryId,
		Online:     online,
		UpdateTime: time.Now(),
	}
	return c.UpdateCategoryById(ctx, categoryDO)
}

func (c *CategoryRepos) QueryAllSubCategory(ctx context.Context) ([]*entity.WechatMallCategoryDO, error) {
	categorys := make([]*entity.WechatMallCategoryDO, 0)
	if err := c.db.Where("is_del = 0 AND online = 1 AND parent_id != 0").Find(&categorys).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	return categorys, nil
}
