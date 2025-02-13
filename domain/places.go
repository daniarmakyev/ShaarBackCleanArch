package domain

import "context"

type Place struct {
	ID        int
	Name      string
	Category  string
	Latitude  float64
	Longitude float64
	Rating    float64
	Price     int
	ImageURL  string
}

type PlacesRepository interface {
	GetAllPlaces() ([]Place, error)
	GetPlacesByCategory(category string) ([]Place, error)
	GetPlacesByPrice(price int) ([]Place, error)
	GetAllCategories() ([]string, error)
}

type PlacesUsecase interface {
	GetAllPlaces(ctx context.Context) ([]Place, error)
	GetPlacesByCategory(ctx context.Context, category string) ([]Place, error)
	GetPlacesByPrice(ctx context.Context, price int) ([]Place, error)
	GetAllCategories(ctx context.Context) ([]string, error)
}
