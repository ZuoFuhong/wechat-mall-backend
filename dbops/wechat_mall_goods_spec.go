package dbops

import (
	"strconv"
	"time"
	"wechat-mall-backend/model"
)

const goodsSpecColumnList = `
id, goods_id, spec_id, is_del, create_time, update_time
`

func GetGoodsSpecList(goodsId int) (*[]model.WechatMallGoodsSpecDO, error) {
	sql := "SELECT " + goodsSpecColumnList + " FROM wechat_mall_goods_spec WHERE is_del = 0 AND goods_id = " + strconv.Itoa(goodsId)
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	var goodsSpecList []model.WechatMallGoodsSpecDO
	for rows.Next() {
		goodsSpec := model.WechatMallGoodsSpecDO{}
		err := rows.Scan(&goodsSpec.Id, &goodsSpec.GoodsId, &goodsSpec.SpecId, &goodsSpec.Del,
			&goodsSpec.CreateTime, &goodsSpec.UpdateTime)
		if err != nil {
			return nil, err
		}
		goodsSpecList = append(goodsSpecList, goodsSpec)
	}
	return &goodsSpecList, nil
}

func DeleteGoodsSpec(goodsId int) error {
	sql := "UPDATE wechat_mall_goods_spec SET is_del = 1 WHERE goods_id = " + strconv.Itoa(goodsId)
	_, err := dbConn.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

func InsertGoodsSpec(goodsSpec *model.WechatMallGoodsSpecDO) error {
	sql := "INSERT INTO wechat_mall_goods_spec(" + goodsSpecColumnList[4:] + ") VALUES(?, ?, ?, ?, ?)"
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(goodsSpec.GoodsId, goodsSpec.SpecId, 0, time.Now(), time.Now())
	if err != nil {
		return err
	}
	return nil
}
