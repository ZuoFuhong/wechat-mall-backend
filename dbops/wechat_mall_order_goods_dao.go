package dbops

import (
	"strconv"
	"time"
	"wechat-mall-backend/model"
)

const orderGoodsColumnList = `
id, order_no, goods_id, sku_id, picture, title, price, specs, num, lock_status, create_time, update_time
`

func QueryOrderGoods(orderNo string) (*[]model.WechatMallOrderGoodsDO, error) {
	sql := "SELECT " + orderColumnList + " FROM wechat_mall_order_goods WHERE order_no = '" + orderNo + "'"
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	goodsList := []model.WechatMallOrderGoodsDO{}
	for rows.Next() {
		goods := model.WechatMallOrderGoodsDO{}
		err := rows.Scan(&goods.Id, &goods.OrderNo, &goods.GoodsId, &goods.SkuId, &goods.Picture, &goods.Title,
			&goods.Price, &goods.Specs, &goods.Num, &goods.LockStatus, &goods.CreateTime, &goods.UpdateTime)
		if err != nil {
			return nil, err
		}
		goodsList = append(goodsList, goods)
	}
	return &goodsList, nil
}

func AddOrderGoods(goods *model.WechatMallOrderGoodsDO) error {
	sql := "INSERT INTO wechat_mall_order_goods (" + orderGoodsColumnList[4:] + " ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(goods.OrderNo, goods.GoodsId, goods.SkuId, goods.Picture, goods.Title, goods.Price, goods.Specs,
		goods.Num, 0, time.Now(), time.Now())
	return err
}

func SumGoodsSaleNum(goodsId, skuId int) (int, error) {
	sql := "SELECT IFNULL(SUM(num), 0) FROM wechat_mall_order_goods WHERE lock_status = 1"
	if goodsId != 0 {
		sql += " AND goods_id = " + strconv.Itoa(goodsId)
	}
	if skuId != 0 {
		sql += " AND sku_id = " + strconv.Itoa(skuId)
	}
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
	return total, err
}
