package redis

import (
	"context"
	"log"

	"github.com/asma12a/challenge-s6/config"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

// GetClient crée une nouvelle connexion Redis et vérifie si la connexion est réussie.
func GetClient() *redis.Client {

	url := config.Env.RedisURL
	opts, err := redis.ParseURL(url)
	if err != nil {
		panic(err)
	}

	rdb := redis.NewClient(opts)

	// Vérification de la connexion avec Redis
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Error connecting to redis: %v", err)
	} else {
		log.Print("✅ Successfully connected to redis!")
	}

	return rdb
}
