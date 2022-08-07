package service

import (
	"context"
	"github.com/pkg/errors"
	"math"
	"strconv"
	"wechat-mall-backend/app/domain/entity"
	"wechat-mall-backend/app/domain/repository"
	"wechat-mall-backend/consts"
)

type ICartService interface {
	DoEditCart(ctx context.Context, userId, goodsId, skuId, num int) error

	GetCartGoods(ctx context.Context, userId, page, size int) ([]*entity.PortalCartGoods, int, error)

	GetCartDOById(ctx context.Context, id int) (*entity.WechatMallUserCartDO, error)

	DeleteCartDOById(ctx context.Context, userId, id int) error

	CountCartGoodsNum(ctx context.Context, userId int) (int, error)
}

type cartService struct {
	repos     repository.IUserCartRepos
	goodsReps repository.IGoodsRepos
	skuRepos  repository.IGoodsSkuRepos
}

func NewCartService(repos repository.IUserCartRepos, goodsReps repository.IGoodsRepos, skuRepos repository.IGoodsSkuRepos) ICartService {
	service := &cartService{
		repos:     repos,
		goodsReps: goodsReps,
		skuRepos:  skuRepos,
	}
	return service
}

func (s *cartService) DoEditCart(ctx context.Context, userId, goodsId, skuId, num int) error {
	if math.Abs(float64(num)) > consts.CartMax {
		return errors.New("cart num than max number")
	}
	goodsDO, err := s.goodsReps.QueryGoodsById(ctx, goodsId)
	if err != nil {
		return err
	}
	if goodsDO.ID == consts.ZERO || goodsDO.Del == consts.DELETE || goodsDO.Online == consts.OFFLINE {
		return errors.New("not found goods record")
	}
	skuDO, err := s.skuRepos.GetSKUById(ctx, skuId)
	if err != nil {
		return err
	}
	if skuDO.ID == consts.ZERO || skuDO.Del == consts.DELETE || skuDO.Online == consts.OFFLINE {
		return errors.New("not found sku record")
	}
	if skuDO.Stock <= 0 {
		return errors.New("Inventory shortage")
	}
	cartDO, err := s.repos.QueryCartByParams(ctx, userId, goodsId, skuId)
	if err != nil {
		return err
	}
	if num > 0 {
		if cartDO.ID == consts.ZERO {
			userCartDO := &entity.WechatMallUserCartDO{
				UserID:  userId,
				GoodsID: goodsId,
				SkuID:   skuId,
				Num:     num,
			}
			err = s.repos.AddUserCart(ctx, userCartDO)
		} else {
			if skuDO.Stock < cartDO.Num+num {
				return errors.New("Inventory shortage")
			}
			if cartDO.Num+num > consts.CartMax {
				cartDO.Num = consts.CartMax
			} else {
				cartDO.Num += num
			}
			err = s.repos.UpdateCartById(ctx, cartDO)
		}
	} else {
		if cartDO.ID == consts.ZERO {
			return errors.New("not found cart goods")
		}
		if cartDO.Num+num >= 1 {
			cartDO.Num += num
			err = s.repos.UpdateCartById(ctx, cartDO)
		}
	}
	return err
}

func (s *cartService) GetCartGoods(ctx context.Context, userId, page, size int) ([]*entity.PortalCartGoods, int, error) {
	cartList, total, err := s.repos.QueryCartList(ctx, userId, page, size)
	if err != nil {
		return nil, 0, err
	}
	voList := make([]*entity.PortalCartGoods, 0)
	for _, cart := range cartList {
		goodsDO, err := s.goodsReps.QueryGoodsById(ctx, cart.GoodsID)
		if err != nil {
			return nil, 0, err
		}
		skuDO, err := s.skuRepos.GetSKUById(ctx, cart.SkuID)
		if err != nil {
			return nil, 0, err
		}
		status := 0
		if goodsDO.ID == consts.ZERO || goodsDO.Del == consts.DELETE || goodsDO.Online == consts.OFFLINE ||
			skuDO.ID == consts.ZERO || skuDO.Del == consts.DELETE || skuDO.Online == consts.OFFLINE {
			status = 2
		} else {
			if skuDO.Stock < cart.Num {
				status = 1
			}
		}
		price, _ := strconv.ParseFloat(skuDO.Price, 2)
		vo := &entity.PortalCartGoods{
			Id:      cart.ID,
			GoodsId: cart.GoodsID,
			SkuId:   cart.SkuID,
			Title:   goodsDO.Title,
			Price:   price,
			Num:     cart.Num,
			Status:  status,
		}
		voList = append(voList, vo)
	}
	return voList, total, nil
}

func (s *cartService) GetCartDOById(ctx context.Context, id int) (*entity.WechatMallUserCartDO, error) {
	return s.repos.SelectCartById(ctx, id)
}

func (s *cartService) DeleteCartDOById(ctx context.Context, userId, id int) error {
	cartDO, err := s.repos.SelectCartById(ctx, id)
	if err != nil {
		return err
	}
	if cartDO.ID == consts.ZERO || cartDO.Del == consts.DELETE || cartDO.UserID != userId {
		return errors.New("not found cart record")
	}
	cartDO.Del = consts.DELETE
	return s.repos.UpdateCartById(ctx, cartDO)
}

func (s *cartService) CountCartGoodsNum(ctx context.Context, userId int) (int, error) {
	return s.repos.CoundCartGoodsNum(ctx, userId)
}
