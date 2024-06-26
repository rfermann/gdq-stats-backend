package models

import (
	"context"
	"github.com/jmoiron/sqlx"
	"strings"
	"time"
)

type Event struct {
	ID             string
	Year           int64
	StartDate      time.Time `db:"start_date"`
	ActiveEvent    bool      `db:"active_event"`
	Viewers        int64
	Donations      float64
	Donors         int64
	CompletedGames int64 `db:"completed_games"`
	TotalGames     int64 `db:"total_games"`
	TwitchChats    int64 `db:"twitch_chats"`
	Tweets         int64
	EventDataCount int64  `db:"event_data_count"`
	ScheduleID     int64  `db:"schedule_id"`
	EventTypeID    string `db:"event_type_id"`
}

type EventsModel struct {
	db *sqlx.DB
}

func (m *EventsModel) Insert(eventInput *Event) (*Event, error) {
	stmt := `
		INSERT INTO events(year, start_date, schedule_id, event_type_id)
		VALUES ($1, $2, $3, $4)
		RETURNING *;
	`

	args := []any{
		eventInput.Year,
		eventInput.StartDate,
		eventInput.ScheduleID,
		eventInput.EventTypeID,
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	var event Event
	err := m.db.GetContext(ctx, &event, stmt, args...)

	return &event, err
}

func (m *EventsModel) GetActive() (*Event, error) {
	stmt := `
		SELECT *
		FROM events
		WHERE active_event = TRUE
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var event Event
	err := m.db.GetContext(ctx, &event, stmt)

	return &event, err
}

func (m *EventsModel) GetAll() ([]*Event, error) {
	stmt := `
		SELECT *
		FROM events
		ORDER BY start_date DESC 
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var events []*Event
	err := m.db.SelectContext(ctx, &events, stmt)

	return events, err
}

func (m *EventsModel) GetById(id string) (*Event, error) {
	stmt := `
		SELECT * 
		FROM events
		WHERE id = $1;
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var event Event
	err := m.db.GetContext(ctx, &event, stmt, id)

	return &event, err
}

func (m *EventsModel) GetInactive() ([]*Event, error) {
	stmt := `
		SELECT *
		FROM events
		WHERE active_event = FALSE
		ORDER BY start_date DESC;
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var events []*Event
	err := m.db.SelectContext(ctx, &events, stmt)

	return events, err
}

func (m *EventsModel) Update(event Event) (*Event, error) {
	stmt := `
		UPDATE events
		SET donations = $1, donors =$2, completed_games = $3, 
		    total_games = $4, tweets = $5, twitch_chats = $6, 
		    viewers = $7, event_data_count = $8
		WHERE id = $9
		RETURNING *;
	`

	args := []any{
		event.Donations,
		event.Donors,
		event.CompletedGames,
		event.TotalGames,
		event.Tweets,
		event.TwitchChats,
		event.Viewers,
		event.EventDataCount,
		event.ID,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var updatedEvent Event
	err := m.db.GetContext(ctx, &updatedEvent, stmt, args...)

	return &updatedEvent, err
}

func (m *EventsModel) GetByScheduleId(id int64) (*Event, error) {
	stmt := `
		SELECT * 
		FROM events 
		WHERE schedule_id = $1;
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var event Event
	err := m.db.GetContext(ctx, &event, stmt, id)

	return &event, err
}

func (m *EventsModel) ActivateById(id string) (*Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	tx, err := m.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}

	activateStmt := `
		UPDATE events s
		SET active_event = TRUE 
		WHERE id = $1 RETURNING *;
	`

	deactivateStmt := `
		UPDATE events
		SET active_event = FALSE
		WHERE id != $1;
	`

	var event Event
	err = tx.GetContext(ctx, &event, activateStmt, id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	_, err = tx.ExecContext(ctx, deactivateStmt, id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &event, err
}

func (m *EventsModel) GetByNameAndYear(name string, year int64) (*Event, error) {
	stmt := `
		SELECT 
		    events.id, year, start_date, active_event, 
		    viewers, donations, donors, completed_games,
		    total_games, twitch_chats, tweets,
		    schedule_id, event_data_count, event_type_id
		FROM events
		INNER JOIN event_types et ON events.event_type_id = et.id
		WHERE et.name = $1
		AND year = $2
	`

	args := []any{
		strings.ToUpper(name),
		year,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var event Event
	err := m.db.GetContext(ctx, &event, stmt, args...)

	return &event, err
}

func (m *EventsModel) GetAlternativeEventsForEventId(id string) ([]*Event, error) {
	stmt := `
		SELECT *
		FROM events
		WHERE id <> $1
		ORDER BY start_date DESC;
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var events []*Event
	err := m.db.SelectContext(ctx, &events, stmt, id)

	return events, err
}
