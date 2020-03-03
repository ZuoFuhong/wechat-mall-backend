package service

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"wechat-mall-backend/dbops"
	"wechat-mall-backend/defs"
	"wechat-mall-backend/utils"
)

type ICMSUserService interface {
	CMSLoginValidate(username, password string) (*CMSUser, error)
	CMSUserRegister(registerReq *defs.CMSRegisterReq) error
	AddCMSUser(username, password, email string) error
}

type CMSUserService struct {
}

func NewCMSUserService() ICMSUserService {
	service := &CMSUserService{}
	return service
}

func (cus *CMSUserService) CMSLoginValidate(username, password string) (*CMSUser, error) {
	user, err := dbops.GetCMSUserByUsername(username)
	if err != nil {
		return nil, err
	}
	if user.Id == 0 {
		return nil, errors.New("用户名不存在！")
	}
	encrpytStr := utils.Md5Encrpyt(password)
	if user.Password != encrpytStr {
		return nil, errors.New("密码错误！")
	}
	return (*CMSUser)(user), nil
}

func (cus *CMSUserService) CMSUserRegister(registerReq *defs.CMSRegisterReq) error {
	user, err := dbops.GetCMSUserByUsername(registerReq.Username)
	if err != nil {
		return err
	}
	if user.Id != 0 {
		return errors.New("用户名已注册！")
	}
	user, err = dbops.GetCMSUserByEmail(registerReq.Email)
	if err != nil {
		return err
	}
	if user.Id != 0 {
		return errors.New("邮箱已注册！")
	}
	code := utils.RandomStr(32)
	data, _ := json.Marshal(registerReq)
	_ = dbops.SetStr(dbops.CMSCodePrefix+code, string(data), dbops.CMSCodeExpire)

	go sendEmailValidate(registerReq.Email, code)
	return nil
}

func sendEmailValidate(email, code string) {
	// todo: 邮箱验证，账号激活
	fmt.Printf("发送验证短信 email = %s, code = %s", email, code)
}

func (cus *CMSUserService) AddCMSUser(username, password, email string) error {
	encrpytStr := utils.Md5Encrpyt(password)
	user := CMSUser{Username: username, Password: encrpytStr, Email: email}
	return dbops.AddCMSUser((*dbops.CMSUser)(&user))
}
