package usercache

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/mch735/education/work5/config"
	"github.com/mch735/education/work5/internal/entity"
)

type UserCache struct {
	db *redis.Client
}

func New(conf *config.Redis) (*UserCache, error) {
	opts, err := redis.ParseURL(conf.ToDSN())
	if err != nil {
		return nil, fmt.Errorf("redis config error: %w", err)
	}

	db := redis.NewClient(opts)

	return &UserCache{db}, nil
}

func (uc *UserCache) Set(ctx context.Context, user *entity.User, expr time.Duration) error {
	err := uc.db.Set(ctx, user.ID, user, expr).Err()
	if err != nil && !errors.Is(err, redis.Nil) {
		return fmt.Errorf("redis save user error: %w", err)
	}

	return nil
}

func (uc *UserCache) Get(ctx context.Context, id string) (*entity.User, error) {
	user := entity.User{}

	err := uc.db.Get(ctx, id).Scan(&user)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}

		return nil, fmt.Errorf("redis find user error: %w", err)
	}

	return &user, nil
}

func (uc *UserCache) Del(ctx context.Context, id string) error {
	err := uc.db.Del(ctx, id).Err()
	if err != nil && !errors.Is(err, redis.Nil) {
		return fmt.Errorf("redis delete user error: %w", err)
	}

	return nil
}

func (uc *UserCache) Close() error {
	err := uc.db.Close()
	if err != nil {
		return fmt.Errorf("redis close error: %w", err)
	}

	return nil
}
