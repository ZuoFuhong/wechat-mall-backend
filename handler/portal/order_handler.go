package portal

import (
	"encoding/json"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/shopspring/decimal"
	"net/http"
	"strconv"
	"wechat-mall-backend/defs"
	"wechat-mall-backend/errs"
)

// 商城-下订单
func (h *Handler) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	req := &defs.PortalCartPlaceOrderReq{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		panic(errs.ErrorParameterValidate)
	}
	userId := r.Context().Value(defs.ContextKey).(int)

	dispatchAmount, err := decimal.NewFromString(req.DispatchAmount)
	if err != nil {
		panic(err)
	}
	expectAmount, err := decimal.NewFromString(req.ExpectAmount)
	if err != nil {
		panic(err)
	}
	resp := h.service.OrderService.GenerateOrder(userId, req.AddressId, req.CouponLogId, dispatchAmount, expectAmount, req.GoodsList)
	defs.SendNormalResponse(w, resp)
}

// 查询订单列表
func (h *Handler) GetOrderList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	status, _ := strconv.Atoi(vars["status"])
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])
	userId := r.Context().Value(defs.ContextKey).(int)

	orderList, total := h.service.OrderService.QueryOrderList(userId, status, page, size)
	resp := make(map[string]interface{})
	resp["list"] = orderList
	resp["total"] = total
	defs.SendNormalResponse(w, resp)
}

// 取消订单
func (h *Handler) CancelOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := r.Context().Value(defs.ContextKey).(int)

	orderId, _ := strconv.Atoi(vars["id"])
	h.service.OrderService.CancelOrder(userId, orderId)
	defs.SendNormalResponse(w, "ok")
}

// 删除订单
func (h *Handler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := r.Context().Value(defs.ContextKey).(int)

	orderId, _ := strconv.Atoi(vars["id"])
	h.service.OrderService.DeleteOrderRecord(userId, orderId)
	defs.SendNormalResponse(w, "ok")
}

// 订单-确认收货
func (h *Handler) ConfirmTakeGoods(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := r.Context().Value(defs.ContextKey).(int)

	orderId, _ := strconv.Atoi(vars["id"])
	h.service.OrderService.ConfirmTakeGoods(userId, orderId)
	defs.SendNormalResponse(w, "ok")
}

// 订单-退款申请
func (h *Handler) RefundApply(w http.ResponseWriter, r *http.Request) {
	req := defs.OrderRefundApplyReq{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		panic(err)
	}
	validate := validator.New()
	if err = validate.Struct(req); err != nil {
		panic(errs.NewParameterError(err.Error()))
	}
	userId := r.Context().Value(defs.ContextKey).(int)

	refundNo := h.service.OrderService.RefundApply(userId, req.OrderNo, req.Reason)
	resp := defs.OrderRefundApplyVO{RefundNo: refundNo}
	defs.SendNormalResponse(w, resp)
}

// 查询订单详情
func (h *Handler) GetOrderDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderNo := vars["orderNo"]
	userId := r.Context().Value(defs.ContextKey).(int)
	orderDetail := h.service.OrderService.QueryOrderDetail(userId, orderNo)
	defs.SendNormalResponse(w, orderDetail)
}

// 微信支付回调通知
func (h *Handler) WxPayNotify(w http.ResponseWriter, r *http.Request) {
	// todo: 解析数据，从 attach 字段获取订单号，响应微信服务器
	h.service.OrderService.OrderPaySuccessNotify("")
}
