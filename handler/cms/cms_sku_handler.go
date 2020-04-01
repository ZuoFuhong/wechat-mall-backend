package cms

import (
	"encoding/json"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"wechat-mall-backend/defs"
	"wechat-mall-backend/errs"
	"wechat-mall-backend/model"
)

// 查询-SKU列表
func (h *Handler) GetSKUList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	goodsId, _ := strconv.Atoi(vars["goodsId"])
	keyword := vars["k"]
	online, _ := strconv.Atoi(vars["o"])
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])

	skuVOs := []defs.CMSSkuListVO{}
	skuList, total := h.service.SKUService.GetSKUList(keyword, goodsId, online, page, size)
	for _, v := range *skuList {
		skuVO := defs.CMSSkuListVO{}
		skuVO.Id = v.Id
		skuVO.Title = v.Title
		skuVO.Price = v.Price
		skuVO.Code = v.Code
		skuVO.Stock = v.Stock
		skuVO.GoodsId = v.GoodsId
		skuVO.Online = v.Online
		skuVO.Picture = v.Picture
		skuVO.Specs = v.Specs
		skuVOs = append(skuVOs, skuVO)
	}
	resp := make(map[string]interface{})
	resp["list"] = skuVOs
	resp["total"] = total
	defs.SendNormalResponse(w, resp)
}

// 查询-单个SKU
func (h *Handler) GetSKU(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	sku := h.service.SKUService.GetSKUById(id)
	if sku.Id == defs.ZERO || sku.Del == defs.DELETE {
		panic(errs.ErrorSKU)
	}
	goodsDO := h.service.GoodsService.GetGoodsById(sku.GoodsId)
	if goodsDO.Id == defs.ZERO || goodsDO.Del == defs.DELETE {
		panic(errs.ErrorGoods)
	}
	categoryDO := h.service.CategoryService.GetCategoryById(goodsDO.CategoryId)
	if categoryDO.Id == defs.ZERO || categoryDO.Del == defs.DELETE {
		panic(errs.ErrorCategory)
	}
	skuVO := defs.CMSSkuDetailVO{}
	skuVO.Id = sku.Id
	skuVO.Title = sku.Title
	skuVO.Price = sku.Price
	skuVO.Code = sku.Code
	skuVO.Stock = sku.Stock
	skuVO.CategoryId = categoryDO.ParentId
	skuVO.SubCategoryId = categoryDO.Id
	skuVO.GoodsId = sku.GoodsId
	skuVO.Online = sku.Online
	skuVO.Picture = sku.Picture
	skuVO.Specs = sku.Specs
	defs.SendNormalResponse(w, skuVO)
}

// 新增/编辑 SKU
func (h *Handler) DoEditSKU(w http.ResponseWriter, r *http.Request) {
	req := defs.CMSSKUReq{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		panic(errs.ErrorParameterValidate)
	}
	validate := validator.New()
	if err = validate.Struct(req); err != nil {
		panic(errs.NewParameterError(err.Error()))
	}
	goods := h.service.GoodsService.GetGoodsById(req.GoodsId)
	if goods.Id == defs.ZERO || goods.Del == defs.DELETE {
		panic(errs.ErrorGoods)
	}
	if req.Id == defs.ZERO {
		sku := model.WechatMallSkuDO{}
		sku.Title = req.Title
		sku.Price = req.Price
		sku.Code = req.Code
		sku.Stock = req.Stock
		sku.GoodsId = req.GoodsId
		sku.Online = req.Online
		sku.Picture = req.Picture
		sku.Specs = req.Specs
		h.service.SKUService.AddSKU(&sku)
	} else {
		sku := h.service.SKUService.GetSKUById(req.Id)
		if sku.Id == defs.ZERO || sku.Del == defs.DELETE {
			panic(errs.ErrorSKU)
		}
		sku.Title = req.Title
		sku.Price = req.Price
		sku.Code = req.Code
		sku.Stock = req.Stock
		sku.GoodsId = req.GoodsId
		sku.Online = req.Online
		sku.Picture = req.Picture
		sku.Specs = req.Specs
		h.service.SKUService.UpdateSKUById(sku)
	}
	defs.SendNormalResponse(w, "ok")
}

// 删除-单个SKU
func (h *Handler) DoDeleteSKU(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	sku := h.service.SKUService.GetSKUById(id)
	if sku.Id == defs.ZERO || sku.Del == defs.DELETE {
		panic(errs.ErrorSKU)
	}
	sku.Del = defs.DELETE
	h.service.SKUService.UpdateSKUById(sku)
	defs.SendNormalResponse(w, "ok")
}
