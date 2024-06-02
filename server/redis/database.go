package redis

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	client *redis.Client
	once   sync.Once
)

func InitRedis() {
	once.Do(func() {
		client = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err := client.Ping(ctx).Err()
		if err != nil {
			fmt.Println("Could not ping Redis:", err)
		} else {
			fmt.Println("Redis ping successful")
		}
	})
}

func RedisDatabase() *redis.Client {
	return client
}
