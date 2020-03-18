package cms

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"wechat-mall-backend/dbops"
	"wechat-mall-backend/defs"
	"wechat-mall-backend/errs"
	"wechat-mall-backend/model"
)

// 查询-商品列表
func (h *Handler) GetGoodsList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	keyword, _ := vars["k"]
	categoryId, _ := strconv.Atoi(vars["c"])
	online, _ := strconv.Atoi(vars["o"])
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])

	goodsVOList := []defs.CMSGoodsListVO{}
	goodsList, total := h.service.GoodsService.GetGoodsList(keyword, categoryId, online, page, size)
	for _, v := range *goodsList {
		categoryDO, err := dbops.QueryCategoryById(v.CategoryId)
		if err != nil {
			panic(errs.ErrorCategory)
		}
		goodsVO := defs.CMSGoodsListVO{}
		goodsVO.Id = v.Id
		goodsVO.BrandName = v.BrandName
		goodsVO.Title = v.Title
		goodsVO.Price = v.Price
		goodsVO.CategoryId = v.CategoryId
		goodsVO.CategoryName = categoryDO.Name
		goodsVO.Online = v.Online
		goodsVO.Picture = v.Picture
		goodsVOList = append(goodsVOList, goodsVO)
	}
	resp := make(map[string]interface{})
	resp["list"] = goodsVOList
	resp["total"] = total
	defs.SendNormalResponse(w, resp)
}

// 查询商品详情
func (h *Handler) GetGoods(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	goodsDO := h.service.GoodsService.GetGoodsById(id)
	if goodsDO.Id == 0 {
		panic(errs.ErrorGoods)
	}
	categoryDO, err := dbops.QueryCategoryById(goodsDO.CategoryId)
	if err != nil {
		panic(errs.ErrorCategory)
	}
	goodsVO := defs.CMSGoodsVO{}
	goodsVO.Id = goodsDO.Id
	goodsVO.BrandName = goodsDO.BrandName
	goodsVO.Title = goodsDO.Title
	goodsVO.Price = goodsDO.Price
	goodsVO.DiscountPrice = goodsDO.DiscountPrice
	goodsVO.CategoryId = categoryDO.ParentId
	goodsVO.SubCategoryId = goodsDO.CategoryId
	goodsVO.CategoryName = categoryDO.Name
	goodsVO.Online = goodsDO.Online
	goodsVO.Picture = goodsDO.Picture
	goodsVO.BannerPicture = goodsDO.BannerPicture
	goodsVO.DetailPicture = goodsDO.DetailPicture
	goodsVO.Tags = goodsDO.Tags
	defs.SendNormalResponse(w, goodsVO)
}

// 新增/编辑商品
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
		h.service.GoodsService.UpdateGoodsById(goodsDO)
		h.service.GoodsService.AddGoodsSpec(req.Id, req.SpecList)
	}
	defs.SendNormalResponse(w, "ok")
}

// 删除商品
func (h *Handler) DoDeleteGoods(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	goodsDO := h.service.GoodsService.GetGoodsById(id)
	if goodsDO.Id == 0 {
		panic(errs.ErrorGoods)
	}
	_, total := h.service.SKUService.GetSKUList(id, 1, 1)
	if total > 0 {
		panic(errs.NewErrorGoods("该商品下有SKU，不能删除！"))
	}
	goodsDO.Del = 1
	h.service.GoodsService.UpdateGoodsById(goodsDO)
	defs.SendNormalResponse(w, "ok")
}

// 查询-商品-规格属性
func (h *Handler) GetGoodsSpecList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	goodsId, _ := strconv.Atoi(vars["id"])
	specList := h.service.GoodsService.GetGoodsSpecList(goodsId)
	defs.SendNormalResponse(w, specList)
}

// Sku查询-全部分类及商品
func (h *Handler) GetChooseCategoryGoods(w http.ResponseWriter, r *http.Request) {
	tmpCategoryList := []map[string]interface{}{}
	categoryList, _ := h.service.CategoryService.GetCategoryList(0, 0, 0)
	for _, v := range *categoryList {
		tmpSubCategoryList := []map[string]interface{}{}
		subCategoryList, _ := h.service.CategoryService.GetCategoryList(v.Id, 0, 0)
		for _, sv := range *subCategoryList {
			tmpGoodsList := []map[string]interface{}{}
			goodsList, _ := h.service.GoodsService.GetGoodsList("", sv.Id, -1, 0, 0)
			for _, g := range *goodsList {
				goodsItem := map[string]interface{}{}
				goodsItem["value"] = g.Id
				goodsItem["label"] = g.Title
				tmpGoodsList = append(tmpGoodsList, goodsItem)
			}
			tmpSubCategory := map[string]interface{}{}
			tmpSubCategory["value"] = sv.Id
			tmpSubCategory["label"] = sv.Name
			tmpSubCategory["children"] = tmpGoodsList
			tmpSubCategoryList = append(tmpSubCategoryList, tmpSubCategory)
		}
		tmpCategory := map[string]interface{}{}
		tmpCategory["value"] = v.Id
		tmpCategory["label"] = v.Name
		tmpCategory["children"] = tmpSubCategoryList
		tmpCategoryList = append(tmpCategoryList, tmpCategory)
	}
	defs.SendNormalResponse(w, tmpCategoryList)
}
