package logger

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/mch735/education/work3/internal/config"
)

func NewLogger(conf config.LoggerConfig) (*slog.Logger, error) {
	if err := conf.Validate(); err != nil {
		return nil, fmt.Errorf("invalid logger settings: %w", err)
	}

	return logger(conf), nil
}

func logger(conf config.LoggerConfig) *slog.Logger {
	level := loggerLevel(conf.Level)

	if conf.Format == config.LogFormatJSON {
		return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level}))
	}

	return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level}))
}

func loggerLevel(level config.LogLevel) slog.Level {
	return map[config.LogLevel]slog.Level{
		config.LogLevelDebug: slog.LevelDebug,
		config.LogLevelInfo:  slog.LevelInfo,
		config.LogLevelWarn:  slog.LevelWarn,
		config.LogLevelError: slog.LevelError,
	}[level]
}
