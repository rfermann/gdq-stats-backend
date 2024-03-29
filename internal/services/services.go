package services

import (
	"github.com/rfermann/gdq-stats-backend/internal/data"
)

type Services struct {
	EventService *EventService
}

func New(models *data.Models) *Services {
	return &Services{
		EventService: &EventService{models: models},
	}
}
