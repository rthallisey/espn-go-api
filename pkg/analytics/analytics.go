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
		exportAvgPosScoreData := map[string]float64{}
		for _, pos := range teams.PositionList() {
			p, _ := teams.PositionStringToID(pos)
			posAVG := teams.TeamAvgPosScore(team.PrimaryOwner, p)
			fmt.Printf("Team %s average points for position %s: %v\n", members[team.PrimaryOwner], pos, posAVG)
			exportAvgPosScoreData[pos] = posAVG
			MapStringFloatToJson("TeamPtsPerPosition/"+members[team.PrimaryOwner]+".json", exportAvgPosScoreData)
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

		// weeks are zeroed in the Team API.  ScoringPeriodId is not
		// zeroed from the Schedule API.

		data, err := teams.TeamMVP(int(team.ID), winScore[int(team.ID)])
		if err != nil {
			return err
		}
		mvp[team.PrimaryOwner] = data
	}

	exportCountData := map[string]int{}
	exportAvgMVPData := map[string]float64{}
	exportMVPPerWinData := map[string]float64{}
	exportNormMVPData := map[string]float64{}
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

		byWin := findMVPByWin(t)

		MapStringIntToJson("TeamMVPByWin/" + members[id] + ".json", byWin)

		for play, wins := range byWin {
			fmt.Printf("Team MVP: %v with %v wins\n", play, wins)
		}

		fmt.Printf("Number of players with MVPs: %d\n", count)
		exportCountData[members[id]] = count

		fmt.Printf("Total wins: %v\n", record.Wins)
		avgMVP := float64(mvpSum) / float64(record.Wins)

		fmt.Printf("Average MVPs per win: %v\n", avgMVP)
		exportAvgMVPData[members[id]] = avgMVP

		mvpsPerWin := float64(mvpSum) / float64(record.Wins)
		fmt.Printf("MVPs Per win: %v\n", mvpsPerWin)
		exportMVPPerWinData[members[id]] = mvpsPerWin

		normMvpsPerWin := float64(mvpSum)/float64(record.Wins)*float64(record.Wins) + float64(count) - float64(record.Wins)
		fmt.Printf("Normalized MVPs Per win: %v\n\n", normMvpsPerWin)
		exportNormMVPData[members[id]] = normMvpsPerWin
	}
	MapStringIntToJson("MVPCount.json", exportCountData)
	MapStringFloatToJson("AvgMVPs.json", exportAvgMVPData)
	MapStringFloatToJson("MVPsPerWin.json", exportMVPPerWinData)
	MapStringFloatToJson("NormalizedMVPsPerWin.json", exportNormMVPData)

	return nil
}

// Return list with MVP(s) that had the most wins
func findMVPByWin(byWin map[string]int) map[string]int {
	mvps := map[string]int{}

	// n * n :(
	max := 0
 	for _, winTotal := range byWin {
		if winTotal >= max {
			max = winTotal
		}
	}
 	for player, winTotal := range byWin {
		if winTotal == max {
			mvps[player] = winTotal
		}
	}
	return mvps
}

func benchPoints(teams espnv3.Team, hc espnv3.LeagueV3) {
	members := hc.LeagueMembers()

	exportData := map[string]float64{}

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
		exportData[members[team.PrimaryOwner]] = benchPts
	}
	MapStringFloatToJson("BencnhPoints.json", exportData)
}

func teamPlayerMostWins(teams espnv3.Team, hc espnv3.LeagueV3) {

}

func teamPlayerMostPoints(teams espnv3.Team, hc espnv3.LeagueV3) {
	members := hc.LeagueMembers()

	exportData := map[string]float64{}
	for _, team := range teams.Generated.Teams {
		highestScorer, highestScore := teams.TeamPlayerMostPoints(team.PrimaryOwner)

		fmt.Printf("Team %s highest scorer is %s with %v points\n", members[team.PrimaryOwner], highestScorer, highestScore)
		exportData[highestScorer+" ("+members[team.PrimaryOwner]+")"] = highestScore
	}
	MapStringFloatToJson("playerMostPoints.json", exportData)

}
