package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/rfermann/gdq-stats-backend/internal/gql"
	"github.com/rfermann/gdq-stats-backend/internal/models"
)

type EventsService struct {
	models *models.Models
}

func (e *EventsService) GetCurrentEvent() (*models.Event, error) {
	event, err := e.models.Events.GetActive()
	if err != nil {
		return nil, ErrRecordNotFound
	}

	return event, nil
}

func (e *EventsService) GetEvents() ([]*models.Event, error) {
	events, err := e.models.Events.GetAll()
	if err != nil {
		return nil, ErrRecordNotFound
	}

	return events, nil
}

func (e *EventsService) GetAlternativeEvents() ([]*models.Event, error) {
	events, err := e.models.Events.GetInactive()
	if err != nil {
		return nil, ErrRecordNotFound
	}

	return events, nil
}

type event struct {
	Name     string    `json:"name"`
	Datetime time.Time `json:"datetime"`
}
type ScheduleResponse struct {
	Event event
}

func (e *EventsService) CreateEvent(input gql.CreateEventInput) (*models.Event, error) {
	event, _ := e.models.Events.GetByScheduleId(input.ScheduleID)
	if event.ID != "" {
		return nil, ErrorEntryAlreadyExists
	}

	eventType, err := e.models.EventTypes.GetByID(input.EventTypeID)
	if err != nil {
		return nil, ErrUnprocessableEntity
	}

	r, err := http.Get(fmt.Sprintf("https://gdq-site.vercel.app/api/schedule/%d", input.ScheduleID))
	if err != nil {
		return nil, ErrUnprocessableEntity
	}

	var scheduleResponse ScheduleResponse
	dec := json.NewDecoder(r.Body)
	err = dec.Decode(&scheduleResponse)
	if err != nil {
		return nil, ErrUnprocessableEntity
	}

	eventInput := &models.Event{
		Year:        int64(scheduleResponse.Event.Datetime.UTC().Year()),
		StartDate:   scheduleResponse.Event.Datetime.UTC(),
		ScheduleID:  input.ScheduleID,
		EventTypeID: eventType.ID,
	}
	return e.models.Events.Insert(eventInput)
}
