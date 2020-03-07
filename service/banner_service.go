package service

import (
	"wechat-mall-backend/dbops"
	"wechat-mall-backend/model"
)

type IBannerService interface {
	GetBannerList(page, size int) (*[]model.Banner, int)
	GetBannerById(id int) *Banner
	AddBanner(banner *Banner)
	UpdateBannerById(banner *Banner)
	GetBannerItemList(bannerId int) *[]model.BannerItem
	GetBannerItemById(id int) *BannerItem
	AddBannerItem(banner *BannerItem)
	UpdateBannerItemById(banner *BannerItem)
}

type bannerService struct {
}

func NewBannerService() IBannerService {
	service := &bannerService{}
	return service
}

func (bs *bannerService) GetBannerList(page, size int) (*[]model.Banner, int) {
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

func (bs *bannerService) GetBannerById(id int) *Banner {
	banner, err := dbops.QueryBannerById(id)
	if err != nil {
		panic(err)
	}
	return (*Banner)(banner)
}

func (bs *bannerService) AddBanner(banner *Banner) {
	_, err := dbops.InsertBanner((*model.Banner)(banner))
	if err != nil {
		panic(err)
	}
}

func (bs *bannerService) UpdateBannerById(banner *Banner) {
	err := dbops.UpdateBannerById((*model.Banner)(banner))
	if err != nil {
		panic(err)
	}
}

func (bs *bannerService) GetBannerItemList(bannerId int) *[]model.BannerItem {
	itemList, err := dbops.QueryBannerItemList(bannerId)
	if err != nil {
		panic(err)
	}
	return itemList
}

func (bs *bannerService) GetBannerItemById(id int) *BannerItem {
	item, err := dbops.QueryBannerItemById(id)
	if err != nil {
		panic(err)
	}
	return (*BannerItem)(item)
}

func (bs *bannerService) AddBannerItem(banner *BannerItem) {
	_, err := dbops.InsertBannerItem((*model.BannerItem)(banner))
	if err != nil {
		panic(err)
	}
}

func (bs *bannerService) UpdateBannerItemById(banner *BannerItem) {
	err := dbops.UpdateBannerItemById((*model.BannerItem)(banner))
	if err != nil {
		panic(err)
	}
}
