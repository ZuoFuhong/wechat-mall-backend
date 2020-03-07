package dbops

import (
	"strconv"
	"time"
	"wechat-mall-backend/model"
)

const specAttrColumnList = `
id, spec_id, value, extend, is_del, create_time, update_time
`

func QuerySpecificationAttrList(page, size int) (*[]model.SpecificationAttr, error) {
	sql := "SELECT " + specAttrColumnList + " FROM wxapp_mall_specification_attr WHERE is_del = 0 LIMIT ?, ?"
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query((page-1)*size, size)
	if err != nil {
		return nil, err
	}
	var attrList []model.SpecificationAttr
	for rows.Next() {
		spec := model.SpecificationAttr{}
		err := rows.Scan(&spec.Id, &spec.SpecId, &spec.Value, &spec.Extend, &spec.Del, &spec.CreateTime, &spec.UpdateTime)
		if err != nil {
			return nil, err
		}
		attrList = append(attrList, spec)
	}
	return &attrList, nil
}

func CountSpecificationAttr() (int, error) {
	sql := "SELECT COUNT(*) FROM wxapp_mall_specification_attr WHERE is_del = 0"
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

func AddSpecificationAttr(spec *model.SpecificationAttr) error {
	sql := "INSERT INTO wxapp_mall_specification_attr ( " + specAttrColumnList[4:] + " ) VALUES(?, ?, ?, ?, ?, ?)"
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(sql, spec.SpecId, spec.Value, spec.Extend, 0, time.Now(), time.Now())
	if err != nil {
		return err
	}
	return nil
}

func QuerySpecificationAttrById(id int) (*model.SpecificationAttr, error) {
	sql := "SELECT " + specAttrColumnList + " FROM wxapp_mall_specification_attr WHERE is_del = 0 AND id = " + strconv.Itoa(id)
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	attr := model.SpecificationAttr{}
	if rows.Next() {
		err := rows.Scan(&attr.Id, &attr.SpecId, &attr.Value, &attr.Extend, &attr.Del, &attr.CreateTime, &attr.UpdateTime)
		if err != nil {
			return nil, err
		}
	}
	return &attr, nil
}

func QuerySpecificationAttrByValue(name string) (*model.SpecificationAttr, error) {
	sql := "SELECT " + specAttrColumnList + " FROM wxapp_mall_specification_attr WHERE is_del = 0 AND value = " + name
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	attr := model.SpecificationAttr{}
	if rows.Next() {
		err := rows.Scan(&attr.Id, &attr.SpecId, &attr.Value, &attr.Extend, &attr.Del, &attr.CreateTime, &attr.UpdateTime)
		if err != nil {
			return nil, err
		}
	}
	return &attr, nil
}

func UpdateSpecificationAttrById(attr *model.SpecificationAttr) error {
	sql := `
UPDATE wxapp_mall_specification_attr 
SET spec_id = ?, value = ?, extend = ?, is_del = ?, update_time = ? 
WHERE id = ?
`
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(attr.SpecId, attr.Value, attr.Extend, attr.Del, time.Now())
	if err != nil {
		return err
	}
	return nil
}
