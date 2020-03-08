package dbops

import (
	"strconv"
	"time"
	"wechat-mall-backend/model"
)

const spuSpecColumnList = `
id, spu_id, spec_id, is_del, create_time, update_time
`

func GetSPUSpecList(spuId int) (*[]model.SPUSpec, error) {
	sql := "SELECT " + spuSpecColumnList + " FROM wxapp_mall_spu_spec WHERE is_del = 0 AND spu_id = " + strconv.Itoa(spuId)
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	var spuSpecList []model.SPUSpec
	for rows.Next() {
		spuSpec := model.SPUSpec{}
		err := rows.Scan(&spuSpec.Id, &spuSpec.SpuId, &spuSpec.SpecId, &spuSpec.Del, &spuSpec.CreateTime, &spuSpec.UpdateTime)
		if err != nil {
			return nil, err
		}
		spuSpecList = append(spuSpecList, spuSpec)
	}
	return &spuSpecList, nil
}

func DeleteSPUSpec(spuId int) error {
	sql := "UPDATE wxapp_mall_spu_spec SET is_del = 1 WHERE spu_id = " + strconv.Itoa(spuId)
	_, err := dbConn.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

func InsertSPUSpec(spuSpec *model.SPUSpec) error {
	sql := "INSERT INTO wxapp_mall_spu_spec(" + spuSpecColumnList[4:] + ") VALUES(?, ?, ?, ?, ?)"
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(spuSpec.SpuId, spuSpec.SpecId, 0, time.Now(), time.Now())
	if err != nil {
		return err
	}
	return nil
}
