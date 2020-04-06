package service

import (
	"time"
	"wechat-mall-backend/dbops"
	"wechat-mall-backend/defs"
	"wechat-mall-backend/errs"
	"wechat-mall-backend/model"
	"wechat-mall-backend/utils"
)

type ICouponService interface {
	GetCouponList(page, size, online int) (*[]model.WechatMallCouponDO, int)
	GetCouponById(id int) *model.WechatMallCouponDO
	AddCoupon(coupon *model.WechatMallCouponDO)
	UpdateCouponById(coupon *model.WechatMallCouponDO)
	QueryCouponLogById(couponLogId int) *model.WechatMallCouponLogDO
	RecordCouponLog(userId, couponId int)
	QueryUserCoupon(userId, status, page, size int) (*[]defs.PortalUserCouponVO, int)
	CountCouponTakeNum(userId, couponId int) int
	DoDeleteCouponLog(couponLog *model.WechatMallCouponLogDO)
	GetAllSubCategory() *[]defs.PortalCategoryVO
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

func (cs *couponService) QueryCouponLogById(couponLogId int) *model.WechatMallCouponLogDO {
	couponLogDO, err := dbops.QueryCouponLogById(couponLogId)
	if err != nil {
		panic(err)
	}
	return couponLogDO
}

func (cs *couponService) RecordCouponLog(userId, couponId int) {
	coupon, err := dbops.QueryCouponById(couponId)
	if err != nil {
		panic(err)
	}
	if coupon.Id == defs.ZERO || coupon.Del == defs.DELETE || coupon.Online == defs.OFFLINE {
		panic(errs.ErrorCoupon)
	}
	couponLog := model.WechatMallCouponLogDO{}
	couponLog.CouponId = couponId
	couponLog.UserId = userId
	couponLog.UseTime = time.Now().Format("2006-01-02 15:04:05")
	couponLog.ExpireTime = coupon.EndTime
	couponLog.Status = 0
	couponLog.Code = utils.RandomNumberStr(12)
	err = dbops.AddCouponLog(&couponLog)
	if err != nil {
		panic(err)
	}
}

func (cs *couponService) QueryUserCoupon(userId, status, page, size int) (*[]defs.PortalUserCouponVO, int) {
	// 刷新券的过期状态（优于调度任务）
	err := dbops.UpdateCouponLogOverdueStatus(userId)
	if err != nil {
		panic(err)
	}
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
		if couponDO.Id == defs.ZERO {
			panic(errs.ErrorCoupon)
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
	total, err := dbops.CountCouponTakeNum(userId, defs.ALL, status, 0)
	if err != nil {
		panic(err)
	}
	return &voList, total
}

// 查询-优惠券领取的数量
func (cs *couponService) CountCouponTakeNum(userId, couponId int) int {
	total, err := dbops.CountCouponTakeNum(userId, couponId, defs.ALL, defs.ALL)
	if err != nil {
		panic(err)
	}
	return total
}

func (cs *couponService) DoDeleteCouponLog(couponLog *model.WechatMallCouponLogDO) {
	couponLog.Del = 1
	err := dbops.UpdateCouponLogById(couponLog)
	if err != nil {
		panic(err)
	}
}

// 查询-所有的二级分类
func (cs *couponService) GetAllSubCategory() *[]defs.PortalCategoryVO {
	categoryList, err := dbops.QueryAllSubCategory()
	if err != nil {
		panic(err)
	}
	categoryVOList := []defs.PortalCategoryVO{}
	for _, v := range *categoryList {
		categoryVO := defs.PortalCategoryVO{}
		categoryVO.Id = v.Id
		categoryVO.Name = v.Name
		categoryVOList = append(categoryVOList, categoryVO)
	}
	return &categoryVOList
}
