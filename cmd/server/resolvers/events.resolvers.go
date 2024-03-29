package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"

	"github.com/rfermann/gdq-stats-backend/internal/data"
	"github.com/rfermann/gdq-stats-backend/internal/gql"
)

// EventType is the resolver for the eventType field.
func (r *eventResolver) EventType(ctx context.Context, obj *data.Event) (*data.EventType, error) {
	return r.Services.EventService.GetEventTypeByID(obj.EventTypeID)
}

// MigrateEventData is the resolver for the migrateEventData field.
func (r *mutationResolver) MigrateEventData(ctx context.Context, input gql.MigrateEventDataInput) (*data.Event, error) {
	return r.Services.EventService.MigrateEventData(input)
}

// GetAlternativeEvents is the resolver for the getAlternativeEvents field.
func (r *queryResolver) GetAlternativeEvents(ctx context.Context) ([]*data.Event, error) {
	return r.Services.EventService.GetAlternativeEvents()
}

// GetCurrentEvent is the resolver for the getCurrentEvent field.
func (r *queryResolver) GetCurrentEvent(ctx context.Context) (*data.Event, error) {
	return r.Services.EventService.GetCurrentEvent()
}

// GetEvents is the resolver for the getEvents field.
func (r *queryResolver) GetEvents(ctx context.Context) ([]*data.Event, error) {
	return r.Services.EventService.GetEvents()
}

// GetEventData is the resolver for the getEventData field.
func (r *queryResolver) GetEventData(ctx context.Context, input *gql.GetEventDataInput) (*gql.EventDataResponse, error) {
	return r.Services.EventService.GetEventData(input)
}

// Event returns gql.EventResolver implementation.
func (r *Resolver) Event() gql.EventResolver { return &eventResolver{r} }

type eventResolver struct{ *Resolver }
