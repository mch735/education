package storages

import (
	"log"

	"github.com/mch735/education/work2/models/user"
)

type MockUserRepo struct{}

func NewMockUserRepo() MockUserRepo {
	return MockUserRepo{}
}

func (s MockUserRepo) Save(user user.User) error {
	log.Printf("Save user: %v\n", user)

	return nil
}

func (s MockUserRepo) FindByID(id int) (user.User, error) {
	log.Printf("Find user by id: %d\n", id)

	return user.User{}, nil //nolint:exhaustruct
}

func (s MockUserRepo) DeleteByID(id int) error {
	log.Printf("Delete user by id: %d\n", id)

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
