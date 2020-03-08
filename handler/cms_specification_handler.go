package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"wechat-mall-backend/defs"
	"wechat-mall-backend/errs"
)

func (h *CMSHandler) GetSpecificationList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])
	specList, total := h.service.SpecificationService.GetSpecificationList(page, size)

	var specVOs []defs.SpecificationVO
	for _, v := range *specList {
		specVO := defs.SpecificationVO{}
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
	sendNormalResponse(w, resp)
}

func (h *CMSHandler) GetSpecification(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	spec := h.service.SpecificationService.GetSpecificationById(id)
	if spec.Id == 0 {
		panic(errs.ErrorSpecification)
	}
	specVO := defs.SpecificationVO{}
	specVO.Id = spec.Id
	specVO.Name = spec.Name
	specVO.Description = spec.Description
	specVO.Unit = spec.Unit
	specVO.Standard = spec.Standard

	sendNormalResponse(w, specVO)
}

func (h *CMSHandler) DoEditSpecification(w http.ResponseWriter, r *http.Request) {
	req := defs.SpecificationReq{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		panic(errs.ErrorParameterValidate)
	}
	if req.Id == 0 {
		spec := h.service.SpecificationService.GetSpecificationByName(req.Name)
		if spec.Id != 0 {
			panic(errs.NewSpecificationError("The name already exists"))
		}
		spec.Name = req.Name
		spec.Description = req.Description
		spec.Unit = req.Unit
		spec.Standard = req.Standard
		h.service.SpecificationService.AddSpecification(spec)
	} else {
		spec := h.service.SpecificationService.GetSpecificationByName(req.Name)
		if spec.Id != 0 && spec.Id != req.Id {
			panic(errs.NewSpecificationError("The name already exists"))
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
	sendNormalResponse(w, "ok")
}

func (h *CMSHandler) DoDeleteSpecification(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	spec := h.service.SpecificationService.GetSpecificationById(id)
	if spec.Id == 0 {
		panic(errs.ErrorSpecification)
	}
	spec.Del = 1
	h.service.SpecificationService.UpdateSpecificationById(spec)
	sendNormalResponse(w, "ok")
}

func (h *CMSHandler) GetSpecificationAttrList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	specId, _ := strconv.Atoi(vars["specId"])
	specAttrList := h.service.SpecificationService.GetSpecificationAttrList(specId)

	attrVOs := []defs.SpecificationAttrVO{}
	for _, v := range *specAttrList {
		attrVO := defs.SpecificationAttrVO{}
		attrVO.Id = v.Id
		attrVO.SpecId = v.SpecId
		attrVO.Value = v.Value
		attrVO.Extend = v.Extend
		attrVOs = append(attrVOs, attrVO)
	}
	sendNormalResponse(w, attrVOs)
}

func (h *CMSHandler) GetSpecificationAttr(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	spec := h.service.SpecificationService.GetSpecificationAttrById(id)
	if spec.Id == 0 {
		panic(errs.ErrorSpecificationAttr)
	}
	attrVO := defs.SpecificationAttrVO{}
	attrVO.Id = spec.Id
	attrVO.SpecId = spec.SpecId
	attrVO.Value = spec.Value
	attrVO.Extend = spec.Extend
	sendNormalResponse(w, attrVO)
}

func (h *CMSHandler) DoEditSpecificationAttr(w http.ResponseWriter, r *http.Request) {
	req := defs.SpecificationAttrReq{}
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
			panic(errs.NewSpecificationAttr("The value already exists"))
		}
		spec.SpecId = req.SpecId
		spec.Value = req.Value
		spec.Extend = req.Extend
		h.service.SpecificationService.AddSpecificationAttr(spec)
	} else {
		spec := h.service.SpecificationService.GetSpecificationAttrByValue(req.Value)
		if spec.Id != 0 && spec.Id != req.Id {
			panic(errs.NewSpecificationError("The name already exists"))
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
	sendNormalResponse(w, "ok")
}

func (h *CMSHandler) DoDeleteSpecificationAttr(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	spec := h.service.SpecificationService.GetSpecificationAttrById(id)
	if spec.Id == 0 {
		panic(errs.ErrorSpecificationAttr)
	}
	spec.Del = 1
	h.service.SpecificationService.UpdateSpecificationAttrById(spec)
	sendNormalResponse(w, "ok")
}
