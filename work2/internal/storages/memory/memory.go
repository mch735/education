package memory

import (
	"github.com/mch735/education/work2/internal/storages"
	"github.com/mch735/education/work2/internal/user"
)

type UserRepo struct {
	data map[string]*user.User
}

func NewUserRepo() *UserRepo {
	return &UserRepo{
		data: make(map[string]*user.User),
	}
}

func (s *UserRepo) Save(user *user.User) error {
	_, exist := s.data[user.ID]
	if exist {
		return storages.ErrUserExist
	}

	s.data[user.ID] = user

	return nil
}

func (s *UserRepo) FindByID(id string) (*user.User, error) {
	result, exist := s.data[id]
	if !exist {
		return nil, storages.ErrUserNotFound
	}

	return result, nil
}

func (s *UserRepo) DeleteByID(id string) error {
	_, exist := s.data[id]
	if !exist {
		return storages.ErrUserNotFound
	}

	delete(s.data, id)

	return nil
}

func (s *UserRepo) FindAll() []*user.User {
	result := make([]*user.User, 0, s.Len())

	for _, v := range s.data {
		result = append(result, v)
	}

	return result
}

func (s *UserRepo) FilterFunc(fn func(user *user.User) bool) []*user.User {
	result := make([]*user.User, 0, s.Len())

	for _, v := range s.data {
		if fn(v) {
			result = append(result, v)
		}
	}

	return result
}

func (s *UserRepo) Len() int {
	return len(s.data)
}
