package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"time"

	db_models "github.com/rfermann/gdq-stats-backend/internal/db/models"
	"github.com/rfermann/gdq-stats-backend/internal/gql"
)

// EventType is the resolver for the eventType field.
func (r *eventResolver) EventType(ctx context.Context, obj *db_models.Event) (*db_models.EventType, error) {
	return r.Services.EventService.GetEventTypeByID(obj.EventTypeID)
}

// StartDate is the resolver for the start_date field.
func (r *eventResolver) StartDate(ctx context.Context, obj *db_models.Event) (*time.Time, error) {
	return &obj.StartDate.Time, nil
}

// EndDate is the resolver for the end_date field.
func (r *eventResolver) EndDate(ctx context.Context, obj *db_models.Event) (*time.Time, error) {
	return &obj.EndDate.Time, nil
}

// CreateEventType is the resolver for the createEventType field.
func (r *mutationResolver) CreateEventType(ctx context.Context, input gql.CreateEventTypeInput) (*db_models.EventType, error) {
	return r.Services.EventService.CreateEventType(input)
}

// DeleteEventType is the resolver for the deleteEventType field.
func (r *mutationResolver) DeleteEventType(ctx context.Context, input gql.DeleteEventTypeInput) (*db_models.EventType, error) {
	return r.Services.EventService.DeleteEventType(input)
}

// UpdateEventType is the resolver for the updateEventType field.
func (r *mutationResolver) UpdateEventType(ctx context.Context, input gql.UpdateEventTypeInput) (*db_models.EventType, error) {
	return r.Services.EventService.UpdateEventType(input)
}

// MigrateEventData is the resolver for the migrateEventData field.
func (r *mutationResolver) MigrateEventData(ctx context.Context, input gql.MigrateEventDataInput) (*db_models.Event, error) {
	return r.Services.EventService.MigrateEventData(input)
}

// GetAlternativeEvents is the resolver for the getAlternativeEvents field.
func (r *queryResolver) GetAlternativeEvents(ctx context.Context) ([]*db_models.Event, error) {
	return r.Services.EventService.GetAlternativeEvents()
}

// GetCurrentEvent is the resolver for the getCurrentEvent field.
func (r *queryResolver) GetCurrentEvent(ctx context.Context) (*db_models.Event, error) {
	return r.Services.EventService.GetCurrentEvent()
}

// GetEvents is the resolver for the getEvents field.
func (r *queryResolver) GetEvents(ctx context.Context) ([]*db_models.Event, error) {
	return r.Services.EventService.GetEvents()
}

// GetEventData is the resolver for the getEventData field.
func (r *queryResolver) GetEventData(ctx context.Context, input *gql.GetEventDataInput) (*gql.EventDataResponse, error) {
	return r.Services.EventService.GetEventData(input)
}

// GetEventTypes is the resolver for the getEventTypes field.
func (r *queryResolver) GetEventTypes(ctx context.Context) ([]*db_models.EventType, error) {
	return r.Services.EventService.GetEventTypes()
}

// GetGames is the resolver for the getGames field.
func (r *queryResolver) GetGames(ctx context.Context, input *gql.EventDataInput) ([]*db_models.Game, error) {
	return r.Services.EventService.GetGames(input)
}

// Event returns gql.EventResolver implementation.
func (r *Resolver) Event() gql.EventResolver { return &eventResolver{r} }

// Mutation returns gql.MutationResolver implementation.
func (r *Resolver) Mutation() gql.MutationResolver { return &mutationResolver{r} }

// Query returns gql.QueryResolver implementation.
func (r *Resolver) Query() gql.QueryResolver { return &queryResolver{r} }

type eventResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
