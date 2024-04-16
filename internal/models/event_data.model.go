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

type EventDatumPayload struct {
	Timestamp time.Time
	Value     float64
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
		ORDER BY timestamp
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var eventData []*EventDatum
	err := m.db.SelectContext(ctx, &eventData, stmt, id)

	return eventData, err
}

func (m *EventDatumModel) GetManyByEventIdAndType(id string, eventDataType EventDataType) ([]*EventDatumPayload, error) {
	stmt := fmt.Sprintf(`
			SELECT timestamp, event_data.%s as value
			FROM event_data
			WHERE event_id = $1
			AND  event_data.%s > 0
			ORDER BY TIMESTAMP
		`, eventDataType, eventDataType)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var eventData []*EventDatumPayload
	err := m.db.SelectContext(ctx, &eventData, stmt, id)

	return eventData, err
}

func (m *EventDatumModel) GetForActiveEventAndType(eventDataType EventDataType) ([]*EventDatumPayload, error) {
	stmt := fmt.Sprintf(`
			SELECT timestamp, event_data.%s as value
			FROM event_data
			INNER JOIN events ON events.id = event_data.event_id
			WHERE active_event = TRUE
			AND  event_data.%s > 0
			ORDER BY TIMESTAMP
		`, eventDataType, eventDataType)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var eventData []*EventDatumPayload
	err := m.db.SelectContext(ctx, &eventData, stmt)

	return eventData, err
}

func (m *EventDatumModel) CheckEventDataExistsForEventId(id string) (bool, error) {
	stmt := `
		SELECT EXISTS(SELECT id FROM event_data WHERE event_id = $1);
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var exists bool
	err := m.db.GetContext(ctx, &exists, stmt, id)

	return exists, err
}

func (m *EventDatumModel) GetViewersCountForEventId(id string) (int64, error) {
	stmt := `
		SELECT MAX(viewers)
		FROM event_data
		WHERE event_id = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var viewers int64
	err := m.db.GetContext(ctx, &viewers, stmt, id)

	return viewers, err
}

func (m *EventDatumModel) GetDonationsCountForEventId(id string) (float64, error) {
	stmt := `
		SELECT MAX(donations)
		FROM event_data
		WHERE event_id = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var donations float64
	err := m.db.GetContext(ctx, &donations, stmt, id)

	return donations, err
}

func (m *EventDatumModel) GetDonorsCountForEventId(id string) (int64, error) {
	stmt := `
		SELECT MAX(donors)
		FROM event_data
		WHERE event_id = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var donors int64
	err := m.db.GetContext(ctx, &donors, stmt, id)

	return donors, err
}

func (m *EventDatumModel) GetTweetsCountForEventId(id string) (int64, error) {
	stmt := `
		SELECT MAX(tweets)
		FROM event_data
		WHERE event_id = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var tweets int64
	err := m.db.GetContext(ctx, &tweets, stmt, id)

	return tweets, err
}

func (m *EventDatumModel) GetTwitchChatsCountForEventId(id string) (int64, error) {
	stmt := `
		SELECT MAX(twitch_chats)
		FROM event_data
		WHERE event_id = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var twitch_chats int64
	err := m.db.GetContext(ctx, &twitch_chats, stmt, id)

	return twitch_chats, err
}

func (m *EventDatumModel) GetEventDataCountForEventId(id string) (int64, error) {
	stmt := `
		SELECT COUNT(id)
		FROM event_data
		WHERE event_id = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var count int64
	err := m.db.GetContext(ctx, &count, stmt, id)

	return count, err
}

func (m *EventDatumModel) DeleteManyByEventId(id string) error {
	stmt := `DELETE FROM event_data WHERE event_id = $1;`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.db.ExecContext(ctx, stmt, id)

	return err
}
