package service

import (
	"strconv"
	"strings"
	"wechat-mall-backend/dbops"
	"wechat-mall-backend/model"
)

type IGoodsService interface {
	GetGoodsList(page, size int) (*[]model.WechatMallGoodsDO, int)
	GetGoodsById(id int) *model.WechatMallGoodsDO
	UpdateGoodsById(goods *model.WechatMallGoodsDO)
	AddGoods(goods *model.WechatMallGoodsDO) int
	GetGoodsSpecList(goodsId int) *[]int
	AddGoodsSpec(goodsId int, specList string)
}

type goodsService struct {
}

func NewGoodsService() IGoodsService {
	service := &goodsService{}
	return service
}

func (s *goodsService) GetGoodsList(page, size int) (*[]model.WechatMallGoodsDO, int) {
	goodsList, err := dbops.QueryGoodsList(page, size)
	if err != nil {
		panic(err)
	}
	total, err := dbops.CountGoods()
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
