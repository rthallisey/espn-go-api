package analytics

import (
	espnv3 "github.com/rthallisey/espn-go-api/pkg/api/v3"

	"fmt"
)

func Start(hc espnv3.LeagueV3, weekly []espnv3.LeagueV3) error {

	members := hc.LeagueMembers()

	teams := espnv3.NewTeam(hc, weekly)
	schedules := espnv3.NewSchedule(hc)

	err := findEveryTeamMVP(schedules, teams, hc)
	if err != nil {
		return err
	}

	benchPoints(teams, hc)
	fmt.Printf("\n")

	for _, team := range teams.Generated.Teams {
		for _, pos := range teams.PositionList() {
			p, _ := teams.PositionStringToID(pos)
			posAVG := teams.TeamAvgPosScore(team.PrimaryOwner, p)
			fmt.Printf("Team %s average poinsts for position %s: %v\n", members[team.PrimaryOwner], pos, posAVG)
		}
		fmt.Printf("\n")
	}

	teamPlayerMostPoints(teams, hc)
	fmt.Printf("\n")

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

func findEveryTeamMVP(schedule espnv3.Schedule, teams espnv3.Team, l espnv3.LeagueV3) error {
	members := l.LeagueMembers()
	mvp := map[string]map[string]int{}

	for _, team := range teams.Generated.Teams {
		if _, ok := mvp[team.PrimaryOwner]; !ok {
			mvp[team.PrimaryOwner] = map[string]int{}
		}

		// winScore -> [teamID][]scoreDiff
		winScore := schedule.TeamGameWinScore(int(team.ID))

		data, err := teams.TeamMVP(int(team.ID), winScore[int(team.ID)])
		if err != nil {
			return err
		}
		mvp[team.PrimaryOwner] = data
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

		fmt.Printf("Number of players with MVPs: %d\n", count)
		fmt.Printf("Total wins: %v\n", record.Wins)
		avgMVP := float64(mvpSum) / float64(record.Wins)
		fmt.Printf("Average MVPs per win: %v\n", avgMVP)
		fmt.Printf("MVPs Per win: %v\n", float64(mvpSum)/float64(record.Wins))
		fmt.Printf("Normalized MVPs Per win: %v\n\n", float64(mvpSum)/float64(record.Wins)*float64(record.Wins)+float64(count)-float64(record.Wins))
	}
	return nil
}

func benchPoints(teams espnv3.Team, hc espnv3.LeagueV3) {
	members := hc.LeagueMembers()

	for _, team := range teams.Generated.Teams {
		benchPts := 0.0
		data, err := teams.PlayerWeeklyScore(team.PrimaryOwner)
		if err != nil {
			fmt.Println(err)
			return
		}
		// Loop through each week
		for _, week := range data {
			for _, pts := range week {
				if pts.Bench {
					benchPts += pts.Score
				}
			}
		}
		fmt.Printf("Team %s Bench Points: %v\n", members[team.PrimaryOwner], benchPts)
	}
}

func teamPlayerMostPoints(teams espnv3.Team, hc espnv3.LeagueV3) {
	members := hc.LeagueMembers()

	for _, team := range teams.Generated.Teams {
		highestScorer, highestScore := teams.TeamPlayerMostPoints(team.PrimaryOwner)

		fmt.Printf("Team %s highest scorer is %s with %v points\n", members[team.PrimaryOwner], highestScorer, highestScore)
	}

}
