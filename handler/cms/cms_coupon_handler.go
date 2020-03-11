package cms

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

func (h *Handler) GetCouponList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])

	couponVOList := []defs.CMSCouponVO{}
	couponList, total := h.service.CouponService.GetCouponList(page, size, 0)
	for _, v := range *couponList {
		couponVO := defs.CMSCouponVO{}
		couponVO.Id = v.Id
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
	resp := make(map[string]interface{})
	resp["list"] = couponVOList
	resp["total"] = total
	defs.SendNormalResponse(w, resp)
}

func (h *Handler) GetCoupon(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	coupon := h.service.CouponService.GetCouponById(id)
	if coupon.Id == 0 {
		panic(errs.ErrorCoupon)
	}
	couponVO := defs.CMSCouponVO{}
	couponVO.Id = coupon.Id
	couponVO.Title = coupon.Title
	couponVO.FullMoney = coupon.FullMoney
	couponVO.Minus = coupon.Minus
	couponVO.Rate = coupon.Rate
	couponVO.Type = coupon.Type
	couponVO.StartTime = coupon.StartTime
	couponVO.EndTime = coupon.EndTime
	couponVO.Description = coupon.Description
	defs.SendNormalResponse(w, couponVO)
}

func (h *Handler) DoEditCoupon(w http.ResponseWriter, r *http.Request) {
	req := defs.CMSCouponReq{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		panic(errs.ErrorParameterValidate)
	}
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		panic(errs.NewParameterError(err.Error()))
	}
	if req.Id == 0 {
		coupon := model.WechatMallCouponDO{}
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
	defs.SendNormalResponse(w, "ok")
}

func (h *Handler) DoDeleteCoupon(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	coupon := h.service.CouponService.GetCouponById(id)
	if coupon.Id == 0 {
		panic(errs.ErrorCoupon)
	}
	coupon.Del = 1
	h.service.CouponService.UpdateCouponById(coupon)
	defs.SendNormalResponse(w, "ok")
}
