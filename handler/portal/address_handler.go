package portal

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"wechat-mall-backend/defs"
	"wechat-mall-backend/errs"
	"wechat-mall-backend/model"
)

// 查询-收货地址列表
func (h *Handler) GetAddressList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])
	userId := r.Context().Value(defs.ContextKey).(int)

	addressList, total := h.service.AddressService.GetAddressList(userId, page, size)
	addressVOList := []defs.PortalAddressVO{}
	for _, v := range *addressList {
		addressVO := defs.PortalAddressVO{}
		addressVO.Id = v.Id
		addressVO.Contacts = v.Contacts
		addressVO.Mobile = v.Mobile
		addressVO.ProvinceId = v.ProvinceId
		addressVO.CityId = v.CityId
		addressVO.AreaId = v.AreaId
		addressVO.ProvinceStr = v.ProvinceStr
		addressVO.CityStr = v.CityStr
		addressVO.AreaStr = v.AreaStr
		addressVO.Address = v.Address
		addressVO.IsDefault = v.IsDefault
		addressVOList = append(addressVOList, addressVO)
	}
	resp := make(map[string]interface{})
	resp["list"] = addressVOList
	resp["total"] = total
	defs.SendNormalResponse(w, resp)
}

// 查询-单个地址
func (h *Handler) GetAddress(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	addressId, _ := strconv.Atoi(vars["id"])
	addressDO := h.service.AddressService.GetAddress(addressId)
	if addressDO.Id == defs.ZERO || addressDO.Del == defs.DELETE {
		panic(errs.ErrorAddress)
	}
	addressVO := defs.PortalAddressVO{}
	addressVO.Id = addressDO.Id
	addressVO.Contacts = addressDO.Contacts
	addressVO.Mobile = addressDO.Mobile
	addressVO.ProvinceId = addressDO.ProvinceId
	addressVO.CityId = addressDO.CityId
	addressVO.AreaId = addressDO.AreaId
	addressVO.ProvinceStr = addressDO.ProvinceStr
	addressVO.CityStr = addressDO.CityStr
	addressVO.AreaStr = addressDO.AreaStr
	addressVO.Address = addressDO.Address
	addressVO.IsDefault = addressDO.IsDefault
	defs.SendNormalResponse(w, addressVO)
}

// 新增/更新-收货地址
func (h *Handler) EditAddress(w http.ResponseWriter, r *http.Request) {
	req := defs.PortalAddressReq{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		panic(errs.ErrorParameterValidate)
	}
	userId := r.Context().Value(defs.ContextKey).(int)
	if req.Id == defs.ZERO {
		address := model.WechatMallUserAddressDO{}
		address.UserId = userId
		address.Contacts = req.Contacts
		address.Mobile = req.Mobile
		address.ProvinceId = req.ProvinceId
		address.CityId = req.CityId
		address.AreaId = req.AreaId
		address.ProvinceStr = req.ProvinceStr
		address.CityStr = req.CityStr
		address.AreaStr = req.AreaStr
		address.Address = req.Address
		address.IsDefault = req.IsDefault
		h.service.AddressService.AddAddress(&address)
	} else {
		address := h.service.AddressService.GetAddress(req.Id)
		if address.Id == defs.ZERO || address.Del == defs.DELETE {
			panic(errs.ErrorAddress)
		}
		address.UserId = userId
		address.Contacts = req.Contacts
		address.Mobile = req.Mobile
		address.ProvinceId = req.ProvinceId
		address.CityId = req.CityId
		address.AreaId = req.AreaId
		address.ProvinceStr = req.ProvinceStr
		address.CityStr = req.CityStr
		address.AreaStr = req.AreaStr
		address.Address = req.Address
		address.IsDefault = req.IsDefault
		h.service.AddressService.UpdateAddress(address)
	}
	defs.SendNormalResponse(w, "ok")
}

// 删除-收货地址
func (h *Handler) DoDeleteAddress(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	addressId, _ := strconv.Atoi(vars["id"])

	address := h.service.AddressService.GetAddress(addressId)
	if address.Id == defs.ZERO || address.Del == defs.DELETE {
		panic(errs.ErrorAddress)
	}
	address.Del = defs.DELETE
	h.service.AddressService.UpdateAddress(address)
	defs.SendNormalResponse(w, "ok")
}
