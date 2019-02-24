package analytics

import (
	"github.com/espn-go-api/pkg/league"
)

func Start(hc league.League) error {
	hc.GetLeague()

	return nil
}
