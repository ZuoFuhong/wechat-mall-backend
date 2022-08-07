package repository

import (
	"context"
	"wechat-mall-backend/app/domain/entity"
)

type IGoodsSkuRepos interface {
	GetSKUList(ctx context.Context, title string, goodsId, online, page, size int) ([]*entity.WechatMallSkuDO, int, error)

	AddSKU(ctx context.Context, skuDO *entity.WechatMallSkuDO) (int, error)

	GetSKUById(ctx context.Context, id int) (*entity.WechatMallSkuDO, error)

	GetSKUByCode(ctx context.Context, code string) (*entity.WechatMallSkuDO, error)

	UpdateSKUById(ctx context.Context, skuDO *entity.WechatMallSkuDO) error

	UpdateSkuStockById(ctx context.Context, id, num int) error

	QuerySellOutSKUList(ctx context.Context, page, size int) ([]*entity.WechatMallSkuDO, error)

	CountSellOutSKUList(ctx context.Context) (int, error)

	AddSkuSpecAttr(ctx context.Context, attrDO *entity.WechatMallSkuSpecAttrDO) error

	RemoveRelatedBySkuId(ctx context.Context, skuId int) error

	CountRelatedByAttrId(ctx context.Context, attrId int) (int, error)

	QuerySpecificationList(ctx context.Context, page, size int) ([]*entity.WechatMallSpecificationDO, error)

	CountSpecification(ctx context.Context) (int, error)

	AddSpecification(ctx context.Context, specDO *entity.WechatMallSpecificationDO) error

	QuerySpecificationById(ctx context.Context, id int) (*entity.WechatMallSpecificationDO, error)

	QuerySpecificationByName(ctx context.Context, name string) (*entity.WechatMallSpecificationDO, error)

	UpdateSpecificationById(ctx context.Context, specDO *entity.WechatMallSpecificationDO) error

	QuerySpecificationAttrList(ctx context.Context, specId int) ([]*entity.WechatMallSpecificationAttrDO, error)

	AddSpecificationAttr(ctx context.Context, attrDO *entity.WechatMallSpecificationAttrDO) error

	QuerySpecificationAttrById(ctx context.Context, id int) (*entity.WechatMallSpecificationAttrDO, error)

	QuerySpecificationAttrByValue(ctx context.Context, name string) (*entity.WechatMallSpecificationAttrDO, error)

	UpdateSpecificationAttrById(ctx context.Context, attrDO *entity.WechatMallSpecificationAttrDO) error
}
