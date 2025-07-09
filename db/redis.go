package db

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedis() (err error) {
	// Start redis connection
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	// Check if connection is successful
	_, err = RedisClient.Ping(context.TODO()).Result()

	return
}

func DisconnectRedis() {
	RedisClient.Close()
}
