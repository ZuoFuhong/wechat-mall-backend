package dbops

import (
	"strconv"
	"time"
	"wechat-mall-backend/defs"
	"wechat-mall-backend/model"
)

const orderColumnList = `
id, order_no, user_id, pay_amount, goods_amount, discount_amount, dispatch_amount, pay_time, deliver_time,
finish_time, status, address_id, address_snapshot, wxapp_prepay_id, transaction_id, remark, is_del, create_time, update_time
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
			&order.DiscountAmount, &order.DispatchAmount, &order.PayTime, &order.DeliverTime, &order.FinishTime,
			&order.Status, &order.AddressId, &order.AddressSnapshot, &order.WxappPrepayId, &order.TransactionId,
			&order.Remark, &order.Del, &order.CreateTime, &order.UpdateTime)
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
			&order.DiscountAmount, &order.DispatchAmount, &order.PayTime, &order.DeliverTime, &order.FinishTime,
			&order.Status, &order.AddressId, &order.AddressSnapshot, &order.WxappPrepayId, &order.TransactionId,
			&order.Remark, &order.Del, &order.CreateTime, &order.UpdateTime)
		if err != nil {
			return nil, err
		}
	}
	return &order, nil
}

func ListOrderByParams(userId, status, page, size int) (*[]model.WechatMallOrderDO, error) {
	sql := "SELECT " + orderColumnList + " FROM wechat_mall_order WHERE is_del = 0 AND user_id = " + strconv.Itoa(userId)
	if status != defs.ALL {
		sql += " AND status = " + strconv.Itoa(status)
	}
	if page > 0 && size > 0 {
		sql += " ORDER BY create_time DESC LIMIT " + strconv.Itoa((page-1)*page) + " , " + strconv.Itoa(size)
	}
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	orderList := []model.WechatMallOrderDO{}
	for rows.Next() {
		order := model.WechatMallOrderDO{}
		err := rows.Scan(&order.Id, &order.OrderNo, &order.UserId, &order.PayAmount, &order.GoodsAmount,
			&order.DiscountAmount, &order.DispatchAmount, &order.PayTime, &order.DeliverTime, &order.FinishTime,
			&order.Status, &order.AddressId, &order.AddressSnapshot, &order.WxappPrepayId, &order.TransactionId,
			&order.Remark, &order.Del, &order.CreateTime, &order.UpdateTime)
		if err != nil {
			return nil, err
		}
		orderList = append(orderList, order)
	}
	return &orderList, nil
}

func CountOrderByParams(userId, status int) (int, error) {
	sql := "SELECT COUNT(*) FROM wechat_mall_order WHERE is_del = 0 AND user_id = " + strconv.Itoa(userId)
	if status != defs.ALL {
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
	sql := "INSERT INTO wechat_mall_order (" + orderColumnList[4:] + ") VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.OrderNo, order.UserId, order.PayAmount, &order.GoodsAmount, &order.DiscountAmount,
		&order.DispatchAmount, &order.PayTime, &order.DeliverTime, &order.FinishTime, &order.Status, &order.AddressId,
		&order.AddressSnapshot, &order.WxappPrepayId, &order.TransactionId, &order.Remark, 0, time.Now(), time.Now())
	return err
}

func UpdateOrderById(order *model.WechatMallOrderDO) error {
	sql := "UPDATE wechat_mall_order SET update_time = now() "
	if order.PayAmount != "" {
		sql += ", pay_amount = '" + order.PayAmount + "'"
	}
	if order.GoodsAmount != "" {
		sql += ", goods_amount = '" + order.GoodsAmount + "'"
	}
	if order.DiscountAmount != "" {
		sql += ", discount_amount = '" + order.DiscountAmount + "'"
	}
	if order.DispatchAmount != "" {
		sql += ", dispatch_amount = '" + order.DispatchAmount + "'"
	}
	if order.PayTime != "" {
		sql += ", pay_time = '" + order.PayTime + "'"
	}
	if order.DeliverTime != "" {
		sql += ", deliver_time = '" + order.DeliverTime + "'"
	}
	if order.FinishTime != "" {
		sql += ", finish_time = '" + order.FinishTime + "'"
	}
	if order.Status != 0 {
		sql += ", status = " + strconv.Itoa(order.Status)
	}
	if order.AddressSnapshot != "" {
		sql += ", address_snapshot = '" + order.AddressSnapshot + "'"
	}
	if order.TransactionId != "" {
		sql += ", transaction_id = '" + order.TransactionId + "'"
	}
	if order.Del != 0 {
		sql += ", is_del = " + strconv.Itoa(order.Del)
	}
	sql += " WHERE id = " + strconv.Itoa(order.Id)
	_, err := dbConn.Exec(sql)
	return err
}

func UpdateOrderRemark(id int, remark string) error {
	sql := "UPDATE wechat_mall_order SET update_time = now(), remark = ? WHERE id = ?"
	stmt, e := dbConn.Prepare(sql)
	if e != nil {
		return e
	}
	_, e = stmt.Exec(remark, id)
	return e
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

func CountOrderNum(userId, status int) (int, error) {
	sql := "SELECT COUNT(*) FROM wechat_mall_order WHERE is_del = 0 AND status = " + strconv.Itoa(status)
	if userId != defs.ALL {
		sql += " AND user_id = " + strconv.Itoa(userId)
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

func SelectCMSOrderList(status, searchType int, keyword, startTime, endTime string, page, size int) (*[]model.WechatMallOrderDO, error) {
	sql := "SELECT " + orderColumnList + " FROM wechat_mall_order WHERE 1 = 1"
	if status != defs.ALL {
		sql += " AND status = " + strconv.Itoa(status)
	}
	if searchType != defs.ALL && keyword != "" {
		switch searchType {
		case 1:
			// 订单号
			sql += " AND order_no LIKE '%" + keyword + "%'"
		case 2, 3:
			// 买家信息
			sql += " AND address_snapshot LIKE '%" + keyword + "%'"
		case 4:
			// 微信支付单号
			sql += " AND transaction_id LIKE '%" + keyword + "%'"
		case 5:
			// 商品名称
			sql += " AND order_no IN (SELECT order_no FROM wechat_mall_order_goods WHERE title LIKE '%" + keyword + "%')"
		default:
		}
	}
	if startTime != "" && endTime != "" {
		sql += " AND create_time BETWEEN '" + startTime + "' AND '" + endTime + "'"
	}
	if page != 0 && size != 0 {
		sql += " LIMIT " + strconv.Itoa((page-1)*size) + ", " + strconv.Itoa(size)
	}
	rows, e := dbConn.Query(sql)
	if e != nil {
		panic(e)
	}
	orderList := []model.WechatMallOrderDO{}
	for rows.Next() {
		orderDO := model.WechatMallOrderDO{}
		err := rows.Scan(&orderDO.Id, &orderDO.OrderNo, &orderDO.UserId, &orderDO.PayAmount, &orderDO.GoodsAmount,
			&orderDO.DiscountAmount, &orderDO.DispatchAmount, &orderDO.PayTime, &orderDO.DeliverTime, &orderDO.FinishTime,
			&orderDO.Status, &orderDO.AddressId, &orderDO.AddressSnapshot, &orderDO.WxappPrepayId, &orderDO.TransactionId,
			&orderDO.Remark, &orderDO.Del, &orderDO.CreateTime, &orderDO.UpdateTime)
		if err != nil {
			return nil, err
		}
		orderList = append(orderList, orderDO)
	}
	return &orderList, nil
}

func SelectCMSOrderNum(status, searchType int, keyword, startTime, endTime string) (int, error) {
	sql := "SELECT COUNT(*) FROM wechat_mall_order WHERE 1 = 1"
	if status != defs.ALL {
		sql += " AND status = " + strconv.Itoa(status)
	}
	if searchType != defs.ALL && keyword != "" {
		switch searchType {
		case 1:
			// 订单号
			sql += " AND order_no LIKE '%" + keyword + "%'"
		case 2, 3:
			// 买家信息
			sql += " AND address_snapshot LIKE '%" + keyword + "%'"
		case 4:
			// 微信支付单号
			sql += " AND transaction_id LIKE '%" + keyword + "%'"
		case 5:
			// 商品名称
			sql += " AND order_no IN (SELECT order_no FROM wechat_mall_order_goods WHERE title LIKE '%" + keyword + "%')"
		default:
		}
	}
	if startTime != "" && endTime != "" {
		sql += " AND create_time BETWEEN '" + startTime + "' AND '" + endTime + "'"
	}
	rows, e := dbConn.Query(sql)
	if e != nil {
		panic(e)
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
