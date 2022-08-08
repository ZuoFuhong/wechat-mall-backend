package service

import (
	"context"
	"wechat-mall-backend/app/domain/entity"
	"wechat-mall-backend/app/domain/repository"
)

type ISpecificationService interface {
	GetSpecificationList(ctx context.Context, page, size int) ([]*entity.WechatMallSpecificationDO, int, error)

	GetSpecificationById(ctx context.Context, id int) (*entity.WechatMallSpecificationDO, error)

	GetSpecificationByName(ctx context.Context, name string) (*entity.WechatMallSpecificationDO, error)

	UpdateSpecificationById(ctx context.Context, spec *entity.WechatMallSpecificationDO) error

	AddSpecification(ctx context.Context, spec *entity.WechatMallSpecificationDO) error

	GetSpecificationAttrList(ctx context.Context, specId int) ([]*entity.WechatMallSpecificationAttrDO, error)

	GetSpecificationAttrById(ctx context.Context, id int) (*entity.WechatMallSpecificationAttrDO, error)

	GetSpecificationAttrByValue(ctx context.Context, value string) (*entity.WechatMallSpecificationAttrDO, error)

	UpdateSpecificationAttrById(ctx context.Context, spec *entity.WechatMallSpecificationAttrDO) error

	AddSpecificationAttr(ctx context.Context, spec *entity.WechatMallSpecificationAttrDO) error
}

type specificationService struct {
	repos repository.IGoodsSkuRepos
}

func NewSpecificationService(repos repository.IGoodsSkuRepos) ISpecificationService {
	service := specificationService{
		repos: repos,
	}
	return &service
}

func (s *specificationService) GetSpecificationList(ctx context.Context, page, size int) ([]*entity.WechatMallSpecificationDO, int, error) {
	specList, err := s.repos.QuerySpecificationList(ctx, page, size)
	if err != nil {
		return nil, 0, err
	}
	total, err := s.repos.CountSpecification(ctx)
	if err != nil {
		return nil, 0, err
	}
	return specList, total, nil
}

func (s *specificationService) GetSpecificationById(ctx context.Context, id int) (*entity.WechatMallSpecificationDO, error) {
	return s.repos.QuerySpecificationById(ctx, id)
}

func (s *specificationService) GetSpecificationByName(ctx context.Context, name string) (*entity.WechatMallSpecificationDO, error) {
	return s.repos.QuerySpecificationByName(ctx, name)
}

func (s *specificationService) UpdateSpecificationById(ctx context.Context, spec *entity.WechatMallSpecificationDO) error {
	return s.repos.UpdateSpecificationById(ctx, spec)
}

func (s *specificationService) AddSpecification(ctx context.Context, spec *entity.WechatMallSpecificationDO) error {
	return s.repos.AddSpecification(ctx, spec)
}

func (s *specificationService) GetSpecificationAttrList(ctx context.Context, specId int) ([]*entity.WechatMallSpecificationAttrDO, error) {
	return s.repos.QuerySpecificationAttrList(ctx, specId)
}

func (s *specificationService) GetSpecificationAttrById(ctx context.Context, id int) (*entity.WechatMallSpecificationAttrDO, error) {
	return s.repos.QuerySpecificationAttrById(ctx, id)
}

func (s *specificationService) GetSpecificationAttrByValue(ctx context.Context, value string) (*entity.WechatMallSpecificationAttrDO, error) {
	return s.repos.QuerySpecificationAttrByValue(ctx, value)
}

func (s *specificationService) UpdateSpecificationAttrById(ctx context.Context, attr *entity.WechatMallSpecificationAttrDO) error {
	return s.repos.UpdateSpecificationAttrById(ctx, attr)
}

func (s *specificationService) AddSpecificationAttr(ctx context.Context, attr *entity.WechatMallSpecificationAttrDO) error {
	return s.repos.AddSpecificationAttr(ctx, attr)
}
