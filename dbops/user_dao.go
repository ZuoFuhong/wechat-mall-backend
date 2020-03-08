package dbops

import "time"

const userColumnList = `
id, openid, nickname, avatar, mobile, city, create_time, update_time
`

func GetUserByOpenid(openid string) (*WxappUser, error) {
	sql := "SELECT " + userColumnList + " FROM wxapp_mall_user WHERE openid = '" + openid + "'"
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	user := WxappUser{}
	if rows.Next() {
		err = rows.Scan(&user.Id, &user.Openid, &user.Nickname, &user.Avatar, &user.Mobile, &user.City, &user.CreateTime, &user.UpdateTime)
		if err != nil {
			return nil, err
		}
	}
	return &user, nil
}

func AddMiniappUser(user *WxappUser) (int64, error) {
	sql := "INSERT INTO wxapp_mall_user(" + userColumnList[4:] + ") VALUES(?, ?, ?, ?, ?, ?, ?)"
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return 0, err
	}
	result, err := stmt.Exec(user.Openid, user.Nickname, user.Avatar, user.Mobile, user.City, time.Now(), time.Now())
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}
