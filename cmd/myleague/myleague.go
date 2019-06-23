package main

import (
	"github.com/rthallisey/espn-go-api/pkg/analytics"
	espnv3 "github.com/rthallisey/espn-go-api/pkg/api/v3"
	"log"

	"fmt"
)

type league struct {
	scoringPeriodID int
	leagueID        int
	seasonID        int
	seasonLength    int

	// URL
	params  string
	espnUrl string

	// Auth
	espnS2 string
	swid   string
}

func main() {
	// 8 team league := 432343
	// 12 team league := 431183
	seasonLength := 15

	l := league{
		espnUrl:         "http://fantasy.espn.com/apis/v3/games/ffl/seasons",
		params:          "view=mBoxscore&view=mMatchupScore&view=mSchedule&view=mScoreboard&view=mSettings&view=mStatus&view=mTeam&view=mRoster&view=modular&view=mNav",
		leagueID:        431183,
		seasonID:        2018,
		seasonLength:    seasonLength,
		espnS2:          "",
		swid:            "",
		scoringPeriodID: 0,
	}

	hc, err := espnv3.NewLeague(l.scoringPeriodID, l.leagueID, l.seasonID, l.params, l.espnUrl)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("League is public:", hc.IsPublic())

	members := hc.LeagueMembers()
	if len(members) > 0 {
		fmt.Println("League members:", members)
	}

	weekly, err := getWeeklyData(l)
	if err != nil {
		log.Fatal(err)
	}

	analytics.Start(hc, weekly)
}

func getWeeklyData(l league) ([]espnv3.LeagueV3, error) {
	weekly := []espnv3.LeagueV3{}

	// seasonLen +2 to cover 2 playoff rounds.  This may not always work
	for week := 1; week <= l.seasonLength+2; week++ {
		fmt.Printf("Gathering week %d data...\n", week)

		hc, err := espnv3.NewLeague(week, l.leagueID, l.seasonID, l.params, l.espnUrl)
		if err != nil {
			return nil, err
		}

		weekly = append(weekly, hc)
	}
	return weekly, nil
}
