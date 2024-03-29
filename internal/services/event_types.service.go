package services

import (
	"github.com/rfermann/gdq-stats-backend/internal/data"
	"github.com/rfermann/gdq-stats-backend/internal/gql"
)

type EventTypesService struct {
	models *data.Models
}

func (e *EventTypesService) CreateEventType(input gql.CreateEventTypeInput) (*data.EventType, error) {
	eventType, err := e.models.EventTypes.Insert(data.EventType{
		Name:        input.Name,
		Description: input.Description,
	})
	if err != nil {
		return nil, ErrUnprocessableEntity
	}

	return eventType, nil
}

func (e *EventTypesService) GetEventTypeByID(id string) (*data.EventType, error) {
	eventType, err := e.models.EventTypes.GetByID(id)
	if err != nil {
		return nil, ErrRecordNotFound
	}

	return eventType, nil
}

func (e *EventTypesService) GetEventTypes() ([]*data.EventType, error) {
	eventTypes, err := e.models.EventTypes.GetAll()
	if err != nil {
		return nil, ErrRecordNotFound
	}

	return eventTypes, nil
}

func (e *EventTypesService) UpdateEventType(input gql.UpdateEventTypeInput) (*data.EventType, error) {
	eventType, err := e.models.EventTypes.GetByID(input.ID)
	if err != nil {
		return nil, ErrRecordNotFound
	}

	if input.Description != nil {
		eventType.Description = *input.Description
	}

	if input.Name != nil {
		eventType.Name = *input.Name
	}

	eventType, err = e.models.EventTypes.Update(*eventType)
	if err != nil {
		return nil, ErrUnprocessableEntity
	}

	return eventType, nil
}

func (e *EventTypesService) DeleteEventType(input gql.DeleteEventTypeInput) (*data.EventType, error) {
	eventType, err := e.models.EventTypes.Delete(input.ID)
	if err != nil {
		return nil, ErrUnprocessableEntity
	}

	return eventType, nil
}
