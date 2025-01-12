package redis

import (
	"github.com/asma12a/challenge-s6/config"
	"github.com/redis/go-redis/v9"
)

func GetClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Env.RedisURL,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return rdb
}
