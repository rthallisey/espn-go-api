package league

import (
	espnv3 "github.com/rthallisey/espn-go-api/pkg/api/v3"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (l League) GetLeagueData() (espnv3.FFL, error) {
	var data espnv3.FFL
	url := l.leagueUrl()

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return data, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return data, err
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return data, err
	}
	json.Unmarshal(bodyText, &data)

	return data, nil
}

func (l League) leagueUrl() string {
	return fmt.Sprintf("%s/%d/segments/0/leagues/%d?%s", l.ESPNUrl, l.SeasonID, l.LeagueID, l.Params)
}
