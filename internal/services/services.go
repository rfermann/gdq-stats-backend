package services

import "database/sql"

type Services struct {
	EventService *EventService
}

func New(db *sql.DB) *Services {
	return &Services{
		EventService: &EventService{db: db},
	}
}
