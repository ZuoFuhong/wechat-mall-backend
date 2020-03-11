package service

import (
	"wechat-mall-backend/dbops"
	"wechat-mall-backend/defs"
	"wechat-mall-backend/model"
	"wechat-mall-backend/utils"
)

type ICouponService interface {
	GetCouponList(page, size, online int) (*[]model.WechatMallCouponDO, int)
	GetCouponById(id int) *model.WechatMallCouponDO
	AddCoupon(coupon *model.WechatMallCouponDO)
	UpdateCouponById(coupon *model.WechatMallCouponDO)
	QueryCouponLog(userId, couponId int) *model.WechatMallCouponLogDO
	RecordCouponLog(userId, couponId int)
	QueryUserCoupon(userId, status, page, size int) (*[]defs.PortalUserCouponVO, int)
	DoDeleteCouponLog(couponLogId int)
}

type couponService struct {
}

func NewCouponService() ICouponService {
	service := couponService{}
	return &service
}

func (cs *couponService) GetCouponList(page, size, online int) (*[]model.WechatMallCouponDO, int) {
	couponList, err := dbops.QueryCouponList(page, size, online)
	if err != nil {
		panic(err)
	}
	total, err := dbops.CountCoupon(online)
	if err != nil {
		panic(err)
	}
	return couponList, total
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

func (cs *couponService) QueryCouponLog(userId, couponId int) *model.WechatMallCouponLogDO {
	couponLog, err := dbops.QueryCouponLog(userId, couponId)
	if err != nil {
		panic(err)
	}
	return couponLog
}

func (cs *couponService) RecordCouponLog(userId, couponId int) {
	coupon, err := dbops.QueryCouponById(couponId)
	if err != nil {
		panic(err)
	}
	couponLog := model.WechatMallCouponLogDO{}
	couponLog.CouponId = couponId
	couponLog.UserId = userId
	couponLog.ExpireTime = coupon.EndTime
	couponLog.Status = 0
	couponLog.Code = utils.RandomNumberStr(12)
	err = dbops.AddCouponLog(&couponLog)
	if err != nil {
		panic(err)
	}
}

func (cs *couponService) QueryUserCoupon(userId, status, page, size int) (*[]defs.PortalUserCouponVO, int) {
	couponLogList, err := dbops.QueryCouponLogList(userId, status, page, size)
	if err != nil {
		panic(err)
	}
	voList := []defs.PortalUserCouponVO{}
	for _, v := range *couponLogList {
		couponDO, err := dbops.QueryCouponById(v.CouponId)
		if err != nil {
			panic(err)
		}
		couponVO := defs.PortalUserCouponVO{}
		couponVO.CLogId = v.Id
		couponVO.CouponId = v.CouponId
		couponVO.Title = couponDO.Title
		couponVO.FullMoney = couponDO.FullMoney
		couponVO.Minus = couponDO.Minus
		couponVO.Rate = couponDO.Rate
		couponVO.Type = couponDO.Type
		couponVO.StartTime = couponDO.StartTime
		couponVO.EndTime = v.ExpireTime
		couponVO.Description = couponDO.Description
		voList = append(voList, couponVO)
	}
	total, err := dbops.CountUserCouponLog(userId, status)
	if err != nil {
		panic(err)
	}
	return &voList, total
}

func (cs *couponService) DoDeleteCouponLog(couponLogId int) {
	couponLog := model.WechatMallCouponLogDO{}
	couponLog.Id = couponLogId
	couponLog.Del = 1
	err := dbops.UpdateCouponLogById(&couponLog)
	if err != nil {
		panic(err)
	}
}
