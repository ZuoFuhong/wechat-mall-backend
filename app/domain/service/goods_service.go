package service

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"strconv"
	"wechat-mall-backend/app/domain/entity"
	"wechat-mall-backend/app/domain/repository"
	"wechat-mall-backend/app/domain/view"
	"wechat-mall-backend/consts"
)

type IGoodsService interface {
	GetGoodsList(ctx context.Context, keyword string, categoryId, online, page, size int) ([]*entity.WechatMallGoodsDO, int, error)

	GetGoodsById(ctx context.Context, id int) (*entity.WechatMallGoodsDO, error)

	UpdateGoodsById(ctx context.Context, goods *entity.WechatMallGoodsDO) error

	AddGoods(ctx context.Context, goods *entity.WechatMallGoodsDO) (int, error)

	GetGoodsSpecList(ctx context.Context, goodsId int) ([]*view.CMSGoodsSpecVO, error)

	AddGoodsSpec(ctx context.Context, goodsId int, specList []int) error

	QueryPortalGoodsList(ctx context.Context, keyword string, sort, categoryId, page, size int) ([]*view.PortalGoodsListVO, int, error)

	QueryPortalGoodsDetail(ctx context.Context, goodsId int) (*view.PortalGoodsInfo, error)

	CountCategoryGoods(ctx context.Context, categoryId int) (int, error)

	CountGoodsSpecBySpecId(ctx context.Context, specId int) (int, error)
}

type goodsService struct {
	repos      repository.IGoodsRepos
	skuRepos   repository.IGoodsSkuRepos
	orderRepos repository.IOrderRepos
}

func NewGoodsService(repos repository.IGoodsRepos, skuRepos repository.IGoodsSkuRepos, orderRepos repository.IOrderRepos) IGoodsService {
	service := &goodsService{
		repos:      repos,
		skuRepos:   skuRepos,
		orderRepos: orderRepos,
	}
	return service
}

func (s *goodsService) GetGoodsList(ctx context.Context, keyword string, categoryId, online, page, size int) ([]*entity.WechatMallGoodsDO, int, error) {
	return s.repos.QueryGoodsList(ctx, keyword, "", categoryId, online, page, size)
}

func (s *goodsService) GetGoodsById(ctx context.Context, id int) (*entity.WechatMallGoodsDO, error) {
	return s.repos.QueryGoodsById(ctx, id)
}

func (s *goodsService) UpdateGoodsById(ctx context.Context, goods *entity.WechatMallGoodsDO) error {
	return s.repos.UpdateGoodsById(ctx, goods)
}

func (s *goodsService) AddGoods(ctx context.Context, goods *entity.WechatMallGoodsDO) (int, error) {
	return s.repos.AddGoods(ctx, goods)
}

func (s *goodsService) GetGoodsSpecList(ctx context.Context, goodsId int) ([]*view.CMSGoodsSpecVO, error) {
	specList, err := s.repos.GetGoodsSpecList(ctx, goodsId)
	if err != nil {
		return nil, err
	}
	specVOList := make([]*view.CMSGoodsSpecVO, 0)
	for _, v := range specList {
		specificationDO, err := s.skuRepos.QuerySpecificationById(ctx, v.SpecID)
		if err != nil {
			return nil, err
		}
		if specificationDO.ID == consts.ZERO || specificationDO.Del == consts.DELETE {
			return nil, errors.New("not found spec attr record")
		}
		attrList, err := s.skuRepos.QuerySpecificationAttrList(ctx, v.SpecID)
		if err != nil {
			return nil, err
		}
		attrVOList := make([]*view.CMSSpecificationAttrVO, 0)
		for _, item := range attrList {
			attrVo := &view.CMSSpecificationAttrVO{
				Id:     item.ID,
				SpecId: item.SpecID,
				Value:  item.Value,
				Extend: item.Extend,
			}
			attrVOList = append(attrVOList, attrVo)
		}
		specVO := &view.CMSGoodsSpecVO{
			SpecId:   v.SpecID,
			Name:     specificationDO.Name,
			AttrList: attrVOList,
		}
		specVOList = append(specVOList, specVO)
	}
	return specVOList, nil
}

func (s *goodsService) AddGoodsSpec(ctx context.Context, goodsId int, specList []int) error {
	err := s.repos.DeleteGoodsSpec(ctx, goodsId)
	if err != nil {
		return err
	}
	for _, specId := range specList {
		spec := &entity.WechatMallGoodsSpecDO{
			GoodsID: goodsId,
			SpecID:  specId,
		}
		if err := s.repos.AddGoodsSpec(ctx, spec); err != nil {
			return err
		}
	}
	return nil
}

func (s *goodsService) QueryPortalGoodsList(ctx context.Context, keyword string, sort, categoryId, page, size int) ([]*view.PortalGoodsListVO, int, error) {
	// 排序：0-综合 1-新品 2-销量 3-价格
	var order string
	switch sort {
	case 1:
		order = "create_time"
	case 2:
		order = "sale_num"
	case 3:
		order = "price"
	default:
		order = ""
	}
	goodsList, total, err := s.repos.QueryGoodsList(ctx, keyword, order, categoryId, consts.ONLINE, page, size)
	if err != nil {
		return nil, 0, err
	}
	voList := make([]*view.PortalGoodsListVO, 0)
	for _, v := range goodsList {
		humanNum, err := s.orderRepos.CountBuyGoodsUserNum(ctx, v.ID)
		if err != nil {
			return nil, 0, err
		}
		price, _ := strconv.ParseFloat(v.Price, 2)
		goodsVO := &view.PortalGoodsListVO{
			Id:       v.ID,
			Title:    v.Title,
			Price:    price,
			Picture:  v.Picture,
			HumanNum: humanNum,
		}
		voList = append(voList, goodsVO)
	}
	return voList, total, nil
}

func (s *goodsService) QueryPortalGoodsDetail(ctx context.Context, goodsId int) (*view.PortalGoodsInfo, error) {
	goodsDO, err := s.repos.QueryGoodsById(ctx, goodsId)
	if err != nil {
		return nil, err
	}
	if goodsDO.ID == consts.ZERO || goodsDO.Del == consts.DELETE || goodsDO.Online == consts.OFFLINE {
		return nil, errors.New("not found goods record")
	}
	skuDOList, _, err := s.skuRepos.GetSKUList(ctx, "", goodsId, consts.ONLINE, 0, 0)
	if err != nil {
		return nil, err
	}
	skuList := extractSkuVOList(skuDOList)
	specList, err := s.extraceSpecVOList(ctx, goodsId, skuDOList)
	if err != nil {
		return nil, err
	}
	price, _ := strconv.ParseFloat(goodsDO.Price, 2)
	goodsVO := &view.PortalGoodsInfo{
		Id:            goodsDO.ID,
		Title:         goodsDO.Title,
		Price:         price,
		Picture:       goodsDO.Picture,
		BannerPicture: goodsDO.BannerPicture,
		DetailPicture: goodsDO.DetailPicture,
		Tags:          goodsDO.Tags,
		Description:   "",
		SkuList:       skuList,
		SpecList:      specList,
	}
	return goodsVO, nil
}

func extractSkuVOList(skuDOList []*entity.WechatMallSkuDO) []*view.PortalSkuVO {
	voList := make([]*view.PortalSkuVO, 0)
	for _, v := range skuDOList {
		price, _ := strconv.ParseFloat(v.Price, 2)
		skuVO := &view.PortalSkuVO{
			Id:      v.ID,
			Picture: v.Picture,
			Title:   v.Title,
			Price:   price,
			Code:    v.Code,
			Stock:   v.Stock,
			Specs:   v.Specs,
		}
		voList = append(voList, skuVO)
	}
	return voList
}

func (s *goodsService) extraceSpecVOList(ctx context.Context, goodsId int, skuDOList []*entity.WechatMallSkuDO) ([]*view.PortalSpecVO, error) {
	specVOMap, specAttrVOMap := extraceSpecAttrVOList(skuDOList)
	specList, err := s.repos.GetGoodsSpecList(ctx, goodsId)
	if err != nil {
		return nil, err
	}
	specVOList := make([]*view.PortalSpecVO, 0)
	for _, v := range specList {
		specId := v.SpecID
		if specVOMap[specId] == "" {
			continue
		}
		specVO := &view.PortalSpecVO{
			SpecId:   specId,
			Name:     specVOMap[specId],
			AttrList: specAttrVOMap[specId],
		}
		specVOList = append(specVOList, specVO)
	}
	return specVOList, nil
}

func extraceSpecAttrVOList(skuDOList []*entity.WechatMallSkuDO) (map[int]string, map[int][]*view.PortalSpecAttrVO) {
	specVOMap := map[int]string{}
	specAttrVOMap := map[int][]*view.PortalSpecAttrVO{}
	for _, v := range skuDOList {
		// [{"key": "颜色", "value": "青芒色", "keyId": 1, "valueId": 42}, {"key": "尺寸", "value": "7英寸", "keyId": 2, "valueId": 5}]
		specs := make([]*entity.SkuSpecs, 0)
		if err := json.Unmarshal([]byte(v.Specs), &specs); err != nil {
			continue
		}
		for _, item := range specs {
			specName := specVOMap[item.KeyId]
			if specName == "" {
				specVOMap[item.KeyId] = item.Key
			}
			attrVOList := specAttrVOMap[item.KeyId]
			if attrVOList == nil {
				attrVOList = make([]*view.PortalSpecAttrVO, 0)
				specAttrVOMap[item.KeyId] = attrVOList
			}
			flag := false
			for _, attrVO := range attrVOList {
				if attrVO.AttrId == item.ValueId {
					flag = true
					break
				}
			}
			if flag {
				continue
			}
			attrVO := &view.PortalSpecAttrVO{
				AttrId: item.ValueId,
				Value:  item.Value,
			}
			attrVOList = append(attrVOList, attrVO)
			specAttrVOMap[item.KeyId] = attrVOList
		}
	}
	return specVOMap, specAttrVOMap
}

func (s *goodsService) CountCategoryGoods(ctx context.Context, categoryId int) (int, error) {
	return s.repos.CountCategoryGoods(ctx, categoryId)
}

func (s *goodsService) CountGoodsSpecBySpecId(ctx context.Context, specId int) (int, error) {
	return s.repos.CountGoodsSpecBySpecId(ctx, specId)
}
