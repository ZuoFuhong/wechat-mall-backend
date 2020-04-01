package dbops

import (
	"strconv"
	"time"
	"wechat-mall-backend/defs"
	"wechat-mall-backend/model"
)

const couponLogColumnList = `
id, coupon_id, user_id, use_time, expire_time, status, code, order_no, is_del, create_time, update_time
`

func QueryCouponLogById(couponLogId int) (*model.WechatMallCouponLogDO, error) {
	sql := "SELECT " + couponLogColumnList + " FROM wechat_mall_coupon_log WHERE id = " + strconv.Itoa(couponLogId)
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	couponLog := model.WechatMallCouponLogDO{}
	for rows.Next() {
		err := rows.Scan(&couponLog.Id, &couponLog.CouponId, &couponLog.UserId, &couponLog.UseTime,
			&couponLog.ExpireTime, &couponLog.Status, &couponLog.Code, &couponLog.OrderNo, &couponLog.Del,
			&couponLog.CreateTime, &couponLog.UpdateTime)
		if err != nil {
			return nil, err
		}
	}
	return &couponLog, nil
}

func QueryCouponLogList(userId, status, page, size int) (*[]model.WechatMallCouponLogDO, error) {
	sql := "SELECT " + couponLogColumnList + " FROM wechat_mall_coupon_log WHERE is_del = 0"
	sql += " AND user_id = " + strconv.Itoa(userId)
	sql += " AND status = " + strconv.Itoa(status)
	if page > 0 && size > 0 {
		sql += " LIMIT " + strconv.Itoa((page-1)*page) + " , " + strconv.Itoa(size)
	}
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	couponLogList := []model.WechatMallCouponLogDO{}
	for rows.Next() {
		couponLog := model.WechatMallCouponLogDO{}
		err := rows.Scan(&couponLog.Id, &couponLog.CouponId, &couponLog.UserId, &couponLog.UseTime,
			&couponLog.ExpireTime, &couponLog.Status, &couponLog.Code, &couponLog.OrderNo, &couponLog.Del,
			&couponLog.CreateTime, &couponLog.UpdateTime)
		if err != nil {
			return nil, err
		}
		couponLogList = append(couponLogList, couponLog)
	}
	return &couponLogList, nil
}

func CountCouponTakeNum(userId, couponId, status int, del int) (int, error) {
	sql := "SELECT COUNT(*) FROM wechat_mall_coupon_log WHERE 1 = 1"
	if userId != defs.ALL {
		sql += " AND user_id = " + strconv.Itoa(userId)
	}
	if couponId != defs.ALL {
		sql += " AND coupon_id = " + strconv.Itoa(couponId)
	}
	if status != defs.ALL {
		sql += " AND status = " + strconv.Itoa(status)
	}
	if del != defs.ALL {
		sql += " AND is_del = " + strconv.Itoa(del)
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

func UpdateCouponLogById(couponLog *model.WechatMallCouponLogDO) error {
	sql := `
UPDATE wechat_mall_coupon_log 
SET coupon_id = ?, user_id = ?, use_time = ?, expire_time = ?, status = ?, code = ?, order_no = ?, 
is_del = ?, update_time = ?
WHERE id = ?
`
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(couponLog.CouponId, couponLog.UserId, couponLog.UseTime, couponLog.ExpireTime, couponLog.Status,
		couponLog.Code, couponLog.OrderNo, couponLog.Del, time.Now(), couponLog.Id)
	return err
}

func AddCouponLog(couponLog *model.WechatMallCouponLogDO) error {
	sql := "INSERT INTO wechat_mall_coupon_log (" + couponLogColumnList[4:] + ") VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(couponLog.CouponId, couponLog.UserId, couponLog.UseTime, couponLog.ExpireTime, couponLog.Status,
		couponLog.Code, "", 0, time.Now(), time.Now())
	return err
}
