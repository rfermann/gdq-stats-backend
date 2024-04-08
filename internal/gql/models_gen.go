// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gql

import (
	"github.com/rfermann/gdq-stats-backend/internal/data"
)

type CreateEventInput struct {
	ScheduleID  int64  `json:"scheduleId"`
	EventTypeID string `json:"eventTypeId"`
}

type CreateEventTypeInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type DeleteEventTypeInput struct {
	ID string `json:"id"`
}

type EventDataInput struct {
	Name string `json:"name"`
	Year int64  `json:"year"`
}

type EventDataResponse struct {
	EventDataType data.EventDataType `json:"eventDataType"`
	EventData     []*data.EventDatum `json:"eventData"`
}

type GetEventDataInput struct {
	EventDataType data.EventDataType `json:"eventDataType"`
	Event         *EventDataInput    `json:"event,omitempty"`
}

type GetGamesInput struct {
	Name string `json:"name"`
	Year int64  `json:"year"`
}

type MigrateEventDataInput struct {
	ID string `json:"id"`
}

type MigrateGamesInput struct {
	ScheduleID int64 `json:"scheduleId"`
}

type Mutation struct {
}

type Query struct {
}

type UpdateEventTypeInput struct {
	ID          string  `json:"id"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}
