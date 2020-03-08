package dbops

import (
	"strconv"
	"time"
	"wechat-mall-backend/model"
)

const bannerColumnList = `
id, picture, name, title, description, is_del, create_time, update_time
`

const bannerItemColumnList = `
id, banner_id, name, picture, keyword, type, is_del, create_time, update_time
`

func QueryBannerList(name string, page, size int) (*[]model.Banner, error) {
	sql := "SELECT " + bannerColumnList + "FROM wxapp_mall_banner WHERE is_del = 0"
	if name != "" {
		sql += " AND name = " + name
	}
	sql += " LIMIT " + strconv.Itoa((page-1)*size) + ", " + strconv.Itoa(size)
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	var bannerList []model.Banner
	for rows.Next() {
		banner := model.Banner{}
		err := rows.Scan(&banner.Id, &banner.Picture, &banner.Name, &banner.Title, &banner.Description, &banner.Del, &banner.CreateTime, &banner.UpdateTime)
		if err != nil {
			return nil, err
		}
		bannerList = append(bannerList, banner)
	}
	return &bannerList, nil
}

func CountBanner(name string) (int, error) {
	sql := "SELECT COUNT(*) FROM wxapp_mall_banner WHERE is_del = 0"
	if name != "" {
		sql += " AND name = " + name
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

func QueryBannerById(id int) (*model.Banner, error) {
	sql := "SELECT " + bannerColumnList + " FROM wxapp_mall_banner WHERE is_del = 0 AND id = " + strconv.Itoa(id)
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	banner := model.Banner{}
	if rows.Next() {
		err := rows.Scan(&banner.Id, &banner.Picture, &banner.Name, &banner.Title, &banner.Description, &banner.Del, &banner.CreateTime, &banner.UpdateTime)
		if err != nil {
			return nil, err
		}
	}
	return &banner, nil
}

func InsertBanner(banner *model.Banner) (int64, error) {
	sql := "INSERT INTO wxapp_mall_banner( " + bannerColumnList[4:] + " ) VALUES (?, ?, ?, ?, ?, ?, ?)"
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

func UpdateBannerById(banner *model.Banner) error {
	sql := `
UPDATE wxapp_mall_banner 
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

func QueryBannerItemList(bannerId int) (*[]model.BannerItem, error) {
	sql := "SELECT " + bannerItemColumnList + " FROM wxapp_mall_banner_item WHERE is_del = 0 AND banner_id = " + strconv.Itoa(bannerId)
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	var bannerItemList []model.BannerItem
	for rows.Next() {
		bannerItem := model.BannerItem{}
		err := rows.Scan(&bannerItem.Id, &bannerItem.BannerId, &bannerItem.Name, &bannerItem.Picture, &bannerItem.Keyword, &bannerItem.Type, &bannerItem.Del, &bannerItem.CreateTime, &bannerItem.UpdateTime)
		if err != nil {
			return nil, err
		}
		bannerItemList = append(bannerItemList, bannerItem)
	}
	return &bannerItemList, nil
}

func QueryBannerItemById(id int) (*model.BannerItem, error) {
	sql := "SELECT " + bannerItemColumnList + " FROM wxapp_mall_banner_item WHERE is_del = 0 AND id = " + strconv.Itoa(id)
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	bannerItem := model.BannerItem{}
	if rows.Next() {
		err := rows.Scan(&bannerItem.Id, &bannerItem.BannerId, &bannerItem.Name, &bannerItem.Picture, &bannerItem.Keyword, &bannerItem.Type, &bannerItem.Del, &bannerItem.CreateTime, &bannerItem.UpdateTime)
		if err != nil {
			return nil, err
		}
	}
	return &bannerItem, nil
}

func InsertBannerItem(bannerItem *model.BannerItem) (int64, error) {
	sql := "INSERT INTO wxapp_mall_banner_item( " + bannerItemColumnList[4:] + " ) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return 0, nil
	}
	result, err := stmt.Exec(bannerItem.BannerId, bannerItem.Name, bannerItem.Picture, bannerItem.Keyword, bannerItem.Type, 0, time.Now(), time.Now())
	if err != nil {
		return 0, err
	}
	autoId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return autoId, nil
}

func UpdateBannerItemById(bannerItem *model.BannerItem) error {
	sql := `
UPDATE wxapp_mall_banner_item 
SET banner_id = ?, name = ?, picture = ?, keyword = ?, type = ?, is_del = ?, update_time = ? 
WHERE id = ?
`
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(bannerItem.BannerId, bannerItem.Name, bannerItem.Picture, bannerItem.Keyword,
		bannerItem.Type, bannerItem.Del, time.Now(), bannerItem.Id)

	if err != nil {
		return err
	}
	return nil
}
