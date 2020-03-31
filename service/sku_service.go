package service

import (
	"wechat-mall-backend/dbops"
	"wechat-mall-backend/model"
)

type ISKUService interface {
	GetSKUList(title string, goodsId, online, page, size int) (*[]model.WechatMallSkuDO, int)
	GetSKUById(id int) *model.WechatMallSkuDO
	GetSKUByCode(code string) *model.WechatMallSkuDO
	AddSKU(sku *model.WechatMallSkuDO)
	UpdateSKUById(sku *model.WechatMallSkuDO)
	CountSellOutSKU() int
	QuerySellOutSKU(page, size int) (*[]model.WechatMallSkuDO, int)
}

type sKUService struct {
}

func NewSKUService() ISKUService {
	service := sKUService{}
	return &service
}

func (s *sKUService) GetSKUList(title string, goodsId, online, page, size int) (*[]model.WechatMallSkuDO, int) {
	skuList, err := dbops.GetSKUList(title, goodsId, online, page, size)
	if err != nil {
		panic(err)
	}
	total, err := dbops.CountSKU(title, goodsId, online)
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

// 统计-售罄的SKU数量
func (s *sKUService) CountSellOutSKU() int {
	total, err := dbops.CountSellOutSKUList()
	if err != nil {
		panic(err)
	}
	return total
}

// 查询-售罄的商品
func (s *sKUService) QuerySellOutSKU(page, size int) (*[]model.WechatMallSkuDO, int) {
	skuList, err := dbops.QuerySellOutSKUList(page, size)
	if err != nil {
		panic(err)
	}
	total, err := dbops.CountSellOutSKUList()
	if err != nil {
		panic(err)
	}
	return skuList, total
}
