package platform

import (
	"github.com/redis/go-redis/v9"
	"web-standalone-template/pkg/configs"
)

func GetNewRedisConn() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     configs.RedisHost,
		Password: configs.RedisPassword,
		DB:       0,
	})
}
