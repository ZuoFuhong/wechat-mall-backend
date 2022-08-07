package view

type CMSUserGroupVO struct {
	Id          int         `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Auths       interface{} `json:"auths"`
}

type WxappLoginVO struct {
	Token string `json:"token"`
}

type WxappUserInfoVO struct {
	Uid      int    `json:"uid"`
	Nickname string `json:"nickName"`
	Avatar   string `json:"avatar"`
	Mobile   string `json:"mobile"`
}

type PortalAddressVO struct {
	Id          int    `json:"id"`          // 地址ID
	Contacts    string `json:"contacts"`    // 联系人
	Mobile      string `json:"mobile"`      // 手机号
	ProvinceId  string `json:"provinceId"`  // 省份编码
	CityId      string `json:"cityId"`      // 城市编码
	AreaId      string `json:"areaId"`      // 地区编码
	ProvinceStr string `json:"provinceStr"` // 省份
	CityStr     string `json:"cityStr"`     // 城市
	AreaStr     string `json:"areaStr"`     // 地区
	Address     string `json:"address"`     // 详细地址
	IsDefault   int    `json:"isDefault"`   // 默认收货地址：0-否 1-是
}
