package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/mch735/education/work5/internal/entity"
	"github.com/mch735/education/work5/internal/repo"
)

type UseCase struct {
	repo  repo.UserRepo
	cache repo.UserCache
	ms    repo.MessageSys
}

func New(repo repo.UserRepo, cache repo.UserCache, ms repo.MessageSys) *UseCase {
	return &UseCase{repo, cache, ms}
}

func (uc *UseCase) GetUsers(ctx context.Context) ([]*entity.User, error) {
	err := uc.ms.Publish("methods", []byte("usecase.GetUsers"))
	if err != nil {
		return nil, fmt.Errorf("uc.ms.Publish: %w", err)
	}

	users, err := uc.repo.GetUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("uc.repo.GetUsers: %w", err)
	}

	return users, nil
}

func (uc *UseCase) GetUserByID(ctx context.Context, id string) (*entity.User, error) {
	err := uc.ms.Publish("methods", []byte("usecase.GetUserByID"))
	if err != nil {
		return nil, fmt.Errorf("uc.ms.Publish: %w", err)
	}

	user, err := uc.cache.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("uc.cache.Get: %w", err)
	}

	if user != nil {
		return user, nil
	}

	user, err = uc.repo.GetUserByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("uc.repo.GetUserByID: %w", err)
	}

	err = uc.cache.Set(ctx, user, time.Minute)
	if err != nil {
		return nil, fmt.Errorf("uc.cache.Set: %w", err)
	}

	return user, nil
}

func (uc *UseCase) Create(ctx context.Context, user *entity.User) error {
	err := uc.ms.Publish("methods", []byte("usecase.Create"))
	if err != nil {
		return fmt.Errorf("uc.ms.Publish: %w", err)
	}

	err = uc.repo.Create(ctx, user)
	if err != nil {
		return fmt.Errorf("uc.repo.Create: %w", err)
	}

	err = uc.cache.Set(ctx, user, time.Minute)
	if err != nil {
		return fmt.Errorf("uc.cache.Set: %w", err)
	}

	return nil
}

func (uc *UseCase) Update(ctx context.Context, user *entity.User) error {
	err := uc.ms.Publish("methods", []byte("usecase.Update"))
	if err != nil {
		return fmt.Errorf("uc.ms.Publish: %w", err)
	}

	err = uc.repo.Update(ctx, user)
	if err != nil {
		return fmt.Errorf("uc.repo.Update: %w", err)
	}

	err = uc.cache.Set(ctx, user, time.Minute)
	if err != nil {
		return fmt.Errorf("uc.cache.Set: %w", err)
	}

	return nil
}

func (uc *UseCase) Delete(ctx context.Context, id string) error {
	err := uc.ms.Publish("methods", []byte("usecase.Delete"))
	if err != nil {
		return fmt.Errorf("uc.ms.Publish: %w", err)
	}

	err = uc.cache.Del(ctx, id)
	if err != nil {
		return fmt.Errorf("uc.cache.Del: %w", err)
	}

	err = uc.repo.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("uc.repo.Delete: %w", err)
	}

	return nil
}
