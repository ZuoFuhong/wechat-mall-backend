package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"wechat-mall-backend/defs"
	"wechat-mall-backend/errs"
	"wechat-mall-backend/model"
)

func (h *CMSHandler) GetSPUList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])

	spuList, total := h.service.SPUService.GetSPUList(page, size)
	var spuVOList []defs.SPUVO
	for _, v := range *spuList {
		spuVO := defs.SPUVO{}
		spuVO.Id = v.Id
		spuVO.BrandName = v.BrandName
		spuVO.Title = v.Title
		spuVO.SubTitle = v.SubTitle
		spuVO.Price = v.Price
		spuVO.DiscountPrice = v.DiscountPrice
		spuVO.CategoryId = v.CategoryId
		spuVO.DefaultSkuId = v.DefaultSkuId
		spuVO.Online = v.Online
		spuVO.Picture = v.Picture
		spuVO.ForThemePicture = v.ForThemePicture
		spuVO.BannerPicture = v.BannerPicture
		spuVO.DetailPicture = v.DetailPicture
		spuVO.Tags = v.Tags
		spuVO.SketchSpecId = v.SketchSpecId
		spuVO.Description = v.Description
		spuVOList = append(spuVOList, spuVO)
	}
	resp := make(map[string]interface{})
	resp["list"] = spuVOList
	resp["total"] = total
	sendNormalResponse(w, resp)
}

func (h *CMSHandler) GetSPU(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	spuDO := h.service.SPUService.GetSPUById(id)
	if spuDO.Id == 0 {
		panic(errs.ErrorSPU)
	}
	spuVO := defs.SPUVO{}
	spuVO.Id = spuDO.Id
	spuVO.BrandName = spuDO.BrandName
	spuVO.Title = spuDO.Title
	spuVO.SubTitle = spuDO.SubTitle
	spuVO.Price = spuDO.Price
	spuVO.DiscountPrice = spuDO.DiscountPrice
	spuVO.CategoryId = spuDO.CategoryId
	spuVO.DefaultSkuId = spuDO.DefaultSkuId
	spuVO.Online = spuDO.Online
	spuVO.Picture = spuDO.Picture
	spuVO.ForThemePicture = spuDO.ForThemePicture
	spuVO.BannerPicture = spuDO.BannerPicture
	spuVO.DetailPicture = spuDO.DetailPicture
	spuVO.Tags = spuDO.Tags
	spuVO.SketchSpecId = spuDO.SketchSpecId
	spuVO.Description = spuDO.Description
	sendNormalResponse(w, spuVO)
}

func (h *CMSHandler) DoEditSPU(w http.ResponseWriter, r *http.Request) {
	req := defs.SPUReq{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		panic(errs.ErrorParameterValidate)
	}
	if req.Id == 0 {
		spuDO := model.SPU{}
		spuDO.BrandName = req.BrandName
		spuDO.Title = req.Title
		spuDO.SubTitle = req.SubTitle
		spuDO.Price = req.Price
		spuDO.DiscountPrice = req.DiscountPrice
		spuDO.CategoryId = req.CategoryId
		spuDO.DefaultSkuId = req.DefaultSkuId
		spuDO.Online = req.Online
		spuDO.Picture = req.Picture
		spuDO.ForThemePicture = req.ForThemePicture
		spuDO.BannerPicture = req.BannerPicture
		spuDO.DetailPicture = req.DetailPicture
		spuDO.Tags = req.Tags
		spuDO.SketchSpecId = req.SketchSpecId
		spuDO.Description = req.Description
		spuId := h.service.SPUService.AddSPU(&spuDO)
		h.service.SPUService.AddSPUSpec(spuId, req.SpecList)
	} else {
		spuDO := h.service.SPUService.GetSPUById(req.Id)
		if spuDO.Id == 0 {
			panic(errs.ErrorSPU)
		}
		spuDO.BrandName = req.BrandName
		spuDO.Title = req.Title
		spuDO.SubTitle = req.SubTitle
		spuDO.Price = req.Price
		spuDO.DiscountPrice = req.DiscountPrice
		spuDO.CategoryId = req.CategoryId
		spuDO.DefaultSkuId = req.DefaultSkuId
		spuDO.Online = req.Online
		spuDO.Picture = req.Picture
		spuDO.ForThemePicture = req.ForThemePicture
		spuDO.BannerPicture = req.BannerPicture
		spuDO.DetailPicture = req.DetailPicture
		spuDO.Tags = req.Tags
		spuDO.SketchSpecId = req.SketchSpecId
		spuDO.Description = req.Description
		h.service.SPUService.UpdateSPUById(spuDO)
		h.service.SPUService.AddSPUSpec(req.Id, req.SpecList)
	}
	sendNormalResponse(w, "ok")
}

func (h *CMSHandler) DoDeleteSPU(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	spuDO := h.service.SPUService.GetSPUById(id)
	if spuDO.Id == 0 {
		panic(errs.ErrorSPU)
	}
	spuDO.Del = 1
	h.service.SPUService.UpdateSPUById(spuDO)
	sendNormalResponse(w, "ok")
}

func (h *CMSHandler) GetSPUSpecList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	specList := h.service.SPUService.GetSPUSpecList(id)
	sendNormalResponse(w, specList)
}
