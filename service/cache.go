package service

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

func InitDistributedCache() *redis.Client {
	redis_uri := os.Getenv("REDIS_URI")

	rdb := redis.NewClient(&redis.Options{
		Addr:     redis_uri,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if err := rdb.Set("key", "value", 0).Err(); err != nil {
		panic(err)
	}
	if val, err := rdb.Get("key").Result(); err != nil {
		panic(err)
	} else {
		fmt.Println("key", val)
	}

	return rdb
}
