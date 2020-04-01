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

// 查询-宫格列表
func (h *Handler) GetGridCategoryList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])

	gcArr, total := h.service.GridCategoryService.GetGridCategoryList(page, size)
	gcVOList := []defs.CMSGridCategoryVO{}
	for _, v := range *gcArr {
		categoryDO := h.service.CategoryService.GetCategoryById(v.CategoryId)
		if categoryDO.Id == defs.ZERO || categoryDO.Del == defs.DELETE {
			panic(errs.ErrorCategory)
		}
		gcVO := defs.CMSGridCategoryVO{}
		gcVO.Id = v.Id
		gcVO.Name = v.Name
		gcVO.CategoryId = v.CategoryId
		gcVO.CategoryName = categoryDO.Name
		gcVO.Picture = v.Picture
		gcVOList = append(gcVOList, gcVO)
	}

	resp := make(map[string]interface{})
	resp["list"] = gcVOList
	resp["total"] = total
	defs.SendNormalResponse(w, resp)
}

// 查询-单个宫格
func (h *Handler) GetGridCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	gridC := h.service.GridCategoryService.GetGridCategoryById(id)
	if gridC.Id == defs.ZERO || gridC.Del == defs.DELETE {
		panic(errs.ErrorGridCategory)
	}
	categoryDO := h.service.CategoryService.GetCategoryById(gridC.CategoryId)
	if categoryDO.Id == defs.ZERO || categoryDO.Del == defs.DELETE {
		panic(errs.ErrorCategory)
	}
	gcVO := defs.CMSGridCategoryDetailVO{}
	gcVO.Id = gridC.Id
	gcVO.Name = gridC.Name
	gcVO.CategoryId = categoryDO.ParentId
	gcVO.SubCategoryId = categoryDO.Id
	gcVO.SubCategoryName = categoryDO.Name
	gcVO.Picture = gridC.Picture
	defs.SendNormalResponse(w, gcVO)
}

// 新增/更新宫格
func (h *Handler) DoEditGridCategory(w http.ResponseWriter, r *http.Request) {
	req := defs.CMSGridCategoryReq{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		panic(errs.ErrorParameterValidate)
	}
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		panic(errs.NewParameterError(err.Error()))
	}
	if req.Id == defs.ZERO {
		gridC := h.service.GridCategoryService.GetGridCategoryByName(req.Name)
		if gridC.Id != defs.ZERO {
			panic(errs.NewGridCategoryError("宫格名称已存在！"))
		}
		gridC.Id = req.Id
		gridC.Name = req.Name
		gridC.CategoryId = req.CategoryId
		gridC.Picture = req.Picture
		h.service.GridCategoryService.AddGridCategory(gridC)
	} else {
		gridC := h.service.GridCategoryService.GetGridCategoryByName(req.Name)
		if gridC.Id != defs.ZERO && gridC.Id != req.Id {
			panic(errs.NewGridCategoryError("宫格名称已存在！"))
		}
		gridC = h.service.GridCategoryService.GetGridCategoryById(req.Id)
		if gridC.Id == defs.ZERO || gridC.Del == defs.DELETE {
			panic(errs.ErrorGridCategory)
		}
		gridC.Name = req.Name
		gridC.CategoryId = req.CategoryId
		gridC.Picture = req.Picture
		h.service.GridCategoryService.UpdateGridCategory(gridC)
	}
	defs.SendNormalResponse(w, "ok")
}

// 删除-单个宫格
func (h *Handler) DoDeleteGridCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	gridC := h.service.GridCategoryService.GetGridCategoryById(id)
	if gridC.Id == defs.ZERO || gridC.Del == defs.DELETE {
		panic(errs.ErrorGridCategory)
	}
	gridC.Del = defs.DELETE
	h.service.GridCategoryService.UpdateGridCategory(gridC)
	defs.SendNormalResponse(w, "ok")
}
