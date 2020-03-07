package service

import (
	"wechat-mall-backend/dbops"
	"wechat-mall-backend/model"
)

type ICouponService interface {
	GetCouponList(activityId int) *[]model.Coupon
	GetCouponById(id int) *model.Coupon
	AddCoupon(coupon *model.Coupon)
	UpdateCouponById(coupon *model.Coupon)
}

type couponService struct {
}

func NewCouponService() ICouponService {
	service := couponService{}
	return &service
}

func (cs *couponService) GetCouponList(activityId int) *[]model.Coupon {
	couponList, err := dbops.QueryCouponList(activityId)
	if err != nil {
		panic(err)
	}
	return couponList
}

func (cs *couponService) GetCouponById(id int) *model.Coupon {
	coupon, err := dbops.QueryCouponById(id)
	if err != nil {
		panic(err)
	}
	return coupon
}

func (cs *couponService) AddCoupon(coupon *model.Coupon) {
	err := dbops.InsertCoupon(coupon)
	if err != nil {
		panic(err)
	}
}

func (cs *couponService) UpdateCouponById(coupon *model.Coupon) {
	err := dbops.UpdateCouponById(coupon)
	if err != nil {
		panic(err)
	}
}
