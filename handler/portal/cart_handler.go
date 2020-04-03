package portal

import (
	"encoding/json"
	"github.com/go-playground/validator"
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

// 添加-购物车商品
func (h *Handler) AddCartGoods(w http.ResponseWriter, r *http.Request) {
	req := defs.PortalCartGoodsReq{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		panic(errs.ErrorParameterValidate)
	}
	validate := validator.New()
	if err = validate.Struct(req); err != nil {
		panic(errs.NewParameterError(err.Error()))
	}
	userId := r.Context().Value(defs.ContextKey).(int)

	h.service.CartService.DoEditCart(userId, req.GoodsId, req.SkuId, req.Num)
	defs.SendNormalResponse(w, "ok")
}

// 编辑-购物车
func (h *Handler) EditCartGoods(w http.ResponseWriter, r *http.Request) {
	req := defs.PortalEditCartReq{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		panic(errs.ErrorParameterValidate)
	}
	userId := r.Context().Value(defs.ContextKey).(int)
	cartDO := h.service.CartService.GetCartDOById(req.Id)
	if cartDO.Id == defs.ZERO || cartDO.Del == defs.DELETE || cartDO.UserId != userId {
		panic(errs.ErrorGoodsCart)
	}
	if req.Num == 0 {
		h.service.CartService.DeleteCartDOById(userId, req.Id)
	} else {
		h.service.CartService.DoEditCart(userId, cartDO.GoodsId, cartDO.SkuId, req.Num)
	}
	defs.SendNormalResponse(w, "ok")
}

// 查询-购物车商品数量
func (h *Handler) GetCartGoodsNum(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(defs.ContextKey).(int)

	goodsNum := h.service.CartService.CountCartGoodsNum(userId)

	resp := map[string]interface{}{}
	resp["num"] = goodsNum
	defs.SendNormalResponse(w, resp)
}
