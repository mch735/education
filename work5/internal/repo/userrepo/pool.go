package userrepo

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/mch735/education/work5/config"
)

const (
	attemptsCount = 5
	timeout       = time.Second
)

type Pool struct {
	*pgxpool.Pool
}

func NewPool(conf *config.PG) (*Pool, error) {
	var pool *pgxpool.Pool

	cfg, err := pgxpool.ParseConfig(conf.ToDSN())
	if err != nil {
		return nil, fmt.Errorf("pgxpool config error: %w", err)
	}

	for range attemptsCount {
		pool, err = pgxpool.NewWithConfig(context.Background(), cfg)
		if err == nil {
			break
		}

		slog.Info("pgxpool is trying to connect ...")
		time.Sleep(timeout)
	}

	if err != nil {
		return nil, fmt.Errorf("pgxpool connect fail: %w", err)
	}

	return &Pool{pool}, nil
}
