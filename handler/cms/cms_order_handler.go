package cms

import (
	"encoding/json"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"wechat-mall-backend/defs"
	"wechat-mall-backend/errs"
)

// 订单-查询列表
func (h *Handler) GetOrderList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	status, _ := strconv.Atoi(vars["status"])
	searchType, _ := strconv.Atoi(vars["stype"])
	keyword := vars["k"]
	startTime := vars["st"]
	endTime := vars["et"]
	page, _ := strconv.Atoi(vars["p"])
	size, _ := strconv.Atoi(vars["s"])

	orderList, total := h.service.OrderService.QueryCMSOrderList(status, searchType, keyword, startTime, endTime, page, size)

	resp := map[string]interface{}{}
	resp["list"] = orderList
	resp["total"] = total
	defs.SendNormalResponse(w, resp)
}

// 订单-查询详情
func (h *Handler) GetOrderDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderNo := vars["orderNo"]
	orderVO := h.service.OrderService.QueryCMSOrderDetail(orderNo)
	if orderVO.OrderNo == "" {
		panic(errs.NewErrorOrder("订单不存在"))
	}
	defs.SendNormalResponse(w, orderVO)
}

// 订单-导出Excel
func (h *Handler) ExportOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	status, _ := strconv.Atoi(vars["status"])
	searchType, _ := strconv.Atoi(vars["stype"])
	keyword := vars["k"]
	startTime := vars["st"]
	endTime := vars["et"]

	ossLink := h.service.OrderService.ExportCMSOrderExcel(status, searchType, keyword, startTime, endTime)

	resp := map[string]interface{}{}
	resp["ossLink"] = ossLink
	defs.SendNormalResponse(w, resp)
}

// 订单-修改状态
func (h *Handler) ModifyOrderStatus(w http.ResponseWriter, r *http.Request) {
	req := defs.CMSModifyOrderStatusReq{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		panic(errs.ErrorParameterValidate)
	}
	validate := validator.New()
	if err = validate.Struct(req); err != nil {
		panic(errs.NewParameterError(err.Error()))
	}

	h.service.OrderService.ModifyOrderStatus(req.OrderNo, req.Otype)
	defs.SendNormalResponse(w, "ok")
}

// 订单-修改备注
func (h *Handler) ModifyOrderRemark(w http.ResponseWriter, r *http.Request) {
	req := defs.CMSModifyOrderRemarkReq{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		panic(errs.ErrorParameterValidate)
	}
	validate := validator.New()
	if err = validate.Struct(req); err != nil {
		panic(errs.NewParameterError(err.Error()))
	}

	h.service.OrderService.ModifyOrderRemark(req.OrderNo, req.Remark)
	defs.SendNormalResponse(w, "ok")
}

// 订单-商品改价
func (h *Handler) ModifyOrderGoods(w http.ResponseWriter, r *http.Request) {
	req := defs.CMSModifyOrderGoodsReq{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		panic(errs.ErrorParameterValidate)
	}
	validate := validator.New()
	if err = validate.Struct(req); err != nil {
		panic(errs.NewParameterError(err.Error()))
	}

	h.service.OrderService.ModifyOrderGoods(req.OrderNo, req.GoodsId, req.Price)
	defs.SendNormalResponse(w, "ok")
}
