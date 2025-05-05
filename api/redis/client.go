package redis

import (
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

var Client *redis.Client

var Ctx = context.Background()

func InitRedis() {
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		log.Fatal("REDIS_URL environment variable is required")
	}

	Client = redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: "", // No password set
		DB:       0,  // Default DB
	})

	_, err := Client.Ping(Ctx).Result()
	if err != nil {
		log.Fatal("Could not connect to Redis: ", err)
	}
	log.Println("Successfully connected to Redis")
}
