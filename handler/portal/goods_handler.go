package portal

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"wechat-mall-backend/defs"
	"wechat-mall-backend/model"
)

// 查询商品列表
func (h *Handler) GetGoodsList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	keyword := vars["k"]
	sort, _ := strconv.Atoi(vars["s"])
	categoryId, _ := strconv.Atoi(vars["c"])
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])

	goodsList, total := h.service.GoodsService.QueryPortalGoodsList(keyword, sort, categoryId, page, size)

	resp := make(map[string]interface{})
	resp["list"] = goodsList
	resp["total"] = total
	defs.SendNormalResponse(w, resp)
}

// 查询商品详情
func (h *Handler) GetGoodsDetail(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(defs.ContextKey).(int)
	vars := mux.Vars(r)
	goodsId, _ := strconv.Atoi(vars["id"])
	goodsInfo := h.service.GoodsService.QueryPortalGoodsDetail(goodsId)
	go h.recordGoodsBrowse(userId, goodsInfo)

	defs.SendNormalResponse(w, goodsInfo)
}

// 浏览商品记录
func (h *Handler) recordGoodsBrowse(userId int, goods *defs.PortalGoodsInfo) {
	defer func() {
		err := recover()
		if err != nil {
			log.Print(err)
		}
	}()
	browse := model.WechatMallGoodsBrowseRecord{}
	browse.UserId = userId
	browse.GoodsId = goods.Id
	browse.Picture = goods.Picture
	browse.Title = goods.Title
	browse.Price = goods.Price
	h.service.BrowseRecordService.AddBrowseRecord(&browse)
}
