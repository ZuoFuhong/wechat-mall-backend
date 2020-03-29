package portal

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"wechat-mall-backend/defs"
)

// 查询首页Banner列表
func (h *Handler) HomeBanner(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])

	bannerList, _ := h.service.BannerService.GetBannerList(1, page, size)
	voList := []defs.PortalBannerVO{}
	for _, v := range *bannerList {
		bannerVO := defs.PortalBannerVO{}
		bannerVO.Id = v.Id
		bannerVO.Picture = v.Picture
		bannerVO.BusinessType = v.BusinessType
		bannerVO.BusinessId = v.BusinessId
		voList = append(voList, bannerVO)
	}
	defs.SendNormalResponse(w, voList)
}

// 查询宫格列表
func (h *Handler) GetGridCategoryList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])
	gridList, _ := h.service.GridCategoryService.GetGridCategoryList(page, size)

	gridVOList := []defs.PortalGridCategoryVO{}
	for _, v := range *gridList {
		gridVO := defs.PortalGridCategoryVO{}
		gridVO.Id = v.Id
		gridVO.Name = v.Name
		gridVO.Picture = v.Picture
		gridVO.CategoryId = v.CategoryId
		gridVOList = append(gridVOList, gridVO)
	}
	defs.SendNormalResponse(w, gridVOList)
}
