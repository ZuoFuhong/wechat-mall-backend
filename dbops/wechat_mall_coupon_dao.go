package dbops

import (
	"strconv"
	"time"
	"wechat-mall-backend/model"
)

const couponColumnList = `
id, activity_id, title, full_money, minus, rate, type, start_time, end_time, description, 
is_del, create_time, update_time
`

func QueryCouponList(activityId, page, size int) (*[]model.WechatMallCouponDO, error) {
	sql := "SELECT " + couponColumnList + " FROM wechat_mall_coupon WHERE is_del = 0 AND activity_id = " + strconv.Itoa(activityId)
	if page > 0 && size > 0 {
		sql += " LIMIT " + strconv.Itoa((page-1)*page) + " , " + strconv.Itoa(size)
	}
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	var couponList []model.WechatMallCouponDO
	for rows.Next() {
		coupon := model.WechatMallCouponDO{}
		err := rows.Scan(&coupon.Id, &coupon.ActivityId, &coupon.Title, &coupon.FullMoney, &coupon.Minus, &coupon.Rate,
			&coupon.Type, &coupon.StartTime, &coupon.EndTime, &coupon.Description, &coupon.Del,
			&coupon.CreateTime, &coupon.UpdateTime)
		if err != nil {
			return nil, err
		}
		couponList = append(couponList, coupon)
	}
	return &couponList, nil
}

func QueryCouponById(id int) (*model.WechatMallCouponDO, error) {
	sql := "SELECT " + couponColumnList + " FROM wechat_mall_coupon WHERE is_del = 0 AND id = " + strconv.Itoa(id)
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	coupon := model.WechatMallCouponDO{}
	if rows.Next() {
		err := rows.Scan(&coupon.Id, &coupon.ActivityId, &coupon.Title, &coupon.FullMoney, &coupon.Minus,
			&coupon.Rate, &coupon.Type, &coupon.StartTime, &coupon.EndTime, &coupon.Description, &coupon.Del,
			&coupon.CreateTime, &coupon.UpdateTime)
		if err != nil {
			return nil, err
		}
	}
	return &coupon, nil
}

func InsertCoupon(coupon *model.WechatMallCouponDO) error {
	sql := "INSERT INTO wechat_mall_coupon( " + couponColumnList[4:] + " ) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(coupon.ActivityId, coupon.Title, coupon.FullMoney, coupon.Minus, coupon.Rate, coupon.Type,
		coupon.StartTime, coupon.EndTime, coupon.Description, 0, time.Now(), time.Now())
	if err != nil {
		return err
	}
	return nil
}

func UpdateCouponById(coupon *model.WechatMallCouponDO) error {
	sql := `
UPDATE wechat_mall_coupon 
SET activity_id = ?, title = ?, full_money = ?, minus = ?, rate = ?, type = ?, start_time = ?, 
    end_time = ?,  description = ?, is_del = ?, update_time = ?
WHERE id = ?
`
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(coupon.ActivityId, coupon.Title, coupon.FullMoney, coupon.Minus, coupon.Rate, coupon.Type,
		coupon.StartTime, coupon.EndTime, coupon.Description, coupon.Del, time.Now(), coupon.Id)
	if err != nil {
		return err
	}
	return nil
}
