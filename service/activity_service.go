package service

import (
	"wechat-mall-backend/dbops"
	"wechat-mall-backend/model"
)

type IActivityService interface {
	GetActivityList(page, size int) (*[]model.Activity, int)
	GetActivityById(id int) (activity *model.Activity)
	GetActivityByName(name string) (activity *model.Activity)
	AddActivity(service *model.Activity)
	UpdateActivity(activity *model.Activity)
}

type activityService struct {
}

func NewActivityService() IActivityService {
	service := activityService{}
	return &service
}

func (as *activityService) GetActivityList(page, size int) (*[]model.Activity, int) {
	activityList, err := dbops.QueryActivityList(page, size)
	if err != nil {
		panic(err)
	}
	total, err := dbops.CountActivity()
	if err != nil {
		panic(err)
	}
	return activityList, total
}

func (as *activityService) GetActivityById(id int) (activity *model.Activity) {
	activity, err := dbops.QueryActivityById(id)
	if err != nil {
		panic(err)
	}
	return activity
}

func (as *activityService) GetActivityByName(name string) (activity *model.Activity) {
	activity, err := dbops.QueryActivityByName(name)
	if err != nil {
		panic(err)
	}
	return activity
}

func (as *activityService) AddActivity(activity *model.Activity) {
	err := dbops.InsertActivity(activity)
	if err != nil {
		panic(err)
	}
}

func (as *activityService) UpdateActivity(activity *model.Activity) {
	err := dbops.UpdateActivityById(activity)
	if err != nil {
		panic(err)
	}
}
