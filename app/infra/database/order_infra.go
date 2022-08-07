package database

import (
	"context"
	"gorm.io/gorm"
	"time"
	"wechat-mall-backend/app/domain/entity"
	"wechat-mall-backend/app/domain/repository"
	"wechat-mall-backend/consts"
	"wechat-mall-backend/pkg/log"
)

type OrderRepos struct {
	db *gorm.DB
}

func NewOrderRepos(db *gorm.DB) repository.IOrderRepos {
	return &OrderRepos{
		db: db,
	}
}

func (c *OrderRepos) QueryOrderByOrderNo(ctx context.Context, orderNo string) (*entity.WechatMallOrderDO, error) {
	orderDO := new(entity.WechatMallOrderDO)
	if err := c.db.Where("order_no = ?", orderNo).Find(orderDO).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	return orderDO, nil
}

func (c *OrderRepos) QueryOrderById(ctx context.Context, orderId int) (*entity.WechatMallOrderDO, error) {
	orderDO := new(entity.WechatMallOrderDO)
	if err := c.db.Where("id = ?", orderId).Find(orderDO).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	return orderDO, nil
}

func (c *OrderRepos) ListOrderByParams(ctx context.Context, userId, status, page, size int) ([]*entity.WechatMallOrderDO, error) {
	tx := c.db.Where("is_del = 0 AND user_id = ?", userId)
	if status != consts.ALL {
		tx = tx.Where("status = ?", status)
	}
	orderList := make([]*entity.WechatMallOrderDO, 0)
	if err := tx.Order("create_time DESC").Offset((page - 1) * size).Limit(size).Find(&orderList).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	return orderList, nil
}

func (c *OrderRepos) CountOrderByParams(ctx context.Context, userId, status int) (int, error) {
	tx := c.db.Where("is_del = 0 AND user_id = ?", userId)
	if status != consts.ALL {
		tx = tx.Where("status = ?", status)
	}
	empty := new(entity.WechatMallOrderDO)
	var total int64
	if err := tx.Table(empty.TableName()).Count(&total).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Count failed, err: %v", err)
		return 0, err
	}
	return int(total), nil
}

func (c *OrderRepos) AddOrder(ctx context.Context, orderDO *entity.WechatMallOrderDO) error {
	if err := c.db.Create(orderDO).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Create failed, err: %v", err)
		return err
	}
	return nil
}

func (c *OrderRepos) UpdateOrderById(ctx context.Context, order *entity.WechatMallOrderDO) error {
	if err := c.db.Where("id = ?", order.ID).Updates(order).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Updates failed, err: %v", err)
		return err
	}
	return nil
}

func (c *OrderRepos) UpdateOrderRemark(ctx context.Context, id int, remark string) error {
	order := &entity.WechatMallOrderDO{
		ID:         id,
		Remark:     remark,
		UpdateTime: time.Now(),
	}
	return c.UpdateOrderById(ctx, order)
}

func (c *OrderRepos) QueryOrderSaleData(ctx context.Context, page, size int) ([]*entity.OrderSaleData, error) {
	empty := new(entity.WechatMallOrderDO)
	saleList := make([]*entity.OrderSaleData, 0)
	if err := c.db.Table(empty.TableName()).Select("DATE_FORMAT(create_time, '%Y-%m-%d') AS order_time,  COUNT(id) AS order_mum, IFNULL( SUM(pay_amount), 0) AS sale_amount").
		Where("status in (1, 2, 3)").Group("order_time").Order("order_time DESC").Offset((page - 1) * size).Limit(size).Find(&saleList).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	return saleList, nil
}

func (c *OrderRepos) CountOrderNum(ctx context.Context, userId, status int) (int, error) {
	empty := new(entity.WechatMallOrderDO)
	tx := c.db.Table(empty.TableName()).Where("is_del = 0 AND status = ?", status)
	if userId != consts.ALL {
		tx = tx.Where("user_id = ?", userId)
	}
	var total int64
	if err := tx.Count(&total).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Count failed, err: %v", err)
		return 0, err
	}
	return int(total), nil
}

func (c *OrderRepos) SelectCMSOrderList(ctx context.Context, status, searchType int, keyword, startTime, endTime string, page, size int) ([]*entity.WechatMallOrderDO, error) {
	tx := c.db.Where("1 = 1")
	if status != consts.ALL {
		tx = tx.Where("status = ?", status)
	}
	if searchType != consts.ALL && keyword != "" {
		switch searchType {
		case 1:
			tx = tx.Where("order_no like ?", "%"+keyword+"%")
		case 2, 3:
			tx = tx.Where("address_snapshot like ?", "%"+keyword+"%")
		case 4:
			tx = tx.Where("transaction_id LIKE ?", "%"+keyword+"%")
		case 5:
			tx = tx.Where("order_no IN (SELECT order_no FROM wechat_mall_order_goods WHERE title LIKE ?)", "%"+keyword+"%")
		}
	}
	if startTime != "" {
		tx = tx.Where("create_time > ?", startTime)
	}
	if endTime != "" {
		tx = tx.Where("create_time < ?", endTime)
	}
	orderList := make([]*entity.WechatMallOrderDO, 0)
	if err := tx.Offset((page - 1) * size).Limit(size).Find(&orderList).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	return orderList, nil
}

func (c *OrderRepos) SelectCMSOrderNum(ctx context.Context, status, searchType int, keyword, startTime, endTime string) (int, error) {
	empty := new(entity.WechatMallOrderDO)
	tx := c.db.Table(empty.TableName()).Where("1 = 1")
	if status != consts.ALL {
		tx = tx.Where("status = ?", status)
	}
	if searchType != consts.ALL && keyword != "" {
		switch searchType {
		case 1:
			tx = tx.Where("order_no like ?", "%"+keyword+"%")
		case 2, 3:
			tx = tx.Where("address_snapshot like ?", "%"+keyword+"%")
		case 4:
			tx = tx.Where("transaction_id LIKE ?", "%"+keyword+"%")
		case 5:
			tx = tx.Where("order_no IN (SELECT order_no FROM wechat_mall_order_goods WHERE title LIKE ?)", "%"+keyword+"%")
		}
	}
	if startTime != "" {
		tx = tx.Where("create_time > ?", startTime)
	}
	if endTime != "" {
		tx = tx.Where("create_time < ?", endTime)
	}
	var total int64
	if err := tx.Count(&total).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Count failed, err: %v", err)
		return 0, err
	}
	return int(total), nil
}

func (c *OrderRepos) QueryOrderGoods(ctx context.Context, orderNo string) ([]*entity.WechatMallOrderGoodsDO, error) {
	goodsList := make([]*entity.WechatMallOrderGoodsDO, 0)
	if err := c.db.Where("order_no = ?", orderNo).Find(&goodsList).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	return goodsList, nil
}

func (c *OrderRepos) AddOrderGoods(ctx context.Context, goods *entity.WechatMallOrderGoodsDO) error {
	if err := c.db.Create(goods).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Create failed, err: %v", err)
		return err
	}
	return nil
}

func (c *OrderRepos) UpdateOrderGoods(ctx context.Context, goods *entity.WechatMallOrderGoodsDO) error {
	if err := c.db.Where("id = ?", goods.ID).Updates(goods).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Updates failed, err: %v", err)
		return err
	}
	return nil
}

func (c *OrderRepos) CountBuyGoodsUserNum(ctx context.Context, goodsId int) (int, error) {
	empty := new(entity.WechatMallOrderDO)
	var total int64
	if err := c.db.Table(empty.TableName()).Select("COUNT(DISTINCT(user_id))").Where("lock_status = 1 AND goods_id = ?", goodsId).Count(&total).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Count failed, err: %v", err)
		return 0, err
	}
	return int(total), nil
}

func (c *OrderRepos) AddRefundRecord(ctx context.Context, orderRefund *entity.WechatMallOrderRefund) error {
	if err := c.db.Create(orderRefund).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Create failed, err: %v", err)
		return err
	}
	return nil
}

func (c *OrderRepos) QueryRefundRecord(ctx context.Context, refundNo string) (*entity.WechatMallOrderRefund, error) {
	refundDO := new(entity.WechatMallOrderRefund)
	if err := c.db.Where("refund_no = ?", refundNo).Find(refundDO).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	return refundDO, nil
}
func (c *OrderRepos) QueryOrderRefundRecord(ctx context.Context, orderNo string) (*entity.WechatMallOrderRefund, error) {
	refundDO := new(entity.WechatMallOrderRefund)
	if err := c.db.Where("order_no = ?", orderNo).Find(refundDO).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Find failed, err: %v", err)
		return nil, err
	}
	return refundDO, nil
}

func (c *OrderRepos) UpdateRefundApply(ctx context.Context, id int, status int) error {
	refundDO := &entity.WechatMallOrderRefund{
		Status:     status,
		UpdateTime: time.Now(),
		RefundTime: time.Now(),
	}
	if err := c.db.Where("status = 0 AND id = ?", id).Updates(refundDO).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Update failed, err: %v", err)
		return err
	}
	return nil
}

func (c *OrderRepos) CountPendingOrderRefund(ctx context.Context) (int, error) {
	empty := new(entity.WechatMallOrderRefund)
	var total int64
	if err := c.db.Table(empty.TableName()).Where("status IN (0, 1)").Count(&total).Error; err != nil {
		log.ErrorContextf(ctx, "call db.Count failed, err: %v", err)
		return 0, err
	}
	return int(total), nil
}
