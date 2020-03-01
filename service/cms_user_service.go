package service

import "wechat-mall-web/store"

type ICMSUserService interface {
	GetByUsername(username string) (*CMSUser, error)
	Create(username, password, email string) (int, error)
}

type CMSUserService struct {
	DbStore    *store.MySQLStore
	RedisStore *store.RedisStore
}

func NewCMSUserService(dbStore *store.MySQLStore, redisStore *store.RedisStore) *CMSUserService {
	return &CMSUserService{DbStore: dbStore, RedisStore: redisStore}
}

func (cus *CMSUserService) GetByUsername(username string) (*CMSUser, error) {
	// todo:

	return nil, nil
}

func (cus *CMSUserService) Create(username, password, email string) (int, error) {
	// todo:

	return 0, nil
}
