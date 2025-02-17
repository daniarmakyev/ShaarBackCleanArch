package domain

import "context"

type SignupRequest struct {
	Email    string `form:"email" binding:"required,email"`
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
	Phone    string `form:"phone"`
	Payment  string `form:"payment"`
	Name     string `form:"name"`
	Surname  string `form:"surname"`
}

type SignupUsecase interface {
	Create(c context.Context, user *User) error
	GetUserByEmail(c context.Context, email string) (bool, error)
}
