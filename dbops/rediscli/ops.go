package rediscli

import (
	"github.com/go-redis/redis"
	"time"
)

var client *redis.Client

func InitRedisCli(addr, passwd string, db int) {
	client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: passwd,
		DB:       db,
	})
	if _, err := client.Ping().Result(); err != nil {
		panic(err)
	}
}

func GetStr(key string) (string, error) {
	d, err := client.Get(key).Result()
	if err == redis.Nil {
		return "", err
	} else if err != nil {
		return "", err
	}
	return d, nil
}

func SetStr(key, val string, exp int64) error {
	err := client.Set(key, val, time.Second*time.Duration(exp)).Err()
	return err
}
