package handler

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

func (h *CMSHandler) GetSKUList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])

	skuList, total := h.service.SKUService.GetSKUList(page, size)

	var skuVOs []model.SKU
	for _, v := range *skuList {
		skuVO := model.SKU{}
		skuVO.Id = v.Id
		skuVO.Title = v.Title
		skuVO.Price = v.Price
		skuVO.Code = v.Code
		skuVO.Stock = v.Stock
		skuVO.SpuId = v.SpuId
		skuVO.Online = v.Online
		skuVO.Picture = v.Picture
		skuVO.Specs = v.Specs
		skuVOs = append(skuVOs, skuVO)
	}
	resp := make(map[string]interface{})
	resp["list"] = skuVOs
	resp["total"] = total
	sendNormalResponse(w, resp)
}

func (h *CMSHandler) GetSKU(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	sku := h.service.SKUService.GetSKUById(id)
	if sku.Id == 0 {
		panic(errs.ErrorSKU)
	}
	skuVO := model.SKU{}
	skuVO.Id = sku.Id
	skuVO.Title = sku.Title
	skuVO.Price = sku.Price
	skuVO.Code = sku.Code
	skuVO.Stock = sku.Stock
	skuVO.SpuId = sku.SpuId
	skuVO.Online = sku.Online
	skuVO.Picture = sku.Picture
	skuVO.Specs = sku.Specs
	sendNormalResponse(w, skuVO)
}

func (h *CMSHandler) DoEditSKU(w http.ResponseWriter, r *http.Request) {
	req := defs.SKUReq{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		panic(errs.ErrorParameterValidate)
	}
	validate := validator.New()
	if err = validate.Struct(req); err != nil {
		panic(errs.NewParameterError(err.Error()))
	}
	if req.Id == 0 {
		sku := h.service.SKUService.GetSKUByCode(req.Code)
		if sku.Id != 0 {
			panic(errs.NewErrorSKU("The code already exists"))
		}
		sku.Title = req.Title
		sku.Price = req.Price
		sku.Code = req.Code
		sku.Stock = req.Stock
		sku.SpuId = req.SpuId
		sku.Online = req.Online
		sku.Picture = req.Picture
		sku.Specs = req.Specs
		h.service.SKUService.AddSKU(sku)
	} else {
		sku := h.service.SKUService.GetSKUByCode(req.Code)
		if sku.Id != req.Id {
			panic(errs.NewErrorSKU("The code already exists"))
		}
		sku = h.service.SKUService.GetSKUById(req.Id)
		if sku.Id == 0 {
			panic(errs.ErrorSKU)
		}
		sku.Title = req.Title
		sku.Price = req.Price
		sku.Code = req.Code
		sku.Stock = req.Stock
		sku.SpuId = req.SpuId
		sku.Online = req.Online
		sku.Picture = req.Picture
		sku.Specs = req.Specs
		h.service.SKUService.UpdateSKUById(sku)
	}
}

func (h *CMSHandler) DoDeleteSKU(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	sku := h.service.SKUService.GetSKUById(id)
	if sku.Id == 0 {
		panic(errs.ErrorSKU)
	}
	sku.Del = 1
	h.service.SKUService.UpdateSKUById(sku)
	sendNormalResponse(w, "ok")
}
