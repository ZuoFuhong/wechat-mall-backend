package repository

import (
	"context"
	"wechat-mall-backend/app/domain/entity"
)

type IUserCartRepos interface {
	QueryCartList(ctx context.Context, userId, page, size int) ([]*entity.WechatMallUserCartDO, int, error)

	CoundCartGoodsNum(ctx context.Context, userId int) (int, error)

	AddUserCart(ctx context.Context, cartDO *entity.WechatMallUserCartDO) error

	QueryCartByParams(ctx context.Context, userId, goodsId, skuId int) (*entity.WechatMallUserCartDO, error)

	UpdateCartById(ctx context.Context, cartDO *entity.WechatMallUserCartDO) error

	SelectCartById(ctx context.Context, id int) (*entity.WechatMallUserCartDO, error)
}
