package services

import (
	"encoding/json"
	"fmt"
	"github.com/rfermann/gdq-stats-backend/internal/gql"
	"github.com/rfermann/gdq-stats-backend/internal/models"
	"net/http"
	"strings"
	"time"
)

type GamesService struct {
	models *models.Models
}

func (e *GamesService) GetGames(input *gql.GetEventInformationInput) ([]*models.Game, error) {
	if input == nil {
		games, err := e.models.Games.GetAllForActiveEvent()
		if err != nil {
			return nil, ErrRecordNotFound
		}

		return games, nil
	}
	fmt.Println("input", input)
	event, err := e.models.Events.GetByNameAndYear(input.Name, input.Year)
	if err != nil {
		fmt.Println("err", err)
		return nil, ErrRecordNotFound
	}

	return e.models.Games.GetAllByEventId(event.ID)
}

type Runner struct {
	RunnerType string `json:"type"`
	Name       string `json:"name"`
}

type Schedule struct {
	GameType  string `json:"type"`
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Runners   []Runner
	StartTime time.Time `json:"starttime"`
	EndTime   time.Time `json:"endtime"`
	Runtime   string    `json:"run_time"`
}

type scheduleResponse struct {
	Schedule []Schedule
}

func (e *GamesService) CreateGames(input gql.MigrateGamesInput) ([]*models.Game, error) {
	r, err := http.Get(fmt.Sprintf("https://gdq-site.vercel.app/api/schedule/%d", input.ScheduleID))
	if err != nil {
		return nil, ErrUnprocessableEntity
	}

	var scheduleResponse scheduleResponse
	dec := json.NewDecoder(r.Body)
	err = dec.Decode(&scheduleResponse)
	if err != nil {
		return nil, ErrUnprocessableEntity
	}

	if len(scheduleResponse.Schedule) == 0 {
		return nil, nil
	}
	event, err := e.models.Events.GetByScheduleId(input.ScheduleID)
	if err != nil {
		return nil, ErrRecordNotFound
	}

	err = e.models.Games.DeleteForEventId(event.ID)
	if err != nil {
		return nil, ErrUnprocessableEntity
	}

	var games []*models.Game
	for _, g := range scheduleResponse.Schedule {
		if g.GameType != "speedrun" {
			continue
		}

		var runnersList []string
		for _, runner := range g.Runners {
			runnersList = append(runnersList, runner.Name)
		}

		game, err := e.models.Games.Insert(&models.Game{
			ID:        "",
			StartDate: g.StartTime,
			EndDate:   g.EndTime,
			Duration:  g.Runtime,
			Name:      g.Name,
			Runners:   strings.Join(runnersList, ", "),
			GdqId:     g.Id,
			EventID:   event.ID,
		})
		if err != nil {
			return nil, ErrUnprocessableEntity
		}

		games = append(games, game)
	}

	return games, nil
}
