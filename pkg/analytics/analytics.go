package analytics

import (
	espnv3 "github.com/rthallisey/espn-go-api/pkg/api/v3"

	"fmt"
)

func Start(hc *espnv3.LeagueV3) error {
	members := hc.LeagueMembers()
	fmt.Println(members)

	teams, err := espnv3.NewTeam(hc)
	if err != nil {
		return err
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

	return nil
}
