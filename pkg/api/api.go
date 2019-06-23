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
}

type Schedule interface {
	// Get a team with match id's total score for each week
	TeamScore(int) ([]float64, error)

	TeamWinLoss(int) ([]string, error)
	TeamOpponent(int) ([]int, error)
}
