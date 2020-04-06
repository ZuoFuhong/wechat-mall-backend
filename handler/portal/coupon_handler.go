package portal

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"wechat-mall-backend/defs"
	"wechat-mall-backend/errs"
)

// 店铺-优惠券列表
func (h *Handler) GetCouponList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])

	couponList, total := h.service.CouponService.GetCouponList(page, size, defs.ONLINE)
	voList := []defs.PortalCouponVO{}
	for _, v := range *couponList {
		couponVO := defs.PortalCouponVO{}
		couponVO.Id = v.Id
		couponVO.Title = v.Title
		couponVO.FullMoney = v.FullMoney
		couponVO.Minus = v.Minus
		couponVO.Rate = v.Rate
		couponVO.Type = v.Type
		couponVO.StartTime = v.StartTime
		couponVO.EndTime = v.EndTime
		couponVO.Description = v.Description
		voList = append(voList, couponVO)
	}
	resp := make(map[string]interface{})
	resp["list"] = voList
	resp["total"] = total
	defs.SendNormalResponse(w, resp)
}

// 领取优惠券
func (h *Handler) TakeCoupon(w http.ResponseWriter, r *http.Request) {
	req := defs.PortalTakeCouponReq{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		panic(errs.ErrorParameterValidate)
	}
	couponId := req.CouponId
	userId := r.Context().Value(defs.ContextKey).(int)

	coupon := h.service.CouponService.GetCouponById(couponId)
	if coupon.Id == defs.ZERO || coupon.Del == defs.DELETE || coupon.Online == defs.OFFLINE {
		panic(errs.NewErrorCoupon("优惠券不存在或下架了"))
	}
	totalTakeNum := h.service.CouponService.CountCouponTakeNum(defs.ALL, couponId)
	if totalTakeNum >= coupon.GrantNum {
		panic(errs.NewErrorCoupon("来晚了，优惠券领光了！"))
	}
	userTakeNum := h.service.CouponService.CountCouponTakeNum(userId, couponId)
	if userTakeNum >= coupon.LimitNum {
		panic(errs.NewErrorCoupon("单人限领！"))
	}
	h.service.CouponService.RecordCouponLog(userId, couponId)
	defs.SendNormalResponse(w, "ok")
}

// 用户领取的优惠券
// status: 0-未使用 1-已使用 2-已过期
func (h *Handler) GetUserCouponList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	status, _ := strconv.Atoi(vars["status"])
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])
	userId := r.Context().Value(defs.ContextKey).(int)

	voList, total := h.service.CouponService.QueryUserCoupon(userId, status, page, size)
	resp := make(map[string]interface{})
	resp["list"] = voList
	resp["total"] = total
	defs.SendNormalResponse(w, resp)
}

// 删除领取的优惠券
func (h *Handler) DoDeleteCouponLog(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	couponLogId, _ := strconv.Atoi(vars["id"])
	userId := r.Context().Value(defs.ContextKey).(int)

	couponLog := h.service.CouponService.QueryCouponLogById(couponLogId)
	if couponLog.Id == defs.ZERO || couponLog.Del == defs.DELETE || couponLog.UserId != userId {
		panic(errs.NewErrorCoupon("未知的记录！"))
	}
	h.service.CouponService.DoDeleteCouponLog(couponLog)
	defs.SendNormalResponse(w, "ok")
}
