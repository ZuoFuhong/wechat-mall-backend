package repository

import (
	"context"
	"wechat-mall-backend/app/domain/entity"
)

type ICategoryRepos interface {
	QueryCategoryList(ctx context.Context, pid, page, size int) ([]*entity.WechatMallCategoryDO, int, error)

	QueryCategoryById(ctx context.Context, id int) (*entity.WechatMallCategoryDO, error)

	QueryCategoryByName(ctx context.Context, cname string) (*entity.WechatMallCategoryDO, error)

	AddCategory(ctx context.Context, categoryDO *entity.WechatMallCategoryDO) error

	UpdateCategoryById(ctx context.Context, categoryDO *entity.WechatMallCategoryDO) error

	QuerySubCategoryByParentId(ctx context.Context, cid int) ([]int, error)

	UpdateSubCategoryOnline(ctx context.Context, categoryId, online int) error

	QueryAllSubCategory(ctx context.Context) ([]*entity.WechatMallCategoryDO, error)
}
