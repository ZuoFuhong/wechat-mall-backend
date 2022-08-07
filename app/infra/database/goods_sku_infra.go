package database

import (
	"context"
	"gorm.io/gorm"
	"time"
	"wechat-mall-backend/app/domain/entity"
	"wechat-mall-backend/app/domain/repository"
	"wechat-mall-backend/pkg/log"
)

type GoodsSkuRepos struct {
	db *gorm.DB
}

func NewGoodsSkuRepos(db *gorm.DB) repository.IGoodsSkuRepos {
	return &GoodsSkuRepos{
		db: db,
	}
}

func (c *GoodsSkuRepos) GetSKUList(ctx context.Context, title string, goodsId, online, page, size int) ([]*entity.WechatMallSkuDO, int, error) {
	tx := c.db.Where("is_del = 0")
	if goodsId != 0 {
		tx = tx.Where("goods_id = ?", goodsId)
	}
	if title != "" {
		tx = tx.Where("title like ?", "%"+title+"%")
	}
	if online == 0 || online == 1 {
		tx = tx.Where("online = ?", online)
	}
	skuList := make([]*entity.WechatMallSkuDO, 0)
	if err := tx.Offset((page - 1) * size).Limit(size).Find(&skuList).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, 0, err
	}
	var total int64
	if err := tx.Count(&total).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Count failed, err: %v", err)
		return nil, 0, err
	}
	return skuList, int(total), nil
}

func (c *GoodsSkuRepos) AddSKU(ctx context.Context, skuDO *entity.WechatMallSkuDO) (int, error) {
	if err := c.db.Create(skuDO).Error; err != nil {
		log.ErrorContextf(ctx, "call db.AddSKU failed, err: %v", err)
		return 0, err
	}
	return skuDO.ID, nil
}

func (c *GoodsSkuRepos) GetSKUById(ctx context.Context, id int) (*entity.WechatMallSkuDO, error) {
	skuDO := new(entity.WechatMallSkuDO)
	if err := c.db.Where("id = ?", id).Find(skuDO).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	return skuDO, nil
}

func (c *GoodsSkuRepos) GetSKUByCode(ctx context.Context, code string) (*entity.WechatMallSkuDO, error) {
	skuDO := new(entity.WechatMallSkuDO)
	if err := c.db.Where("is_del = 0 AND code = ?", code).Find(skuDO).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	return skuDO, nil
}

func (c *GoodsSkuRepos) UpdateSKUById(ctx context.Context, skuDO *entity.WechatMallSkuDO) error {
	if err := c.db.Where("id = ?", skuDO.ID).Updates(skuDO).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Updates failed, err: %v", err)
		return err
	}
	return nil
}

func (c *GoodsSkuRepos) UpdateSkuStockById(ctx context.Context, id, num int) error {
	if err := c.db.Where("id = ? AND stock >= ?", id, num).Update("update_time = ?", time.Now()).Update("stock = stock - ?", num).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Update failed, err: %v", err)
		return err
	}
	return nil
}

func (c *GoodsSkuRepos) QuerySellOutSKUList(ctx context.Context, page, size int) ([]*entity.WechatMallSkuDO, error) {
	skuList := make([]*entity.WechatMallSkuDO, 0)
	if err := c.db.Where("is_del = 0 AND stock = 0").Offset((page - 1) * size).Limit(size).Find(&skuList).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	return skuList, nil
}

func (c *GoodsSkuRepos) CountSellOutSKUList(ctx context.Context) (int, error) {
	empty := new(entity.WechatMallSkuDO)
	var total int64
	if err := c.db.Table(empty.TableName()).Where("is_del = 0 AND stock = 0").Count(&total).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Count failed, err: %v", err)
		return 0, err
	}
	return int(total), nil
}

func (c *GoodsSkuRepos) AddSkuSpecAttr(ctx context.Context, attrDO *entity.WechatMallSkuSpecAttrDO) error {
	if err := c.db.Create(attrDO).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Create failed, err: %v", err)
		return err
	}
	return nil
}

func (c *GoodsSkuRepos) RemoveRelatedBySkuId(ctx context.Context, skuId int) error {
	attrDO := &entity.WechatMallSkuSpecAttrDO{
		Del:        1,
		UpdateTime: time.Now(),
	}
	if err := c.db.Where("sku_id = ?", skuId).Updates(attrDO).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Updates failed, err: %v", err)
		return err
	}
	return nil
}

func (c *GoodsSkuRepos) CountRelatedByAttrId(ctx context.Context, attrId int) (int, error) {
	empty := new(entity.WechatMallSkuSpecAttrDO)
	var total int64
	if err := c.db.Table(empty.TableName()).Where("is_del = 0 AND attr_id = ?", attrId).Count(&total).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Count failed, err: %v", err)
		return 0, err
	}
	return int(total), nil
}

func (c *GoodsSkuRepos) QuerySpecificationList(ctx context.Context, page, size int) ([]*entity.WechatMallSpecificationDO, error) {
	specList := make([]*entity.WechatMallSpecificationDO, 0)
	if err := c.db.Where("is_del = 0").Offset((page - 1) * size).Limit(size).Find(&specList).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	return specList, nil
}

func (c *GoodsSkuRepos) CountSpecification(ctx context.Context) (int, error) {
	empty := new(entity.WechatMallSpecificationDO)
	var total int64
	if err := c.db.Table(empty.TableName()).Where("is_del = 0").Count(&total).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Count failed, err: %v", err)
		return 0, err
	}
	return int(total), nil
}

func (c *GoodsSkuRepos) AddSpecification(ctx context.Context, specDO *entity.WechatMallSpecificationDO) error {
	if err := c.db.Create(specDO).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Create failed, err: %v", err)
		return err
	}
	return nil
}

func (c *GoodsSkuRepos) QuerySpecificationById(ctx context.Context, id int) (*entity.WechatMallSpecificationDO, error) {
	specDO := new(entity.WechatMallSpecificationDO)
	if err := c.db.Where("id = ?", id).Find(specDO).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	return specDO, nil
}

func (c *GoodsSkuRepos) QuerySpecificationByName(ctx context.Context, name string) (*entity.WechatMallSpecificationDO, error) {
	specDO := new(entity.WechatMallSpecificationDO)
	if err := c.db.Where("is_del = 0 AND name = ?", name).Find(specDO).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	return specDO, nil
}

func (c *GoodsSkuRepos) UpdateSpecificationById(ctx context.Context, specDO *entity.WechatMallSpecificationDO) error {
	if err := c.db.Where("id = ?", specDO.ID).Updates(specDO).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Updates failed, err: %v", err)
		return err
	}
	return nil
}

func (c *GoodsSkuRepos) QuerySpecificationAttrList(ctx context.Context, specId int) ([]*entity.WechatMallSpecificationAttrDO, error) {
	attrList := make([]*entity.WechatMallSpecificationAttrDO, 0)
	if err := c.db.Where("is_del = 0 AND spec_id = ?", specId).Find(&attrList).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	return attrList, nil
}

func (c *GoodsSkuRepos) AddSpecificationAttr(ctx context.Context, attrDO *entity.WechatMallSpecificationAttrDO) error {
	if err := c.db.Create(attrDO).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Create failed, err: %v", err)
		return err
	}
	return nil
}

func (c *GoodsSkuRepos) QuerySpecificationAttrById(ctx context.Context, id int) (*entity.WechatMallSpecificationAttrDO, error) {
	attrDO := new(entity.WechatMallSpecificationAttrDO)
	if err := c.db.Where("id = ?", id).Find(attrDO).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	return attrDO, nil
}

func (c *GoodsSkuRepos) QuerySpecificationAttrByValue(ctx context.Context, name string) (*entity.WechatMallSpecificationAttrDO, error) {
	attrDO := new(entity.WechatMallSpecificationAttrDO)
	if err := c.db.Where("is_del = 0 AND value = ?", name).Find(attrDO).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	return attrDO, nil
}

func (c *GoodsSkuRepos) UpdateSpecificationAttrById(ctx context.Context, attrDO *entity.WechatMallSpecificationAttrDO) error {
	if err := c.db.Where("id = ?", attrDO.ID).Updates(attrDO).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Updates failed, err: %v", err)
		return err
	}
	return nil
}
