package user

import (
	"errors"
	"fmt"
	"net/mail"
	"slices"
	"strings"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidRole  = errors.New("invalid role")
	ErrInvalidName  = errors.New("invalid name")
	ErrInvalidEmail = errors.New("invalid email")
)

type Repository interface {
	Save(user *User) error
	FindByID(id string) (*User, error)
	FindAll() []*User
	DeleteByID(id string) error
	FilterFunc(fun func(user *User) bool) []*User
}

type Service struct {
	storage Repository
}

func NewService(repo Repository) *Service {
	return &Service{storage: repo}
}

func (s *Service) CreateUser(name, email, role string) (*User, error) {
	record := &User{
		ID:        uuid.NewString(),
		Name:      name,
		Email:     email,
		Role:      role,
		CreatedAt: time.Now(),
	}

	err := s.validate(record)
	if err != nil {
		return nil, fmt.Errorf("user not valid: %w", err)
	}

	err = s.storage.Save(record)
	if err != nil {
		return nil, fmt.Errorf("user not created: %w", err)
	}

	return record, nil
}

func (s *Service) GetUser(id string) (*User, error) {
	record, err := s.storage.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	return record, nil
}

func (s *Service) RemoveUser(id string) error {
	err := s.storage.DeleteByID(id)
	if err != nil {
		return fmt.Errorf("user not removed: %w", err)
	}

	return nil
}

func (s *Service) ListUsers() []*User {
	return s.storage.FindAll()
}

func (s *Service) ListUsersWithRole(role string) []*User {
	return s.storage.FilterFunc(func(user *User) bool {
		return user.Role == role
	})
}

func (s *Service) validate(user *User) error {
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
