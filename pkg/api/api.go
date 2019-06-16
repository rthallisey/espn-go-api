package api

type League interface {
	// Get League members
	LeagueMembers() map[string]string
	LeagueMembersAbbrev() map[string]string

	// Public or Private league
	IsPublic() bool
}

type Team interface {
	AllRosters() map[string][]string

	// Get a team's roster
	Roster(string) ([]string, error)
}x

type Schedule interface {
	// Get teamID on the schedule
	TeamID() ([]string, error)
}
