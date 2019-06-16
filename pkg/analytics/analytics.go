package analytics

import (
	espnv3 "github.com/rthallisey/espn-go-api/pkg/api/v3"

	"fmt"
)

func Start(hc espnv3.LeagueV3, weekly []espnv3.LeagueV3) error {

	members := hc.LeagueMembers()
	membersAbbrev := hc.LeagueMembersAbbrev()

	teams := espnv3.NewTeam(hc, weekly)
	// schedules := espnv3.NewSchedule(hc)

	weeklyPoints, err := teams.TeamWeeklyScore(membersAbbrev["GBL"])
	if err != nil {
		return err
	}

	for week, w := range weeklyPoints {
		// Weeks aren't 0'd
		fmt.Printf("Week %d\n", (week + 1))

		fmt.Println(w)
	}

	rosters := teams.AllRosters()
	for rosterID, _ := range rosters {
		fmt.Println(members[rosterID])

		r, err := teams.Roster(rosterID)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(r)
		}
	}

	// schedule, err := schedules.TeamID()
	// if err != nil {
	// 	return err
	// }
	// fmt.Println(schedule)

	return nil
}
