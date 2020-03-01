package store

import "time"

func (m *MySQLStore) GetUserByOpenid(openid string) (*WxappUser, error) {
	stmt, err := m.client.Prepare("SELECT * FROM wxapp_mall_user WHERE openid = ?")
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query(openid)
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

func (m *MySQLStore) AddMiniappUser(user *WxappUser) (int64, error) {
	sql := `
INSERT INTO wxapp_mall_user(openid, nickname, avatar, mobile, city, create_time, update_time) 
VALUES(?, ?, ?, ?, ?, ?, ?)
`
	stmt, err := m.client.Prepare(sql)
	if err != nil {
		return 0, err
	}
	result, err := stmt.Exec(user.Openid, user.Nickname, user.Avatar, user.Mobile, user.City, time.Now(), time.Now())
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}
