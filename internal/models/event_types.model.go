package models

import (
	"context"
	"github.com/jmoiron/sqlx"
	"time"
)

type EventType struct {
	ID          string
	Name        string
	Description string
}

type EventTypesModel struct {
	db *sqlx.DB
}

func (m *EventTypesModel) Insert(input EventType) (*EventType, error) {
	stmt := `
		INSERT INTO event_types (name, description)
		VALUES ($1, $2)
		RETURNING id, name, description;
	`

	args := []any{
		input.Name,
		input.Description,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var eventType EventType
	err := m.db.GetContext(ctx, &eventType, stmt, args...)

	return &eventType, err
}

func (m *EventTypesModel) GetAll() ([]*EventType, error) {
	stmt := `
		SELECT id, name, description
		FROM event_types
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var eventTypes []*EventType
	err := m.db.SelectContext(ctx, &eventTypes, stmt)

	return eventTypes, err
}

func (m *EventTypesModel) GetByID(id string) (*EventType, error) {
	stmt := `
		SELECT id, name, description
		FROM event_types
		WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var eventType EventType
	err := m.db.GetContext(ctx, &eventType, stmt, id)

	return &eventType, err
}

func (m *EventTypesModel) Update(input EventType) (*EventType, error) {
	updateStmt := `
		UPDATE event_types
		SET name = $1, description = $2
		WHERE id = $3
		RETURNING id, name, description;
	`

	args := []any{
		input.Name,
		input.Description,
		input.ID,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var eventType EventType
	err := m.db.GetContext(ctx, &eventType, updateStmt, args...)

	return &eventType, err
}

func (m *EventTypesModel) Delete(id string) (*EventType, error) {
	stmt := `
		DELETE FROM event_types
		WHERE id = $1
		RETURNING id, name, description;
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var eventType EventType
	err := m.db.GetContext(ctx, &eventType, stmt, id)

	return &eventType, err
}
