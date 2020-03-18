package cms

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"wechat-mall-backend/defs"
)

// 查询订单列表
func (h *Handler) GetOrderList(w http.ResponseWriter, r *http.Request) {

	// todo: 查询订单列表
}

// 查询-订单报表数据
func (h *Handler) GetSaleTableData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])

	saleData := h.service.OrderService.QueryOrderSaleData(page, size)
	defs.SendNormalResponse(w, saleData)
}
