package data

import (
	"context"
	"github.com/jmoiron/sqlx"
	"time"
)

type Game struct {
	ID        string
	StartDate time.Time `db:"start_date"`
	EndDate   time.Time `db:"end_date"`
	Duration  string
	Name      string
	Runners   string
	GdqId     int
	EventID   string `db:"event_id"`
}

type GameModel struct {
	db *sqlx.DB
}

func (m *GameModel) Insert(game *Game) (*Game, error) {
	stmt := `
		INSERT INTO games (start_date, end_date, duration, name, runners,gdq_id, event_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, start_date, end_date, duration, name, runners, event_id;
	`

	args := []any{
		game.StartDate,
		game.EndDate,
		game.Duration,
		game.Name,
		game.Runners,
		game.GdqId,
		game.EventID,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.db.GetContext(ctx, game, stmt, args...)

	return game, err
}

func (m *GameModel) GetAllForActiveEvent() ([]*Game, error) {
	stmt := `
		SELECT g.id, g.start_date, g.end_date, duration, name, runners, event_id
		FROM games g
		INNER JOIN events e ON e.id = g.event_id
		WHERE e.active_event = TRUE;
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var games []*Game
	err := m.db.SelectContext(ctx, &games, stmt)

	return games, err
}

func (m *GameModel) DeleteForEventId(id string) error {
	stmt := `DELETE FROM games WHERE event_id = $1;`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.db.ExecContext(ctx, stmt, id)

	return err
}
