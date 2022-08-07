package repository

import (
	"context"
	"wechat-mall-backend/app/domain/entity"
)

type IBannerRepos interface {
	QueryBannerList(ctx context.Context, status, page, size int) ([]*entity.WechatMallBannerDO, int, error)

	QueryBannerById(ctx context.Context, bid int) (*entity.WechatMallBannerDO, error)

	AddBanner(ctx context.Context, bannerDO *entity.WechatMallBannerDO) error

	UpdateBanner(ctx context.Context, bannerDO *entity.WechatMallBannerDO) error
}
