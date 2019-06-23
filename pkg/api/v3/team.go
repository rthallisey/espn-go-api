package v3

import (
	"errors"
	"fmt"

	"github.com/rthallisey/espn-go-api/pkg/api/v3/generated"
)

type Team struct {
	GeneratedList []generated.Teams

	Generated generated.Teams
}

func NewTeam(l LeagueV3, leagueList []LeagueV3) Team {
	teamList := []generated.Teams{}
	for _, league := range leagueList {
		teamList = append(teamList, generated.Teams{league.Data.Teams})
	}

	return Team{Generated: generated.Teams{l.Data.Teams}, GeneratedList: teamList}
}

func (t Team) Roster(id string) ([]string, error) {
	var a []string
	for _, r := range t.Generated.Teams {
		if id == r.PrimaryOwner {
			var a []string
			for _, j := range r.Roster.Entries {
				a = append(a, j.PlayerPoolEntry.Player.FullName)
			}
			return a, nil
		}
	}
	return a, errors.New(fmt.Sprintf("The teams with ID %s is not in the league", id))
}

func (t Team) AllRosters() map[string][]string {
	roster := make(map[string][]string)

	for _, r := range t.Generated.Teams {
		var a []string
		for _, j := range r.Roster.Entries {
			a = append(a, j.PlayerPoolEntry.Player.FullName)
		}
		roster[r.PrimaryOwner] = a
	}

	return roster
}

// TODO: add a Roster object to re use roster check code

type playerPoints struct {
	projectedPoints float64
	Score           float64
	seasonAverage   float64
}

// Players and their scores for a wekk
func (t Team) PlayerWeekScore(id string, week int) (map[string]playerPoints, error) {
	points := map[string]playerPoints{}

	// Loop through each team
	for _, team := range t.GeneratedList[week].Teams {

		// Match team owner with their team
		//
		// **This has no check to verify a team was found**
		// errors.New(fmt.Sprintf("The teams with ID %s is not in the league", id))
		if id == team.PrimaryOwner {

			// Loop through players on their roster
			for _, player := range team.Roster.Entries {

				trackPoints := playerPoints{}

				// Loop through a player's stats and find pts scored
				//   Projected points for a week:
				//       appliedTotal
				//       statSourceId = 1
				//       statSplitId  = 1
				//   Scored points for a week:
				//       appliedTotal
				//       statSourceId = 0
				//       statSplitId  = 1
				//   Average points for the season:
				//       appliedAverage
				//       statSourceId = 0
				//       statSplitId  = 0
				for _, statSource := range player.PlayerPoolEntry.Player.Stats {
					if statSource.StatSourceID == 1 && statSource.StatSplitTypeID == 1 && int(statSource.ScoringPeriodID) == week {
						trackPoints.projectedPoints = statSource.AppliedTotal

					} else if statSource.StatSourceID == 0 && statSource.StatSplitTypeID == 1 && int(statSource.ScoringPeriodID) == week {
						trackPoints.Score = statSource.AppliedTotal
					} else if statSource.StatSourceID == 0 && statSource.StatSplitTypeID == 0 {
						trackPoints.seasonAverage = statSource.AppliedAverage
					}
				}
				points[player.PlayerPoolEntry.Player.FullName] = trackPoints
			}
			return points, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("The teams with ID %s is not in the league", id))
}

func (t Team) PlayerWeeklyScore(id string) ([]map[string]playerPoints, error) {
	weekly := []map[string]playerPoints{}

	// Loop through each week
	for w, _ := range t.GeneratedList {
		// Weeks aren't 0'd
		w += 1
		weekScore, err := t.PlayerWeekScore(id, w)
		if err != nil {
			return nil, err
		}
		weekly = append(weekly, weekScore)
	}
	return weekly, nil
}
