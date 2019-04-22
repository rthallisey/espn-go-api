package api

type League interface {
	LeagueMembers() map[string]string
}

type Team interface {
	Roster(string) ([]string, error)
	AllRosters() map[string][]string
}
