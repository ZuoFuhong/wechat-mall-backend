package dbops

import (
	"strconv"
	"time"
	"wechat-mall-backend/model"
)

const gridCategoryColumnList = `
id, title, name, category_id, picture, is_del, create_time, update_time
`

func QueryGridCategoryList(page, size int) (*[]model.WechatMallGridCategoryDO, error) {
	sql := "SELECT " + gridCategoryColumnList + " FROM wechat_mall_grid_category WHERE is_del = 0"
	if page > 0 && size > 0 {
		sql += " LIMIT " + strconv.Itoa((page-1)*size) + " ," + strconv.Itoa(size)
	}
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	var gridCList []model.WechatMallGridCategoryDO
	for rows.Next() {
		gridC := model.WechatMallGridCategoryDO{}
		err := rows.Scan(&gridC.Id, &gridC.Title, &gridC.Name, &gridC.CategoryId, &gridC.Picture, &gridC.Del, &gridC.CreateTime, &gridC.UpdateTime)
		if err != nil {
			return nil, err
		}
		gridCList = append(gridCList, gridC)
	}
	return &gridCList, nil
}

func CountGridCategory() (int, error) {
	sql := "SELECT COUNT(*) FROM wechat_mall_grid_category WHERE is_del = 0"
	rows, err := dbConn.Query(sql)
	if err != nil {
		return 0, err
	}
	total := 0
	if rows.Next() {
		err := rows.Scan(&total)
		if err != nil {
			return 0, err
		}
	}
	return total, nil
}

func InsertGridCategory(gridC *model.WechatMallGridCategoryDO) error {
	sql := "INSERT INTO wechat_mall_grid_category( " + gridCategoryColumnList[4:] + " ) VALUES(?, ?, ?, ?, ?, ?, ?)"
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(gridC.Title, gridC.Name, gridC.CategoryId, gridC.Picture, 0, time.Now(), time.Now())
	if err != nil {
		return err
	}
	return nil
}

func QueryGridCategoryById(id int) (*model.WechatMallGridCategoryDO, error) {
	sql := "SELECT " + gridCategoryColumnList + " FROM wechat_mall_grid_category WHERE is_del = 0 AND id = " + strconv.Itoa(id)
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	gridC := model.WechatMallGridCategoryDO{}
	if rows.Next() {
		err := rows.Scan(&gridC.Id, &gridC.Title, &gridC.Name, &gridC.CategoryId, &gridC.Picture, &gridC.Del, &gridC.CreateTime, &gridC.UpdateTime)
		if err != nil {
			return nil, err
		}
	}
	return &gridC, nil
}

func QueryGridCategoryByName(name string) (*model.WechatMallGridCategoryDO, error) {
	sql := "SELECT " + gridCategoryColumnList + " FROM wechat_mall_grid_category WHERE is_del = 0 AND name = '" + name + "'"
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	gridC := model.WechatMallGridCategoryDO{}
	if rows.Next() {
		err := rows.Scan(&gridC.Id, &gridC.Title, &gridC.Name, &gridC.CategoryId, &gridC.Picture, &gridC.Del, &gridC.CreateTime, &gridC.UpdateTime)
		if err != nil {
			return nil, err
		}
	}
	return &gridC, nil
}

func UpdateGridCategoryById(gridC *model.WechatMallGridCategoryDO) error {
	sql := `
UPDATE wechat_mall_grid_category 
SET title = ?, name = ?, category_id = ?, picture = ?, is_del = ?, update_time = ? 
WHERE id = ?
`
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(gridC.Title, gridC.Name, gridC.CategoryId, gridC.Picture, gridC.Del, time.Now(), gridC.Id)
	if err != nil {
		return err
	}
	return nil
}
