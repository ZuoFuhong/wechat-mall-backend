package dbops

import (
	"strconv"
	"time"
	"wechat-mall-backend/model"
)

const couponLogColumnList = `
id, coupon_id, user_id, use_time, expire_time, status, code, order_no, is_del, create_time, update_time
`

func QueryCouponLog(userId, couponId int) (*model.WechatMallCouponLogDO, error) {
	sql := "SELECT " + couponLogColumnList + " FROM wechat_mall_coupon_log WHERE user_id = ? AND coupon_id = ?"
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query(userId, couponId)
	if err != nil {
		return nil, err
	}
	couponLog := model.WechatMallCouponLogDO{}
	if rows.Next() {
		err := rows.Scan(&couponLog.Id, &couponLog.CouponId, &couponLog.UserId, &couponLog.UseTime,
			&couponLog.ExpireTime, &couponLog.Status, &couponLog.Code, &couponLog.OrderNo, &couponLog.Del,
			&couponLog.CreateTime, &couponLog.UpdateTime)
		if err != nil {
			return nil, err
		}
	}
	return &couponLog, nil
}

func QueryCouponLogList(userId, page, size int) (*[]model.WechatMallCouponLogDO, error) {
	sql := "SELECT " + couponLogColumnList + " FROM wechat_mall_coupon_log WHERE is_del = 0 AND user_id = " + strconv.Itoa(userId)
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
	sql := "INSERT INTO wechat_mall_coupon_log ( " + couponLogColumnList[4:] + " ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(couponLog.CouponId, couponLog.UserId, couponLog.UseTime, couponLog.ExpireTime, couponLog.Status,
		couponLog.Code, couponLog.OrderNo, couponLog.Del, couponLog.CreateTime, couponLog.UpdateTime)
	return err
}
