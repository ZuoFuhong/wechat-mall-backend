package repository

import (
	"context"
	"wechat-mall-backend/app/domain/entity"
)

type IGoodsRepos interface {
	QueryGoodsList(ctx context.Context, keyword, order string, categoryId, online, page, size int) ([]*entity.WechatMallGoodsDO, int, error)

	AddGoods(ctx context.Context, goods *entity.WechatMallGoodsDO) (int, error)

	QueryGoodsById(ctx context.Context, id int) (*entity.WechatMallGoodsDO, error)

	UpdateGoodsById(ctx context.Context, goods *entity.WechatMallGoodsDO) error

	CountCategoryGoods(ctx context.Context, categoryId int) (int, error)

	UpdateCategoryGoodsOnlineStatus(ctx context.Context, categoryId, online int) error

	UpdateGoodsSaleNum(ctx context.Context, goodsId, num int) error

	GetGoodsSpecList(ctx context.Context, goodsId int) ([]*entity.WechatMallGoodsSpecDO, error)

	CountGoodsSpecBySpecId(ctx context.Context, specId int) (int, error)

	DeleteGoodsSpec(ctx context.Context, goodsId int) error

	AddGoodsSpec(ctx context.Context, specDO *entity.WechatMallGoodsSpecDO) error
}
