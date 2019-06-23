package analytics

import (
	espnv3 "github.com/rthallisey/espn-go-api/pkg/api/v3"

	"fmt"
)

func Start(hc espnv3.LeagueV3, weekly []espnv3.LeagueV3) error {

	//members := hc.LeagueMembers()

	teams := espnv3.NewTeam(hc, weekly)
	schedules := espnv3.NewSchedule(hc)

	FindTeamMVP(schedules, teams, hc)

	// rosters := teams.AllRosters()
	// for rosterID, _ := range rosters {
	// 	fmt.Println(members[rosterID])

	// 	r, err := teams.Roster(rosterID)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	} else {
	// 		fmt.Println(r)
	// 	}
	// }

	return nil
}

// Score difference and who lost
type scoreLoss struct {
	difference float64
	loser      int
}

func FindTeamMVP(schedule espnv3.Schedule, teams espnv3.Team, l espnv3.LeagueV3) error {
	membersID := l.LeagueMembersID()
	members := l.LeagueMembers()

	// mvp['team uuid']['player']
	mvp := map[string]map[string]int{}

	// Winner is the key for sl values
	for game, _ := range schedule.Generated.Schedule {
		//winner, diff, loser
		winner, diff, _ := schedule.GameWinLossScore(game)

		w := schedule.GameToWeek(game)
		for uuid, id := range membersID {
			if _, ok := mvp[uuid]; !ok {
				mvp[uuid] = map[string]int{}
			}
			if id == winner {
				//fmt.Printf("TEAM %v won over TEAM %v by %v points\n", id, loser, diff)
				playerWeekPoints, err := teams.PlayerWeekScore(uuid, w)
				if err != nil {
					return err
				}

				for player, pts := range playerWeekPoints {
					if pts.Score > diff {
						if _, ok := mvp[uuid][player]; ok {
							mvp[uuid][player] += 1
						} else {
							mvp[uuid][player] = 1
						}
					}
				}
			}
		}
	}

	for id, t := range mvp {
		fmt.Printf("Teams %s MVPs\n", members[id])
		for p, total := range t {
			fmt.Printf("%s: %d\n", p, total)
		}
	}
	return nil
}
