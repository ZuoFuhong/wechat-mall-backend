package service

import (
	"encoding/json"
	"wechat-mall-backend/dbops"
	"wechat-mall-backend/defs"
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
	CountAttrRelatedSku(attrId int) int
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
	skuId, err := dbops.AddSKU(sku)
	if err != nil {
		panic(err)
	}
	syncSkuSpecAttrRecord(int(skuId), sku.Specs)
}

func (s *sKUService) UpdateSKUById(sku *model.WechatMallSkuDO) {
	err := dbops.UpdateSKUById(sku)
	if err != nil {
		panic(err)
	}
	if sku.Del == 1 {
		err = dbops.RemoveRelatedBySkuId(sku.Id)
		if err != nil {
			panic(err)
		}
	} else {
		syncSkuSpecAttrRecord(sku.Id, sku.Specs)
	}
}

// 同步-关联SKU属性
func syncSkuSpecAttrRecord(skuId int, specs string) {
	err := dbops.RemoveRelatedBySkuId(skuId)
	if err != nil {
		panic(err)
	}
	skuSpecList := []defs.SkuSpecs{}
	err = json.Unmarshal([]byte(specs), &skuSpecList)
	if err != nil {
		panic(err)
	}
	for _, v := range skuSpecList {
		attrDO := model.WechatMallSkuSpecAttrDO{}
		attrDO.SkuId = model.ID(skuId)
		attrDO.SpecId = v.KeyId
		attrDO.AttrId = v.ValueId
		err := dbops.InsertSkuSpecAttr(&attrDO)
		if err != nil {
			panic(err)
		}
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

func (s *sKUService) CountAttrRelatedSku(attrId int) int {
	total, err := dbops.CountRelatedByAttrId(attrId)
	if err != nil {
		panic(err)
	}
	return total
}
