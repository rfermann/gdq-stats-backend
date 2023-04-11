package services

import (
	"context"
	"database/sql"

	db_models "github.com/rfermann/gdq-stats-backend/internal/db/models"
	"github.com/rfermann/gdq-stats-backend/internal/errors"
)

type EventService struct {
	db *sql.DB
}

func (e *EventService) GetCurrentEvent() (*db_models.Event, error) {
	event, err := db_models.Events(db_models.EventWhere.ActiveEvent.EQ(true)).One(context.Background(), e.db)
	if err != nil {
		return nil, errors.ErrRecordNotFound
	}

	return event, nil
}
