package data

import "time"

type EventDatum struct {
	ID                   string
	Timestamp            time.Time
	Donations            float64
	DonationsPerMinute   float64 `db:"donations_per_minute"`
	Donors               int64
	Tweets               int64
	TweetsPerMinute      int64 `db:"tweets_per_minute"`
	TwitchChats          int64 `db:"twitch_chats"`
	TwitchChatsPerMinute int64 `db:"twitch_chats_per_minute"`
	Viewers              int64
	EventID              string `db:"event_id"`
}
