package interfaces

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"wechat-mall-backend/app/domain/view"
	"wechat-mall-backend/consts"
	"wechat-mall-backend/errcode"
	"wechat-mall-backend/pkg/config"
	"wechat-mall-backend/pkg/utils"
)

// GetOSSPolicyToken 生成-OSSPolicyToken
func (m *MallHttpServiceImpl) GetOSSPolicyToken(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dir := vars["dir"]

	ossConf := config.GlobalConfig().Oss
	oss := utils.OSSPolicyToken{
		AccessKeyId:     ossConf.AccessKeyId,
		AccessKeySecret: ossConf.AccessKeySecret,
		Host:            "https://" + ossConf.BucketName + ".oss-cn-hangzhou.aliyuncs.com",
		UploadDir:       dir,
		ExpireTime:      30,
	}
	response := oss.GetPolicyToken()
	Ok(w, response)
}

// GetMarketMetrics CMS-统计商城指标信息
func (m *MallHttpServiceImpl) GetMarketMetrics(w http.ResponseWriter, r *http.Request) {
	visitorNum := m.userService.QueryTodayUniqueVisitor(r.Context())
	sellOutSKU, _ := m.skuService.CountSellOutSKU(r.Context())
	waitingOrder, _ := m.orderService.CountOrderNum(r.Context(), consts.ALL, 1)
	refundOrder, _ := m.orderService.CountPendingOrderRefund(r.Context())

	metricsVO := &view.CMSMarketMetricsVO{}
	metricsVO.VisitorNum = visitorNum
	metricsVO.SellOutSKUNum = sellOutSKU
	metricsVO.WaitingOrder = waitingOrder
	metricsVO.ActivistOrder = refundOrder
	Ok(w, metricsVO)
}

// GetSaleTableData 查询-订单报表数据
func (m *MallHttpServiceImpl) GetSaleTableData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])

	saleData, err := m.orderService.QueryOrderSaleData(r.Context(), page, size)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	Ok(w, saleData)
}
