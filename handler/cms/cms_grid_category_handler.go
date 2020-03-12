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

func (h *Handler) GetGridCategoryList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])

	gcArr, total := h.service.GridCategoryService.GetGridCategoryList(page, size)
	gcVOList := []defs.CMSGridCategoryVO{}
	for _, v := range *gcArr {
		gcVO := defs.CMSGridCategoryVO{}
		gcVO.Id = v.Id
		gcVO.Title = v.Title
		gcVO.Name = v.Name
		gcVO.CategoryId = v.CategoryId
		gcVO.Picture = v.Picture
		gcVOList = append(gcVOList, gcVO)
	}

	resp := make(map[string]interface{})
	resp["list"] = gcVOList
	resp["total"] = total
	defs.SendNormalResponse(w, resp)
}

func (h *Handler) GetGridCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	gridC := h.service.GridCategoryService.GetGridCategoryById(id)
	if gridC.Id == 0 {
		panic(errs.ErrorGridCategory)
	}
	gcVO := defs.CMSGridCategoryVO{}
	gcVO.Id = gridC.Id
	gcVO.Title = gridC.Title
	gcVO.Name = gridC.Name
	gcVO.CategoryId = gridC.CategoryId
	gcVO.Picture = gridC.Picture
	defs.SendNormalResponse(w, gcVO)
}

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
	if req.Id == 0 {
		gridC := h.service.GridCategoryService.GetGridCategoryByName(req.Name)
		if gridC.Id != 0 {
			panic(errs.NewGridCategoryError("The name already exists"))
		}
		gridC.Id = req.Id
		gridC.Title = req.Title
		gridC.Name = req.Name
		gridC.CategoryId = req.CategoryId
		gridC.Picture = req.Picture
		h.service.GridCategoryService.AddGridCategory(gridC)
	} else {
		gridC := h.service.GridCategoryService.GetGridCategoryByName(req.Name)
		if gridC.Id != 0 && gridC.Id != req.Id {
			panic(errs.NewGridCategoryError("The name already exists"))
		}
		gridC = h.service.GridCategoryService.GetGridCategoryById(req.Id)
		if gridC.Id == 0 {
			panic(errs.ErrorGridCategory)
		}
		gridC.Title = req.Title
		gridC.Name = req.Name
		gridC.CategoryId = req.CategoryId
		gridC.Picture = req.Picture
		h.service.GridCategoryService.UpdateGridCategory(gridC)
	}
	defs.SendNormalResponse(w, "ok")
}

func (h *Handler) DoDeleteGridCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	gridC := h.service.GridCategoryService.GetGridCategoryById(id)
	if gridC.Id == 0 {
		panic(errs.ErrorGridCategory)
	}
	gridC.Del = 1
	h.service.GridCategoryService.UpdateGridCategory(gridC)
	defs.SendNormalResponse(w, "ok")
}
