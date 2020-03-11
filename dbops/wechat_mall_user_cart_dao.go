package dbops

import (
	"strconv"
	"time"
	"wechat-mall-backend/model"
)

const cartColumnList = `
id, user_id, goods_id, sku_id, num, is_del, create_time, update_time
`

func QueryCartList(userId, page, size int) (*[]model.WechatMallUserCartDO, error) {
	sql := "SELECT " + cartColumnList + " FROM wechat_mall_user_cart WHERE is_del = 0 AND user_id = " + strconv.Itoa(userId)
	if page > 0 && size > 0 {
		sql += " LIMIT " + strconv.Itoa((page-1)*page) + " , " + strconv.Itoa(size)
	}
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	cartList := []model.WechatMallUserCartDO{}
	for rows.Next() {
		cartDO := model.WechatMallUserCartDO{}
		err := rows.Scan(&cartDO.Id, &cartDO.UserId, &cartDO.GoodsId, &cartDO.SkuId, &cartDO.Num, &cartDO.Del, &cartDO.CreateTime, &cartDO.UpdateTime)
		if err != nil {
			return nil, err
		}
		cartList = append(cartList, cartDO)
	}
	return &cartList, nil
}

func AddUserCart(cartDO *model.WechatMallUserCartDO) error {
	sql := "INSERT INTO wechat_mall_user_cart ( " + cartColumnList[4:] + " ) VALUES (?, ?, ?, ?, ?, ?, ?)"
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(cartDO.UserId, cartDO.GoodsId, cartDO.SkuId, cartDO.Num, 0, time.Now(), time.Now())
	return err
}

func QueryCartByParams(userId, goodsId, skuId int) (*model.WechatMallUserCartDO, error) {
	sql := "SELECT " + cartColumnList + " FROM wechat_mall_user_cart WHERE is_del = 0 AND user_id = ? AND goods_id = ? AND sku_id = ?"
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query(userId, goodsId, skuId)
	if err != nil {
		return nil, err
	}
	cartDO := model.WechatMallUserCartDO{}
	if rows.Next() {
		err := rows.Scan(&cartDO.UserId, &cartDO.GoodsId, &cartDO.Num, &cartDO.Del, &cartDO.CreateTime, &cartDO.UpdateTime)
		if err != nil {
			return nil, err
		}
	}
	return &cartDO, nil
}

func UpdateCartById(cartDO *model.WechatMallUserCartDO) error {
	sql := `
UPDATE wechat_mall_user_cart 
SET user_id = ?, goods_id = ?, num = ?, is_del = ?, update_time = ?
WHERE id = ?
`
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(cartDO.UserId, cartDO.GoodsId, cartDO.Num, cartDO.Del, time.Now(), cartDO.Id)
	return err
}
