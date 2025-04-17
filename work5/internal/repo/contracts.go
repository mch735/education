package repo

import (
	"context"
	"time"

	"github.com/mch735/education/work5/internal/entity"
)

type (
	UserRepo interface {
		GetUsers() ([]*entity.User, error)
		GetUserByID(id string) (*entity.User, error)
		Create(user *entity.User) error
		Update(user *entity.User) error
		Delete(id string) error
	}

	UserCache interface {
		Set(ctx context.Context, user *entity.User, expr time.Duration) error
		Get(ctx context.Context, id string) (*entity.User, error)
		Del(ctx context.Context, id string) error
	}

	MessageSys interface {
		Publish(subj string, data []byte) error
	}
)
