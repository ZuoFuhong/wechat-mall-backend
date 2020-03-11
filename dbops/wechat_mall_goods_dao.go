package dbops

import (
	"strconv"
	"time"
	"wechat-mall-backend/model"
)

const goodsColunList = `
id, brand_name, title, subtitle, price, discount_price, category_id, default_sku_id, online, picture, 
banner_picture, detail_picture, tags, sketch_spec_id, description, is_del, create_time, update_time
`

func QueryGoodsList(page, size int) (*[]model.WechatMallGoodsDO, error) {
	sql := "SELECT" + goodsColunList + " FROM wechat_mall_goods WHERE is_del = 0 LIMIT ?, ?"
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query((page-1)*size, size)
	if err != nil {
		return nil, err
	}
	var goodsList []model.WechatMallGoodsDO
	for rows.Next() {
		goods := model.WechatMallGoodsDO{}
		err := rows.Scan(&goods.Id, &goods.BrandName, &goods.Title, &goods.SubTitle, &goods.Price, &goods.DiscountPrice,
			&goods.CategoryId, &goods.DefaultSkuId, &goods.Online, &goods.Picture, &goods.BannerPicture,
			&goods.DetailPicture, &goods.Tags, &goods.SketchSpecId, &goods.Description, &goods.Del,
			&goods.CreateTime, &goods.UpdateTime)
		if err != nil {
			return nil, err
		}
		goodsList = append(goodsList, goods)
	}
	return &goodsList, nil
}

func CountGoods() (int, error) {
	sql := "SELECT COUNT(*) FROM wechat_mall_goods WHERE is_del = 0"
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

func AddGoods(goods *model.WechatMallGoodsDO) (int64, error) {
	sql := "INSERT INTO wechat_mall_goods ( " + goodsColunList[4:] + " ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return 0, err
	}
	result, err := stmt.Exec(goods.BrandName, goods.Title, goods.SubTitle, goods.Price, goods.DiscountPrice,
		goods.CategoryId, goods.DefaultSkuId, goods.Online, goods.Picture, goods.BannerPicture, goods.DetailPicture,
		goods.Tags, goods.SketchSpecId, goods.Description, 0, time.Now(), time.Now())
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
	if rows.Next() {
		err := rows.Scan(&goods.Id, &goods.BrandName, &goods.Title, &goods.SubTitle, &goods.Price,
			&goods.DiscountPrice, &goods.CategoryId, &goods.DefaultSkuId, &goods.Online, &goods.Picture,
			&goods.BannerPicture, &goods.DetailPicture, &goods.Tags, &goods.SketchSpecId, &goods.Description,
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
SET brand_name = ?, title = ?, subtitle = ?, price = ?, discount_price = ?, category_id = ?, default_sku_id = ?, 
online = ?, picture = ?, banner_picture = ?, detail_picture = ?, tags = ?, sketch_spec_id = ?,
description = ?, is_del = ?, update_time = ?
WHERE id = ?
`
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(goods.BrandName, goods.Title, goods.SubTitle, goods.Price, goods.DiscountPrice,
		goods.CategoryId, goods.DefaultSkuId, goods.Online, goods.Picture, goods.BannerPicture, goods.DetailPicture,
		goods.Tags, goods.SketchSpecId, goods.Description, goods.Del, time.Now(), goods.Id)
	if err != nil {
		return err
	}
	return nil
}
