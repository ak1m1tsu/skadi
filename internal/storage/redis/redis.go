package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Storage interface {
	Set(ctx context.Context, key string, val interface{}) error
	Get(ctx context.Context, key string) (string, error)
}

type Config struct {
	Addr     string
	Password string
	DB       int
	TTL      time.Duration
}

type RedisStorage struct {
	client     *redis.Client
	expiration time.Duration
}

func New(cfg *Config) *RedisStorage {
	return &RedisStorage{
		client: redis.NewClient(&redis.Options{
			Addr:     cfg.Addr,
			Password: cfg.Password,
			DB:       cfg.DB,
		}),
		expiration: cfg.TTL,
	}
}

func (storage *RedisStorage) Get(ctx context.Context, key string) (string, error) {
	return storage.client.Get(ctx, key).Result()
}

func (storage *RedisStorage) Set(ctx context.Context, key string, val interface{}) error {
	return storage.client.Set(ctx, key, val, storage.expiration).Err()
}
