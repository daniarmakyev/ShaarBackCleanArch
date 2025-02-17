package usecase

import (
	"context"
	"errors"
	"shaar/domain"
)

type UserUsecase struct {
	UserRepo domain.UserRepository
}

func NewUserUsecase(userRepo domain.UserRepository) domain.UserUseCase {
	return &UserUsecase{
		UserRepo: userRepo,
	}
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
