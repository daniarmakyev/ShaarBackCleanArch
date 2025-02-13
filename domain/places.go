package domain

import "context"

type Place struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Category  string  `json:"category"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Rating    float64 `json:"rating"`
	Price     int     `json:"price"`
	ImageURL  string  `json:"imageUrl"`
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
