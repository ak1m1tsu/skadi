package client

import (
	"fmt"
	"net"

	"github.com/romankravchuk/skadi/internal/log"
)

type Client struct {
	conn     net.Conn
	logger   log.Logger
	connType string
	connAddr string
}

type Config struct {
	ConnectionType string
	Host           string
	Port           string
	Logger         log.Logger
}

func New(cfg *Config) *Client {
	return &Client{
		logger:   cfg.Logger,
		connType: cfg.ConnectionType,
		connAddr: cfg.Host + ":" + cfg.Port,
	}
}

func (c *Client) Connect() error {
	var err error
	c.conn, err = net.Dial(c.connType, c.connAddr)
	if err != nil {
		return fmt.Errorf("error connecting to server: %w", err)
	}
	return nil
}

func (c *Client) Close() error {
	return fmt.Errorf("error closing connection: %w", c.conn.Close())
}

func (c *Client) Send(msg []byte) ([]byte, error) {
	buffer := make([]byte, 1024)
	_, err := c.conn.Write(msg)
	if err != nil {
		return nil, fmt.Errorf("error sending message: %w", err)
	}
	mLen, err := c.conn.Read(buffer)
	if err != nil {
		return nil, fmt.Errorf("error reading message: %w", err)
	}
	c.logger.Info(string(buffer[:mLen]))
	return buffer[:mLen], nil
}
