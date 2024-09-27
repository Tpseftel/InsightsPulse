package teamgenerator

import (
	"errors"
	"insights-pulse/src/models/insights/teaminsights"
	"insights-pulse/src/models/responses"
	"insights-pulse/src/utils"
	"time"
)

type HomeAwayMetricsGenerator struct {
	*InsightGeneratorBase
}

func (h *HomeAwayMetricsGenerator) GetConfig() InsightConfig {
	return InsightConfig{
		Type:            "HomeAwayMetricsGenerator",
		TableName:       "home_away_metrics",
		Api:             "https://v3.football.api-sports.io",
		Endpoints:       []string{"/teams/statistics"},
		UpdateFrequency: 7 * 24 * time.Hour, //
	}
}

func (h *HomeAwayMetricsGenerator) GenerateAndSaveInsights(imeta teaminsights.StatsMetaData) error {

	params := dataclients.QueryParameters{
		TeamId:   imeta.TeamId,
		LeagueId: imeta.LeagueId,
		Season:   imeta.Season,
	}
	// INFO: Step 1. Get Season statistics
	resp := h.TeamClient.GetTeamSeasonStats(params)
	if resp == nil {
		return errors.New("something went wrong while fetching GetTeamSeasonStats")
	}
	// INFO: Step 2. Parse to HomeAwayMetrics and calculate Points
	homeAwayMetrics := h.parseToHomeAwayMetrics(resp)
	// INFO: Step 3. Save to Database
	h.TeamRepo.SaveHomeAwayMetrics(imeta, homeAwayMetrics)

	h.LogInfo(h.GetConfig(), imeta)
	return nil
}

func (h *HomeAwayMetricsGenerator) parseToHomeAwayMetrics(resp *responses.TeamStatsResponse) *teaminsights.HomeAwayMetrics {
	homeAwayMetrics := teaminsights.NewHomeAwayMetrics()

	homeAwayMetrics.Fixtures.Home = utils.ConvertToFloat64Ptr(resp.Response.Fixtures.Played.Home)
	homeAwayMetrics.Fixtures.Away = utils.ConvertToFloat64Ptr(resp.Response.Fixtures.Played.Away)
	homeAwayMetrics.Fixtures.Total = utils.ConvertToFloat64Ptr(resp.Response.Fixtures.Played.Total)
	homeAwayMetrics.Wins.Home = utils.ConvertToFloat64Ptr(resp.Response.Fixtures.Wins.Home)
	homeAwayMetrics.Wins.Away = utils.ConvertToFloat64Ptr(resp.Response.Fixtures.Wins.Away)
	homeAwayMetrics.Wins.Total = utils.ConvertToFloat64Ptr(resp.Response.Fixtures.Wins.Total)
	homeAwayMetrics.Draws.Home = utils.ConvertToFloat64Ptr(resp.Response.Fixtures.Draws.Home)
	homeAwayMetrics.Draws.Away = utils.ConvertToFloat64Ptr(resp.Response.Fixtures.Draws.Away)
	homeAwayMetrics.Draws.Total = utils.ConvertToFloat64Ptr(resp.Response.Fixtures.Draws.Total)
	homeAwayMetrics.Loses.Home = utils.ConvertToFloat64Ptr(resp.Response.Fixtures.Loses.Home)
	homeAwayMetrics.Loses.Away = utils.ConvertToFloat64Ptr(resp.Response.Fixtures.Loses.Away)
	homeAwayMetrics.Loses.Total = utils.ConvertToFloat64Ptr(resp.Response.Fixtures.Loses.Total)
	homeAwayMetrics.GoalsForTotal.Home = utils.ConvertToFloat64Ptr(resp.Response.Goals.For.Total.Home)
	homeAwayMetrics.GoalsForTotal.Away = utils.ConvertToFloat64Ptr(resp.Response.Goals.For.Total.Away)
	homeAwayMetrics.GoalsForTotal.Total = utils.ConvertToFloat64Ptr(resp.Response.Goals.For.Total.Total)
	homeAwayMetrics.GoalsForAverage.Home = utils.ConvertToFloat64Ptr(resp.Response.Goals.For.Average.Home)
	homeAwayMetrics.GoalsForAverage.Away = utils.ConvertToFloat64Ptr(resp.Response.Goals.For.Average.Away)
	homeAwayMetrics.GoalsForAverage.Total = utils.ConvertToFloat64Ptr(resp.Response.Goals.For.Average.Total)
	homeAwayMetrics.GoalsForMinute = resp.Response.Goals.For.Minute
	homeAwayMetrics.GoalsAgainstTotal.Home = utils.ConvertToFloat64Ptr(resp.Response.Goals.Against.Total.Home)
	homeAwayMetrics.GoalsAgainstTotal.Away = utils.ConvertToFloat64Ptr(resp.Response.Goals.Against.Total.Away)
	homeAwayMetrics.GoalsAgainstTotal.Total = utils.ConvertToFloat64Ptr(resp.Response.Goals.Against.Total.Total)
	homeAwayMetrics.GoalsAgainstAverage.Home = utils.ConvertToFloat64Ptr(resp.Response.Goals.Against.Average.Home)
	homeAwayMetrics.GoalsAgainstAverage.Away = utils.ConvertToFloat64Ptr(resp.Response.Goals.Against.Average.Away)
	homeAwayMetrics.GoalsAgainstAverage.Total = utils.ConvertToFloat64Ptr(resp.Response.Goals.Against.Average.Total)
	homeAwayMetrics.GoalsAgainstMinute = resp.Response.Goals.Against.Minute
	homeAwayMetrics.CleanSheets.Home = utils.ConvertToFloat64Ptr(resp.Response.CleanSheets.Home)
	homeAwayMetrics.CleanSheets.Away = utils.ConvertToFloat64Ptr(resp.Response.CleanSheets.Away)
	homeAwayMetrics.CleanSheets.Total = utils.ConvertToFloat64Ptr(resp.Response.CleanSheets.Total)
	homeAwayMetrics.FailedToScore.Home = utils.ConvertToFloat64Ptr(resp.Response.FailedToScore.Home)
	homeAwayMetrics.FailedToScore.Away = utils.ConvertToFloat64Ptr(resp.Response.FailedToScore.Away)
	homeAwayMetrics.FailedToScore.Total = utils.ConvertToFloat64Ptr(resp.Response.FailedToScore.Total)
	homeAwayMetrics.PointsPerGame.Away = utils.ConvertToFloat64Ptr(calculatePoints(*homeAwayMetrics.Wins.Away, *homeAwayMetrics.Draws.Away))
	homeAwayMetrics.PointsPerGame.Home = utils.ConvertToFloat64Ptr(calculatePoints(*homeAwayMetrics.Wins.Home, *homeAwayMetrics.Draws.Home))
	homeAwayMetrics.PointsPerGame.Total = utils.ConvertToFloat64Ptr(calculatePoints(*homeAwayMetrics.Wins.Total, *homeAwayMetrics.Draws.Total))

	return homeAwayMetrics
}

func calculatePoints(wins, draws float64) float64 {
	winsPoints := wins * 3
	drawsPoints := draws * 1
	return winsPoints + drawsPoints
}
