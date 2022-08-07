package entity

import "time"

// WechatMallUserDO 小程序-用户表
type WechatMallUserDO struct {
	ID         int       `gorm:"column:id;primary_key;AUTO_INCREMENT"`                  // 主键
	Openid     string    `gorm:"column:openid;NOT NULL"`                                // 微信openid
	Nickname   string    `gorm:"column:nickname;NOT NULL"`                              // 昵称
	Avatar     string    `gorm:"column:avatar;NOT NULL"`                                // 微信头像
	Mobile     string    `gorm:"column:mobile;NOT NULL"`                                // 手机号
	City       string    `gorm:"column:city;NOT NULL"`                                  // 城市
	Province   string    `gorm:"column:province;NOT NULL"`                              // 省份
	Country    string    `gorm:"column:country;NOT NULL"`                               // 国家
	Gender     int       `gorm:"column:gender;default:0;NOT NULL"`                      // 性别 0：未知、1：男、2：女
	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 更新时间
}

func (m *WechatMallUserDO) TableName() string {
	return "wechat_mall_user"
}
