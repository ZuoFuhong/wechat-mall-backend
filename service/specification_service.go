package service

import (
	"wechat-mall-backend/dbops"
	"wechat-mall-backend/model"
)

type ISpecificationService interface {
	GetSpecificationList(page, size int) (*[]model.Specification, int)
	GetSpecificationById(id int) *model.Specification
	GetSpecificationByName(name string) *model.Specification
	UpdateSpecificationById(spec *model.Specification)
	AddSpecification(spec *model.Specification)
	GetSpecificationAttrList(page, size int) (*[]model.SpecificationAttr, int)
	GetSpecificationAttrById(id int) *model.SpecificationAttr
	GetSpecificationAttrByValue(value string) *model.SpecificationAttr
	UpdateSpecificationAttrById(spec *model.SpecificationAttr)
	AddSpecificationAttr(spec *model.SpecificationAttr)
}

type specificationService struct {
}

func NewSpecificationService() ISpecificationService {
	service := specificationService{}
	return &service
}

func (ss *specificationService) GetSpecificationList(page, size int) (*[]model.Specification, int) {
	specList, err := dbops.QuerySpecificationList(page, size)
	if err != nil {
		panic(err)
	}
	total, err := dbops.CountSpecification()
	if err != nil {
		panic(err)
	}
	return specList, total
}

func (ss *specificationService) GetSpecificationById(id int) *model.Specification {
	spec, err := dbops.QuerySpecificationById(id)
	if err != nil {
		panic(err)
	}
	return spec
}

func (ss *specificationService) GetSpecificationByName(name string) *model.Specification {
	spec, err := dbops.QuerySpecificationByName(name)
	if err != nil {
		panic(err)
	}
	return spec
}

func (ss *specificationService) UpdateSpecificationById(spec *model.Specification) {
	err := dbops.UpdateSpecificationById(spec)
	if err != nil {
		panic(err)
	}
}

func (ss *specificationService) AddSpecification(spec *model.Specification) {
	err := dbops.AddSpecification(spec)
	if err != nil {
		panic(err)
	}
}

func (ss *specificationService) GetSpecificationAttrList(page, size int) (*[]model.SpecificationAttr, int) {
	attrList, err := dbops.QuerySpecificationAttrList(page, size)
	if err != nil {
		panic(err)
	}
	total, err := dbops.CountSpecificationAttr()
	if err != nil {
		panic(err)
	}
	return attrList, total
}

func (ss *specificationService) GetSpecificationAttrById(id int) *model.SpecificationAttr {
	attr, err := dbops.QuerySpecificationAttrById(id)
	if err != nil {
		panic(err)
	}
	return attr
}

func (ss *specificationService) GetSpecificationAttrByValue(value string) *model.SpecificationAttr {
	spec, err := dbops.QuerySpecificationAttrByValue(value)
	if err != nil {
		panic(err)
	}
	return spec
}

func (ss *specificationService) UpdateSpecificationAttrById(attr *model.SpecificationAttr) {
	err := dbops.UpdateSpecificationAttrById(attr)
	if err != nil {
		panic(err)
	}
}

func (ss *specificationService) AddSpecificationAttr(attr *model.SpecificationAttr) {
	err := dbops.AddSpecificationAttr(attr)
	if err != nil {
		panic(err)
	}
}
