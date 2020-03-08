package dbops

import (
	"strconv"
	"time"
	"wechat-mall-backend/model"
)

const specAttrColumnList = `
id, spec_id, value, extend, is_del, create_time, update_time
`

func QuerySpecificationAttrList(specId int) (*[]model.SpecificationAttr, error) {
	sql := "SELECT " + specAttrColumnList + " FROM wxapp_mall_specification_attr WHERE is_del = 0 AND spec_id = " + strconv.Itoa(specId)
	rows, err := dbConn.Query(sql)
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

func AddSpecificationAttr(spec *model.SpecificationAttr) error {
	sql := "INSERT INTO wxapp_mall_specification_attr ( " + specAttrColumnList[4:] + " ) VALUES(?, ?, ?, ?, ?, ?)"
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(spec.SpecId, spec.Value, spec.Extend, 0, time.Now(), time.Now())
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
	sql := "SELECT " + specAttrColumnList + " FROM wxapp_mall_specification_attr WHERE is_del = 0 AND value = '" + name + "'"
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
	_, err = stmt.Exec(attr.SpecId, attr.Value, attr.Extend, attr.Del, time.Now(), attr.Id)
	if err != nil {
		return err
	}
	return nil
}
