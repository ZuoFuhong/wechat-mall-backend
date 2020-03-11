package service

import (
	"wechat-mall-backend/dbops"
	"wechat-mall-backend/model"
)

type ISKUService interface {
	GetSKUList(page, size int) (*[]model.WechatMallSkuDO, int)
	GetSKUById(id int) *model.WechatMallSkuDO
	GetSKUByCode(code string) *model.WechatMallSkuDO
	AddSKU(sku *model.WechatMallSkuDO)
	UpdateSKUById(sku *model.WechatMallSkuDO)
}

type sKUService struct {
}

func NewSKUService() ISKUService {
	service := sKUService{}
	return &service
}

func (s *sKUService) GetSKUList(page, size int) (*[]model.WechatMallSkuDO, int) {
	skuList, err := dbops.GetSKUList(0, page, size)
	if err != nil {
		panic(err)
	}
	total, err := dbops.CountSKU()
	if err != nil {
		panic(err)
	}
	return skuList, total
}

func (s *sKUService) GetSKUById(id int) *model.WechatMallSkuDO {
	sku, err := dbops.GetSKUById(id)
	if err != nil {
		panic(err)
	}
	return sku
}

func (s *sKUService) GetSKUByCode(code string) *model.WechatMallSkuDO {
	sku, err := dbops.GetSKUByCode(code)
	if err != nil {
		panic(err)
	}
	return sku
}

func (s *sKUService) AddSKU(sku *model.WechatMallSkuDO) {
	err := dbops.AddSKU(sku)
	if err != nil {
		panic(err)
	}
}

func (s *sKUService) UpdateSKUById(sku *model.WechatMallSkuDO) {
	err := dbops.UpdateSKUById(sku)
	if err != nil {
		panic(err)
	}
}
