package service

import (
	"context"
	"encoding/json"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"log"
	"os"
	"strconv"
	"time"
	"wechat-mall-backend/app/domain/entity"
	"wechat-mall-backend/app/domain/repository"
	"wechat-mall-backend/app/domain/view"
	"wechat-mall-backend/consts"
	"wechat-mall-backend/pkg/config"
	"wechat-mall-backend/pkg/utils"
)

type IOrderService interface {
	GenerateOrder(ctx context.Context, userId, addressId, couponLogId int, dispatchAmount, expectAmount decimal.Decimal, goodsList []*entity.CartGoods) (*view.PortalPlaceOrderVO, error)

	QueryOrderList(ctx context.Context, userId, status, page, size int) ([]*view.PortalOrderListVO, int, error)

	QueryOrderDetail(ctx context.Context, userId int, orderNo string) (*view.PortalOrderDetailVO, error)

	OrderPaySuccessNotify(ctx context.Context, orderNo string) error

	QueryOrderSaleData(ctx context.Context, page, size int) ([]*entity.OrderSaleData, error)

	CountOrderNum(ctx context.Context, userId, status int) (int, error)

	CountPendingOrderRefund(ctx context.Context) (int, error)

	CancelOrder(ctx context.Context, userId, orderId int) error

	DeleteOrderRecord(ctx context.Context, userId, orderId int) error

	ConfirmTakeGoods(ctx context.Context, userId, orderId int) error

	RefundApply(ctx context.Context, userId int, orderNo, reason string) (string, error)

	QueryRefundDetail(ctx context.Context, userId int, refundNo string) (*view.OrderRefundDetailVO, error)

	UndoRefundApply(ctx context.Context, userId int, refundNo string) error

	QueryCMSOrderList(ctx context.Context, status, searchType int, keyword, startTime, endTime string, page, size int) ([]*view.CMSOrderInfoVO, int, error)

	ExportCMSOrderExcel(ctx context.Context, status, searchType int, keyword, startTime, endTime string) (string, error)

	QueryCMSOrderDetail(ctx context.Context, orderNo string) (*view.CMSOrderInfoVO, error)

	ModifyOrderStatus(ctx context.Context, orderNo string, otype int) error

	ModifyOrderRemark(ctx context.Context, orderNo, remark string) error

	ModifyOrderGoods(ctx context.Context, orderNo string, goodsId int, price string) error
}

type orderService struct {
	repos       repository.IOrderRepos
	goodsRepos  repository.IGoodsRepos
	userRepos   repository.IUserRepos
	cartRepos   repository.IUserCartRepos
	skuRepos    repository.IGoodsSkuRepos
	couponRepos repository.ICouponRepos
}

func NewOrderService(repos repository.IOrderRepos, goodsRepos repository.IGoodsRepos, userRepos repository.IUserRepos, cartRepos repository.IUserCartRepos, skuRepos repository.IGoodsSkuRepos, couponRepos repository.ICouponRepos) IOrderService {
	service := orderService{
		repos:       repos,
		goodsRepos:  goodsRepos,
		userRepos:   userRepos,
		cartRepos:   cartRepos,
		skuRepos:    skuRepos,
		couponRepos: couponRepos,
	}
	return &service
}

func (s *orderService) GenerateOrder(ctx context.Context, userId, addressId, couponLogId int, dispatchAmount, expectAmount decimal.Decimal,
	goodsList []*entity.CartGoods) (*view.PortalPlaceOrderVO, error) {
	goodsAmount, err := s.checkCartGoodsAndStock(ctx, goodsList)
	if err != nil {
		return nil, err
	}
	discountAmount, err := s.calcGoodsDiscountAmount(ctx, goodsAmount, userId, couponLogId)
	if err != nil {
		return nil, err
	}
	if !goodsAmount.Sub(discountAmount).Add(dispatchAmount).Equal(expectAmount) {
		return nil, errors.New("订单金额不符")
	}
	addressSnap, err := s.getAddressSnapshot(ctx, addressId)
	if err != nil {
		return nil, err
	}
	orderNo := time.Now().Format("20060102150405") + utils.RandomNumberStr(6)
	prepayId, err := s.generateWxpayPrepayId(ctx, orderNo, expectAmount.String())
	if err != nil {
		return nil, err
	}
	orderDO := &entity.WechatMallOrderDO{}
	orderDO.OrderNo = orderNo
	orderDO.UserID = userId
	orderDO.PayAmount = goodsAmount.Sub(discountAmount).Add(dispatchAmount).String()
	orderDO.GoodsAmount = goodsAmount.String()
	orderDO.DiscountAmount = discountAmount.String()
	orderDO.DispatchAmount = dispatchAmount.String()
	orderDO.PayTime = time.Unix(1136185445, 0)
	orderDO.DeliverTime = time.Unix(1136185445, 0)
	orderDO.FinishTime = time.Unix(1136185445, 0)
	orderDO.Status = 0
	orderDO.AddressID = addressId
	orderDO.AddressSnapshot = addressSnap
	orderDO.WxappPrepayID = prepayId
	if err := s.repos.AddOrder(ctx, orderDO); err != nil {
		return nil, err
	}
	if err := s.orderGoodsSnapshot(ctx, userId, orderNo, goodsList); err != nil {
		return nil, err
	}
	if err := s.clearUserCart(ctx, goodsList); err != nil {
		return nil, err
	}
	if err := s.couponCannel(ctx, couponLogId); err != nil {
		return nil, err
	}
	return &view.PortalPlaceOrderVO{OrderNo: orderNo, PrepayId: prepayId}, nil
}

// 检查-购物车以及商品的库存
func (s *orderService) checkCartGoodsAndStock(ctx context.Context, goodsList []*entity.CartGoods) (decimal.Decimal, error) {
	goodsAmount := decimal.NewFromInt(0)
	for _, v := range goodsList {
		if v.CartId != 0 {
			cartDO, err := s.cartRepos.SelectCartById(ctx, v.CartId)
			if err != nil {
				return goodsAmount, err
			}
			if cartDO.ID == consts.ZERO || cartDO.Del == consts.DELETE {
				return goodsAmount, errors.New("not found cart record")
			}
		}
		goodsDO, err := s.goodsRepos.QueryGoodsById(ctx, v.GoodsId)
		if err != nil {
			return goodsAmount, err

		}
		if goodsDO.ID == consts.ZERO || goodsDO.Del == consts.DELETE || goodsDO.Online == consts.OFFLINE {
			return goodsAmount, errors.New("商品下架，无法售出")
		}
		skuDO, err := s.skuRepos.GetSKUById(ctx, v.SkuId)
		if err != nil {
			return goodsAmount, err
		}
		if skuDO.ID == consts.ZERO || skuDO.Del == consts.DELETE || skuDO.Online == consts.OFFLINE {
			return goodsAmount, errors.New("商品下架，无法售出")
		}
		if skuDO.Stock < v.Num {
			return goodsAmount, errors.New("商品库存不足")
		}
		price, err := decimal.NewFromString(skuDO.Price)
		if err != nil {
			return goodsAmount, err
		}
		num := decimal.NewFromInt(int64(v.Num))
		goodsAmount = goodsAmount.Add(price.Mul(num))
	}
	return goodsAmount, nil
}

// 计算优惠金额
func (s *orderService) calcGoodsDiscountAmount(ctx context.Context, goodsAmount decimal.Decimal, userId, couponLogId int) (decimal.Decimal, error) {
	zero := decimal.NewFromInt(0)
	if couponLogId == 0 {
		return zero, nil
	}
	couponLog, err := s.couponRepos.QueryCouponLogById(ctx, couponLogId)
	if err != nil {
		return zero, err
	}
	if couponLog.CouponID == consts.ZERO || couponLog.Del == consts.DELETE || couponLog.Status != 0 || couponLog.UserID != userId {
		return zero, errors.New("无效的优惠券")
	}
	coupon, err := s.couponRepos.QueryCouponById(ctx, couponLog.CouponID)
	if err != nil {
		return zero, err
	}
	if coupon.ID == consts.ZERO {
		return zero, errors.New("not found coupon record")
	}
	var discountAmount decimal.Decimal
	switch coupon.Type {
	case 1:
		fullMoney, err := decimal.NewFromString(coupon.FullMoney)
		if err != nil {
			return zero, err
		}
		if goodsAmount.LessThan(fullMoney) {
			return zero, errors.New("未达到满减要求")
		}
		minus, err := decimal.NewFromString(coupon.Minus)
		if err != nil {
			return zero, err
		}
		discountAmount = minus
	case 2:
		rate, err := decimal.NewFromString(coupon.Rate)
		if err != nil {

			return zero, err

		}
		discountAmount = goodsAmount.Sub(goodsAmount.Mul(rate).Round(2))
	case 3:
		minus, err := decimal.NewFromString(coupon.Minus)
		if err != nil {
			return zero, err
		}
		discountAmount = minus
	case 4:
		fullMoney, err := decimal.NewFromString(coupon.FullMoney)
		if err != nil {
			return zero, err
		}
		if goodsAmount.LessThan(fullMoney) {
			return zero, errors.New("未达到满减要求")
		}
		rate, err := decimal.NewFromString(coupon.Rate)
		if err != nil {
			return zero, err

		}
		discountAmount = goodsAmount.Sub(goodsAmount.Mul(rate).Round(2))
	default:
		discountAmount = decimal.NewFromInt(0)
	}
	if discountAmount.GreaterThan(goodsAmount) {
		discountAmount = goodsAmount
	}
	return discountAmount, nil
}

func (s *orderService) getAddressSnapshot(ctx context.Context, addressId int) (string, error) {
	addressDO, err := s.userRepos.QueryUserAddressById(ctx, addressId)
	if err != nil {
		return "", err
	}
	if addressDO.ID == consts.ZERO || addressDO.Del == consts.DELETE {
		return "", errors.New("not found address record")
	}
	snapshot := &view.AddressSnapshot{}
	snapshot.Contacts = addressDO.Contacts
	snapshot.Mobile = addressDO.Mobile
	snapshot.ProvinceId = addressDO.ProvinceID
	snapshot.ProvinceStr = addressDO.ProvinceStr
	snapshot.CityId = addressDO.CityID
	snapshot.CityStr = addressDO.CityStr
	snapshot.AreaStr = addressDO.AreaStr
	snapshot.Address = addressDO.Address
	bytes, err := json.Marshal(snapshot)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// 支付-获取微信预支付ID
func (s *orderService) generateWxpayPrepayId(ctx context.Context, orderNo string, payAmount string) (string, error) {
	// todo: 请求微信支付订单
	return "prepay_id:" + orderNo, nil
}

// 订单详情-快照
func (s *orderService) orderGoodsSnapshot(ctx context.Context, userId int, orderNo string, goodsList []*entity.CartGoods) error {
	for _, v := range goodsList {
		goodsDO, _ := s.goodsRepos.QueryGoodsById(ctx, v.GoodsId)
		skuDO, _ := s.skuRepos.GetSKUById(ctx, v.SkuId)

		orderGoodsDO := &entity.WechatMallOrderGoodsDO{}
		orderGoodsDO.OrderNo = orderNo
		orderGoodsDO.UserID = userId
		orderGoodsDO.GoodsID = v.GoodsId
		orderGoodsDO.SkuID = v.SkuId
		orderGoodsDO.Picture = skuDO.Picture
		orderGoodsDO.Title = goodsDO.Title
		orderGoodsDO.Price = skuDO.Price
		orderGoodsDO.Specs = skuDO.Specs
		orderGoodsDO.Num = v.Num
		orderGoodsDO.LockStatus = 0

		if err := s.repos.AddOrderGoods(ctx, orderGoodsDO); err != nil {
			return err
		}
		// 减库存
		if err := s.skuRepos.UpdateSkuStockById(ctx, v.SkuId, v.Num); err != nil {
			return err
		}
		// 商品销量
		if err := s.goodsRepos.UpdateGoodsSaleNum(ctx, v.GoodsId, v.Num); err != nil {
			return err
		}
	}
	return nil
}

// 下单成功-清理购物车
func (s *orderService) clearUserCart(ctx context.Context, goodsList []*entity.CartGoods) error {
	for _, v := range goodsList {
		if v.CartId != 0 {
			cartDO, _ := s.cartRepos.SelectCartById(ctx, v.CartId)
			if cartDO.ID == consts.ZERO || cartDO.Del == consts.DELETE {
				continue
			}
			cartDO.Del = consts.DELETE

			if err := s.cartRepos.UpdateCartById(ctx, cartDO); err != nil {
				return err
			}
		}
	}
	return nil
}

// 优惠券-核销
func (s *orderService) couponCannel(ctx context.Context, couponLogId int) error {
	if couponLogId == 0 {
		return nil
	}
	couponLogDO, _ := s.couponRepos.QueryCouponLogById(ctx, couponLogId)
	if couponLogDO.ID == 0 {
		return errors.New("not found coupon log record")
	}
	couponLogDO.Status = 1
	couponLogDO.UseTime = time.Unix(1136185445, 0)

	if err := s.couponRepos.UpdateCouponLogById(ctx, couponLogDO); err != nil {
		return err
	}
	return nil
}

func (s *orderService) QueryOrderList(ctx context.Context, userId, status, page, size int) ([]*view.PortalOrderListVO, int, error) {
	orderList, err := s.repos.ListOrderByParams(ctx, userId, status, page, size)
	if err != nil {
		return nil, 0, err
	}
	total, err := s.repos.CountOrderByParams(ctx, userId, status)
	if err != nil {

		return nil, 0, err

	}
	voList := make([]*view.PortalOrderListVO, 0)
	for _, v := range orderList {
		price, _ := strconv.ParseFloat(v.PayAmount, 2)
		goodsList, goodsNum, err := s.extraceOrderGoods(ctx, v.OrderNo)
		if err != nil {
			return nil, 0, err
		}
		orderVO := &view.PortalOrderListVO{
			Id:        v.ID,
			OrderNo:   v.OrderNo,
			PlaceTime: utils.FormatTime(v.CreateTime),
			PayAmount: price,
			Status:    v.Status,
			GoodsList: goodsList,
			GoodsNum:  goodsNum,
		}
		voList = append(voList, orderVO)
	}
	return voList, total, nil
}

func (s *orderService) extraceOrderGoods(ctx context.Context, orderNo string) ([]*view.PortalOrderGoodsVO, int, error) {
	orderGoodsList, err := s.repos.QueryOrderGoods(ctx, orderNo)
	if err != nil {
		return nil, 0, err
	}
	goodsNum := 0
	voList := make([]*view.PortalOrderGoodsVO, 0)
	for _, v := range orderGoodsList {
		specList := make([]*entity.SkuSpecs, 0)
		if err := json.Unmarshal([]byte(v.Specs), &specList); err != nil {
			return nil, 0, err
		}
		specs := ""
		for _, v := range specList {
			specs += v.Value + "; "
		}
		if len(specs) > 2 {
			specs = specs[0 : len(specs)-2]
		}
		price, _ := strconv.ParseFloat(v.Price, 2)
		goodsVO := &view.PortalOrderGoodsVO{
			GoodsId: v.GoodsID,
			Title:   v.Title,
			Price:   price,
			Picture: v.Picture,
			SkuId:   v.SkuID,
			Specs:   specs,
			Num:     v.Num,
		}
		voList = append(voList, goodsVO)
		goodsNum += v.Num
	}
	return voList, goodsNum, nil
}

func (s *orderService) QueryOrderDetail(ctx context.Context, userId int, orderNo string) (*view.PortalOrderDetailVO, error) {
	orderDO, err := s.repos.QueryOrderByOrderNo(ctx, orderNo)
	if err != nil {
		return nil, err
	}
	if orderDO.ID == consts.ZERO || orderDO.Del == consts.DELETE {
		return nil, errors.New("not found order record")
	}
	if orderDO.UserID != userId {
		return nil, errors.New("not found order record")
	}
	// 订单信息
	snapshot := new(view.AddressSnapshot)
	if err = json.Unmarshal([]byte(orderDO.AddressSnapshot), snapshot); err != nil {
		return nil, err
	}
	orderGoods, orderGoodsNum, err := s.extraceOrderGoods(ctx, orderDO.OrderNo)
	if err != nil {
		return nil, err
	}
	// 退款信息
	refundDO, err := s.repos.QueryOrderRefundRecord(ctx, orderNo)
	if err != nil {
		return nil, err
	}
	refundApply := &view.OrderRefundApplyVO{RefundNo: refundDO.RefundNo}
	orderVO := &view.PortalOrderDetailVO{}
	orderVO.Id = orderDO.ID
	orderVO.OrderNo = orderDO.OrderNo
	orderVO.GoodsAmount, _ = strconv.ParseFloat(orderDO.GoodsAmount, 2)
	orderVO.DiscountAmount, _ = strconv.ParseFloat(orderDO.DiscountAmount, 2)
	orderVO.DispatchAmount, _ = strconv.ParseFloat(orderDO.DispatchAmount, 2)
	orderVO.PayAmount, _ = strconv.ParseFloat(orderDO.PayAmount, 2)
	orderVO.Status = orderDO.Status
	orderVO.PlaceTime = utils.FormatTime(orderDO.CreateTime)
	orderVO.PayTime = utils.FormatTime(orderDO.PayTime)
	orderVO.DeliverTime = utils.FormatTime(orderDO.DeliverTime)
	orderVO.FinishTime = utils.FormatTime(orderDO.FinishTime)
	orderVO.GoodsList = orderGoods
	orderVO.GoodsNum = orderGoodsNum
	orderVO.Address = snapshot
	orderVO.RefundApply = refundApply
	return orderVO, nil
}

func (s *orderService) OrderPaySuccessNotify(ctx context.Context, orderNo string) error {
	orderDO, err := s.repos.QueryOrderByOrderNo(ctx, orderNo)
	if err != nil {
		return err
	}
	if orderDO.ID == consts.ZERO || orderDO.Del == consts.DELETE {

		return errors.New("not found order record")
	}
	if orderDO.Status != 0 {
		log.Printf("orderNo = %v 重复回调", orderDO)
		return nil
	}
	orderDO.Status = 1
	orderDO.PayTime = time.Unix(1136185445, 0)
	return s.repos.UpdateOrderById(ctx, orderDO)
}

func (s *orderService) QueryOrderSaleData(ctx context.Context, page, size int) ([]*entity.OrderSaleData, error) {
	return s.repos.QueryOrderSaleData(ctx, page, size)
}

func (s *orderService) CountOrderNum(ctx context.Context, userId, status int) (int, error) {
	return s.repos.CountOrderNum(ctx, userId, status)
}

func (s *orderService) CountPendingOrderRefund(ctx context.Context) (int, error) {
	return s.repos.CountPendingOrderRefund(ctx)
}

func (s *orderService) CancelOrder(ctx context.Context, userId, orderId int) error {
	orderDO, err := s.repos.QueryOrderById(ctx, orderId)
	if err != nil {
		return err
	}
	if orderDO.ID == consts.ZERO || orderDO.Del == consts.DELETE {
		return errors.New("not found order record")
	}
	if orderDO.UserID != userId {
		return errors.New("非法操作")

	}
	if orderDO.Status != 0 {
		return errors.New("进行中的订单，无法取消")
	}
	orderDO.Status = -1
	orderDO.FinishTime = time.Unix(1136185445, 0)
	err = s.repos.UpdateOrderById(ctx, orderDO)
	if err != nil {
		return err
	}
	return s.orderStockRollback(ctx, orderDO.OrderNo)
}

// 订单-库存回滚
// 场景：取消订单（手动取消、超时未支付）、订单退款
func (s *orderService) orderStockRollback(ctx context.Context, orderNo string) error {
	orderGoods, err := s.repos.QueryOrderGoods(ctx, orderNo)
	if err != nil {
		return err
	}
	for _, v := range orderGoods {
		if err := s.skuRepos.UpdateSkuStockById(ctx, v.SkuID, -v.Num); err != nil {
			return err
		}
		if err = s.goodsRepos.UpdateGoodsSaleNum(ctx, v.GoodsID, -v.Num); err != nil {
			return err
		}
	}
	return nil
}

func (s *orderService) DeleteOrderRecord(ctx context.Context, userId, orderId int) error {
	orderDO, err := s.repos.QueryOrderById(ctx, orderId)
	if err != nil {
		return err
	}
	if orderDO.ID == consts.ZERO || orderDO.Del == consts.DELETE {
		return errors.New("not found order record")
	}
	if orderDO.UserID != userId {
		return errors.New("非法操作")
	}
	if orderDO.Status == -1 || orderDO.Status == 3 || orderDO.Status == 5 {
		orderDO.Del = 1
		if err := s.repos.UpdateOrderById(ctx, orderDO); err != nil {
			return err
		}
	} else {
		return errors.New("进行中的订单，无法删除")
	}
	return nil
}

func (s *orderService) ConfirmTakeGoods(ctx context.Context, userId, orderId int) error {
	orderDO, err := s.repos.QueryOrderById(ctx, orderId)
	if err != nil {
		return err
	}
	if orderDO.ID == consts.ZERO || orderDO.Del == consts.DELETE {
		return errors.New("not found order record")

	}
	if orderDO.UserID != userId {
		return errors.New("非法操作")
	}
	if orderDO.Status == 2 {
		orderDO.Status = 3
		orderDO.FinishTime = time.Unix(1136185445, 0)

		if err := s.repos.UpdateOrderById(ctx, orderDO); err != nil {
			return err

		}
	}
	return nil
}

func (s *orderService) RefundApply(ctx context.Context, userId int, orderNo, reason string) (string, error) {
	orderDO, err := s.repos.QueryOrderByOrderNo(ctx, orderNo)
	if err != nil {
		return "", err
	}
	if orderDO.ID == consts.ZERO || orderDO.Del == consts.DELETE {
		return "", errors.New("not found order record")
	}
	if orderDO.UserID != userId {
		return "", errors.New("非法操作")

	}
	if orderDO.Status != 1 {
		return "", errors.New("非法操作")
	}
	refundNo := time.Now().Format("20060102150405") + utils.RandomNumberStr(6)
	refund := &entity.WechatMallOrderRefund{}
	refund.RefundNo = refundNo
	refund.UserID = userId
	refund.OrderNo = orderNo
	refund.Reason = reason
	refund.RefundTime = time.Unix(1136185445, 0)
	refund.RefundAmount = orderDO.PayAmount

	if err = s.repos.AddRefundRecord(ctx, refund); err != nil {
		return "", err
	}
	orderDO.Status = 4
	err = s.repos.UpdateOrderById(ctx, orderDO)
	if err != nil {
		return "", err
	}
	return refundNo, nil
}

func (s *orderService) QueryRefundDetail(ctx context.Context, userId int, refundNo string) (*view.OrderRefundDetailVO, error) {
	refundDO, err := s.repos.QueryRefundRecord(ctx, refundNo)
	if err != nil {
		return nil, err
	}
	if refundDO.ID == consts.ZERO || refundDO.Del == consts.DELETE {
		return nil, errors.New("not found order refund record")
	}
	if refundDO.UserID != userId {
		return nil, errors.New("非法操作")
	}
	goodsList, _, err := s.extraceOrderGoods(ctx, refundDO.OrderNo)
	if err != nil {
		return nil, err
	}
	refundAmount, _ := strconv.ParseFloat(refundDO.RefundAmount, 2)
	refundVO := &view.OrderRefundDetailVO{
		RefundNo:     refundDO.RefundNo,
		Reason:       refundDO.Reason,
		RefundAmount: refundAmount,
		Status:       refundDO.Status,
		ApplyTime:    utils.FormatTime(refundDO.CreateTime),
		RefundTime:   utils.FormatTime(refundDO.RefundTime),
		GoodsList:    goodsList,
	}
	return refundVO, nil
}

func (s *orderService) UndoRefundApply(ctx context.Context, userId int, refundNo string) error {
	refundDO, err := s.repos.QueryRefundRecord(ctx, refundNo)
	if err != nil {
		return err
	}
	if refundDO.ID == consts.ZERO || refundDO.Del == consts.DELETE {
		return errors.New("not found order refund record")
	}
	if refundDO.UserID != userId {
		return errors.New("非法操作")

	}
	if refundDO.Status != 0 {
		return errors.New("状态异常")
	}
	// 订单：待收货
	orderDO, err := s.repos.QueryOrderByOrderNo(ctx, refundDO.OrderNo)
	if err != nil {
		return err
	}
	orderDO.Status = 1
	if err = s.repos.UpdateOrderById(ctx, orderDO); err != nil {
		return err
	}
	if err = s.repos.UpdateRefundApply(ctx, refundDO.ID, 2); err != nil {
		return err

	}
	return nil
}

func (s *orderService) QueryCMSOrderList(ctx context.Context, status, searchType int, keyword, startTime, endTime string, page, size int) ([]*view.CMSOrderInfoVO, int, error) {
	orderList, err := s.repos.SelectCMSOrderList(ctx, status, searchType, keyword, startTime, endTime, page, size)
	if err != nil {
		return nil, 0, err
	}
	total, err := s.repos.SelectCMSOrderNum(ctx, status, searchType, keyword, startTime, endTime)
	if err != nil {
		return nil, 0, err
	}
	orderVOList := make([]*view.CMSOrderInfoVO, 0)
	for _, v := range orderList {
		address := extractOrderAddress(v.AddressSnapshot)
		buyer, err := s.extractOrderBuyer(ctx, v.UserID)
		if err != nil {
			return nil, 0, err
		}
		orderVO := &view.CMSOrderInfoVO{}
		orderVO.OrderNo = v.OrderNo
		orderVO.PlaceTime = v.CreateTime.Format("2006-01-02 15:04:05")
		orderVO.Address = address
		orderVO.PayAmount, _ = strconv.ParseFloat(v.PayAmount, 2)
		orderVO.GoodsAmount, _ = strconv.ParseFloat(v.GoodsAmount, 2)
		orderVO.DiscountAmount, _ = strconv.ParseFloat(v.DiscountAmount, 2)
		orderVO.DispatchAmount, _ = strconv.ParseFloat(v.DispatchAmount, 2)
		orderVO.Status = v.Status
		orderVO.TransactionId = v.TransactionID
		orderVO.Remark = v.Remark
		orderVO.PayTime = v.PayTime.Format("2006-01-02 15:04:05")
		orderVO.DeliverTime = v.DeliverTime.Format("2006-01-02 15:04:05")
		orderVO.FinishTime = v.FinishTime.Format("2006-01-02 15:04:05")
		orderVO.Buyer = buyer
		orderVOList = append(orderVOList, orderVO)
	}
	return orderVOList, total, nil
}

func (s *orderService) ExportCMSOrderExcel(ctx context.Context, status, searchType int, keyword, startTime, endTime string) (string, error) {
	orderList, err := s.repos.SelectCMSOrderList(ctx, status, searchType, keyword, startTime, endTime, 0, 0)
	if err != nil {
		return "", err
	}
	excelData, err := s.extractOrderExcelData(ctx, orderList)
	if err != nil {
		return "", err
	}
	filepath, filename := generateExcel(excelData)
	ossLink, err := uploadFileToOSS(filepath, filename)
	if err != nil {
		return "", err
	}
	_ = os.Remove(filepath)
	return ossLink, nil
}

func uploadFileToOSS(filepath, filename string) (string, error) {
	ossConf := config.GlobalConfig().Oss
	endpoint := "https://oss-cn-hangzhou.aliyuncs.com"
	client, err := oss.New(endpoint, ossConf.AccessKeyId, ossConf.AccessKeySecret)
	if err != nil {
		return "", err
	}
	bucket, err := client.Bucket(ossConf.BucketName)
	if err != nil {
		return "", err

	}
	objectKey := "order-excel/" + filename
	if err := bucket.PutObjectFromFile(objectKey, filepath); err != nil {
		return "", err
	}
	ossLink := "https://" + ossConf.BucketName + ".oss-cn-hangzhou.aliyuncs.com/" + objectKey
	return ossLink, nil
}

// 订单-提取Excel数据
func (s *orderService) extractOrderExcelData(ctx context.Context, orderList []*entity.WechatMallOrderDO) (map[string]string, error) {
	excelData := map[string]string{}
	excelData["A1"] = "订单号"
	excelData["B1"] = "微信支付单号"
	excelData["C1"] = "下单时间"
	excelData["D1"] = "订单金额"
	excelData["E1"] = "实付金额"
	excelData["F1"] = "订单状态"
	excelData["G1"] = "买家信息"
	excelData["H1"] = "收货地址"
	excelData["I1"] = "备注"
	excelData["J1"] = "商品名称"
	excelData["K1"] = "规格"
	excelData["L1"] = "数量"
	rowNum := 1
	for _, v := range orderList {
		goodsList, err := s.extractOrderGoodsVO(ctx, v.OrderNo)
		if err != nil {
			return nil, err
		}
		goodsAmount, _ := decimal.NewFromString(v.GoodsAmount)
		dispatchAmount, _ := decimal.NewFromString(v.DispatchAmount)
		orderAmount := goodsAmount.Add(dispatchAmount).String()
		statusStr := extractOrderStatus(v.Status)
		addressSnap := extractOrderAddress(v.AddressSnapshot)

		for _, g := range goodsList {
			rowNum += 1
			rowNumStr := strconv.Itoa(rowNum)
			excelData["A"+rowNumStr] = v.OrderNo
			excelData["B"+rowNumStr] = v.TransactionID
			excelData["C"+rowNumStr] = v.CreateTime.Format("2006-01-02 15:04:05")
			excelData["D"+rowNumStr] = orderAmount
			excelData["E"+rowNumStr] = v.PayAmount
			excelData["F"+rowNumStr] = statusStr
			excelData["G"+rowNumStr] = addressSnap.Contacts + "（" + addressSnap.Mobile + "）"
			excelData["H"+rowNumStr] = addressSnap.ProvinceStr + addressSnap.CityStr + addressSnap.AreaStr + addressSnap.Address
			excelData["I"+rowNumStr] = v.Remark
			excelData["J"+rowNumStr] = g.Title
			excelData["K"+rowNumStr] = g.Specs
			excelData["L"+rowNumStr] = strconv.Itoa(g.Num)
		}
	}
	return excelData, nil
}

// 生成Excel文件，保存在本地
func generateExcel(excelData map[string]string) (string, string) {
	xlsx := excelize.NewFile()
	sheet := xlsx.NewSheet("sheet1")
	for k, v := range excelData {
		xlsx.SetCellValue("sheet1", k, v)
	}
	xlsx.SetActiveSheet(sheet)
	filename := "订单列表_" + utils.FormatDatetime(time.Now(), utils.YYYYMMDDHHMMSS) + ".xlsx"
	filepath := "/tmp/wechat-mall/" + filename
	_ = utils.CheckFileDirExists(filepath)
	err := xlsx.SaveAs(filepath)
	if err != nil {
		// ignore error
	}
	return filepath, filename
}

// 订单-状态
func extractOrderStatus(status int) string {
	statusStr := ""
	switch status {
	case 0:
		statusStr = "待付款"
	case 1:
		statusStr = "待发货"
	case 2:
		statusStr = "待收货"
	case 3:
		statusStr = "已完成"
	}
	return statusStr
}

// 订单-商品规格
func extractOrderGoodsSpecs(specs string) string {
	specList := make([]*entity.SkuSpecs, 0)
	if err := json.Unmarshal([]byte(specs), &specList); err != nil {
		return ""
	}
	specStr := ""
	for _, v := range specList {
		specStr += v.Value + ";"
	}
	return specStr
}

func (s *orderService) QueryCMSOrderDetail(ctx context.Context, orderNo string) (*view.CMSOrderInfoVO, error) {
	orderDO, err := s.repos.QueryOrderByOrderNo(ctx, orderNo)
	if err != nil {
		return nil, err
	}
	goodsList, err := s.extractOrderGoodsVO(ctx, orderNo)
	if err != nil {
		return nil, err
	}
	address := extractOrderAddress(orderDO.AddressSnapshot)
	buyer, err := s.extractOrderBuyer(ctx, orderDO.UserID)
	if err != nil {
		return nil, err
	}

	orderVO := &view.CMSOrderInfoVO{}
	orderVO.OrderNo = orderDO.OrderNo
	orderVO.PlaceTime = orderDO.CreateTime.Format("2006-01-02 15:04:05")
	orderVO.Address = address
	orderVO.PayAmount, _ = strconv.ParseFloat(orderDO.PayAmount, 2)
	orderVO.GoodsAmount, _ = strconv.ParseFloat(orderDO.GoodsAmount, 2)
	orderVO.DiscountAmount, _ = strconv.ParseFloat(orderDO.DiscountAmount, 2)
	orderVO.DispatchAmount, _ = strconv.ParseFloat(orderDO.DispatchAmount, 2)
	orderVO.Status = orderDO.Status
	orderVO.TransactionId = orderDO.TransactionID
	orderVO.Remark = orderDO.Remark
	orderVO.PayTime = orderDO.PayTime.Format("2006-01-02 15:04:05")
	orderVO.DeliverTime = orderDO.DeliverTime.Format("2006-01-02 15:04:05")
	orderVO.FinishTime = orderDO.FinishTime.Format("2006-01-02 15:04:05")
	orderVO.Buyer = buyer
	orderVO.GoodsList = goodsList
	return orderVO, nil
}

// 提取订单中的商品
func (s *orderService) extractOrderGoodsVO(ctx context.Context, orderNo string) ([]*view.CMSOrderGoodsVO, error) {
	goodsDOList, err := s.repos.QueryOrderGoods(ctx, orderNo)
	if err != nil {
		return nil, err
	}
	goodsVOList := make([]*view.CMSOrderGoodsVO, 0)
	for _, v := range goodsDOList {
		specs := extractOrderGoodsSpecs(v.Specs)

		goodsVO := &view.CMSOrderGoodsVO{}
		goodsVO.Picture = v.Picture
		goodsVO.Title = v.Title
		goodsVO.Price, _ = strconv.ParseFloat(v.Price, 2)
		goodsVO.Specs = specs
		goodsVO.Num = v.Num
		goodsVOList = append(goodsVOList, goodsVO)
	}
	return goodsVOList, nil
}

// 订单-提取收货人
func extractOrderAddress(addressSnapshot string) *view.AddressSnapshot {
	address := new(view.AddressSnapshot)
	if err := json.Unmarshal([]byte(addressSnapshot), address); err != nil {
		// ignore error
	}
	return address
}

// 订单-提取买家信息
func (s *orderService) extractOrderBuyer(ctx context.Context, uid int) (*view.BasicUser, error) {
	userDO, err := s.userRepos.GetUserById(ctx, uid)
	if err != nil {
		return nil, err
	}
	basicUser := &view.BasicUser{
		UserId:   userDO.ID,
		Nickname: userDO.Nickname,
		Avatar:   userDO.Avatar,
	}
	return basicUser, nil
}

func (s *orderService) ModifyOrderStatus(ctx context.Context, orderNo string, otype int) error {
	/*
		待发货：1-确认发货
		待收货：2-确认收货
		待付款：3-确认付款
	*/
	orderDO, err := s.repos.QueryOrderByOrderNo(ctx, orderNo)
	if err != nil {
		return err
	}
	newStatus := 0
	switch otype {
	case 1:
		if orderDO.Status != 1 {
			return errors.New("非法操作")
		}
		newStatus = 2
	case 2:
		if orderDO.Status != 2 {
			return errors.New("非法操作")
		}
		newStatus = 3
	case 3:
		if orderDO.Status != 0 {
			return errors.New("非法操作")
		}
		newStatus = 1
	default:
		return errors.New("非法操作")
	}
	params := &entity.WechatMallOrderDO{}
	params.ID = orderDO.ID
	params.Status = newStatus
	return s.repos.UpdateOrderById(ctx, params)
}

func (s *orderService) ModifyOrderRemark(ctx context.Context, orderNo, remark string) error {
	orderDO, err := s.repos.QueryOrderByOrderNo(ctx, orderNo)
	if err != nil {
		return err
	}
	return s.repos.UpdateOrderRemark(ctx, orderDO.ID, remark)
}

func (s *orderService) ModifyOrderGoods(ctx context.Context, orderNo string, goodsId int, price string) error {
	orderDO, err := s.repos.QueryOrderByOrderNo(ctx, orderNo)
	if err != nil {
		return err
	}
	if orderDO.Status != 0 {
		return errors.New("非法操作")
	}
	goodsList, err := s.repos.QueryOrderGoods(ctx, orderNo)
	if err != nil {
		return err
	}
	orderGoods := &entity.WechatMallOrderGoodsDO{}
	for _, v := range goodsList {
		if v.GoodsID == goodsId {
			orderGoods = v
		}
	}
	if orderGoods.ID == 0 {
		return errors.New("订单商品异常")
	}
	// 更新商品价格
	params := &entity.WechatMallOrderGoodsDO{}
	params.ID = orderGoods.ID
	params.Price = price

	if err = s.repos.UpdateOrderGoods(ctx, params); err != nil {
		return err
	}
	// 订单金额，计算差价
	newGoodsPrice, _ := decimal.NewFromString(price)
	oldGoodsPrice, _ := decimal.NewFromString(orderGoods.Price)
	diffAmount := newGoodsPrice.Sub(oldGoodsPrice).Mul(decimal.NewFromInt(int64(orderGoods.Num)))
	payAmount, _ := decimal.NewFromString(orderDO.PayAmount)
	goodsAmount, _ := decimal.NewFromString(orderDO.GoodsAmount)
	orderParams := &entity.WechatMallOrderDO{}
	orderParams.ID = orderDO.ID
	orderParams.PayAmount = payAmount.Add(diffAmount).String()
	orderParams.GoodsAmount = goodsAmount.Add(diffAmount).String()
	return s.repos.UpdateOrderById(ctx, orderParams)
}
