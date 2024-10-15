package entities

import (
	"time"

	"github.com/mehmetolgundev/nba-project/domain/match/constants"
)

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

func (e Event) IsScoreEvent() bool {
	return e.Type == constants.EventType_Score
}
func (e Event) IsMatchStartedEvent() bool {
	return e.Type == constants.EventType_MatchStarted
}
func (e Event) IsMatchStartedFinished() bool {
	return e.Type == constants.EventType_MatchFinished
}
