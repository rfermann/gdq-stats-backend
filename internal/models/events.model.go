package models

import (
	"context"
	"github.com/jmoiron/sqlx"
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
	CompletedGames int64 `db:"games_completed"`
	TotalGames     int64 `db:"games_completed"`
	TwitchChats    int64 `db:"twitch_chats"`
	Tweets         int64
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
		SET donations = $1, donors =$2, games_completed = $3, 
		    tweets = $4, twitch_chats = $5, viewers = $6
		WHERE id = $7
		RETURNING *;
	`

	args := []any{
		event.Donations,
		event.Donors,
		event.CompletedGames,
		event.Tweets,
		event.TwitchChats,
		event.Viewers,
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
