package service

import (
	"wechat-mall-backend/dbops"
	"wechat-mall-backend/model"
)

type ISpecificationService interface {
	GetSpecificationList(page, size int) (*[]model.WechatMallSpecificationDO, int)
	GetSpecificationById(id int) *model.WechatMallSpecificationDO
	GetSpecificationByName(name string) *model.WechatMallSpecificationDO
	UpdateSpecificationById(spec *model.WechatMallSpecificationDO)
	AddSpecification(spec *model.WechatMallSpecificationDO)
	GetSpecificationAttrList(specId int) *[]model.WechatMallSpecificationAttrDO
	GetSpecificationAttrById(id int) *model.WechatMallSpecificationAttrDO
	GetSpecificationAttrByValue(value string) *model.WechatMallSpecificationAttrDO
	UpdateSpecificationAttrById(spec *model.WechatMallSpecificationAttrDO)
	AddSpecificationAttr(spec *model.WechatMallSpecificationAttrDO)
}

type specificationService struct {
}

func NewSpecificationService() ISpecificationService {
	service := specificationService{}
	return &service
}

func (ss *specificationService) GetSpecificationList(page, size int) (*[]model.WechatMallSpecificationDO, int) {
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

func (ss *specificationService) GetSpecificationById(id int) *model.WechatMallSpecificationDO {
	spec, err := dbops.QuerySpecificationById(id)
	if err != nil {
		panic(err)
	}
	return spec
}

func (ss *specificationService) GetSpecificationByName(name string) *model.WechatMallSpecificationDO {
	spec, err := dbops.QuerySpecificationByName(name)
	if err != nil {
		panic(err)
	}
	return spec
}

func (ss *specificationService) UpdateSpecificationById(spec *model.WechatMallSpecificationDO) {
	err := dbops.UpdateSpecificationById(spec)
	if err != nil {
		panic(err)
	}
}

func (ss *specificationService) AddSpecification(spec *model.WechatMallSpecificationDO) {
	err := dbops.AddSpecification(spec)
	if err != nil {
		panic(err)
	}
}

func (ss *specificationService) GetSpecificationAttrList(specId int) *[]model.WechatMallSpecificationAttrDO {
	attrList, err := dbops.QuerySpecificationAttrList(specId)
	if err != nil {
		panic(err)
	}
	return attrList
}

func (ss *specificationService) GetSpecificationAttrById(id int) *model.WechatMallSpecificationAttrDO {
	attr, err := dbops.QuerySpecificationAttrById(id)
	if err != nil {
		panic(err)
	}
	return attr
}

func (ss *specificationService) GetSpecificationAttrByValue(value string) *model.WechatMallSpecificationAttrDO {
	spec, err := dbops.QuerySpecificationAttrByValue(value)
	if err != nil {
		panic(err)
	}
	return spec
}

func (ss *specificationService) UpdateSpecificationAttrById(attr *model.WechatMallSpecificationAttrDO) {
	err := dbops.UpdateSpecificationAttrById(attr)
	if err != nil {
		panic(err)
	}
}

func (ss *specificationService) AddSpecificationAttr(attr *model.WechatMallSpecificationAttrDO) {
	err := dbops.AddSpecificationAttr(attr)
	if err != nil {
		panic(err)
	}
}
