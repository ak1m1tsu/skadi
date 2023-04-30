package main

import (
	"time"

	logger "github.com/romankravchuk/skadi/internal/log"
	"github.com/romankravchuk/skadi/internal/server"
	"github.com/romankravchuk/skadi/internal/storage/redis"
)

func main() {
	logger := logger.New()
	redisClient := redis.New(&redis.Config{
		Addr:     ":6379",
		Password: "",
		DB:       0,
		TTL:      1 * time.Minute,
	})
	server := server.New(&server.Config{
		Storage: redisClient,
		Logger:  logger,
		Type:    "tcp",
		Host:    "localhost",
		Port:    "8090",
	})
	logger.Error(server.Run().Error())
}
