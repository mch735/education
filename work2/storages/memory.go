package storages

import "github.com/mch735/education/work2/models/user"

type InMemoryUserRepo struct {
	data map[string]user.User
}

func NewInMemoryUserRepo() InMemoryUserRepo {
	return InMemoryUserRepo{
		data: make(map[string]user.User),
	}
}

func (s InMemoryUserRepo) Save(user user.User) error {
	_, exist := s.data[user.ID]
	if exist {
		return ErrUserExist
	}

	s.data[user.ID] = user

	return nil
}

func (s InMemoryUserRepo) FindByID(id string) (user.User, error) {
	result, exist := s.data[id]
	if !exist {
		return user.User{}, ErrUserNotFound
	}

	return result, nil
}

func (s InMemoryUserRepo) DeleteByID(id string) error {
	_, exist := s.data[id]
	if !exist {
		return ErrUserNotFound
	}

	delete(s.data, id)

	return nil
}

func (s InMemoryUserRepo) FindAll() []user.User {
	result := make([]user.User, 0, s.Len())

	for _, v := range s.data {
		result = append(result, v)
	}

	return result
}

func (s InMemoryUserRepo) FilterFunc(fn func(user user.User) bool) []user.User {
	result := make([]user.User, 0, s.Len())

	for _, v := range s.data {
		if fn(v) {
			result = append(result, v)
		}
	}

	return result
}

func (s InMemoryUserRepo) Len() int {
	return len(s.data)
}
