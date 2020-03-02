package service

//import (
//	"errors"
//	"github.com/dgrijalva/jwt-go"
//	"github.com/prometheus/common/log"
//	"time"
//	"wechat-mall-web/store"
//)
//
//type ITokenService interface {
//	CreateToken(uid ,expirationInMinutes int) (string, error)
//	ValidateToken(tokenStr string) bool
//	ParseToken(tokenStr string) (*jwt.StandardClaims, error)
//}
//
//type tokenService struct {
//	secretKey  string
//	redisStore *store.RedisStore
//}
//
//func NewTokenService (redisStore *store.RedisStore) ITokenService {
//	return &tokenService{
//		secretKey: "123456",
//		redisStore: redisStore,
//	}
//}
//
//func (ts *tokenService) CreateToken(uid, expirationInMinutes int) (string, error) {
//	claims := jwt.StandardClaims {
//		ExpiresAt: int64(time.Now().Add(time.Hour * time.Duration(expirationInMinutes)).Unix()),
//		Issuer: "dazuo",
//		NotBefore: time.Now().Unix(),
//		IssuedAt: time.Now().Unix(),
//	}
//	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(ts.secretKey))
//	if err != nil {
//		log.Error(err)
//		return "", errors.New("token无效")
//	}
//	return token, nil
//}
//
//func (ts *tokenService) ValidateToken (tokenStr string) bool {
//	token, err := jwt.ParseWithClaims(tokenStr, &jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, e error) {
//		return []byte(ts.secretKey), nil
//	})
//	if err != nil {
//		log.Error(err)
//		return false
//	}
//	return token.Valid
//}
//
//func (ts *tokenService) ParseToken (tokenStr string) (*jwt.StandardClaims, error) {
//	token, err := jwt.ParseWithClaims(tokenStr, &jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, e error) {
//		return []byte(ts.secretKey), nil
//	})
//	if err != nil {
//		log.Error(err)
//		return nil, errors.New("token解析失败")
//	}
//	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
//		return claims, nil
//	} else {
//		return nil, errors.New("token无效")
//	}
//}
