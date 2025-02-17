package domain

import (
	"context"
)

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Payment  string `json:"payment"`
	Name     string `json:"firstName"`
	Surname  string `json:"lastName"`
}

type UserUpdateRequest struct {
	Username string `json:"username" binding:"required,min=3,max=20"`
	Email    string `json:"email" binding:"required,email"`
	Phone    string `json:"phone" binding:"e164"`
	Payment  string `json:"payment"`
	Name     string `json:"firstName" binding:"alpha"`
	Surname  string `json:"lastName" binding:"alpha"`
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetByEmail(ctx context.Context, email string) (User, error)
	GetByID(ctx context.Context, id int64) (*User, error)
	Update(ctx context.Context, user *User) error
}

type UserUseCase interface {
	UpdateUser(ctx context.Context, id int64, updates UserUpdateRequest) (*User, error)
	GetUserByID(ctx context.Context, id int64) (*User, error)
	ExtractIDFromToken(requestToken string, secret string) (int64, error)
}
