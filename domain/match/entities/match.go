package entities

import "time"

type Match struct {
	Id       string    `bson:"_id,omitempty"`
	HomeTeam string    `json:"home_team"`
	AwayTeam string    `json:"away_team"`
	DateTime time.Time `json:"match_date" bson:"MatchDate"`
	Events   []Event
}
