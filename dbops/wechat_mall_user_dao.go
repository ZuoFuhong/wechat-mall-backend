package dbops

import (
	"strconv"
	"time"
	"wechat-mall-backend/model"
)

const userColumnList = `
id, openid, nickname, avatar, mobile, city, province, country, gender, create_time, update_time
`

func GetUserByOpenid(openid string) (*model.WechatMallUserDO, error) {
	sql := "SELECT " + userColumnList + " FROM wechat_mall_user WHERE openid = '" + openid + "'"
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	user := model.WechatMallUserDO{}
	if rows.Next() {
		err = rows.Scan(&user.Id, &user.Openid, &user.Nickname, &user.Avatar, &user.Mobile, &user.City,
			&user.Province, &user.Country, &user.Gender, &user.CreateTime, &user.UpdateTime)
		if err != nil {
			return nil, err
		}
	}
	return &user, nil
}

func GetUserById(id int) (*model.WechatMallUserDO, error) {
	sql := "SELECT " + userColumnList + " FROM wechat_mall_user WHERE id = " + strconv.Itoa(id)
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	user := model.WechatMallUserDO{}
	if rows.Next() {
		err = rows.Scan(&user.Id, &user.Openid, &user.Nickname, &user.Avatar, &user.Mobile, &user.City,
			&user.Province, &user.Country, &user.Gender, &user.CreateTime, &user.UpdateTime)
		if err != nil {
			return nil, err
		}
	}
	return &user, nil
}

func AddMiniappUser(user *model.WechatMallUserDO) (int64, error) {
	sql := "INSERT INTO wechat_mall_user(" + userColumnList[4:] + ") VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return 0, err
	}
	result, err := stmt.Exec(user.Openid, user.Nickname, user.Avatar, user.Mobile, user.City, user.Province,
		user.Country, user.Gender, time.Now(), time.Now())
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func UpdateUserById(user *model.WechatMallUserDO) error {
	sql := `
UPDATE wechat_mall_user
SET nickname = ?, avatar = ?, mobile = ?, city = ?, province = ?, country = ?, gender = ?, update_time = ?
WHERE id = ?
`
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user.Nickname, user.Avatar, user.Mobile, user.City, user.Province, user.Country,
		user.Gender, time.Now(), user.Id)
	return err
}
