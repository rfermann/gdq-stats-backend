package data

import "time"

type Game struct {
	ID        string
	StartDate time.Time `db:"start_date"`
	EndDate   time.Time `db:"end_date"`
	Duration  string
	Name      string
	Runner    string
	EventID   string `db:"event_id"`
}
