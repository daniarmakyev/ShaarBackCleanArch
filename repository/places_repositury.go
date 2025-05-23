package repository

import (
	"context"
	"database/sql"
	"shaar/domain"
)

type PlacesRepository struct {
	db *sql.DB
}

func NewPlacesRepository(db *sql.DB) *PlacesRepository {
	return &PlacesRepository{db: db}
}

func (r *PlacesRepository) GetAllPlaces(ctx context.Context) ([]domain.Place, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, name, category, latitude, longitude, rating, price, image_url FROM places")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var places []domain.Place
	for rows.Next() {
		var place domain.Place
		if err := rows.Scan(&place.ID, &place.Name, &place.Category, &place.Latitude, &place.Longitude, &place.Rating, &place.Price, &place.ImageURL); err != nil {
			return nil, err
		}
		places = append(places, place)
	}
	return places, nil
}

func (r *PlacesRepository) GetPlacesByCategory(ctx context.Context, category string) ([]domain.Place, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, name, category, latitude, longitude, rating, price, image_url FROM places WHERE category = $1", category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var places []domain.Place
	for rows.Next() {
		var place domain.Place
		if err := rows.Scan(&place.ID, &place.Name, &place.Category, &place.Latitude, &place.Longitude, &place.Rating, &place.Price, &place.ImageURL); err != nil {
			return nil, err
		}
		places = append(places, place)
	}
	return places, nil
}

func (r *PlacesRepository) GetPlacesByPrice(ctx context.Context, price int) ([]domain.Place, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, name, category, latitude, longitude, rating, price, image_url FROM places WHERE price = $1", price)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var places []domain.Place
	for rows.Next() {
		var place domain.Place
		if err := rows.Scan(&place.ID, &place.Name, &place.Category, &place.Latitude, &place.Longitude, &place.Rating, &place.Price, &place.ImageURL); err != nil {
			return nil, err
		}
		places = append(places, place)
	}
	return places, nil
}

func (r *PlacesRepository) GetAllCategories(ctx context.Context) ([]string, error) {
	var categories []string
	rows, err := r.db.QueryContext(ctx, "SELECT DISTINCT category FROM places")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var category string
		if err := rows.Scan(&category); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}
