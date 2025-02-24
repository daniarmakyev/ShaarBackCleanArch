package domain

import "context"

type EventRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Address     string `json:"address"`
}

type Event struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Address     string `json:"address"`
}

type EventRepository interface {
	Create(ctx context.Context, event *EventRequest) error
	GetAllEvents(ctx context.Context, page int, linit int) ([]Event, error)
}

type EventUsecase interface {
	Create(ctx context.Context, event *EventRequest) error
	GetAllEvents(ctx context.Context, page int, limit int) ([]Event, error)
}
