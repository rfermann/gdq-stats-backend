package services

import (
	"errors"
	"github.com/rfermann/gdq-stats-backend/internal/data"
)

var (
	ErrRecordNotFound      = errors.New("record not found")
	ErrUnprocessableEntity = errors.New("unprocessable entity")
)

type Services struct {
	EventService *EventService
}

func New(models *data.Models) *Services {
	return &Services{
		EventService: &EventService{models: models},
	}
}
