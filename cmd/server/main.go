package main

import (
	"log"
	"os"
	"time"

	logger "github.com/romankravchuk/skadi/internal/log"
	"github.com/romankravchuk/skadi/internal/server"
	"github.com/romankravchuk/skadi/internal/storage/redis"
	"github.com/rs/zerolog"
)

func main() {
	redisClient := redis.New(&redis.Config{
		Addr:     ":6379",
		Password: "",
		DB:       0,
		TTL:      1 * time.Minute,
	})
	server := server.New(&server.Config{
		Storage: redisClient,
		Logger:  logger.New(&logger.Config{Writer: zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2006-01-02 15:04:05"}}),
		Type:    "tcp",
		Host:    "localhost",
		Port:    "8090",
	})
	log.Fatal(server.Run())
}
