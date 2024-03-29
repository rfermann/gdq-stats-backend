package services

import (
	"github.com/rfermann/gdq-stats-backend/internal/data"
	"github.com/rfermann/gdq-stats-backend/internal/gql"
)

type GamesService struct {
	models *data.Models
}

func (e *GamesService) GetGames(input *gql.GetGamesInput) ([]*data.Game, error) {
	if input == nil {
		games, err := e.models.Games.GetAllForActiveEvent()
		if err != nil {
			return nil, ErrRecordNotFound
		}

		return games, nil
	}

	return nil, nil
}
