package view

type OrderRemindVO struct {
	WaitPay     int `json:"waitPay"`     // 待付款
	NotExpress  int `json:"notExpress"`  // 待发货
	WaitReceive int `json:"waitReceive"` // 待收货
}

type PortalPlaceOrderVO struct {
	OrderNo  string `json:"orderNo"`  // 订单号
	PrepayId string `json:"prepayId"` // 预支付ID
}

type PortalOrderListVO struct {
	Id        int                   `json:"id"`        // 订单ID
	OrderNo   string                `json:"orderNo"`   // 订单号
	PlaceTime string                `json:"placeTime"` // 下单时间
	PayAmount float64               `json:"payAmount"` // 支付金额
	Status    int                   `json:"status"`    // 订单状态 -1 已取消 0-待付款 1-待发货 2-待收货 3-已完成
	GoodsNum  int                   `json:"goodsNum"`  // 商品数量
	GoodsList []*PortalOrderGoodsVO `json:"goodsList"`
}

type PortalOrderGoodsVO struct {
	GoodsId int     `json:"goodsId"` // 商品ID
	Title   string  `json:"title"`   // 标题
	Price   float64 `json:"price"`   // 价格
	Picture string  `json:"picture"` // 图片
	SkuId   int     `json:"skuId"`   // skuId
	Specs   string  `json:"specs"`   // specs值
	Num     int     `json:"num"`     // 数量
}

type PortalOrderDetailVO struct {
	Id             int                   `json:"id"`             // 订单ID
	OrderNo        string                `json:"orderNo"`        // 订单号
	GoodsAmount    float64               `json:"goodsAmount"`    // 商品小计
	DiscountAmount float64               `json:"discountAmount"` // 优惠金额
	DispatchAmount float64               `json:"dispatchAmount"` // 运费
	PayAmount      float64               `json:"payAmount"`      // 实付款
	Status         int                   `json:"status"`         // 订单状态 -1 已取消 0-待付款 1-待发货 2-待收货 3-已完成 4-（待发货）退款申请 5-已退款
	GoodsNum       int                   `json:"goodsNum"`       // 商品数量
	PlaceTime      string                `json:"placeTime"`      // 下单时间
	PayTime        string                `json:"payTime"`        // 付款时间
	DeliverTime    string                `json:"deliverTime"`    // 发货时间
	FinishTime     string                `json:"finishTime"`     // 完成时间
	GoodsList      []*PortalOrderGoodsVO `json:"goodsList"`      // 订单商品
	Address        *AddressSnapshot      `json:"address"`        // 收货地址
	RefundApply    *OrderRefundApplyVO   `json:"refundApply"`    // 退款申请
}

type OrderRefundApplyVO struct {
	RefundNo string `json:"refundNo"` // 退款编号
}

type OrderRefundDetailVO struct {
	RefundNo     string                `json:"refundNo"`     // 退款编号
	Reason       string                `json:"reason"`       // 退款原因
	RefundAmount float64               `json:"refundAmount"` // 退款金额
	Status       int                   `json:"status"`       // 状态：0-退款申请 1-商家处理申请 2-退款完成
	ApplyTime    string                `json:"applyTime"`    // 申请时间
	RefundTime   string                `json:"refundTime"`   // 退款时间
	GoodsList    []*PortalOrderGoodsVO `json:"goodsList"`    // 订单商品
}

type AddressSnapshot struct {
	Contacts    string `json:"contacts"`    // 联系人
	Mobile      string `json:"mobile"`      // 手机号
	ProvinceId  string `json:"provinceId"`  // 省份编码
	CityId      string `json:"cityId"`      // 城市编码
	AreaId      string `json:"areaId"`      // 地区编码
	ProvinceStr string `json:"provinceStr"` // 省份
	CityStr     string `json:"cityStr"`     // 城市
	AreaStr     string `json:"areaStr"`     // 地区
	Address     string `json:"address"`     // 详细地址
}

type CMSOrderInfoVO struct {
	OrderNo        string             `json:"orderNo"`
	PlaceTime      string             `json:"placeTime"`
	PayAmount      float64            `json:"payAmount"`
	GoodsAmount    float64            `json:"goodsAmount"`
	DiscountAmount float64            `json:"discountAmount"`
	DispatchAmount float64            `json:"dispatchAmount"`
	Status         int                `json:"status"`
	TransactionId  string             `json:"transactionId"`
	Remark         string             `json:"remark"`
	PayTime        string             `json:"payTime"`
	DeliverTime    string             `json:"deliverTime"`
	FinishTime     string             `json:"finishTime"`
	Buyer          *BasicUser         `json:"buyer"`
	Address        *AddressSnapshot   `json:"address"`
	GoodsList      []*CMSOrderGoodsVO `json:"goodsList"`
}

type BasicUser struct {
	UserId   int    `json:"userId"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

type CMSOrderGoodsVO struct {
	Picture string  `json:"picture"`
	Title   string  `json:"title"`
	Price   float64 `json:"price"`
	Specs   string  `json:"specs"`
	Num     int     `json:"num"`
}

type CMSMarketMetricsVO struct {
	VisitorNum    int `json:"visitorNum"`
	SellOutSKUNum int `json:"sellOutSKUNum"`
	WaitingOrder  int `json:"waitingOrder"`
	ActivistOrder int `json:"activistOrder"`
}
