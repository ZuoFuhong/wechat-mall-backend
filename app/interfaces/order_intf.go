package interfaces

import (
	"encoding/json"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/shopspring/decimal"
	"net/http"
	"strconv"
	"wechat-mall-backend/app/domain/view"
	"wechat-mall-backend/consts"
	"wechat-mall-backend/errcode"
	"wechat-mall-backend/pkg/log"
)

// PlaceOrder 订单-C端下订单
func (m *MallHttpServiceImpl) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	req := new(PortalCartPlaceOrderReq)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		Error(w, errcode.BadRequestParam, "参数解析失败")
		return
	}
	userId := r.Context().Value(consts.ContextKey).(int)
	dispatchAmount, err := decimal.NewFromString(req.DispatchAmount)
	if err != nil {
		Error(w, errcode.BadRequestParam, "金额错误")
		return
	}
	expectAmount, err := decimal.NewFromString(req.ExpectAmount)
	if err != nil {
		Error(w, errcode.BadRequestParam, "金额错误")
		return
	}
	data, err := m.orderService.GenerateOrder(r.Context(), userId, req.AddressId, req.CouponLogId, dispatchAmount, expectAmount, req.GoodsList)
	if err != nil {
		log.ErrorContextf(r.Context(), "call GenerateOrder failed, err: %v", err)
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	Ok(w, data)
}

// GetOrderList 订单-C端列表
func (m *MallHttpServiceImpl) GetOrderList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	status, _ := strconv.Atoi(vars["status"])
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])
	userId := r.Context().Value(consts.ContextKey).(int)

	orderList, total, err := m.orderService.QueryOrderList(r.Context(), userId, status, page, size)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	data := make(map[string]interface{})
	data["list"] = orderList
	data["total"] = total
	Ok(w, data)
}

// CancelOrder 订单-C端取消订单
func (m *MallHttpServiceImpl) CancelOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := r.Context().Value(consts.ContextKey).(int)

	orderId, _ := strconv.Atoi(vars["id"])
	if err := m.orderService.CancelOrder(r.Context(), userId, orderId); err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	Ok(w, "ok")
}

// DeleteOrder 订单-C端删除订单
func (m *MallHttpServiceImpl) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := r.Context().Value(consts.ContextKey).(int)

	orderId, _ := strconv.Atoi(vars["id"])
	if err := m.orderService.DeleteOrderRecord(r.Context(), userId, orderId); err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	Ok(w, "ok")
}

// ConfirmTakeGoods 订单-C端确认收货
func (m *MallHttpServiceImpl) ConfirmTakeGoods(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := r.Context().Value(consts.ContextKey).(int)

	orderId, _ := strconv.Atoi(vars["id"])
	if err := m.orderService.ConfirmTakeGoods(r.Context(), userId, orderId); err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	Ok(w, "ok")
}

// GetOrderDetail 订单-C端订单详情
func (m *MallHttpServiceImpl) GetOrderDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderNo := vars["orderNo"]
	userId := r.Context().Value(consts.ContextKey).(int)
	orderDetail, err := m.orderService.QueryOrderDetail(r.Context(), userId, orderNo)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	Ok(w, orderDetail)
}

// GetOrderRemind 订单-红点提醒
func (m *MallHttpServiceImpl) GetOrderRemind(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(consts.ContextKey).(int)

	remindVO := new(view.OrderRemindVO)
	remindVO.WaitPay, _ = m.orderService.CountOrderNum(r.Context(), userId, 0)
	remindVO.NotExpress, _ = m.orderService.CountOrderNum(r.Context(), userId, 1)
	remindVO.WaitReceive, _ = m.orderService.CountOrderNum(r.Context(), userId, 2)
	Ok(w, remindVO)
}

// WxPayNotify 订单-微信支付回调
func (m *MallHttpServiceImpl) WxPayNotify(w http.ResponseWriter, r *http.Request) {
	// todo: 解析数据，从 attach 字段获取订单号，响应微信服务器
	_ = m.orderService.OrderPaySuccessNotify(r.Context(), "")
}

// RefundApply 退款-C端退款申请
func (m *MallHttpServiceImpl) RefundApply(w http.ResponseWriter, r *http.Request) {
	req := new(OrderRefundApplyReq)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		Error(w, errcode.BadRequestParam, "参数解析失败")
		return
	}
	if err := validator.New().Struct(req); err != nil {
		Error(w, errcode.BadRequestParam, "缺少参数")
		return
	}
	userId := r.Context().Value(consts.ContextKey).(int)
	refundNo, err := m.orderService.RefundApply(r.Context(), userId, req.OrderNo, req.Reason)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	resp := view.OrderRefundApplyVO{RefundNo: refundNo}
	Ok(w, resp)
}

// RefundDetail 退款-C端退款详情
func (m *MallHttpServiceImpl) RefundDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	refundNo := vars["refundNo"]
	userId := r.Context().Value(consts.ContextKey).(int)

	refundDetail, err := m.orderService.QueryRefundDetail(r.Context(), userId, refundNo)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	Ok(w, refundDetail)
}

// UndoRefundApply 退款-C端撤销申请
func (m *MallHttpServiceImpl) UndoRefundApply(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	refundNo := vars["refundNo"]
	userId := r.Context().Value(consts.ContextKey).(int)
	if err := m.orderService.UndoRefundApply(r.Context(), userId, refundNo); err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	Ok(w, "ok")
}
