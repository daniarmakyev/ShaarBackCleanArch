package usecase

import (
	"context"
	"fmt"
	"os"
	"shaar/domain"
	"time"
)

type eventUsecase struct {
	eventRepository domain.EventRepository
	contextTimeout  time.Duration
}

func NewEventUsecase(er domain.EventRepository, timeout time.Duration) domain.EventUsecase {
	return &eventUsecase{
		eventRepository: er,
		contextTimeout:  timeout,
	}
}

func (eu *eventUsecase) Create(ctx context.Context, event *domain.EventRequest) error {
	ctx, cancel := context.WithTimeout(ctx, eu.contextTimeout)
	defer cancel()
	return eu.eventRepository.Create(ctx, event)
}

func (eu *eventUsecase) GetAllEvents(ctx context.Context, page int, limit int) ([]domain.Event, error) {
	ctx, cancel := context.WithTimeout(ctx, eu.contextTimeout)
	defer cancel()
	events, err := eu.eventRepository.GetAllEvents(ctx, page, limit)
	if err != nil {
		if os.IsTimeout(err) {
			return nil, fmt.Errorf("request timed out")
		}
		return nil, err
	}
	return events, nil
}
