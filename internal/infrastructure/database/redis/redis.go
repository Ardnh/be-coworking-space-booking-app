package redis

import (
	"context"
	"fmt"
	"log"

	"github.com/Ardnh/be-coworking-space-booking-app/internal/config"
	"github.com/redis/go-redis/v9"
)

type RedisDB struct {
	client *redis.Client
}

var Ctx = context.Background()

func NewRedisDB(cfg *config.Config) *redis.Client {
	hostPort := fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port)
	rdb := redis.NewClient(&redis.Options{
		Addr:     hostPort,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	// Test koneksi
	_, err := rdb.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("❌ Failed to connect to Redis: %v", err)
	}

	log.Println("✅ Redis connected successfully!")
	return rdb
}

func (r *RedisDB) Close() error {
	return r.client.Close()
}
