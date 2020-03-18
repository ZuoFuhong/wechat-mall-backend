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

// 查询-Banner列表
func (h *Handler) GetBannerList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])
	bannerList, total := h.service.BannerService.GetBannerList(page, size)

	voList := []defs.CMSBannerVO{}
	for _, v := range *bannerList {
		vo := defs.CMSBannerVO{}
		vo.Id = v.Id
		vo.Picture = v.Picture
		vo.Title = v.Title
		voList = append(voList, vo)
	}
	resp := make(map[string]interface{}, 0)
	resp["list"] = voList
	resp["total"] = total
	defs.SendNormalResponse(w, resp)
}

// 查询-Banner详情
func (h *Handler) GetBanner(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	banner := h.service.BannerService.GetBannerById(id)
	if banner.Id == 0 {
		panic(errs.ErrorBannerNotExist)
	}
	bVO := defs.CMSBannerVO{}
	bVO.Id = banner.Id
	bVO.Picture = banner.Picture
	bVO.Title = banner.Title
	defs.SendNormalResponse(w, bVO)
}

// 新增/编辑 Banner
func (h *Handler) DoEditBanner(w http.ResponseWriter, r *http.Request) {
	req := defs.CMSBannerReq{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		panic(err)
	}
	validate := validator.New()
	if err = validate.Struct(req); err != nil {
		panic(errs.NewParameterError(err.Error()))
	}
	if req.Id == 0 {
		banner := model.WechatMallBannerDO{}
		banner.Picture = req.Picture
		banner.Title = req.Title
		h.service.BannerService.AddBanner(&banner)
	} else {
		banner := h.service.BannerService.GetBannerById(req.Id)
		if banner.Id == 0 {
			panic(errs.ErrorBannerNotExist)
		}
		banner.Picture = req.Picture
		banner.Title = req.Title
		h.service.BannerService.UpdateBannerById(banner)
	}
	defs.SendNormalResponse(w, "ok")
}

// 删除Banner
func (h *Handler) DoDeleteBanner(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	banner := h.service.BannerService.GetBannerById(id)
	if banner.Id == 0 {
		panic(errs.ErrorBannerNotExist)
	}
	banner.Del = 1
	h.service.BannerService.UpdateBannerById(banner)
	defs.SendNormalResponse(w, "ok")
}
