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

// 查询-分类列表
func (h *Handler) GetCategoryList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, _ := strconv.Atoi(vars["pid"])
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])

	cateVOList := []defs.CMSCategoryVO{}
	cateList, total := h.service.CategoryService.GetCategoryList(pid, page, size)
	for _, v := range *cateList {
		cateVO := defs.CMSCategoryVO{}
		cateVO.Id = v.Id
		cateVO.ParentId = v.ParentId
		cateVO.Name = v.Name
		cateVO.Sort = v.Sort
		cateVO.Online = v.Online
		cateVO.Picture = v.Picture
		cateVO.Description = v.Description
		cateVOList = append(cateVOList, cateVO)
	}

	resp := make(map[string]interface{})
	resp["list"] = cateVOList
	resp["total"] = total
	defs.SendNormalResponse(w, resp)
}

// 查询单个分类
func (h *Handler) GetCategoryById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	category := h.service.CategoryService.GetCategoryById(id)
	if category.Id == 0 {
		panic(errs.ErrorCategory)
	}
	cateVO := defs.CMSCategoryVO{}
	cateVO.Id = category.Id
	cateVO.ParentId = category.ParentId
	cateVO.Name = category.Name
	cateVO.Sort = category.Sort
	cateVO.Online = category.Online
	cateVO.Picture = category.Picture
	cateVO.Description = category.Description
	defs.SendNormalResponse(w, cateVO)
}

// 新增/编辑 分类
func (h *Handler) DoEditCategory(w http.ResponseWriter, r *http.Request) {
	req := defs.CMSCategoryReq{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		panic(errs.ErrorParameterValidate)
	}
	validate := validator.New()
	if err = validate.Struct(req); err != nil {
		panic(errs.NewParameterError(err.Error()))
	}
	if req.Id == 0 {
		category := h.service.CategoryService.GetCategoryByName(req.Name)
		if category.Id != 0 {
			panic(errs.NewCategoryError("分类名已存在！"))
		}
		category.ParentId = req.ParentId
		category.Name = req.Name
		category.Sort = req.Sort
		category.Online = req.Online
		category.Picture = req.Picture
		category.Description = req.Description
		h.service.CategoryService.AddCategory(category)
	} else {
		category := h.service.CategoryService.GetCategoryByName(req.Name)
		if category.Id != 0 && category.Id != req.Id {
			panic(errs.NewCategoryError("分类名已存在！"))
		}
		category = h.service.CategoryService.GetCategoryById(req.Id)
		if category.Id == 0 {
			panic(errs.ErrorCategory)
		}
		category.ParentId = req.ParentId
		category.Name = req.Name
		category.Sort = req.Sort
		category.Online = req.Online
		category.Picture = req.Picture
		category.Description = req.Description
		h.service.CategoryService.UpdateCategory(category)
	}
	defs.SendNormalResponse(w, "ok")
}

// 删除分类
func (h *Handler) DoDeleteCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	category := h.service.CategoryService.GetCategoryById(id)
	if category.Id == 0 {
		panic(errs.ErrorCategory)
	}
	if category.ParentId == 0 {
		_, total := h.service.CategoryService.GetCategoryList(id, 1, 1)
		if total > 0 {
			panic(errs.NewCategoryError("该分类下有子分类，不能删除！"))
		}
	} else {
		total := h.service.GoodsService.CountCategoryGoods(id)
		if total > 0 {
			panic(errs.NewCategoryError("该分类下有商品，不能删除！"))
		}
	}
	category.Del = 1
	h.service.CategoryService.UpdateCategory(category)
	defs.SendNormalResponse(w, "ok")
}

// CMS-查询-全部分类
func (h *Handler) GetChooseCategory(w http.ResponseWriter, r *http.Request) {
	tmpCategoryList := []map[string]interface{}{}
	categoryList, _ := h.service.CategoryService.GetCategoryList(0, 0, 0)
	for _, v := range *categoryList {
		tmpSubCategoryList := []map[string]interface{}{}
		subCategoryList, _ := h.service.CategoryService.GetCategoryList(v.Id, 0, 0)
		for _, sv := range *subCategoryList {
			tmpSubCategory := map[string]interface{}{}
			tmpSubCategory["value"] = sv.Id
			tmpSubCategory["label"] = sv.Name
			tmpSubCategoryList = append(tmpSubCategoryList, tmpSubCategory)
		}
		tmpCategory := map[string]interface{}{}
		tmpCategory["value"] = v.Id
		tmpCategory["label"] = v.Name
		tmpCategory["children"] = tmpSubCategoryList
		tmpCategoryList = append(tmpCategoryList, tmpCategory)
	}
	defs.SendNormalResponse(w, tmpCategoryList)
}
