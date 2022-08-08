package service

import (
	"context"
	"github.com/pkg/errors"
	"time"
	"wechat-mall-backend/app/domain/entity"
	"wechat-mall-backend/app/domain/repository"
	"wechat-mall-backend/app/domain/view"
	"wechat-mall-backend/consts"
	"wechat-mall-backend/pkg/utils"
)

type ICouponService interface {
	GetCouponList(ctx context.Context, page, size, online int) ([]*entity.WechatMallCouponDO, int, error)

	GetCouponById(ctx context.Context, id int) (*entity.WechatMallCouponDO, error)

	AddCoupon(ctx context.Context, coupon *entity.WechatMallCouponDO) error

	UpdateCouponById(ctx context.Context, coupon *entity.WechatMallCouponDO) error

	QueryCouponLogById(ctx context.Context, couponLogId int) (*entity.WechatMallCouponLogDO, error)

	RecordCouponLog(ctx context.Context, userId, couponId int) error

	QueryUserCoupon(ctx context.Context, userId, status, page, size int) ([]*view.PortalUserCouponVO, int, error)

	CountCouponTakeNum(ctx context.Context, userId, couponId int) (int, error)

	DoDeleteCouponLog(ctx context.Context, couponLog *entity.WechatMallCouponLogDO) error

	GetAllSubCategory(ctx context.Context) ([]*view.PortalCategoryVO, error)
}

type couponService struct {
	repos     repository.ICouponRepos
	cateRepos repository.ICategoryRepos
}

func NewCouponService(repos repository.ICouponRepos, cateRepos repository.ICategoryRepos) ICouponService {
	service := couponService{
		repos:     repos,
		cateRepos: cateRepos,
	}
	return &service
}

func (s *couponService) GetCouponList(ctx context.Context, page, size, online int) ([]*entity.WechatMallCouponDO, int, error) {
	return s.repos.QueryCouponList(ctx, page, size, online)
}

func (s *couponService) GetCouponById(ctx context.Context, id int) (*entity.WechatMallCouponDO, error) {
	return s.repos.QueryCouponById(ctx, id)
}

func (s *couponService) AddCoupon(ctx context.Context, coupon *entity.WechatMallCouponDO) error {
	return s.repos.AddCoupon(ctx, coupon)
}

func (s *couponService) UpdateCouponById(ctx context.Context, coupon *entity.WechatMallCouponDO) error {
	return s.repos.UpdateCouponById(ctx, coupon)
}

func (s *couponService) QueryCouponLogById(ctx context.Context, couponLogId int) (*entity.WechatMallCouponLogDO, error) {
	return s.repos.QueryCouponLogById(ctx, couponLogId)
}

func (s *couponService) RecordCouponLog(ctx context.Context, userId, couponId int) error {
	coupon, err := s.repos.QueryCouponById(ctx, couponId)
	if err != nil {
		return err
	}
	if coupon.ID == consts.ZERO || coupon.Del == consts.DELETE || coupon.Online == consts.OFFLINE {
		return errors.New("not found coupon record")
	}
	couponLog := &entity.WechatMallCouponLogDO{
		CouponID:   couponId,
		UserID:     userId,
		UseTime:    time.Now(),
		ExpireTime: coupon.EndTime,
		Status:     0,
		Code:       utils.RandomNumberStr(12),
	}
	return s.repos.AddCouponLog(ctx, couponLog)
}

func (s *couponService) QueryUserCoupon(ctx context.Context, userId, status, page, size int) ([]*view.PortalUserCouponVO, int, error) {
	// 刷新券的过期状态（优于调度任务）
	err := s.repos.UpdateCouponLogOverdueStatus(ctx, userId)
	if err != nil {
		return nil, 0, err
	}
	couponLogList, err := s.repos.QueryCouponLogList(ctx, userId, status, page, size)
	if err != nil {
		return nil, 0, err
	}
	voList := make([]*view.PortalUserCouponVO, 0)
	for _, v := range couponLogList {
		couponDO, err := s.repos.QueryCouponById(ctx, v.CouponID)
		if err != nil {
			return nil, 0, err
		}
		if couponDO.ID == consts.ZERO {
			return nil, 0, errors.New("not found coupon record")
		}
		couponVO := &view.PortalUserCouponVO{
			CLogId:      v.ID,
			CouponId:    v.CouponID,
			Title:       couponDO.Title,
			FullMoney:   couponDO.FullMoney,
			Minus:       couponDO.Minus,
			Rate:        couponDO.Rate,
			Type:        couponDO.Type,
			StartTime:   couponDO.StartTime.Format("2006-01-02 15:04:05"),
			EndTime:     v.ExpireTime.Format("2006-01-02 15:04:05"),
			Description: couponDO.Description,
		}
		voList = append(voList, couponVO)
	}
	total, err := s.repos.CountCouponTakeNum(ctx, userId, consts.ALL, status, 0)
	if err != nil {
		return nil, 0, err
	}
	return voList, total, nil
}

func (s *couponService) CountCouponTakeNum(ctx context.Context, userId, couponId int) (int, error) {
	return s.repos.CountCouponTakeNum(ctx, userId, couponId, consts.ALL, consts.ALL)
}

func (s *couponService) DoDeleteCouponLog(ctx context.Context, couponLog *entity.WechatMallCouponLogDO) error {
	couponLog.Del = 1
	return s.repos.UpdateCouponLogById(ctx, couponLog)
}

func (s *couponService) GetAllSubCategory(ctx context.Context) ([]*view.PortalCategoryVO, error) {
	categoryList, err := s.cateRepos.QueryAllSubCategory(ctx)
	if err != nil {
		return nil, err
	}
	voList := make([]*view.PortalCategoryVO, 0)
	for _, do := range categoryList {
		vo := &view.PortalCategoryVO{
			Id:   do.ID,
			Name: do.Name,
		}
		voList = append(voList, vo)
	}
	return voList, nil
}
