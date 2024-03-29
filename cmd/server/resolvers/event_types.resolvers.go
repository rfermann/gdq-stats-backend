package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"

	"github.com/rfermann/gdq-stats-backend/internal/data"
	"github.com/rfermann/gdq-stats-backend/internal/gql"
)

// CreateEventType is the resolver for the createEventType field.
func (r *mutationResolver) CreateEventType(ctx context.Context, input gql.CreateEventTypeInput) (*data.EventType, error) {
	return r.Services.EventTypesService.CreateEventType(input)
}

// DeleteEventType is the resolver for the deleteEventType field.
func (r *mutationResolver) DeleteEventType(ctx context.Context, input gql.DeleteEventTypeInput) (*data.EventType, error) {
	return r.Services.EventTypesService.DeleteEventType(input)
}

// UpdateEventType is the resolver for the updateEventType field.
func (r *mutationResolver) UpdateEventType(ctx context.Context, input gql.UpdateEventTypeInput) (*data.EventType, error) {
	return r.Services.EventTypesService.UpdateEventType(input)
}

// GetEventTypes is the resolver for the getEventTypes field.
func (r *queryResolver) GetEventTypes(ctx context.Context) ([]*data.EventType, error) {
	return r.Services.EventTypesService.GetEventTypes()
}

// Mutation returns gql.MutationResolver implementation.
func (r *Resolver) Mutation() gql.MutationResolver { return &mutationResolver{r} }

// Query returns gql.QueryResolver implementation.
func (r *Resolver) Query() gql.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
