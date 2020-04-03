package dbops

import (
	"strconv"
	"time"
	"wechat-mall-backend/model"
)

const orderGoodsColumnList = `
id, order_no, user_id, goods_id, sku_id, picture, title, price, specs, num, lock_status, create_time, update_time
`

func QueryOrderGoods(orderNo string) (*[]model.WechatMallOrderGoodsDO, error) {
	sql := "SELECT " + orderGoodsColumnList + " FROM wechat_mall_order_goods WHERE order_no = '" + orderNo + "'"
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	goodsList := []model.WechatMallOrderGoodsDO{}
	for rows.Next() {
		goods := model.WechatMallOrderGoodsDO{}
		err := rows.Scan(&goods.Id, &goods.OrderNo, &goods.UserId, &goods.GoodsId, &goods.SkuId, &goods.Picture, &goods.Title,
			&goods.Price, &goods.Specs, &goods.Num, &goods.LockStatus, &goods.CreateTime, &goods.UpdateTime)
		if err != nil {
			return nil, err
		}
		goodsList = append(goodsList, goods)
	}
	return &goodsList, nil
}

func AddOrderGoods(goods *model.WechatMallOrderGoodsDO) error {
	sql := "INSERT INTO wechat_mall_order_goods (" + orderGoodsColumnList[4:] + " ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(goods.OrderNo, goods.UserId, goods.GoodsId, goods.SkuId, goods.Picture, goods.Title, goods.Price, goods.Specs,
		goods.Num, 0, time.Now(), time.Now())
	return err
}

func UpdateOrderGoodsLockStatus(id, status int) error {
	sql := "UPDATE SET lock_status = " + strconv.Itoa(status) + " WHERE id = " + strconv.Itoa(id)
	_, err := dbConn.Exec(sql)
	return err
}

// 商品-统计购买人数
func CountBuyGoodsUserNum(goodsId int) (int, error) {
	sql := "SELECT COUNT(DISTINCT(user_id)) FROM wechat_mall_order_goods WHERE lock_status = 1 AND goods_id = " + strconv.Itoa(goodsId)
	rows, err := dbConn.Query(sql)
	if err != nil {
		return 0, err
	}
	humanNum := 0
	for rows.Next() {
		err := rows.Scan(&humanNum)
		if err != nil {
			return 0, err
		}
	}
	return humanNum, nil
}
