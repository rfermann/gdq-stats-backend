package models

import (
	"github.com/jmoiron/sqlx"
)

type Models struct {
	Events     EventsModel
	EventData  EventDatumModel
	EventTypes EventTypesModel
	Games      GameModel
}

func NewModels(db *sqlx.DB) *Models {
	return &Models{
		Events:     EventsModel{db: db},
		EventData:  EventDatumModel{db: db},
		EventTypes: EventTypesModel{db: db},
		Games:      GameModel{db: db},
	}
}
