package interfaces

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"wechat-mall-backend/app/domain/entity"
	"wechat-mall-backend/app/domain/view"
	"wechat-mall-backend/consts"
	"wechat-mall-backend/errcode"
)

// GetCmsGoodsList 查询-商品列表
func (m *MallHttpServiceImpl) GetCmsGoodsList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	keyword, _ := vars["k"]
	categoryId, _ := strconv.Atoi(vars["c"])
	online, _ := strconv.Atoi(vars["o"])
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])

	if categoryId == 0 {
		categoryId = consts.ALL
	}
	if online == -1 {
		online = consts.ALL
	}
	goodsList, total, err := m.goodsService.GetGoodsList(r.Context(), keyword, categoryId, online, page, size)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	goodsVOList := make([]*view.CMSGoodsListVO, 0)
	for _, v := range goodsList {
		categoryDO, err := m.categoryService.GetCategoryById(r.Context(), v.CategoryID)
		if err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
		if categoryDO.ID == consts.ZERO || categoryDO.Del == consts.DELETE {
			Error(w, errcode.NotFoundCategory, "分类不存在")
			return
		}
		goodsVO := &view.CMSGoodsListVO{}
		goodsVO.Id = v.ID
		goodsVO.BrandName = v.BrandName
		goodsVO.Title = v.Title
		goodsVO.Price = v.Price
		goodsVO.CategoryId = v.CategoryID
		goodsVO.CategoryName = categoryDO.Name
		goodsVO.Online = v.Online
		goodsVO.Picture = v.Picture
		goodsVOList = append(goodsVOList, goodsVO)
	}
	data := make(map[string]interface{})
	data["list"] = goodsVOList
	data["total"] = total
	Ok(w, data)
}

// GetGoods 查询商品详情
func (m *MallHttpServiceImpl) GetGoods(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	goodsDO, err := m.goodsService.GetGoodsById(r.Context(), id)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	if goodsDO.ID == consts.ZERO || goodsDO.Del == consts.DELETE {
		Error(w, errcode.NotFoundGoods, "商品不存在")
		return
	}
	categoryDO, err := m.categoryService.GetCategoryById(r.Context(), goodsDO.CategoryID)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	if categoryDO.ID == consts.ZERO || categoryDO.Del == consts.DELETE {
		Error(w, errcode.NotFoundCategory, "分类不存在")
		return
	}
	goodsVO := &view.CMSGoodsVO{}
	goodsVO.Id = goodsDO.ID
	goodsVO.BrandName = goodsDO.BrandName
	goodsVO.Title = goodsDO.Title
	goodsVO.Price = goodsDO.Price
	goodsVO.DiscountPrice = goodsDO.DiscountPrice
	goodsVO.CategoryId = categoryDO.ParentID
	goodsVO.SubCategoryId = goodsDO.CategoryID
	goodsVO.CategoryName = categoryDO.Name
	goodsVO.Online = goodsDO.Online
	goodsVO.Picture = goodsDO.Picture
	goodsVO.BannerPicture = goodsDO.BannerPicture
	goodsVO.DetailPicture = goodsDO.DetailPicture
	goodsVO.Tags = goodsDO.Tags
	Ok(w, goodsDO)
}

// DoEditGoods 新增/编辑商品
func (m *MallHttpServiceImpl) DoEditGoods(w http.ResponseWriter, r *http.Request) {
	req := new(CMSGoodsReq)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		Error(w, errcode.BadRequestParam, "参数解析失败")
		return
	}
	category, err := m.categoryService.GetCategoryById(r.Context(), req.CategoryId)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	if category.ID == consts.ZERO || category.Del == consts.DELETE {
		Error(w, errcode.NotFoundCategory, "分类不存在")
		return
	}
	if req.Id == consts.ZERO {
		goodsDO := &entity.WechatMallGoodsDO{}
		goodsDO.BrandName = req.BrandName
		goodsDO.Title = req.Title
		goodsDO.Price = req.Price
		goodsDO.DiscountPrice = req.DiscountPrice
		goodsDO.CategoryID = req.CategoryId
		goodsDO.Online = req.Online
		goodsDO.Picture = req.Picture
		goodsDO.BannerPicture = req.BannerPicture
		goodsDO.DetailPicture = req.DetailPicture
		goodsDO.Tags = req.Tags
		goodsId, err := m.goodsService.AddGoods(r.Context(), goodsDO)
		if err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
		if err := m.goodsService.AddGoodsSpec(r.Context(), goodsId, req.SpecList); err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
	} else {
		goodsDO, err := m.goodsService.GetGoodsById(r.Context(), req.Id)
		if err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
		if goodsDO.ID == consts.ZERO || goodsDO.Del == consts.DELETE {
			Error(w, errcode.ErrorInternalFaults, "商品不存在")
			return
		}
		goodsDO.BrandName = req.BrandName
		goodsDO.Title = req.Title
		goodsDO.Price = req.Price
		goodsDO.DiscountPrice = req.DiscountPrice
		goodsDO.CategoryID = req.CategoryId
		goodsDO.Online = req.Online
		goodsDO.Picture = req.Picture
		goodsDO.BannerPicture = req.BannerPicture
		goodsDO.DetailPicture = req.DetailPicture
		goodsDO.Tags = req.Tags
		if err := m.goodsService.UpdateGoodsById(r.Context(), goodsDO); err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
		if err := m.goodsService.AddGoodsSpec(r.Context(), req.Id, req.SpecList); err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
	}
	Ok(w, "ok")
}

// DoDeleteGoods 删除商品
func (m *MallHttpServiceImpl) DoDeleteGoods(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	goodsDO, err := m.goodsService.GetGoodsById(r.Context(), id)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	if goodsDO.ID == consts.ZERO || goodsDO.Del == consts.DELETE {
		Error(w, errcode.ErrorInternalFaults, "商品不存在")
		return
	}
	_, total, err := m.skuService.GetSKUList(r.Context(), "", id, consts.ALL, 1, 1)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	if total > 0 {
		Error(w, errcode.NotAllowOperation, "该商品下有SKU，不能删除")
		return
	}
	goodsDO.Del = consts.DELETE
	if err := m.goodsService.UpdateGoodsById(r.Context(), goodsDO); err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	// 同步：清理关联的规格
	if err := m.goodsService.AddGoodsSpec(r.Context(), id, []int{}); err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	Ok(w, "ok")
}

// GetGoodsSpecList 查询-商品-规格属性
func (m *MallHttpServiceImpl) GetGoodsSpecList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	goodsId, _ := strconv.Atoi(vars["id"])
	specList, err := m.goodsService.GetGoodsSpecList(r.Context(), goodsId)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	Ok(w, specList)
}

// GetChooseCategoryGoods Sku查询-全部分类及商品
func (m *MallHttpServiceImpl) GetChooseCategoryGoods(w http.ResponseWriter, r *http.Request) {
	var tmpCategoryList []map[string]interface{}
	categoryList, _, err := m.categoryService.GetCategoryList(r.Context(), 0, 0, 0)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	for _, v := range categoryList {
		var tmpSubCategoryList []map[string]interface{}
		subCategoryList, _, err := m.categoryService.GetCategoryList(r.Context(), v.ID, 0, 0)
		if err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
		for _, sv := range subCategoryList {
			var tmpGoodsList []map[string]interface{}
			goodsList, _, err := m.goodsService.GetGoodsList(r.Context(), "", sv.ID, consts.ALL, 0, 0)
			if err != nil {
				Error(w, errcode.ErrorInternalFaults, "系统繁忙")
				return
			}
			for _, g := range goodsList {
				goodsItem := map[string]interface{}{}
				goodsItem["value"] = g.ID
				goodsItem["label"] = g.Title
				tmpGoodsList = append(tmpGoodsList, goodsItem)
			}
			tmpSubCategory := map[string]interface{}{}
			tmpSubCategory["value"] = sv.ID
			tmpSubCategory["label"] = sv.Name
			tmpSubCategory["children"] = tmpGoodsList
			tmpSubCategoryList = append(tmpSubCategoryList, tmpSubCategory)
		}
		tmpCategory := map[string]interface{}{}
		tmpCategory["value"] = v.ID
		tmpCategory["label"] = v.Name
		tmpCategory["children"] = tmpSubCategoryList
		tmpCategoryList = append(tmpCategoryList, tmpCategory)
	}
	Ok(w, tmpCategoryList)
}
