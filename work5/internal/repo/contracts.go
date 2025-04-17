package repo

import (
	"context"
	"time"

	"github.com/mch735/education/work5/internal/entity"
)

type (
	UserRepo interface {
		GetUsers(ctx context.Context) ([]*entity.User, error)
		GetUserByID(ctx context.Context, id string) (*entity.User, error)
		Create(ctx context.Context, user *entity.User) error
		Update(ctx context.Context, user *entity.User) error
		Delete(ctx context.Context, id string) error
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
