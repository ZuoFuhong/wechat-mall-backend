package service

import (
	"wechat-mall-backend/dbops"
	"wechat-mall-backend/model"
)

type IGridCategoryService interface {
	GetGridCategoryList(page, size int) (*[]model.WechatMallGridCategoryDO, int)
	GetGridCategoryById(id int) *model.WechatMallGridCategoryDO
	GetGridCategoryByName(name string) *model.WechatMallGridCategoryDO
	AddGridCategory(gridC *model.WechatMallGridCategoryDO)
	UpdateGridCategory(gridC *model.WechatMallGridCategoryDO)
	CountCategoryBindGrid(categoryId int) int
}

type gridCategoryService struct {
}

func NewGridCategoryService() IGridCategoryService {
	service := gridCategoryService{}
	return &service
}

func (g *gridCategoryService) GetGridCategoryList(page, size int) (*[]model.WechatMallGridCategoryDO, int) {
	gridCList, err := dbops.QueryGridCategoryList(page, size)
	if err != nil {
		panic(err)
	}
	total, err := dbops.CountGridCategory()
	if err != nil {
		panic(err)
	}
	return gridCList, total
}

func (g *gridCategoryService) GetGridCategoryById(id int) *model.WechatMallGridCategoryDO {
	gridC, err := dbops.QueryGridCategoryById(id)
	if err != nil {
		panic(err)
	}
	return gridC
}

func (g *gridCategoryService) GetGridCategoryByName(name string) *model.WechatMallGridCategoryDO {
	gridC, err := dbops.QueryGridCategoryByName(name)
	if err != nil {
		panic(err)
	}
	return gridC
}

func (g *gridCategoryService) AddGridCategory(gridC *model.WechatMallGridCategoryDO) {
	err := dbops.InsertGridCategory(gridC)
	if err != nil {
		panic(err)
	}
}

func (g *gridCategoryService) UpdateGridCategory(gridC *model.WechatMallGridCategoryDO) {
	err := dbops.UpdateGridCategoryById(gridC)
	if err != nil {
		panic(err)
	}
}

// 统计分类绑定的宫格
func (g *gridCategoryService) CountCategoryBindGrid(categoryId int) int {
	total, err := dbops.CountGridByCategoryId(categoryId)
	if err != nil {
		panic(err)
	}
	return total
}
