package v3

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rthallisey/espn-go-api/pkg/api/v3/generated"
)

type LeagueV3 struct {
	Data generated.FFL

	LeagueID     int
	SeasonID     int
	SeasonLength int
	Params       string
	ESPNUrl      string
	EspnS2       string
	Swid         string
}

func NewLeague(l *LeagueV3) error {
	url := leagueUrl(l)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	json.Unmarshal(bodyText, &l.Data)

	return nil
}

func leagueUrl(l *LeagueV3) string {
	return fmt.Sprintf("%s/%d/segments/0/leagues/%d?%s", l.ESPNUrl, l.SeasonID, l.LeagueID, l.Params)
}

func (l LeagueV3) LeagueMembers() ([]string, error) {
	var s []string
	for _, i := range l.Data.Members {
		s = append(s, fmt.Sprintf("%s %s", i.FirstName, i.LastName))
	}

	return s, nil
}
