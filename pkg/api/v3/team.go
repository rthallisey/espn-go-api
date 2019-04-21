package v3

import (
	"github.com/rthallisey/espn-go-api/pkg/api/v3/generated"
)

type Team struct {
	generated.Teams
}

func NewTeam(l *LeagueV3) (*Team, error) {
	return &Team{generated.Teams{l.Data.Teams}}, nil
}

func (t *Team) TeamData() ([]string, error) {
	var a []string
	return a, nil
}
