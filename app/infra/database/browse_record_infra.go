package database

import (
	"context"
	"gorm.io/gorm"
	"wechat-mall-backend/app/domain/entity"
	"wechat-mall-backend/app/domain/repository"
	"wechat-mall-backend/pkg/log"
)

type BrowseRepos struct {
	db *gorm.DB
}

func NewBrowseRepos(db *gorm.DB) repository.IBrowseRepos {
	return &BrowseRepos{
		db: db,
	}
}

func (c *BrowseRepos) AddBrowseRecord(ctx context.Context, record *entity.WechatMallGoodsBrowseRecord) error {
	if err := c.db.Create(record).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Insert failed, err: %v", err)
		return err
	}
	return nil
}

func (c *BrowseRepos) SelectGoodsBrowse(ctx context.Context, userId, goodsId int) (*entity.WechatMallGoodsBrowseRecord, error) {
	record := new(entity.WechatMallGoodsBrowseRecord)
	if err := c.db.Where("is_del = 0 AND user_id = ? AND goods_id = ?", userId, goodsId).Find(record).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	return record, nil
}

func (c *BrowseRepos) DeleteBrowseRecordById(ctx context.Context, id int) error {
	record := new(entity.WechatMallGoodsBrowseRecord)
	if err := c.db.Table(record.TableName()).Where("id = ?", id).Update("is_del = ?", 1).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Update failed, err: %v", err)
		return err
	}
	return nil
}

func (c *BrowseRepos) SelectGoodsBrowseByUserId(ctx context.Context, userId, page, size int) ([]*entity.WechatMallGoodsBrowseRecord, int, error) {
	records := make([]*entity.WechatMallGoodsBrowseRecord, 0)
	if err := c.db.Where("is_del = 0 AND user_id = ?", userId).Order("update_time DESC").Offset((page - 1) * size).Limit(size).Find(&records).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, 0, err
	}
	empty := new(entity.WechatMallGoodsBrowseRecord)
	var total int64
	if err := c.db.Table(empty.TableName()).Where("is_del = 0 AND user_id = ?", userId).Count(&total).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Count failed, err: %v", err)
		return nil, 0, err
	}
	return records, int(total), nil
}
