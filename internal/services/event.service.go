package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rfermann/gdq-stats-backend/internal/data"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/rfermann/gdq-stats-backend/internal/errors"
	"github.com/rfermann/gdq-stats-backend/internal/gql"
	"github.com/samber/lo"
)

type EventService struct {
	db *sqlx.DB
}

func (e *EventService) GetCurrentEvent() (*data.Event, error) {
	stmt := `
		SELECT id, year, start_date, end_date, 
		       active_event, viewers, donations, 
		       donors, games_completed, twitch_chats, 
		       tweets, schedule_id, event_type_id
		FROM events
		WHERE active_event = TRUE
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var event data.Event
	err := e.db.GetContext(ctx, &event, stmt)
	if err != nil {
		fmt.Println("error getting current event:", err)
		return nil, errors.ErrRecordNotFound
	}

	return &event, nil
}

func (e *EventService) GetEvents() ([]*data.Event, error) {
	stmt := `
		SELECT id, year, start_date, end_date, 
		       active_event, viewers, donations, 
		       donors, games_completed, twitch_chats, 
		       tweets, schedule_id, event_type_id
		FROM events
		ORDER BY end_date DESC 
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var events []*data.Event
	err := e.db.SelectContext(ctx, &events, stmt)
	if err != nil {
		return nil, errors.ErrRecordNotFound
	}

	return events, nil
}

func (e *EventService) GetEventData(input *gql.GetEventDataInput) (*gql.EventDataResponse, error) {
	if input.Event == nil {
		stmt := fmt.Sprintf(`
			SELECT timestamp, event_data.%s
			FROM event_data
			INNER JOIN events ON events.id = event_data.event_id
			WHERE active_event = TRUE
			AND  event_data.%s > 0
			ORDER BY TIMESTAMP
		`, input.EventDataType.String(), input.EventDataType.String())

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		var eventData []*data.EventDatum
		err := e.db.SelectContext(ctx, &eventData, stmt)
		if err != nil {
			return nil, errors.ErrRecordNotFound
		}

		eventData = lo.Filter(eventData, func(item *data.EventDatum, index int) bool {
			return index%2 == 0
		})

		return &gql.EventDataResponse{
			EventDataType: input.EventDataType,
			EventData:     eventData,
		}, nil
	} else {
		fmt.Println("input", input)
	}
	return nil, nil
}

func (e *EventService) GetEventTypeByID(id string) (*data.EventType, error) {
	stmt := `
		SELECT id, name, description
		FROM event_types
		WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var eventType data.EventType
	err := e.db.GetContext(ctx, &eventType, stmt, id)
	if err != nil {
		return nil, errors.ErrRecordNotFound
	}

	return &eventType, nil
}

func (e *EventService) GetEventTypes() ([]*data.EventType, error) {
	stmt := `
		SELECT id, name, description
		FROM event_types
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var eventTypes []*data.EventType
	err := e.db.SelectContext(ctx, &eventTypes, stmt)
	if err != nil {
		return nil, errors.ErrRecordNotFound
	}

	return eventTypes, nil
}

func (e *EventService) CreateEventType(input gql.CreateEventTypeInput) (*data.EventType, error) {
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

	var eventType data.EventType
	err := e.db.GetContext(ctx, &eventType, stmt, args...)
	if err != nil {
		return nil, err
	}

	return &eventType, nil
}

func (e *EventService) DeleteEventType(input gql.DeleteEventTypeInput) (*data.EventType, error) {
	stmt := `
		DELETE FROM event_types
		WHERE id = $1
		RETURNING id, name, description;
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var eventType data.EventType
	err := e.db.GetContext(ctx, &eventType, stmt, input.ID)
	if err != nil {
		return nil, err
	}

	return &eventType, nil
}

func (e *EventService) UpdateEventType(input gql.UpdateEventTypeInput) (*data.EventType, error) {
	readStmt := `
		SELECT id, name, description
		FROM event_types
		WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var eventType data.EventType
	err := e.db.GetContext(ctx, &eventType, readStmt, input.ID)
	if err != nil {
		return nil, errors.ErrRecordNotFound
	}

	if input.Description != nil {
		eventType.Description = *input.Description
	}

	if input.Name != nil {
		eventType.Name = *input.Name
	}

	updateStmt := `
		UPDATE event_types
		SET name = $1, description = $2
		WHERE id = $3
		RETURNING id, name, description;
	`

	args := []any{
		input.ID,
		eventType.Name,
		eventType.Description,
	}

	ctx, cancel = context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err = e.db.GetContext(ctx, &eventType, updateStmt, args...)

	return &eventType, nil
}

type eventDataStruct struct {
	Timestamp      time.Time `json:"time"`
	Donations      *float64  `json:"m"`
	Donors         int64     `json:"d"`
	GamesCompleted int64
	Tweets         int64 `json:"t"`
	TwitchChats    int64 `json:"c"`
	Viewers        int64 `json:"v"`
}

type scheduleDataStruct struct {
	Duration  string `json:"duration"`
	Runner    string `json:"runners"`
	Title     string `json:"name"`
	StartTime string `json:"start_time"`
}

func (e *EventService) MigrateEventData(input gql.MigrateEventDataInput) (*data.Event, error) {
	getEventStmt := `
		SELECT * 
		FROM events
		WHERE id = $1;
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var event data.Event
	err := e.db.GetContext(ctx, &event, getEventStmt, input.ID)
	if err != nil {
		return nil, errors.ErrRecordNotFound
	}

	var eventData []eventDataStruct
	var scheduleData []scheduleDataStruct

	getEventTypeStmt := `
		SELECT name
		FROM event_types
		WHERE id = $1;
	`

	var eventTypeName string
	err = e.db.GetContext(ctx, &eventTypeName, getEventTypeStmt, event.EventTypeID)
	if err != nil {
		return nil, errors.ErrRecordNotFound
	}

	if eventTypeName == "SGDQ" && event.Year == 2016 {
		eventData, scheduleData, err = extractEventDataSGDQ2016()
		if err != nil {
			return nil, err
		}
	} else {
		eventDataUrl := ""
		scheduleDataUrl := ""
		if event.ActiveEvent {
			eventDataUrl = "https://storage.gdqstats.com/latest.json"
			scheduleDataUrl = "https://storage.gdqstats.com/schedule.json"
		} else {
			eventDataUrl = fmt.Sprintf("https://gdqstats.com/data/%d/%s_final/latest.json", event.Year, strings.ToLower(eventTypeName))
			scheduleDataUrl = fmt.Sprintf("https://gdqstats.com/data/%d/%s_final/schedule.json", event.Year, strings.ToLower(eventTypeName))
		}

		r, err := http.Get(eventDataUrl)
		if err != nil {
			return nil, err
		}

		dec := json.NewDecoder(r.Body)
		err = dec.Decode(&eventData)
		if err != nil {
			return nil, err
		}

		r, err = http.Get(scheduleDataUrl)
		if err != nil {
			return nil, err
		}

		dec = json.NewDecoder(r.Body)
		err = dec.Decode(&scheduleData)
		if err != nil {
			return nil, err
		}
	}

	deleteEventDataStmt := `DELETE FROM event_data WHERE event_id = $1;`

	ctx, cancel = context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err = e.db.ExecContext(ctx, deleteEventDataStmt, event.ID)
	if err != nil {
		return nil, err
	}

	deleteGameDataStmt := `DELETE FROM games WHERE event_id = $1;`

	ctx, cancel = context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err = e.db.ExecContext(ctx, deleteGameDataStmt, event.ID)
	if err != nil {
		return nil, err
	}

	var eventStatsData *eventDataStruct
	fmt.Println("event.ID: ", event.ID)
	eventStatsData, err = extractEventData(event.ID, eventData, scheduleData, e.db)
	if err != nil {
		return nil, err
	}

	if eventStatsData.Donations != nil {
		event.Donations = *eventStatsData.Donations
	}

	event.Donors = eventStatsData.Donors
	event.GamesCompleted = eventStatsData.GamesCompleted
	event.Tweets = eventStatsData.Tweets
	event.TwitchChats = eventStatsData.TwitchChats
	event.Viewers = eventStatsData.Viewers

	updateEventStmt := `
		UPDATE events
		SET donations = $1, donors =$2, games_completed = $3, tweets = $4, twitch_chats = $5, viewers = $6
		WHERE id = $7;
	`

	args := []any{
		event.Donations,
		event.Donors,
		event.GamesCompleted,
		event.Tweets,
		event.TwitchChats,
		event.Viewers,
		event.ID,
	}

	ctx, cancel = context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err = e.db.ExecContext(ctx, updateEventStmt, args...)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (e *EventService) GetGames(input *gql.EventDataInput) ([]*data.Game, error) {
	fmt.Println("input", input)
	if input == nil {
		stmt := `
			SELECT g.id, g.start_date, g.end_date, duration, name, runner, event_id
			FROM games g
			INNER JOIN events e ON e.id = g.event_id
			WHERE e.active_event = TRUE;
		`

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		var games []*data.Game
		err := e.db.SelectContext(ctx, &games, stmt)
		if err != nil {
			return nil, errors.ErrRecordNotFound
		}

		return games, nil
	}

	return nil, nil
}

// TODO: check if this method can be combined with GetEvents
// TODO: simplify query: it only fetches non-active events; no need for being overly complicated
func (e *EventService) GetAlternativeEvents() ([]*data.Event, error) {
	stmt := `
			SELECT c.id, p.*
			FROM events p, events c
			WHERE c.active_event = TRUE
			  AND p.id <> c.id
			ORDER BY p.start_date DESC;
		`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var events []*data.Event
	err := e.db.SelectContext(ctx, &events, stmt)
	if err != nil {
		return nil, errors.ErrRecordNotFound
	}

	return events, nil
}

func extractEventData(eventId string, eventData []eventDataStruct, scheduleData []scheduleDataStruct, db *sqlx.DB) (*eventDataStruct, error) {
	lastDonation := float64(0)

	donationsPerMinute := float64(0)
	eventDataSum := lo.Reduce(eventData, func(agg eventDataStruct, eventItem eventDataStruct, count int) eventDataStruct {
		if eventItem.Viewers > agg.Viewers {
			agg.Viewers = eventItem.Viewers
		}

		if eventItem.Donations != nil {
			donationsPerMinute = *eventItem.Donations - lastDonation
			if lastDonation < *eventItem.Donations {
				lastDonation = *eventItem.Donations
			}
		}

		insertEventDatumStmt := `
			INSERT INTO event_data (timestamp, donations, donations_per_minute, donors, tweets, tweets_per_minute, twitch_chats, twitch_chats_per_minute, viewers, event_id)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		`
		args := []any{
			eventItem.Timestamp,
			lastDonation,
			donationsPerMinute,
			eventItem.Donors,
			agg.Tweets + eventItem.Tweets,
			eventItem.Tweets,
			agg.TwitchChats + eventItem.TwitchChats,
			eventItem.TwitchChats,
			eventItem.Viewers,
			eventId,
		}

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		_, _ = db.ExecContext(ctx, insertEventDatumStmt, args...)

		return eventDataStruct{
			Donations:   eventItem.Donations,
			Donors:      eventItem.Donors,
			Tweets:      agg.Tweets + eventItem.Tweets,
			TwitchChats: agg.TwitchChats + eventItem.TwitchChats,
			Viewers:     agg.Viewers,
		}
	}, eventDataStruct{})

	now := time.Now()

	completedGamesCount := lo.Reduce(scheduleData, func(agg int64, scheduleItem scheduleDataStruct, _ int) int64 {
		parsedDuration, err := parseStringDuration(scheduleItem.Duration)
		if err != nil {
			return agg
		}

		startTime, err := time.Parse("2006-01-02T15:04:05", scheduleItem.StartTime)
		if err != nil {
			return agg
		}

		endDate := startTime.Add(parsedDuration)

		insertGameDataStmt := `
			INSERT INTO games ( start_date, end_date, duration, name, runner, event_id)
			VALUES ($1, $2, $3, $4, $5, $6)
		`
		args := []any{
			startTime,
			endDate,
			scheduleItem.Duration,
			scheduleItem.Title,
			scheduleItem.Runner,
			eventId,
		}

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		_, _ = db.ExecContext(ctx, insertGameDataStmt, args...)

		if endDate.Before(now) {
			return agg + 1
		}
		return agg
	}, 0)

	return &eventDataStruct{
		GamesCompleted: completedGamesCount,
		TwitchChats:    eventDataSum.TwitchChats,
		Donations:      eventDataSum.Donations,
		Donors:         eventDataSum.Donors,
		Viewers:        eventDataSum.Viewers,
		Tweets:         eventDataSum.Tweets,
	}, nil
}

func extractEventDataSGDQ2016() ([]eventDataStruct, []scheduleDataStruct, error) {
	type statsStruct struct {
		ChatCount      int64   `json:"total_chats"`
		DonationAmount float64 `json:"total_donations"`
		DonationCount  int64   `json:"num_donators"`
		TweetsCount    int64   `json:"total_tweets"`
	}

	type responseStruct struct {
		Data  map[string]map[string]interface{}
		Games map[string]map[string]interface{}
		Stats statsStruct
	}

	var responseData responseStruct
	r, err := http.Get("https://gdqstats.com/data/2016/sgdq2016final.json")
	if err != nil {
		return nil, nil, err
	}

	dec := json.NewDecoder(r.Body)
	err = dec.Decode(&responseData)
	if err != nil {
		return nil, nil, err
	}

	var eventData []eventDataStruct
	dataKeys := reflect.ValueOf(responseData.Data)
	if dataKeys.Kind() == reflect.Map {
		for _, v := range dataKeys.MapKeys() {
			value := dataKeys.MapIndex(v).Interface()
			currentValue := value.(map[string]interface{})["v"]
			if reflect.TypeOf(currentValue) != nil {
				eventData = append(eventData, eventDataStruct{
					Viewers: int64(currentValue.(float64)),
				})
			}
		}
	}

	eventData = append(eventData, eventDataStruct{
		TwitchChats: responseData.Stats.ChatCount,
		Donations:   &responseData.Stats.DonationAmount,
		Donors:      responseData.Stats.DonationCount,
		Tweets:      responseData.Stats.TweetsCount,
	})

	var scheduleData []scheduleDataStruct
	scheduleKeys := reflect.ValueOf(responseData.Games)
	if scheduleKeys.Kind() == reflect.Map {
		for _, v := range scheduleKeys.MapKeys() {
			duration := time.Duration(int64(responseData.Games[v.String()]["start_time"].(float64)) * 1000 * 1000)
			startTime := time.Unix(0, 0).Add(duration)

			scheduleData = append(scheduleData, scheduleDataStruct{
				Duration:  responseData.Games[v.String()]["duration"].(string),
				Runner:    responseData.Games[v.String()]["runner"].(string),
				StartTime: startTime.String(),
				Title:     responseData.Games[v.String()]["title"].(string),
			})
		}
	}

	return eventData, scheduleData, err
}

func parseStringDuration(duration string) (time.Duration, error) {
	newDuration := strings.Replace(duration, ":", "h", 1)
	newDuration = strings.Replace(newDuration, ":", "m", 1)
	newDuration = newDuration + "s"

	return time.ParseDuration(newDuration)
}
