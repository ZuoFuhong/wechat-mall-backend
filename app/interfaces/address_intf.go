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

// GetAddressList 查询-收货地址列表
func (m *MallHttpServiceImpl) GetAddressList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])
	userId := r.Context().Value(consts.ContextKey).(int)

	addressList, total, err := m.addressService.GetAddressList(r.Context(), userId, page, size)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	voList := make([]*view.PortalAddressVO, 0)
	for _, v := range addressList {
		addressVO := &view.PortalAddressVO{}
		addressVO.Id = v.ID
		addressVO.Contacts = v.Contacts
		addressVO.Mobile = v.Mobile
		addressVO.ProvinceId = v.ProvinceID
		addressVO.CityId = v.CityID
		addressVO.AreaId = v.AreaID
		addressVO.ProvinceStr = v.ProvinceStr
		addressVO.CityStr = v.CityStr
		addressVO.AreaStr = v.AreaStr
		addressVO.Address = v.Address
		addressVO.IsDefault = v.IsDefault
		voList = append(voList, addressVO)
	}
	data := make(map[string]interface{})
	data["list"] = voList
	data["total"] = total
	Ok(w, data)
}

// GetAddress 查询-单个地址
func (m *MallHttpServiceImpl) GetAddress(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	addressId, _ := strconv.Atoi(vars["id"])
	addressDO, err := m.addressService.GetAddress(r.Context(), addressId)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	if addressDO.ID == consts.ZERO || addressDO.Del == consts.DELETE {
		Error(w, errcode.NotFoundUserAddress, "地址不存在")
		return
	}
	addressVO := &view.PortalAddressVO{
		Id:          addressDO.ID,
		Contacts:    addressDO.Contacts,
		Mobile:      addressDO.Mobile,
		ProvinceId:  addressDO.ProvinceID,
		CityId:      addressDO.CityID,
		AreaId:      addressDO.AreaID,
		ProvinceStr: addressDO.ProvinceStr,
		CityStr:     addressDO.CityStr,
		AreaStr:     addressDO.AreaStr,
		Address:     addressDO.Address,
		IsDefault:   addressDO.IsDefault,
	}
	Ok(w, addressVO)
}

// EditAddress 新增/更新-收货地址
func (m *MallHttpServiceImpl) EditAddress(w http.ResponseWriter, r *http.Request) {
	req := new(PortalAddressReq)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		Error(w, errcode.BadRequestParam, "参数解析失败")
		return
	}
	userId := r.Context().Value(consts.ContextKey).(int)
	if req.Id == consts.ZERO {
		addressDO, err := m.addressService.GetDefaultAddress(r.Context(), userId)
		if err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
		isDefault := 0
		if addressDO.ID == 0 {
			isDefault = 1
		}
		address := &entity.WechatMallUserAddressDO{}
		address.UserID = userId
		address.Contacts = req.Contacts
		address.Mobile = req.Mobile
		address.ProvinceID = req.ProvinceId
		address.CityID = req.CityId
		address.AreaID = req.AreaId
		address.ProvinceStr = req.ProvinceStr
		address.CityStr = req.CityStr
		address.AreaStr = req.AreaStr
		address.Address = req.Address
		address.IsDefault = isDefault
		if err := m.addressService.AddAddress(r.Context(), address); err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
	} else {
		address, err := m.addressService.GetAddress(r.Context(), req.Id)
		if err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
		if address.ID == consts.ZERO || address.Del == consts.DELETE {
			Error(w, errcode.NotFoundUserAddress, "地址不存在")
			return
		}
		address.UserID = userId
		address.Contacts = req.Contacts
		address.Mobile = req.Mobile
		address.ProvinceID = req.ProvinceId
		address.CityID = req.CityId
		address.AreaID = req.AreaId
		address.ProvinceStr = req.ProvinceStr
		address.CityStr = req.CityStr
		address.AreaStr = req.AreaStr
		address.Address = req.Address
		address.IsDefault = req.IsDefault
		if err := m.addressService.UpdateAddress(r.Context(), address); err != nil {
			Error(w, errcode.ErrorInternalFaults, "系统繁忙")
			return
		}
	}
	Ok(w, "ok")
}

// DoDeleteAddress 删除-收货地址
func (m *MallHttpServiceImpl) DoDeleteAddress(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	addressId, _ := strconv.Atoi(vars["id"])

	address, err := m.addressService.GetAddress(r.Context(), addressId)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	if address.ID == consts.ZERO || address.Del == consts.DELETE {
		Error(w, errcode.NotFoundUserAddress, "地址不存在")
		return
	}
	address.Del = consts.DELETE
	if err := m.addressService.UpdateAddress(r.Context(), address); err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	Ok(w, "ok")
}

// GetDefaultAddress 查询-默认收货地址
func (m *MallHttpServiceImpl) GetDefaultAddress(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(consts.ContextKey).(int)
	addressDO, err := m.addressService.GetDefaultAddress(r.Context(), userId)
	if err != nil {
		Error(w, errcode.ErrorInternalFaults, "系统繁忙")
		return
	}
	addressVO := &view.PortalAddressVO{}
	addressVO.Id = addressDO.ID
	addressVO.Contacts = addressDO.Contacts
	addressVO.Mobile = addressDO.Mobile
	addressVO.ProvinceId = addressDO.ProvinceID
	addressVO.CityId = addressDO.CityID
	addressVO.AreaId = addressDO.AreaID
	addressVO.ProvinceStr = addressDO.ProvinceStr
	addressVO.CityStr = addressDO.CityStr
	addressVO.AreaStr = addressDO.AreaStr
	addressVO.Address = addressDO.Address
	addressVO.IsDefault = addressDO.IsDefault
	Ok(w, addressVO)
}
