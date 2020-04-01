package service

import (
	"encoding/json"
	"github.com/shopspring/decimal"
	"log"
	"time"
	"wechat-mall-backend/dbops"
	"wechat-mall-backend/defs"
	"wechat-mall-backend/errs"
	"wechat-mall-backend/model"
	"wechat-mall-backend/utils"
)

type IOrderService interface {
	GenerateOrder(userId, addressId, couponLogId int, dispatchAmount, expectAmount decimal.Decimal, goodsList []defs.PortalCartGoods) string
	QueryOrderList(userId, status, page, size int) (*[]defs.PortalOrderListVO, int)
	QueryOrderDetail(userId, orderId int) *defs.PortalOrderDetailVO
	OrderPaySuccessNotify(orderNo string)
	QueryOrderSaleData(page, size int) *[]defs.OrderSaleData
	CountWaitingOrderNum(status int) int
	CountPendingOrderRefund() int
}

type orderService struct {
}

func NewOrderService() IOrderService {
	service := orderService{}
	return &service
}

func (s *orderService) GenerateOrder(userId, addressId, couponLogId int, dispatchAmount, expectAmount decimal.Decimal,
	goodsList []defs.PortalCartGoods) string {
	goodsAmount := checkGoodsStock(goodsList)
	discountAmount := calcGoodsDiscountAmount(goodsAmount, userId, couponLogId)
	if !goodsAmount.Sub(discountAmount).Add(dispatchAmount).Equal(expectAmount) {
		panic(errs.NewErrorOrder("订单金额不符！"))
	}
	addressSnap := getAddressSnapshot(addressId)
	orderNo := time.Now().Format("20060102150405") + utils.RandomNumberStr(6)
	prepayId := s.generateWxpayPrepayId(orderNo, expectAmount.String())

	orderDO := model.WechatMallOrderDO{}
	orderDO.OrderNo = orderNo
	orderDO.UserId = userId
	orderDO.PayAmount = goodsAmount.Sub(discountAmount).Add(dispatchAmount).String()
	orderDO.GoodsAmount = goodsAmount.String()
	orderDO.DiscountAmount = discountAmount.String()
	orderDO.DispatchAmount = dispatchAmount.String()
	orderDO.PayTime = time.Now().Format("2006-01-02 15:04:05")
	orderDO.Status = 0
	orderDO.AddressId = addressId
	orderDO.AddressSnapshot = addressSnap
	orderDO.WxappPrePayId = prepayId
	err := dbops.AddOrder(&orderDO)
	if err != nil {
		panic(err)
	}
	orderGoodsSnapshot(orderNo, goodsList)
	clearUserCart(goodsList)
	couponCannel(couponLogId)
	return prepayId
}

func checkGoodsStock(goodsList []defs.PortalCartGoods) decimal.Decimal {
	goodsAmount := decimal.NewFromInt(0)
	for _, v := range goodsList {
		cartDO, err := dbops.SelectCartById(v.CartId)
		if err != nil {
			panic(err)
		}
		if cartDO.Id == defs.ZERO || cartDO.Del == defs.DELETE {
			panic(errs.ErrorGoodsCart)
		}
		goodsDO, err := dbops.QueryGoodsById(cartDO.GoodsId)
		if err != nil {
			panic(err)
		}
		if goodsDO.Id == defs.ZERO || goodsDO.Del == defs.DELETE || goodsDO.Online == defs.OFFLINE {
			panic(errs.NewErrorOrder("商品下架，无法售出"))
		}
		skuDO, err := dbops.GetSKUById(cartDO.SkuId)
		if err != nil {
			panic(err)
		}
		if skuDO.Id == defs.ZERO || skuDO.Del == defs.DELETE || skuDO.Online == defs.OFFLINE {
			panic(errs.NewErrorOrder("商品下架，无法售出"))
		}
		if skuDO.Stock < v.Num {
			panic(errs.NewErrorOrder("商品库存不足！"))
		}
		price, err := decimal.NewFromString(skuDO.Price)
		if err != nil {
			panic(err)
		}
		num := decimal.NewFromInt(int64(v.Num))
		goodsAmount = goodsAmount.Add(price.Mul(num))
	}
	return goodsAmount
}

// 计算优惠金额
func calcGoodsDiscountAmount(goodsAmount decimal.Decimal, userId, couponLogId int) decimal.Decimal {
	if couponLogId == 0 {
		return decimal.NewFromInt(0)
	}
	couponLog, err := dbops.QueryCouponLogById(couponLogId)
	if err != nil {
		panic(err)
	}
	if couponLog.Id == defs.ZERO || couponLog.Del == defs.DELETE || couponLog.Status != 0 || couponLog.UserId != userId {
		panic(errs.NewErrorCoupon("无效的优惠券！"))
	}
	coupon, err := dbops.QueryCouponById(couponLog.CouponId)
	if err != nil {
		panic(err)
	}
	if coupon.Id == defs.ZERO {
		panic(err)
	}
	var discountAmount decimal.Decimal
	switch coupon.Type {
	case 1:
		fullMoney, err := decimal.NewFromString(coupon.FullMoney)
		if err != nil {
			panic(err)
		}
		if goodsAmount.LessThan(fullMoney) {
			panic(errs.NewErrorCoupon("未达到满减要求！"))
		}
		minus, err := decimal.NewFromString(coupon.Minus)
		if err != nil {
			panic(err)
		}
		discountAmount = minus
	case 2:
		rate, err := decimal.NewFromString(coupon.Rate)
		if err != nil {
			panic(err)
		}
		discountAmount = goodsAmount.Mul(rate).Round(2)
	case 3:
		minus, err := decimal.NewFromString(coupon.Minus)
		if err != nil {
			panic(err)
		}
		discountAmount = minus
	case 4:
		fullMoney, err := decimal.NewFromString(coupon.FullMoney)
		if err != nil {
			panic(err)
		}
		if goodsAmount.LessThan(fullMoney) {
			panic(errs.NewErrorCoupon("未达到满减要求！"))
		}
		rate, err := decimal.NewFromString(coupon.Rate)
		if err != nil {
			panic(err)
		}
		discountAmount = goodsAmount.Mul(rate).Round(2)
	default:
		discountAmount = decimal.NewFromInt(0)
	}
	if discountAmount.GreaterThan(goodsAmount) {
		discountAmount = goodsAmount
	}
	return discountAmount
}

func getAddressSnapshot(addressId int) string {
	addressDO, err := dbops.QueryUserAddressById(addressId)
	if err != nil {
		panic(err)
	}
	if addressDO.Id == defs.ZERO || addressDO.Del == defs.DELETE {
		panic(errs.ErrorAddress)
	}
	snapshot := defs.AddressSnapshot{}
	snapshot.Contacts = addressDO.Contacts
	snapshot.Mobile = addressDO.Mobile
	snapshot.ProvinceId = addressDO.ProvinceId
	snapshot.ProvinceStr = addressDO.ProvinceStr
	snapshot.CityId = addressDO.CityId
	snapshot.CityStr = addressDO.CityStr
	snapshot.AreaStr = addressDO.AreaStr
	snapshot.Address = addressDO.Address
	bytes, err := json.Marshal(snapshot)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

// 支付-获取微信预支付ID
func (s *orderService) generateWxpayPrepayId(orderNo string, payAmount string) string {
	// todo: 请求微信支付订单

	return "prepay_id:" + orderNo
}

// 订单详情-快照
func orderGoodsSnapshot(orderNo string, goodsList []defs.PortalCartGoods) {
	for _, v := range goodsList {
		cartDO, _ := dbops.SelectCartById(v.CartId)
		goodsDO, _ := dbops.QueryGoodsById(cartDO.GoodsId)
		skuDO, _ := dbops.GetSKUById(cartDO.SkuId)

		orderGoodsDO := model.WechatMallOrderGoodsDO{}
		orderGoodsDO.OrderNo = orderNo
		orderGoodsDO.GoodsId = cartDO.GoodsId
		orderGoodsDO.SkuId = cartDO.SkuId
		orderGoodsDO.Picture = skuDO.Picture
		orderGoodsDO.Title = goodsDO.Title
		orderGoodsDO.Price = skuDO.Price
		orderGoodsDO.Specs = skuDO.Specs
		orderGoodsDO.Num = v.Num
		orderGoodsDO.LockStatus = 0
		err := dbops.AddOrderGoods(&orderGoodsDO)
		if err != nil {
			panic(err)
		}
		err = dbops.UpdateSkuStockById(cartDO.SkuId, v.Num)
		if err != nil {
			panic(err)
		}
	}
}

// 下单成功-清理购物车
func clearUserCart(goodsList []defs.PortalCartGoods) {
	for _, v := range goodsList {
		if v.CartId != 0 {
			cartDO, _ := dbops.SelectCartById(v.CartId)
			if cartDO.Id == defs.ZERO || cartDO.Del == defs.DELETE {
				continue
			}
			cartDO.Del = defs.DELETE
			err := dbops.UpdateCartById(cartDO)
			if err != nil {
				panic(err)
			}
		}
	}
}

// 优惠券-核销
func couponCannel(couponLogId int) {
	couponLogDO, _ := dbops.QueryCouponLogById(couponLogId)
	couponLogDO.Del = 1
	couponLogDO.UseTime = time.Now().Format("2006-01-02 15:04:05")
	err := dbops.UpdateCouponLogById(couponLogDO)
	if err != nil {
		panic(err)
	}
}

func (s *orderService) QueryOrderList(userId, status, page, size int) (*[]defs.PortalOrderListVO, int) {
	orderList, err := dbops.ListOrderByParams(userId, status, page, size)
	if err != nil {
		panic(err)
	}
	total, err := dbops.CountOrderByParams(userId, status)
	if err != nil {
		panic(err)
	}
	orderVOList := []defs.PortalOrderListVO{}
	for _, v := range *orderList {
		orderVO := defs.PortalOrderListVO{}
		orderVO.Id = v.Id
		orderVO.OrderNo = v.OrderNo
		orderVO.PlaceTime = v.CreateTime
		orderVO.Status = v.Status
		orderVO.GoodsList = extraceOrderGoods(v.OrderNo)
		orderVOList = append(orderVOList, orderVO)
	}
	return &orderVOList, total
}

func extraceOrderGoods(orderNo string) []defs.PortalOrderGoodsVO {
	orderGoodsList, err := dbops.QueryOrderGoods(orderNo)
	if err != nil {
		panic(err)
	}
	goodsVOList := []defs.PortalOrderGoodsVO{}
	for _, v := range *orderGoodsList {
		goodsVO := defs.PortalOrderGoodsVO{}
		goodsVO.GoodsId = v.GoodsId
		goodsVO.Title = v.Title
		goodsVO.Price = v.Price
		goodsVO.Picture = v.Price
		goodsVO.SkuId = v.SkuId
		goodsVO.Specs = v.Specs
		goodsVO.Num = v.Num
		goodsVOList = append(goodsVOList, goodsVO)
	}
	return goodsVOList
}

func (s *orderService) QueryOrderDetail(userId, orderId int) *defs.PortalOrderDetailVO {
	orderDO, err := dbops.QueryOrderById(orderId)
	if err != nil {
		panic(err)
	}
	if orderDO.Id == defs.ZERO || orderDO.Del == defs.DELETE {
		panic(errs.ErrorOrder)
	}
	if orderDO.UserId != userId {
		panic(errs.ErrorOrder)
	}
	snapshot := defs.AddressSnapshot{}
	err = json.Unmarshal([]byte(orderDO.AddressSnapshot), &snapshot)
	if err != nil {
		panic(err)
	}
	orderVO := defs.PortalOrderDetailVO{}
	orderVO.Id = orderDO.Id
	orderVO.OrderNo = orderDO.OrderNo
	orderVO.PlaceTime = orderDO.CreateTime
	orderVO.Status = orderDO.Status
	orderVO.GoodsList = extraceOrderGoods(orderDO.OrderNo)
	orderVO.Address = snapshot
	return &orderVO
}

func (s *orderService) OrderPaySuccessNotify(orderNo string) {
	orderDO, err := dbops.QueryOrderByOrderNo(orderNo)
	if err != nil {
		panic(err)
	}
	if orderDO.Id == defs.ZERO || orderDO.Del == defs.DELETE {
		panic(errs.ErrorOrder)
	}
	if orderDO.Status != 0 {
		log.Printf("orderNo = %v 重复回调", orderDO)
		return
	}
	orderDO.Status = 1
	err = dbops.UpdateOrderById(orderDO)
	if err != nil {
		panic(err)
	}
}

func (s *orderService) QueryOrderSaleData(page, size int) *[]defs.OrderSaleData {
	saleData, err := dbops.QueryOrderSaleData(page, size)
	if err != nil {
		panic(err)
	}
	return saleData
}

// 统计-待发货订单数量
func (s *orderService) CountWaitingOrderNum(status int) int {
	orderNum, err := dbops.CountOrderNum(status)
	if err != nil {
		return 0
	}
	return orderNum
}

// 统计-待处理的退款订单数量
func (s *orderService) CountPendingOrderRefund() int {
	total, err := dbops.CountPendingOrderRefund()
	if err != nil {
		panic(err)
	}
	return total
}
