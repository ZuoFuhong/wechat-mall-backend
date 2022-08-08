package interfaces

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/shopspring/decimal"
	"log"
	"net/http"
	"strconv"
	"wechat-mall-backend/app/domain/entity"
	"wechat-mall-backend/app/domain/view"
	"wechat-mall-backend/consts"
	"wechat-mall-backend/errcode"
)

// GetGoodsList 查询商品列表
func (m *MallHttpServiceImpl) GetGoodsList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	keyword := vars["k"]
	sort, _ := strconv.Atoi(vars["s"])
	categoryId, _ := strconv.Atoi(vars["c"])
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])

	if categoryId == 0 {
		categoryId = consts.ALL
	}
	goodsList, total, err := m.goodsService.QueryPortalGoodsList(r.Context(), keyword, sort, categoryId, page, size)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	data := make(map[string]interface{})
	data["list"] = goodsList
	data["total"] = total
	Ok(w, data)
}

// GetGoodsDetail 查询商品详情
func (m *MallHttpServiceImpl) GetGoodsDetail(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(consts.ContextKey).(int)
	vars := mux.Vars(r)
	goodsId, _ := strconv.Atoi(vars["id"])
	goodsInfo, err := m.goodsService.QueryPortalGoodsDetail(r.Context(), goodsId)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	go m.recordGoodsBrowse(r.Context(), userId, goodsInfo)

	Ok(w, goodsInfo)
}

// 浏览商品记录
func (m *MallHttpServiceImpl) recordGoodsBrowse(ctx context.Context, userId int, goods *view.PortalGoodsInfo) {
	defer func() {
		err := recover()
		if err != nil {
			log.Print(err)
		}
	}()
	browse := &entity.WechatMallGoodsBrowseRecord{}
	browse.UserID = userId
	browse.GoodsID = goods.Id
	browse.Picture = goods.Picture
	browse.Title = goods.Title
	browse.Price = decimal.NewFromFloat(goods.Price).String()
	_ = m.browseService.AddBrowseRecord(ctx, browse)
}

// ClearBrowseHistory 清理-浏览历史
func (m *MallHttpServiceImpl) ClearBrowseHistory(w http.ResponseWriter, r *http.Request) {
	var ids []int
	if err := json.NewDecoder(r.Body).Decode(&ids); err != nil {
		Error(w, errcode.BadRequestParam, "参数解析失败")
		return
	}
	m.browseService.ClearBrowseHistory(r.Context(), ids)
	Ok(w, "ok")
}
