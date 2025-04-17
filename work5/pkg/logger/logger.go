package logger

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/mch735/education/work5/config"
)

func New(conf *config.Log) (*slog.Logger, error) {
	var level slog.Level

	err := level.UnmarshalText([]byte(conf.Level))
	if err != nil {
		return nil, fmt.Errorf("l.UnmarshalText: %w", err)
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: conf.AddSource, Level: level}))
	slog.SetDefault(logger)

	return logger, nil
}
