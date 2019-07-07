package v3

type scoreDiff struct {
	difference      float64
	scoringPeriodID int
}

type playerPoints struct {
	Score float64

	// The default place where a player is played e.g. WR
	DefaultPositionID int

	// Player was on the bench
	Bench bool

	// Drafted, Added, Traded for
	AcquisitionType string

	projectedPoints float64
	seasonAverage   float64
}

type record struct {
	Wins   int
	Losses int
}
