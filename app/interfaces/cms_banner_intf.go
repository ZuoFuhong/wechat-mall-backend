package interfaces

import (
	"encoding/json"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"wechat-mall-backend/app/domain/entity"
	"wechat-mall-backend/app/domain/view"
	"wechat-mall-backend/consts"
	"wechat-mall-backend/errcode"
)

// GetBannerList 查询-Banner列表
func (m *MallHttpServiceImpl) GetBannerList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])
	bannerList, total, err := m.bannerService.GetBannerList(r.Context(), consts.ALL, page, size)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	voList := make([]*view.CMSBannerVO, 0)
	for _, v := range bannerList {
		vo := &view.CMSBannerVO{
			Id:      v.ID,
			Picture: v.Picture,
			Name:    v.Name,
			Status:  v.Status,
		}
		voList = append(voList, vo)
	}
	data := make(map[string]interface{}, 0)
	data["list"] = voList
	data["total"] = total
	Ok(w, data)
}

// GetBanner 查询-Banner详情（当前仅支持关联商品）
func (m *MallHttpServiceImpl) GetBanner(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	banner, err := m.bannerService.GetBannerById(r.Context(), id)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	if banner.ID == consts.ZERO || banner.Del == consts.DELETE {
		Error(w, errcode.NotFoundBanner, "banner 不存在")
		return
	}
	// 关联商品：可选
	goodsDO, err := m.goodsService.GetGoodsById(r.Context(), banner.BusinessID)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	categoryDO := new(entity.WechatMallCategoryDO)
	if goodsDO.ID != consts.ZERO {
		categoryDO, err = m.categoryService.GetCategoryById(r.Context(), goodsDO.CategoryID)
		if err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
		if categoryDO.ID == consts.ZERO || categoryDO.Del == consts.DELETE {
			Error(w, errcode.NotFoundCategory, "类目不存在")
			return
		}
	}
	bVO := view.CMSGoodsBannerVO{
		Id:            banner.ID,
		Picture:       banner.Picture,
		Name:          banner.Name,
		GoodsId:       goodsDO.ID,
		CategoryId:    categoryDO.ParentID,
		SubCategoryId: categoryDO.ID,
		Status:        banner.Status,
	}
	Ok(w, bVO)
}

// DoEditBanner 新增/编辑 Banner
func (m *MallHttpServiceImpl) DoEditBanner(w http.ResponseWriter, r *http.Request) {
	req := CMSBannerReq{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		Error(w, errcode.BadRequestParam, "参数解析失败")
		return
	}
	if err := validator.New().Struct(req); err != nil {
		Error(w, errcode.BadRequestParam, "缺少参数")
		return
	}
	if req.Id == consts.ZERO {
		banner := &entity.WechatMallBannerDO{
			Picture:      req.Picture,
			Name:         req.Name,
			BusinessType: req.BusinessType,
			BusinessID:   req.BusinessId,
			Status:       req.Status,
		}
		if err := m.bannerService.AddBanner(r.Context(), banner); err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
	} else {
		banner, err := m.bannerService.GetBannerById(r.Context(), req.Id)
		if err != nil {
			Error(w, errcode.BadRequestParam, "系统繁忙")
			return
		}
		if banner.ID == consts.ZERO || banner.Del == consts.DELETE {
			Error(w, errcode.NotFoundBanner, "banner不存在")
			return
		}
		banner.Picture = req.Picture
		banner.Name = req.Name
		banner.BusinessType = req.BusinessType
		banner.BusinessID = req.BusinessId
		banner.Status = req.Status
		if err := m.bannerService.UpdateBannerById(r.Context(), banner); err != nil {
			Error(w, errcode.BadRequestParam, "系统繁忙")
			return
		}
	}
	Ok(w, "ok")
}

// DoDeleteBanner 删除Banner
func (m *MallHttpServiceImpl) DoDeleteBanner(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	banner, err := m.bannerService.GetBannerById(r.Context(), id)
	if err != nil {
		Error(w, errcode.BadRequestParam, "系统繁忙")
		return
	}
	if banner.ID == consts.ZERO || banner.Del == consts.DELETE {
		Error(w, errcode.NotFoundBanner, "banner不存在")
		return
	}
	banner.Del = consts.DELETE
	if err := m.bannerService.UpdateBannerById(r.Context(), banner); err != nil {
		Error(w, errcode.BadRequestParam, "系统繁忙")
		return
	}
	Ok(w, "ok")
}
