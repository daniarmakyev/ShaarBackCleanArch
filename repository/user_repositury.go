package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"shaar/domain"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) Create(ctx context.Context, user *domain.User) error {
	query := "INSERT INTO users (username, email, password, avatar) VALUES ($1, $2, $3, $4)"
	_, err := ur.db.ExecContext(ctx, query, user.Username, user.Email, user.Password, user.Avatar)
	return err
}

func (ur *userRepository) getUser(ctx context.Context, condition string, arg interface{}) (*domain.User, error) {
	query := fmt.Sprintf("SELECT id, username, email, password, avatar FROM users WHERE %s", condition)
	row := ur.db.QueryRowContext(ctx, query, arg)

	var user domain.User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Avatar)
	if err == sql.ErrNoRows {
		return nil, sql.ErrNoRows
	}
	if err != nil {
		return nil, fmt.Errorf("failed to scan user: %w", err)
	}

	return &user, nil
}

func (ur *userRepository) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	user, err := ur.getUser(ctx, "email = $1", email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, nil
		}
		return domain.User{}, err
	}
	return *user, nil
}

func (ur *userRepository) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	return ur.getUser(ctx, "id = $1", id)
}
