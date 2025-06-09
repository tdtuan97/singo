package cache

import (
	"context"
	"os"
	"singo/util"
	"strconv"

	"github.com/redis/go-redis/v9"
)

// RedisClient Redis cache client singleton
var RedisClient *redis.Client

// Redis Initialize redis connection in middleware
func Redis() {
	db, _ := strconv.ParseUint(os.Getenv("REDIS_DB"), 10, 64)
	addr := os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT")
	client := redis.NewClient(&redis.Options{
		Addr:       addr,
		Password:   os.Getenv("REDIS_PASSWORD"),
		DB:         int(db),
		MaxRetries: 1,
	})

	_, err := client.Ping(context.Background()).Result()

	if err != nil {
		util.Log().Panic("Failed to connect to Redis", err)
	}

	RedisClient = client
}
