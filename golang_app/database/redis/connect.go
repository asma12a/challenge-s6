package redis

import (
	"github.com/redis/go-redis/v9"
)

func GetClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	
	return rdb
}
