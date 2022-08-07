package interfaces

import (
	"encoding/json"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"wechat-mall-backend/errcode"
)

// GetCmsOrderList GetOrderList 订单-查询列表
func (m *MallHttpServiceImpl) GetCmsOrderList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	status, _ := strconv.Atoi(vars["status"])
	searchType, _ := strconv.Atoi(vars["stype"])
	keyword := vars["k"]
	startTime := vars["st"]
	endTime := vars["et"]
	page, _ := strconv.Atoi(vars["p"])
	size, _ := strconv.Atoi(vars["s"])

	orderList, total, err := m.orderService.QueryCMSOrderList(r.Context(), status, searchType, keyword, startTime, endTime, page, size)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	data := map[string]interface{}{}
	data["list"] = orderList
	data["total"] = total
	Ok(w, data)
}

// GetCmsOrderDetail 订单-查询详情
func (m *MallHttpServiceImpl) GetCmsOrderDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderNo := vars["orderNo"]
	orderVO, err := m.orderService.QueryCMSOrderDetail(r.Context(), orderNo)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	if orderVO.OrderNo == "" {
		Error(w, errcode.NotFoundOrderRefund, "订单不存在")
		return
	}
	Ok(w, orderVO)
}

// ExportOrder 订单-导出Excel
func (m *MallHttpServiceImpl) ExportOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	status, _ := strconv.Atoi(vars["status"])
	searchType, _ := strconv.Atoi(vars["stype"])
	keyword := vars["k"]
	startTime := vars["st"]
	endTime := vars["et"]

	ossLink, err := m.orderService.ExportCMSOrderExcel(r.Context(), status, searchType, keyword, startTime, endTime)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	data := map[string]interface{}{}
	data["ossLink"] = ossLink
	Ok(w, data)
}

// ModifyOrderStatus 订单-修改状态
func (m *MallHttpServiceImpl) ModifyOrderStatus(w http.ResponseWriter, r *http.Request) {
	req := new(CMSModifyOrderStatusReq)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		Error(w, errcode.BadRequestParam, "参数解析失败")
		return
	}
	if err := validator.New().Struct(req); err != nil {
		Error(w, errcode.BadRequestParam, "缺少参数")
		return
	}
	if err := m.orderService.ModifyOrderStatus(r.Context(), req.OrderNo, req.Otype); err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	Ok(w, "ok")
}

// ModifyOrderRemark 订单-修改备注
func (m *MallHttpServiceImpl) ModifyOrderRemark(w http.ResponseWriter, r *http.Request) {
	req := new(CMSModifyOrderRemarkReq)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		Error(w, errcode.BadRequestParam, "参数解析失败")
		return
	}
	if err := validator.New().Struct(req); err != nil {
		Error(w, errcode.BadRequestParam, "缺少参数")
		return
	}
	if err := m.orderService.ModifyOrderRemark(r.Context(), req.OrderNo, req.Remark); err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	Ok(w, "ok")
}

// ModifyOrderGoods 订单-商品改价
func (m *MallHttpServiceImpl) ModifyOrderGoods(w http.ResponseWriter, r *http.Request) {
	req := new(CMSModifyOrderGoodsReq)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		Error(w, errcode.BadRequestParam, "参数解析失败")
		return
	}
	if err := validator.New().Struct(req); err != nil {
		Error(w, errcode.BadRequestParam, "缺少参数")
		return
	}
	if err := m.orderService.ModifyOrderGoods(r.Context(), req.OrderNo, req.GoodsId, req.Price); err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	Ok(w, "ok")
}
