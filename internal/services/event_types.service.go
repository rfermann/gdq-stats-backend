package services

import (
	"github.com/rfermann/gdq-stats-backend/internal/gql"
	"github.com/rfermann/gdq-stats-backend/internal/models"
)

type EventTypesService struct {
	models *models.Models
}

func (e *EventTypesService) CreateEventType(input gql.CreateEventTypeInput) (*models.EventType, error) {
	eventType, err := e.models.EventTypes.Insert(models.EventType{
		Name:        input.Name,
		Description: input.Description,
	})
	if err != nil {
		return nil, ErrUnprocessableEntity
	}

	return eventType, nil
}

func (e *EventTypesService) GetEventTypeByID(id string) (*models.EventType, error) {
	eventType, err := e.models.EventTypes.GetByID(id)
	if err != nil {
		return nil, ErrRecordNotFound
	}

	return eventType, nil
}

func (e *EventTypesService) GetEventTypes() ([]*models.EventType, error) {
	eventTypes, err := e.models.EventTypes.GetAll()
	if err != nil {
		return nil, ErrRecordNotFound
	}

	return eventTypes, nil
}

func (e *EventTypesService) UpdateEventType(input gql.UpdateEventTypeInput) (*models.EventType, error) {
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

func (e *EventTypesService) DeleteEventType(input gql.DeleteEventTypeInput) (*models.EventType, error) {
	eventType, err := e.models.EventTypes.Delete(input.ID)
	if err != nil {
		return nil, ErrUnprocessableEntity
	}

	return eventType, nil
}
