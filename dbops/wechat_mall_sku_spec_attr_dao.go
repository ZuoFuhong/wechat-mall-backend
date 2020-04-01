package dbops

import (
	"strconv"
	"time"
	"wechat-mall-backend/model"
)

const skuSpecAttrColumnList = `
id, sku_id, spec_id, attr_id, is_del, create_time, update_time
`

func InsertSkuSpecAttr(skuSpecAttrDO *model.WechatMallSkuSpecAttrDO) error {
	sql := "INSERT INTO wechat_mall_sku_spec_attr (" + skuSpecAttrColumnList[4:] + ") VALUES (?, ?, ?, ?, ?, ?)"
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(skuSpecAttrDO.SkuId, skuSpecAttrDO.SpecId, skuSpecAttrDO.AttrId, 0, time.Now(), time.Now())
	return err
}

func RemoveRelatedBySkuId(skuId int) error {
	sql := "UPDATE wechat_mall_sku_spec_attr SET update_time = now(), is_del = 1 WHERE sku_id = " + strconv.Itoa(skuId)
	_, err := dbConn.Exec(sql)
	return err
}

func CountRelatedByAttrId(attrId int) (int, error) {
	sql := "SELECT COUNT(*) FROM wechat_mall_sku_spec_attr WHERE is_del = 0 AND attr_id = " + strconv.Itoa(attrId)
	rows, err := dbConn.Query(sql)
	if err != nil {
		return 0, err
	}
	total := 0
	for rows.Next() {
		err := rows.Scan(&total)
		if err != nil {
			panic(err)
		}
	}
	return total, nil
}
