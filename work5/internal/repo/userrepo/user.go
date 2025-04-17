package userrepo

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/mch735/education/work5/config"
	"github.com/mch735/education/work5/internal/entity"
)

type UserRepo struct {
	pool *pgxpool.Pool
}

func New(conf *config.PG) (*UserRepo, error) {
	pool, err := pgxpool.New(context.Background(), conf.ToDSN())
	if err != nil {
		return nil, fmt.Errorf("pgxpool.New: %w", err)
	}

	return &UserRepo{pool: pool}, nil
}

func (ur *UserRepo) GetUsers(ctx context.Context) ([]*entity.User, error) {
	sql, args, err := StatementBuilder.
		Select("id, name, email, role, created_at, updated_at").
		From("users").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("StatementBuilder: %w", err)
	}

	rows, err := ur.pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("ur.pool.Query: %w", err)
	}
	defer rows.Close()

	users := make([]*entity.User, 0, 100) //nolint:mnd

	for rows.Next() {
		user := &entity.User{}

		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan: %w", err)
		}

		users = append(users, user)
	}

	return users, nil
}

func (ur *UserRepo) GetUserByID(ctx context.Context, id string) (*entity.User, error) {
	sql, args, err := StatementBuilder.
		Select("id, name, email, role, created_at, updated_at").
		From("users").
		Where("id = ?", id).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("StatementBuilder: %w", err)
	}

	user := &entity.User{}

	err = ur.pool.QueryRow(ctx, sql, args...).
		Scan(&user.ID, &user.Name, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("ur.pool.QueryRow: %w", err)
	}

	return user, nil
}

func (ur *UserRepo) Create(ctx context.Context, user *entity.User) error {
	sql, args, err := StatementBuilder.
		Insert("users").
		Columns("name", "email", "role").
		Values(user.Name, user.Email, user.Role).
		Suffix("RETURNING id, created_at, updated_at").
		ToSql()
	if err != nil {
		return fmt.Errorf("StatementBuilder: %w", err)
	}

	err = ur.pool.QueryRow(ctx, sql, args...).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return fmt.Errorf("ur.pool.QueryRow: %w", err)
	}

	return nil
}

func (ur *UserRepo) Update(ctx context.Context, user *entity.User) error {
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
		return fmt.Errorf("StatementBuilder: %w", err)
	}

	err = ur.pool.QueryRow(ctx, sql, args...).Scan(&user.UpdatedAt)
	if err != nil {
		return fmt.Errorf("ur.pool.QueryRow: %w", err)
	}

	return nil
}

func (ur *UserRepo) Delete(ctx context.Context, id string) error {
	sql, args, err := StatementBuilder.
		Delete("users").
		Where("id = ?", id).
		ToSql()
	if err != nil {
		return fmt.Errorf("StatementBuilder: %w", err)
	}

	_, err = ur.pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("ur.pool.Exec: %w", err)
	}

	return nil
}

func (ur *UserRepo) Close() {
	ur.pool.Close()
}
