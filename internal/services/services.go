package services

import (
	"errors"
	"github.com/rfermann/gdq-stats-backend/internal/models"
)

var (
	ErrRecordNotFound       = errors.New("record not found")
	ErrUnprocessableEntity  = errors.New("unprocessable entity")
	ErrorEntryAlreadyExists = errors.New("entry already exists")
)

type Services struct {
	EventDataService  *EventDataService
	EventsService     *EventsService
	EventTypesService *EventTypesService
	GamesService      *GamesService
}

func New(models *models.Models) *Services {
	return &Services{
		EventDataService:  &EventDataService{models: models},
		EventsService:     &EventsService{models: models},
		EventTypesService: &EventTypesService{models: models},
		GamesService:      &GamesService{models: models},
	}
}
