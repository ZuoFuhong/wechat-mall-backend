package dbops

import (
	"strconv"
	"time"
	"wechat-mall-backend/model"
)

const cmsUserColumnList = `
id, username, password, email, mobile, avatar, group_id, is_del, create_time, update_time
`

const cmsUserGroupColumnList = `
id, name, description, is_del, create_time, update_time
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
			&cmsUser.Avatar, &cmsUser.GroupId, &cmsUser.Del, &cmsUser.CreateTime, &cmsUser.UpdateTime)
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
			&cmsUser.Avatar, &cmsUser.GroupId, &cmsUser.Del, &cmsUser.CreateTime, &cmsUser.UpdateTime)
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
	_, err = stmt.Exec(user.Username, user.Password, user.Email, user.Mobile, user.Avatar, user.GroupId, 0, time.Now(), time.Now())
	if err != nil {
		return err
	}
	return nil
}

func CountGroupUser(groupId int) (int, error) {
	sql := "SELECT COUNT(*) FROM wechat_mall_cms_user WHERE is_del = 0 AND group_id = " + strconv.Itoa(groupId)
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
	return total, err
}

func QueryCMSUser(id int) (*model.WechatMallCMSUserDO, error) {
	sql := "SELECT " + cmsUserColumnList + " FROM wechat_mall_cms_user WHERE is_del = 0 AND id = " + strconv.Itoa(id)
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	userDO := model.WechatMallCMSUserDO{}
	if rows.Next() {
		err := rows.Scan(&userDO.Id, &userDO.Username, &userDO.Password, &userDO.Email, &userDO.Mobile, &userDO.Avatar,
			&userDO.GroupId, &userDO.Del, &userDO.CreateTime, &userDO.UpdateTime)
		if err != nil {
			return nil, err
		}
	}
	return &userDO, nil
}

func UpdateCMSUserById(userDO *model.WechatMallCMSUserDO) error {
	sql := `
UPDATE wechat_mall_cms_user
SET username = ?, password = ?, email = ?, mobile = ?, avatar = ?, group_id = ?, is_del = ?, update_time = ?
WHERE id = ?
`
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(userDO.Username, userDO.Password, userDO.Email, userDO.Mobile, userDO.Avatar, userDO.GroupId,
		userDO.Del, time.Now(), userDO.Id)
	return err
}

func ListCMSUser(page, size int) (*[]model.WechatMallCMSUserDO, error) {
	sql := "SELECT " + cmsUserColumnList + " FROM wechat_mall_cms_user WHERE is_del = 0 AND id != 1"
	if page > 0 && size > 0 {
		sql += " LIMIT " + strconv.Itoa((page-1)*size) + ", " + strconv.Itoa(size)
	}
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	userList := []model.WechatMallCMSUserDO{}
	for rows.Next() {
		userDO := model.WechatMallCMSUserDO{}
		err := rows.Scan(&userDO.Id, &userDO.Username, &userDO.Password, &userDO.Email, &userDO.Mobile, &userDO.Avatar,
			&userDO.GroupId, &userDO.Del, &userDO.CreateTime, &userDO.UpdateTime)
		if err != nil {
			return nil, err
		}
		userList = append(userList, userDO)
	}
	return &userList, nil
}

func CountCMSUser() (int, error) {
	sql := "SELECT COUNT(*) FROM wechat_mall_cms_user WHERE is_del = 0 AND id != 1"
	rows, err := dbConn.Query(sql)
	if err != nil {
		return 0, nil
	}
	total := 0
	if rows.Next() {
		err := rows.Scan(&total)
		if err != nil {
			return 0, nil
		}
	}
	return total, nil
}

func AddUserGroup(group *model.WechatMallUserGroupDO) (int64, error) {
	sql := "INSERT INTO wechat_mall_user_group ( " + cmsUserGroupColumnList[4:] + " ) VALUES (?, ?, ?, ?, ?)"
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return 0, err
	}
	result, err := stmt.Exec(group.Name, group.Description, 0, time.Now(), time.Now())
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func QueryUserGroupById(id int) (*model.WechatMallUserGroupDO, error) {
	sql := "SELECT " + cmsUserGroupColumnList + " FROM wechat_mall_user_group WHERE is_del = 0 AND id = " + strconv.Itoa(id)
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	groupDO := model.WechatMallUserGroupDO{}
	if rows.Next() {
		err := rows.Scan(&groupDO.Id, &groupDO.Name, &groupDO.Description, &groupDO.Del, &groupDO.CreateTime, &groupDO.UpdateTime)
		if err != nil {
			return nil, err
		}
	}
	return &groupDO, nil
}

func QueryUserGroupByName(name string) (*model.WechatMallUserGroupDO, error) {
	sql := "SELECT " + cmsUserGroupColumnList + " FROM wechat_mall_user_group WHERE is_del = 0 AND name = '" + name + "'"
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	groupDO := model.WechatMallUserGroupDO{}
	if rows.Next() {
		err := rows.Scan(&groupDO.Id, &groupDO.Name, &groupDO.Description, &groupDO.Del, &groupDO.CreateTime, &groupDO.UpdateTime)
		if err != nil {
			return nil, err
		}
	}
	return &groupDO, nil
}

func QueryGroupList(page, size int) (*[]model.WechatMallUserGroupDO, error) {
	sql := "SELECT " + cmsUserGroupColumnList + " FROM wechat_mall_user_group WHERE is_del = 0"
	if page > 0 && size > 0 {
		sql += " LIMIT " + strconv.Itoa((page-1)*size) + ", " + strconv.Itoa(size)
	}
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	groupList := []model.WechatMallUserGroupDO{}
	for rows.Next() {
		groupDO := model.WechatMallUserGroupDO{}
		err := rows.Scan(&groupDO.Id, &groupDO.Name, &groupDO.Description, &groupDO.Del, &groupDO.CreateTime, &groupDO.UpdateTime)
		if err != nil {
			return nil, err
		}
		groupList = append(groupList, groupDO)
	}
	return &groupList, nil
}

func CountUserCoupon() (int, error) {
	sql := "SELECT COUNT(*) FROM wechat_mall_user_group WHERE is_del = 0"
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

func UpdateGroupById(group *model.WechatMallUserGroupDO) error {
	sql := `
UPDATE wechat_mall_user_group
SET name = ?, description = ?, is_del = ?, update_time = ?
WHERE id = ?
`
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(group.Name, group.Description, group.Del, time.Now(), group.Id)
	return err
}
