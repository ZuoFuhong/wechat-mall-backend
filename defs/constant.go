package defs

const (
	ContextKey         = "uid"
	MiniappTokenPrefix = "miniappToken:"
	AccessTokenExpire  = 2 * 3600
	RefreshTokenExpire = 30 * 24 * 3600
	CMSCodePrefix      = "cmscode:"
	CMSCodeExpire      = 300
)

type SkuSpecs struct {
	Key     string `json:"key"`
	Value   string `json:"value"`
	KeyId   int    `json:"keyId"`
	ValueId int    `json:"valueId"`
}

type AddressSnapshot struct {
	Contacts    string `json:"contacts"`    // 联系人
	Mobile      string `json:"mobile"`      // 手机号
	ProvinceId  string `json:"provinceId"`  // 省份编码
	CityId      string `json:"cityId"`      // 城市编码
	AreaId      string `json:"areaId"`      // 地区编码
	ProvinceStr string `json:"provinceStr"` // 省份
	CityStr     string `json:"cityStr"`     // 城市
	AreaStr     string `json:"areaStr"`     // 地区
	Address     string `json:"address"`     // 详细地址
}
