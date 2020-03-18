package dbops

import (
	"strconv"
	"time"
	"wechat-mall-backend/model"
)

const goodsColunList = `
id, brand_name, title, price, discount_price, category_id, online, picture, 
banner_picture, detail_picture, tags, is_del, create_time, update_time
`

func QueryGoodsList(keyword string, categoryId, online, page, size int) (*[]model.WechatMallGoodsDO, error) {
	sql := "SELECT " + goodsColunList + " FROM wechat_mall_goods WHERE is_del = 0"
	if keyword != "" {
		sql += " AND title LIKE '%" + keyword + "%'"
	}
	if categoryId != 0 {
		sql += " AND category_id = " + strconv.Itoa(categoryId)
	}
	if online == 0 || online == 1 {
		sql += " AND online = " + strconv.Itoa(online)
	}
	if page > 0 && size > 0 {
		sql += " LIMIT " + strconv.Itoa((page-1)*size) + ", " + strconv.Itoa(size)
	}
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	goodsList := []model.WechatMallGoodsDO{}
	for rows.Next() {
		goods := model.WechatMallGoodsDO{}
		err := rows.Scan(&goods.Id, &goods.BrandName, &goods.Title, &goods.Price, &goods.DiscountPrice,
			&goods.CategoryId, &goods.Online, &goods.Picture, &goods.BannerPicture, &goods.DetailPicture,
			&goods.Tags, &goods.Del, &goods.CreateTime, &goods.UpdateTime)
		if err != nil {
			return nil, err
		}
		goodsList = append(goodsList, goods)
	}
	return &goodsList, nil
}

func CountGoods(keyword string, categoryId, online int) (int, error) {
	sql := "SELECT COUNT(*) FROM wechat_mall_goods WHERE is_del = 0"
	if keyword != "" {
		sql += " AND title LIKE '%" + keyword + "%'"
	}
	if categoryId != 0 {
		sql += " AND category_id = " + strconv.Itoa(categoryId)
	}
	if online == 0 || online == 1 {
		sql += " AND online = " + strconv.Itoa(online)
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

func AddGoods(goods *model.WechatMallGoodsDO) (int64, error) {
	sql := "INSERT INTO wechat_mall_goods ( " + goodsColunList[4:] + " ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return 0, err
	}
	result, err := stmt.Exec(goods.BrandName, goods.Title, goods.Price, goods.DiscountPrice,
		goods.CategoryId, goods.Online, goods.Picture, goods.BannerPicture, goods.DetailPicture,
		goods.Tags, 0, time.Now(), time.Now())
	if err != nil {
		return 0, err
	}
	i, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return i, nil
}

func QueryGoodsById(id int) (*model.WechatMallGoodsDO, error) {
	sql := "SELECT " + goodsColunList + " FROM wechat_mall_goods WHERE is_del = 0 AND id = " + strconv.Itoa(id)
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	goods := model.WechatMallGoodsDO{}
	for rows.Next() {
		err := rows.Scan(&goods.Id, &goods.BrandName, &goods.Title, &goods.Price,
			&goods.DiscountPrice, &goods.CategoryId, &goods.Online, &goods.Picture,
			&goods.BannerPicture, &goods.DetailPicture, &goods.Tags,
			&goods.Del, &goods.CreateTime, &goods.UpdateTime)
		if err != nil {
			return nil, err
		}
	}
	return &goods, nil
}

func UpdateGoodsById(goods *model.WechatMallGoodsDO) error {
	sql := `
UPDATE wechat_mall_goods 
SET brand_name = ?, title = ?, price = ?, discount_price = ?, category_id = ?,
online = ?, picture = ?, banner_picture = ?, detail_picture = ?, tags = ?,
is_del = ?, update_time = ?
WHERE id = ?
`
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(goods.BrandName, goods.Title, goods.Price, goods.DiscountPrice,
		goods.CategoryId, goods.Online, goods.Picture, goods.BannerPicture, goods.DetailPicture,
		goods.Tags, goods.Del, time.Now(), goods.Id)
	if err != nil {
		return err
	}
	return nil
}

func CountCategoryGoods(categoryId int) (int, error) {
	sql := "SELECT COUNT(*) FROM wechat_mall_category WHERE is_del = 0 AND category_id = " + strconv.Itoa(categoryId)
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
