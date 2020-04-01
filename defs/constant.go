package defs

const (
	ContextKey         = "uid"
	MiniappTokenPrefix = "miniappToken:"
	AccessTokenExpire  = 2 * 3600
	RefreshTokenExpire = 30 * 24 * 3600
	CMSCodePrefix      = "cmscode:"
	CMSCodeExpire      = 300
	ALL                = -999
	CartMax            = 99
	CouponMax          = 10
	ADMIN              = 1
	ZERO               = 0
	DELETE             = 1
	ONLINE             = 1
	OFFLINE            = 0
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

// 报表数据-订单统计
type OrderSaleData struct {
	Time       string `json:"time"`       // 下单时间
	OrderNum   int    `json:"orderNum"`   // 订单数
	SaleAmount string `json:"saleAmount"` // 销售额
}
