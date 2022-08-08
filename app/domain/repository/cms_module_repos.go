package repository

import (
	"context"
	"wechat-mall-backend/app/domain/entity"
)

type ICmsModuleRepos interface {
	QueryModuleList(ctx context.Context) ([]*entity.WechatMallModuleDO, error)

	QueryModuleById(ctx context.Context, mid int) (*entity.WechatMallModuleDO, error)

	ListModulePage(ctx context.Context, mid int) ([]*entity.WechatMallModulePageDO, error)

	QueryModulePageById(ctx context.Context, pageId int) (*entity.WechatMallModulePageDO, error)

	ListGroupPagePermission(ctx context.Context, groupId int) ([]*entity.WechatMallGroupPagePermission, error)

	AddGroupPagePermission(ctx context.Context, pageId, groupId int) error

	RemoveGroupAllPagePermission(ctx context.Context, groupId int) error
}
