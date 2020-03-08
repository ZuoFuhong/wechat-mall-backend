package service

import (
	"strconv"
	"strings"
	"wechat-mall-backend/dbops"
	"wechat-mall-backend/model"
)

type ISPUService interface {
	GetSPUList(page, size int) (*[]model.SPU, int)
	GetSPUById(id int) *model.SPU
	UpdateSPUById(spu *model.SPU)
	AddSPU(spu *model.SPU) int
	GetSPUSpecList(spuId int) *[]int
	AddSPUSpec(spuId int, specList string)
}

type sPUService struct {
}

func NewSPUService() ISPUService {
	service := &sPUService{}
	return service
}

func (s *sPUService) GetSPUList(page, size int) (*[]model.SPU, int) {
	spuList, err := dbops.QuerySPUList(page, size)
	if err != nil {
		panic(err)
	}
	total, err := dbops.CountSPU()
	if err != nil {
		panic(err)
	}
	return spuList, total
}

func (s *sPUService) GetSPUById(id int) *model.SPU {
	spu, err := dbops.QuerySPUById(id)
	if err != nil {
		panic(err)
	}
	return spu
}

func (s *sPUService) UpdateSPUById(spu *model.SPU) {
	err := dbops.UpdateSPUById(spu)
	if err != nil {
		panic(err)
	}
}

func (s *sPUService) AddSPU(spu *model.SPU) int {
	id, err := dbops.AddSPU(spu)
	if err != nil {
		panic(err)
	}
	return int(id)
}

func (s *sPUService) GetSPUSpecList(spuId int) *[]int {
	specList, err := dbops.GetSPUSpecList(spuId)
	if err != nil {
		panic(err)
	}
	specIds := []int{}
	for _, v := range *specList {
		specIds = append(specIds, v.SpecId)
	}
	return &specIds
}

func (s *sPUService) AddSPUSpec(spuId int, specList string) {
	err := dbops.DeleteSPUSpec(spuId)
	if err != nil {
		panic(err)
	}
	specIds := strings.Split(specList, ",")
	for _, v := range specIds {
		spec := model.SPUSpec{}
		spec.SpuId = spuId
		spec.SpecId, _ = strconv.Atoi(v)
		err := dbops.InsertSPUSpec(&spec)
		if err != nil {
			panic(err)
		}
	}
}
