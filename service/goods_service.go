package service

import (
	"encoding/json"
	"strconv"
	"strings"
	"wechat-mall-backend/dbops"
	"wechat-mall-backend/defs"
	"wechat-mall-backend/errs"
	"wechat-mall-backend/model"
)

type IGoodsService interface {
	GetGoodsList(page, size int) (*[]model.WechatMallGoodsDO, int)
	GetGoodsById(id int) *model.WechatMallGoodsDO
	UpdateGoodsById(goods *model.WechatMallGoodsDO)
	AddGoods(goods *model.WechatMallGoodsDO) int
	GetGoodsSpecList(goodsId int) *[]int
	AddGoodsSpec(goodsId int, specList string)
	QueryGoodsList(goodsName string, order string, categoryId, page, size int) (*[]defs.PortalGoodsListVO, int)
	QueryGoodsDetail(goodsId int) *defs.PortalGoodsInfo
}

type goodsService struct {
}

func NewGoodsService() IGoodsService {
	service := &goodsService{}
	return service
}

func (s *goodsService) GetGoodsList(page, size int) (*[]model.WechatMallGoodsDO, int) {
	goodsList, err := dbops.QueryGoodsList("", "", 0, page, size)
	if err != nil {
		panic(err)
	}
	total, err := dbops.CountGoods("", 0)
	if err != nil {
		panic(err)
	}
	return goodsList, total
}

func (s *goodsService) GetGoodsById(id int) *model.WechatMallGoodsDO {
	goods, err := dbops.QueryGoodsById(id)
	if err != nil {
		panic(err)
	}
	return goods
}

func (s *goodsService) UpdateGoodsById(goods *model.WechatMallGoodsDO) {
	err := dbops.UpdateGoodsById(goods)
	if err != nil {
		panic(err)
	}
}

func (s *goodsService) AddGoods(goods *model.WechatMallGoodsDO) int {
	id, err := dbops.AddGoods(goods)
	if err != nil {
		panic(err)
	}
	return int(id)
}

func (s *goodsService) GetGoodsSpecList(goodsId int) *[]int {
	specList, err := dbops.GetGoodsSpecList(goodsId)
	if err != nil {
		panic(err)
	}
	specIds := []int{}
	for _, v := range *specList {
		specIds = append(specIds, v.SpecId)
	}
	return &specIds
}

func (s *goodsService) AddGoodsSpec(goodsId int, specList string) {
	err := dbops.DeleteGoodsSpec(goodsId)
	if err != nil {
		panic(err)
	}
	specIds := strings.Split(specList, ",")
	for _, v := range specIds {
		spec := model.WechatMallGoodsSpecDO{}
		spec.GoodsId = goodsId
		spec.SpecId, _ = strconv.Atoi(v)
		err := dbops.InsertGoodsSpec(&spec)
		if err != nil {
			panic(err)
		}
	}
}

func (s *goodsService) QueryGoodsList(keyword string, order string, categoryId, page, size int) (*[]defs.PortalGoodsListVO, int) {
	goodsList, err := dbops.QueryGoodsList(keyword, order, categoryId, page, size)
	if err != nil {
		panic(err)
	}
	total, err := dbops.CountGoods(keyword, categoryId)
	if err != nil {
		panic(err)
	}
	goodsVOList := []defs.PortalGoodsListVO{}
	for _, v := range *goodsList {
		saleNum, err := dbops.SumGoodsSaleNum(v.Id, 0)
		if err != nil {
			panic(err)
		}
		goodsVO := defs.PortalGoodsListVO{}
		goodsVO.Id = v.Id
		goodsVO.Title = v.Title
		goodsVO.Price = v.Price
		goodsVO.DiscountPrice = v.DiscountPrice
		goodsVO.Picture = v.Picture
		goodsVO.Tags = v.Tags
		goodsVO.SaleNum = saleNum
		goodsVOList = append(goodsVOList, goodsVO)
	}
	return &goodsVOList, total
}

func (s *goodsService) QueryGoodsDetail(goodsId int) *defs.PortalGoodsInfo {
	goodsDO, err := dbops.QueryGoodsById(goodsId)
	if err != nil {
		panic(err)
	}
	if goodsDO.Id == 0 {
		panic(errs.ErrorGoods)
	}
	skuDOList, err := dbops.GetSKUList(goodsId, 0, 0)
	if err != nil {
		panic(err)
	}
	muptiplePicture := extractMuptiplePicture(skuDOList)
	skuList := extractSkuVOList(skuDOList)
	specList := extraceSpecVOList(goodsId, skuDOList)

	goodsVO := defs.PortalGoodsInfo{}
	goodsVO.Id = goodsDO.Id
	goodsVO.BrandName = goodsDO.BrandName
	goodsVO.Title = goodsDO.Title
	goodsVO.Price = goodsDO.Price
	goodsVO.DiscountPrice = goodsDO.DiscountPrice
	goodsVO.Picture = goodsDO.Picture
	goodsVO.BannerPicture = goodsDO.BannerPicture
	goodsVO.DetailPicture = goodsDO.DetailPicture
	goodsVO.Tags = goodsDO.Tags
	goodsVO.Description = goodsDO.Description
	goodsVO.MultiplePicture = muptiplePicture
	goodsVO.SkuList = skuList
	goodsVO.SpecList = specList
	return &goodsVO
}

func extractMuptiplePicture(skuDOList *[]model.WechatMallSkuDO) []string {
	multiplePicture := []string{}
	for _, v := range *skuDOList {
		multiplePicture = append(multiplePicture, v.Picture)
	}
	return multiplePicture
}

func extractSkuVOList(skuDOList *[]model.WechatMallSkuDO) []defs.PortalSkuVO {
	skuList := []defs.PortalSkuVO{}
	for _, v := range *skuDOList {
		skuVO := defs.PortalSkuVO{}
		skuVO.Id = v.Id
		skuVO.Title = v.Title
		skuVO.Price = v.Price
		skuVO.Code = v.Code
		skuVO.Stock = v.Stock
		skuVO.Specs = v.Specs
		skuList = append(skuList, skuVO)
	}
	return skuList
}

func extraceSpecVOList(goodsId int, skuDOList *[]model.WechatMallSkuDO) []defs.PortalSpecVO {
	specVOMap := extraceSpecAttrVOList(skuDOList)
	specList, err := dbops.GetGoodsSpecList(goodsId)
	if err != nil {
		panic(err)
	}
	specVOList := []defs.PortalSpecVO{}
	for _, v := range *specList {
		specId := v.Id
		specDO, err := dbops.QuerySpecificationById(specId)
		if err != nil {
			panic(err)
		}
		specVO := defs.PortalSpecVO{}
		specVO.SpecId = specId
		specVO.Name = specDO.Name
		specVO.AttrList = specVOMap[specId]
		specVOList = append(specVOList, specVO)
	}
	return specVOList
}

func extraceSpecAttrVOList(skuDOList *[]model.WechatMallSkuDO) map[int][]defs.PortalSpecAttrVO {
	specVOMap := map[int][]defs.PortalSpecAttrVO{}
	for _, v := range *skuDOList {
		// [{"key": "颜色", "value": "青芒色", "key_id": 1, "value_id": 42}, {"key": "尺寸", "value": "7英寸", "key_id": 2, "value_id": 5}]
		specs := []defs.SkuSpecs{}
		err := json.Unmarshal([]byte(v.Specs), &specs)
		if err != nil {
			panic(err)
		}
		for _, item := range specs {
			attrVOList := specVOMap[item.KeyId]
			if attrVOList == nil {
				attrVOList = []defs.PortalSpecAttrVO{}
				specVOMap[item.KeyId] = attrVOList
			}
			for _, attrVO := range attrVOList {
				if attrVO.AttrId == item.ValueId {
					break
				}
				attrVO := defs.PortalSpecAttrVO{}
				attrVO.AttrId = item.ValueId
				attrVO.Value = item.Value
				attrVOList = append(attrVOList, attrVO)
			}
		}
	}
	return specVOMap
}
