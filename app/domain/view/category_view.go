package view

type CMSCategoryVO struct {
	Id          int    `json:"id"`
	ParentId    int    `json:"parentId"`
	Name        string `json:"name"`
	Sort        int    `json:"sort"`
	Online      int    `json:"online"`
	Picture     string `json:"picture"`
	Description string `json:"description"`
}

type PortalGridCategoryVO struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`       // 宫格标题
	CategoryId int    `json:"categoryId"` // 关联的分类
	Picture    string `json:"picture"`    // 宫格图标
}

type PortalCategoryVO struct {
	Id   int    `json:"id"`   // 分类ID
	Name string `json:"name"` // 分类名称
}

type CMSGridCategoryVO struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	CategoryId   int    `json:"categoryId"`
	CategoryName string `json:"categoryName"`
	Picture      string `json:"picture"`
}

type CMSGridCategoryDetailVO struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	CategoryId      int    `json:"categoryId"`
	SubCategoryId   int    `json:"subCategoryId"`
	SubCategoryName string `json:"subCategoryName"`
	Picture         string `json:"picture"`
}
