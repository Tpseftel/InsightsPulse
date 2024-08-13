package insightsgenerator

import (
	"insights-pulse/src/models/responses"
	u "insights-pulse/src/utils"
)

type TeamMetricsPerSeasonTotal struct {
	CleanSheets    string `json:"clean_sheets"`
	GoalDifference string `json:"goal_difference"`
}

// FIXME: Merge this with the AverageInsightsPerGame struct
type TeamAveragePerformanceMetricsPerGame struct {
	TeamID               string  `json:"team_id"`
	LeagueID             string  `json:"league_id"`
	Season               string  `json:"season"`
	GoalsScored          string  `json:"goals_scored"`
	GoalsConceded        string  `json:"goals_conceded"`
	PointsPerGame        string  `json:"points_per_game"`
	PossessionPercent    string  `json:"possession_percent"`
	PassAccuracy         string  `json:"pass_accuracy"`
	ShotsOnTarget        string  `json:"shots_on_target"`
	TacklesInterceptions string  `json:"tackles_interceptions"`
	ExpectedGoals        float64 `json:"expected_goals"`
}

// Team Average Metrics Per Game
func GenerateMeanPerformance(t *responses.TeamStatsResponse) *TeamAveragePerformanceMetricsPerGame {

	return &TeamAveragePerformanceMetricsPerGame{
		TeamID:        u.ConvToString(t.Response.Team.ID),
		LeagueID:      u.ConvToString(t.Response.League.ID),
		Season:        u.ConvToString(t.Response.League.Season),
		GoalsScored:   u.ConvToString(t.Response.Goals.For.Average.Total),
		GoalsConceded: u.ConvToString(t.Response.Goals.Against.Average.Total),
		PointsPerGame: u.ConvToString(
			getPointsPerGame(
				t.Response.Fixtures.Wins.Total,
				t.Response.Fixtures.Draws.Total,
				t.Response.Fixtures.Played.Total,
			)),
		// TODO: Add PossessionPercent,
		// TODO: Add PassAccuracy,
		// TODO: Add ShotsOnTarget,
		// TODO: Add CleanSheets,
		// TODO: Add TacklesInterceptions,
		// TODO: Add ExpectedGoalsFor,
		// TODO: Add ExpectedGoalsAgainst,
	}
}

func getPointsPerGame(wins, draws, games int) float32 {
	// INFO: Formula (wins * 3) + (draws * 1) / games
	winsPoints := wins * 3
	drawsPoints := draws * 1

	return (float32(winsPoints) + float32(drawsPoints)) / float32(games)
}
