package repository

import (
	"context"
	"time"
	"wechat-mall-backend/app/domain/entity"
)

type IUserRepos interface {
	GetUserByOpenid(ctx context.Context, openid string) (*entity.WechatMallUserDO, error)

	GetUserById(ctx context.Context, uid int) (*entity.WechatMallUserDO, error)

	AddUser(ctx context.Context, userDO *entity.WechatMallUserDO) (int, error)

	UpdateUser(ctx context.Context, userDO *entity.WechatMallUserDO) error

	AddUserAddress(ctx context.Context, address *entity.WechatMallUserAddressDO) error

	ListUserAddress(ctx context.Context, userId, page, size int) ([]*entity.WechatMallUserAddressDO, int, error)

	QueryUserAddressById(ctx context.Context, id int) (*entity.WechatMallUserAddressDO, error)

	UpdateUserAddress(ctx context.Context, address *entity.WechatMallUserAddressDO) error

	QueryDefaultAddress(ctx context.Context, userId int) (*entity.WechatMallUserAddressDO, error)

	AddVisitorRecord(ctx context.Context, userId int, ip string) error

	CountUniqueVisitor(ctx context.Context, startTime, endTime time.Time) (int, error)
}
