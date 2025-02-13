package usecase

import (
	"context"
	"fmt"
	"os"
	"shaar/domain"
	"time"
)

type placesUsecase struct {
	placesRepository domain.PlacesRepository
	contextTimeout   time.Duration
}

func NewPlacesUsecase(pr domain.PlacesRepository, timeout time.Duration) domain.PlacesUsecase {
	return &placesUsecase{
		placesRepository: pr,
		contextTimeout:   timeout,
	}
}

func (pu *placesUsecase) GetAllPlaces(ctx context.Context) ([]domain.Place, error) {
	ctx, cancel := context.WithTimeout(ctx, pu.contextTimeout)
	defer cancel()
	places, err := pu.placesRepository.GetAllPlaces()
	if err != nil {
		if os.IsTimeout(err) {
			return nil, fmt.Errorf("request timed out")
		}
		return nil, err
	}
	return places, nil
}

func (pu *placesUsecase) GetPlacesByCategory(ctx context.Context, category string) ([]domain.Place, error) {
	ctx, cancel := context.WithTimeout(ctx, pu.contextTimeout)
	defer cancel()
	places, err := pu.placesRepository.GetPlacesByCategory(category)
	if err != nil {
		if os.IsTimeout(err) {
			return nil, fmt.Errorf("request timed out")
		}
		return nil, err
	}
	return places, nil
}

func (pu *placesUsecase) GetPlacesByPrice(ctx context.Context, price int) ([]domain.Place, error) {
	ctx, cancel := context.WithTimeout(ctx, pu.contextTimeout)
	defer cancel()
	places, err := pu.placesRepository.GetPlacesByPrice(price)
	if err != nil {
		if os.IsTimeout(err) {
			return nil, fmt.Errorf("request timed out")
		}
		return nil, err
	}
	return places, nil
}

func (pu *placesUsecase) GetAllCategories(ctx context.Context) ([]string, error) {
	ctx, cancel := context.WithTimeout(ctx, pu.contextTimeout)
	defer cancel()
	categories, err := pu.placesRepository.GetAllCategories()
	if err != nil {
		if os.IsTimeout(err) {
			return nil, fmt.Errorf("request timed out")
		}
		return nil, err
	}
	return categories, nil
}
