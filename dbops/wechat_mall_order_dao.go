package dbops

import (
	"strconv"
	"time"
	"wechat-mall-backend/defs"
	"wechat-mall-backend/model"
)

const orderColumnList = `
id, order_no, user_id, pay_amount, goods_amount, discount_amount, dispatch_amount, pay_time, status, address_id,
address_snapshot, wxapp_prepay_id, is_del, create_time, update_time
`

func QueryOrderByOrderNo(orderNo string) (*model.WechatMallOrderDO, error) {
	sql := "SELECT " + orderColumnList + " FROM wechat_mall_order WHERE order_no = '" + orderNo + "'"
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	order := model.WechatMallOrderDO{}
	for rows.Next() {
		err := rows.Scan(&order.Id, &order.OrderNo, &order.UserId, &order.PayAmount, &order.GoodsAmount,
			&order.DiscountAmount, &order.DispatchAmount, &order.PayTime, &order.Status, &order.AddressId,
			&order.AddressSnapshot, &order.WxappPrePayId, &order.Del, &order.CreateTime, &order.UpdateTime)
		if err != nil {
			return nil, err
		}
	}
	return &order, nil
}

func QueryOrderById(id int) (*model.WechatMallOrderDO, error) {
	sql := "SELECT " + orderColumnList + " FROM wechat_mall_order WHERE id = " + strconv.Itoa(id)
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	order := model.WechatMallOrderDO{}
	for rows.Next() {
		err := rows.Scan(&order.Id, &order.OrderNo, &order.UserId, &order.PayAmount, &order.GoodsAmount,
			&order.DiscountAmount, &order.DispatchAmount, &order.PayTime, &order.Status, &order.AddressId,
			&order.AddressSnapshot, &order.WxappPrePayId, &order.Del, &order.CreateTime, &order.UpdateTime)
		if err != nil {
			return nil, err
		}
	}
	return &order, nil
}

func ListOrderByParams(userId, status, page, size int) (*[]model.WechatMallOrderDO, error) {
	sql := "SELECT " + orderColumnList + " FROM wechat_mall_order WHERE is_del = 0 AND user_id = " + strconv.Itoa(userId)
	if status != 999 {
		sql += " AND status = " + strconv.Itoa(status)
	}
	if page > 0 && size > 0 {
		sql += " LIMIT " + strconv.Itoa((page-1)*page) + " , " + strconv.Itoa(size)
	}
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	orderList := []model.WechatMallOrderDO{}
	for rows.Next() {
		order := model.WechatMallOrderDO{}
		err := rows.Scan(&order.Id, &order.OrderNo, &order.UserId, &order.PayAmount, &order.GoodsAmount,
			&order.DiscountAmount, &order.DispatchAmount, &order.PayTime, &order.Status, &order.AddressId,
			&order.AddressSnapshot, &order.WxappPrePayId, &order.Del, &order.CreateTime, &order.UpdateTime)
		if err != nil {
			return nil, err
		}
		orderList = append(orderList, order)
	}
	return &orderList, nil
}

func CountOrderByParams(userId, status int) (int, error) {
	sql := "SELECT COUNT(*) FROM wechat_mall_order WHERE is_del = 0 AND user_id = " + strconv.Itoa(userId)
	if status != 999 {
		sql += " AND status = " + strconv.Itoa(status)
	}
	rows, err := dbConn.Query(sql)
	if err != nil {
		return 0, err
	}
	total := 0
	for rows.Next() {
		err := rows.Scan(&total)
		if err != nil {
			return 0, err
		}
	}
	return total, nil
}

func AddOrder(order *model.WechatMallOrderDO) error {
	sql := "INSERT INTO wechat_mall_order (" + orderColumnList[4:] + ") VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.OrderNo, order.UserId, order.PayAmount, &order.GoodsAmount, &order.DiscountAmount,
		&order.DispatchAmount, &order.PayTime, &order.Status, &order.AddressId, &order.AddressSnapshot,
		&order.WxappPrePayId, 0, time.Now(), time.Now())
	return err
}

func UpdateOrderById(order *model.WechatMallOrderDO) error {
	sql := `
UPDATE wechat_mall_order
SET order_no = ?, user_id = ?, pay_amount = ?, discount_amount = ?, dispatch_amount = ?, pay_time = ?, status = ?, 
address_id = ?, address_snapshot = ?, wxapp_prepay_id = ?, is_del = ?, update_time = ?
WHERE id = ?
`
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.OrderNo, order.UserId, order.PayAmount, order.DiscountAmount, order.DispatchAmount,
		order.PayTime, order.Status, order.AddressId, order.AddressSnapshot, order.WxappPrePayId, order.Del,
		time.Now(), order.Id)
	return err
}

func QueryOrderSaleData(page, size int) (*[]defs.OrderSaleData, error) {
	sql := `
SELECT 
   DATE_FORMAT(create_time, '%Y-%m-%d') AS createTime, 
   COUNT(id) AS orderNum, 
   IFNULL( SUM(pay_amount), 0) AS saleAmount
FROM wechat_mall_order 
WHERE status IN (1, 2, 3)
GROUP BY createTime
ORDER BY createTime DESC
`
	if page > 0 && size > 0 {
		sql += " LIMIT " + strconv.Itoa((page-1)*size) + ", " + strconv.Itoa(size)
	}
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	saleDataList := []defs.OrderSaleData{}
	for rows.Next() {
		saleData := defs.OrderSaleData{}
		err := rows.Scan(&saleData.Time, &saleData.OrderNum, &saleData.SaleAmount)
		if err != nil {
			return nil, err
		}
		saleDataList = append(saleDataList, saleData)
	}
	return &saleDataList, nil
}

func CountOrderNum(status int) (int, error) {
	sql := "SELECT COUNT(*) FROM wechat_mall_order WHERE status = " + strconv.Itoa(status)
	rows, err := dbConn.Query(sql)
	if err != nil {
		return 0, err
	}
	total := 0
	for rows.Next() {
		err := rows.Scan(&total)
		if err != nil {
			return 0, err
		}
	}
	return total, nil
}
