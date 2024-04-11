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

func (e *EventsService) GetAlternativeEvents(input *gql.GetAlternativeEventsInput) ([]*models.Event, error) {
	if input == nil {
		events, err := e.models.Events.GetInactive()
		if err != nil {
			return nil, ErrRecordNotFound
		}
		return events, nil
	}

	event, err := e.models.Events.GetByNameAndYear(input.Name, input.Year)
	if err != nil {
		return nil, ErrRecordNotFound
	}

	return e.models.Events.GetAlternativeEventsForEventId(event.ID)
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

func (e *EventsService) AggregateEventStatistics(input *gql.AggregateEventStatisticsInput) (*models.Event, error) {
	event, err := e.models.Events.GetById(input.ID)
	if err != nil {
		return nil, ErrRecordNotFound
	}

	var eventDataExists bool

	eventDataExists, err = e.models.EventData.CheckEventDataExistsForEventId(input.ID)
	if err != nil || !eventDataExists {
		return nil, ErrRecordNotFound
	}

	event.Viewers, err = e.models.EventData.GetViewersCountForEventId(input.ID)
	event.Donations, err = e.models.EventData.GetDonationsCountForEventId(input.ID)
	event.Donors, err = e.models.EventData.GetDonorsCountForEventId(input.ID)
	event.Tweets, err = e.models.EventData.GetTweetsCountForEventId(input.ID)
	event.TwitchChats, err = e.models.EventData.GetTwitchChatsCountForEventId(input.ID)
	event.EventDataCount, err = e.models.EventData.GetEventDataCountForEventId(input.ID)
	event.CompletedGames, err = e.models.Games.GetCompletedGamesCountForEventId(input.ID)
	event.TotalGames, err = e.models.Games.GetTotalGamesCountForEventId(input.ID)

	return e.models.Events.Update(*event)
}

func (e *EventsService) ActivateEvent(input gql.ActivateEventInput) (*models.Event, error) {
	event, err := e.models.Events.ActivateById(input.ID)
	if err != nil {
		return nil, ErrUnprocessableEntity
	}

	return event, nil
}
