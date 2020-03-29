package dbops

import (
	"strconv"
	"time"
	"wechat-mall-backend/model"
)

const browseColumnList = `
id, user_id, goods_id, picture, title, price, is_del, create_time, update_time
`

func InsertBrowseRecord(record *model.WechatMallGoodsBrowseRecord) error {
	sql := "INSERT INTO wechat_mall_goods_browse_record (" + browseColumnList[4:] + ") VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(record.UserId, record.GoodsId, record.Picture, record.Title, record.Price, 0, time.Now(), time.Now())
	return err
}

func SelectGoodsBrowse(userId, goodsId int) (*model.WechatMallGoodsBrowseRecord, error) {
	sql := "SELECT " + browseColumnList + " FROM wechat_mall_goods_browse_record WHERE is_del = 0 AND user_id = " +
		strconv.Itoa(userId) + " AND goods_id = " + strconv.Itoa(goodsId)

	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	record := model.WechatMallGoodsBrowseRecord{}
	for rows.Next() {
		err := rows.Scan(&record.Id, &record.UserId, &record.GoodsId, &record.Picture, &record.Title, &record.Price,
			&record.Del, &record.CreateTime, &record.UpdateTime)
		if err != nil {
			return nil, err
		}
	}
	return &record, nil
}

func DeleteBrowseRecordById(id int) error {
	sql := "UPDATE wechat_mall_goods_browse_record SET is_del = 1 WHERE id = " + strconv.Itoa(id)
	_, err := dbConn.Exec(sql)
	return err
}

func SelectGoodsBrowseByUserId(userId, page, size int) (*[]model.WechatMallGoodsBrowseRecord, error) {
	sql := "SELECT " + browseColumnList + " FROM wechat_mall_goods_browse_record WHERE is_del = 0 AND user_id = " + strconv.Itoa(userId)
	if page > 0 && size > 0 {
		sql += " LIMIT " + strconv.Itoa((page-1)*size) + ", " + strconv.Itoa(size)
	}
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	records := []model.WechatMallGoodsBrowseRecord{}
	for rows.Next() {
		record := model.WechatMallGoodsBrowseRecord{}
		err := rows.Scan(&record.Id, &record.UserId, &record.GoodsId, &record.Picture, &record.Title, &record.Price,
			&record.Del, &record.CreateTime, &record.UpdateTime)
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}
	return &records, nil
}

func CountGoodsBrowseByUserId(userId int) (int, error) {
	sql := "SELECT COUNT(*) FROM wechat_mall_goods_browse_record WHERE is_del = 0 AND user_id " + strconv.Itoa(userId)
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
