package usecase

import (
	"context"
	"fmt"
	"os"
	"shaar/domain"
	"shaar/internal/tokenutil"
	"time"
)

type signinUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewSigninUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.SigninUsecase {
	return &signinUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (su *signinUsecase) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, su.contextTimeout)
	defer cancel()
	existingUser, err := su.userRepository.GetByEmail(ctx, email)
	if err != nil {
		if os.IsTimeout(err) {
			return domain.User{}, fmt.Errorf("request timed out")
		}
		return domain.User{}, fmt.Errorf("error checking existing user: %v", err)
	}
	return existingUser, nil
}

func (su *signinUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (su *signinUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}
