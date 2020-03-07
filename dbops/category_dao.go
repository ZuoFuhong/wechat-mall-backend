package dbops

import (
	"strconv"
	"time"
	"wechat-mall-backend/model"
)

const categoryColumnList = "id, parent_id, name, sort, online, picture, description, is_del, create_time, update_time"

func QueryCategoryList(page, size int) (*[]model.Category, error) {
	sql := "SELECT " + categoryColumnList + " FROM wxapp_mall_spu_category WHERE is_del = 0 LIMIT ?, ?"
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query((page-1)*size, size)
	if err != nil {
		return nil, err
	}
	var cateList []model.Category
	for rows.Next() {
		category := model.Category{}
		err := rows.Scan(&category.Id, &category.ParentId, &category.Name, &category.Sort, &category.Online, &category.Picture, &category.Description, &category.Del, &category.CreateTime, &category.UpdateTime)
		if err != nil {
			return nil, err
		}
		cateList = append(cateList, category)
	}
	return &cateList, nil
}

func CountCategory() (int, error) {
	sql := "SELECT COUNT(*) FROM wxapp_mall_spu_category WHERE is_del = 0"
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

func QueryCategoryById(id int) (*model.Category, error) {
	sql := "SELECT " + categoryColumnList + " FROM wxapp_mall_spu_category WHERE is_del = 0 AND id = " + strconv.Itoa(id)
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	category := model.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.ParentId, &category.Name, &category.Sort, &category.Online, &category.Picture, &category.Description, &category.Del, &category.CreateTime, &category.UpdateTime)
		if err != nil {
			return nil, err
		}
	}
	return &category, nil
}

func QueryCategoryByName(name string) (*model.Category, error) {
	sql := "SELECT " + categoryColumnList + " FROM wxapp_mall_spu_category WHERE is_del = 0 AND name = " + name
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	category := model.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.ParentId, &category.Name, &category.Sort, &category.Online, &category.Picture, &category.Description, &category.Del, &category.CreateTime, &category.UpdateTime)
		if err != nil {
			return nil, err
		}
	}
	return &category, nil
}

func InsertCategory(category *model.Category) error {
	sql := "INSERT INTO wxapp_mall_spu_category ( " + categoryColumnList[4:] + " ) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)"
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(category.ParentId, category.Name, category.Sort, category.Online, category.Picture, category.Description, 0, time.Now(), time.Now())
	if err != nil {
		return err
	}
	return nil
}

func UpdateCategoryById(category *model.Category) error {
	sql := `
UPDATE wxapp_mall_spu_category
SET parent_id = ?, name = ?, sort = ?, online = ?, picture = ?, description = ?, is_del = ?, update_time = ? 
WHERE id = ?
`
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(category.ParentId, category.Name, category.Sort, category.Online, category.Picture, category.Description, category.Del, time.Now(), category.Id)
	if err != nil {
		return err
	}
	return nil
}
