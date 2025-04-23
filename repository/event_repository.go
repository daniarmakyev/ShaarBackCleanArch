package repository

import (
	"context"
	"database/sql"
	"fmt"
	"shaar/domain"
)

type EventRepository struct {
	db *sql.DB
}

func NewEventRepository(db *sql.DB) *EventRepository {
	return &EventRepository{db: db}
}

func (er *EventRepository) Create(ctx context.Context, event *domain.EventRequest) error {
	query := "INSERT INTO events (name, description, date, address) VALUES ($1, $2, $3, $4)"
	_, err := er.db.ExecContext(ctx, query, event.Name, event.Description, event.Date, event.Address)
	return err
}

func (er *EventRepository) GetAllEvents(ctx context.Context, page, limit int) ([]domain.Event, int, error) {
	offset := (page - 1) * limit
	query := fmt.Sprintf("SELECT id, name, description, TO_CHAR(date, 'YYYY-MM-DD'), address FROM events ORDER BY date ASC LIMIT %d OFFSET %d", limit, offset)
	rows, err := er.db.QueryContext(ctx, query)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var events []domain.Event
	for rows.Next() {
		var event domain.Event
		if err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Date, &event.Address); err != nil {
			return nil, 0, err
		}
		events = append(events, event)
	}

	var total int
	err = er.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM events").Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	return events, total, nil
}
