package storages

import (
	"log"

	"github.com/mch735/education/work2/models/user"
)

type MockUserRepo struct {
	fail bool
}

func NewMockSuccessUserRepo() *MockUserRepo {
	return &MockUserRepo{false}
}

func NewMockErrorUserRepo() *MockUserRepo {
	return &MockUserRepo{true}
}

func (s MockUserRepo) Save(user user.User) error {
	log.Printf("Save user: %v\n", user)

	if s.fail {
		return ErrUserExist
	}

	return nil
}

func (s MockUserRepo) FindByID(id string) (user.User, error) {
	log.Printf("Find user by id: %s\n", id)

	if s.fail {
		return user.User{}, ErrUserNotFound
	}

	return user.User{}, nil //nolint:exhaustruct
}

func (s MockUserRepo) DeleteByID(id string) error {
	log.Printf("Delete user by id: %s\n", id)

	if s.fail {
		return ErrUserNotFound
	}

	return nil
}

func (s MockUserRepo) FindAll() []user.User {
	log.Printf("Find all users\n")

	return []user.User{}
}

func (s MockUserRepo) FilterFunc(_ func(user user.User) bool) []user.User {
	log.Printf("Filter users by func\n")

	return []user.User{}
}
