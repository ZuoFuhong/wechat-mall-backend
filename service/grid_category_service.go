package service

import (
	"wechat-mall-backend/dbops"
	"wechat-mall-backend/model"
)

type IGridCategoryService interface {
	GetGridCategoryList(page, size int) (*[]model.GridCategory, int)
	GetGridCategoryById(id int) *model.GridCategory
	GetGridCategoryByName(name string) *model.GridCategory
	AddGridCategory(gridC *model.GridCategory)
	UpdateGridCategory(gridC *model.GridCategory)
}

type gridCategoryService struct {
}

func NewGridCategoryService() IGridCategoryService {
	service := gridCategoryService{}
	return &service
}

func (g *gridCategoryService) GetGridCategoryList(page, size int) (*[]model.GridCategory, int) {
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

func (g *gridCategoryService) GetGridCategoryById(id int) *model.GridCategory {
	gridC, err := dbops.QueryGridCategoryById(id)
	if err != nil {
		panic(err)
	}
	return gridC
}

func (g *gridCategoryService) GetGridCategoryByName(name string) *model.GridCategory {
	gridC, err := dbops.QueryGridCategoryByName(name)
	if err != nil {
		panic(err)
	}
	return gridC
}

func (g *gridCategoryService) AddGridCategory(gridC *model.GridCategory) {
	err := dbops.InsertGridCategory(gridC)
	if err != nil {
		panic(err)
	}
}

func (g *gridCategoryService) UpdateGridCategory(gridC *model.GridCategory) {
	err := dbops.UpdateGridCategoryById(gridC)
	if err != nil {
		panic(err)
	}
}
