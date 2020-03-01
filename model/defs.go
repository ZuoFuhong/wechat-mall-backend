package model

type ID = int

type WxappUser struct {
	Id         ID
	Openid     string
	Nickname   string
	Avatar     string
	Mobile     string
	City       string
	CreateTime string
	UpdateTime string
}

type CMSUser struct {
	Id         ID
	Username   string
	Password   string
	Email      string
	Mobile     string
	Avatar     string
	CreateTime string
	UpdateTime string
}
