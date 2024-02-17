package handlers

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func getRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis-container:6379",
		Password: "",
		DB:       0,
	})
	return client
}

func setRedisData(client *redis.Client, key string, value string) {
	ctx := context.Background()
	err := client.Set(ctx, key, value, 0).Err()
	if err != nil {
		panic(err)
	}
}

func getRedisData(client *redis.Client, key string) string {
	ctx := context.Background()
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		if err.Error() == "redis: nil" {
			return ""
		} else {
			panic(err)
		}
	}
	return val
}
