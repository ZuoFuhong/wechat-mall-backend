package service

import (
	"wechat-mall-backend/dbops"
	"wechat-mall-backend/defs"
	"wechat-mall-backend/model"
)

type IBrowseRecordService interface {
	AddBrowseRecord(record *model.WechatMallGoodsBrowseRecord)
	ListBrowseRecord(userId, page, size int) (*[]defs.PortalBrowseRecordVO, int)
}

type browseRecordService struct {
}

func NewBrowseRecordService() IBrowseRecordService {
	service := &browseRecordService{}
	return service
}

func (s *browseRecordService) AddBrowseRecord(record *model.WechatMallGoodsBrowseRecord) {
	recordDO, err := dbops.SelectGoodsBrowse(record.UserId, record.GoodsId)
	if err != nil {
		panic(err)
	}
	if recordDO.Id != 0 {
		err := dbops.DeleteBrowseRecordById(recordDO.Id)
		if err != nil {
			panic(err)
		}
	}
	err = dbops.InsertBrowseRecord(record)
	if err != nil {
		panic(err)
	}
}

func (s *browseRecordService) ListBrowseRecord(userId, page, size int) (*[]defs.PortalBrowseRecordVO, int) {
	records, err := dbops.SelectGoodsBrowseByUserId(userId, page, size)
	if err != nil {
		panic(err)
	}
	recordVOs := []defs.PortalBrowseRecordVO{}
	for _, recordDO := range *records {
		recordVO := defs.PortalBrowseRecordVO{}
		recordVO.Id = recordDO.Id
		recordVO.Picture = recordDO.Picture
		recordVO.Title = recordDO.Title
		recordVO.Price = recordDO.Price
		recordVO.CreateTime = recordDO.CreateTime
		recordVOs = append(recordVOs, recordVO)
	}
	total, err := dbops.CountGoodsBrowseByUserId(userId)
	if err != nil {
		panic(err)
	}
	return &recordVOs, total
}
