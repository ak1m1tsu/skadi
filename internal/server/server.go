package server

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/romankravchuk/skadi/internal/log"
	"github.com/romankravchuk/skadi/internal/skadi"
	"github.com/romankravchuk/skadi/internal/storage/redis"
)

type Server struct {
	connType   string
	listenAddr string
	storage    redis.Storage
	logger     log.Logger
	clients    []net.Conn
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
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	listener, err := net.Listen(s.connType, s.listenAddr)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}
	defer listener.Close()

	go func(listener net.Listener) {
		<-sigChan
		s.logger.Info("Shutting down...")
		for _, client := range s.clients {
			client.Close()
		}
		if err := listener.Close(); err != nil {
			s.logger.Error(fmt.Sprintf("failed to close: %v", err))
		}
	}(listener)

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
	for {
		var req skadi.Request
		if err := s.read(conn, &req); err != nil {
			s.logger.Error(fmt.Sprintf("failed to read data: %v", err))
			return
		}
		s.logger.Info(fmt.Sprintf("Received: %v", req))

		var resp skadi.Response
		switch req.CommandCode {
		case skadi.Message:
			resp.Body = []byte("Hello!")
		case skadi.Close:
			resp.Body = []byte("Bye!")
			defer func() {
				conn.Close()
				s.logger.Info("Connection closed")
			}()
		default:
			resp.Body = []byte("Unknown command")
		}

		if err := s.write(conn, &resp); err != nil {
			s.logger.Error(fmt.Sprintf("failed to write data: %v", err))
			return
		}

		if req.CommandCode == skadi.Close {
			return
		}
	}
}

func (s *Server) read(conn net.Conn, data interface{}) error {
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		return err
	}
	return skadi.Decode(buffer, data)
}

func (s *Server) write(conn net.Conn, data interface{}) error {
	body, err := skadi.Encode(data)
	if err != nil {
		return err
	}
	if _, err := conn.Write(body); err != nil {
		return err
	}
	return nil
}
