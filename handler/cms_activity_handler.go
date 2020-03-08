package handler

import (
	"encoding/json"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"wechat-mall-backend/defs"
	"wechat-mall-backend/errs"
)

func (h *CMSHandler) GetActivityList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["page"])
	size, _ := strconv.Atoi(vars["size"])

	activityVOList := []defs.ActivityVO{}
	aList, total := h.service.ActivityService.GetActivityList(page, size)
	for _, v := range *aList {
		aVO := defs.ActivityVO{}
		aVO.Id = v.Id
		aVO.Title = v.Title
		aVO.Name = v.Name
		aVO.Remark = v.Remark
		aVO.Online = v.Online
		aVO.StartTime = v.StartTime
		aVO.EndTime = v.EndTime
		aVO.Description = v.Description
		aVO.EntrancePicture = v.EntrancePicture
		aVO.InternalTopPicture = v.InternalTopPicture
		activityVOList = append(activityVOList, aVO)
	}
	resp := make(map[string]interface{})
	resp["list"] = activityVOList
	resp["total"] = total
	sendNormalResponse(w, resp)
}

func (h *CMSHandler) GetActivity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	activity := h.service.ActivityService.GetActivityById(id)
	if activity.Id == 0 {
		panic(errs.ErrorActivity)
	}
	aVO := defs.ActivityVO{}
	aVO.Id = activity.Id
	aVO.Title = activity.Title
	aVO.Name = activity.Name
	aVO.Remark = activity.Remark
	aVO.Online = activity.Online
	aVO.StartTime = activity.StartTime
	aVO.EndTime = activity.EndTime
	aVO.Description = activity.Description
	aVO.EntrancePicture = activity.EntrancePicture
	aVO.InternalTopPicture = activity.InternalTopPicture
	sendNormalResponse(w, aVO)
}

func (h *CMSHandler) DoEditActivity(w http.ResponseWriter, r *http.Request) {
	req := defs.ActivityReq{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		panic(errs.ErrorParameterValidate)
	}
	validate := validator.New()
	if err = validate.Struct(req); err != nil {
		panic(errs.NewParameterError(err.Error()))
	}
	if req.Id == 0 {
		activity := h.service.ActivityService.GetActivityByName(req.Name)
		if activity.Id != 0 {
			panic(errs.NewErrorActivity("The name already exists"))
		}
		activity.Title = req.Title
		activity.Name = req.Name
		activity.Remark = req.Remark
		activity.Online = req.Online
		activity.StartTime = req.StartTime
		activity.EndTime = req.EndTime
		activity.Description = req.Description
		activity.EntrancePicture = req.EntrancePicture
		activity.InternalTopPicture = req.InternalTopPicture
		h.service.ActivityService.AddActivity(activity)
	} else {
		activity := h.service.ActivityService.GetActivityByName(req.Name)
		if activity.Id != 0 && activity.Id != req.Id {
			panic(errs.NewErrorActivity("The name already exists"))
		}
		activity = h.service.ActivityService.GetActivityById(req.Id)
		if activity.Id == 0 {
			panic(errs.ErrorActivity)
		}
		activity.Title = req.Title
		activity.Name = req.Name
		activity.Remark = req.Remark
		activity.Online = req.Online
		activity.StartTime = req.StartTime
		activity.EndTime = req.EndTime
		activity.Description = req.Description
		activity.EntrancePicture = req.EntrancePicture
		activity.InternalTopPicture = req.InternalTopPicture
		h.service.ActivityService.UpdateActivity(activity)
	}
	sendNormalResponse(w, "ok")
}

func (h *CMSHandler) DoDeleteActivity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	activity := h.service.ActivityService.GetActivityById(id)
	if activity.Id == 0 {
		panic(errs.ErrorActivity)
	}
	activity.Del = 1
	h.service.ActivityService.UpdateActivity(activity)
	sendNormalResponse(w, "ok")
}
