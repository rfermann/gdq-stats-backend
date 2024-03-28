package data

import (
	"time"
)

type Event struct {
	ID             string
	Year           int64
	StartDate      time.Time `db:"start_date"`
	EndDate        time.Time `db:"end_date"`
	ActiveEvent    bool      `db:"active_event"`
	Viewers        int64
	Donations      float64
	Donors         int64
	GamesCompleted int64 `db:"games_completed"`
	TwitchChats    int64 `db:"twitch_chats"`
	Tweets         int64
	ScheduleID     int64  `db:"schedule_id"`
	EventTypeID    string `db:"event_type_id"`
}
