package repository

import (
	"context"
	"wechat-mall-backend/app/domain/entity"
)

type IBrowseRepos interface {
	AddBrowseRecord(ctx context.Context, record *entity.WechatMallGoodsBrowseRecord) error

	SelectGoodsBrowse(ctx context.Context, userId, goodsId int) (*entity.WechatMallGoodsBrowseRecord, error)

	DeleteBrowseRecordById(ctx context.Context, id int) error

	SelectGoodsBrowseByUserId(ctx context.Context, userId, page, size int) ([]*entity.WechatMallGoodsBrowseRecord, int, error)
}
