package service

import (
	"encoding/json"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/shopspring/decimal"
	"log"
	"os"
	"strconv"
	"time"
	"wechat-mall-backend/dbops"
	"wechat-mall-backend/defs"
	"wechat-mall-backend/env"
	"wechat-mall-backend/errs"
	"wechat-mall-backend/model"
	"wechat-mall-backend/utils"
)

type IOrderService interface {
	GenerateOrder(userId, addressId, couponLogId int, dispatchAmount, expectAmount decimal.Decimal, goodsList []defs.PortalCartGoods) defs.PortalPlaceOrderVO
	QueryOrderList(userId, status, page, size int) (*[]defs.PortalOrderListVO, int)
	QueryOrderDetail(userId int, orderNo string) *defs.PortalOrderDetailVO
	OrderPaySuccessNotify(orderNo string)
	QueryOrderSaleData(page, size int) *[]defs.OrderSaleData
	CountOrderNum(userId, status int) int
	CountPendingOrderRefund() int
	CancelOrder(userId, orderId int)
	DeleteOrderRecord(userId, orderId int)
	ConfirmTakeGoods(userId, orderId int)
	RefundApply(userId int, orderNo, reason string) string
	QueryRefundDetail(userId int, refundNo string) *defs.OrderRefundDetailVO
	UndoRefundApply(userId int, refundNo string)
	QueryCMSOrderList(status, searchType int, keyword, startTime, endTime string, page, size int) (*[]defs.CMSOrderInfoVO, int)
	ExportCMSOrderExcel(status, searchType int, keyword, startTime, endTime string) string
	QueryCMSOrderDetail(orderNo string) *defs.CMSOrderInfoVO
	ModifyOrderStatus(orderNo string, otype int)
	ModifyOrderRemark(orderNo, remark string)
	ModifyOrderGoods(orderNo string, goodsId int, price string)
}

type orderService struct {
}

func NewOrderService() IOrderService {
	service := orderService{}
	return &service
}

func (s *orderService) GenerateOrder(userId, addressId, couponLogId int, dispatchAmount, expectAmount decimal.Decimal,
	goodsList []defs.PortalCartGoods) defs.PortalPlaceOrderVO {
	goodsAmount := checkCartGoodsAndStock(goodsList)
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
	orderDO.PayTime = "2006-01-02 15:04:05"
	orderDO.DeliverTime = "2006-01-02 15:04:05"
	orderDO.FinishTime = "2006-01-02 15:04:05"
	orderDO.Status = 0
	orderDO.AddressId = addressId
	orderDO.AddressSnapshot = addressSnap
	orderDO.WxappPrepayId = prepayId
	err := dbops.AddOrder(&orderDO)
	if err != nil {
		panic(err)
	}
	orderGoodsSnapshot(userId, orderNo, goodsList)
	clearUserCart(goodsList)
	couponCannel(couponLogId)
	return defs.PortalPlaceOrderVO{OrderNo: orderNo, PrepayId: prepayId}
}

// 检查-购物车以及商品的库存
func checkCartGoodsAndStock(goodsList []defs.PortalCartGoods) decimal.Decimal {
	goodsAmount := decimal.NewFromInt(0)
	for _, v := range goodsList {
		if v.CartId != 0 {
			cartDO, err := dbops.SelectCartById(v.CartId)
			if err != nil {
				panic(err)
			}
			if cartDO.Id == defs.ZERO || cartDO.Del == defs.DELETE {
				panic(errs.ErrorGoodsCart)
			}
		}
		goodsDO, err := dbops.QueryGoodsById(v.GoodsId)
		if err != nil {
			panic(err)
		}
		if goodsDO.Id == defs.ZERO || goodsDO.Del == defs.DELETE || goodsDO.Online == defs.OFFLINE {
			panic(errs.NewErrorOrder("商品下架，无法售出"))
		}
		skuDO, err := dbops.GetSKUById(v.SkuId)
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
		discountAmount = goodsAmount.Sub(goodsAmount.Mul(rate).Round(2))
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
		discountAmount = goodsAmount.Sub(goodsAmount.Mul(rate).Round(2))
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
func orderGoodsSnapshot(userId int, orderNo string, goodsList []defs.PortalCartGoods) {
	for _, v := range goodsList {
		goodsDO, _ := dbops.QueryGoodsById(v.GoodsId)
		skuDO, _ := dbops.GetSKUById(v.SkuId)

		orderGoodsDO := model.WechatMallOrderGoodsDO{}
		orderGoodsDO.OrderNo = orderNo
		orderGoodsDO.UserId = userId
		orderGoodsDO.GoodsId = v.GoodsId
		orderGoodsDO.SkuId = v.SkuId
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
		// 减库存
		err = dbops.UpdateSkuStockById(v.SkuId, v.Num)
		if err != nil {
			panic(err)
		}
		// 商品销量
		err = dbops.UpdateGoodsSaleNum(v.GoodsId, v.Num)
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
	if couponLogDO.Id == 0 {
		return
	}
	couponLogDO.Status = 1
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
		orderVO.PayAmount, _ = strconv.ParseFloat(v.PayAmount, 2)
		orderVO.Status = v.Status
		orderVO.GoodsList, orderVO.GoodsNum = extraceOrderGoods(v.OrderNo)
		orderVOList = append(orderVOList, orderVO)
	}
	return &orderVOList, total
}

func extraceOrderGoods(orderNo string) ([]defs.PortalOrderGoodsVO, int) {
	orderGoodsList, err := dbops.QueryOrderGoods(orderNo)
	if err != nil {
		panic(err)
	}
	goodsNum := 0
	goodsVOList := []defs.PortalOrderGoodsVO{}
	for _, v := range *orderGoodsList {
		specList := []defs.SkuSpecs{}
		err := json.Unmarshal([]byte(v.Specs), &specList)
		if err != nil {
			panic(err)
		}
		specs := ""
		for _, v := range specList {
			specs += v.Value + "; "
		}
		if len(specs) > 2 {
			specs = specs[0 : len(specs)-2]
		}
		goodsVO := defs.PortalOrderGoodsVO{}
		goodsVO.GoodsId = v.GoodsId
		goodsVO.Title = v.Title
		goodsVO.Price, _ = strconv.ParseFloat(v.Price, 2)
		goodsVO.Picture = v.Picture
		goodsVO.SkuId = v.SkuId
		goodsVO.Specs = specs
		goodsVO.Num = v.Num
		goodsVOList = append(goodsVOList, goodsVO)
		goodsNum += v.Num
	}
	return goodsVOList, goodsNum
}

func (s *orderService) QueryOrderDetail(userId int, orderNo string) *defs.PortalOrderDetailVO {
	orderDO, err := dbops.QueryOrderByOrderNo(orderNo)
	if err != nil {
		panic(err)
	}
	if orderDO.Id == defs.ZERO || orderDO.Del == defs.DELETE {
		panic(errs.ErrorOrder)
	}
	if orderDO.UserId != userId {
		panic(errs.ErrorOrder)
	}
	// 订单信息
	snapshot := defs.AddressSnapshot{}
	err = json.Unmarshal([]byte(orderDO.AddressSnapshot), &snapshot)
	if err != nil {
		panic(err)
	}
	orderGoods, orderGoodsNum := extraceOrderGoods(orderDO.OrderNo)
	// 退款信息
	refundDO, err := dbops.QueryOrderRefundRecord(orderNo)
	if err != nil {
		panic(err)
	}
	refundApply := defs.OrderRefundApplyVO{RefundNo: refundDO.RefundNo}

	orderVO := defs.PortalOrderDetailVO{}
	orderVO.Id = orderDO.Id
	orderVO.OrderNo = orderDO.OrderNo
	orderVO.GoodsAmount, _ = strconv.ParseFloat(orderDO.GoodsAmount, 2)
	orderVO.DiscountAmount, _ = strconv.ParseFloat(orderDO.DiscountAmount, 2)
	orderVO.DispatchAmount, _ = strconv.ParseFloat(orderDO.DispatchAmount, 2)
	orderVO.PayAmount, _ = strconv.ParseFloat(orderDO.PayAmount, 2)
	orderVO.Status = orderDO.Status
	orderVO.PlaceTime = orderDO.CreateTime
	orderVO.PayTime = orderDO.PayTime
	orderVO.DeliverTime = orderDO.DeliverTime
	orderVO.FinishTime = orderDO.FinishTime
	orderVO.GoodsList = orderGoods
	orderVO.GoodsNum = orderGoodsNum
	orderVO.Address = snapshot
	orderVO.RefundApply = refundApply
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
	orderDO.PayTime = time.Now().Format("2006-01-02 15:04:05")
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

// 统计-订单数量
func (s *orderService) CountOrderNum(userId, status int) int {
	orderNum, err := dbops.CountOrderNum(userId, status)
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

// 订单-取消订单
func (s *orderService) CancelOrder(userId, orderId int) {
	orderDO, err := dbops.QueryOrderById(orderId)
	if err != nil {
		panic(err)
	}
	if orderDO.Id == defs.ZERO || orderDO.Del == defs.DELETE {
		panic(errs.ErrorOrder)
	}
	if orderDO.UserId != userId {
		panic(errs.NewErrorOrder("非法操作"))
	}
	if orderDO.Status != 0 {
		panic(errs.NewErrorOrder("进行中的订单，无法取消"))
	}
	orderDO.Status = -1
	orderDO.FinishTime = time.Now().Format("2006-01-02 15:04:05")
	err = dbops.UpdateOrderById(orderDO)
	if err != nil {
		panic(err)
	}
	orderStockRollback(orderDO.OrderNo)
}

// 订单-库存回滚
// 场景：取消订单（手动取消、超时未支付）、订单退款
func orderStockRollback(orderNo string) {
	orderGoods, err := dbops.QueryOrderGoods(orderNo)
	if err != nil {
		panic(err)
	}
	for _, v := range *orderGoods {
		err := dbops.UpdateSkuStockById(v.SkuId, -v.Num)
		if err != nil {
			panic(err)
		}
		err = dbops.UpdateGoodsSaleNum(v.GoodsId, -v.Num)
		if err != nil {
			panic(err)
		}
	}
}

// 订单-删除记录
func (s *orderService) DeleteOrderRecord(userId, orderId int) {
	orderDO, err := dbops.QueryOrderById(orderId)
	if err != nil {
		panic(err)
	}
	if orderDO.Id == defs.ZERO || orderDO.Del == defs.DELETE {
		panic(errs.ErrorOrder)
	}
	if orderDO.UserId != userId {
		panic(errs.NewErrorOrder("非法操作"))
	}
	if orderDO.Status == -1 || orderDO.Status == 3 || orderDO.Status == 5 {
		orderDO.Del = 1
		err := dbops.UpdateOrderById(orderDO)
		if err != nil {
			panic(err)
		}
	} else {
		panic(errs.NewErrorOrder("进行中的订单，无法删除"))
	}
}

// 订单-确认收货
func (s *orderService) ConfirmTakeGoods(userId, orderId int) {
	orderDO, err := dbops.QueryOrderById(orderId)
	if err != nil {
		panic(err)
	}
	if orderDO.Id == defs.ZERO || orderDO.Del == defs.DELETE {
		panic(errs.ErrorOrder)
	}
	if orderDO.UserId != userId {
		panic(errs.NewErrorOrder("非法操作"))
	}
	if orderDO.Status == 2 {
		orderDO.Status = 3
		orderDO.FinishTime = time.Now().Format("2006-01-02 15:04:05")
		err := dbops.UpdateOrderById(orderDO)
		if err != nil {
			panic(err)
		}
	}
}

func (s *orderService) RefundApply(userId int, orderNo, reason string) string {
	orderDO, err := dbops.QueryOrderByOrderNo(orderNo)
	if err != nil {
		panic(err)
	}
	if orderDO.Id == defs.ZERO || orderDO.Del == defs.DELETE {
		panic(errs.ErrorOrder)
	}
	if orderDO.UserId != userId {
		panic(errs.NewErrorOrder("非法操作"))
	}
	if orderDO.Status != 1 {
		panic(errs.NewErrorOrder("非法操作"))
	}
	refundNo := time.Now().Format("20060102150405") + utils.RandomNumberStr(6)
	refund := model.WechatMallOrderRefund{}
	refund.RefundNo = refundNo
	refund.UserId = userId
	refund.OrderNo = orderNo
	refund.Reason = reason
	refund.RefundTime = "2006-01-02 15:04:05"
	refund.RefundAmount = orderDO.PayAmount
	err = dbops.AddRefundRecord(&refund)
	if err != nil {
		panic(err)
	}
	orderDO.Status = 4
	err = dbops.UpdateOrderById(orderDO)
	if err != nil {
		panic(err)
	}
	return refundNo
}

func (s *orderService) QueryRefundDetail(userId int, refundNo string) *defs.OrderRefundDetailVO {
	refundDO, err := dbops.QueryRefundRecord(refundNo)
	if err != nil {
		panic(err)
	}
	if refundDO.Id == defs.ZERO || refundDO.Del == defs.DELETE {
		panic(errs.ErrorOrderRefund)
	}
	if refundDO.UserId != userId {
		panic(errs.NewErrorOrder("非法操作"))
	}
	refundVO := defs.OrderRefundDetailVO{}
	refundVO.RefundNo = refundDO.RefundNo
	refundVO.Reason = refundDO.Reason
	refundVO.RefundAmount, _ = strconv.ParseFloat(refundDO.RefundAmount, 2)
	refundVO.Status = refundDO.Status
	refundVO.ApplyTime = refundDO.CreateTime
	refundVO.RefundTime = refundDO.RefundTime
	refundVO.GoodsList, _ = extraceOrderGoods(refundDO.OrderNo)
	return &refundVO
}

func (s *orderService) UndoRefundApply(userId int, refundNo string) {
	refundDO, err := dbops.QueryRefundRecord(refundNo)
	if err != nil {
		panic(err)
	}
	if refundDO.Id == defs.ZERO || refundDO.Del == defs.DELETE {
		panic(errs.ErrorOrderRefund)
	}
	if refundDO.UserId != userId {
		panic(errs.NewErrorOrderRefund("非法操作"))
	}
	if refundDO.Status != 0 {
		panic(errs.NewErrorOrderRefund("状态异常"))
	}
	// 订单：待收货
	orderDO, err := dbops.QueryOrderByOrderNo(refundDO.OrderNo)
	if err != nil {
		panic(err)
	}
	orderDO.Status = 1
	err = dbops.UpdateOrderById(orderDO)
	if err != nil {
		panic(err)
	}
	err = dbops.UpdateRefundApply(refundDO.Id, 2)
	if err != nil {
		panic(err)
	}
}

func (s *orderService) QueryCMSOrderList(status, searchType int, keyword, startTime, endTime string, page, size int) (*[]defs.CMSOrderInfoVO, int) {
	orderList, err := dbops.SelectCMSOrderList(status, searchType, keyword, startTime, endTime, page, size)
	if err != nil {
		panic(err)
	}
	total, err := dbops.SelectCMSOrderNum(status, searchType, keyword, startTime, endTime)
	if err != nil {
		panic(err)
	}
	orderVOList := []defs.CMSOrderInfoVO{}
	for _, v := range *orderList {
		address := extractOrderAddress(v.AddressSnapshot)
		buyer := extractOrderBuyer(v.UserId)

		orderVO := defs.CMSOrderInfoVO{}
		orderVO.OrderNo = v.OrderNo
		orderVO.PlaceTime = v.CreateTime
		orderVO.Address = *address
		orderVO.PayAmount, _ = strconv.ParseFloat(v.PayAmount, 2)
		orderVO.GoodsAmount, _ = strconv.ParseFloat(v.GoodsAmount, 2)
		orderVO.DiscountAmount, _ = strconv.ParseFloat(v.DiscountAmount, 2)
		orderVO.DispatchAmount, _ = strconv.ParseFloat(v.DispatchAmount, 2)
		orderVO.Status = v.Status
		orderVO.TransactionId = v.TransactionId
		orderVO.Remark = v.Remark
		orderVO.PayTime = v.PayTime
		orderVO.DeliverTime = v.DeliverTime
		orderVO.FinishTime = v.FinishTime
		orderVO.Buyer = *buyer
		orderVOList = append(orderVOList, orderVO)
	}
	return &orderVOList, total
}

func (s *orderService) ExportCMSOrderExcel(status, searchType int, keyword, startTime, endTime string) string {
	orderList, err := dbops.SelectCMSOrderList(status, searchType, keyword, startTime, endTime, 0, 0)
	if err != nil {
		panic(err)
	}
	excelData := extractOrderExcelData(*orderList)
	filepath, filename := generateExcel(excelData)
	ossLink := uploadFileToOSS(filepath, filename)
	_ = os.Remove(filepath)

	return ossLink
}

func uploadFileToOSS(filepath, filename string) string {
	conf := env.LoadConf()
	ossConf := conf.Oss
	endpoint := "https://oss-cn-hangzhou.aliyuncs.com"

	client, e := oss.New(endpoint, ossConf.AccessKeyId, ossConf.AccessKeySecret)
	if e != nil {
		panic(e)
	}
	bucket, e := client.Bucket(ossConf.BucketName)
	if e != nil {
		panic(e)
	}
	objectKey := "order-excel/" + filename
	e = bucket.PutObjectFromFile(objectKey, filepath)
	if e != nil {
		panic(e)
	}
	ossLink := "https://" + ossConf.BucketName + ".oss-cn-hangzhou.aliyuncs.com/" + objectKey
	return ossLink
}

// 订单-提取Excel数据
func extractOrderExcelData(orderList []model.WechatMallOrderDO) *map[string]string {
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
		goodsList := extractOrderGoodsVO(v.OrderNo)
		goodsAmount, _ := decimal.NewFromString(v.GoodsAmount)
		dispatchAmount, _ := decimal.NewFromString(v.DispatchAmount)
		orderAmount := goodsAmount.Add(dispatchAmount).String()
		statusStr := extractOrderStatus(v.Status)
		addressSnap := extractOrderAddress(v.AddressSnapshot)

		for _, g := range goodsList {
			rowNum += 1
			rowNumStr := strconv.Itoa(rowNum)
			excelData["A"+rowNumStr] = v.OrderNo
			excelData["B"+rowNumStr] = v.TransactionId
			excelData["C"+rowNumStr] = v.CreateTime
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
	return &excelData
}

// 生成Excel文件，保存在本地
func generateExcel(excelData *map[string]string) (string, string) {
	xlsx := excelize.NewFile()
	sheet := xlsx.NewSheet("sheet1")
	for k, v := range *excelData {
		xlsx.SetCellValue("sheet1", k, v)
	}
	xlsx.SetActiveSheet(sheet)
	filename := "订单列表_" + utils.FormatDatetime(time.Now(), utils.YYYYMMDDHHMMSS) + ".xlsx"
	filepath := "/tmp/wechat-mall/" + filename
	utils.CheckFileDirExists(filepath)
	err := xlsx.SaveAs(filepath)
	if err != nil {
		panic(err)
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
	specList := []defs.SkuSpecs{}
	err := json.Unmarshal([]byte(specs), &specList)
	if err != nil {
		panic(err)
	}
	specStr := ""
	for _, v := range specList {
		specStr += v.Value + ";"
	}
	return specStr
}

func (s *orderService) QueryCMSOrderDetail(orderNo string) *defs.CMSOrderInfoVO {
	orderDO, err := dbops.QueryOrderByOrderNo(orderNo)
	if err != nil {
		panic(err)
	}
	goodsList := extractOrderGoodsVO(orderNo)
	address := extractOrderAddress(orderDO.AddressSnapshot)
	buyer := extractOrderBuyer(orderDO.UserId)

	orderVO := defs.CMSOrderInfoVO{}
	orderVO.OrderNo = orderDO.OrderNo
	orderVO.PlaceTime = orderDO.CreateTime
	orderVO.Address = *address
	orderVO.PayAmount, _ = strconv.ParseFloat(orderDO.PayAmount, 2)
	orderVO.GoodsAmount, _ = strconv.ParseFloat(orderDO.GoodsAmount, 2)
	orderVO.DiscountAmount, _ = strconv.ParseFloat(orderDO.DiscountAmount, 2)
	orderVO.DispatchAmount, _ = strconv.ParseFloat(orderDO.DispatchAmount, 2)
	orderVO.Status = orderDO.Status
	orderVO.TransactionId = orderDO.TransactionId
	orderVO.Remark = orderDO.Remark
	orderVO.PayTime = orderDO.PayTime
	orderVO.DeliverTime = orderDO.DeliverTime
	orderVO.FinishTime = orderDO.FinishTime
	orderVO.Buyer = *buyer
	orderVO.GoodsList = goodsList
	return &orderVO
}

// 提取订单中的商品
func extractOrderGoodsVO(orderNo string) []defs.CMSOrderGoodsVO {
	goodsDOList, err := dbops.QueryOrderGoods(orderNo)
	if err != nil {
		panic(err)
	}
	goodsVOList := []defs.CMSOrderGoodsVO{}
	for _, v := range *goodsDOList {
		specs := extractOrderGoodsSpecs(v.Specs)

		goodsVO := defs.CMSOrderGoodsVO{}
		goodsVO.Picture = v.Picture
		goodsVO.Title = v.Title
		goodsVO.Price, _ = strconv.ParseFloat(v.Price, 2)
		goodsVO.Specs = specs
		goodsVO.Num = v.Num
		goodsVOList = append(goodsVOList, goodsVO)
	}
	return goodsVOList
}

// 订单-提取收货人
func extractOrderAddress(addressSnapshot string) *defs.AddressSnapshot {
	address := defs.AddressSnapshot{}
	err := json.Unmarshal([]byte(addressSnapshot), &address)
	if err != nil {
		panic(err)
	}
	return &address
}

// 订单-提取买家信息
func extractOrderBuyer(uid int) *defs.BasicUser {
	userDO, e := dbops.GetUserById(uid)
	if e != nil {
		panic(e)
	}
	basicUser := defs.BasicUser{}
	basicUser.UserId = userDO.Id
	basicUser.Nickname = userDO.Nickname
	basicUser.Avatar = userDO.Avatar
	return &basicUser
}

func (s *orderService) ModifyOrderStatus(orderNo string, otype int) {
	/*
		待发货：1-确认发货
		待收货：2-确认收货
		待付款：3-确认付款
	*/
	orderDO, err := dbops.QueryOrderByOrderNo(orderNo)
	if err != nil {
		panic(errs.NewErrorOrder("订单不存在"))
	}
	newStatus := 0
	switch otype {
	case 1:
		if orderDO.Status != 1 {
			panic(errs.NewErrorOrder("非法操作"))
		}
		newStatus = 2
	case 2:
		if orderDO.Status != 2 {
			panic(errs.NewErrorOrder("非法操作"))
		}
		newStatus = 3
	case 3:
		if orderDO.Status != 0 {
			panic(errs.NewErrorOrder("非法操作"))
		}
		newStatus = 1
	default:
		panic(errs.NewErrorOrder("非法操作"))
	}
	params := model.WechatMallOrderDO{}
	params.Id = orderDO.Id
	params.Status = newStatus
	err = dbops.UpdateOrderById(&params)
	if err != nil {
		panic(err)
	}
}

func (s *orderService) ModifyOrderRemark(orderNo, remark string) {
	orderDO, err := dbops.QueryOrderByOrderNo(orderNo)
	if err != nil {
		panic(errs.NewErrorOrder("订单不存在"))
	}
	err = dbops.UpdateOrderRemark(orderDO.Id, remark)
	if err != nil {
		panic(err)
	}
}

func (s *orderService) ModifyOrderGoods(orderNo string, goodsId int, price string) {
	orderDO, err := dbops.QueryOrderByOrderNo(orderNo)
	if err != nil {
		panic(errs.NewErrorOrder("订单不存在"))
	}
	if orderDO.Status != 0 {
		panic(errs.NewErrorOrder("非法操作"))
	}
	goodsList, err := dbops.QueryOrderGoods(orderNo)
	if err != nil {
		panic(err)
	}
	orderGoods := model.WechatMallOrderGoodsDO{}
	for _, v := range *goodsList {
		if v.GoodsId == goodsId {
			orderGoods = v
		}
	}
	if orderGoods.Id == 0 {
		panic(errs.NewErrorOrder("订单商品异常！"))
	}
	// 更新商品价格
	params := model.WechatMallOrderGoodsDO{}
	params.Id = orderGoods.Id
	params.Price = price
	err = dbops.UpdateOrderGoods(&params)
	if err != nil {
		panic(err)
	}
	// 订单金额，计算差价
	newGoodsPrice, _ := decimal.NewFromString(price)
	oldGoodsPrice, _ := decimal.NewFromString(orderGoods.Price)
	diffAmount := newGoodsPrice.Sub(oldGoodsPrice).Mul(decimal.NewFromInt(int64(orderGoods.Num)))

	payAmount, _ := decimal.NewFromString(orderDO.PayAmount)
	goodsAmount, _ := decimal.NewFromString(orderDO.GoodsAmount)

	orderParams := model.WechatMallOrderDO{}
	orderParams.Id = orderDO.Id
	orderParams.PayAmount = payAmount.Add(diffAmount).String()
	orderParams.GoodsAmount = goodsAmount.Add(diffAmount).String()
	err = dbops.UpdateOrderById(&orderParams)
	if err != nil {
		panic(err)
	}
}
