package service

import (
	"wechat-mall-backend/dbops"
	"wechat-mall-backend/model"
)

type ICouponService interface {
	GetCouponList(page, size int) *[]model.WechatMallCouponDO
	GetCouponById(id int) *model.WechatMallCouponDO
	AddCoupon(coupon *model.WechatMallCouponDO)
	UpdateCouponById(coupon *model.WechatMallCouponDO)
}

type couponService struct {
}

func NewCouponService() ICouponService {
	service := couponService{}
	return &service
}

func (cs *couponService) GetCouponList(page, size int) *[]model.WechatMallCouponDO {
	couponList, err := dbops.QueryCouponList(page, size)
	if err != nil {
		panic(err)
	}
	return couponList
}

func (cs *couponService) GetCouponById(id int) *model.WechatMallCouponDO {
	coupon, err := dbops.QueryCouponById(id)
	if err != nil {
		panic(err)
	}
	return coupon
}

func (cs *couponService) AddCoupon(coupon *model.WechatMallCouponDO) {
	err := dbops.InsertCoupon(coupon)
	if err != nil {
		panic(err)
	}
}

func (cs *couponService) UpdateCouponById(coupon *model.WechatMallCouponDO) {
	err := dbops.UpdateCouponById(coupon)
	if err != nil {
		panic(err)
	}
}
