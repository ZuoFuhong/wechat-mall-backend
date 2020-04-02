package dbops

import (
	"time"
	"wechat-mall-backend/model"
)

const refundColumnList = `
id, refund_no, user_id, order_no, reason, refund_amount, status, is_del, create_time, update_time
`

func AddRefundRecord(record *model.WechatMallOrderRefund) error {
	sql := "INSERT INTO wechat_mall_order_refund (" + refundColumnList[4:] + ") VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)"
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(record.RefundNo, record.UserId, record.OrderNo, record.Reason, record.RefundAmount, 0, 0, time.Now(), time.Now())
	return err
}

func QueryRefundRecord(refundNo string) (*model.WechatMallOrderRefund, error) {
	sql := "SELECT " + refundColumnList + " FROM wechat_mall_order_refund WHERE refund_no = '" + refundNo + "'"
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	record := model.WechatMallOrderRefund{}
	for rows.Next() {
		err := rows.Scan(&record.RefundNo, &record.UserId, &record.OrderNo, &record.Reason, &record.RefundAmount, &record.Status, &record.Del, &record.CreateTime, &record.UpdateTime)
		if err != nil {
			return nil, err
		}
	}
	return &record, nil
}

func CountPendingOrderRefund() (int, error) {
	sql := "SELECT COUNT(*) FROM wechat_mall_order_refund WHERE status IN (0, 1)"
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
