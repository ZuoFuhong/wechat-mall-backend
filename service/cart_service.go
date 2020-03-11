package service

import (
	"encoding/json"
	"wechat-mall-backend/dbops"
	"wechat-mall-backend/defs"
	"wechat-mall-backend/model"
)

type ICartService interface {
	DoEditCart(userId, goodsId, skuId, num int)
	GetCartGoods(userId, page, size int) (*[]defs.PortalCartGoodsVO, int)
}

type cartService struct {
}

func NewCartService() ICartService {
	service := cartService{}
	return &service
}

func (s *cartService) DoEditCart(userId, goodsId, skuId, num int) {
	if num == 0 {
		return
	}
	cartDO, err := dbops.QueryCartByParams(userId, goodsId, skuId)
	if err != nil {
		panic(err)
	}
	if num > 0 {
		if cartDO.Id == 0 {
			userCartDO := model.WechatMallUserCartDO{}
			userCartDO.UserId = userId
			userCartDO.GoodsId = goodsId
			userCartDO.SkuId = skuId
			userCartDO.Num = num
			err = dbops.AddUserCart(&userCartDO)
		} else {
			cartDO.Num += num
			err = dbops.UpdateCartById(cartDO)
		}
	} else {
		if cartDO.Id == 0 {
			return
		}
		if cartDO.Num+num > 0 {
			cartDO.Num += num
			err = dbops.UpdateCartById(cartDO)
		} else {
			cartDO.Del = 1
			err = dbops.UpdateCartById(cartDO)
		}
	}
	panic(err)
}

func (s *cartService) GetCartGoods(userId, page, size int) (*[]defs.PortalCartGoodsVO, int) {
	cartList, err := dbops.QueryCartList(userId, page, size)
	if err != nil {
		panic(err)
	}
	total, err := dbops.CountCartGoods(userId)
	if err != nil {
		panic(err)
	}
	cartGoodsVOList := []defs.PortalCartGoodsVO{}
	for _, v := range *cartList {
		goodsDO, err := dbops.QueryGoodsById(v.GoodsId)
		if err != nil {
			panic(err)
		}
		cartGoodsVO := defs.PortalCartGoodsVO{}
		cartGoodsVO.GoodsId = v.Id
		cartGoodsVO.Title = goodsDO.Title
		cartGoodsVO.Price = goodsDO.Price
		cartGoodsVO.DiscountPrice = goodsDO.DiscountPrice
		cartGoodsVO.Picture = goodsDO.Picture
		cartGoodsVO.Tags = goodsDO.Tags
		cartGoodsVO.SkuId = v.SkuId
		cartGoodsVO.SkuSpecs = extractSkuSpecs(v.SkuId)
		cartGoodsVO.Num = v.Num
		cartGoodsVOList = append(cartGoodsVOList, cartGoodsVO)
	}
	return &cartGoodsVOList, total
}

func extractSkuSpecs(skuId int) []defs.SkuSpecs {
	skuDO, err := dbops.GetSKUById(skuId)
	if err != nil {
		panic(err)
	}
	skuSpecs := []defs.SkuSpecs{}
	err = json.Unmarshal([]byte(skuDO.Specs), &skuSpecs)
	if err != nil {
		panic(err)
	}
	return skuSpecs
}
