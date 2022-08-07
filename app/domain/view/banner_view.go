package view

type PortalBannerVO struct {
	Id           int    `json:"id"`
	Picture      string `json:"picture"`
	BusinessType int    `json:"businessType"`
	BusinessId   int    `json:"businessId"`
}

type CMSGoodsBannerVO struct {
	Id            int    `json:"id"`
	Picture       string `json:"picture"`
	Name          string `json:"name"`
	GoodsId       int    `json:"goodsId"`
	CategoryId    int    `json:"categoryId"`
	SubCategoryId int    `json:"subCategoryId"`
	Status        int    `json:"status"`
}

type CMSBannerVO struct {
	Id      int    `json:"id"`
	Picture string `json:"picture"`
	Name    string `json:"name"`
	Status  int    `json:"status"`
}
