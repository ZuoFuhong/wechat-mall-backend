package cms

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"wechat-mall-backend/defs"
)

// CMS-统计商城指标信息
func (h *Handler) GetMarketMetrics(w http.ResponseWriter, r *http.Request) {
	visitorNum := h.service.UserService.QueryTodayUniqueVisitor()
	sellOutSKU := h.service.SKUService.CountSellOutSKU()
	waitingOrder := h.service.OrderService.CountWaitingOrderNum(1)
	refundOrder := h.service.OrderService.CountPendingOrderRefund()

	metricsVO := defs.CMSMarketMetricsVO{}
	metricsVO.VisitorNum = visitorNum
	metricsVO.SellOutSKUNum = sellOutSKU
	metricsVO.WaitingOrder = waitingOrder
	metricsVO.ActivistOrder = refundOrder
	defs.SendNormalResponse(w, metricsVO)
}

// 查询-订单报表数据
func (h *Handler) GetSaleTableData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])

	saleData := h.service.OrderService.QueryOrderSaleData(page, size)
	defs.SendNormalResponse(w, saleData)
}
