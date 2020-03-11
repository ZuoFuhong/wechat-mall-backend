package cms

import (
	"encoding/json"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"wechat-mall-backend/defs"
	"wechat-mall-backend/errs"
)

func (h *Handler) GetSKUList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])

	skuVOs := []defs.CMSSKUVO{}
	skuList, total := h.service.SKUService.GetSKUList(page, size)
	for _, v := range *skuList {
		skuVO := defs.CMSSKUVO{}
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

func (h *Handler) GetSKU(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	sku := h.service.SKUService.GetSKUById(id)
	if sku.Id == 0 {
		panic(errs.ErrorSKU)
	}
	skuVO := defs.CMSSKUVO{}
	skuVO.Id = sku.Id
	skuVO.Title = sku.Title
	skuVO.Price = sku.Price
	skuVO.Code = sku.Code
	skuVO.Stock = sku.Stock
	skuVO.GoodsId = sku.GoodsId
	skuVO.Online = sku.Online
	skuVO.Picture = sku.Picture
	skuVO.Specs = sku.Specs
	defs.SendNormalResponse(w, skuVO)
}

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
	if goods.Id == 0 {
		panic(errs.ErrorGoods)
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
		sku.GoodsId = req.GoodsId
		sku.Online = req.Online
		sku.Picture = req.Picture
		sku.Specs = req.Specs
		h.service.SKUService.AddSKU(sku)
	} else {
		sku := h.service.SKUService.GetSKUByCode(req.Code)
		if sku.Id != 0 && sku.Id != req.Id {
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
		sku.GoodsId = req.GoodsId
		sku.Online = req.Online
		sku.Picture = req.Picture
		sku.Specs = req.Specs
		h.service.SKUService.UpdateSKUById(sku)
	}
	defs.SendNormalResponse(w, "ok")
}

func (h *Handler) DoDeleteSKU(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	sku := h.service.SKUService.GetSKUById(id)
	if sku.Id == 0 {
		panic(errs.ErrorSKU)
	}
	sku.Del = 1
	h.service.SKUService.UpdateSKUById(sku)
	defs.SendNormalResponse(w, "ok")
}
