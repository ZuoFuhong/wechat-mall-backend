package defs

const (
	AccessTokenExpire  = 2 * 3600
	RefreshTokenExpire = 30 * 24 * 3600
	MiniappTokenPrefix = "miniappToken:"
	MiniappTokenExpire = 2 * 3600
	CMSCodePrefix      = "cmscode:"
	CMSCodeExpire      = 300
	ContextKey         = "uid"
)

type SkuSpecs struct {
	Key     string `json:"key"`
	Value   string `json:"value"`
	KeyId   int    `json:"key_id"`
	ValueId int    `json:"value_id"`
}

type AddressSnapshot struct {
	Contacts    string `json:"contacts"`    // 联系人
	Mobile      string `json:"mobile"`      // 手机号
	ProvinceId  int    `json:"provinceId"`  // 省份编码
	CityId      int    `json:"cityId"`      // 城市编码
	AreaId      int    `json:"areaId"`      // 地区编码
	ProvinceStr string `json:"provinceStr"` // 省份
	CityStr     string `json:"cityStr"`     // 城市
	AreaStr     string `json:"areaStr"`     // 地区
	Address     string `json:"address"`     // 详细地址
}
