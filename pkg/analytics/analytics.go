package analytics

import (
	espnv3 "github.com/rthallisey/espn-go-api/pkg/api/v3"

	"fmt"
)

func Start(hc *espnv3.LeagueV3) error {
	members, err := hc.LeagueMembers()
	if err != nil {
		return err
	}
	fmt.Println(members)

	teams, err := espnv3.NewTeam(hc)
	if err != nil {
		return err
	}
	teams.TeamData()

	fmt.Println(hc.Data.DraftDetail.Drafted)
	return nil
}
