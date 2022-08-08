package interfaces

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"wechat-mall-backend/app/domain/view"
	"wechat-mall-backend/consts"
	"wechat-mall-backend/errcode"
	"wechat-mall-backend/pkg/utils"
)

// GetCouponList 店铺-优惠券列表
func (m *MallHttpServiceImpl) GetCouponList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])

	couponList, total, err := m.couponService.GetCouponList(r.Context(), page, size, consts.ONLINE)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	voList := make([]*view.PortalCouponVO, 0)
	for _, v := range couponList {
		couponVO := &view.PortalCouponVO{}
		couponVO.Id = v.ID
		couponVO.Title = v.Title
		couponVO.FullMoney = v.FullMoney
		couponVO.Minus = v.Minus
		couponVO.Rate = v.Rate
		couponVO.Type = v.Type
		couponVO.StartTime = utils.FormatTime(v.StartTime)
		couponVO.EndTime = utils.FormatTime(v.EndTime)
		couponVO.Description = v.Description
		voList = append(voList, couponVO)
	}
	data := make(map[string]interface{})
	data["list"] = voList
	data["total"] = total
	Ok(w, data)
}

// TakeCoupon 领取优惠券
func (m *MallHttpServiceImpl) TakeCoupon(w http.ResponseWriter, r *http.Request) {
	req := new(PortalTakeCouponReq)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		Error(w, errcode.BadRequestParam, "参数解析失败")
		return
	}
	couponId := req.CouponId
	userId := r.Context().Value(consts.ContextKey).(int)
	coupon, err := m.couponService.GetCouponById(r.Context(), couponId)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	if coupon.ID == consts.ZERO || coupon.Del == consts.DELETE || coupon.Online == consts.OFFLINE {
		Error(w, errcode.NotFoundCoupon, "优惠券不存在或下架了")
		return
	}
	totalTakeNum, err := m.couponService.CountCouponTakeNum(r.Context(), consts.ALL, couponId)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	if totalTakeNum >= coupon.GrantNum {
		Error(w, errcode.NotAllowOperation, "来晚了，优惠券领光了")
		return
	}
	userTakeNum, err := m.couponService.CountCouponTakeNum(r.Context(), userId, couponId)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	if userTakeNum >= coupon.LimitNum {
		Error(w, errcode.NotAllowOperation, "单人限领")
		return
	}
	if err := m.couponService.RecordCouponLog(r.Context(), userId, couponId); err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	Ok(w, "ok")
}

// GetUserCouponList 用户领取的优惠券
// status: 0-未使用 1-已使用 2-已过期
func (m *MallHttpServiceImpl) GetUserCouponList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	status, _ := strconv.Atoi(vars["status"])
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])
	userId := r.Context().Value(consts.ContextKey).(int)

	voList, total, err := m.couponService.QueryUserCoupon(r.Context(), userId, status, page, size)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	data := make(map[string]interface{})
	data["list"] = voList
	data["total"] = total
	Ok(w, data)
}

// DoDeleteCouponLog 删除领取的优惠券
func (m *MallHttpServiceImpl) DoDeleteCouponLog(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	couponLogId, _ := strconv.Atoi(vars["id"])
	userId := r.Context().Value(consts.ContextKey).(int)

	couponLog, err := m.couponService.QueryCouponLogById(r.Context(), couponLogId)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	if couponLog.ID == consts.ZERO || couponLog.Del == consts.DELETE || couponLog.UserID != userId {
		Error(w, errcode.NotFoundCoupon, "未知的记录")
		return
	}
	if err := m.couponService.DoDeleteCouponLog(r.Context(), couponLog); err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	Ok(w, "ok")
}
