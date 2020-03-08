package handler

import (
	"encoding/json"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"wechat-mall-backend/defs"
	"wechat-mall-backend/errs"
	"wechat-mall-backend/model"
)

func (h *CMSHandler) GetCouponList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	activityId, _ := strconv.Atoi(vars["activityId"])

	couponVOList := []defs.CouponVO{}
	couponList := h.service.CouponService.GetCouponList(activityId)
	for _, v := range *couponList {
		couponVO := defs.CouponVO{}
		couponVO.Id = v.Id
		couponVO.ActivityId = v.ActivityId
		couponVO.Title = v.Title
		couponVO.FullMoney = v.FullMoney
		couponVO.Minus = v.Minus
		couponVO.Rate = v.Rate
		couponVO.Type = v.Type
		couponVO.StartTime = v.StartTime
		couponVO.EndTime = v.EndTime
		couponVO.Description = v.Description
		couponVOList = append(couponVOList, couponVO)
	}
	sendNormalResponse(w, couponVOList)
}

func (h *CMSHandler) GetCoupon(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	coupon := h.service.CouponService.GetCouponById(id)
	if coupon.Id == 0 {
		panic(errs.ErrorCoupon)
	}
	couponVO := defs.CouponVO{}
	couponVO.Id = coupon.Id
	couponVO.ActivityId = coupon.ActivityId
	couponVO.Title = coupon.Title
	couponVO.FullMoney = coupon.FullMoney
	couponVO.Minus = coupon.Minus
	couponVO.Rate = coupon.Rate
	couponVO.Type = coupon.Type
	couponVO.StartTime = coupon.StartTime
	couponVO.EndTime = coupon.EndTime
	couponVO.Description = coupon.Description

	sendNormalResponse(w, couponVO)
}

func (h *CMSHandler) DoEditCoupon(w http.ResponseWriter, r *http.Request) {
	req := defs.CouponReq{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		panic(errs.ErrorParameterValidate)
	}
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		panic(errs.NewParameterError(err.Error()))
	}
	activity := h.service.ActivityService.GetActivityById(req.ActivityId)
	if activity.Id == 0 {
		panic(errs.ErrorActivity)
	}
	if req.Id == 0 {
		coupon := model.Coupon{}
		coupon.ActivityId = req.ActivityId
		coupon.Title = req.Title
		coupon.FullMoney = req.FullMoney
		coupon.Minus = req.Minus
		coupon.Rate = req.Rate
		coupon.Type = req.Type
		coupon.StartTime = req.StartTime
		coupon.EndTime = req.EndTime
		coupon.Description = req.Description
		h.service.CouponService.AddCoupon(&coupon)
	} else {
		coupon := h.service.CouponService.GetCouponById(req.Id)
		if coupon.Id == 0 {
			panic(errs.ErrorCoupon)
		}
		coupon.ActivityId = req.ActivityId
		coupon.Title = req.Title
		coupon.FullMoney = req.FullMoney
		coupon.Minus = req.Minus
		coupon.Rate = req.Rate
		coupon.Type = req.Type
		coupon.StartTime = req.StartTime
		coupon.EndTime = req.EndTime
		coupon.Description = req.Description
		h.service.CouponService.UpdateCouponById(coupon)
	}
	sendNormalResponse(w, "ok")
}

func (h *CMSHandler) DoDeleteCoupon(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	coupon := h.service.CouponService.GetCouponById(id)
	if coupon.Id == 0 {
		panic(errs.ErrorCoupon)
	}
	coupon.Del = 1
	h.service.CouponService.UpdateCouponById(coupon)
	sendNormalResponse(w, "ok")
}
