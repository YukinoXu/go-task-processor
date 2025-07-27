package cache

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/xuexiangxu/go-task-processor/internal/config"
)

var (
	RDB *redis.Client
	ctx = context.Background()
)

func InitRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     config.Cfg.RedisAddr,
		Password: "",
		DB:       0,
	})

	_, err := RDB.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Redis connected")
}

func SetIfNotExist(key string) bool {
	ok, err := RDB.SetNX(ctx, key, "1", 5*time.Minute).Result()
	if err != nil {
		log.Printf("Redis SetNX failed: %v", err)
		return false
	}
	return ok
}
