package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"shaar/domain"
	"time"
)

type signupUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewSignupUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.SignupUsecase {
	return &signupUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (su *signupUsecase) Create(ctx context.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(ctx, su.contextTimeout)
	defer cancel()
	existingUser, err := su.userRepository.GetByEmail(ctx, user.Email)

	if err != nil {
		if os.IsTimeout(err) {
			return fmt.Errorf("request timed out")
		}
		if !errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("error checking existing user: %v", err)
		}
	}

	if existingUser != (domain.User{}) {
		return fmt.Errorf("user with email %s already exists", user.Email)
	}

	return su.userRepository.Create(ctx, user)
}

func (su *signupUsecase) GetUserByEmail(ctx context.Context, email string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, su.contextTimeout)
	defer cancel()
	existingUser, err := su.userRepository.GetByEmail(ctx, email)
	if err != nil {
		if os.IsTimeout(err) {
			return false, fmt.Errorf("request timed out")
		}
		return false, fmt.Errorf("error checking existing user: %v", err)
	}
	return existingUser != (domain.User{}), nil
}
