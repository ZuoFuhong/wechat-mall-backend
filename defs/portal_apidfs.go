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

type PortalActivityVO struct {
	Id              int    `json:"id"`
	Online          int    `json:"online"`
	StartTime       string `json:"start_time"`
	EndTime         string `json:"end_time"`
	EntrancePicture string `json:"entrance_picture"`
}
