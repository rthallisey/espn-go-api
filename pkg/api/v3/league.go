package v3

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rthallisey/espn-go-api/pkg/api/v3/generated"
)

type LeagueV3 struct {
	Data *generated.League

	ScoreID int
}

func NewLeague(scoringPeriodID int, leagueID int, seasonID int, params string, espnUrl string) (LeagueV3, error) {
	url := fmt.Sprintf("%s/%d/segments/0/leagues/%d?%s", espnUrl, seasonID, leagueID, params)

	if scoringPeriodID > 0 {
		url = fmt.Sprintf("%s&scoringPeriodId=%d", url, scoringPeriodID)
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LeagueV3{}, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return LeagueV3{}, err
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return LeagueV3{}, err
	}

	d := &generated.League{}
	json.Unmarshal(bodyText, d)

	return LeagueV3{Data: d, ScoreID: scoringPeriodID}, nil
}

func (l LeagueV3) LeagueMembers() map[string]string {
	members := make(map[string]string)
	for _, i := range l.Data.Members {
		members[i.ID] = fmt.Sprintf("%s %s", i.FirstName, i.LastName)
	}

	return members
}

// Return IDs pairs with team abbreviations.
// Abbreviations are more human readable.
func (l LeagueV3) LeagueMembersAbbrev() map[string]string {
	members := make(map[string]string)

	for _, i := range l.Data.Members {
		for _, j := range l.Data.Teams {
			if j.PrimaryOwner == i.ID {
				members[j.Abbrev] = i.ID
			}
		}
	}

	return members
}

// Return IDs pairs with team IDs.
func (l LeagueV3) LeagueMembersID() map[string]int {
	members := make(map[string]int)

	for _, i := range l.Data.Teams {
		members[i.PrimaryOwner] = int(i.ID)
	}

	return members
}

func (l LeagueV3) IsPublic() bool {
	return l.Data.Status.IsViewable
}
