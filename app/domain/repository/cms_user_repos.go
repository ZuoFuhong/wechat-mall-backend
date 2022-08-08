package repository

import (
	"context"
	"wechat-mall-backend/app/domain/entity"
)

type ICmsUserRepos interface {
	GetCMSUserByUsername(ctx context.Context, uname string) (*entity.WechatMallCMSUserDO, error)

	GetCMSUserByMobile(ctx context.Context, mobile string) (*entity.WechatMallCMSUserDO, error)

	GetCMSUserByEmail(ctx context.Context, email string) (*entity.WechatMallCMSUserDO, error)

	AddCMSUser(ctx context.Context, user *entity.WechatMallCMSUserDO) error

	CountGroupUser(ctx context.Context, groupId int) (int, error)

	QueryCMSUser(ctx context.Context, id int) (*entity.WechatMallCMSUserDO, error)

	UpdateCMSUserById(ctx context.Context, user *entity.WechatMallCMSUserDO) error

	ListCMSUser(ctx context.Context, page, size int) ([]*entity.WechatMallCMSUserDO, error)

	CountCMSUser(ctx context.Context) (int, error)

	AddUserGroup(ctx context.Context, user *entity.WechatMallUserGroupDO) (int, error)

	QueryUserGroupById(ctx context.Context, id int) (*entity.WechatMallUserGroupDO, error)

	QueryUserGroupByName(ctx context.Context, gname string) (*entity.WechatMallUserGroupDO, error)

	QueryGroupList(ctx context.Context, page, size int) ([]*entity.WechatMallUserGroupDO, error)

	CountUserCoupon(ctx context.Context) (int, error)

	UpdateGroupById(ctx context.Context, userDO *entity.WechatMallUserGroupDO) error
}
