// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gql

import (
	"github.com/rfermann/gdq-stats-backend/internal/models"
)

type ActivateEventInput struct {
	ID string `json:"id"`
}

type AggregateEventStatisticsInput struct {
	ID string `json:"id"`
}

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
	EventDataType models.EventDataType `json:"eventDataType"`
	EventData     []*models.EventDatum `json:"eventData"`
}

type GetAlternativeEventsInput struct {
	Name string `json:"name"`
	Year int64  `json:"year"`
}

type GetEventDataInput struct {
	EventDataType models.EventDataType `json:"eventDataType"`
	Event         *EventDataInput      `json:"event,omitempty"`
}

type GetGamesInput struct {
	Name string `json:"name"`
	Year int64  `json:"year"`
}

type MigrateEventDataInput struct {
	EventID string `json:"eventId"`
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
