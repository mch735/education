package usecase

import "github.com/mch735/education/work5/internal/entity"

type User interface {
	GetUsers() ([]*entity.User, error)
	GetUserByID(id string) (*entity.User, error)
	Create(user *entity.User) error
	Update(user *entity.User) error
	Delete(id string) error
}
