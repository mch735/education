package userrepo

import (
	"context"
	"fmt"
	"time"

	"github.com/mch735/education/work5/config"
	"github.com/mch735/education/work5/internal/entity"
)

type UserRepo struct {
	pool *Pool
}

func New(conf *config.PG) (*UserRepo, error) {
	pool, err := NewPool(conf)
	if err != nil {
		return nil, fmt.Errorf("user repo build error: %w", err)
	}

	return &UserRepo{pool: pool}, nil
}

func (ur *UserRepo) GetUsers() ([]*entity.User, error) {
	sql, args, err := StatementBuilder.
		Select("id, name, email, role, created_at, updated_at").
		From("users").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("user repo get users sql build error: %w", err)
	}

	rows, err := ur.pool.Query(context.Background(), sql, args...)
	if err != nil {
		return nil, fmt.Errorf("user repo get users query error: %w", err)
	}
	defer rows.Close()

	users := make([]*entity.User, 0, 100) //nolint:mnd

	for rows.Next() {
		user := &entity.User{}

		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("user repo get users scan error: %w", err)
		}

		users = append(users, user)
	}

	return users, nil
}

func (ur *UserRepo) GetUserByID(id string) (*entity.User, error) {
	sql, args, err := StatementBuilder.
		Select("id, name, email, role, created_at, updated_at").
		From("users").
		Where("id = ?", id).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("user repo get user by id sql build error: %w", err)
	}

	user := &entity.User{}

	err = ur.pool.QueryRow(context.Background(), sql, args...).
		Scan(&user.ID, &user.Name, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("user repo get user by id scan error: %w", err)
	}

	return user, nil
}

func (ur *UserRepo) Create(user *entity.User) error {
	sql, args, err := StatementBuilder.
		Insert("users").
		Columns("name", "email", "role").
		Values(user.Name, user.Email, user.Role).
		Suffix("RETURNING id, created_at, updated_at").
		ToSql()
	if err != nil {
		return fmt.Errorf("user repo create user sql build error: %w", err)
	}

	err = ur.pool.QueryRow(context.Background(), sql, args...).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return fmt.Errorf("user repo create user scan error: %w", err)
	}

	return nil
}

func (ur *UserRepo) Update(user *entity.User) error {
	sql, args, err := StatementBuilder.
		Update("users").
		Set("name", user.Name).
		Set("email", user.Email).
		Set("role", user.Role).
		Set("updated_at", time.Now()).
		Where("id = ?", user.ID).
		Suffix("RETURNING updated_at").
		ToSql()
	if err != nil {
		return fmt.Errorf("user repo update user sql build error: %w", err)
	}

	err = ur.pool.QueryRow(context.Background(), sql, args...).Scan(&user.UpdatedAt)
	if err != nil {
		return fmt.Errorf("user repo update user scan error: %w", err)
	}

	return nil
}

func (ur *UserRepo) Delete(id string) error {
	sql, args, err := StatementBuilder.
		Delete("users").
		Where("id = ?", id).
		ToSql()
	if err != nil {
		return fmt.Errorf("user repo delete user sql build error: %w", err)
	}

	_, err = ur.pool.Exec(context.Background(), sql, args...)
	if err != nil {
		return fmt.Errorf("user repo update user scan error: %w", err)
	}

	return nil
}

func (ur *UserRepo) Close() {
	ur.pool.Close()
}
