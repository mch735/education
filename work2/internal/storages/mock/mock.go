package mock

import (
	"log"

	"github.com/mch735/education/work2/internal/storages"
	"github.com/mch735/education/work2/internal/user"
)

type UserRepo struct {
	fail bool
}

func NewSuccessUserRepo() *UserRepo {
	return &UserRepo{false}
}

func NewErrorUserRepo() *UserRepo {
	return &UserRepo{true}
}

func (s *UserRepo) Save(user *user.User) error {
	log.Printf("Save user: %v\n", user)

	if s.fail {
		return storages.ErrUserExist
	}

	return nil
}

func (s *UserRepo) FindByID(id string) (*user.User, error) {
	log.Printf("Find user by id: %s\n", id)

	if s.fail {
		return nil, storages.ErrUserNotFound
	}

	return &user.User{}, nil //nolint:exhaustruct
}

func (s *UserRepo) DeleteByID(id string) error {
	log.Printf("Delete user by id: %s\n", id)

	if s.fail {
		return storages.ErrUserNotFound
	}

	return nil
}

func (s *UserRepo) FindAll() []*user.User {
	log.Printf("Find all users\n")

	return []*user.User{}
}

func (s *UserRepo) FilterFunc(_ func(user *user.User) bool) []*user.User {
	log.Printf("Filter users by func\n")

	return []*user.User{}
}

func (s *UserRepo) Len() int {
	return 0
}
