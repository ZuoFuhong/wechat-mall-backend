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

// GetSKUList 查询-SKU列表
func (m *MallHttpServiceImpl) GetSKUList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	goodsId, _ := strconv.Atoi(vars["goodsId"])
	keyword := vars["k"]
	online, _ := strconv.Atoi(vars["o"])
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])

	skuVOs := make([]*view.CMSSkuListVO, 0)
	skuList, total, err := m.skuService.GetSKUList(r.Context(), keyword, goodsId, online, page, size)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	for _, v := range skuList {
		skuVO := &view.CMSSkuListVO{}
		skuVO.Id = v.ID
		skuVO.Title = v.Title
		skuVO.Price = v.Price
		skuVO.Code = v.Code
		skuVO.Stock = v.Stock
		skuVO.GoodsId = v.GoodsID
		skuVO.Online = v.Online
		skuVO.Picture = v.Picture
		skuVO.Specs = v.Specs
		skuVOs = append(skuVOs, skuVO)
	}
	data := make(map[string]interface{})
	data["list"] = skuVOs
	data["total"] = total
	Ok(w, data)
}

// GetSKU 查询-单个SKU
func (m *MallHttpServiceImpl) GetSKU(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	sku, err := m.skuService.GetSKUById(r.Context(), id)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	if sku.ID == consts.ZERO || sku.Del == consts.DELETE {
		Error(w, errcode.NotFoundGoodsSku, "商品SKU不存在")
		return
	}
	goodsDO, err := m.goodsService.GetGoodsById(r.Context(), sku.GoodsID)
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
		Error(w, errcode.NotFoundGoods, "类目不存在")
		return
	}
	skuVO := &view.CMSSkuDetailVO{}
	skuVO.Id = sku.ID
	skuVO.Title = sku.Title
	skuVO.Price = sku.Price
	skuVO.Code = sku.Code
	skuVO.Stock = sku.Stock
	skuVO.CategoryId = categoryDO.ParentID
	skuVO.SubCategoryId = categoryDO.ID
	skuVO.GoodsId = sku.GoodsID
	skuVO.Online = sku.Online
	skuVO.Picture = sku.Picture
	skuVO.Specs = sku.Specs
	Ok(w, skuVO)
}

// DoEditSKU 新增/编辑 SKU
func (m *MallHttpServiceImpl) DoEditSKU(w http.ResponseWriter, r *http.Request) {
	req := new(CMSSKUReq)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		Error(w, errcode.BadRequestParam, "参数解析失败")
		return
	}
	if err := validator.New().Struct(req); err != nil {
		Error(w, errcode.BadRequestParam, "缺少参数")
		return
	}
	goods, err := m.goodsService.GetGoodsById(r.Context(), req.GoodsId)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	if goods.ID == consts.ZERO || goods.Del == consts.DELETE {
		Error(w, errcode.NotFoundGoods, "商品不存在")
		return
	}
	if req.Id == consts.ZERO {
		sku := &entity.WechatMallSkuDO{}
		sku.Title = req.Title
		sku.Price = req.Price
		sku.Code = req.Code
		sku.Stock = req.Stock
		sku.GoodsID = req.GoodsId
		sku.Online = req.Online
		sku.Picture = req.Picture
		sku.Specs = req.Specs
		m.skuService.AddSKU(r.Context(), sku)
	} else {
		sku, err := m.skuService.GetSKUById(r.Context(), req.Id)
		if err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
		if sku.ID == consts.ZERO || sku.Del == consts.DELETE {
			Error(w, errcode.NotFoundGoodsSku, "商品SKU不存在")
			return
		}
		sku.Title = req.Title
		sku.Price = req.Price
		sku.Code = req.Code
		sku.Stock = req.Stock
		sku.GoodsID = req.GoodsId
		sku.Online = req.Online
		sku.Picture = req.Picture
		sku.Specs = req.Specs
		if err := m.skuService.UpdateSKUById(r.Context(), sku); err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
	}
	Ok(w, "ok")
}

// DoDeleteSKU 删除-单个SKU
func (m *MallHttpServiceImpl) DoDeleteSKU(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	sku, err := m.skuService.GetSKUById(r.Context(), id)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	if sku.ID == consts.ZERO || sku.Del == consts.DELETE {
		Error(w, errcode.NotFoundGoodsSku, "商品SKU不存在")
		return
	}
	sku.Del = consts.DELETE
	if err := m.skuService.UpdateSKUById(r.Context(), sku); err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	Ok(w, "ok")
}

// GetSpecificationList 查询-规格列表
func (m *MallHttpServiceImpl) GetSpecificationList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])
	specList, total, err := m.specService.GetSpecificationList(r.Context(), page, size)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	specVOs := make([]*view.CMSSpecificationVO, 0)
	for _, v := range specList {
		specVO := &view.CMSSpecificationVO{}
		specVO.Id = v.ID
		specVO.Name = v.Name
		specVO.Description = v.Description
		specVO.Unit = v.Unit
		specVO.Standard = v.Standard
		specVOs = append(specVOs, specVO)
	}
	data := make(map[string]interface{})
	data["list"] = specVOs
	data["total"] = total
	Ok(w, data)
}

// GetSpecification 查询-单个规格
func (m *MallHttpServiceImpl) GetSpecification(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	spec, err := m.specService.GetSpecificationById(r.Context(), id)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	if spec.ID == consts.ZERO || spec.Del == consts.DELETE {
		Error(w, errcode.NotFoundSpecification, "规格不存在")
		return
	}
	specVO := &view.CMSSpecificationVO{
		Id:          spec.ID,
		Name:        spec.Name,
		Description: spec.Description,
		Unit:        spec.Unit,
		Standard:    spec.Standard,
	}
	Ok(w, specVO)
}

// DoEditSpecification 新增/编辑-规格
func (m *MallHttpServiceImpl) DoEditSpecification(w http.ResponseWriter, r *http.Request) {
	req := new(CMSSpecificationReq)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		Error(w, errcode.BadRequestParam, "参数解析失败")
		return
	}
	if req.Id == consts.ZERO {
		spec, err := m.specService.GetSpecificationByName(r.Context(), req.Name)
		if err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
		if spec.ID != consts.ZERO {
			Error(w, errcode.NotAllowOperation, "规格名已存在")
			return
		}
		spec.Name = req.Name
		spec.Description = req.Description
		spec.Unit = req.Unit
		spec.Standard = req.Standard
		m.specService.AddSpecification(r.Context(), spec)
	} else {
		spec, err := m.specService.GetSpecificationByName(r.Context(), req.Name)
		if err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
		if spec.ID != consts.ZERO && spec.ID != req.Id {
			Error(w, errcode.NotAllowOperation, "规格名已存在")
		}
		spec, err = m.specService.GetSpecificationById(r.Context(), req.Id)
		if err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
		if spec.ID == consts.ZERO || spec.Del == consts.DELETE {
			Error(w, errcode.NotFoundSpecification, "规格不存在")
		}
		spec.Name = req.Name
		spec.Description = req.Description
		spec.Unit = req.Unit
		spec.Standard = req.Standard
		if err := m.specService.UpdateSpecificationById(r.Context(), spec); err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
	}
	Ok(w, "ok")
}

// DoDeleteSpecification 删除-规格
func (m *MallHttpServiceImpl) DoDeleteSpecification(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	spec, err := m.specService.GetSpecificationById(r.Context(), id)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	if spec.ID == consts.ZERO || spec.Del == consts.DELETE {
		Error(w, errcode.NotFoundSpecification, "规格不存在")
		return
	}
	attrList, err := m.specService.GetSpecificationAttrList(r.Context(), id)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	if len(attrList) > 0 {
		Error(w, errcode.NotAllowOperation, "该规格下有属性，不能删除")
	}
	goodsNum, err := m.goodsService.CountGoodsSpecBySpecId(r.Context(), id)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	if goodsNum > 0 {
		Error(w, errcode.NotAllowOperation, "部分商品正在使用该规格，不能删除")
	}
	spec.Del = consts.DELETE
	if err := m.specService.UpdateSpecificationById(r.Context(), spec); err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	Ok(w, "ok")
}

// GetSpecificationAttrList 查询-单个规格-全部属性
func (m *MallHttpServiceImpl) GetSpecificationAttrList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	specId, _ := strconv.Atoi(vars["specId"])
	specAttrList, err := m.specService.GetSpecificationAttrList(r.Context(), specId)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	attrVOs := make([]*view.CMSSpecificationAttrVO, 0)
	for _, v := range specAttrList {
		attrVO := &view.CMSSpecificationAttrVO{}
		attrVO.Id = v.ID
		attrVO.SpecId = v.SpecID
		attrVO.Value = v.Value
		attrVO.Extend = v.Extend
		attrVOs = append(attrVOs, attrVO)
	}
	Ok(w, attrVOs)
}

// GetSpecificationAttr 查询-单个规格-单个属性
func (m *MallHttpServiceImpl) GetSpecificationAttr(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	spec, err := m.specService.GetSpecificationAttrById(r.Context(), id)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	if spec.ID == consts.ZERO || spec.Del == consts.DELETE {
		Error(w, errcode.NotFoundSpecificationAttr, "规格属性不存在")
		return
	}
	attrVO := &view.CMSSpecificationAttrVO{
		Id:     spec.ID,
		SpecId: spec.SpecID,
		Value:  spec.Value,
		Extend: spec.Extend,
	}
	Ok(w, attrVO)
}

// DoEditSpecificationAttr 新增/更新-规格-单个属性
func (m *MallHttpServiceImpl) DoEditSpecificationAttr(w http.ResponseWriter, r *http.Request) {
	req := new(CMSSpecificationAttrReq)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		Error(w, errcode.BadRequestParam, "参数解析失败")
		return
	}
	spec, err := m.specService.GetSpecificationById(r.Context(), req.SpecId)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	if spec.ID == consts.ZERO || spec.Del == consts.DELETE {
		Error(w, errcode.NotFoundSpecificationAttr, "规格属性不存在")
		return
	}
	if req.Id == consts.ZERO {
		spec, err := m.specService.GetSpecificationAttrByValue(r.Context(), req.Value)
		if err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
		if spec.ID != consts.ZERO {
			Error(w, errcode.NotAllowOperation, "属性名已存在")
			return
		}
		spec.SpecID = req.SpecId
		spec.Value = req.Value
		spec.Extend = req.Extend
		if err := m.specService.AddSpecificationAttr(r.Context(), spec); err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
	} else {
		spec, err := m.specService.GetSpecificationAttrByValue(r.Context(), req.Value)
		if err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
		if spec.ID != consts.ZERO && spec.ID != req.Id {
			Error(w, errcode.NotAllowOperation, "属性名已存在")
			return
		}
		spec, err = m.specService.GetSpecificationAttrById(r.Context(), req.Id)
		if err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
		if spec.ID == consts.ZERO || spec.Del == consts.DELETE {
			Error(w, errcode.NotFoundSpecificationAttr, "规格属性不存在")
			return
		}
		spec.SpecID = req.SpecId
		spec.Value = req.Value
		spec.Extend = req.Extend
		if err := m.specService.UpdateSpecificationAttrById(r.Context(), spec); err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
	}
	Ok(w, "ok")
}

// DoDeleteSpecificationAttr 删除-单个规格-单个属性
func (m *MallHttpServiceImpl) DoDeleteSpecificationAttr(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	spec, err := m.specService.GetSpecificationAttrById(r.Context(), id)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	if spec.ID == consts.ZERO || spec.Del == consts.DELETE {
		Error(w, errcode.NotFoundSpecificationAttr, "规格属性不存在")
		return
	}
	total, err := m.skuService.CountAttrRelatedSku(r.Context(), id)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	if total > 0 {
		Error(w, errcode.NotAllowOperation, "部分商品正在使用该属性，禁止删除")
		return
	}
	spec.Del = consts.DELETE
	if err := m.specService.UpdateSpecificationAttrById(r.Context(), spec); err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	Ok(w, "ok")
}
