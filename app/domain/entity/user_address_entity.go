package entity

import "time"

// WechatMallUserAddressDO 商城-用户收货地址表
type WechatMallUserAddressDO struct {
	ID          int       `gorm:"column:id;primary_key;AUTO_INCREMENT"`                  // 主键
	UserID      int       `gorm:"column:user_id;default:0;NOT NULL"`                     // 用户ID
	Contacts    string    `gorm:"column:contacts;NOT NULL"`                              // 联系人
	Mobile      string    `gorm:"column:mobile;NOT NULL"`                                // 手机号
	ProvinceID  string    `gorm:"column:province_id;NOT NULL"`                           // 省份编码
	CityID      string    `gorm:"column:city_id;NOT NULL"`                               // 城市编码
	AreaID      string    `gorm:"column:area_id;NOT NULL"`                               // 地区编码
	ProvinceStr string    `gorm:"column:province_str;NOT NULL"`                          // 省份
	CityStr     string    `gorm:"column:city_str;NOT NULL"`                              // 城市
	AreaStr     string    `gorm:"column:area_str;NOT NULL"`                              // 地区
	Address     string    `gorm:"column:address;NOT NULL"`                               // 详细地址
	IsDefault   int       `gorm:"column:is_default;default:0;NOT NULL"`                  // 默认收货地址：0-否 1-是
	Del         int       `gorm:"column:is_del;default:0;NOT NULL"`                      // 是否删除：0-否 1-是
	CreateTime  time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 创建时间
	UpdateTime  time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 修改时间
}

func (m *WechatMallUserAddressDO) TableName() string {
	return "wechat_mall_user_address"
}
