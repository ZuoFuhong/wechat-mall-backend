package view

type CMSModuleVO struct {
	Id          int                `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	PageList    []*CMSModulePageVO `json:"pageList"`
}

type CMSModulePageVO struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
