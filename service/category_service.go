package service

import (
	"wechat-mall-backend/dbops"
	"wechat-mall-backend/model"
)

type ICategoryService interface {
	GetCategoryList(page, size int) (*[]model.Category, int)
	GetCategoryById(id int) *Category
	GetCategoryByName(name string) *Category
	AddCategory(category *model.Category)
	UpdateCategory(category *model.Category)
}

type categoryService struct {
}

func NewCategoryService() ICategoryService {
	service := &categoryService{}
	return service
}

func (cs *categoryService) GetCategoryList(page, size int) (*[]model.Category, int) {
	cateList, err := dbops.QueryCategoryList(page, size)
	if err != nil {
		panic(err)
	}
	total, err := dbops.CountCategory()
	if err != nil {
		panic(err)
	}
	return cateList, total
}

func (cs *categoryService) GetCategoryById(id int) *Category {
	category, err := dbops.QueryCategoryById(id)
	if err != nil {
		panic(err)
	}
	return (*Category)(category)
}

func (cs *categoryService) GetCategoryByName(name string) *Category {
	category, err := dbops.QueryCategoryByName(name)
	if err != nil {
		panic(err)
	}
	return (*Category)(category)
}

func (cs *categoryService) AddCategory(category *model.Category) {
	err := dbops.InsertCategory(category)
	if err != nil {
		panic(err)
	}
}

func (cs *categoryService) UpdateCategory(category *model.Category) {
	err := dbops.UpdateCategoryById(category)
	if err != nil {
		panic(err)
	}
}
