package dbops

import (
	"time"
	"wechat-mall-backend/model"
)

const cmsUserColumnList = `
id, username, password, email, mobile, avatar, is_del, create_time, update_time
`

func GetCMSUserByUsername(username string) (*model.WechatMallCMSUserDO, error) {
	sql := "SELECT " + cmsUserColumnList + " FROM wechat_mall_cms_user WHERE username = '" + username + "'"
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	cmsUser := model.WechatMallCMSUserDO{}
	if rows.Next() {
		err := rows.Scan(&cmsUser.Id, &cmsUser.Username, &cmsUser.Password, &cmsUser.Email, &cmsUser.Mobile,
			&cmsUser.Avatar, &cmsUser.Del, &cmsUser.CreateTime, &cmsUser.UpdateTime)
		if err != nil {
			return nil, err
		}
	}
	return &cmsUser, nil
}

func GetCMSUserByEmail(email string) (*model.WechatMallCMSUserDO, error) {
	sql := "SELECT " + cmsUserColumnList + " FROM wechat_mall_cms_user WHERE email = '" + email + "'"
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	cmsUser := model.WechatMallCMSUserDO{}
	if rows.Next() {
		err := rows.Scan(&cmsUser.Id, &cmsUser.Username, &cmsUser.Password, &cmsUser.Email, &cmsUser.Mobile,
			&cmsUser.Avatar, &cmsUser.Del, &cmsUser.CreateTime, &cmsUser.UpdateTime)
		if err != nil {
			return nil, err
		}
	}
	return &cmsUser, nil
}

func AddCMSUser(user *model.WechatMallCMSUserDO) error {
	sql := "INSERT INTO wechat_mall_cms_user( " + cmsUserColumnList[4:] + " ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user.Username, user.Password, user.Email, user.Mobile, user.Avatar, 0, time.Now(), time.Now())
	if err != nil {
		return err
	}
	return nil
}
