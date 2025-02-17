package usecase

import (
	"context"
	"fmt"
	"os"
	"shaar/domain"
	"shaar/internal/tokenutil"
	"time"
)

type refreshTokenUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewRefreshTokenUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.RefreshTokenUsecase {
	return &refreshTokenUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (rtu *refreshTokenUsecase) GetUserByID(ctx context.Context, id int64) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, rtu.contextTimeout)
	defer cancel()
	user, err := rtu.userRepository.GetByID(ctx, id)
	if err != nil {
		if os.IsTimeout(err) {
			return domain.User{}, fmt.Errorf("request timed out")
		}
		return domain.User{}, err
	}
	return *user, nil
}

func (rtu *refreshTokenUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (rtu *refreshTokenUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}

func (rtu *refreshTokenUsecase) ExtractIDFromToken(requestToken string, secret string) (int64, error) {
	return tokenutil.ExtractIDFromToken(requestToken, secret)
}
