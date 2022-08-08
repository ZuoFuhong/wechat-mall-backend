package repository

import (
	"context"
	"wechat-mall-backend/app/domain/entity"
)

type IGridCategoryRepos interface {
	QueryGridCategoryList(ctx context.Context, page, size int) ([]*entity.WechatMallGridCategoryDO, error)

	CountGridCategory(ctx context.Context) (int, error)

	AddGridCategory(ctx context.Context, gridDO *entity.WechatMallGridCategoryDO) error

	QueryGridCategoryById(ctx context.Context, id int) (*entity.WechatMallGridCategoryDO, error)

	QueryGridCategoryByName(ctx context.Context, name string) (*entity.WechatMallGridCategoryDO, error)

	UpdateGridCategoryById(ctx context.Context, gridDO *entity.WechatMallGridCategoryDO) error

	CountGridByCategoryId(ctx context.Context, categoryId int) (int, error)
}
