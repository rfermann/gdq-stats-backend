package services

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/rfermann/gdq-stats-backend/internal/gql"
	"github.com/rfermann/gdq-stats-backend/internal/models"
	"github.com/samber/lo"
)

type EventDataService struct {
	models *models.Models
}

func (e *EventDataService) GetEventData(input *gql.GetEventDataInput) (*gql.EventDataResponse, error) {
	if input.Event == nil {
		eventData, err := e.models.EventData.GetForActiveEventAndType(input.EventDataType)
		if err != nil {
			return nil, ErrRecordNotFound
		}

		eventData = lo.Filter(eventData, func(item *models.EventDatumPayload, index int) bool {
			return index%2 == 0
		})

		return &gql.EventDataResponse{
			EventDataType: input.EventDataType,
			EventData:     eventData,
		}, nil
	} else {
		event, err := e.models.Events.GetByNameAndYear(input.Event.Name, input.Event.Year)
		if err != nil {
			return nil, ErrRecordNotFound
		}

		eventData, err := e.models.EventData.GetManyByEventIdAndType(event.ID, input.EventDataType)
		if err != nil {
			fmt.Printf("err", err)
			return nil, ErrRecordNotFound
		}

		eventData = lo.Filter(eventData, func(item *models.EventDatumPayload, index int) bool {
			return index%2 == 0
		})

		return &gql.EventDataResponse{
			EventDataType: input.EventDataType,
			EventData:     eventData,
		}, nil
	}
}

func (e *EventDataService) MigrateEventData(input gql.MigrateEventDataInput) ([]*models.EventDatum, error) {
	event, err := e.models.Events.GetById(input.EventID)
	if err != nil {
		return nil, ErrRecordNotFound
	}

	eventType, err := e.models.EventTypes.GetByID(event.EventTypeID)
	if err != nil {
		return nil, ErrRecordNotFound
	}

	_ = e.models.EventData.DeleteManyByEventId(event.ID)

	var extractedEventData []*models.EventDatum

	// support old format of event data
	if eventType.Name == "SGDQ" && event.Year == 2016 {
		extractedEventData, err = extractEventDataSGDQ2016()
		if err != nil {
			return nil, ErrUnprocessableEntity
		}
	} else {
		extractedEventData, err = extractEventData(event, eventType)
		if err != nil {
			return nil, ErrUnprocessableEntity
		}
	}

	var eventData []*models.EventDatum
	var lastDonation float64 = 0.0
	var tweets int64 = 0
	var twitchChats int64 = 0
	for _, extractedEventDatum := range extractedEventData {
		donationsPerMinute := extractedEventDatum.Donations - lastDonation
		if lastDonation < extractedEventDatum.Donations {
			lastDonation = extractedEventDatum.Donations
		}
		twitchChats = twitchChats + extractedEventDatum.TwitchChatsPerMinute
		tweets = tweets + extractedEventDatum.TweetsPerMinute

		eventDatum := &models.EventDatum{
			Timestamp:            extractedEventDatum.Timestamp,
			Donations:            extractedEventDatum.Donations,
			DonationsPerMinute:   donationsPerMinute,
			Donors:               extractedEventDatum.Donors,
			Tweets:               tweets,
			TweetsPerMinute:      extractedEventDatum.TweetsPerMinute,
			TwitchChats:          twitchChats,
			TwitchChatsPerMinute: extractedEventDatum.TwitchChatsPerMinute,
			Viewers:              extractedEventDatum.Viewers,
			EventID:              event.ID,
		}
		eventData = append(eventData, eventDatum)
	}

	err = e.models.EventData.InsertBulk(eventData)
	if err != nil {
		return nil, ErrUnprocessableEntity
	}

	return e.models.EventData.GetManyByEventId(event.ID)
}

type responseStruct struct {
	Data   map[int64]datum `json:"data"`
	Extras map[int64]extra `json:"extras"`
}

type datum struct {
	D *int64   `json:"d,omitempty"` // donators
	M *float64 `json:"m,omitempty"` // donations
	V *int64   `json:"v,omitempty"` // viewers
}

type extra struct {
	C int64  `json:"c"`           // twitch chats
	E int64  `json:"e"`           // twitch emotes // not needed
	T *int64 `json:"t,omitempty"` // tweets
}

func extractEventDataSGDQ2016() ([]*models.EventDatum, error) {
	var responseData responseStruct
	responseData, err := readJsonResponse[responseStruct]("https://gdqstats.com/data/2016/sgdq2016final.json")
	if err != nil {
		return nil, ErrUnprocessableEntity
	}

	var dates []int64

	for date := range responseData.Data {
		dates = append(dates, date)
	}

	slices.Sort(dates)

	var eventData []*models.EventDatum
	for _, date := range dates {
		data := responseData.Data[date]
		extras := responseData.Extras[date]

		var donations float64 = 0.0
		var donors int64 = 0
		var tweetsPerMinute int64 = 0
		var twitchChatsPerMinute = extras.C
		var viewers int64 = 0

		if data.D != nil {
			donors = *data.D
		}

		if data.M != nil {
			donations = *data.M
		}

		if extras.T != nil {
			tweetsPerMinute = *extras.T
		}

		if data.V != nil {
			viewers = *data.V
		}

		eventDatum := &models.EventDatum{
			Timestamp:            time.UnixMilli(date),
			Donations:            donations,
			Donors:               donors,
			TweetsPerMinute:      tweetsPerMinute,
			TwitchChatsPerMinute: twitchChatsPerMinute,
			Viewers:              viewers,
		}
		eventData = append(eventData, eventDatum)
	}

	return eventData, nil
}

type eventDatumPayload struct {
	Time time.Time `json:"time"` // timestamp
	V    *int64    `json:"v"`    // viewers
	T    int64     `json:"t"`    // tweets
	C    int64     `json:"c"`    // twitch chats
	E    int64     `json:"e"`    // twitch emotes // not needed
	D    *int64    `json:"d"`    // donators
	M    *float64  `json:"m"`    // donations
}

func extractEventData(event *models.Event, eventType *models.EventType) ([]*models.EventDatum, error) {
	var eventDataPayload []eventDatumPayload
	var eventDataUrl string

	if event.ActiveEvent {
		eventDataUrl = "https://storage.gdqstats.com/latest.json"
	} else {
		eventDataUrl = fmt.Sprintf("https://gdqstats.com/data/%d/%s_final/latest.json", event.Year, strings.ToLower(eventType.Name))
	}

	eventDataPayload, err := readJsonResponse[[]eventDatumPayload](eventDataUrl)
	if err != nil {
		return nil, err
	}

	var eventData []*models.EventDatum
	for _, eventDatum := range eventDataPayload {
		var donations float64 = 0.0
		var donors int64 = 0
		var viewers int64 = 0

		if eventDatum.D != nil {
			donors = *eventDatum.D
		}

		if eventDatum.M != nil {
			donations = *eventDatum.M
		}

		if eventDatum.V != nil {
			viewers = *eventDatum.V
		}

		eventDatum := &models.EventDatum{
			Timestamp:            eventDatum.Time,
			Donations:            donations,
			Donors:               donors,
			TweetsPerMinute:      eventDatum.T,
			TwitchChatsPerMinute: eventDatum.C,
			Viewers:              viewers,
		}
		eventData = append(eventData, eventDatum)
	}

	return eventData, nil
}
