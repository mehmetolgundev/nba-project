package entities

import "time"

type MatchDTO struct {
	HomeTeam        string    `json:"home_team"`
	AwayTeam        string    `json:"away_team"`
	DateTime        time.Time `json:"match_date"`
	Time            int64     `json:"time"`
	HomeTeamScore   int64     `json:"home_team_score"`
	AwayTeamScore   int64     `json:"away_team_score"`
	HomeTeamAssists string    `json:"home_team_assists"`
	AwayTeamAssists string    `json:"away_team_assists"`
	HomeTeamScorer  string    `json:"home_team_scorer"`
	AwayTeamScorer  string    `json:"away_team_scorer"`
	MatchStatus     string    `json:"match_status"`
}
