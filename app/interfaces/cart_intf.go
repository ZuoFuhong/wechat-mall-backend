package interfaces

import (
	"encoding/json"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"wechat-mall-backend/consts"
	"wechat-mall-backend/errcode"
	"wechat-mall-backend/pkg/log"
)

// GetCartGoodsList 查询-购物车商品
func (m *MallHttpServiceImpl) GetCartGoodsList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])
	userId := r.Context().Value(consts.ContextKey).(int)
	cartGoodsList, total, err := m.cartService.GetCartGoods(r.Context(), userId, page, size)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	data := make(map[string]interface{})
	data["list"] = cartGoodsList
	data["total"] = total
	Ok(w, data)
}

// AddCartGoods 添加-购物车商品
func (m *MallHttpServiceImpl) AddCartGoods(w http.ResponseWriter, r *http.Request) {
	req := new(PortalCartGoodsReq)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		Error(w, errcode.BadRequestParam, "参数解析失败")
		return
	}
	if err := validator.New().Struct(req); err != nil {
		Error(w, errcode.BadRequestParam, "缺少参数")
		return
	}
	userId := r.Context().Value(consts.ContextKey).(int)
	if err := m.cartService.DoEditCart(r.Context(), userId, req.GoodsId, req.SkuId, req.Num); err != nil {
		log.ErrorContextf(r.Context(), "call DoEditCart failed, err: %v", err)
		Error(w, errcode.NotAllowOperation, err.Error())
		return
	}
	Ok(w, "ok")
}

// EditCartGoods 编辑-购物车
func (m *MallHttpServiceImpl) EditCartGoods(w http.ResponseWriter, r *http.Request) {
	req := new(PortalEditCartReq)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		Error(w, errcode.BadRequestParam, "参数解析失败")
		return
	}
	userId := r.Context().Value(consts.ContextKey).(int)
	cartDO, err := m.cartService.GetCartDOById(r.Context(), req.Id)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	if cartDO.ID == consts.ZERO || cartDO.Del == consts.DELETE || cartDO.UserID != userId {
		Error(w, errcode.NotFoundCartGoods, "商品不存在")
		return
	}
	if req.Num == 0 {
		if err := m.cartService.DeleteCartDOById(r.Context(), userId, req.Id); err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
	} else {
		if err := m.cartService.DoEditCart(r.Context(), userId, cartDO.GoodsID, cartDO.SkuID, req.Num); err != nil {
			Error(w, errcode.NotAllowOperation, err.Error())
			return
		}
	}
	Ok(w, "ok")
}

// GetCartGoodsNum 查询-购物车商品数量
func (m *MallHttpServiceImpl) GetCartGoodsNum(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(consts.ContextKey).(int)
	goodsNum, err := m.cartService.CountCartGoodsNum(r.Context(), userId)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	data := map[string]interface{}{}
	data["num"] = goodsNum
	Ok(w, data)
}
