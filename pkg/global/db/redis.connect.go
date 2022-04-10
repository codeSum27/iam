package db

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

var client *redis.Client

func RedisInit() {
	fmt.Println("Start to initialize Redis Connection")

	dsn := os.Getenv("REDIS_HOST")
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

	fmt.Println("Success to initialize Redis Connection")
}

func GetRedisClient() *redis.Client {
	return client
}