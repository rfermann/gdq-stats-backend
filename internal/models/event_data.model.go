package models

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

type EventDatum struct {
	ID                   string
	Timestamp            time.Time
	Donations            float64
	DonationsPerMinute   float64 `db:"donations_per_minute"`
	Donors               int64
	Tweets               int64
	TweetsPerMinute      int64 `db:"tweets_per_minute"`
	TwitchChats          int64 `db:"twitch_chats"`
	TwitchChatsPerMinute int64 `db:"twitch_chats_per_minute"`
	Viewers              int64
	EventID              string `db:"event_id"`
}

type EventDatumModel struct {
	db *sqlx.DB
}

type EventDataType string

func (m *EventDatumModel) Insert(eventDatum EventDatum) (*EventDatum, error) {
	stmt := `
		INSERT INTO event_data (
			timestamp, donations, donations_per_minute, 
		    donors, tweets, tweets_per_minute, twitch_chats, 
			twitch_chats_per_minute, viewers, event_id
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING 
		    id, timestamp, donations, donations_per_minute, 
		    donors, tweets, tweets_per_minute, twitch_chats, 
			twitch_chats_per_minute, viewers, event_id;
	`

	args := []any{
		eventDatum.Timestamp,
		eventDatum.Donations,
		eventDatum.DonationsPerMinute,
		eventDatum.Donors,
		eventDatum.Tweets,
		eventDatum.TweetsPerMinute,
		eventDatum.TwitchChats,
		eventDatum.TwitchChatsPerMinute,
		eventDatum.Viewers,
		eventDatum.EventID,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var newEventDatum EventDatum
	err := m.db.GetContext(ctx, &newEventDatum, stmt, args...)

	return &newEventDatum, err
}

func (m *EventDatumModel) InsertBulk(eventData []*EventDatum) error {
	stmt := `
		INSERT INTO event_data (
			timestamp, donations, donations_per_minute, 
		    donors, tweets, tweets_per_minute, twitch_chats, 
			twitch_chats_per_minute, viewers, event_id
		)
		VALUES (
		    :timestamp, :donations, :donations_per_minute, 
		    :donors, :tweets, :tweets_per_minute, :twitch_chats, 
			:twitch_chats_per_minute, :viewers, :event_id
		)
		RETURNING 
		    id, timestamp, donations, donations_per_minute, 
		    donors, tweets, tweets_per_minute, twitch_chats, 
			twitch_chats_per_minute, viewers, event_id;
	`

	batchSize := 2000
	batches := make([][]*EventDatum, 0, (len(eventData)+batchSize-1)/batchSize)

	for batchSize < len(eventData) {
		eventData, batches = eventData[batchSize:], append(batches, eventData[0:batchSize:batchSize])
	}
	batches = append(batches, eventData)

	for _, batch := range batches {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		_, err := m.db.NamedExecContext(ctx, stmt, batch)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *EventDatumModel) GetManyByEventId(id string) ([]*EventDatum, error) {
	stmt := `
		SELECT 
		    id, timestamp, donations, donations_per_minute, 
		    donors, tweets, tweets_per_minute, twitch_chats,
		    twitch_chats_per_minute, viewers, event_id
		FROM event_data
		WHERE event_id = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var eventData []*EventDatum
	err := m.db.SelectContext(ctx, &eventData, stmt, id)

	return eventData, err
}

func (m *EventDatumModel) GetForActiveEvent(eventDataType EventDataType) ([]*EventDatum, error) {
	stmt := fmt.Sprintf(`
			SELECT timestamp, event_data.%s
			FROM event_data
			INNER JOIN events ON events.id = event_data.event_id
			WHERE active_event = TRUE
			AND  event_data.%S > 0
			ORDER BY TIMESTAMP
		`, eventDataType, eventDataType)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var eventData []*EventDatum
	err := m.db.SelectContext(ctx, &eventData, stmt)

	return eventData, err
}

func (m *EventDatumModel) DeleteManyByEventId(id string) error {
	stmt := `DELETE FROM event_data WHERE event_id = $1;`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.db.ExecContext(ctx, stmt, id)

	return err
}
