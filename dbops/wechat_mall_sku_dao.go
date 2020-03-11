package dbops

import (
	"strconv"
	"time"
	"wechat-mall-backend/model"
)

const skuColumnList = `
id, title, price, code, stock, goods_id, online, picture, specs, is_del, create_time, update_time
`

func GetSKUList(page, size int) (*[]model.WechatMallSkuDO, error) {
	sql := "SELECT " + skuColumnList + " FROM wechat_mall_sku WHERE is_del = 0 LIMIT ?, ?"
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query((page-1)*size, size)
	if err != nil {
		return nil, err
	}
	var skuList []model.WechatMallSkuDO
	for rows.Next() {
		sku := model.WechatMallSkuDO{}
		err := rows.Scan(&sku.Id, &sku.Title, &sku.Price, &sku.Code, &sku.Stock, &sku.GoodsId, &sku.Online, &sku.Picture,
			&sku.Specs, &sku.Del, &sku.CreateTime, &sku.UpdateTime)
		if err != nil {
			return nil, err
		}
		skuList = append(skuList, sku)
	}
	return &skuList, nil
}

func CountSKU() (int, error) {
	sql := "SELECT COUNT(*) FROM wechat_mall_sku WHERE is_del = 0"
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

func AddSKU(sku *model.WechatMallSkuDO) error {
	sql := "INSERT INTO wechat_mall_sku( " + skuColumnList[4:] + " ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(sku.Title, sku.Price, sku.Code, sku.Stock, sku.GoodsId, sku.Online, sku.Picture, sku.Specs, 0, time.Now(), time.Now())
	if err != nil {
		return err
	}
	return nil
}

func GetSKUById(id int) (*model.WechatMallSkuDO, error) {
	sql := "SELECT " + skuColumnList + " FROM wechat_mall_sku WHERE is_del = 0 AND id = " + strconv.Itoa(id)
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	sku := model.WechatMallSkuDO{}
	if rows.Next() {
		err := rows.Scan(&sku.Id, &sku.Title, &sku.Price, &sku.Code, &sku.Stock, &sku.GoodsId, &sku.Online, &sku.Picture,
			&sku.Specs, &sku.Del, &sku.CreateTime, &sku.UpdateTime)
		if err != nil {
			return nil, err
		}
	}
	return &sku, nil
}

func GetSKUByCode(code string) (*model.WechatMallSkuDO, error) {
	sql := "SELECT " + skuColumnList + " FROM wechat_mall_sku WHERE is_del = 0 AND code = '" + code + "'"
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	sku := model.WechatMallSkuDO{}
	if rows.Next() {
		err := rows.Scan(&sku.Id, &sku.Title, &sku.Price, &sku.Code, &sku.Stock, &sku.GoodsId, &sku.Online, &sku.Picture,
			&sku.Specs, &sku.Del, &sku.CreateTime, &sku.UpdateTime)
		if err != nil {
			return nil, err
		}
	}
	return &sku, nil
}

func UpdateSKUById(sku *model.WechatMallSkuDO) error {
	sql := `
UPDATE wechat_mall_sku
SET title = ?, price = ?, code = ?, stock = ?, goods_id = ?, online = ?, picture = ?, specs = ?, is_del = ?, update_time = ?
WHERE id = ?
`
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(sku.Title, sku.Price, sku.Code, sku.Stock, sku.GoodsId, sku.Online, sku.Picture, sku.Specs, sku.Del, time.Now(), sku.Id)
	if err != nil {
		return err
	}
	return nil
}
