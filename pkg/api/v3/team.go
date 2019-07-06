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

const WideReceiver = "WR"
const RunningBack = "RB"
const QuarterBack = "QB"
const TightEnd = "TE"
const Kicker = "K"
const Defense = "DEF"

func (t Team) PositionList() []string {
	return []string{WideReceiver, RunningBack, QuarterBack, TightEnd, Kicker, Defense}
}

func (t Team) PositionStringToID(pos string) (int, error) {
	switch pos {
	case RunningBack:
		return 2, nil
	case WideReceiver:
		return 3, nil
	case QuarterBack:
		return 1, nil
	case TightEnd:
		return 4, nil
	case Kicker:
		return 5, nil
	case Defense:
		return 16, nil
	default:
		return 0, errors.New(fmt.Sprintf("Position %s does not exist. Valid positions: %v", pos, t.PositionList()))
	}
}

// Average score for position of players who were rostered for a week
func (t Team) AvgPosScore(posID int) float64 {
	p := map[string]int{}
	var count float64
	var avg float64

	for _, team := range t.Generated.Teams {
		s, err := t.PlayerWeeklyScore(team.PrimaryOwner)
		if err != nil {
			return 0
		}
		// Loop through each week
		for _, week := range s {
			for player, pts := range week {
				// Verify the player's position
				if pts.DefaultPositionID == posID {
					// Only tally a player's avg that hasn't been seen yet
					if _, ok := p[player]; !ok {
						p[player] = 0
						count += 1
						avg += pts.seasonAverage
					}
				}
			}
		}
	}
	return avg / count
}

type playerPoints struct {
	Score float64

	// The default place where a player is played e.g. WR
	DefaultPositionID int

	projectedPoints float64
	seasonAverage   float64
}

type record struct {
	Wins   int
	Losses int
}

// Get a team's record
func (t Team) TeamRecord(id string) (record, error) {
	teamRecord := record{}
	for _, r := range t.Generated.Teams {
		if id == r.PrimaryOwner {
			teamRecord.Losses = int(r.Record.Overall.Losses)
			teamRecord.Wins = int(r.Record.Overall.Wins)
			return teamRecord, nil
		}
	}
	return record{}, errors.New(fmt.Sprintf("The teams with ID %s is not in the league", id))
}

// Players and their scores for a specific week
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

				// Register a players default position ID
				trackPoints.DefaultPositionID = int(player.PlayerPoolEntry.Player.DefaultPositionID)

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

// Player scores for each week from a specific team
func (t Team) PlayerWeeklyScore(id string) ([]map[string]playerPoints, error) {
	weekly := []map[string]playerPoints{}

	// Loop through each week
	for w, _ := range t.GeneratedList {
		weekScore, err := t.PlayerWeekScore(id, w)
		if err != nil {
			return nil, err
		}
		weekly = append(weekly, weekScore)
	}
	return weekly, nil
}
