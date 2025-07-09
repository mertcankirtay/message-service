package db

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func InitRedis() (err error) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	_, err = redisClient.Ping(context.TODO()).Result()

	return
}

func DisconnectRedis() {
	redisClient.Close()
}
