package dbops

import "time"

const visitorColumnList = `
id, user_id, ip, create_time, update_time
`

func AddVisitorRecord(userId int, ip string) error {
	sql := "INSERT INTO wechat_mall_visitor_record (" + visitorColumnList[4:] + ") VALUES (?, ?, ?, ?)"
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(userId, ip, time.Now(), time.Now())
	return err
}

func CountUniqueVisitor(startTime, endTime time.Time) (int, error) {
	sql := "SELECT COUNT(DISTINCT(user_id)) FROM wechat_mall_visitor_record WHERE create_time BETWEEN ? AND ?"
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return 0, err
	}
	rows, err := stmt.Query(startTime, endTime)
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
