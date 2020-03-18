package cms

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"wechat-mall-backend/defs"
	"wechat-mall-backend/errs"
)

// 查询-规格列表
func (h *Handler) GetSpecificationList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])
	specList, total := h.service.SpecificationService.GetSpecificationList(page, size)

	specVOs := []defs.CMSSpecificationVO{}
	for _, v := range *specList {
		specVO := defs.CMSSpecificationVO{}
		specVO.Id = v.Id
		specVO.Name = v.Name
		specVO.Description = v.Description
		specVO.Unit = v.Unit
		specVO.Standard = v.Standard
		specVOs = append(specVOs, specVO)
	}
	resp := make(map[string]interface{})
	resp["list"] = specVOs
	resp["total"] = total
	defs.SendNormalResponse(w, resp)
}

// 查询-单个规格
func (h *Handler) GetSpecification(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	spec := h.service.SpecificationService.GetSpecificationById(id)
	if spec.Id == 0 {
		panic(errs.ErrorSpecification)
	}
	specVO := defs.CMSSpecificationVO{}
	specVO.Id = spec.Id
	specVO.Name = spec.Name
	specVO.Description = spec.Description
	specVO.Unit = spec.Unit
	specVO.Standard = spec.Standard
	defs.SendNormalResponse(w, specVO)
}

// 新增/编辑-规格
func (h *Handler) DoEditSpecification(w http.ResponseWriter, r *http.Request) {
	req := defs.CMSSpecificationReq{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		panic(errs.ErrorParameterValidate)
	}
	if req.Id == 0 {
		spec := h.service.SpecificationService.GetSpecificationByName(req.Name)
		if spec.Id != 0 {
			panic(errs.NewSpecificationError("规格名已存在！"))
		}
		spec.Name = req.Name
		spec.Description = req.Description
		spec.Unit = req.Unit
		spec.Standard = req.Standard
		h.service.SpecificationService.AddSpecification(spec)
	} else {
		spec := h.service.SpecificationService.GetSpecificationByName(req.Name)
		if spec.Id != 0 && spec.Id != req.Id {
			panic(errs.NewSpecificationError("规格名已存在！"))
		}
		spec = h.service.SpecificationService.GetSpecificationById(req.Id)
		if spec.Id == 0 {
			panic(errs.ErrorSpecification)
		}
		spec = h.service.SpecificationService.GetSpecificationById(req.Id)
		spec.Name = req.Name
		spec.Description = req.Description
		spec.Unit = req.Unit
		spec.Standard = req.Standard
		h.service.SpecificationService.UpdateSpecificationById(spec)
	}
	defs.SendNormalResponse(w, "ok")
}

// 删除-规格
func (h *Handler) DoDeleteSpecification(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	spec := h.service.SpecificationService.GetSpecificationById(id)
	if spec.Id == 0 {
		panic(errs.ErrorSpecification)
	}
	attrList := h.service.SpecificationService.GetSpecificationAttrList(id)
	if len(*attrList) > 0 {
		panic(errs.NewSpecificationError("该规格下有属性，不能删除！"))
	}
	spec.Del = 1
	h.service.SpecificationService.UpdateSpecificationById(spec)
	defs.SendNormalResponse(w, "ok")
}

// 查询-单个规格-全部属性
func (h *Handler) GetSpecificationAttrList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	specId, _ := strconv.Atoi(vars["specId"])
	specAttrList := h.service.SpecificationService.GetSpecificationAttrList(specId)

	attrVOs := []defs.CMSSpecificationAttrVO{}
	for _, v := range *specAttrList {
		attrVO := defs.CMSSpecificationAttrVO{}
		attrVO.Id = v.Id
		attrVO.SpecId = v.SpecId
		attrVO.Value = v.Value
		attrVO.Extend = v.Extend
		attrVOs = append(attrVOs, attrVO)
	}
	defs.SendNormalResponse(w, attrVOs)
}

// 查询-单个规格-单个属性
func (h *Handler) GetSpecificationAttr(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	spec := h.service.SpecificationService.GetSpecificationAttrById(id)
	if spec.Id == 0 {
		panic(errs.ErrorSpecificationAttr)
	}
	attrVO := defs.CMSSpecificationAttrVO{}
	attrVO.Id = spec.Id
	attrVO.SpecId = spec.SpecId
	attrVO.Value = spec.Value
	attrVO.Extend = spec.Extend
	defs.SendNormalResponse(w, attrVO)
}

// 新增/更新-规格-单个属性
func (h *Handler) DoEditSpecificationAttr(w http.ResponseWriter, r *http.Request) {
	req := defs.CMSSpecificationAttrReq{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		panic(errs.ErrorParameterValidate)
	}
	spec := h.service.SpecificationService.GetSpecificationById(req.SpecId)
	if spec.Id == 0 {
		panic(errs.ErrorSpecificationAttr)
	}
	if req.Id == 0 {
		spec := h.service.SpecificationService.GetSpecificationAttrByValue(req.Value)
		if spec.Id != 0 {
			panic(errs.NewSpecificationAttr("属性名已存在！"))
		}
		spec.SpecId = req.SpecId
		spec.Value = req.Value
		spec.Extend = req.Extend
		h.service.SpecificationService.AddSpecificationAttr(spec)
	} else {
		spec := h.service.SpecificationService.GetSpecificationAttrByValue(req.Value)
		if spec.Id != 0 && spec.Id != req.Id {
			panic(errs.NewSpecificationError("属性名已存在！"))
		}
		spec = h.service.SpecificationService.GetSpecificationAttrById(req.Id)
		if spec.Id == 0 {
			panic(errs.ErrorSpecificationAttr)
		}
		spec = h.service.SpecificationService.GetSpecificationAttrById(req.Id)
		spec.SpecId = req.SpecId
		spec.Value = req.Value
		spec.Extend = req.Extend
		h.service.SpecificationService.UpdateSpecificationAttrById(spec)
	}
	defs.SendNormalResponse(w, "ok")
}

// 删除-单个规格-单个属性
func (h *Handler) DoDeleteSpecificationAttr(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	spec := h.service.SpecificationService.GetSpecificationAttrById(id)
	if spec.Id == 0 {
		panic(errs.ErrorSpecificationAttr)
	}
	spec.Del = 1
	h.service.SpecificationService.UpdateSpecificationAttrById(spec)
	defs.SendNormalResponse(w, "ok")
}
