package defs

type WxappLoginResp struct {
	Token               string `json:"token" validate:"required"`
	ExpirationInMinutes int    `json:"expiration_in_minutes" validate:"required"`
}

type PortalBannerVO struct {
	Id      int    `json:"id"`
	Picture string `json:"picture"`
}

type PortalGridCategoryVO struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	CategoryId int    `json:"category"`
	Picture    string `json:"picture"`
}
