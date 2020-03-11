package dbops

import (
	"time"
	"wechat-mall-backend/model"
)

const orderGoodsColumnList = `
id, order_no, goods_id, sku_id, picture, title, price, specs, num, create_time, update_time
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
			&goods.Price, &goods.Specs, &goods.Num, &goods.CreateTime, &goods.UpdateTime)
		if err != nil {
			return nil, err
		}
		goodsList = append(goodsList, goods)
	}
	return &goodsList, nil
}

func AddOrderGoods(goods *model.WechatMallOrderGoodsDO) error {
	sql := "INSERT INTO wechat_mall_order_goods (" + orderGoodsColumnList[4:] + " ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(goods.OrderNo, goods.GoodsId, goods.SkuId, goods.Picture, goods.Title, goods.Price, goods.Specs,
		goods.Num, time.Now(), time.Now())
	return err
}
