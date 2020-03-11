package service

import (
	"wechat-mall-backend/dbops"
	"wechat-mall-backend/model"
)

type IBannerService interface {
	GetBannerList(page, size int) (*[]model.WechatMallBannerDO, int)
	GetBannerById(id int) *model.WechatMallBannerDO
	AddBanner(banner *model.WechatMallBannerDO)
	UpdateBannerById(banner *model.WechatMallBannerDO)
}

type bannerService struct {
}

func NewBannerService() IBannerService {
	service := &bannerService{}
	return service
}

func (bs *bannerService) GetBannerList(page, size int) (*[]model.WechatMallBannerDO, int) {
	bannerList, err := dbops.QueryBannerList("", page, size)
	if err != nil {
		panic(err)
	}
	total, err := dbops.CountBanner("")
	if err != nil {
		panic(err)
	}
	return bannerList, total
}

func (bs *bannerService) GetBannerById(id int) *model.WechatMallBannerDO {
	banner, err := dbops.QueryBannerById(id)
	if err != nil {
		panic(err)
	}
	return banner
}

func (bs *bannerService) AddBanner(banner *model.WechatMallBannerDO) {
	_, err := dbops.InsertBanner(banner)
	if err != nil {
		panic(err)
	}
}

func (bs *bannerService) UpdateBannerById(banner *model.WechatMallBannerDO) {
	err := dbops.UpdateBannerById(banner)
	if err != nil {
		panic(err)
	}
}
