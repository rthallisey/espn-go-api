package analytics

import (
	//"github.com/rthallisey/espn-go-api/pkg/api"
	espnv3 "github.com/rthallisey/espn-go-api/pkg/api/v3"
	"github.com/rthallisey/espn-go-api/pkg/league"

	"fmt"
)

func Start(hc league.League) error {
	ffl, err := hc.GetLeagueData()
	if err != nil {
		return err
	}

	// API V3
	l := espnv3.LeagueV3{Data: ffl}
	members, err := l.LeagueMembers()
	if err != nil {
		return err
	}
	fmt.Println(members)
	// api.Team.TeamData()

	fmt.Println(ffl.DraftDetail.Drafted)
	return nil
}
