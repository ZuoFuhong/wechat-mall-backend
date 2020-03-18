package portal

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"wechat-mall-backend/defs"
)

// 查询商品列表
func (h *Handler) GetGoodsList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	keyword := vars["k"]
	categoryId, _ := strconv.Atoi(vars["c"])
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])

	goodsList, total := h.service.GoodsService.QueryPortalGoodsList(keyword, categoryId, page, size)

	resp := make(map[string]interface{})
	resp["list"] = goodsList
	resp["total"] = total
	defs.SendNormalResponse(w, resp)
}

// 查询商品详情
func (h *Handler) GetGoodsDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	goodsId, _ := strconv.Atoi(vars["id"])
	goodsInfo := h.service.GoodsService.QueryPortalGoodsDetail(goodsId)

	defs.SendNormalResponse(w, goodsInfo)
}
