package database

import (
	"context"
	"gorm.io/gorm"
	"time"
	"wechat-mall-backend/app/domain/entity"
	"wechat-mall-backend/app/domain/repository"
	"wechat-mall-backend/consts"
	"wechat-mall-backend/pkg/log"
)

type GoodsRepos struct {
	db *gorm.DB
}

func NewGoodsRepos(db *gorm.DB) repository.IGoodsRepos {
	return &GoodsRepos{
		db: db,
	}
}

func (c *GoodsRepos) QueryGoodsList(ctx context.Context, keyword, order string, categoryId, online, page, size int) ([]*entity.WechatMallGoodsDO, int, error) {
	tx := c.db.Where("is_del = 0")
	if keyword != "" {
		tx = tx.Where("title LIKE ?", "%"+keyword+"%")
	}
	if categoryId != consts.ALL {
		tx = tx.Where("category_id = ?", categoryId)
	}
	if online != consts.ALL {
		tx = tx.Where("online = ?", online)
	}
	if order != "" {
		tx = tx.Order(order + " DESC")
	}
	goodsList := make([]*entity.WechatMallGoodsDO, 0)
	if err := tx.Offset((page - 1) * size).Limit(size).Find(&goodsList).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, 0, err
	}

	var total int64
	if err := tx.Count(&total).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, 0, err
	}
	return goodsList, int(total), nil
}

func (c *GoodsRepos) AddGoods(ctx context.Context, goods *entity.WechatMallGoodsDO) (int, error) {
	if err := c.db.Create(goods).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Create failed, err: %v", err)
		return 0, err
	}
	return goods.ID, nil
}

func (c *GoodsRepos) QueryGoodsById(ctx context.Context, id int) (*entity.WechatMallGoodsDO, error) {
	goods := new(entity.WechatMallGoodsDO)
	if err := c.db.Where("id = ?", id).Find(goods).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	return goods, nil
}

func (c *GoodsRepos) UpdateGoodsById(ctx context.Context, goods *entity.WechatMallGoodsDO) error {
	if err := c.db.Where("id = ?", goods.ID).Updates(goods).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Updates failed, err: %v", err)
		return err
	}
	return nil
}

func (c *GoodsRepos) CountCategoryGoods(ctx context.Context, categoryId int) (int, error) {
	empty := new(entity.WechatMallGoodsDO)
	var total int64
	if err := c.db.Table(empty.TableName()).Where("is_del = 0 AND category_id = ?", categoryId).Count(&total).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Count failed, err: %v", err)
		return 0, err
	}
	return int(total), nil
}

func (c *GoodsRepos) UpdateCategoryGoodsOnlineStatus(ctx context.Context, categoryId, online int) error {
	goods := new(entity.WechatMallGoodsDO)
	if err := c.db.Table(goods.TableName()).Where("is_del = 0 AND category_id = ?", categoryId).Update("update_time", time.Now()).Update("online", online).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Update failed, err: %v", err)
		return err
	}
	return nil
}

func (c *GoodsRepos) UpdateGoodsSaleNum(ctx context.Context, goodsId, num int) error {
	goods := &entity.WechatMallGoodsDO{
		ID:      goodsId,
		SaleNum: num,
	}
	return c.UpdateGoodsById(ctx, goods)
}

func (c *GoodsRepos) GetGoodsSpecList(ctx context.Context, goodsId int) ([]*entity.WechatMallGoodsSpecDO, error) {
	specList := make([]*entity.WechatMallGoodsSpecDO, 0)
	if err := c.db.Where("is_del = 0 AND goods_id = ?", goodsId).Find(&specList).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	return specList, nil
}

func (c *GoodsRepos) CountGoodsSpecBySpecId(ctx context.Context, specId int) (int, error) {
	empty := new(entity.WechatMallGoodsSpecDO)
	var total int64
	if err := c.db.Table(empty.TableName()).Where("is_del = 0 AND spec_id = ?", specId).Count(&total).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Count failed, err: %v", err)
		return 0, err
	}
	return int(total), nil
}

func (c *GoodsRepos) DeleteGoodsSpec(ctx context.Context, goodsId int) error {
	empty := new(entity.WechatMallGoodsSpecDO)
	if err := c.db.Table(empty.TableName()).Where("goods_id = ?", goodsId).Update("is_del", 1).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Update failed, err: %v", err)
		return err
	}
	return nil
}

func (c *GoodsRepos) AddGoodsSpec(ctx context.Context, specDO *entity.WechatMallGoodsSpecDO) error {
	if err := c.db.Create(specDO).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Create failed, err: %v", err)
		return err
	}
	return nil
}
