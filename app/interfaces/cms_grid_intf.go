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

// GetCmsGridCategoryList 查询-宫格列表
func (m *MallHttpServiceImpl) GetCmsGridCategoryList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])

	gcArr, total, err := m.gridService.GetGridCategoryList(r.Context(), page, size)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	gcVOList := make([]*view.CMSGridCategoryVO, 0)
	for _, v := range gcArr {
		categoryDO, err := m.categoryService.GetCategoryById(r.Context(), v.CategoryID)
		if err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
		if categoryDO.ID == consts.ZERO || categoryDO.Del == consts.DELETE {
			Error(w, errcode.NotFoundCategory, "分类不存在")
			return
		}
		gcVO := &view.CMSGridCategoryVO{
			Id:           v.ID,
			Name:         v.Name,
			CategoryId:   v.CategoryID,
			CategoryName: categoryDO.Name,
			Picture:      v.Picture,
		}
		gcVOList = append(gcVOList, gcVO)
	}
	data := make(map[string]interface{})
	data["list"] = gcVOList
	data["total"] = total
	Ok(w, data)
}

// GetGridCategory 查询-单个宫格
func (m *MallHttpServiceImpl) GetGridCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	gridC, err := m.gridService.GetGridCategoryById(r.Context(), id)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	if gridC.ID == consts.ZERO || gridC.Del == consts.DELETE {
		Error(w, errcode.NotFoundGridCategory, "宫格不存在")
		return
	}
	categoryDO, err := m.categoryService.GetCategoryById(r.Context(), gridC.CategoryID)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	if categoryDO.ID == consts.ZERO || categoryDO.Del == consts.DELETE {
		Error(w, errcode.NotFoundCategory, "分类不存在")
		return
	}
	gcVO := &view.CMSGridCategoryDetailVO{}
	gcVO.Id = gridC.ID
	gcVO.Name = gridC.Name
	gcVO.CategoryId = categoryDO.ParentID
	gcVO.SubCategoryId = categoryDO.ID
	gcVO.SubCategoryName = categoryDO.Name
	gcVO.Picture = gridC.Picture
	Ok(w, gcVO)
}

// DoEditGridCategory 新增/更新宫格
func (m *MallHttpServiceImpl) DoEditGridCategory(w http.ResponseWriter, r *http.Request) {
	req := new(CMSGridCategoryReq)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		Error(w, errcode.BadRequestParam, "参数解析失败")
		return
	}
	if err := validator.New().Struct(req); err != nil {
		Error(w, errcode.BadRequestParam, "缺少参数")
		return
	}
	if req.Id == consts.ZERO {
		gridC, err := m.gridService.GetGridCategoryByName(r.Context(), req.Name)
		if err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
		if gridC.ID != consts.ZERO {
			Error(w, errcode.NotAllowOperation, "宫格名称已存在")
			return
		}
		gridC.ID = req.Id
		gridC.Name = req.Name
		gridC.CategoryID = req.CategoryId
		gridC.Picture = req.Picture
		if err := m.gridService.AddGridCategory(r.Context(), gridC); err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
	} else {
		gridC, err := m.gridService.GetGridCategoryByName(r.Context(), req.Name)
		if err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
		if gridC.ID != consts.ZERO && gridC.ID != req.Id {
			Error(w, errcode.NotAllowOperation, "宫格名称已存在")
			return
		}
		gridC, err = m.gridService.GetGridCategoryById(r.Context(), req.Id)
		if err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
		if gridC.ID == consts.ZERO || gridC.Del == consts.DELETE {
			Error(w, errcode.NotFoundGridCategory, "宫格不存在")
			return
		}
		gridC.Name = req.Name
		gridC.CategoryID = req.CategoryId
		gridC.Picture = req.Picture
		if err := m.gridService.UpdateGridCategory(r.Context(), gridC); err != nil {
			if err != nil {
				Error(w, errcode.ErrorInternalFaults, "系统繁忙")
				return
			}
		}
	}
	Ok(w, "ok")
}

// DoDeleteGridCategory 删除-单个宫格
func (m *MallHttpServiceImpl) DoDeleteGridCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	gridC, err := m.gridService.GetGridCategoryById(r.Context(), id)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	if gridC.ID == consts.ZERO || gridC.Del == consts.DELETE {
		Error(w, errcode.NotFoundGridCategory, "宫格不存在")
		return
	}
	gridC.Del = consts.DELETE
	if err := m.gridService.UpdateGridCategory(r.Context(), gridC); err != nil {
		if err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
	}
	Ok(w, "ok")
}
