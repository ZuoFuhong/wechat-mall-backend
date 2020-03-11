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
