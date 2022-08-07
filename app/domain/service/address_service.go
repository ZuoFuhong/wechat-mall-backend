package service

import (
	"context"
	"wechat-mall-backend/app/domain/entity"
	"wechat-mall-backend/app/domain/repository"
	"wechat-mall-backend/consts"
)

type IAddressService interface {
	GetAddressList(ctx context.Context, userId, page, size int) ([]*entity.WechatMallUserAddressDO, int, error)

	GetAddress(ctx context.Context, id int) (*entity.WechatMallUserAddressDO, error)

	GetDefaultAddress(ctx context.Context, userId int) (*entity.WechatMallUserAddressDO, error)

	AddAddress(ctx context.Context, address *entity.WechatMallUserAddressDO) error

	UpdateAddress(ctx context.Context, address *entity.WechatMallUserAddressDO) error
}

type addressService struct {
	repos repository.IUserRepos
}

func NewAddressService(repos repository.IUserRepos) IAddressService {
	service := addressService{
		repos: repos,
	}
	return &service
}

func (s *addressService) GetAddressList(ctx context.Context, userId, page, size int) ([]*entity.WechatMallUserAddressDO, int, error) {
	addressList, total, err := s.repos.ListUserAddress(ctx, userId, page, size)
	if err != nil {
		return nil, 0, err
	}
	return addressList, total, nil
}

func (s *addressService) GetAddress(ctx context.Context, id int) (*entity.WechatMallUserAddressDO, error) {
	return s.repos.QueryUserAddressById(ctx, id)
}

func (s *addressService) GetDefaultAddress(ctx context.Context, userId int) (*entity.WechatMallUserAddressDO, error) {
	return s.repos.QueryDefaultAddress(ctx, userId)
}

func (s *addressService) AddAddress(ctx context.Context, address *entity.WechatMallUserAddressDO) error {
	if address.IsDefault == 1 {
		if err := s.clearDefaultAddress(ctx, address.UserID); err != nil {
			return err
		}
	}
	return s.repos.AddUserAddress(ctx, address)
}

func (s *addressService) UpdateAddress(ctx context.Context, address *entity.WechatMallUserAddressDO) error {
	if address.IsDefault == 1 {
		if err := s.clearDefaultAddress(ctx, address.UserID); err != nil {
			return err
		}
	}
	if err := s.repos.UpdateUserAddress(ctx, address); err != nil {
		return err
	}
	return nil
}

func (s *addressService) clearDefaultAddress(ctx context.Context, userId int) error {
	addressDO, err := s.repos.QueryDefaultAddress(ctx, userId)
	if err != nil {
		return err
	}
	if addressDO.ID != consts.ZERO {
		addressDO.IsDefault = 0
		return s.repos.UpdateUserAddress(ctx, addressDO)
	}
	return nil
}
