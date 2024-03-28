package services

import "github.com/jmoiron/sqlx"

type Services struct {
	EventService *EventService
}

func New(db *sqlx.DB) *Services {
	return &Services{
		EventService: &EventService{db: db},
	}
}
