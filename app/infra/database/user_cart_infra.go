package database

import (
	"context"
	"gorm.io/gorm"
	"wechat-mall-backend/app/domain/entity"
	"wechat-mall-backend/app/domain/repository"
	"wechat-mall-backend/pkg/log"
)

type UserCartRepos struct {
	db *gorm.DB
}

func NewUserCart(db *gorm.DB) repository.IUserCartRepos {
	return &UserCartRepos{
		db: db,
	}
}

func (u *UserCartRepos) QueryCartList(ctx context.Context, userId, page, size int) ([]*entity.WechatMallUserCartDO, int, error) {
	cartList := make([]*entity.WechatMallUserCartDO, 0)
	if err := u.db.Where("is_del = 0 AND user_id = ?", userId).Order("update_time DESC").Offset((page - 1) * size).Limit(size).Find(&cartList).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, 0, err
	}
	empty := new(entity.WechatMallUserCartDO)
	var total int64
	if err := u.db.Table(empty.TableName()).Where("is_del = 0 AND user_id = ?", userId).Count(&total).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Count failed, err: %v", err)
		return nil, 0, err
	}
	return cartList, int(total), nil
}

func (u *UserCartRepos) CoundCartGoodsNum(ctx context.Context, userId int) (int, error) {
	empty := new(entity.WechatMallUserCartDO)
	var total int64
	if err := u.db.Table(empty.TableName()).Where("is_del = 0 AND user_id = ?", userId).Count(&total).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Count failed, err: %v", err)
		return 0, err
	}
	return int(total), nil
}

func (u *UserCartRepos) AddUserCart(ctx context.Context, cartDO *entity.WechatMallUserCartDO) error {
	if err := u.db.Create(cartDO).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Create failed, err: %v", err)
		return err
	}
	return nil
}

func (u *UserCartRepos) QueryCartByParams(ctx context.Context, userId, goodsId, skuId int) (*entity.WechatMallUserCartDO, error) {
	cartDO := new(entity.WechatMallUserCartDO)
	if err := u.db.Where("is_del = 0 AND user_id = ? AND goods_id = ? AND sku_id = ?", userId, goodsId, skuId).Find(cartDO).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	return cartDO, nil
}

func (u *UserCartRepos) UpdateCartById(ctx context.Context, cartDO *entity.WechatMallUserCartDO) error {
	if err := u.db.Where("id = ?", cartDO.ID).Updates(cartDO).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Updates failed, err: %v", err)
		return err
	}
	return nil
}

func (u *UserCartRepos) SelectCartById(ctx context.Context, id int) (*entity.WechatMallUserCartDO, error) {
	cartDO := new(entity.WechatMallUserCartDO)
	if err := u.db.Where("id = ?", id).Find(cartDO).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	return cartDO, nil
}
