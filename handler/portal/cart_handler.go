package portal

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"wechat-mall-backend/defs"
	"wechat-mall-backend/errs"
)

// 查询-购物车商品
func (h *Handler) GetCartGoodsList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])
	userId := r.Context().Value(defs.ContextKey).(int)
	cartGoodsList, total := h.service.CartService.GetCartGoods(userId, page, size)

	resp := make(map[string]interface{})
	resp["list"] = cartGoodsList
	resp["total"] = total
	defs.SendNormalResponse(w, resp)
}

// 新增/更新-购物车商品
func (h *Handler) EditCartGoods(w http.ResponseWriter, r *http.Request) {
	req := defs.PortalCartGoodsReq{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		panic(errs.ErrorParameterValidate)
	}
	userId := r.Context().Value(defs.ContextKey).(int)

	h.service.CartService.DoEditCart(userId, req.GoodsId, req.SkuId, req.Num)
	defs.SendNormalResponse(w, "ok")
}
