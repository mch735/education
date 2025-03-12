package services

import (
	"errors"
	"fmt"
	"net/mail"
	"slices"
	"strings"
	"time"

	"github.com/mch735/education/work2/models/user"
)

var (
	ErrInvalidRole  = errors.New("invalid role")
	ErrInvalidName  = errors.New("invalid name")
	ErrInvalidEmail = errors.New("invalid email")
)

type UserService struct {
	storage user.Repository
	ids     int
}

func NewUserService(repo user.Repository) *UserService {
	return &UserService{storage: repo, ids: 0}
}

func (us *UserService) CreateUser(name, email, role string) (user.User, error) {
	record := user.User{
		ID:        us.id(),
		Name:      name,
		Email:     email,
		Role:      role,
		CreatedAt: time.Now(),
	}

	err := us.validate(record)
	if err != nil {
		return user.User{}, fmt.Errorf("user not valid: %w", err)
	}

	err = us.storage.Save(record)
	if err != nil {
		return user.User{}, fmt.Errorf("user not created: %w", err)
	}

	return record, nil
}

func (us *UserService) GetUser(id int) (user.User, error) {
	record, err := us.storage.FindByID(id)
	if err != nil {
		return user.User{}, fmt.Errorf("user not found: %w", err)
	}

	return record, nil
}

func (us *UserService) RemoveUser(id int) error {
	err := us.storage.DeleteByID(id)
	if err != nil {
		return fmt.Errorf("user not removed: %w", err)
	}

	return nil
}

func (us *UserService) ListUsers() []user.User {
	return us.storage.FindAll()
}

func (us *UserService) ListUsersWithRole(role string) []user.User {
	return us.storage.FilterFunc(func(user user.User) bool {
		return user.Role == role
	})
}

func (us *UserService) id() int {
	us.ids++

	return us.ids
}

func (us *UserService) validate(user user.User) error {
	roles := []string{"admin", "user", "guest"}

	exist := slices.Contains(roles, user.Role)
	if !exist {
		return ErrInvalidRole
	}

	text := strings.TrimSpace(user.Name)
	if text == "" {
		return ErrInvalidName
	}

	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		return ErrInvalidEmail
	}

	return nil
}
