package portal

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"wechat-mall-backend/defs"
)

// 查询Banner列表
func (h *Handler) GetBannerList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])

	bannerList, total := h.service.BannerService.GetBannerList(page, size)
	voList := []defs.PortalBannerVO{}
	for _, v := range *bannerList {
		bannerVO := defs.PortalBannerVO{}
		bannerVO.Id = v.Id
		bannerVO.Picture = v.Picture
		voList = append(voList, bannerVO)
	}
	resp := make(map[string]interface{})
	resp["list"] = voList
	resp["total"] = total
	defs.SendNormalResponse(w, resp)
}

// 查询宫格列表
func (h *Handler) GetGridCategoryList(w http.ResponseWriter, r *http.Request) {
	gridList, total := h.service.GridCategoryService.GetGridCategoryList(0, 0)

	gridVOList := []defs.PortalGridCategoryVO{}
	for _, v := range *gridList {
		gridVO := defs.PortalGridCategoryVO{}
		gridVO.Id = v.Id
		gridVO.Title = v.Title
		gridVO.Picture = v.Picture
		gridVO.CategoryId = v.CategoryId
		gridVOList = append(gridVOList, gridVO)
	}
	resp := make(map[string]interface{})
	resp["list"] = gridVOList
	resp["total"] = total
	defs.SendNormalResponse(w, resp)
}
