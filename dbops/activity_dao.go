package dbops

import (
	"strconv"
	"time"
	"wechat-mall-backend/model"
)

const activityColumnList = `
id, title, name, remark, online, start_time, end_time, description, entrance_picture, internal_top_picture, 
is_del, create_time, update_time
`

func QueryActivityList(page, size int) (*[]model.Activity, error) {
	sql := "SELECT " + activityColumnList + " FROM wxapp_mall_activity WHERE is_del = 0 LIMIT ?, ?"
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query((page-1)*size, size)
	if err != nil {
		return nil, err
	}
	var aList []model.Activity
	for rows.Next() {
		activity := model.Activity{}
		err := rows.Scan(&activity.Id, &activity.Title, &activity.Name, &activity.Remark, &activity.Online, &activity.StartTime,
			&activity.EndTime, &activity.Description, &activity.EntrancePicture, &activity.InternalTopPicture, &activity.Del,
			&activity.CreateTime, &activity.UpdateTime)
		if err != nil {
			return nil, err
		}
		aList = append(aList, activity)
	}
	return &aList, nil
}

func CountActivity() (int, error) {
	sql := " SELECT COUNT(*) FROM wxapp_mall_activity WHERE is_del = 0"
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

func InsertActivity(activity *model.Activity) error {
	sql := "INSERT INTO wxapp_mall_activity (" + activityColumnList[4:] + " ) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(activity.Title, activity.Name, activity.Remark, activity.Online, activity.StartTime,
		activity.EndTime, activity.Description, activity.EntrancePicture, activity.InternalTopPicture,
		0, time.Now(), time.Now())

	if err != nil {
		return err
	}
	return nil
}

func QueryActivityById(id int) (*model.Activity, error) {
	sql := "SELECT " + activityColumnList + " FROM wxapp_mall_activity WHERE is_del = 0 AND id = " + strconv.Itoa(id)
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	activity := model.Activity{}
	if rows.Next() {
		err := rows.Scan(&activity.Id, &activity.Title, &activity.Name, &activity.Remark, &activity.Online, &activity.StartTime,
			&activity.EndTime, &activity.Description, &activity.EntrancePicture, &activity.InternalTopPicture, &activity.Del,
			&activity.CreateTime, &activity.EndTime)
		if err != nil {
			return nil, err
		}
	}
	return &activity, nil
}

func QueryActivityByName(name string) (*model.Activity, error) {
	sql := "SELECT " + activityColumnList + " FROM wxapp_mall_activity WHERE is_del = 0 AND name = '" + name + "'"
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	activity := model.Activity{}
	if rows.Next() {
		err := rows.Scan(&activity.Id, &activity.Title, &activity.Name, &activity.Remark, &activity.Online, &activity.StartTime,
			&activity.EndTime, &activity.Description, &activity.EntrancePicture, &activity.InternalTopPicture, &activity.Del,
			&activity.CreateTime, &activity.EndTime)
		if err != nil {
			return nil, err
		}
	}
	return &activity, nil
}

func UpdateActivityById(activity *model.Activity) error {
	sql := `
UPDATE wxapp_mall_activity 
SET title = ?, name = ?, remark = ?, online = ?, start_time = ?, end_time = ?, description = ?, entrance_picture = ?, 
internal_top_picture = ?, is_del = ?, update_time = ? WHERE id = ?
`
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(activity.Title, activity.Name, activity.Remark, activity.Online, activity.StartTime,
		activity.EndTime, activity.Description, activity.EntrancePicture, activity.InternalTopPicture,
		activity.Del, time.Now(), activity.Id)
	if err != nil {
		return err
	}
	return nil
}
