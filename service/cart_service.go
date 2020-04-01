package service

import (
	"math"
	"strconv"
	"wechat-mall-backend/dbops"
	"wechat-mall-backend/defs"
	"wechat-mall-backend/errs"
	"wechat-mall-backend/model"
)

type ICartService interface {
	DoEditCart(userId, goodsId, skuId, num int)
	GetCartGoods(userId, page, size int) (*[]defs.PortalCartGoodsVO, int)
	GetCartDOById(id int) *model.WechatMallUserCartDO
	DeleteCartDOById(userId, id int)
}

type cartService struct {
}

func NewCartService() ICartService {
	service := cartService{}
	return &service
}

func (s *cartService) DoEditCart(userId, goodsId, skuId, num int) {
	if math.Abs(float64(num)) > defs.CartMax {
		panic(errs.ErrorParameterValidate)
	}
	goodsDO, err := dbops.QueryGoodsById(goodsId)
	if err != nil {
		panic(err)
	}
	if goodsDO.Id == defs.ZERO || goodsDO.Del == defs.DELETE || goodsDO.Online == defs.OFFLINE {
		panic(errs.ErrorGoods)
	}
	skuDO, err := dbops.GetSKUById(skuId)
	if err != nil {
		panic(err)
	}
	if skuDO.Id == defs.ZERO || skuDO.Del == defs.DELETE || skuDO.Online == defs.OFFLINE {
		panic(errs.ErrorSKU)
	}
	if skuDO.Stock <= 0 {
		panic(errs.NewErrorGoodsCart("库存不足！"))
	}
	cartDO, err := dbops.QueryCartByParams(userId, goodsId, skuId)
	if err != nil {
		panic(err)
	}
	if num > 0 {
		if cartDO.Id == defs.ZERO {
			userCartDO := model.WechatMallUserCartDO{}
			userCartDO.UserId = userId
			userCartDO.GoodsId = goodsId
			userCartDO.SkuId = skuId
			userCartDO.Num = num
			err = dbops.AddUserCart(&userCartDO)
		} else {
			if skuDO.Stock < cartDO.Num+num {
				panic(errs.NewErrorGoodsCart("库存不足！"))
			}
			if cartDO.Num+num > defs.CartMax {
				cartDO.Num = defs.CartMax
			} else {
				cartDO.Num += num
			}
			err = dbops.UpdateCartById(cartDO)
		}
	} else {
		if cartDO.Id == defs.ZERO {
			panic(errs.ErrorGoodsCart)
		}
		if cartDO.Num+num >= 1 {
			cartDO.Num += num
			err = dbops.UpdateCartById(cartDO)
		}
	}
	if err != nil {
		panic(err)
	}
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
		skuDO, err := dbops.GetSKUById(v.SkuId)
		if err != nil {
			panic(err)
		}
		status := 0
		if goodsDO.Id == defs.ZERO || goodsDO.Del == defs.DELETE || goodsDO.Online == defs.OFFLINE ||
			skuDO.Id == defs.ZERO || skuDO.Del == defs.DELETE || skuDO.Online == defs.OFFLINE {
			status = 2
		} else {
			if skuDO.Stock < v.Num {
				status = 1
			}
		}
		cartGoodsVO := defs.PortalCartGoodsVO{}
		cartGoodsVO.Id = v.Id
		cartGoodsVO.GoodsId = v.GoodsId
		cartGoodsVO.SkuId = v.SkuId
		cartGoodsVO.Title = goodsDO.Title
		cartGoodsVO.Price, _ = strconv.ParseFloat(skuDO.Price, 2)
		cartGoodsVO.Picture = skuDO.Picture
		cartGoodsVO.Specs = skuDO.Specs
		cartGoodsVO.Num = v.Num
		cartGoodsVO.Status = status
		cartGoodsVOList = append(cartGoodsVOList, cartGoodsVO)
	}
	return &cartGoodsVOList, total
}

func (s *cartService) GetCartDOById(id int) *model.WechatMallUserCartDO {
	cartDO, err := dbops.SelectCartById(id)
	if err != nil {
		panic(err)
	}
	return cartDO
}

func (s *cartService) DeleteCartDOById(userId, id int) {
	cartDO, err := dbops.SelectCartById(id)
	if err != nil {
		panic(err)
	}
	if cartDO.Id == defs.ZERO || cartDO.Del == defs.DELETE || cartDO.UserId != userId {
		panic(errs.ErrorGoodsCart)
	}
	cartDO.Del = defs.DELETE
	err = dbops.UpdateCartById(cartDO)
	if err != nil {
		panic(err)
	}
}
