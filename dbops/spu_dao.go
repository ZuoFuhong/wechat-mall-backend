package dbops

import (
	"strconv"
	"time"
	"wechat-mall-backend/model"
)

const SPUColunList = `
id, brand_name, title, subtitle, price, discount_price, category_id, default_sku_id, online, picture, 
for_theme_picture, banner_picture, detail_picture, tags, sketch_spec_id, description, is_del, 
create_time, update_time
`

func QuerySPUList(page, size int) (*[]model.SPU, error) {
	sql := "SELECT" + SPUColunList + " FROM wxapp_mall_spu WHERE is_del = 0 LIMIT ?, ?"
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query((page-1)*size, size)
	if err != nil {
		return nil, err
	}
	var spuList []model.SPU
	for rows.Next() {
		spu := model.SPU{}
		err := rows.Scan(&spu.Id, &spu.BrandName, &spu.Title, &spu.SubTitle, &spu.Price, &spu.DiscountPrice,
			&spu.CategoryId, &spu.DefaultSkuId, &spu.Online, &spu.Picture, &spu.ForThemePicture,
			&spu.BannerPicture, &spu.DetailPicture, &spu.Tags, &spu.SketchSpecId, &spu.Description, &spu.Del,
			&spu.CreateTime, &spu.UpdateTime)
		if err != nil {
			return nil, err
		}
		spuList = append(spuList, spu)
	}
	return &spuList, nil
}

func CountSPU() (int, error) {
	sql := "SELECT COUNT(*) FROM wxapp_mall_spu WHERE is_del = 0"
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

func AddSPU(spu *model.SPU) (int64, error) {
	sql := "INSERT INTO wxapp_mall_spu ( " + SPUColunList[4:] + " ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return 0, err
	}
	result, err := stmt.Exec(spu.BrandName, spu.Title, spu.SubTitle, spu.Price, spu.DiscountPrice, spu.CategoryId,
		spu.DefaultSkuId, spu.Online, spu.Picture, spu.ForThemePicture, spu.BannerPicture, spu.DetailPicture, spu.Tags,
		spu.SketchSpecId, spu.Description, 0, time.Now(), time.Now())
	if err != nil {
		return 0, err
	}
	i, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return i, nil
}

func QuerySPUById(id int) (*model.SPU, error) {
	sql := "SELECT " + SPUColunList + " FROM wxapp_mall_spu WHERE is_del = 0 AND id = " + strconv.Itoa(id)
	rows, err := dbConn.Query(sql)
	if err != nil {
		return nil, err
	}
	spu := model.SPU{}
	if rows.Next() {
		err := rows.Scan(&spu.Id, &spu.BrandName, &spu.Title, &spu.SubTitle, &spu.Price, &spu.DiscountPrice,
			&spu.CategoryId, &spu.DefaultSkuId, &spu.Online, &spu.Picture, &spu.ForThemePicture,
			&spu.BannerPicture, &spu.DetailPicture, &spu.Tags, &spu.SketchSpecId, &spu.Description, &spu.Del,
			&spu.CreateTime, &spu.UpdateTime)
		if err != nil {
			return nil, err
		}
	}
	return &spu, nil
}

func UpdateSPUById(spu *model.SPU) error {
	sql := `
UPDATE wxapp_mall_spu 
SET brand_name = ?, title = ?, subtitle = ?, price = ?, discount_price = ?, category_id = ?, default_sku_id = ?, 
online = ?, picture = ?, for_theme_picture = ?, banner_picture = ?, detail_picture = ?, tags = ?, sketch_spec_id = ?,
description = ?, is_del = ?, update_time = ?
WHERE id = ?
`
	stmt, err := dbConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(spu.BrandName, spu.Title, spu.SubTitle, spu.Price, spu.DiscountPrice, spu.CategoryId, spu.DefaultSkuId,
		spu.Online, spu.Picture, spu.ForThemePicture, spu.BannerPicture, spu.DetailPicture, spu.Tags, spu.SketchSpecId,
		spu.Description, spu.Del, time.Now(), spu.Id)
	if err != nil {
		return err
	}
	return nil
}
