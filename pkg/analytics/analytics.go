package analytics

import (
	espnv3 "github.com/rthallisey/espn-go-api/pkg/api/v3"

	"fmt"
)

func Start(hc espnv3.LeagueV3, weekly []espnv3.LeagueV3) error {

	members := hc.LeagueMembers()

	teams := espnv3.NewTeam(hc, weekly)
	schedules := espnv3.NewSchedule(hc)

	mvp, err := FindEveryTeamMVP(schedules, teams, hc)
	if err != nil {
		return err
	}

	for id, t := range mvp {
		fmt.Printf("Team %s MVPs\n", members[id])
		count := 0
		mvpSum := 0
		for p, total := range t {
			count += 1
			mvpSum += total
			fmt.Printf("%s: %d\n", p, total)
		}

		record, err := teams.TeamRecord(id)
		if err != nil {
			return err
		}

		fmt.Printf("Total MVPs: %d\n", count)
		fmt.Printf("Total wins: %v\n", record.Wins)
		avgMVP := float64(mvpSum) / float64(record.Wins)
		fmt.Printf("Average MVPs per win: %v\n", avgMVP)
		fmt.Printf("MVPs Per win: %v\n", float64(count)/float64(record.Wins))
		fmt.Printf("Normalized MVPs Per win: %v\n\n", float64(count)/float64(record.Wins)*float64(record.Wins))

	}

	for _, pos := range teams.PositionList() {
		p, _ := teams.PositionStringToID(pos)
		posAVG := teams.AvgPosScore(p)
		fmt.Printf("%s average points: %v\n", pos, posAVG)
	}
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

func FindEveryTeamMVP(schedule espnv3.Schedule, teams espnv3.Team, l espnv3.LeagueV3) (map[string]map[string]int, error) {
	membersID := l.LeagueMembersID()

	// mvp['team uuid']['player']
	mvp := map[string]map[string]int{}

	posAVG := map[int]float64{}
	for _, pos := range teams.PositionList() {
		p, _ := teams.PositionStringToID(pos)
		posAVG[p] = teams.AvgPosScore(p)
	}

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
					return nil, err
				}

				for player, pts := range playerWeekPoints {
					if pts.Score > diff && pts.Score > posAVG[pts.DefaultPositionID] {
						mvp[uuid][player] += 1
					}
				}
			}
		}
	}

	return mvp, nil
}
