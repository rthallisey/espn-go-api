package main

import (
	"github.com/rthallisey/espn-go-api/pkg/analytics"
	"github.com/rthallisey/espn-go-api/pkg/league"
)

// 361793,
func main() {
	hc := league.League{
		ESPNUrl:      "http://fantasy.espn.com/apis/v3/games/ffl/seasons",
		Params:       "view=mBoxscore&view=mMatchupScore&view=mSchedule&view=mScoreboard&view=mSettings&view=mStatus&view=mTeam&view=mRoster&view=modular&view=mNav",
		LeagueID:     432343,
		SeasonID:     2018,
		SeasonLength: 16,
		EspnS2:       "AEBUBSDyP6%2FXzKzc88QaBaBq7iJUzklumOk5JLbx0DKGcLvTHkUJGTXmmqOOEoxPNjvvEszYuAHflwr7lPURWhtNhLgvWID%2BHhGc7u%2FdWAQM0yRsTR0resXumgJS1tMSEn4UdFy%2FE32XGAWcg4A1PZHYGkK7FFKxEP1HnBxUgYHNFWbPo7cJGPIBE7oDadz0CLlLHrRf9E9oNKzmL5FLny57unKhynNnr9gVR%2FGmZEGCYFbX%2BRTmJPIHIf%2FhnwQbZzY4mNG3%2F7jk9h3cq6ZW7kuL",
		Swid:         "{7A703CA3-FB06-4A12-A868-46266E40A33D}",
	}

	analytics.Start(hc)
}
