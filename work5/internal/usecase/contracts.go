package usecase

import (
	"context"

	"github.com/mch735/education/work5/internal/entity"
)

type User interface {
	GetUsers(ctx context.Context) ([]*entity.User, error)
	GetUserByID(ctx context.Context, id string) (*entity.User, error)
	Create(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id string) error
}
