package service

import (
	"encoding/json"
	"fmt"
	"github.com/prometheus/common/log"
	"wechat-mall-backend/dbops"
	"wechat-mall-backend/dbops/rediscli"
	"wechat-mall-backend/defs"
	"wechat-mall-backend/env"
	"wechat-mall-backend/errs"
	"wechat-mall-backend/model"
	"wechat-mall-backend/utils"
)

type IUserService interface {
	LoginCodeAuth(code string) string
	DoWxUserPhoneSignature(userId int, sessionKey, encryptedData, iv string)
	DoUserAuthInfo(userId int, req defs.WxappAuthUserInfoReq)
}

type UserService struct {
	Conf *env.Conf
}

func NewUserService(conf *env.Conf) IUserService {
	return &UserService{Conf: conf}
}

func (s *UserService) LoginCodeAuth(code string) string {
	baseUrl := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	url := fmt.Sprintf(baseUrl, s.Conf.Wxapp.Appid, s.Conf.Wxapp.Appsecret, code)

	tmpVal, err := utils.HttpGet(url)
	if err != nil {
		panic(err)
	}

	result := make(map[string]interface{})
	err = json.Unmarshal([]byte(tmpVal), &result)
	if err != nil {
		panic(errs.ErrorWechatError)
	}
	// {"session_key":"TppZM2zEd6\/dGzkqbbrriQ==","expires_in":7200,"openid":"oQOru0EUuLdidBZH0r_F8fDURPjI"}
	if result["errcode"] != nil {
		log.Error("微信内部异常：", result)
		panic(errs.ErrorWechatError)
	}
	userId := registerUser(result["openid"].(string))
	token, err := utils.CreateToken(userId, defs.AccessTokenExpire)
	if err != nil {
		panic(err)
	}
	err = rediscli.SetStr(defs.MiniappTokenPrefix+token, tmpVal, defs.AccessTokenExpire)
	if err != nil {
		panic(errs.ErrorRedisError)
	}
	return token
}

func registerUser(openid string) int {
	user, err := dbops.GetUserByOpenid(openid)
	if err != nil {
		panic(err)
	}
	if user.Id == 0 {
		user = &model.WechatMallUserDO{Openid: openid}
		uid, err := dbops.AddMiniappUser(user)
		if err != nil {
			panic(err)
		}
		user.Id = model.ID(uid)
	}
	return user.Id
}

func (s *UserService) DoWxUserPhoneSignature(userId int, sessionKey, encryptedData, iv string) {
	appid := s.Conf.Wxapp.Appid
	wxSensitiveData := utils.WxSensitiveData{AppId: appid, SessionKey: sessionKey, Iv: iv, EncryptedData: encryptedData}
	decrypt, err := wxSensitiveData.Decrypt()
	if err != nil {
		panic(err)
	}
	userDO, err := dbops.GetUserById(userId)
	if err != nil {
		panic(err)
	}
	userDO.Mobile = decrypt["phoneNumber"].(string)
	err = dbops.UpdateUserById(userDO)
	if err != nil {
		panic(err)
	}
}

func (s *UserService) DoUserAuthInfo(userId int, req defs.WxappAuthUserInfoReq) {
	userDO, err := dbops.GetUserById(userId)
	if err != nil {
		panic(err)
	}
	userDO.Nickname = req.NickName
	userDO.Avatar = req.AvatarUrl
	userDO.Gender = req.Gender
	userDO.Country = req.Country
	userDO.Province = req.Province
	userDO.City = req.City
	err = dbops.UpdateUserById(userDO)
	if err != nil {
		panic(err)
	}
}
