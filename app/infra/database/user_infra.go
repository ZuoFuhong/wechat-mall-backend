package database

import (
	"context"
	"gorm.io/gorm"
	"time"
	"wechat-mall-backend/app/domain/entity"
	"wechat-mall-backend/app/domain/repository"
	"wechat-mall-backend/pkg/log"
)

type UserRepos struct {
	db *gorm.DB
}

func NewUserRepos(db *gorm.DB) repository.IUserRepos {
	return &UserRepos{
		db: db,
	}
}

func (u *UserRepos) GetUserByOpenid(ctx context.Context, openid string) (*entity.WechatMallUserDO, error) {
	userDO := new(entity.WechatMallUserDO)
	if err := u.db.Where("openid = ?", openid).Find(userDO).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	return userDO, nil
}

func (u *UserRepos) GetUserById(ctx context.Context, uid int) (*entity.WechatMallUserDO, error) {
	userDO := new(entity.WechatMallUserDO)
	if err := u.db.Where("id = ?", uid).Find(userDO).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	return userDO, nil
}

func (u *UserRepos) AddUser(ctx context.Context, userDO *entity.WechatMallUserDO) (int, error) {
	if err := u.db.Create(userDO).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Create failed, err: %v", err)
		return 0, err
	}
	return userDO.ID, nil
}

func (u *UserRepos) UpdateUser(ctx context.Context, userDO *entity.WechatMallUserDO) error {
	if err := u.db.Where("id = ?", userDO.ID).Updates(userDO).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Updates failed, err: %v", err)
		return err
	}
	return nil
}

func (u *UserRepos) AddUserAddress(ctx context.Context, address *entity.WechatMallUserAddressDO) error {
	if err := u.db.Create(address).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Create failed, err: %v", err)
		return err
	}
	return nil
}

func (u *UserRepos) ListUserAddress(ctx context.Context, userId, page, size int) ([]*entity.WechatMallUserAddressDO, int, error) {
	addressList := make([]*entity.WechatMallUserAddressDO, 0)
	if err := u.db.Where("is_del = 0 AND user_id = ?", userId).Order("update_time DESC").Offset((page - 1) * size).Limit(size).Find(&addressList).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, 0, err
	}
	empty := new(entity.WechatMallUserAddressDO)
	var total int64
	if err := u.db.Table(empty.TableName()).Where("is_del = 0 AND user_id = ?", userId).Find(&total).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Count failed, err: %v", err)
		return nil, 0, err
	}
	return addressList, int(total), nil
}

func (u *UserRepos) QueryUserAddressById(ctx context.Context, id int) (*entity.WechatMallUserAddressDO, error) {
	address := new(entity.WechatMallUserAddressDO)
	if err := u.db.Table("id = ?", id).Find(address).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	return address, nil
}

func (u *UserRepos) UpdateUserAddress(ctx context.Context, address *entity.WechatMallUserAddressDO) error {
	if err := u.db.Where("id = ?", address.ID).Updates(address).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Updates failed, err: %v", err)
		return err
	}
	return nil
}

func (u *UserRepos) QueryDefaultAddress(ctx context.Context, userId int) (*entity.WechatMallUserAddressDO, error) {
	address := new(entity.WechatMallUserAddressDO)
	if err := u.db.Table("is_del = 0 AND is_default = 1 AND user_id = ?", userId).Find(address).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	return address, nil
}

func (u *UserRepos) AddVisitorRecord(ctx context.Context, userId int, ip string) error {
	record := &entity.WechatMallVisitorRecord{
		UserID:     userId,
		IP:         ip,
		CreateTime: time.Now(),
	}
	if err := u.db.Create(record).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Create failed, err: %v", err)
		return err
	}
	return nil
}

func (u *UserRepos) CountUniqueVisitor(ctx context.Context, startTime, endTime time.Time) (int, error) {
	empty := new(entity.WechatMallVisitorRecord)
	var total int64
	if err := u.db.Table(empty.TableName()).Select("COUNT(DISTINCT(user_id))").Where("create_time BETWEEN ? AND ?", startTime, endTime).Count(&total).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Count failed, err: %v", err)
		return 0, err
	}
	return int(total), nil
}
