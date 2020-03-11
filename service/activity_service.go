package service

import (
	"wechat-mall-backend/dbops"
	"wechat-mall-backend/model"
)

type IActivityService interface {
	GetActivityList(page, size, online int) (*[]model.WechatMallActivityDO, int)
	GetActivityById(id int) (activity *model.WechatMallActivityDO)
	GetActivityByName(name string) (activity *model.WechatMallActivityDO)
	AddActivity(service *model.WechatMallActivityDO)
	UpdateActivity(activity *model.WechatMallActivityDO)
}

type activityService struct {
}

func NewActivityService() IActivityService {
	service := activityService{}
	return &service
}

func (as *activityService) GetActivityList(page, size, online int) (*[]model.WechatMallActivityDO, int) {
	activityList, err := dbops.QueryActivityList(page, size, online)
	if err != nil {
		panic(err)
	}
	total, err := dbops.CountActivity()
	if err != nil {
		panic(err)
	}
	return activityList, total
}

func (as *activityService) GetActivityById(id int) (activity *model.WechatMallActivityDO) {
	activity, err := dbops.QueryActivityById(id)
	if err != nil {
		panic(err)
	}
	return activity
}

func (as *activityService) GetActivityByName(name string) (activity *model.WechatMallActivityDO) {
	activity, err := dbops.QueryActivityByName(name)
	if err != nil {
		panic(err)
	}
	return activity
}

func (as *activityService) AddActivity(activity *model.WechatMallActivityDO) {
	err := dbops.InsertActivity(activity)
	if err != nil {
		panic(err)
	}
}

func (as *activityService) UpdateActivity(activity *model.WechatMallActivityDO) {
	err := dbops.UpdateActivityById(activity)
	if err != nil {
		panic(err)
	}
}
