package server

import (
	"fmt"
	"net"
	"net/http"

	"github.com/romankravchuk/skadi/internal/log"
	"github.com/romankravchuk/skadi/internal/storage/redis"
)

type Server struct {
	http.Handler
	connType   string
	listenAddr string
	storage    redis.Storage
	logger     log.Logger
}

type Config struct {
	Storage redis.Storage
	Logger  log.Logger
	Type    string
	Host    string
	Port    string
}

func New(cfg *Config) *Server {
	return &Server{
		listenAddr: cfg.Host + ":" + cfg.Port,
		storage:    cfg.Storage,
		logger:     cfg.Logger,
		connType:   cfg.Type,
	}
}

func (s *Server) Run() error {
	listener, err := net.Listen(s.connType, s.listenAddr)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}
	defer listener.Close()
	s.logger.Info(fmt.Sprintf("Listening on %s", s.listenAddr))
	s.logger.Info("Waiting for connections...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			return fmt.Errorf("failed to accept: %w", err)
		}
		s.logger.Info(fmt.Sprintf("Connection from %s", conn.RemoteAddr()))
		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1024)
	mLen, err := conn.Read(buffer)
	if err != nil {
		s.logger.Error(fmt.Sprintf("failed to read: %v", err))
	}
	s.logger.Info(fmt.Sprintf("Received: %s", buffer[:mLen]))
	if _, err = conn.Write([]byte("OK")); err != nil {
		s.logger.Error(fmt.Sprintf("failed to write: %v", err))
	}
	s.logger.Info("Connection closed")
}
