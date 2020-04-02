package portal

import (
	"encoding/json"
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

// 查询订单详情
func (h *Handler) GetOrderDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderId, _ := strconv.Atoi(vars["id"])
	userId := r.Context().Value(defs.ContextKey).(int)
	orderDetail := h.service.OrderService.QueryOrderDetail(userId, orderId)
	defs.SendNormalResponse(w, orderDetail)
}

// 微信支付回调通知
func (h *Handler) WxPayNotify(w http.ResponseWriter, r *http.Request) {
	// todo: 解析数据，从 attach 字段获取订单号，响应微信服务器

	h.service.OrderService.OrderPaySuccessNotify("")
}
