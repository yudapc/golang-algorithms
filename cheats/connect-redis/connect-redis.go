package connectredis

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

var ctx = context.Background()

func ConnectRedis() error {
	godotenv.Load(".env")
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("RedisHost"),
		Password: os.Getenv("RedisPassword"),
		DB:       0, // use default DB
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		return err
	}
	return nil
}
