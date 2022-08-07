package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/allegro/bigcache/v3"
	"github.com/pkg/errors"
	"time"
	"wechat-mall-backend/app/domain/entity"
	"wechat-mall-backend/app/domain/repository"
	"wechat-mall-backend/consts"
	"wechat-mall-backend/pkg/config"
	"wechat-mall-backend/pkg/log"
	"wechat-mall-backend/pkg/utils"
)

type IUserService interface {
	LoginCodeAuth(ctx context.Context, code string) (string, int, error)

	DoWxUserPhoneSignature(ctx context.Context, userId int, accessToken, encryptedData, iv string) error

	DoUserAuthInfo(ctx context.Context, userDO *entity.WechatMallUserDO) error

	DoAddVisitorRecord(ctx context.Context, userId int, ip string)

	QueryTodayUniqueVisitor(ctx context.Context) int

	QueryUserInfo(ctx context.Context, userId int) (*entity.WechatMallUserDO, error)
}

type UserService struct {
	repos      repository.IUserRepos
	tokenCache *bigcache.BigCache
}

func NewUserService(repos repository.IUserRepos) IUserService {
	cache, err := bigcache.NewBigCache(bigcache.DefaultConfig(consts.AccessTokenExpire * time.Second))
	if err != nil {
		log.Fatalf("bigcache.NewBigCache failed, err: %v", err)
	}
	return &UserService{
		repos:      repos,
		tokenCache: cache,
	}
}

func (s *UserService) LoginCodeAuth(ctx context.Context, code string) (string, int, error) {
	wxapp := config.GlobalConfig().Wxapp
	baseUrl := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	url := fmt.Sprintf(baseUrl, wxapp.Appid, wxapp.Appsecret, code)
	tmpVal, err := utils.HttpGet(url)
	if err != nil {
		log.ErrorContextf(ctx, "call HttpGet failed, err: %v", err)
		return "", 0, err
	}
	data := make(map[string]interface{})
	err = json.Unmarshal([]byte(tmpVal), &data)
	if err != nil {
		return "", 0, err
	}
	// {"session_key":"TppZM2zEd6\/dGzkqbbrriQ==","expires_in":7200,"openid":"oQOru0EUuLdidBZH0r_F8fDURPjI"}
	if data["errcode"] != nil {
		return "", 0, errors.New("wechat internal error")
	}
	userId, err := s.registerUser(ctx, data["openid"].(string))
	if err != nil {
		return "", 0, err
	}
	token, err := utils.CreateToken(userId, consts.AccessTokenExpire)
	if err != nil {
		return "", 0, err
	}
	_ = s.tokenCache.Set(consts.MiniappTokenPrefix+token, []byte(tmpVal))
	return token, userId, nil
}

func (s *UserService) registerUser(ctx context.Context, openid string) (int, error) {
	user, err := s.repos.GetUserByOpenid(ctx, openid)
	if err != nil {
		return 0, err
	}
	if user.ID == 0 {
		user = &entity.WechatMallUserDO{
			Openid:     openid,
			Nickname:   "新用户",
			CreateTime: time.Now(),
		}
		uid, err := s.repos.AddUser(ctx, user)
		if err != nil {
			return 0, err
		}
		user.ID = uid
	}
	return user.ID, nil
}

func (s *UserService) DoWxUserPhoneSignature(ctx context.Context, userId int, accessToken, encryptedData, iv string) error {
	cacheData, err := s.tokenCache.Get(consts.MiniappTokenPrefix + accessToken)
	if err != nil {
		return errors.New("invalid access_token")
	}
	data := make(map[string]interface{})
	if err = json.Unmarshal(cacheData, &data); err != nil {
		return errors.New("invalid access_token")
	}
	wxapp := config.GlobalConfig().Wxapp
	wxSensitiveData := utils.WxSensitiveData{AppId: wxapp.Appid, SessionKey: data["session_key"].(string), Iv: iv, EncryptedData: encryptedData}
	decrypt, err := wxSensitiveData.Decrypt()
	if err != nil {
		return err
	}
	userDO, err := s.repos.GetUserById(ctx, userId)
	if err != nil {
		return err
	}
	if userDO.ID == consts.ZERO {
		return errors.New("not found user record")
	}
	userDO.Mobile = decrypt["phoneNumber"].(string)
	return s.repos.UpdateUser(ctx, userDO)
}

func (s *UserService) DoUserAuthInfo(ctx context.Context, user *entity.WechatMallUserDO) error {
	userDO, err := s.repos.GetUserById(ctx, user.ID)
	if err != nil {
		return err
	}
	if userDO.ID == consts.ZERO {
		return errors.New("not found user record")
	}
	userDO.Nickname = user.Nickname
	userDO.Avatar = user.Avatar
	userDO.Gender = user.Gender
	userDO.Country = user.Country
	userDO.Province = user.Province
	userDO.City = user.City
	return s.repos.UpdateUser(ctx, userDO)
}

func (s *UserService) DoAddVisitorRecord(ctx context.Context, userId int, ip string) {
	err := s.repos.AddVisitorRecord(ctx, userId, ip)
	if err != nil {
		// ignore error
		log.ErrorContextf(ctx, "call AddVisitorRecord failed, err: %v", err)
	}
}

func (s *UserService) QueryTodayUniqueVisitor(ctx context.Context) int {
	todayStr := utils.FormatDatetime(time.Now(), utils.YYYYMMDD)
	today, _ := utils.ParseDatetime(todayStr, utils.YYYYMMDD)
	startTime := time.Unix(today.Unix()-28800, 0)
	endTime := time.Unix(today.Unix()+57600, 0)
	total, err := s.repos.CountUniqueVisitor(ctx, startTime, endTime)
	if err != nil {
		// ignore error
		log.ErrorContextf(ctx, "call CountUniqueVisitor failed, err: %v", err)
	}
	return total
}

func (s *UserService) QueryUserInfo(ctx context.Context, userId int) (*entity.WechatMallUserDO, error) {
	return s.repos.GetUserById(ctx, userId)
}
