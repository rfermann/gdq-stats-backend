package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"

	db_models "github.com/rfermann/gdq-stats-backend/internal/db/models"
	"github.com/rfermann/gdq-stats-backend/internal/errors"
	"github.com/rfermann/gdq-stats-backend/internal/gql"
	"github.com/samber/lo"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type EventService struct {
	db *sql.DB
}

func (e *EventService) GetCurrentEvent() (*db_models.Event, error) {
	event, err := db_models.Events(db_models.EventWhere.ActiveEvent.EQ(true)).One(context.Background(), e.db)
	if err != nil {
		return nil, errors.ErrRecordNotFound
	}

	return event, nil
}

func (e *EventService) GetEvents() ([]*db_models.Event, error) {
	events, err := db_models.Events(qm.OrderBy(fmt.Sprintf("%s desc", db_models.EventColumns.EndDate))).All(context.Background(), e.db)
	if err != nil {
		return nil, errors.ErrRecordNotFound
	}

	return events, nil
}

func (e *EventService) GetEventTypeByID(id string) (*db_models.EventType, error) {
	eventType, err := db_models.FindEventType(context.Background(), e.db, id)
	if err != nil {
		return nil, errors.ErrRecordNotFound
	}

	return eventType, nil
}

func (e *EventService) GetEventTypes() ([]*db_models.EventType, error) {
	eventTypes, err := db_models.EventTypes().All(context.Background(), e.db)
	if err != nil {
		return nil, errors.ErrRecordNotFound
	}

	return eventTypes, nil
}

func (e *EventService) CreateEventType(input gql.CreateEventTypeInput) (*db_models.EventType, error) {
	eventType := db_models.EventType{
		Description: input.Description,
		Name:        input.Name,
	}

	err := eventType.Insert(context.Background(), e.db, boil.Infer())
	if err != nil {
		return nil, err
	}

	return &eventType, nil
}

func (e *EventService) DeleteEventType(input gql.DeleteEventTypeInput) (*db_models.EventType, error) {
	eventType, err := db_models.FindEventType(context.Background(), e.db, input.ID)
	if err != nil {
		return nil, err
	}

	eventType.Delete(context.Background(), e.db)

	return eventType, nil
}

func (e *EventService) UpdateEventType(input gql.UpdateEventTypeInput) (*db_models.EventType, error) {
	eventType, err := db_models.FindEventType(context.Background(), e.db, input.ID)
	if err != nil {
		return nil, errors.ErrRecordNotFound
	}

	if input.Description != nil {
		eventType.Description = *input.Description
	}

	if input.Name != nil {
		eventType.Name = *input.Name
	}

	_, err = eventType.Update(context.Background(), e.db, boil.Infer())
	if err != nil {
		return nil, errors.ErrRecordNotFound
	}

	return eventType, nil
}

type eventDataStruct struct {
	ChatCount       int64   `json:"c"`
	DonationAmount  float64 `json:"m"`
	DonationCount   int64   `json:"d"`
	MaxViewersCount int64   `json:"v"`
	TweetsCount     int64   `json:"t"`
}

type scheduleDataStruct struct {
	Duration  string
	Runner    string
	Title     string
	StartTime time.Time
}

type eventStatsStruct struct {
	ChatCount           int64
	CompletedGamesCount int64
	DonationAmount      float64
	DonationCount       int64
	MaxViewersCount     int64
	TweetsCount         int64
}

func (e *EventService) MigrateEventData(input gql.MigrateEventDataInput) (*db_models.Event, error) {
	event, err := db_models.Events(
		db_models.EventWhere.ID.EQ(input.ID),
		qm.InnerJoin(
			db_models.TableNames.EventTypes+" on "+
				db_models.TableNames.Events+"."+
				db_models.EventColumns.EventTypeID+"="+
				db_models.TableNames.EventTypes+"."+
				db_models.EventTypeColumns.ID,
		),
		qm.Load(db_models.EventRels.EventType),
	).One(context.Background(), e.db)
	if err != nil {
		return nil, errors.ErrRecordNotFound
	}

	var eventData []eventDataStruct
	var scheduleData []scheduleDataStruct

	eventName := event.R.EventType.Name
	if eventName == "SGDQ" && event.Year == 2016 {
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
			eventDataUrl = fmt.Sprintf("https://gdqstats.com/data/%d/%s_final/latest.json", event.Year, strings.ToLower(eventName))
			scheduleDataUrl = fmt.Sprintf("https://gdqstats.com/data/%d/%s_final/schedule.json", event.Year, strings.ToLower(eventName))
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

	var eventStatsData *eventStatsStruct
	eventStatsData, err = extractEventData(event, eventData, scheduleData)
	if err != nil {
		return nil, err
	}

	event.CompletedGamesCount = eventStatsData.CompletedGamesCount
	event.DonationAmount = eventStatsData.DonationAmount
	event.DonationCount = eventStatsData.DonationCount
	event.MaxViewersCount = eventStatsData.MaxViewersCount
	event.TweetsCount = eventStatsData.TweetsCount
	event.TwitchChatsCount = eventStatsData.ChatCount

	event.Update(
		context.Background(),
		e.db,
		boil.Whitelist(
			db_models.EventColumns.CompletedGamesCount,
			db_models.EventColumns.DonationAmount,
			db_models.EventColumns.DonationCount,
			db_models.EventColumns.MaxViewersCount,
			db_models.EventColumns.TweetsCount,
			db_models.EventColumns.TwitchChatsCount,
		),
	)

	return event, nil
}

func extractEventData(event *db_models.Event, eventData []eventDataStruct, scheduleData []scheduleDataStruct) (*eventStatsStruct, error) {
	eventDataSum := lo.Reduce(eventData, func(agg eventDataStruct, eventItem eventDataStruct, _ int) eventDataStruct {
		if eventItem.MaxViewersCount > agg.MaxViewersCount {
			agg.MaxViewersCount = eventItem.MaxViewersCount
		}
		return eventDataStruct{
			ChatCount:       agg.ChatCount + eventItem.ChatCount,
			DonationCount:   eventItem.DonationCount,
			MaxViewersCount: agg.MaxViewersCount,
			TweetsCount:     agg.TweetsCount + eventItem.TweetsCount,
			DonationAmount:  eventItem.DonationAmount,
		}
	}, eventDataStruct{})

	now := time.Now()

	completedGamesCount := lo.Reduce(scheduleData, func(agg int64, scheduleItem scheduleDataStruct, _ int) int64 {
		parsedDuration, err := parseStringDuration(scheduleItem.Duration)
		if err != nil {
			return agg
		}
		endDate := scheduleItem.StartTime.Add(parsedDuration)

		if endDate.Before(now) {
			return agg + 1
		}
		return agg
	}, 0)

	return &eventStatsStruct{
		CompletedGamesCount: completedGamesCount,
		ChatCount:           eventDataSum.ChatCount,
		DonationAmount:      eventDataSum.DonationAmount,
		DonationCount:       eventDataSum.DonationCount,
		MaxViewersCount:     eventDataSum.MaxViewersCount,
		TweetsCount:         eventDataSum.TweetsCount,
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
					MaxViewersCount: int64(currentValue.(float64)),
				})
			}
		}
	}

	eventData = append(eventData, eventDataStruct{
		ChatCount:      responseData.Stats.ChatCount,
		DonationAmount: responseData.Stats.DonationAmount,
		DonationCount:  responseData.Stats.DonationCount,
		TweetsCount:    responseData.Stats.TweetsCount,
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
				StartTime: startTime,
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
