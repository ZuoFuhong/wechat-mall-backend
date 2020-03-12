package dbops

import (
	"strconv"
	"time"
	"wechat-mall-backend/model"
)

const couponColumnList = `
id, title, full_money, minus, rate, type, start_time, end_time, description, online, is_del, create_time, update_time
`

func QueryCouponList(page, size, online int) (*[]model.WechatMallCouponDO, error) {
	sql := "SELECT " + couponColumnList + " FROM wechat_mall_coupon WHERE is_del = 0"
	if online != 0 {
		sql += " AND online = " + strconv.Itoa(online)
	}
	if page > 0 && size > 0 {
		sql += " LIMIT " + strconv.Itoa((page-1)*page) + " ," + strconv.Itoa(size)
	}
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	var couponList []model.WechatMallCouponDO
	for rows.Next() {
		coupon := model.WechatMallCouponDO{}
		err := rows.Scan(&coupon.Id, &coupon.Title, &coupon.FullMoney, &coupon.Minus, &coupon.Rate,
			&coupon.Type, &coupon.StartTime, &coupon.EndTime, &coupon.Description, &coupon.Online,
			&coupon.Del, &coupon.CreateTime, &coupon.UpdateTime)
		if err != nil {
			return nil, err
		}
		couponList = append(couponList, coupon)
	}
	return &couponList, nil
}

func CountCoupon(online int) (int, error) {
	sql := "SELECT COUNT(*) FROM wechat_mall_coupon WHERE is_del = 0"
	if online != 0 {
		sql += " AND online = " + strconv.Itoa(online)
	}
	rows, err := dbConn.Query(sql)
	if err != nil {
		return 0, err
	}
	total := 0
	if rows.Next() {
		err := rows.Scan(&total)
		if err != nil {
			return 0, err
		}
	}
	return total, nil
}

func QueryCouponById(id int) (*model.WechatMallCouponDO, error) {
	sql := "SELECT " + couponColumnList + " FROM wechat_mall_coupon WHERE id = " + strconv.Itoa(id)
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	coupon := model.WechatMallCouponDO{}
	if rows.Next() {
		err := rows.Scan(&coupon.Id, &coupon.Title, &coupon.FullMoney, &coupon.Minus,
			&coupon.Rate, &coupon.Type, &coupon.StartTime, &coupon.EndTime, &coupon.Description,
			&coupon.Online, &coupon.Del, &coupon.CreateTime, &coupon.UpdateTime)
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
	_, err = stmt.Exec(coupon.Title, coupon.FullMoney, coupon.Minus, coupon.Rate, coupon.Type,
		coupon.StartTime, coupon.EndTime, coupon.Description, coupon.Online, 0, time.Now(), time.Now())
	if err != nil {
		return err
	}
	return nil
}

func UpdateCouponById(coupon *model.WechatMallCouponDO) error {
	sql := `
UPDATE wechat_mall_coupon 
SET title = ?, full_money = ?, minus = ?, rate = ?, type = ?, start_time = ?, 
    end_time = ?,  description = ?, online = ?, is_del = ?, update_time = ?
WHERE id = ?
`
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(coupon.Title, coupon.FullMoney, coupon.Minus, coupon.Rate, coupon.Type,
		coupon.StartTime, coupon.EndTime, coupon.Description, coupon.Online, coupon.Del, time.Now(), coupon.Id)
	if err != nil {
		return err
	}
	return nil
}
