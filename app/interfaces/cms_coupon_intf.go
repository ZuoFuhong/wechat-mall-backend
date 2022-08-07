package interfaces

import (
	"encoding/json"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
	"wechat-mall-backend/app/domain/entity"
	"wechat-mall-backend/app/domain/view"
	"wechat-mall-backend/consts"
	"wechat-mall-backend/errcode"
	"wechat-mall-backend/pkg/utils"
)

// GetCmsCouponList 查询-优惠券列表
func (m *MallHttpServiceImpl) GetCmsCouponList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])

	couponVOList := make([]*view.CMSCouponVO, 0)
	couponList, total, err := m.couponService.GetCouponList(r.Context(), page, size, consts.ALL)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	for _, v := range couponList {
		couponVO := &view.CMSCouponVO{}
		couponVO.Id = v.ID
		couponVO.Title = v.Title
		couponVO.FullMoney = v.FullMoney
		couponVO.Minus = v.Minus
		couponVO.Rate = v.Rate
		couponVO.Type = v.Type
		couponVO.GrantNum = v.GrantNum
		couponVO.LimitNum = v.LimitNum
		couponVO.StartTime = utils.FormatTime(v.StartTime)
		couponVO.EndTime = utils.FormatTime(v.EndTime)
		couponVO.Description = v.Description
		couponVO.Online = v.Online
		couponVOList = append(couponVOList, couponVO)
	}
	data := make(map[string]interface{})
	data["list"] = couponVOList
	data["total"] = total
	Ok(w, data)
}

// GetCoupon 查询优惠券
func (m *MallHttpServiceImpl) GetCoupon(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	coupon, err := m.couponService.GetCouponById(r.Context(), id)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	if coupon.ID == consts.ZERO || coupon.Del == consts.DELETE {
		Error(w, errcode.NotFoundCoupon, "优惠券不存在")
		return
	}
	couponVO := &view.CMSCouponVO{}
	couponVO.Id = coupon.ID
	couponVO.Title = coupon.Title
	couponVO.FullMoney = coupon.FullMoney
	couponVO.Minus = coupon.Minus
	couponVO.Rate = coupon.Rate
	couponVO.Type = coupon.Type
	couponVO.GrantNum = coupon.GrantNum
	couponVO.LimitNum = coupon.LimitNum
	couponVO.StartTime = utils.FormatTime(coupon.StartTime)
	couponVO.EndTime = utils.FormatTime(coupon.EndTime)
	couponVO.Description = coupon.Description
	couponVO.Online = coupon.Online
	Ok(w, couponVO)
}

// DoEditCoupon 新增/删除 优惠券
func (m *MallHttpServiceImpl) DoEditCoupon(w http.ResponseWriter, r *http.Request) {
	req := new(CMSCouponReq)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		Error(w, errcode.BadRequestParam, "参数解析失败")
		return
	}
	if err := validator.New().Struct(req); err != nil {
		Error(w, errcode.BadRequestParam, "缺少参数")
		return
	}
	startTime, _ := time.Parse("2006-01-02 15:04:05", req.StartTime)
	endTime, _ := time.Parse("2006-01-02 15:04:05", req.EndTime)
	if req.Id == consts.ZERO {
		_, total, err := m.couponService.GetCouponList(r.Context(), 1, consts.CouponMax, consts.ALL)
		if err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
		if total >= consts.CouponMax {
			Error(w, errcode.NotAllowOperation, "最多只能添加"+strconv.Itoa(consts.CouponMax)+"张优惠券")
			return
		}
		coupon := &entity.WechatMallCouponDO{}
		coupon.Title = req.Title
		coupon.FullMoney = req.FullMoney
		coupon.Minus = req.Minus
		coupon.Rate = req.Rate
		coupon.Type = req.Type
		coupon.GrantNum = req.GrantNum
		coupon.LimitNum = req.LimitNum
		coupon.StartTime = startTime
		coupon.EndTime = endTime
		coupon.Description = req.Description
		coupon.Online = req.Online
		if err := m.couponService.AddCoupon(r.Context(), coupon); err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
	} else {
		coupon, err := m.couponService.GetCouponById(r.Context(), req.Id)
		if err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
		if coupon.ID == consts.ZERO || coupon.Del == consts.DELETE {
			Error(w, errcode.NotFoundCoupon, "优惠券不存在")
			return
		}
		coupon.Title = req.Title
		coupon.FullMoney = req.FullMoney
		coupon.Minus = req.Minus
		coupon.Rate = req.Rate
		coupon.Type = req.Type
		coupon.GrantNum = req.GrantNum
		coupon.LimitNum = req.LimitNum
		coupon.StartTime = startTime
		coupon.EndTime = endTime
		coupon.Description = req.Description
		coupon.Online = req.Online
		if err := m.couponService.UpdateCouponById(r.Context(), coupon); err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
	}
	Ok(w, "ok")
}

// DoDeleteCoupon 删除优惠券
func (m *MallHttpServiceImpl) DoDeleteCoupon(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	coupon, err := m.couponService.GetCouponById(r.Context(), id)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	if coupon.ID == consts.ZERO || coupon.Del == consts.DELETE {
		Error(w, errcode.NotFoundCoupon, "优惠券不存在")
		return
	}
	coupon.Del = consts.DELETE
	if err := m.couponService.UpdateCouponById(r.Context(), coupon); err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	Ok(w, "ok")
}
