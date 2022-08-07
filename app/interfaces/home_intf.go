package interfaces

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"wechat-mall-backend/app/domain/view"
	"wechat-mall-backend/consts"
	"wechat-mall-backend/errcode"
)

// HomeBanner 查询首页Banner列表
func (m *MallHttpServiceImpl) HomeBanner(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])

	bannerList, _, err := m.bannerService.GetBannerList(r.Context(), consts.ONLINE, page, size)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	voList := make([]*view.PortalBannerVO, 0)
	for _, v := range bannerList {
		bannerVO := &view.PortalBannerVO{}
		bannerVO.Id = v.ID
		bannerVO.Picture = v.Picture
		bannerVO.BusinessType = v.BusinessType
		bannerVO.BusinessId = v.BusinessID
		voList = append(voList, bannerVO)
	}
	Ok(w, voList)
}

// GetGridCategoryList 查询宫格列表
func (m *MallHttpServiceImpl) GetGridCategoryList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])
	gridList, _, err := m.gridService.GetGridCategoryList(r.Context(), page, size)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}

	gridVOList := make([]*view.PortalGridCategoryVO, 0)
	for _, v := range gridList {
		gridVO := &view.PortalGridCategoryVO{}
		gridVO.Id = v.ID
		gridVO.Name = v.Name
		gridVO.Picture = v.Picture
		gridVO.CategoryId = v.CategoryID
		gridVOList = append(gridVOList, gridVO)
	}
	Ok(w, gridList)
}

// GetSubCategoryList 查询-全部二级分类
func (m *MallHttpServiceImpl) GetSubCategoryList(w http.ResponseWriter, r *http.Request) {
	categoryList, err := m.couponService.GetAllSubCategory(r.Context())
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	Ok(w, categoryList)
}
