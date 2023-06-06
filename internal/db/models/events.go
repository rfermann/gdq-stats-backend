// Code generated by SQLBoiler 4.14.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package db_models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Event is an object representing the database table.
type Event struct {
	ID             string    `boil:"id" json:"id" toml:"id" yaml:"id"`
	Year           int64     `boil:"year" json:"year" toml:"year" yaml:"year"`
	StartDate      null.Time `boil:"start_date" json:"start_date,omitempty" toml:"start_date" yaml:"start_date,omitempty"`
	EndDate        null.Time `boil:"end_date" json:"end_date,omitempty" toml:"end_date" yaml:"end_date,omitempty"`
	ActiveEvent    bool      `boil:"active_event" json:"active_event" toml:"active_event" yaml:"active_event"`
	Viewers        int64     `boil:"viewers" json:"viewers" toml:"viewers" yaml:"viewers"`
	Donations      float64   `boil:"donations" json:"donations" toml:"donations" yaml:"donations"`
	Donors         int64     `boil:"donors" json:"donors" toml:"donors" yaml:"donors"`
	GamesCompleted int64     `boil:"games_completed" json:"games_completed" toml:"games_completed" yaml:"games_completed"`
	TwitchChats    int64     `boil:"twitch_chats" json:"twitch_chats" toml:"twitch_chats" yaml:"twitch_chats"`
	Tweets         int64     `boil:"tweets" json:"tweets" toml:"tweets" yaml:"tweets"`
	ScheduleID     int64     `boil:"schedule_id" json:"schedule_id" toml:"schedule_id" yaml:"schedule_id"`
	EventTypeID    string    `boil:"event_type_id" json:"event_type_id" toml:"event_type_id" yaml:"event_type_id"`

	R *eventR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L eventL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var EventColumns = struct {
	ID             string
	Year           string
	StartDate      string
	EndDate        string
	ActiveEvent    string
	Viewers        string
	Donations      string
	Donors         string
	GamesCompleted string
	TwitchChats    string
	Tweets         string
	ScheduleID     string
	EventTypeID    string
}{
	ID:             "id",
	Year:           "year",
	StartDate:      "start_date",
	EndDate:        "end_date",
	ActiveEvent:    "active_event",
	Viewers:        "viewers",
	Donations:      "donations",
	Donors:         "donors",
	GamesCompleted: "games_completed",
	TwitchChats:    "twitch_chats",
	Tweets:         "tweets",
	ScheduleID:     "schedule_id",
	EventTypeID:    "event_type_id",
}

var EventTableColumns = struct {
	ID             string
	Year           string
	StartDate      string
	EndDate        string
	ActiveEvent    string
	Viewers        string
	Donations      string
	Donors         string
	GamesCompleted string
	TwitchChats    string
	Tweets         string
	ScheduleID     string
	EventTypeID    string
}{
	ID:             "events.id",
	Year:           "events.year",
	StartDate:      "events.start_date",
	EndDate:        "events.end_date",
	ActiveEvent:    "events.active_event",
	Viewers:        "events.viewers",
	Donations:      "events.donations",
	Donors:         "events.donors",
	GamesCompleted: "events.games_completed",
	TwitchChats:    "events.twitch_chats",
	Tweets:         "events.tweets",
	ScheduleID:     "events.schedule_id",
	EventTypeID:    "events.event_type_id",
}

// Generated where

type whereHelpernull_Time struct{ field string }

func (w whereHelpernull_Time) EQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Time) NEQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Time) LT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Time) LTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Time) GT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Time) GTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_Time) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Time) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

type whereHelperbool struct{ field string }

func (w whereHelperbool) EQ(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperbool) NEQ(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperbool) LT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperbool) LTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperbool) GT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperbool) GTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

var EventWhere = struct {
	ID             whereHelperstring
	Year           whereHelperint64
	StartDate      whereHelpernull_Time
	EndDate        whereHelpernull_Time
	ActiveEvent    whereHelperbool
	Viewers        whereHelperint64
	Donations      whereHelperfloat64
	Donors         whereHelperint64
	GamesCompleted whereHelperint64
	TwitchChats    whereHelperint64
	Tweets         whereHelperint64
	ScheduleID     whereHelperint64
	EventTypeID    whereHelperstring
}{
	ID:             whereHelperstring{field: "\"events\".\"id\""},
	Year:           whereHelperint64{field: "\"events\".\"year\""},
	StartDate:      whereHelpernull_Time{field: "\"events\".\"start_date\""},
	EndDate:        whereHelpernull_Time{field: "\"events\".\"end_date\""},
	ActiveEvent:    whereHelperbool{field: "\"events\".\"active_event\""},
	Viewers:        whereHelperint64{field: "\"events\".\"viewers\""},
	Donations:      whereHelperfloat64{field: "\"events\".\"donations\""},
	Donors:         whereHelperint64{field: "\"events\".\"donors\""},
	GamesCompleted: whereHelperint64{field: "\"events\".\"games_completed\""},
	TwitchChats:    whereHelperint64{field: "\"events\".\"twitch_chats\""},
	Tweets:         whereHelperint64{field: "\"events\".\"tweets\""},
	ScheduleID:     whereHelperint64{field: "\"events\".\"schedule_id\""},
	EventTypeID:    whereHelperstring{field: "\"events\".\"event_type_id\""},
}

// EventRels is where relationship names are stored.
var EventRels = struct {
	EventType string
	EventData string
}{
	EventType: "EventType",
	EventData: "EventData",
}

// eventR is where relationships are stored.
type eventR struct {
	EventType *EventType      `boil:"EventType" json:"EventType" toml:"EventType" yaml:"EventType"`
	EventData EventDatumSlice `boil:"EventData" json:"EventData" toml:"EventData" yaml:"EventData"`
}

// NewStruct creates a new relationship struct
func (*eventR) NewStruct() *eventR {
	return &eventR{}
}

func (r *eventR) GetEventType() *EventType {
	if r == nil {
		return nil
	}
	return r.EventType
}

func (r *eventR) GetEventData() EventDatumSlice {
	if r == nil {
		return nil
	}
	return r.EventData
}

// eventL is where Load methods for each relationship are stored.
type eventL struct{}

var (
	eventAllColumns            = []string{"id", "year", "start_date", "end_date", "active_event", "viewers", "donations", "donors", "games_completed", "twitch_chats", "tweets", "schedule_id", "event_type_id"}
	eventColumnsWithoutDefault = []string{"year", "event_type_id"}
	eventColumnsWithDefault    = []string{"id", "start_date", "end_date", "active_event", "viewers", "donations", "donors", "games_completed", "twitch_chats", "tweets", "schedule_id"}
	eventPrimaryKeyColumns     = []string{"id"}
	eventGeneratedColumns      = []string{}
)

type (
	// EventSlice is an alias for a slice of pointers to Event.
	// This should almost always be used instead of []Event.
	EventSlice []*Event
	// EventHook is the signature for custom Event hook methods
	EventHook func(context.Context, boil.ContextExecutor, *Event) error

	eventQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	eventType                 = reflect.TypeOf(&Event{})
	eventMapping              = queries.MakeStructMapping(eventType)
	eventPrimaryKeyMapping, _ = queries.BindMapping(eventType, eventMapping, eventPrimaryKeyColumns)
	eventInsertCacheMut       sync.RWMutex
	eventInsertCache          = make(map[string]insertCache)
	eventUpdateCacheMut       sync.RWMutex
	eventUpdateCache          = make(map[string]updateCache)
	eventUpsertCacheMut       sync.RWMutex
	eventUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var eventAfterSelectHooks []EventHook

var eventBeforeInsertHooks []EventHook
var eventAfterInsertHooks []EventHook

var eventBeforeUpdateHooks []EventHook
var eventAfterUpdateHooks []EventHook

var eventBeforeDeleteHooks []EventHook
var eventAfterDeleteHooks []EventHook

var eventBeforeUpsertHooks []EventHook
var eventAfterUpsertHooks []EventHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Event) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Event) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Event) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Event) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Event) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Event) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Event) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Event) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Event) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddEventHook registers your hook function for all future operations.
func AddEventHook(hookPoint boil.HookPoint, eventHook EventHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		eventAfterSelectHooks = append(eventAfterSelectHooks, eventHook)
	case boil.BeforeInsertHook:
		eventBeforeInsertHooks = append(eventBeforeInsertHooks, eventHook)
	case boil.AfterInsertHook:
		eventAfterInsertHooks = append(eventAfterInsertHooks, eventHook)
	case boil.BeforeUpdateHook:
		eventBeforeUpdateHooks = append(eventBeforeUpdateHooks, eventHook)
	case boil.AfterUpdateHook:
		eventAfterUpdateHooks = append(eventAfterUpdateHooks, eventHook)
	case boil.BeforeDeleteHook:
		eventBeforeDeleteHooks = append(eventBeforeDeleteHooks, eventHook)
	case boil.AfterDeleteHook:
		eventAfterDeleteHooks = append(eventAfterDeleteHooks, eventHook)
	case boil.BeforeUpsertHook:
		eventBeforeUpsertHooks = append(eventBeforeUpsertHooks, eventHook)
	case boil.AfterUpsertHook:
		eventAfterUpsertHooks = append(eventAfterUpsertHooks, eventHook)
	}
}

// One returns a single event record from the query.
func (q eventQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Event, error) {
	o := &Event{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "db_models: failed to execute a one query for events")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Event records from the query.
func (q eventQuery) All(ctx context.Context, exec boil.ContextExecutor) (EventSlice, error) {
	var o []*Event

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "db_models: failed to assign all query results to Event slice")
	}

	if len(eventAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Event records in the query.
func (q eventQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "db_models: failed to count events rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q eventQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "db_models: failed to check if events exists")
	}

	return count > 0, nil
}

// EventType pointed to by the foreign key.
func (o *Event) EventType(mods ...qm.QueryMod) eventTypeQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.EventTypeID),
	}

	queryMods = append(queryMods, mods...)

	return EventTypes(queryMods...)
}

// EventData retrieves all the event_datum's EventData with an executor.
func (o *Event) EventData(mods ...qm.QueryMod) eventDatumQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"event_data\".\"event_id\"=?", o.ID),
	)

	return EventData(queryMods...)
}

// LoadEventType allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (eventL) LoadEventType(ctx context.Context, e boil.ContextExecutor, singular bool, maybeEvent interface{}, mods queries.Applicator) error {
	var slice []*Event
	var object *Event

	if singular {
		var ok bool
		object, ok = maybeEvent.(*Event)
		if !ok {
			object = new(Event)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeEvent)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeEvent))
			}
		}
	} else {
		s, ok := maybeEvent.(*[]*Event)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeEvent)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeEvent))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &eventR{}
		}
		args = append(args, object.EventTypeID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &eventR{}
			}

			for _, a := range args {
				if a == obj.EventTypeID {
					continue Outer
				}
			}

			args = append(args, obj.EventTypeID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`event_types`),
		qm.WhereIn(`event_types.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load EventType")
	}

	var resultSlice []*EventType
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice EventType")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for event_types")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for event_types")
	}

	if len(eventTypeAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.EventType = foreign
		if foreign.R == nil {
			foreign.R = &eventTypeR{}
		}
		foreign.R.Events = append(foreign.R.Events, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.EventTypeID == foreign.ID {
				local.R.EventType = foreign
				if foreign.R == nil {
					foreign.R = &eventTypeR{}
				}
				foreign.R.Events = append(foreign.R.Events, local)
				break
			}
		}
	}

	return nil
}

// LoadEventData allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (eventL) LoadEventData(ctx context.Context, e boil.ContextExecutor, singular bool, maybeEvent interface{}, mods queries.Applicator) error {
	var slice []*Event
	var object *Event

	if singular {
		var ok bool
		object, ok = maybeEvent.(*Event)
		if !ok {
			object = new(Event)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeEvent)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeEvent))
			}
		}
	} else {
		s, ok := maybeEvent.(*[]*Event)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeEvent)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeEvent))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &eventR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &eventR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`event_data`),
		qm.WhereIn(`event_data.event_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load event_data")
	}

	var resultSlice []*EventDatum
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice event_data")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on event_data")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for event_data")
	}

	if len(eventDatumAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.EventData = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &eventDatumR{}
			}
			foreign.R.Event = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.EventID {
				local.R.EventData = append(local.R.EventData, foreign)
				if foreign.R == nil {
					foreign.R = &eventDatumR{}
				}
				foreign.R.Event = local
				break
			}
		}
	}

	return nil
}

// SetEventType of the event to the related item.
// Sets o.R.EventType to related.
// Adds o to related.R.Events.
func (o *Event) SetEventType(ctx context.Context, exec boil.ContextExecutor, insert bool, related *EventType) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"events\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"event_type_id"}),
		strmangle.WhereClause("\"", "\"", 2, eventPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.EventTypeID = related.ID
	if o.R == nil {
		o.R = &eventR{
			EventType: related,
		}
	} else {
		o.R.EventType = related
	}

	if related.R == nil {
		related.R = &eventTypeR{
			Events: EventSlice{o},
		}
	} else {
		related.R.Events = append(related.R.Events, o)
	}

	return nil
}

// AddEventData adds the given related objects to the existing relationships
// of the event, optionally inserting them as new records.
// Appends related to o.R.EventData.
// Sets related.R.Event appropriately.
func (o *Event) AddEventData(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*EventDatum) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.EventID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"event_data\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"event_id"}),
				strmangle.WhereClause("\"", "\"", 2, eventDatumPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.EventID = o.ID
		}
	}

	if o.R == nil {
		o.R = &eventR{
			EventData: related,
		}
	} else {
		o.R.EventData = append(o.R.EventData, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &eventDatumR{
				Event: o,
			}
		} else {
			rel.R.Event = o
		}
	}
	return nil
}

// Events retrieves all the records using an executor.
func Events(mods ...qm.QueryMod) eventQuery {
	mods = append(mods, qm.From("\"events\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"events\".*"})
	}

	return eventQuery{q}
}

// FindEvent retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindEvent(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Event, error) {
	eventObj := &Event{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"events\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, eventObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "db_models: unable to select from events")
	}

	if err = eventObj.doAfterSelectHooks(ctx, exec); err != nil {
		return eventObj, err
	}

	return eventObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Event) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("db_models: no events provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(eventColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	eventInsertCacheMut.RLock()
	cache, cached := eventInsertCache[key]
	eventInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			eventAllColumns,
			eventColumnsWithDefault,
			eventColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(eventType, eventMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(eventType, eventMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"events\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"events\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "db_models: unable to insert into events")
	}

	if !cached {
		eventInsertCacheMut.Lock()
		eventInsertCache[key] = cache
		eventInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Event.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Event) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	eventUpdateCacheMut.RLock()
	cache, cached := eventUpdateCache[key]
	eventUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			eventAllColumns,
			eventPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("db_models: unable to update events, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"events\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, eventPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(eventType, eventMapping, append(wl, eventPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "db_models: unable to update events row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db_models: failed to get rows affected by update for events")
	}

	if !cached {
		eventUpdateCacheMut.Lock()
		eventUpdateCache[key] = cache
		eventUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q eventQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "db_models: unable to update all for events")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db_models: unable to retrieve rows affected for events")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o EventSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("db_models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), eventPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"events\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, eventPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "db_models: unable to update all in event slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db_models: unable to retrieve rows affected all in update all event")
	}
	return rowsAff, nil
}

// Delete deletes a single Event record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Event) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("db_models: no Event provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), eventPrimaryKeyMapping)
	sql := "DELETE FROM \"events\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "db_models: unable to delete from events")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db_models: failed to get rows affected by delete for events")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q eventQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("db_models: no eventQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "db_models: unable to delete all from events")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db_models: failed to get rows affected by deleteall for events")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o EventSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(eventBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), eventPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"events\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, eventPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "db_models: unable to delete all from event slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db_models: failed to get rows affected by deleteall for events")
	}

	if len(eventAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Event) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindEvent(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *EventSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := EventSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), eventPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"events\".* FROM \"events\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, eventPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "db_models: unable to reload all in EventSlice")
	}

	*o = slice

	return nil
}

// EventExists checks if the Event row exists.
func EventExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"events\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "db_models: unable to check if events exists")
	}

	return exists, nil
}

// Exists checks if the Event row exists.
func (o *Event) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return EventExists(ctx, exec, o.ID)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Event) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("db_models: no events provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(eventColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	eventUpsertCacheMut.RLock()
	cache, cached := eventUpsertCache[key]
	eventUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			eventAllColumns,
			eventColumnsWithDefault,
			eventColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			eventAllColumns,
			eventPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("db_models: unable to upsert events, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(eventPrimaryKeyColumns))
			copy(conflict, eventPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryCockroachDB(dialect, "\"events\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(eventType, eventMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(eventType, eventMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		_, _ = fmt.Fprintln(boil.DebugWriter, cache.query)
		_, _ = fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // CockcorachDB doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "db_models: unable to upsert events")
	}

	if !cached {
		eventUpsertCacheMut.Lock()
		eventUpsertCache[key] = cache
		eventUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}
