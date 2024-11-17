package middlewares

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedis() *redis.Client {
	ctx, cancel := context.WithTimeout(context.Background(),
		30*time.Second)

	defer cancel()

	opt, err := redis.ParseURL(os.Getenv("REDIS_URI"))
	if err != nil {
		panic(err)
	}

	RedisClient = redis.NewClient(opt)

	pong, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Error connecting to Redis:", err)
	}
	fmt.Println("Connected to Redis:", pong)

	return RedisClient
}

func GetCache(ctx context.Context, redis *redis.Client, key string, endpoint string) (stringData string, err error) {

	keyRedis := key + "@" + endpoint

	return redis.Get(ctx, keyRedis).Result()
}

func SetCache(ctx context.Context, redis *redis.Client, key string, endpoint string, data interface{}, expiredTime time.Duration) error {

	keyRedis := key + "@" + endpoint

	json, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return redis.Set(ctx, keyRedis, json, expiredTime).Err()
}
