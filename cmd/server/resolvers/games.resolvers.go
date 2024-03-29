package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"

	"github.com/rfermann/gdq-stats-backend/internal/data"
	"github.com/rfermann/gdq-stats-backend/internal/gql"
)

// GetGames is the resolver for the getGames field.
func (r *queryResolver) GetGames(ctx context.Context, input *gql.GetGamesInput) ([]*data.Game, error) {
	return r.Services.GamesService.GetGames(input)
}
