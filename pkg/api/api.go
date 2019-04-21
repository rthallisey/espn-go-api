package api

type League interface {
	LeagueMembers() ([]string, error)
}

type Team interface {
	TeamRosters() map[string][]string
}
