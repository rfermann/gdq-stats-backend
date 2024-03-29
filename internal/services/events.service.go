package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/rfermann/gdq-stats-backend/internal/data"
	"github.com/rfermann/gdq-stats-backend/internal/gql"
	"github.com/samber/lo"
)

type EventsService struct {
	models *data.Models
}

func (e *EventsService) GetCurrentEvent() (*data.Event, error) {
	event, err := e.models.Events.GetActive()
	if err != nil {
		return nil, ErrRecordNotFound
	}

	return event, nil
}

func (e *EventsService) GetEvents() ([]*data.Event, error) {
	events, err := e.models.Events.GetAll()
	if err != nil {
		return nil, ErrRecordNotFound
	}

	return events, nil
}

func (e *EventsService) GetEventData(input *gql.GetEventDataInput) (*gql.EventDataResponse, error) {
	if input.Event == nil {
		eventData, err := e.models.EventData.GetForActiveEvent(input.EventDataType)
		if err != nil {
			return nil, ErrRecordNotFound
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

func (e *EventsService) MigrateEventData(input gql.MigrateEventDataInput) (*data.Event, error) {
	event, err := e.models.Events.GetById(input.ID)
	fmt.Println("event", event)
	if err != nil {
		return nil, ErrRecordNotFound
	}

	var eventData []eventDataStruct
	var scheduleData []scheduleDataStruct

	eventType, err := e.models.EventTypes.GetByID(event.EventTypeID)
	if err != nil {
		return nil, ErrRecordNotFound
	}

	fmt.Println("eventType", eventType.Name)
	if eventType.Name == "SGDQ" && event.Year == 2016 {
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
			eventDataUrl = fmt.Sprintf("https://gdqstats.com/data/%d/%s_final/latest.json", event.Year, strings.ToLower(eventType.Name))
			scheduleDataUrl = fmt.Sprintf("https://gdqstats.com/data/%d/%s_final/schedule.json", event.Year, strings.ToLower(eventType.Name))
		}
		fmt.Println("eventDataUrl", eventDataUrl)
		fmt.Println("scheduleDataUrl", scheduleDataUrl)
		r, err := http.Get(eventDataUrl)
		if err != nil {
			fmt.Println("err:eventData", err)
			return nil, err
		}

		dec := json.NewDecoder(r.Body)
		err = dec.Decode(&eventData)
		if err != nil {
			fmt.Println("err:eventDataDecode", err)
			return nil, err
		}

		r, err = http.Get(scheduleDataUrl)
		if err != nil {
			fmt.Println("err:scheduleData", err)
			return nil, err
		}

		dec = json.NewDecoder(r.Body)
		err = dec.Decode(&scheduleData)
		if err != nil {
			fmt.Println("err:scheduleDataDecode", err)
			return nil, err
		}
	}

	_ = e.models.EventData.DeleteManyByEventId(event.ID)
	_ = e.models.Games.DeleteForEventId(event.ID)

	var eventStatsData *eventDataStruct
	eventStatsData, err = extractEventData(event.ID, eventData, scheduleData, e.models)
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

	_, err = e.models.Events.Update(*event)
	if err != nil {
		return nil, ErrUnprocessableEntity
	}

	return event, nil
}

func (e *EventsService) GetAlternativeEvents() ([]*data.Event, error) {
	events, err := e.models.Events.GetInactive()
	if err != nil {
		return nil, ErrRecordNotFound
	}

	return events, nil
}

func extractEventData(eventId string, eventData []eventDataStruct, scheduleData []scheduleDataStruct, models *data.Models) (*eventDataStruct, error) {
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

		eventDatum := data.EventDatum{
			ID:                   "",
			Timestamp:            eventItem.Timestamp,
			Donations:            lastDonation,
			DonationsPerMinute:   donationsPerMinute,
			Donors:               eventItem.Donors,
			Tweets:               agg.Tweets + eventItem.Tweets,
			TweetsPerMinute:      eventItem.Tweets,
			TwitchChats:          agg.TwitchChats + eventItem.TwitchChats,
			TwitchChatsPerMinute: eventItem.TwitchChats,
			Viewers:              eventItem.Viewers,
			EventID:              eventId,
		}

		_, _ = models.EventData.Insert(eventDatum)

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

		game := &data.Game{
			ID:        "",
			StartDate: startTime,
			EndDate:   endDate,
			Duration:  scheduleItem.Duration,
			Name:      scheduleItem.Title,
			Runner:    scheduleItem.Runner,
			EventID:   eventId,
		}

		_, _ = models.Games.Insert(game)

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
