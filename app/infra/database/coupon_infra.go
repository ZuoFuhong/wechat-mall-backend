package database

import (
	"context"
	"gorm.io/gorm"
	"wechat-mall-backend/app/domain/entity"
	"wechat-mall-backend/app/domain/repository"
	"wechat-mall-backend/consts"
	"wechat-mall-backend/pkg/log"
)

type CouponRepos struct {
	db *gorm.DB
}

func NewCouponRepos(db *gorm.DB) repository.ICouponRepos {
	return &CouponRepos{
		db: db,
	}
}

func (c *CouponRepos) QueryCouponList(ctx context.Context, page, size, online int) ([]*entity.WechatMallCouponDO, int, error) {
	coupons := make([]*entity.WechatMallCouponDO, 0)
	tx := c.db.Where("is_del = 0")
	if online != consts.ALL {
		tx = tx.Where("online = ?", online)
	}
	if online == 1 {
		tx = tx.Where("start_time < now() AND end_time > now()")
	}
	if err := tx.Offset((page - 1) * size).Limit(size).Find(&coupons).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, 0, err
	}
	var total int64
	if err := tx.Count(&total).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Count failed, err: %v", err)
		return nil, 0, err
	}
	return coupons, int(total), nil
}

func (c *CouponRepos) QueryCouponById(ctx context.Context, id int) (*entity.WechatMallCouponDO, error) {
	coupon := new(entity.WechatMallCouponDO)
	if err := c.db.Where("id = ?", id).Find(coupon).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	return coupon, nil
}

func (c *CouponRepos) AddCoupon(ctx context.Context, coupon *entity.WechatMallCouponDO) error {
	if err := c.db.Create(coupon).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Create failed, err: %v", err)
		return err
	}
	return nil
}

func (c *CouponRepos) UpdateCouponById(ctx context.Context, coupon *entity.WechatMallCouponDO) error {
	if err := c.db.Where("id = ?", coupon.ID).Updates(coupon).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Updates failed, err: %v", err)
		return err
	}
	return nil
}

func (c *CouponRepos) QueryCouponLogById(ctx context.Context, couponLogId int) (*entity.WechatMallCouponLogDO, error) {
	couponLog := new(entity.WechatMallCouponLogDO)
	if err := c.db.Where("id = ?", couponLogId).Find(couponLog).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	return couponLog, nil
}

func (c *CouponRepos) QueryCouponLogList(ctx context.Context, userId, status, page, size int) ([]*entity.WechatMallCouponLogDO, error) {
	couponLogs := make([]*entity.WechatMallCouponLogDO, 0)
	if err := c.db.Where("is_del = 0 AND user_id = ? AND status = ?", userId, status).Order("update_time DESC").Offset((page - 1) * size).Limit(size).Find(&couponLogs).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	return couponLogs, nil
}

func (c *CouponRepos) CountCouponTakeNum(ctx context.Context, userId, couponId, status int, del int) (int, error) {
	empty := new(entity.WechatMallCouponLogDO)
	tx := c.db.Table(empty.TableName())
	if userId != consts.ALL {
		tx = tx.Where("user_id = ?", userId)
	}
	if couponId != consts.ALL {
		tx = tx.Where("coupon_id = ?", couponId)
	}
	if status != consts.ALL {
		tx = tx.Where("status = ?", status)
	}
	if del != consts.ALL {
		tx = tx.Where("is_del = ?", del)
	}
	var total int64
	if err := tx.Count(&total).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Count failed, err: %v", err)
		return 0, err
	}
	return int(total), nil
}

func (c *CouponRepos) UpdateCouponLogById(ctx context.Context, couponLog *entity.WechatMallCouponLogDO) error {
	if err := c.db.Where("id = ?", couponLog.ID).Updates(couponLog).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Updates failed, err: %v", err)
		return err
	}
	return nil
}

func (c *CouponRepos) AddCouponLog(ctx context.Context, couponLog *entity.WechatMallCouponLogDO) error {
	if err := c.db.Create(couponLog).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Create failed, err: %v", err)
		return err
	}
	return nil
}

func (c *CouponRepos) UpdateCouponLogOverdueStatus(ctx context.Context, userId int) error {
	empty := new(entity.WechatMallCouponLogDO)
	if err := c.db.Table(empty.TableName()).Where("is_del = 0 AND expire_time < now() AND user_id = ?", userId).Update("status", 2).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Update failed, err: %v", err)
		return err
	}
	return nil
}
