package defs

type WxappLoginResp struct {
	Token               string `json:"token"`
	ExpirationInMinutes int    `json:"expiration_in_minutes"`
}

type CMSLoginReq struct {
	Username string
	Password string
}

type CMSLoginResp struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
