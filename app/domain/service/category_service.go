package service

import (
	"context"
	"wechat-mall-backend/app/domain/entity"
	"wechat-mall-backend/app/domain/repository"
)

type ICategoryService interface {
	GetCategoryList(ctx context.Context, pid, page, size int) ([]*entity.WechatMallCategoryDO, int, error)

	GetCategoryById(ctx context.Context, id int) (*entity.WechatMallCategoryDO, error)

	GetCategoryByName(ctx context.Context, name string) (*entity.WechatMallCategoryDO, error)

	AddCategory(ctx context.Context, category *entity.WechatMallCategoryDO) error

	UpdateCategory(ctx context.Context, category *entity.WechatMallCategoryDO) error
}

type categoryService struct {
	repos      repository.ICategoryRepos
	goodsRepos repository.IGoodsRepos
}

func NewCategoryService(repos repository.ICategoryRepos, goodsRepos repository.IGoodsRepos) ICategoryService {
	service := &categoryService{
		repos:      repos,
		goodsRepos: goodsRepos,
	}
	return service
}

func (s *categoryService) GetCategoryList(ctx context.Context, pid, page, size int) ([]*entity.WechatMallCategoryDO, int, error) {
	return s.repos.QueryCategoryList(ctx, pid, page, size)
}

func (s *categoryService) GetCategoryById(ctx context.Context, id int) (*entity.WechatMallCategoryDO, error) {
	return s.repos.QueryCategoryById(ctx, id)
}

func (s *categoryService) GetCategoryByName(ctx context.Context, name string) (*entity.WechatMallCategoryDO, error) {
	return s.repos.QueryCategoryByName(ctx, name)
}

func (s *categoryService) AddCategory(ctx context.Context, category *entity.WechatMallCategoryDO) error {
	return s.repos.AddCategory(ctx, category)
}

func (s *categoryService) UpdateCategory(ctx context.Context, category *entity.WechatMallCategoryDO) error {
	if err := s.repos.UpdateCategoryById(ctx, category); err != nil {
		return err
	}
	return s.syncSubCategoryAndGoodsOnline(ctx, category.ParentID, category.ID, category.Online)
}

// 同步其子分类和商品的上下架状态
func (s *categoryService) syncSubCategoryAndGoodsOnline(ctx context.Context, parentId, categoryId, online int) error {
	if parentId == 0 {
		if err := s.repos.UpdateSubCategoryOnline(ctx, categoryId, online); err != nil {
			return err
		}
		ids, err := s.repos.QuerySubCategoryByParentId(ctx, categoryId)
		if err != nil {
			return err
		}
		for _, id := range ids {
			if err := s.goodsRepos.UpdateCategoryGoodsOnlineStatus(ctx, id, online); err != nil {
				return err
			}
		}
	} else {
		if err := s.goodsRepos.UpdateCategoryGoodsOnlineStatus(ctx, categoryId, online); err != nil {
			return err
		}
	}
	return nil
}
