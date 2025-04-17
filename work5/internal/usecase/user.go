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

func (uc *UseCase) GetUsers() ([]*entity.User, error) {
	err := uc.ms.Publish("methods", []byte("usecase.GetUsers"))
	if err != nil {
		return nil, fmt.Errorf("usecase get users publish error: %w", err)
	}

	users, err := uc.repo.GetUsers()
	if err != nil {
		return nil, fmt.Errorf("usecase get users request error: %w", err)
	}

	return users, nil
}

func (uc *UseCase) GetUserByID(id string) (*entity.User, error) {
	err := uc.ms.Publish("methods", []byte("usecase.GetUserByID"))
	if err != nil {
		return nil, fmt.Errorf("usecase get user by id publish error: %w", err)
	}

	user, err := uc.cache.Get(context.Background(), id)
	if err != nil {
		return nil, fmt.Errorf("usecase get user by id load from cache error: %w", err)
	}

	if user != nil {
		return user, nil
	}

	user, err = uc.repo.GetUserByID(id)
	if err != nil {
		return nil, fmt.Errorf("usecase get user by id in repo error: %w", err)
	}

	err = uc.cache.Set(context.Background(), user, time.Minute)
	if err != nil {
		return nil, fmt.Errorf("usecase get user by id store in cache error: %w", err)
	}

	return user, nil
}

func (uc *UseCase) Create(user *entity.User) error {
	err := uc.ms.Publish("methods", []byte("usecase.Create"))
	if err != nil {
		return fmt.Errorf("usecase create user publish error: %w", err)
	}

	err = uc.repo.Create(user)
	if err != nil {
		return fmt.Errorf("usecase create user request error: %w", err)
	}

	err = uc.cache.Set(context.Background(), user, time.Minute)
	if err != nil {
		return fmt.Errorf("usecase create user store in cache error: %w", err)
	}

	return nil
}

func (uc *UseCase) Update(user *entity.User) error {
	err := uc.ms.Publish("methods", []byte("usecase.Update"))
	if err != nil {
		return fmt.Errorf("usecase update user publish error: %w", err)
	}

	err = uc.repo.Update(user)
	if err != nil {
		return fmt.Errorf("usecase update user request error: %w", err)
	}

	err = uc.cache.Set(context.Background(), user, time.Minute)
	if err != nil {
		return fmt.Errorf("usecase update user store in cache error: %w", err)
	}

	return nil
}

func (uc *UseCase) Delete(id string) error {
	err := uc.ms.Publish("methods", []byte("usecase.Delete"))
	if err != nil {
		return fmt.Errorf("usecase delete user publish error: %w", err)
	}

	err = uc.cache.Del(context.Background(), id)
	if err != nil {
		return fmt.Errorf("usecase delete user remove in cache error: %w", err)
	}

	err = uc.repo.Delete(id)
	if err != nil {
		return fmt.Errorf("usecase delete user request error: %w", err)
	}

	return nil
}
