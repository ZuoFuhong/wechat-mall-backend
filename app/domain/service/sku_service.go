package service

import (
	"context"
	"encoding/json"
	"wechat-mall-backend/app/domain/entity"
	"wechat-mall-backend/app/domain/repository"
)

type ISKUService interface {
	GetSKUList(ctx context.Context, title string, goodsId, online, page, size int) ([]*entity.WechatMallSkuDO, int, error)

	GetSKUById(ctx context.Context, id int) (*entity.WechatMallSkuDO, error)

	GetSKUByCode(ctx context.Context, code string) (*entity.WechatMallSkuDO, error)

	AddSKU(ctx context.Context, sku *entity.WechatMallSkuDO) error

	UpdateSKUById(ctx context.Context, sku *entity.WechatMallSkuDO) error

	CountSellOutSKU(ctx context.Context) (int, error)

	QuerySellOutSKU(ctx context.Context, page, size int) ([]*entity.WechatMallSkuDO, int, error)

	CountAttrRelatedSku(ctx context.Context, attrId int) (int, error)
}

type skuService struct {
	repos repository.IGoodsSkuRepos
}

func NewSKUService(repos repository.IGoodsSkuRepos) ISKUService {
	service := skuService{
		repos: repos,
	}
	return &service
}

func (s *skuService) GetSKUList(ctx context.Context, title string, goodsId, online, page, size int) ([]*entity.WechatMallSkuDO, int, error) {
	return s.repos.GetSKUList(ctx, title, goodsId, online, page, size)
}

func (s *skuService) GetSKUById(ctx context.Context, id int) (*entity.WechatMallSkuDO, error) {
	return s.repos.GetSKUById(ctx, id)
}

func (s *skuService) GetSKUByCode(ctx context.Context, code string) (*entity.WechatMallSkuDO, error) {
	return s.repos.GetSKUByCode(ctx, code)
}

func (s *skuService) AddSKU(ctx context.Context, sku *entity.WechatMallSkuDO) error {
	skuId, err := s.repos.AddSKU(ctx, sku)
	if err != nil {
		return err
	}
	return s.syncSkuSpecAttrRecord(ctx, skuId, sku.Specs)
}

func (s *skuService) UpdateSKUById(ctx context.Context, sku *entity.WechatMallSkuDO) error {
	err := s.repos.UpdateSKUById(ctx, sku)
	if err != nil {
		return err
	}
	if sku.Del == 1 {
		return s.repos.RemoveRelatedBySkuId(ctx, sku.ID)
	} else {
		return s.syncSkuSpecAttrRecord(ctx, sku.ID, sku.Specs)
	}
}

// 同步-关联SKU属性
func (s *skuService) syncSkuSpecAttrRecord(ctx context.Context, skuId int, specs string) error {
	if err := s.repos.RemoveRelatedBySkuId(ctx, skuId); err != nil {
		return err
	}
	skuSpecList := make([]*entity.SkuSpecs, 0)
	if err := json.Unmarshal([]byte(specs), &skuSpecList); err != nil {
		return err
	}
	for _, v := range skuSpecList {
		attrDO := &entity.WechatMallSkuSpecAttrDO{
			SkuID:  skuId,
			SpecID: v.KeyId,
			AttrID: v.ValueId,
		}
		if err := s.repos.AddSkuSpecAttr(ctx, attrDO); err != nil {
			return err
		}
	}
	return nil
}

func (s *skuService) CountSellOutSKU(ctx context.Context) (int, error) {
	return s.repos.CountSellOutSKUList(ctx)
}

func (s *skuService) QuerySellOutSKU(ctx context.Context, page, size int) ([]*entity.WechatMallSkuDO, int, error) {
	skuList, err := s.repos.QuerySellOutSKUList(ctx, page, size)
	if err != nil {
		return nil, 0, err
	}
	total, err := s.repos.CountSellOutSKUList(ctx)
	if err != nil {
		return nil, 0, err
	}
	return skuList, total, nil
}

func (s *skuService) CountAttrRelatedSku(ctx context.Context, attrId int) (int, error) {
	return s.repos.CountRelatedByAttrId(ctx, attrId)
}
