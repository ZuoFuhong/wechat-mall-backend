package dbops

import (
	"strconv"
	"time"
	"wechat-mall-backend/model"
)

const (
	moduleColumnList = `
id, name, description, is_del, create_time, update_time
`
	modulePageColumnList = `
id, module_id, name, description, is_del, create_time, update_time
`
	pagePermissionColumnList = `
id, group_id, page_id, is_del, create_time, update_time
`
)

func QueryModuleList() (*[]model.WechatMallModuleDO, error) {
	sql := "SELECT " + moduleColumnList + " FROM wechat_mall_module WHERE is_del = 0"
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	moduleList := []model.WechatMallModuleDO{}
	for rows.Next() {
		moduleDO := model.WechatMallModuleDO{}
		err := rows.Scan(&moduleDO.Id, &moduleDO.Name, &moduleDO.Description, &moduleDO.Del, &moduleDO.CreateTime, &moduleDO.UpdateTime)
		if err != nil {
			return nil, err
		}
		moduleList = append(moduleList, moduleDO)
	}
	return &moduleList, nil
}

func ListModulePage(moduleId int) (*[]model.WechatMallModulePageDO, error) {
	sql := "SELECT " + modulePageColumnList + " FROM wechat_mall_module_page WHERE is_del = 0 AND module_id = " + strconv.Itoa(moduleId)
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	pageList := []model.WechatMallModulePageDO{}
	for rows.Next() {
		pageDO := model.WechatMallModulePageDO{}
		err := rows.Scan(&pageDO.Id, &pageDO.ModuleId, &pageDO.Name, &pageDO.Description, &pageDO.Del, &pageDO.CreateTime, &pageDO.UpdateTime)
		if err != nil {
			return nil, err
		}
		pageList = append(pageList, pageDO)
	}
	return &pageList, nil
}

func QueryModulePageById(pageId int) (*model.WechatMallModulePageDO, error) {
	sql := "SELECT " + modulePageColumnList + " FROM wechat_mall_module_page WHERE is_del = 0 AND id = " + strconv.Itoa(pageId)
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	pageDO := model.WechatMallModulePageDO{}
	for rows.Next() {
		err := rows.Scan(&pageDO.Id, &pageDO.ModuleId, &pageDO.Name, &pageDO.Description, &pageDO.Del, &pageDO.CreateTime, &pageDO.UpdateTime)
		if err != nil {
			return nil, err
		}
	}
	return &pageDO, nil
}

func ListGroupPagePermission(groupId int) (*[]model.WechatMallGroupPagePermission, error) {
	sql := "SELECT " + pagePermissionColumnList + " FROM wechat_mall_group_page_permission WHERE is_del = 0 AND group_id = " + strconv.Itoa(groupId)
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	auths := []model.WechatMallGroupPagePermission{}
	for rows.Next() {
		auth := model.WechatMallGroupPagePermission{}
		err := rows.Scan(&auth.Id, &auth.GroupId, &auth.PageId, &auth.Del, &auth.CreateTime, &auth.UpdateTime)
		if err != nil {
			return nil, err
		}
		auths = append(auths, auth)
	}
	return &auths, nil
}

func AddGroupPagePermission(pageId, groupId int) error {
	sql := "INSERT INTO wechat_mall_group_page_permission ( " + pagePermissionColumnList[4:] + " ) VALUES (?, ?, ?, ?, ?)"
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return nil
	}
	_, err = stmt.Exec(groupId, pageId, 0, time.Now(), time.Now())
	return err
}

func RemoveGroupAllPagePermission(groupId int) error {
	sql := "UPDATE wechat_mall_group_page_permission SET is_del = 1 WHERE group_id = " + strconv.Itoa(groupId)
	_, err := dbConn.Exec(sql)
	return err
}
