package utils

import "github.com/go-redis/redis/v9"

func RedisConnection() *redis.Client {
	redis := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return redis
}
