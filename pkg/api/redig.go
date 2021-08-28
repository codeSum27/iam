package api

import (
	"github.com/go-redis/redis"
	"os"
)

var client *redis.Client

func RedisInit() {
	dsn := os.Getenv("REDIS_DSN")
	if len(dsn) == 0 {
		dsn = "localhost:6379"
	}

	client = redis.NewClient(&redis.Options{
		Addr: dsn,
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
}

func GetRedisClient() *redis.Client {
	return client
}