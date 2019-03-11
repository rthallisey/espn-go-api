package v3

import (
	"fmt"
)

type LeagueV3 struct {
	Data FFL
}

func (l LeagueV3) LeagueMembers() ([]string, error) {
	var s []string
	for _, i := range l.Data.Members {
		s = append(s, fmt.Sprintf("%s %s", i.FirstName, i.LastName))
	}

	return s, nil
}
