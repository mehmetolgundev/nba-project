package entities

import "time"

type Match struct {
	Id       string    `bson:"_id,omitempty"`
	HomeTeam string    `json:"home_team"`
	AwayTeam string    `json:"away_team"`
	DateTime time.Time `json:"match_date" bson:"MatchDate"`
	Events   []Event
}
type Event struct {
	EventId     string    `bson:"EventId,omitempty"`
	Type        string    `json:"type"`
	Team        string    `json:"team"`
	Player      string    `json:"player"`
	AsistPlayer string    `json:"asistPlayer"`
	Point       int64     `json:"point"`
	Time        int64     `json:"time"`
	DateTime    time.Time `json:"date_time"`
}
