package interfaces

import (
	"encoding/json"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"wechat-mall-backend/app/domain/view"
	"wechat-mall-backend/consts"
	"wechat-mall-backend/errcode"
)

// GetCategoryList 查询-分类列表
func (m *MallHttpServiceImpl) GetCategoryList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, _ := strconv.Atoi(vars["pid"])
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])

	cateList, total, err := m.categoryService.GetCategoryList(r.Context(), pid, page, size)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	cateVOList := make([]*view.CMSCategoryVO, 0)
	for _, v := range cateList {
		cateVO := &view.CMSCategoryVO{
			Id:          v.ID,
			ParentId:    v.ParentID,
			Name:        v.Name,
			Sort:        v.Sort,
			Online:      v.Online,
			Picture:     v.Picture,
			Description: v.Description,
		}
		cateVOList = append(cateVOList, cateVO)
	}

	data := make(map[string]interface{})
	data["list"] = cateVOList
	data["total"] = total
	Ok(w, data)
}

// GetCategoryById 查询单个分类
func (m *MallHttpServiceImpl) GetCategoryById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	category, err := m.categoryService.GetCategoryById(r.Context(), id)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	if category.ID == consts.ZERO || category.Del == consts.DELETE {
		Error(w, errcode.NotFoundCategory, "分类不存在")
		return
	}
	cateVO := view.CMSCategoryVO{
		Id:          category.ID,
		ParentId:    category.ParentID,
		Name:        category.Name,
		Sort:        category.Sort,
		Online:      category.Online,
		Picture:     category.Picture,
		Description: category.Description,
	}
	Ok(w, cateVO)
}

// DoEditCategory 新增/编辑 分类
func (m *MallHttpServiceImpl) DoEditCategory(w http.ResponseWriter, r *http.Request) {
	req := new(CMSCategoryReq)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		Error(w, errcode.BadRequestParam, "参数解析失败")
		return
	}
	if err := validator.New().Struct(req); err != nil {
		Error(w, errcode.BadRequestParam, "缺少参数")
		return
	}
	if req.Id == consts.ZERO {
		category, err := m.categoryService.GetCategoryByName(r.Context(), req.Name)
		if err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
		if category.ID != consts.ZERO {
			Error(w, errcode.BadRequestParam, "分类名已存在")
			return
		}
		category.ParentID = req.ParentId
		category.Name = req.Name
		category.Sort = req.Sort
		category.Online = req.Online
		category.Picture = req.Picture
		category.Description = req.Description
		if err := m.categoryService.AddCategory(r.Context(), category); err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
	} else {
		category, err := m.categoryService.GetCategoryByName(r.Context(), req.Name)
		if err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
		if category.ID != consts.ZERO && category.ID != req.Id {
			Error(w, errcode.NotAllowOperation, "分类名已存在")
			return
		}
		category, err = m.categoryService.GetCategoryById(r.Context(), req.Id)
		if err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
		if category.ID == consts.ZERO || category.Del == consts.DELETE {
			Error(w, errcode.BadRequestParam, "分类不存在")
			return
		}
		category.ParentID = req.ParentId
		category.Name = req.Name
		category.Sort = req.Sort
		category.Online = req.Online
		category.Picture = req.Picture
		category.Description = req.Description
		if err := m.categoryService.UpdateCategory(r.Context(), category); err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
	}
	Ok(w, "ok")
}

// DoDeleteCategory 删除分类
func (m *MallHttpServiceImpl) DoDeleteCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	category, err := m.categoryService.GetCategoryById(r.Context(), id)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	if category.ID == consts.ZERO || category.Del == consts.DELETE {
		Error(w, errcode.BadRequestParam, "分类不存在")
		return
	}
	if category.ParentID == consts.ZERO {
		_, total, err := m.categoryService.GetCategoryList(r.Context(), id, 1, 1)
		if err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
		if total > 0 {
			Error(w, errcode.NotAllowOperation, "该分类下有子分类，不能删除")
			return
		}
	} else {
		total, err := m.goodsService.CountCategoryGoods(r.Context(), id)
		if err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
		if total > 0 {
			Error(w, errcode.NotAllowOperation, "该分类下有商品，不能删除")
			return
		}
		total, err = m.gridService.CountCategoryBindGrid(r.Context(), id)
		if err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
		if total > 0 {
			Error(w, errcode.NotAllowOperation, "该分类绑定了宫格，不能删除")
			return
		}
	}
	category.Del = consts.DELETE
	if err := m.categoryService.UpdateCategory(r.Context(), category); err != nil {
		if err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
	}
	Ok(w, "ok")
}

// GetChooseCategory CMS-查询-全部分类
func (m *MallHttpServiceImpl) GetChooseCategory(w http.ResponseWriter, r *http.Request) {
	var tmpCategoryList []map[string]interface{}
	categoryList, _, err := m.categoryService.GetCategoryList(r.Context(), 0, 1, 1000)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	for _, v := range categoryList {
		var tmpSubCategoryList []map[string]interface{}
		subCategoryList, _, err := m.categoryService.GetCategoryList(r.Context(), v.ID, 0, 1000)
		if err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
		for _, sv := range subCategoryList {
			tmpSubCategory := map[string]interface{}{}
			tmpSubCategory["value"] = sv.ID
			tmpSubCategory["label"] = sv.Name
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
