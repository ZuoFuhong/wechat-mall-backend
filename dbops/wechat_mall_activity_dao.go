package dbops

import (
	"strconv"
	"time"
	"wechat-mall-backend/model"
)

const activityColumnList = `
id, title, name, remark, online, start_time, end_time, description, entrance_picture,
is_del, create_time, update_time
`

func QueryActivityList(page, size, online int) (*[]model.WechatMallActivityDO, error) {
	sql := "SELECT " + activityColumnList + " FROM wechat_mall_activity WHERE is_del = 0"
	if online != 0 {
		sql += " AND online = " + strconv.Itoa(online)
	}
	if page > 0 && size > 0 {
		sql += " LIMIT " + strconv.Itoa((page-1)*size) + ", " + strconv.Itoa(size)
	}
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	var aList []model.WechatMallActivityDO
	for rows.Next() {
		activity := model.WechatMallActivityDO{}
		err := rows.Scan(&activity.Id, &activity.Title, &activity.Name, &activity.Remark, &activity.Online, &activity.StartTime,
			&activity.EndTime, &activity.Description, &activity.EntrancePicture, &activity.Del,
			&activity.CreateTime, &activity.UpdateTime)
		if err != nil {
			return nil, err
		}
		aList = append(aList, activity)
	}
	return &aList, nil
}

func CountActivity() (int, error) {
	sql := " SELECT COUNT(*) FROM wechat_mall_activity WHERE is_del = 0"
	rows, err := dbConn.Query(sql)
	if err != nil {
		return 0, err
	}
	total := 0
	for rows.Next() {
		err := rows.Scan(&total)
		if err != nil {
			return 0, err
		}
	}
	return total, nil
}

func InsertActivity(activity *model.WechatMallActivityDO) error {
	sql := "INSERT INTO wechat_mall_activity (" + activityColumnList[4:] + " ) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(activity.Title, activity.Name, activity.Remark, activity.Online, activity.StartTime,
		activity.EndTime, activity.Description, activity.EntrancePicture, 0, time.Now(), time.Now())

	if err != nil {
		return err
	}
	return nil
}

func QueryActivityById(id int) (*model.WechatMallActivityDO, error) {
	sql := "SELECT " + activityColumnList + " FROM wechat_mall_activity WHERE is_del = 0 AND id = " + strconv.Itoa(id)
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	activity := model.WechatMallActivityDO{}
	if rows.Next() {
		err := rows.Scan(&activity.Id, &activity.Title, &activity.Name, &activity.Remark, &activity.Online,
			&activity.StartTime, &activity.EndTime, &activity.Description, &activity.EntrancePicture, &activity.Del,
			&activity.CreateTime, &activity.EndTime)
		if err != nil {
			return nil, err
		}
	}
	return &activity, nil
}

func QueryActivityByName(name string) (*model.WechatMallActivityDO, error) {
	sql := "SELECT " + activityColumnList + " FROM wechat_mall_activity WHERE is_del = 0 AND name = '" + name + "'"
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	activity := model.WechatMallActivityDO{}
	if rows.Next() {
		err := rows.Scan(&activity.Id, &activity.Title, &activity.Name, &activity.Remark, &activity.Online,
			&activity.StartTime, &activity.EndTime, &activity.Description, &activity.EntrancePicture, &activity.Del,
			&activity.CreateTime, &activity.EndTime)
		if err != nil {
			return nil, err
		}
	}
	return &activity, nil
}

func UpdateActivityById(activity *model.WechatMallActivityDO) error {
	sql := `
UPDATE wechat_mall_activity 
SET title = ?, name = ?, remark = ?, online = ?, start_time = ?, end_time = ?, description = ?, entrance_picture = ?, 
internal_top_picture = ?, is_del = ?, update_time = ? WHERE id = ?
`
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(activity.Title, activity.Name, activity.Remark, activity.Online, activity.StartTime,
		activity.EndTime, activity.Description, activity.EntrancePicture, activity.Del, time.Now(), activity.Id)
	if err != nil {
		return err
	}
	return nil
}
