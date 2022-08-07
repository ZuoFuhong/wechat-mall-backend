package repository

import (
	"context"
	"wechat-mall-backend/app/domain/entity"
)

type IOrderRepos interface {
	QueryOrderByOrderNo(ctx context.Context, orderNo string) (*entity.WechatMallOrderDO, error)

	QueryOrderById(ctx context.Context, orderId int) (*entity.WechatMallOrderDO, error)

	ListOrderByParams(ctx context.Context, userId, status, page, size int) ([]*entity.WechatMallOrderDO, error)

	CountOrderByParams(ctx context.Context, userId, status int) (int, error)

	AddOrder(ctx context.Context, orderDO *entity.WechatMallOrderDO) error

	UpdateOrderById(ctx context.Context, order *entity.WechatMallOrderDO) error

	UpdateOrderRemark(ctx context.Context, id int, remark string) error

	QueryOrderSaleData(ctx context.Context, page, size int) ([]*entity.OrderSaleData, error)

	CountOrderNum(ctx context.Context, userId, status int) (int, error)

	SelectCMSOrderList(ctx context.Context, status, searchType int, keyword, startTime, endTime string, page, size int) ([]*entity.WechatMallOrderDO, error)

	SelectCMSOrderNum(ctx context.Context, status, searchType int, keyword, startTime, endTime string) (int, error)

	QueryOrderGoods(ctx context.Context, orderNo string) ([]*entity.WechatMallOrderGoodsDO, error)

	AddOrderGoods(ctx context.Context, goods *entity.WechatMallOrderGoodsDO) error

	UpdateOrderGoods(ctx context.Context, goods *entity.WechatMallOrderGoodsDO) error

	CountBuyGoodsUserNum(ctx context.Context, goodsId int) (int, error)

	AddRefundRecord(ctx context.Context, orderRefund *entity.WechatMallOrderRefund) error

	QueryRefundRecord(ctx context.Context, refundNo string) (*entity.WechatMallOrderRefund, error)

	QueryOrderRefundRecord(ctx context.Context, orderNo string) (*entity.WechatMallOrderRefund, error)

	UpdateRefundApply(ctx context.Context, id int, status int) error

	CountPendingOrderRefund(ctx context.Context) (int, error)
}
