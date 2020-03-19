package dbops

const refundColumnList = `
id, refund_no, user_id, order_no, reason, refund_amount, status, is_del, create_time, update_time
`

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
