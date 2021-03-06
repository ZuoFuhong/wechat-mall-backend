package service

import (
	"encoding/json"
	"strconv"
	"wechat-mall-backend/dbops"
	"wechat-mall-backend/defs"
	"wechat-mall-backend/errs"
	"wechat-mall-backend/model"
)

type IGoodsService interface {
	GetGoodsList(keyword string, categoryId, online, page, size int) (*[]model.WechatMallGoodsDO, int)
	GetGoodsById(id int) *model.WechatMallGoodsDO
	UpdateGoodsById(goods *model.WechatMallGoodsDO)
	AddGoods(goods *model.WechatMallGoodsDO) int
	GetGoodsSpecList(goodsId int) *[]defs.CMSGoodsSpecVO
	AddGoodsSpec(goodsId int, specList []int)
	QueryPortalGoodsList(keyword string, sort, categoryId, page, size int) (*[]defs.PortalGoodsListVO, int)
	QueryPortalGoodsDetail(goodsId int) *defs.PortalGoodsInfo
	CountCategoryGoods(categoryId int) int
	CountGoodsSpecBySpecId(specId int) int
}

type goodsService struct {
}

func NewGoodsService() IGoodsService {
	service := &goodsService{}
	return service
}

func (s *goodsService) GetGoodsList(keyword string, categoryId, online, page, size int) (*[]model.WechatMallGoodsDO, int) {
	goodsList, err := dbops.QueryGoodsList(keyword, "", categoryId, online, page, size)
	if err != nil {
		panic(err)
	}
	total, err := dbops.CountGoods(keyword, categoryId, online)
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

func (s *goodsService) GetGoodsSpecList(goodsId int) *[]defs.CMSGoodsSpecVO {
	specList, err := dbops.GetGoodsSpecList(goodsId)
	if err != nil {
		panic(err)
	}
	specVOList := []defs.CMSGoodsSpecVO{}
	for _, v := range *specList {
		specificationDO, err := dbops.QuerySpecificationById(v.SpecId)
		if err != nil {
			panic(err)
		}
		if specificationDO.Id == defs.ZERO || specificationDO.Del == defs.DELETE {
			panic(errs.ErrorSpecificationAttr)
		}
		attrList, err := dbops.QuerySpecificationAttrList(v.SpecId)
		if err != nil {
			panic(err)
		}
		attrVOList := []defs.CMSSpecificationAttrVO{}
		for _, item := range *attrList {
			attrVO := defs.CMSSpecificationAttrVO{}
			attrVO.Id = item.Id
			attrVO.SpecId = item.SpecId
			attrVO.Value = item.Value
			attrVO.Extend = item.Extend
			attrVOList = append(attrVOList, attrVO)
		}
		specVO := defs.CMSGoodsSpecVO{}
		specVO.SpecId = v.SpecId
		specVO.Name = specificationDO.Name
		specVO.AttrList = attrVOList
		specVOList = append(specVOList, specVO)
	}
	return &specVOList
}

func (s *goodsService) AddGoodsSpec(goodsId int, specList []int) {
	err := dbops.DeleteGoodsSpec(goodsId)
	if err != nil {
		panic(err)
	}
	for _, v := range specList {
		spec := model.WechatMallGoodsSpecDO{}
		spec.GoodsId = goodsId
		spec.SpecId = v
		err := dbops.InsertGoodsSpec(&spec)
		if err != nil {
			panic(err)
		}
	}
}

func (s *goodsService) QueryPortalGoodsList(keyword string, sort, categoryId, page, size int) (*[]defs.PortalGoodsListVO, int) {
	// 排序：0-综合 1-新品 2-销量 3-价格
	var order string
	switch sort {
	case 1:
		order = "create_time"
	case 2:
		order = "sale_num"
	case 3:
		order = "price"
	default:
		order = ""
	}
	goodsList, err := dbops.QueryGoodsList(keyword, order, categoryId, defs.ONLINE, page, size)
	if err != nil {
		panic(err)
	}
	total, err := dbops.CountGoods(keyword, categoryId, defs.ONLINE)
	if err != nil {
		panic(err)
	}
	goodsVOList := []defs.PortalGoodsListVO{}
	for _, v := range *goodsList {
		humanNum, err := dbops.CountBuyGoodsUserNum(v.Id)
		if err != nil {
			panic(err)
		}
		goodsVO := defs.PortalGoodsListVO{}
		goodsVO.Id = v.Id
		goodsVO.Title = v.Title
		goodsVO.Price, _ = strconv.ParseFloat(v.Price, 2)
		goodsVO.Picture = v.Picture
		goodsVO.HumanNum = humanNum
		goodsVOList = append(goodsVOList, goodsVO)
	}
	return &goodsVOList, total
}

func (s *goodsService) QueryPortalGoodsDetail(goodsId int) *defs.PortalGoodsInfo {
	goodsDO, err := dbops.QueryGoodsById(goodsId)
	if err != nil {
		panic(err)
	}
	if goodsDO.Id == defs.ZERO || goodsDO.Del == defs.DELETE || goodsDO.Online == defs.OFFLINE {
		panic(errs.ErrorGoods)
	}
	skuDOList, err := dbops.GetSKUList("", goodsId, defs.ONLINE, 0, 0)
	if err != nil {
		panic(err)
	}
	skuList := extractSkuVOList(skuDOList)
	specList := extraceSpecVOList(goodsId, skuDOList)

	goodsVO := defs.PortalGoodsInfo{}
	goodsVO.Id = goodsDO.Id
	goodsVO.Title = goodsDO.Title
	goodsVO.Price, _ = strconv.ParseFloat(goodsDO.Price, 2)
	goodsVO.Picture = goodsDO.Picture
	goodsVO.BannerPicture = goodsDO.BannerPicture
	goodsVO.DetailPicture = goodsDO.DetailPicture
	goodsVO.Tags = goodsDO.Tags
	goodsVO.Description = goodsDO.Description
	goodsVO.SkuList = skuList
	goodsVO.SpecList = specList
	return &goodsVO
}

func extractSkuVOList(skuDOList *[]model.WechatMallSkuDO) []defs.PortalSkuVO {
	skuList := []defs.PortalSkuVO{}
	for _, v := range *skuDOList {
		skuVO := defs.PortalSkuVO{}
		skuVO.Id = v.Id
		skuVO.Picture = v.Picture
		skuVO.Title = v.Title
		skuVO.Price, _ = strconv.ParseFloat(v.Price, 2)
		skuVO.Code = v.Code
		skuVO.Stock = v.Stock
		skuVO.Specs = v.Specs
		skuList = append(skuList, skuVO)
	}
	return skuList
}

func extraceSpecVOList(goodsId int, skuDOList *[]model.WechatMallSkuDO) []defs.PortalSpecVO {
	specVOMap, specAttrVOMap := extraceSpecAttrVOList(skuDOList)
	specList, err := dbops.GetGoodsSpecList(goodsId)
	if err != nil {
		panic(err)
	}
	specVOList := []defs.PortalSpecVO{}
	for _, v := range *specList {
		specId := v.SpecId
		if specVOMap[specId] == "" {
			continue
		}
		specVO := defs.PortalSpecVO{}
		specVO.SpecId = specId
		specVO.Name = specVOMap[specId]
		specVO.AttrList = specAttrVOMap[specId]
		specVOList = append(specVOList, specVO)
	}
	return specVOList
}

func extraceSpecAttrVOList(skuDOList *[]model.WechatMallSkuDO) (map[int]string, map[int][]defs.PortalSpecAttrVO) {
	specVOMap := map[int]string{}
	specAttrVOMap := map[int][]defs.PortalSpecAttrVO{}
	for _, v := range *skuDOList {
		// [{"key": "颜色", "value": "青芒色", "keyId": 1, "valueId": 42}, {"key": "尺寸", "value": "7英寸", "keyId": 2, "valueId": 5}]
		specs := []defs.SkuSpecs{}
		err := json.Unmarshal([]byte(v.Specs), &specs)
		if err != nil {
			panic(err)
		}
		for _, item := range specs {
			specName := specVOMap[item.KeyId]
			if specName == "" {
				specVOMap[item.KeyId] = item.Key
			}
			attrVOList := specAttrVOMap[item.KeyId]
			if attrVOList == nil {
				attrVOList = []defs.PortalSpecAttrVO{}
				specAttrVOMap[item.KeyId] = attrVOList
			}
			flag := false
			for _, attrVO := range attrVOList {
				if attrVO.AttrId == item.ValueId {
					flag = true
					break
				}
			}
			if flag {
				continue
			}
			attrVO := defs.PortalSpecAttrVO{}
			attrVO.AttrId = item.ValueId
			attrVO.Value = item.Value
			attrVOList = append(attrVOList, attrVO)
			specAttrVOMap[item.KeyId] = attrVOList
		}
	}
	return specVOMap, specAttrVOMap
}

func (s *goodsService) CountCategoryGoods(categoryId int) int {
	total, err := dbops.CountCategoryGoods(categoryId)
	if err != nil {
		panic(err)
	}
	return total
}

// 统计-商品规格-关联的商品
func (s *goodsService) CountGoodsSpecBySpecId(specId int) int {
	goodsNum, err := dbops.CountGoodsSpecBySpecId(specId)
	if err != nil {
		panic(err)
	}
	return goodsNum
}
