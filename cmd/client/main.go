package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/romankravchuk/skadi/internal/client"
	"github.com/romankravchuk/skadi/internal/log"
	"github.com/romankravchuk/skadi/internal/skadi"
)

func main() {
	logger := log.New()
	client := client.New(&client.Config{
		ConnectionType: "tcp",
		Host:           "localhost",
		Port:           "8090",
		Logger:         logger,
	})
	client.Connect()
	defer client.Close()
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		msg, _ := skadi.Encode(&skadi.Request{
			CommandCode: skadi.Close,
			Body:        []byte(strings.Replace(text, "\n", "", -1)),
		})
		resp, err := client.Send(msg)
		if err != nil {
			logger.Error(err.Error())
		}
		if string(resp) == "Bye!" {
			logger.Info("Connection closed")
			break
		}
	}
}
