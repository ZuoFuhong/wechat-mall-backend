package service

import (
	"context"
	"wechat-mall-backend/app/domain/entity"
	"wechat-mall-backend/app/domain/repository"
	"wechat-mall-backend/app/domain/view"
	"wechat-mall-backend/pkg/utils"
)

type IBrowseRecordService interface {
	AddBrowseRecord(ctx context.Context, record *entity.WechatMallGoodsBrowseRecord) error

	ListBrowseRecord(ctx context.Context, userId, page, size int) ([]*view.PortalBrowseRecordVO, int, error)

	ClearBrowseHistory(ctx context.Context, ids []int)
}

type browseRecordService struct {
	repos repository.IBrowseRepos
}

func NewBrowseRecordService(repos repository.IBrowseRepos) IBrowseRecordService {
	service := &browseRecordService{
		repos: repos,
	}
	return service
}

func (s *browseRecordService) AddBrowseRecord(ctx context.Context, record *entity.WechatMallGoodsBrowseRecord) error {
	recordDO, err := s.repos.SelectGoodsBrowse(ctx, record.UserID, record.GoodsID)
	if err != nil {
		return err
	}
	if recordDO.ID != 0 {
		if err := s.repos.DeleteBrowseRecordById(ctx, recordDO.ID); err != nil {
			return err
		}
	}
	return s.repos.AddBrowseRecord(ctx, record)
}

func (s *browseRecordService) ListBrowseRecord(ctx context.Context, userId, page, size int) ([]*view.PortalBrowseRecordVO, int, error) {
	records, total, err := s.repos.SelectGoodsBrowseByUserId(ctx, userId, page, size)
	if err != nil {
		return nil, 0, err
	}
	voList := make([]*view.PortalBrowseRecordVO, 0)
	for _, recordDO := range records {
		recordVO := &view.PortalBrowseRecordVO{
			Id:         recordDO.ID,
			GoodsId:    recordDO.GoodsID,
			Picture:    recordDO.Picture,
			Title:      recordDO.Title,
			Price:      recordDO.Price,
			BrowseTime: utils.FormatTime(recordDO.UpdateTime),
		}
		voList = append(voList, recordVO)
	}
	return voList, total, nil
}

func (s *browseRecordService) ClearBrowseHistory(ctx context.Context, ids []int) {
	for _, id := range ids {
		if err := s.repos.DeleteBrowseRecordById(ctx, id); err != nil {
			return
		}
	}
}
