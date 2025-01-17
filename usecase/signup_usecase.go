package usecase

import (
	"context"
	"fmt"
	"shaar/domain"
)

type signupUsecase struct {
	userRepository domain.UserRepository
}

func NewSignupUsecase(userRepository domain.UserRepository) domain.SignupUsecase {
	return &signupUsecase{userRepository: userRepository}
}
func (su *signupUsecase) Create(ctx context.Context, user *domain.User) error {
	existingUser, err := su.userRepository.GetByEmail(ctx, user.Email)
	if err != nil {
		return fmt.Errorf("error checking existing user: %v", err)
	}
	if existingUser != (domain.User{}) {
		return fmt.Errorf("user with email %s already exists", user.Email)
	}
	return su.userRepository.Create(ctx, user)
}

func (su *signupUsecase) GetUserByEmail(ctx context.Context, email string) (bool, error) {
	existingUser, err := su.userRepository.GetByEmail(ctx, email)
	if err != nil {
		return false, fmt.Errorf("error checking existing user: %v", err)
	}
	return existingUser != (domain.User{}), nil
}
