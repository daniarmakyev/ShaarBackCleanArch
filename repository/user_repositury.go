package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"shaar/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) Create(ctx context.Context, user *domain.User) error {
	query := "INSERT INTO users (username, email, password, phone, payment, name, surname) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	_, err := ur.db.ExecContext(ctx, query, user.Username, user.Email, user.Password, user.Phone, user.Payment, user.Name, user.Surname)
	return err
}

func (ur *UserRepository) getUser(ctx context.Context, condition string, arg interface{}) (*domain.User, error) {
	query := fmt.Sprintf("SELECT id, username, email, password, phone, payment, name, surname FROM users WHERE %s", condition)
	row := ur.db.QueryRowContext(ctx, query, arg)

	var user domain.User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Phone, &user.Payment, &user.Name, &user.Surname)
	if err == sql.ErrNoRows {
		return nil, sql.ErrNoRows
	}
	if err != nil {
		return nil, fmt.Errorf("failed to scan user: %w", err)
	}

	return &user, nil
}

func (ur *UserRepository) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	user, err := ur.getUser(ctx, "email = $1", email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, nil
		}
		return domain.User{}, err
	}
	return *user, nil
}

func (ur *UserRepository) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	return ur.getUser(ctx, "id = $1", id)
}

func (ur *UserRepository) Update(ctx context.Context, user *domain.User) error {
	query := `UPDATE users 
	          SET username = $1, email = $2, phone = $3, payment = $4, name = $5, surname = $6 
	          WHERE id = $7`
	_, err := ur.db.ExecContext(ctx, query, user.Username, user.Email, user.Phone, user.Payment, user.Name, user.Surname, user.ID)
	if err != nil {
		return fmt.Errorf("userRepository.Update: %w", err)
	}
	return nil
}
