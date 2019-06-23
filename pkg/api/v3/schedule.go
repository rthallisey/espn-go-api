package v3

import (
	"errors"
	"math"

	"github.com/rthallisey/espn-go-api/pkg/api/v3/generated"
)

type Schedule struct {
	Generated generated.Schedule
}

func NewSchedule(l LeagueV3) Schedule {
	return Schedule{Generated: generated.Schedule{l.Data.Schedule}}
}

func (s Schedule) TeamScore(id int) ([]float64, error) {
	var score []float64
	found := false

	for _, sched := range s.Generated.Schedule {
		if int(sched.Home.TeamID) == id {
			found = true
			score = append(score, sched.Home.TotalPoints)
		} else if int(sched.Away.TeamID) == id {
			found = true
			score = append(score, sched.Away.TotalPoints)
		}
	}
	if !found {
		return nil, errors.New("Did not find team with matching Id.")
	}

	return score, nil
}

// Determine team's (id) opponent for each week by teamID
func (s Schedule) TeamOpponent(id int) ([]int, error) {
	var opponent []int
	found := false

	for _, sched := range s.Generated.Schedule {
		if int(sched.Home.TeamID) == id {
			found = true
			opponent = append(opponent, int(sched.Away.TeamID))
		} else if int(sched.Away.TeamID) == id {
			found = true
			opponent = append(opponent, int(sched.Home.TeamID))
		}
	}
	if !found {
		return nil, errors.New("Did not find team with matching Id.")
	}
	return opponent, nil
}

// Translate a game ID to week (MatchupPeriodID)
func (s Schedule) GameToWeek(game int) int {
	return int(s.Generated.Schedule[game].MatchupPeriodID)
}

// Determine who won, the score, and the losers score for a particular game
// teamID
// difference
// teamID
func (s Schedule) GameWinLossScore(game int) (int, float64, int) {
	winner := s.Generated.Schedule[game].Winner
	var winnerTeamID int
	var loserTeamID int
	var difference float64
	if winner == "AWAY" {
		difference = s.Generated.Schedule[game].Away.TotalPoints - s.Generated.Schedule[game].Home.TotalPoints
		winnerTeamID = int(s.Generated.Schedule[game].Away.TeamID)
		loserTeamID = int(s.Generated.Schedule[game].Home.TeamID)
	} else if winner == "HOME" {
		difference = s.Generated.Schedule[game].Home.TotalPoints - s.Generated.Schedule[game].Away.TotalPoints
		winnerTeamID = int(s.Generated.Schedule[game].Home.TeamID)
		loserTeamID = int(s.Generated.Schedule[game].Away.TeamID)
	}
	difference = math.Round(difference*100) / 100
	return winnerTeamID, difference, loserTeamID
}

// Determine if a team (id) won or loss for a particular week
func (s Schedule) TeamWinLoss(id int) ([]string, error) {
	var winloss []string
	found := false

	for _, sched := range s.Generated.Schedule {
		if int(sched.Home.TeamID) == id {
			found = true
			if sched.Winner == "HOME" {
				winloss = append(winloss, "win")
			} else if sched.Winner == "AWAY" {
				winloss = append(winloss, "loss")
			}
		} else if int(sched.Away.TeamID) == id {
			found = true
			if sched.Winner == "AWAY" {
				winloss = append(winloss, "win")
			} else if sched.Winner == "HOME" {
				winloss = append(winloss, "loss")
			}
		}
	}
	if !found {
		return nil, errors.New("Did not find team with matching Id.")
	}
	return winloss, nil
}
