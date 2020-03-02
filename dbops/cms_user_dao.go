package dbops

import "time"

func GetCMSUserByUsername(username string) (*CMSUser, error) {
	stmt, err := dbConn.Prepare("SELECT id, username, `password`, `email`, `mobile`, `avatar` FROM wxapp_mall_cms_user WHERE username = ?")
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query(username)
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
	stmt, err := dbConn.Prepare("SELECT id, username, `password`, `email`, `mobile`, `avatar` FROM wxapp_mall_cms_user WHERE email = ?")
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query(email)
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
	stmt, err := dbConn.Prepare("INSERT INTO wxapp_mall_cms_user(`username`, `password`, `email`, `mobile`, `avatar`, `create_time`, `update_time`) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user.Username, user.Password, user.Email, "", "", time.Now(), time.Now())
	if err != nil {
		return err
	}
	return nil
}
