package match

import (
	"context"
	"fmt"

	"github.com/mehmetolgundev/nba-project/domain/match/entities"
)

type MatchService struct {
	repository MatchRepository
}

func NewService(repository MatchRepository) MatchService {
	return MatchService{repository: repository}
}
func (s MatchService) GetMatches(ctx context.Context, currentTime int) []entities.MatchDTO {
	matches := s.repository.GetMatches(ctx)

	var result []entities.MatchDTO
	for _, match := range matches {
		matchDTO := entities.MatchDTO{
			HomeTeam: match.HomeTeam,
			AwayTeam: match.AwayTeam,
			DateTime: match.DateTime,
		}
		availableEvents := s.getAvailableEvents(match, currentTime)
		for _, event := range availableEvents {
			if event.Type == "Score" {
				matchDTO.MatchStatus = fmt.Sprintf("Continue : %s", getRealTime(currentTime))
				if event.Team == "Home" {
					matchDTO.HomeTeamAssists += fmt.Sprintf("%s <br>", event.AsistPlayer)
					matchDTO.HomeTeamScorer += fmt.Sprintf("%s (%d) <br>", event.Player, event.Point)
					matchDTO.HomeTeamScore += event.Point

				} else {
					matchDTO.AwayTeamAssists += fmt.Sprintf("%s <br>", event.AsistPlayer)
					matchDTO.AwayTeamScorer += fmt.Sprintf("%s (%d) <br>", event.Player, event.Point)
					matchDTO.AwayTeamScore += event.Point
				}
			} else if event.Type == "MatchStarted" {
				matchDTO.MatchStatus = fmt.Sprintf("Match started at %s <br> CurrentTime: %s", event.DateTime.Format("2006-01-02T15:04:05.000Z"), getRealTime(currentTime))

			} else if event.Type == "MatchFinished" {
				matchDTO.MatchStatus = fmt.Sprintf("Match finished at %s", event.DateTime.Format("2006-01-02T15:04:05.000Z"))
			}
		}
		result = append(result, matchDTO)
	}
	return result
}
func (s MatchService) getAvailableEvents(match entities.Match, currentTime int) []entities.Event {
	var events []entities.Event
	for _, event := range match.Events {
		if event.Time <= int64(currentTime) {
			events = append(events, event)
		}
	}
	return events
}
func getRealTime(currentTime int) string {
	return fmt.Sprintf("%d min", currentTime/5)
}
