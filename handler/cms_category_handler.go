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
)

func (h *CMSHandler) GetCategoryList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])

	var cateVOList []defs.CategoryVO
	cateList, total := h.service.CategoryService.GetCategoryList(page, size)
	for _, v := range *cateList {
		cateVO := defs.CategoryVO{}
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
	sendNormalResponse(w, resp)
}

func (h *CMSHandler) GetCategoryById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	category := h.service.CategoryService.GetCategoryById(id)
	if category.Id == 0 {
		panic(errs.ErrorCategory)
	}
	cateVO := defs.CategoryVO{}
	cateVO.Id = category.Id
	cateVO.ParentId = category.ParentId
	cateVO.Name = category.Name
	cateVO.Sort = category.Sort
	cateVO.Online = category.Online
	cateVO.Picture = category.Picture
	cateVO.Description = category.Description

	sendNormalResponse(w, cateVO)
}

func (h *CMSHandler) DoEditCategory(w http.ResponseWriter, r *http.Request) {
	req := defs.CategoryReq{}
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
			panic(errs.NewCategoryError("The category name alreadly exists"))
		}
		category.ParentId = req.ParentId
		category.Name = req.Name
		category.Sort = req.Sort
		category.Online = req.Online
		category.Picture = req.Picture
		category.Description = req.Description
		h.service.CategoryService.AddCategory((*model.Category)(category))
	} else {
		category := h.service.CategoryService.GetCategoryByName(req.Name)
		if category.Id != 0 && category.Id != req.Id {
			panic(errs.NewCategoryError("The category name alreadly exists"))
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
		h.service.CategoryService.UpdateCategory((*model.Category)(category))
	}
	sendNormalResponse(w, "ok")
}

func (h *CMSHandler) DoDeleteCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	category := h.service.CategoryService.GetCategoryById(id)
	if category.Id == 0 {
		panic(errs.ErrorCategory)
	}
	category.Del = 1
	h.service.CategoryService.UpdateCategory((*model.Category)(category))
	sendNormalResponse(w, "ok")
}
