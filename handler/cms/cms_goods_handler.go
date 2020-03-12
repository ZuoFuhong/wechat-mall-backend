package cms

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"wechat-mall-backend/defs"
	"wechat-mall-backend/errs"
	"wechat-mall-backend/model"
)

func (h *Handler) GetGoodsList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])

	goodsVOList := []defs.CMSGoodsVO{}
	goodsList, total := h.service.GoodsService.GetGoodsList(page, size)
	for _, v := range *goodsList {
		goodsVO := defs.CMSGoodsVO{}
		goodsVO.Id = v.Id
		goodsVO.BrandName = v.BrandName
		goodsVO.Title = v.Title
		goodsVO.Price = v.Price
		goodsVO.DiscountPrice = v.DiscountPrice
		goodsVO.CategoryId = v.CategoryId
		goodsVO.Online = v.Online
		goodsVO.Picture = v.Picture
		goodsVO.BannerPicture = v.BannerPicture
		goodsVO.DetailPicture = v.DetailPicture
		goodsVO.Tags = v.Tags
		goodsVO.Description = v.Description
		goodsVOList = append(goodsVOList, goodsVO)
	}
	resp := make(map[string]interface{})
	resp["list"] = goodsVOList
	resp["total"] = total
	defs.SendNormalResponse(w, resp)
}

func (h *Handler) GetGoods(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	goodsDO := h.service.GoodsService.GetGoodsById(id)
	if goodsDO.Id == 0 {
		panic(errs.ErrorGoods)
	}
	goodsVO := defs.CMSGoodsVO{}
	goodsVO.Id = goodsDO.Id
	goodsVO.BrandName = goodsDO.BrandName
	goodsVO.Title = goodsDO.Title
	goodsVO.Price = goodsDO.Price
	goodsVO.DiscountPrice = goodsDO.DiscountPrice
	goodsVO.CategoryId = goodsDO.CategoryId
	goodsVO.Online = goodsDO.Online
	goodsVO.Picture = goodsDO.Picture
	goodsVO.BannerPicture = goodsDO.BannerPicture
	goodsVO.DetailPicture = goodsDO.DetailPicture
	goodsVO.Tags = goodsDO.Tags
	goodsVO.Description = goodsDO.Description
	defs.SendNormalResponse(w, goodsVO)
}

func (h *Handler) DoEditGoods(w http.ResponseWriter, r *http.Request) {
	req := defs.CMSGoodsReq{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		panic(errs.ErrorParameterValidate)
	}
	category := h.service.CategoryService.GetCategoryById(req.CategoryId)
	if category.Id == 0 {
		panic(errs.ErrorCategory)
	}
	if req.Id == 0 {
		goodsDO := model.WechatMallGoodsDO{}
		goodsDO.BrandName = req.BrandName
		goodsDO.Title = req.Title
		goodsDO.Price = req.Price
		goodsDO.DiscountPrice = req.DiscountPrice
		goodsDO.CategoryId = req.CategoryId
		goodsDO.Online = req.Online
		goodsDO.Picture = req.Picture
		goodsDO.BannerPicture = req.BannerPicture
		goodsDO.DetailPicture = req.DetailPicture
		goodsDO.Tags = req.Tags
		goodsDO.Description = req.Description
		goodsId := h.service.GoodsService.AddGoods(&goodsDO)
		h.service.GoodsService.AddGoodsSpec(goodsId, req.SpecList)
	} else {
		goodsDO := h.service.GoodsService.GetGoodsById(req.Id)
		if goodsDO.Id == 0 {
			panic(errs.ErrorGoods)
		}
		goodsDO.BrandName = req.BrandName
		goodsDO.Title = req.Title
		goodsDO.Price = req.Price
		goodsDO.DiscountPrice = req.DiscountPrice
		goodsDO.CategoryId = req.CategoryId
		goodsDO.Online = req.Online
		goodsDO.Picture = req.Picture
		goodsDO.BannerPicture = req.BannerPicture
		goodsDO.DetailPicture = req.DetailPicture
		goodsDO.Tags = req.Tags
		goodsDO.Description = req.Description
		h.service.GoodsService.UpdateGoodsById(goodsDO)
		h.service.GoodsService.AddGoodsSpec(req.Id, req.SpecList)
	}
	defs.SendNormalResponse(w, "ok")
}

func (h *Handler) DoDeleteGoods(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	goodsDO := h.service.GoodsService.GetGoodsById(id)
	if goodsDO.Id == 0 {
		panic(errs.ErrorGoods)
	}
	goodsDO.Del = 1
	h.service.GoodsService.UpdateGoodsById(goodsDO)
	defs.SendNormalResponse(w, "ok")
}

func (h *Handler) GetGoodsSpecList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	goodsId, _ := strconv.Atoi(vars["id"])
	specList := h.service.GoodsService.GetGoodsSpecList(goodsId)
	defs.SendNormalResponse(w, specList)
}
