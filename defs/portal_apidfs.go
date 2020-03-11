package defs

type WxappLoginResp struct {
	Token               string `json:"token" validate:"required"`
	ExpirationInMinutes int    `json:"expiration_in_minutes" validate:"required"`
}

type PortalBannerVO struct {
	Id      int    `json:"id"`
	Picture string `json:"picture"`
}

type PortalGridCategoryVO struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	CategoryId int    `json:"category"`
	Picture    string `json:"picture"`
}

type PortalCouponVO struct {
	Id          int    `json:"id"`          // 优惠券ID
	Title       string `json:"title"`       // 标题
	FullMoney   string `json:"full_money"`  // 满减额
	Minus       string `json:"minus"`       // 优惠额
	Rate        string `json:"rate"`        // 折扣
	Type        int    `json:"type"`        // 券类型：1-满减券 2-折扣券 3-无门槛券 4-满金额折扣券
	StartTime   string `json:"start_time"`  // 开始时间
	EndTime     string `json:"end_time"`    // 结束时间
	Description string `json:"description"` // 描述
	Status      int    `json:"status"`      // 领取状态：
}

type TakeCouponReq struct {
	UserId   int `json:"user_id" validate:"required"`
	CouponId int `json:"coupon_id" validate:"required"`
}

type PortalUserCouponVO struct {
	CLogId      int    `json:"c_log_id"`    // 领取记录ID
	CouponId    int    `json:"coupon_id"`   // 优惠券ID
	Title       string `json:"title"`       // 标题
	FullMoney   string `json:"full_money"`  // 满减额
	Minus       string `json:"minus"`       // 优惠额
	Rate        string `json:"rate"`        // 折扣
	Type        int    `json:"type"`        // 券类型：1-满减券 2-折扣券 3-无门槛券 4-满金额折扣券
	StartTime   string `json:"start_time"`  // 开始时间
	EndTime     string `json:"end_time"`    // 结束时间
	Description string `json:"description"` // 描述
}
