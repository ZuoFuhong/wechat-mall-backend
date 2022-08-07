package entity

import "time"

// WechatMallCMSUserDO CMS后台用户表
type WechatMallCMSUserDO struct {
	ID         int       `gorm:"column:id;primary_key;AUTO_INCREMENT"`                  // 主键
	Username   string    `gorm:"column:username;NOT NULL"`                              // 用户名
	Password   string    `gorm:"column:password;NOT NULL"`                              // 密码
	Email      string    `gorm:"column:email;NOT NULL"`                                 // 邮箱
	Mobile     string    `gorm:"column:mobile;NOT NULL"`                                // 手机号
	Avatar     string    `gorm:"column:avatar;NOT NULL"`                                // 头像
	GroupID    int       `gorm:"column:group_id;default:0;NOT NULL"`                    // 分组ID
	Del        int       `gorm:"column:is_del;default:0;NOT NULL"`                      // 是否删除：0-否 1-是
	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 修改时间
}

func (m *WechatMallCMSUserDO) TableName() string {
	return "wechat_mall_cms_user"
}
