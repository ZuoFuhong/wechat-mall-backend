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
	"wechat-mall-backend/service"
)

func (h *CMSHandler) GetBannerList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])
	bannerList, total := h.service.BannerService.GetBannerList(page, size)

	var voList []defs.BannerVO
	for _, v := range *bannerList {
		vo := defs.BannerVO{}
		vo.Id = v.Id
		vo.Name = v.Name
		vo.Picture = v.Picture
		vo.Title = v.Title
		vo.Description = v.Description
		voList = append(voList, vo)
	}
	resp := make(map[string]interface{}, 0)
	resp["list"] = voList
	resp["total"] = total
	sendNormalResponse(w, resp)
}

func (h *CMSHandler) GetBanner(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	banner := h.service.BannerService.GetBannerById(id)
	if banner.Id == 0 {
		panic(errs.ErrorBannerNotExist)
	}
	bVO := defs.BannerVO{}
	bVO.Id = banner.Id
	bVO.Picture = banner.Picture
	bVO.Name = banner.Name
	bVO.Title = banner.Title
	bVO.Description = banner.Description
	sendNormalResponse(w, bVO)
}

func (h *CMSHandler) DoEditBanner(w http.ResponseWriter, r *http.Request) {
	req := defs.CMSBannerReq{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		panic(errs.ErrorRequestBodyParseFailed)
	}
	validate := validator.New()
	if err = validate.Struct(req); err != nil {
		panic(errs.NewParameterError(err.Error()))
	}
	if req.Id == 0 {
		banner := model.Banner{}
		banner.Picture = req.Picture
		banner.Name = req.Name
		banner.Title = req.Title
		banner.Description = req.Description
		h.service.BannerService.AddBanner((*service.Banner)(&banner))
	} else {
		banner := h.service.BannerService.GetBannerById(req.Id)
		if banner.Id == 0 {
			panic(errs.ErrorBannerNotExist)
		}
		banner.Picture = req.Picture
		banner.Name = req.Name
		banner.Title = req.Title
		banner.Description = req.Description
		h.service.BannerService.UpdateBannerById(banner)
	}
	sendNormalResponse(w, "ok")
}

func (h *CMSHandler) DoDeleteBanner(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	banner := h.service.BannerService.GetBannerById(id)
	if banner.Id == 0 {
		panic(errs.ErrorBannerNotExist)
	}
	banner.Del = 1
	h.service.BannerService.UpdateBannerById(banner)
	sendNormalResponse(w, "ok")
}

func (h *CMSHandler) GetBannerItemList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bannerId, _ := strconv.Atoi(vars["id"])
	itemList := h.service.BannerService.GetBannerItemList(bannerId)

	var itemVOList []defs.BannerItemVO
	for _, v := range *itemList {
		itemVO := defs.BannerItemVO{}
		itemVO.Id = v.Id
		itemVO.BannerId = v.BannerId
		itemVO.Name = v.Name
		itemVO.Picture = v.Picture
		itemVO.Keyword = v.Keyword
		itemVO.Type = v.Type
		itemVOList = append(itemVOList, itemVO)
	}
	sendNormalResponse(w, itemVOList)
}

func (h *CMSHandler) GetBannerItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bannerItemId, _ := strconv.Atoi(vars["id"])
	bannerItem := h.service.BannerService.GetBannerItemById(bannerItemId)
	if bannerItem.Id == 0 {
		panic(errs.ErrorBannerNotExist)
	}
	itemVO := defs.BannerItemVO{}
	itemVO.Id = bannerItem.Id
	itemVO.BannerId = bannerItem.BannerId
	itemVO.Name = bannerItem.Name
	itemVO.Picture = bannerItem.Picture
	itemVO.Keyword = bannerItem.Keyword
	itemVO.Type = bannerItem.Type
	sendNormalResponse(w, itemVO)
}

func (h *CMSHandler) DoEditBannerItem(w http.ResponseWriter, r *http.Request) {
	req := defs.BannerItemReq{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		panic(errs.ErrorParameterValidate)
	}
	validate := validator.New()
	if err = validate.Struct(req); err != nil {
		panic(errs.NewParameterError(err.Error()))
	}
	if req.Id == 0 {
		item := model.BannerItem{}
		item.BannerId = req.BannerId
		item.Name = req.Name
		item.Picture = req.Picture
		item.Keyword = req.Keyword
		item.Type = req.Type
		h.service.BannerService.AddBannerItem((*service.BannerItem)(&item))
	} else {
		item := h.service.BannerService.GetBannerItemById(req.Id)
		if item.Id == 0 {
			panic(errs.ErrorBannerNotExist)
		}
		item.BannerId = req.BannerId
		item.Name = req.Name
		item.Picture = req.Picture
		item.Keyword = req.Keyword
		item.Type = req.Type
		h.service.BannerService.UpdateBannerItemById(item)
	}
	sendNormalResponse(w, "ok")
}

func (h *CMSHandler) DoDeleteBannerItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	bannerItem := h.service.BannerService.GetBannerItemById(id)
	if bannerItem.Id == 0 {
		panic(errs.ErrorBannerNotExist)
	}
	bannerItem.Del = 1
	h.service.BannerService.UpdateBannerItemById(bannerItem)
	sendNormalResponse(w, "ok")
}
