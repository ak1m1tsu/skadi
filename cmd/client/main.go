package main

import (
	"os"

	"github.com/romankravchuk/skadi/internal/client"
	"github.com/romankravchuk/skadi/internal/log"
	"github.com/rs/zerolog"
)

func main() {
	client := client.New(&client.Config{
		ConnectionType: "tcp",
		Host:           "localhost",
		Port:           "8090",
		Logger: log.New(&log.Config{
			Writer: zerolog.ConsoleWriter{
				Out:        os.Stdout,
				TimeFormat: "2006-01-02 15:04:05",
			},
		}),
	})
	client.Connect()
	defer client.Close()
	client.Send("hello world")
}
