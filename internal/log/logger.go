package log

import (
	"io"

	"github.com/rs/zerolog"
)

type Logger interface {
	Info(msg string)
	Error(msg string)
}

type Config struct {
	Writer io.Writer
}

type logger struct {
	zlog *zerolog.Logger
}

func New(cfg *Config) *logger {
	zlog := zerolog.New(cfg.Writer).With().Timestamp().Logger()
	return &logger{zlog: &zlog}
}

func (l *logger) Info(msg string) {
	l.zlog.Info().Msg(msg)
}

func (l *logger) Error(msg string) {
	l.zlog.Error().Msg(msg)
}
