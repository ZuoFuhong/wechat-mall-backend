package service

import (
	"context"
	"wechat-mall-backend/app/domain/entity"
	"wechat-mall-backend/app/domain/repository"
)

type IBannerService interface {
	GetBannerList(ctx context.Context, status, page, size int) ([]*entity.WechatMallBannerDO, int, error)

	GetBannerById(ctx context.Context, id int) (*entity.WechatMallBannerDO, error)

	AddBanner(ctx context.Context, banner *entity.WechatMallBannerDO) error

	UpdateBannerById(ctx context.Context, banner *entity.WechatMallBannerDO) error
}

type bannerService struct {
	repos repository.IBannerRepos
}

func NewBannerService(repos repository.IBannerRepos) IBannerService {
	service := &bannerService{
		repos: repos,
	}
	return service
}

func (s *bannerService) GetBannerList(ctx context.Context, status, page, size int) ([]*entity.WechatMallBannerDO, int, error) {
	return s.repos.QueryBannerList(ctx, status, page, size)
}

func (s *bannerService) GetBannerById(ctx context.Context, id int) (*entity.WechatMallBannerDO, error) {
	return s.repos.QueryBannerById(ctx, id)
}

func (s *bannerService) AddBanner(ctx context.Context, banner *entity.WechatMallBannerDO) error {
	return s.repos.AddBanner(ctx, banner)
}

func (s *bannerService) UpdateBannerById(ctx context.Context, banner *entity.WechatMallBannerDO) error {
	return s.repos.UpdateBanner(ctx, banner)
}
