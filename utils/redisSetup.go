package utils

import "github.com/go-redis/redis"

func getRedisConnection() *redis.Client {
	redis := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	return redis
}
