package v3

import (
	"errors"
	"fmt"

	"github.com/rthallisey/espn-go-api/pkg/api/v3/generated"
)

type Team struct {
	Generated generated.Teams
}

func NewTeam(l *LeagueV3) (*Team, error) {
	return &Team{Generated: generated.Teams{l.Data.Teams}}, nil
}

func (t *Team) Roster(id string) ([]string, error) {
	var a []string
	for _, r := range t.Generated.Teams {
		if id == r.PrimaryOwner {
			var a []string
			for _, j := range r.Roster.Entries {
				a = append(a, j.PlayerPoolEntry.Player.FullName)
			}
			return a, nil
		}
	}
	return a, errors.New(fmt.Sprintf("The teams with ID %s is not in the league", id))
}

func (t *Team) AllRosters() map[string][]string {
	roster := make(map[string][]string)

	for _, r := range t.Generated.Teams {
		var a []string
		for _, j := range r.Roster.Entries {
			a = append(a, j.PlayerPoolEntry.Player.FullName)
		}
		roster[r.PrimaryOwner] = a
	}

	return roster
}
