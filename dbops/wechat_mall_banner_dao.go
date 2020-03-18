package dbops

import (
	"strconv"
	"time"
	"wechat-mall-backend/model"
)

const bannerColumnList = `
id, picture, name, title, description, is_del, create_time, update_time
`

func QueryBannerList(name string, page, size int) (*[]model.WechatMallBannerDO, error) {
	sql := "SELECT " + bannerColumnList + "FROM wechat_mall_banner WHERE is_del = 0"
	if name != "" {
		sql += " AND name = " + name
	}
	sql += " LIMIT " + strconv.Itoa((page-1)*size) + ", " + strconv.Itoa(size)
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	var bannerList []model.WechatMallBannerDO
	for rows.Next() {
		banner := model.WechatMallBannerDO{}
		err := rows.Scan(&banner.Id, &banner.Picture, &banner.Name, &banner.Title, &banner.Description, &banner.Del, &banner.CreateTime, &banner.UpdateTime)
		if err != nil {
			return nil, err
		}
		bannerList = append(bannerList, banner)
	}
	return &bannerList, nil
}

func CountBanner(name string) (int, error) {
	sql := "SELECT COUNT(*) FROM wechat_mall_banner WHERE is_del = 0"
	if name != "" {
		sql += " AND name = " + name
	}
	rows, err := dbConn.Query(sql)
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

func QueryBannerById(id int) (*model.WechatMallBannerDO, error) {
	sql := "SELECT " + bannerColumnList + " FROM wechat_mall_banner WHERE is_del = 0 AND id = " + strconv.Itoa(id)
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	banner := model.WechatMallBannerDO{}
	for rows.Next() {
		err := rows.Scan(&banner.Id, &banner.Picture, &banner.Name, &banner.Title, &banner.Description, &banner.Del, &banner.CreateTime, &banner.UpdateTime)
		if err != nil {
			return nil, err
		}
	}
	return &banner, nil
}

func InsertBanner(banner *model.WechatMallBannerDO) (int64, error) {
	sql := "INSERT INTO wechat_mall_banner( " + bannerColumnList[4:] + " ) VALUES (?, ?, ?, ?, ?, ?, ?)"
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return 0, err
	}
	result, err := stmt.Exec(banner.Picture, banner.Name, banner.Title, banner.Description, 0, time.Now(), time.Now())
	if err != nil {
		return 0, err
	}
	autoId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return autoId, nil
}

func UpdateBannerById(banner *model.WechatMallBannerDO) error {
	sql := `
UPDATE wechat_mall_banner 
SET picture = ?, name = ?, title = ?, description = ?, is_del = ?, update_time = ?
WHERE id = ?
`
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(banner.Picture, banner.Name, banner.Title, banner.Description, banner.Del, time.Now(), banner.Id)
	if err != nil {
		return err
	}
	return nil
}
