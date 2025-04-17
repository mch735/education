package logger

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/mch735/education/work5/config"
)

func New(conf *config.Log) (*slog.Logger, error) {
	var l slog.Level

	err := l.UnmarshalText([]byte(conf.Level))
	if err != nil {
		return nil, fmt.Errorf("logger error: %w", err)
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: l}))
	slog.SetDefault(logger)

	return logger, nil
}
