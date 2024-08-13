package insights

import (
	"insights-pulse/src/models/responses"
	u "insights-pulse/src/utils"
)

type StatDetail struct {
	Sum float64
	Num float64
	Avg float64
}

type StatsMetadata struct {
	TeamId   string `json:"teamId"`
	Season   string `json:"season"`
	LeagueId string `json:"leagueId"`
}

type AvgMetricsPerGame struct {
	TeamId           string      `json:"teamId"`
	Season           string      `json:"season"`
	LeagueId         string      `json:"league"`
	ShotsOnGoal      *StatDetail `json:"Shots on Goal"`
	ShotsOffGoal     *StatDetail `json:"Shots off Goal"`
	TotalShots       *StatDetail `json:"Total Shots"`
	BlockedShots     *StatDetail `json:"Blocked Shots"`
	ShotsInsideBox   *StatDetail `json:"Shots insidebox"`
	ShotsOutsideBox  *StatDetail `json:"Shots outsidebox"`
	Fouls            *StatDetail `json:"Fouls"`
	CornerKicks      *StatDetail `json:"Corner Kicks"`
	Offsides         *StatDetail `json:"Offsides"`
	BallPossession   *StatDetail `json:"Ball Possession"`
	YellowCards      *StatDetail `json:"Yellow Cards"`
	RedCards         *StatDetail `json:"Red Cards"`
	GoalkeeperSaves  *StatDetail `json:"Goalkeeper Saves"`
	TotalPasses      *StatDetail `json:"Total passes"`
	PassesAccurate   *StatDetail `json:"Passes accurate"`
	PassesPercentage *StatDetail `json:"Passes %"`
	ExpectedGoals    *StatDetail `json:"expected_goals"`
}

func NewAverageInsightsPerGame(teamId, leagueId, season string) *AvgMetricsPerGame {

	return &AvgMetricsPerGame{
		TeamId:           teamId,
		LeagueId:         leagueId,
		Season:           season,
		ShotsOnGoal:      &StatDetail{},
		ShotsOffGoal:     &StatDetail{},
		TotalShots:       &StatDetail{},
		BlockedShots:     &StatDetail{},
		ShotsInsideBox:   &StatDetail{},
		ShotsOutsideBox:  &StatDetail{},
		Fouls:            &StatDetail{},
		CornerKicks:      &StatDetail{},
		Offsides:         &StatDetail{},
		BallPossession:   &StatDetail{},
		YellowCards:      &StatDetail{},
		RedCards:         &StatDetail{},
		GoalkeeperSaves:  &StatDetail{},
		TotalPasses:      &StatDetail{},
		PassesAccurate:   &StatDetail{},
		PassesPercentage: &StatDetail{},
		ExpectedGoals:    &StatDetail{},
	}

}

// INFO: Insights New
type MatchStatsDetail struct {
	Sum float64
	Num float64
	Avg float64
}

type StatsMetaData struct {
	TeamId   string `json:"teamId"`
	Season   string `json:"season"`
	LeagueId string `json:"league"`
}

type MatchMetrics struct {
	ShotsOnGoal      *MatchStatsDetail `json:"Shots on Goal"`
	ShotsOffGoal     *MatchStatsDetail `json:"Shots off Goal"`
	TotalShots       *MatchStatsDetail `json:"Total Shots"`
	BlockedShots     *MatchStatsDetail `json:"Blocked Shots"`
	ShotsInsideBox   *MatchStatsDetail `json:"Shots insidebox"`
	ShotsOutsideBox  *MatchStatsDetail `json:"Shots outsidebox"`
	Fouls            *MatchStatsDetail `json:"Fouls"`
	CornerKicks      *MatchStatsDetail `json:"Corner Kicks"`
	Offsides         *MatchStatsDetail `json:"Offsides"`
	BallPossession   *MatchStatsDetail `json:"Ball Possession"`
	YellowCards      *MatchStatsDetail `json:"Yellow Cards"`
	RedCards         *MatchStatsDetail `json:"Red Cards"`
	GoalkeeperSaves  *MatchStatsDetail `json:"Goalkeeper Saves"`
	TotalPasses      *MatchStatsDetail `json:"Total passes"`
	PassesAccurate   *MatchStatsDetail `json:"Passes accurate"`
	PassesPercentage *MatchStatsDetail `json:"Passes %"`
	ExpectedGoals    *MatchStatsDetail `json:"expected_goals"`
}

func NewMatchMetrics() *MatchMetrics {
	return &MatchMetrics{
		ShotsOnGoal:      &MatchStatsDetail{},
		ShotsOffGoal:     &MatchStatsDetail{},
		TotalShots:       &MatchStatsDetail{},
		BlockedShots:     &MatchStatsDetail{},
		ShotsInsideBox:   &MatchStatsDetail{},
		ShotsOutsideBox:  &MatchStatsDetail{},
		Fouls:            &MatchStatsDetail{},
		CornerKicks:      &MatchStatsDetail{},
		Offsides:         &MatchStatsDetail{},
		BallPossession:   &MatchStatsDetail{},
		YellowCards:      &MatchStatsDetail{},
		RedCards:         &MatchStatsDetail{},
		GoalkeeperSaves:  &MatchStatsDetail{},
		TotalPasses:      &MatchStatsDetail{},
		PassesAccurate:   &MatchStatsDetail{},
		PassesPercentage: &MatchStatsDetail{},
		ExpectedGoals:    &MatchStatsDetail{},
	}
}

// FIXME: Merge this with the AverageInsightsPerGame struct
type TeamMetricsPerSeasonTotal struct {
	CleanSheets    string `json:"clean_sheets"`
	GoalDifference string `json:"goal_difference"`
}

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
