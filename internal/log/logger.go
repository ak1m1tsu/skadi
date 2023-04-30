package log

import (
	"io"
	"os"

	"github.com/rs/zerolog"
)

type Logger interface {
	Info(msg string)
	Error(msg string)
}

var DefaultConfig Config = Config{
	Writer: zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: "2006-01-02 15:04:05",
	},
}

type Config struct {
	Writer io.Writer
}

type logger struct {
	zlog *zerolog.Logger
}

func New(cfg ...Config) *logger {
	var zlog zerolog.Logger
	if cfg == nil {
		zlog = zerolog.New(DefaultConfig.Writer).With().Timestamp().Logger()
	} else if len(cfg) > 0 {
		zlog = zerolog.New(cfg[0].Writer).With().Timestamp().Logger()
	}
	return &logger{zlog: &zlog}
}

func (l *logger) Info(msg string) {
	l.zlog.Info().Msg(msg)
}

func (l *logger) Error(msg string) {
	l.zlog.Error().Msg(msg)
}
