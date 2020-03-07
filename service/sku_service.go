package service

import (
	"wechat-mall-backend/dbops"
	"wechat-mall-backend/model"
)

type ISKUService interface {
	GetSKUList(page, size int) (*[]model.SKU, int)
	GetSKUById(id int) *model.SKU
	GetSKUByCode(code string) *model.SKU
	AddSKU(sku *model.SKU)
	UpdateSKUById(sku *model.SKU)
}

type sKUService struct {
}

func NewSKUService() ISKUService {
	service := sKUService{}
	return &service
}

func (s *sKUService) GetSKUList(page, size int) (*[]model.SKU, int) {
	skuList, err := dbops.GetSKUList(page, size)
	if err != nil {
		panic(err)
	}
	total, err := dbops.CountSKU()
	if err != nil {
		panic(err)
	}
	return skuList, total
}

func (s *sKUService) GetSKUById(id int) *model.SKU {
	sku, err := dbops.GetSKUById(id)
	if err != nil {
		panic(err)
	}
	return sku
}

func (s *sKUService) GetSKUByCode(code string) *model.SKU {
	sku, err := dbops.GetSKUByCode(code)
	if err != nil {
		panic(err)
	}
	return sku
}

func (s *sKUService) AddSKU(sku *model.SKU) {
	err := dbops.AddSKU(sku)
	if err != nil {
		panic(err)
	}
}

func (s *sKUService) UpdateSKUById(sku *model.SKU) {
	err := dbops.UpdateSKUById(sku)
	if err != nil {
		panic(err)
	}
}
