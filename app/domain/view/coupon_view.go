package view

type PortalUserCouponVO struct {
	CLogId      int    `json:"CLogId"`      // 领取记录ID
	CouponId    int    `json:"couponId"`    // 优惠券ID
	Title       string `json:"title"`       // 标题
	FullMoney   string `json:"fullMoney"`   // 满减额
	Minus       string `json:"minus"`       // 优惠额
	Rate        string `json:"rate"`        // 折扣
	Type        int    `json:"type"`        // 券类型：1-满减券 2-折扣券 3-代金券 4-满金额折扣券
	StartTime   string `json:"startTime"`   // 开始时间
	EndTime     string `json:"endTime"`     // 结束时间
	Description string `json:"description"` // 描述
}

type CMSCouponVO struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	FullMoney   string `json:"fullMoney"`
	Minus       string `json:"minus"`
	Rate        string `json:"rate"`
	Type        int    `json:"type"`
	GrantNum    int    `json:"grantNum"`
	LimitNum    int    `json:"limitNum"`
	StartTime   string `json:"startTime"`
	EndTime     string `json:"endTime"`
	Description string `json:"description"`
	Online      int    `json:"online"`
}

type CMSUserVO struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Mobile    string `json:"mobile"`
	Avatar    string `json:"avatar"`
	GroupId   int    `json:"groupId"`
	GroupName string `json:"groupName"`
}

type PortalCouponVO struct {
	Id          int    `json:"id"`          // 优惠券ID
	Title       string `json:"title"`       // 标题
	FullMoney   string `json:"fullMoney"`   // 满减额
	Minus       string `json:"minus"`       // 优惠额
	Rate        string `json:"rate"`        // 折扣
	Type        int    `json:"type"`        // 券类型：1-满减券 2-折扣券 3-代金券 4-满金额折扣券
	StartTime   string `json:"startTime"`   // 开始时间
	EndTime     string `json:"endTime"`     // 结束时间
	Description string `json:"description"` // 描述
}
