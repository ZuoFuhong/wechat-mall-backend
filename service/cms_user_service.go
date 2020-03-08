package service

import (
	"encoding/json"
	"fmt"
	"wechat-mall-backend/dbops"
	"wechat-mall-backend/defs"
	"wechat-mall-backend/errs"
	"wechat-mall-backend/utils"
)

type ICMSUserService interface {
	CMSLoginValidate(username, password string) *CMSUser
	CMSUserRegister(registerReq *defs.CMSRegisterReq)
	AddCMSUser(username, password, email string) error
}

type CMSUserService struct {
}

func NewCMSUserService() ICMSUserService {
	service := &CMSUserService{}
	return service
}

func (cus *CMSUserService) CMSLoginValidate(username, password string) *CMSUser {
	user, err := dbops.GetCMSUserByUsername(username)
	if err != nil {
		panic(err)
	}
	if user.Id == 0 {
		panic(errs.NewAuthUserError("用户名不存在！"))
	}
	encrpytStr := utils.Md5Encrpyt(password)
	if user.Password != encrpytStr {
		panic(errs.NewAuthUserError("密码错误！"))
	}
	return (*CMSUser)(user)
}

func (cus *CMSUserService) CMSUserRegister(registerReq *defs.CMSRegisterReq) {
	user, err := dbops.GetCMSUserByUsername(registerReq.Username)
	if err != nil {
		panic(err)
	}
	if user.Id != 0 {
		panic(errs.NewAuthUserError("用户名已注册！"))
	}
	user, err = dbops.GetCMSUserByEmail(registerReq.Email)
	if err != nil {
		panic(err)
	}
	if user.Id != 0 {
		panic(errs.NewAuthUserError("邮箱已注册！"))
	}
	code := utils.RandomStr(32)
	data, _ := json.Marshal(registerReq)
	_ = dbops.SetStr(dbops.CMSCodePrefix+code, string(data), dbops.CMSCodeExpire)

	go sendEmailValidate(registerReq.Email, code)
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
