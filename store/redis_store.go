package store

import (
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"time"
)

const (
	MiniappTokenPrefix = "miniappToken:"
	MiniappTokenExpire = 7200
)

type RedisStore struct {
	client *redis.Client
}

func NewRedisStore(addr, passwd string, db int) *RedisStore {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: passwd,
		DB:       db,
	})
	if _, err := client.Ping().Result(); err != nil {
		panic(err)
	}
	return &RedisStore{client: client}
}

func (r *RedisStore) GetStr(key string) (string, error) {
	d, err := r.client.Get(key).Result()
	if err == redis.Nil {
		return "", errors.New("unknown short URL")
	} else if err != nil {
		return "", err
	}
	return d, nil
}

func (r *RedisStore) SetStr(key, val string, exp int64) error {
	err := r.client.Set(key, val, time.Second*time.Duration(exp)).Err()
	return err
}
