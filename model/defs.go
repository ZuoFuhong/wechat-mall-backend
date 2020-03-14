package model

type ID = int

// 小程序用户
type WechatMallUserDO struct {
	Id         ID
	Openid     string // 微信openid
	Nickname   string // 昵称
	Avatar     string // 头像
	Mobile     string // 手机号
	City       string // 城市编码
	Province   string // 省份
	Country    string // 国家
	Gender     int    // 性别 0：未知、1：男、2：女
	CreateTime string
	UpdateTime string
}

// CMS后台用户
type WechatMallCMSUserDO struct {
	Id         ID
	Username   string // 用户名
	Password   string // 密码
	Email      string // 邮箱
	Mobile     string // 手机号
	Avatar     string // 头像
	GroupId    int    // 用户分组
	Del        int
	CreateTime string
	UpdateTime string
}

// CMS后台用户分组
type WechatMallUserGroupDO struct {
	Id          ID
	Name        string // 名称
	Description string // 描述
	Del         int
	CreateTime  string
	UpdateTime  string
}

// CMS-组成模块
type WechatMallModuleDO struct {
	Id          ID
	Name        string // 名称
	Description string // 描述
	Del         int
	CreateTime  string
	UpdateTime  string
}

type WechatMallModulePageDO struct {
	Id          ID
	ModuleId    ID     // 模块ID
	Name        string // 名称
	Description string // 描述
	Del         int
	CreateTime  string
	UpdateTime  string
}

type WechatMallGroupPagePermission struct {
	Id         ID
	GroupId    ID // 分组ID
	PageId     ID // 页面ID
	Del        int
	CreateTime string
	UpdateTime string
}

// 商城Banner
type WechatMallBannerDO struct {
	Id          ID
	Picture     string // 图片
	Name        string // 名称
	Title       string // 标题
	Description string // 描述
	Del         int
	CreateTime  string
	UpdateTime  string
}

// 商城-分类
type WechatMallCategoryDO struct {
	Id          ID
	ParentId    int    // 父级分类
	Name        string // 分类名称
	Sort        int    // 排序
	Online      int    // 是否上线
	Picture     string // 图标
	Description string // 描述
	Del         int
	CreateTime  string
	UpdateTime  string
}

// 小程序-宫格
type WechatMallGridCategoryDO struct {
	Id         ID
	Title      string // 标题
	Name       string // 名称
	CategoryId int    // 分类ID
	Picture    string // 图标
	Del        int
	CreateTime string
	UpdateTime string
}

// 商城-规格
type WechatMallSpecificationDO struct {
	Id          ID
	Name        string // 名称
	Description string // 描述
	Unit        string // 单位
	Standard    int    // 是否标准
	Del         int
	CreateTime  string
	UpdateTime  string
}

// 商城-规格属性表
type WechatMallSpecificationAttrDO struct {
	Id         ID
	SpecId     ID     // 规格主键
	Value      string // 规格值
	Extend     string // 扩展
	Del        int
	CreateTime string
	UpdateTime string
}

// 商城-商品
type WechatMallGoodsDO struct {
	Id            ID
	BrandName     string // 品牌
	Title         string // 标题
	Price         string // 价格
	DiscountPrice string // 折扣
	CategoryId    int    // 分类ID
	Online        int    // 是否上线
	Picture       string // 主图
	BannerPicture string // 详情图
	DetailPicture string // 轮播图
	Tags          string // 标签
	Description   string // 详情
	Del           int
	CreateTime    string
	UpdateTime    string
}

// 商城-商品规格
type WechatMallGoodsSpecDO struct {
	Id         ID
	GoodsId    ID  // 商品ID
	SpecId     int // 规格ID
	Del        int
	CreateTime string
	UpdateTime string
}

// 商城-SKU
type WechatMallSkuDO struct {
	Id         ID
	Title      string // 标题
	Price      string // 价格
	Code       string // 编码
	Stock      int    // 库存量
	GoodsId    int    // 商品ID
	Online     int    // 是否上线
	Picture    string // 图片
	Specs      string // 多规格属性
	Del        int
	CreateTime string
	UpdateTime string
}

// 商城-优惠券
type WechatMallCouponDO struct {
	Id          ID
	Title       string // 标题
	FullMoney   string // 满减额
	Minus       string // 优惠金额
	Rate        string // 折扣
	Type        int    // 类型
	StartTime   string // 开始时间
	EndTime     string // 截止时间
	Description string // 规则描述
	Online      int    // 是否上线
	Del         int
	CreateTime  string
	UpdateTime  string
}

// 商城-优惠券领取记录
type WechatMallCouponLogDO struct {
	Id         ID
	CouponId   ID     // 优惠券ID
	UserId     ID     // 用户ID
	UseTime    string // 核销时间
	ExpireTime string // 过期时间
	Status     int    // 状态：0-未使用 1-已使用 2-已过期
	Code       string // 券码
	OrderNo    string // 核销的订单号
	Del        int    // 是否删除
	CreateTime string
	UpdateTime string
}

// 商城-购物车
type WechatMallUserCartDO struct {
	Id         ID
	UserId     ID  // 用户ID
	GoodsId    ID  // 商品ID
	SkuId      ID  // sku ID
	Num        int // 数量
	Del        int
	CreateTime string
	UpdateTime string
}

// 商城-用户收货地址
type WechatMallUserAddressDO struct {
	Id          ID
	UserId      ID     // 用户ID
	Contacts    string // 联系人
	Mobile      string // 手机号
	ProvinceId  string // 省份编码
	CityId      string // 城市编码
	AreaId      string // 地区编码
	ProvinceStr string // 省份
	CityStr     string // 城市
	AreaStr     string // 地区
	Address     string // 详细地址
	IsDefault   int    // 默认收货地址：0-否 1-是
	Del         int
	CreateTime  string
	UpdateTime  string
}

// 商城-订单表
type WechatMallOrderDO struct {
	Id              ID
	OrderNo         string // 订单号
	UserId          ID     // 用户ID
	PayAmount       string // 付款金额
	GoodsAmount     string // 商品小计
	DiscountAmount  string // 优惠金额
	DispatchAmount  string // 运费
	PayTime         string // 支付时间
	Status          int    // 状态 -1 已取消 0-待付款 1-待发货 2-待收货 3-已完成
	AddressId       int    // 地址ID
	AddressSnapshot string // 地址快照
	WxappPrePayId   string // 微信预支付ID
	Del             int
	CreateTime      string
	UpdateTime      string
}

type WechatMallOrderGoodsDO struct {
	Id         ID
	OrderNo    string // 订单号
	GoodsId    int    // 商品ID
	SkuId      int    // sku ID
	Picture    string // 图片
	Title      string // 标题
	Price      string // 价格
	Specs      string // sku属性
	Num        int    // 数量
	LockStatus int    // 锁定状态：0-预定 1-付款 2-取消
	CreateTime string
	UpdateTime string
}
