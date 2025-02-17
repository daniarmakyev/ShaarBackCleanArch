package usecase

import (
	"context"
	"errors"
	"fmt"
	"os"
	"shaar/domain"
	"shaar/internal/tokenutil"
	"time"
)

type UserUsecase struct {
	UserRepo       domain.UserRepository
	contextTimeout time.Duration
}

func NewUserUsecase(userRepo domain.UserRepository, timeout time.Duration) domain.UserUseCase {
	return &UserUsecase{
		UserRepo:       userRepo,
		contextTimeout: timeout,
	}
}

func (rtu *UserUsecase) GetUserByID(ctx context.Context, id int64) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, rtu.contextTimeout)
	defer cancel()

	user, err := rtu.UserRepo.GetByID(ctx, id)
	if err != nil {
		if os.IsTimeout(err) || errors.Is(err, context.DeadlineExceeded) {
			return nil, fmt.Errorf("request timed out")
		}
		return nil, err
	}
	return user, nil
}

func (uc *UserUsecase) UpdateUser(ctx context.Context, id int64, updates domain.UserUpdateRequest) (*domain.User, error) {
	user, err := uc.UserRepo.GetByID(ctx, id)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if updates.Username != "" {
		user.Username = updates.Username
	}
	if updates.Email != "" {
		user.Email = updates.Email
	}

	if updates.Phone != "" {
		user.Phone = updates.Phone
	}
	if updates.Payment != "" {
		user.Payment = updates.Payment
	}
	if updates.Name != "" {
		user.Name = updates.Name
	}
	if updates.Surname != "" {
		user.Surname = updates.Surname
	}

	err = uc.UserRepo.Update(ctx, user)
	if err != nil {
		return nil, errors.New("error saving updated user")
	}

	return user, nil
}

func (rtu *UserUsecase) ExtractIDFromToken(requestToken string, secret string) (int64, error) {
	return tokenutil.ExtractIDFromToken(requestToken, secret)
}
