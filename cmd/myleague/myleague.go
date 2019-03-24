package main

import (
	"github.com/rthallisey/espn-go-api/pkg/analytics"
	espnv3 "github.com/rthallisey/espn-go-api/pkg/api/v3"

	"log"
)

func main() {
	hc := &espnv3.LeagueV3{
		ESPNUrl:      "http://fantasy.espn.com/apis/v3/games/ffl/seasons",
		Params:       "view=mBoxscore&view=mMatchupScore&view=mSchedule&view=mScoreboard&view=mSettings&view=mStatus&view=mTeam&view=mRoster&view=modular&view=mNav",
		LeagueID:     432343,
		SeasonID:     2018,
		SeasonLength: 16,
		EspnS2:       "",
		Swid:         "",
	}

	err := espnv3.NewLeague(hc)
	if err != nil {
		log.Fatal(err)
	}
	analytics.Start(hc)
}
