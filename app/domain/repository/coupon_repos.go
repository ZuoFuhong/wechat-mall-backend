package repository

import (
	"context"
	"wechat-mall-backend/app/domain/entity"
)

type ICouponRepos interface {
	QueryCouponList(ctx context.Context, page, size, online int) ([]*entity.WechatMallCouponDO, int, error)

	QueryCouponById(ctx context.Context, id int) (*entity.WechatMallCouponDO, error)

	AddCoupon(ctx context.Context, coupon *entity.WechatMallCouponDO) error

	UpdateCouponById(ctx context.Context, coupon *entity.WechatMallCouponDO) error

	QueryCouponLogById(ctx context.Context, couponLogId int) (*entity.WechatMallCouponLogDO, error)

	QueryCouponLogList(ctx context.Context, userId, status, page, size int) ([]*entity.WechatMallCouponLogDO, error)

	CountCouponTakeNum(ctx context.Context, userId, couponId, status int, del int) (int, error)

	UpdateCouponLogById(ctx context.Context, couponLog *entity.WechatMallCouponLogDO) error

	AddCouponLog(ctx context.Context, couponLog *entity.WechatMallCouponLogDO) error

	UpdateCouponLogOverdueStatus(ctx context.Context, userId int) error
}
