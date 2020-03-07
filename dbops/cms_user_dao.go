package dbops

import "time"

const cmsUserColumnList = "id, username, `password`, `email`, `mobile`, `avatar`"

func GetCMSUserByUsername(username string) (*CMSUser, error) {
	sql := "SELECT " + cmsUserColumnList + " FROM wxapp_mall_cms_user WHERE username = " + username
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	cmsUser := CMSUser{}
	if rows.Next() {
		err := rows.Scan(&cmsUser.Id, &cmsUser.Username, &cmsUser.Password, &cmsUser.Email, &cmsUser.Mobile, &cmsUser.Avatar)
		if err != nil {
			return nil, err
		}
	}
	return &cmsUser, nil
}

func GetCMSUserByEmail(email string) (*CMSUser, error) {
	sql := "SELECT " + cmsUserColumnList + " FROM wxapp_mall_cms_user WHERE email = " + email
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	cmsUser := CMSUser{}
	if rows.Next() {
		err := rows.Scan(&cmsUser.Id, &cmsUser.Username, &cmsUser.Password)
		if err != nil {
			return nil, err
		}
	}
	return &cmsUser, nil
}

func AddCMSUser(user *CMSUser) error {
	sql := "INSERT INTO wxapp_mall_cms_user( " + cmsUserColumnList[4:] + " ) VALUES (?, ?, ?, ?, ?, ?, ?)"
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user.Username, user.Password, user.Email, "", "", time.Now(), time.Now())
	if err != nil {
		return err
	}
	return nil
}
