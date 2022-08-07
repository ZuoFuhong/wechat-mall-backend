package service

import (
	"context"
	"wechat-mall-backend/app/domain/entity"
	"wechat-mall-backend/app/domain/repository"
)

type IGridCategoryService interface {
	GetGridCategoryList(ctx context.Context, page, size int) ([]*entity.WechatMallGridCategoryDO, int, error)

	GetGridCategoryById(ctx context.Context, id int) (*entity.WechatMallGridCategoryDO, error)

	GetGridCategoryByName(ctx context.Context, name string) (*entity.WechatMallGridCategoryDO, error)

	AddGridCategory(ctx context.Context, gridC *entity.WechatMallGridCategoryDO) error

	UpdateGridCategory(ctx context.Context, gridC *entity.WechatMallGridCategoryDO) error

	CountCategoryBindGrid(ctx context.Context, categoryId int) (int, error)
}

type gridCategoryService struct {
	repos repository.IGridCategoryRepos
}

func NewGridCategoryService(repos repository.IGridCategoryRepos) IGridCategoryService {
	service := &gridCategoryService{
		repos: repos,
	}
	return service
}

func (s *gridCategoryService) GetGridCategoryList(ctx context.Context, page, size int) ([]*entity.WechatMallGridCategoryDO, int, error) {
	gridCList, err := s.repos.QueryGridCategoryList(ctx, page, size)
	if err != nil {
		return nil, 0, err
	}
	total, err := s.repos.CountGridCategory(ctx)
	if err != nil {
		return nil, 0, err
	}
	return gridCList, total, nil
}

func (s *gridCategoryService) GetGridCategoryById(ctx context.Context, id int) (*entity.WechatMallGridCategoryDO, error) {
	return s.repos.QueryGridCategoryById(ctx, id)
}

func (s *gridCategoryService) GetGridCategoryByName(ctx context.Context, name string) (*entity.WechatMallGridCategoryDO, error) {
	return s.repos.QueryGridCategoryByName(ctx, name)
}

func (s *gridCategoryService) AddGridCategory(ctx context.Context, gridC *entity.WechatMallGridCategoryDO) error {
	return s.repos.AddGridCategory(ctx, gridC)
}

func (s *gridCategoryService) UpdateGridCategory(ctx context.Context, gridC *entity.WechatMallGridCategoryDO) error {
	return s.repos.UpdateGridCategoryById(ctx, gridC)
}

func (s *gridCategoryService) CountCategoryBindGrid(ctx context.Context, categoryId int) (int, error) {
	return s.repos.CountGridByCategoryId(ctx, categoryId)
}
