package teaminsights

type HomeAwayData struct {
	Home  float64 `json:"home"`
	Away  float64 `json:"away"`
	Total float64 `json:"total"`
}

// INFO: Main struct
// TODO: FInish me
type HomeAwayMetrics struct {
	Fixtures            *HomeAwayData            `json:"gamesPlayed"`
	Wins                *HomeAwayData            `json:"wins"`
	Draws               *HomeAwayData            `json:"draws"`
	Loses               *HomeAwayData            `json:"loses"`
	GoalsForTotal       *HomeAwayData            `json:"goalsForTotal"`
	GoalsForAverage     *HomeAwayData            `json:"goalsForAverage"`
	GoalsForMinute      map[string]*HomeAwayData `json:"goalsForMinute"` // FIXME: Create separate struct for this
	GoalsAgainstTotal   *HomeAwayData            `json:"goalsAgainstTotal"`
	GoalsAgainstAverage *HomeAwayData            `json:"goalsAgainstAverage"`
	GoalsAgainstMinute  map[string]*HomeAwayData `json:"goalsAgainstMinute"` // FIXME: Create separate struct for this
	CleanSheets         *HomeAwayData            `json:"cleanSheets"`
	FailedToScore       *HomeAwayData            `json:"failedToScore"`
}

func getPointsPerGame(wins, draws, games int) float32 {
	// INFO: Formula (wins * 3) + (draws * 1) / games
	winsPoints := wins * 3
	drawsPoints := draws * 1

	return (float32(winsPoints) + float32(drawsPoints)) / float32(games)
}
