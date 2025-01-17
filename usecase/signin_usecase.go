package usecase

import (
	"context"
	"fmt"
	"shaar/domain"
	"shaar/internal/tokenutil"
)

type signinUsecase struct {
	userRepository domain.UserRepository
}

func NewSigninUsecase(userRepository domain.UserRepository) domain.SigninUsecase {
	return &signinUsecase{userRepository: userRepository}
}

func (su *signinUsecase) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	existingUser, err := su.userRepository.GetByEmail(ctx, email)
	if err != nil {
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
