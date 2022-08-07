package database

import (
	"context"
	"gorm.io/gorm"
	"wechat-mall-backend/app/domain/entity"
	"wechat-mall-backend/app/domain/repository"
	"wechat-mall-backend/pkg/log"
)

type BannerRepos struct {
	db *gorm.DB
}

func NewBannerRepos(db *gorm.DB) repository.IBannerRepos {
	return &BannerRepos{
		db: db,
	}
}

func (b *BannerRepos) QueryBannerList(ctx context.Context, status, page, size int) ([]*entity.WechatMallBannerDO, int, error) {
	banners := make([]*entity.WechatMallBannerDO, 0)
	if err := b.db.Where("status = ? AND is_del = 0", status).Offset((page - 1) * size).Limit(size).Find(&banners).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, 0, err
	}
	empty := new(entity.WechatMallBannerDO)
	var total int64
	if err := b.db.Table(empty.TableName()).Where("is_del = 0 AND status = ?", status).Count(&total).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Count failed, err: %v", err)
		return nil, 0, err
	}
	return banners, int(total), nil
}

func (b *BannerRepos) QueryBannerById(ctx context.Context, id int) (*entity.WechatMallBannerDO, error) {
	bannerDO := new(entity.WechatMallBannerDO)
	if err := b.db.Where("id = ?", id).Find(bannerDO).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	return bannerDO, nil
}

func (b *BannerRepos) AddBanner(ctx context.Context, bannerDO *entity.WechatMallBannerDO) error {
	if err := b.db.Create(bannerDO).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Create failed, err: %v", err)
		return err
	}
	return nil
}

func (b *BannerRepos) UpdateBanner(ctx context.Context, bannerDO *entity.WechatMallBannerDO) error {
	if err := b.db.Where("id = ?", bannerDO.ID).Updates(bannerDO).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Updates failed, err: %v", err)
		return err
	}
	return nil
}
