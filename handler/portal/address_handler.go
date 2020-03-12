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

// 新增/更新-收货地址
func (h *Handler) EditAddress(w http.ResponseWriter, r *http.Request) {
	req := defs.PortalAddressReq{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		panic(errs.ErrorParameterValidate)
	}
	userId := r.Context().Value(defs.ContextKey).(int)
	if req.Id == 0 {
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
		if address.Id == 0 {
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
	if address.Id == 0 {
		panic(errs.ErrorAddress)
	}
	address.Del = 1
	h.service.AddressService.UpdateAddress(address)
	defs.SendNormalResponse(w, "ok")
}
