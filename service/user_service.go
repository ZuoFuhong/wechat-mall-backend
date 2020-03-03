package service

import (
	"encoding/json"
	"fmt"
	"github.com/prometheus/common/log"
	"wechat-mall-backend/dbops"
	"wechat-mall-backend/defs"
	"wechat-mall-backend/env"
	"wechat-mall-backend/errs"
	"wechat-mall-backend/utils"
)

type IUserService interface {
	LoginCodeAuth(code string) (defs.WxappLoginResp, error)
}

type UserService struct {
	Conf *env.Conf
}

func NewUserService(conf *env.Conf) IUserService {
	return &UserService{Conf: conf}
}

func (service *UserService) LoginCodeAuth(code string) (defs.WxappLoginResp, error) {
	baseUrl := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	url := fmt.Sprintf(baseUrl, service.Conf.Wxapp.Appid, service.Conf.Wxapp.Appsecret, code)

	tmpVal, err := utils.HttpGet(url)
	if err != nil {
		panic(err)
	}

	result := make(map[string]interface{})
	err = json.Unmarshal([]byte(tmpVal), &result)
	if err != nil {
		panic(errs.ErrorWechatError)
	}
	if result["errcode"] != nil {
		log.Error("微信内部异常：", result)
		panic(errs.ErrorWechatError)
	}

	// {"session_key":"TppZM2zEd6\/dGzkqbbrriQ==","expires_in":7200,"openid":"oQOru0EUuLdidBZH0r_F8fDURPjI"}
	token := utils.RandomStr(32)
	err = dbops.SetStr(dbops.MiniappTokenPrefix+token, tmpVal, dbops.MiniappTokenExpire)
	if err != nil {
		panic(errs.ErrorRedisError)
	}
	registerUser(result["openid"].(string))

	resp := defs.WxappLoginResp{Token: token, ExpirationInMinutes: dbops.MiniappTokenExpire}
	return resp, nil
}

func registerUser(openid string) {
	user, err := dbops.GetUserByOpenid(openid)
	if err != nil {
		panic(err)
	}
	if user.Id == 0 {
		newUser := &WxappUser{Openid: openid, Nickname: "", Avatar: "", Mobile: "", City: ""}
		_, err := dbops.AddMiniappUser((*dbops.WxappUser)(newUser))
		if err != nil {
			panic(err)
		}
	}
}
