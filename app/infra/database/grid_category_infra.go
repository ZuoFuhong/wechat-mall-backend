package database

import (
	"context"
	"gorm.io/gorm"
	"wechat-mall-backend/app/domain/entity"
	"wechat-mall-backend/app/domain/repository"
	"wechat-mall-backend/pkg/log"
)

type GridCategoryRepos struct {
	db *gorm.DB
}

func NewGridCategoryRepos(db *gorm.DB) repository.IGridCategoryRepos {
	return &GridCategoryRepos{
		db: db,
	}
}

func (c *GridCategoryRepos) QueryGridCategoryList(ctx context.Context, page, size int) ([]*entity.WechatMallGridCategoryDO, error) {
	gridList := make([]*entity.WechatMallGridCategoryDO, 0)
	if err := c.db.Where("is_del = 0").Offset((page - 1) * size).Limit(size).Find(&gridList).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	return gridList, nil
}

func (c *GridCategoryRepos) CountGridCategory(ctx context.Context) (int, error) {
	empty := new(entity.WechatMallGridCategoryDO)
	var total int64
	if err := c.db.Table(empty.TableName()).Where("is_del = 0").Count(&total).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Count failed, err: %v", err)
		return 0, err
	}
	return int(total), nil
}

func (c *GridCategoryRepos) AddGridCategory(ctx context.Context, gridDO *entity.WechatMallGridCategoryDO) error {
	if err := c.db.Create(gridDO).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Create failed, err: %v", err)
		return err
	}
	return nil
}

func (c *GridCategoryRepos) QueryGridCategoryById(ctx context.Context, id int) (*entity.WechatMallGridCategoryDO, error) {
	gridDO := new(entity.WechatMallGridCategoryDO)
	if err := c.db.Where("id = ?", id).Find(gridDO).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	return gridDO, nil
}

func (c *GridCategoryRepos) QueryGridCategoryByName(ctx context.Context, name string) (*entity.WechatMallGridCategoryDO, error) {
	gridDO := new(entity.WechatMallGridCategoryDO)
	if err := c.db.Where("is_del = 0 AND name = ?", name).Find(gridDO).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	return gridDO, nil
}

func (c *GridCategoryRepos) UpdateGridCategoryById(ctx context.Context, gridDO *entity.WechatMallGridCategoryDO) error {
	if err := c.db.Where("id = ?", gridDO.ID).Updates(gridDO).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Updates failed, err: %v", err)
		return err
	}
	return nil
}

func (c *GridCategoryRepos) CountGridByCategoryId(ctx context.Context, categoryId int) (int, error) {
	empty := new(entity.WechatMallGridCategoryDO)
	var total int64
	if err := c.db.Table(empty.TableName()).Where("is_del = 0 AND category_id = ?", categoryId).Count(&total).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Count failed, err: %v", err)
		return 0, err
	}
	return int(total), nil
}
