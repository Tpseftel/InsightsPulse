package main

import (
	"insights-pulse/src/apiclients"
	"insights-pulse/src/dataclients"
	"insights-pulse/src/db"
	"insights-pulse/src/insightsgenerator/teamgenerator"
	"insights-pulse/src/models/insights/teaminsights"
	"insights-pulse/src/repositories/sqlrepo"
)

func main() {
	// Initialize Database Global Variable
	db.InitDb()

	// config, err := config.GetConfig()
	// if err != nil {
	// 	logger.GetLogger().Error("Cannot load Configuration variables: " + err.Error())
	// 	panic(err)
	// }

	// INFO: Initialize clients
	apiClient := apiclients.NewApiFootballClientImp()
	teamClient := dataclients.NewTeamClient(apiClient)
	teamRepo := sqlrepo.NewTeamRepository(db.DB)

	var insigeneratorBase *teamgenerator.InsightGeneratorBase = teamgenerator.NewInsightGeneratorBase(teamClient, teamRepo)

	// INFO: Initialize the insight generators
	// var homeAwayMetricsGen teamgenerator.InsightsGenerator = &teamgenerator.HomeAwayMetricsGenerator{
	// 	InsightGeneratorBase: insigeneratorBase,
	// }
	// var avgSeasonMetricsGen teamgenerator.InsightsGenerator = &teamgenerator.AvgMatchMetricsGenerator{
	// 	InsightGeneratorBase: insigeneratorBase,
	// }
	var teamsInfo teamgenerator.InsightsGenerator = &teamgenerator.TeamsInfoGenerator{
		InsightGeneratorBase: insigeneratorBase,
	}

	// var leagueCollector *collectors.LeagueCollector = &collectors.LeagueCollector{
	// 	TeamClient: teamClient,
	// 	Config:     config,
	// }

	// INFO: Get League teams
	teamsInfo.GenerateAndSaveInsights(teaminsights.StatsMetaData{
		LeagueId: "39",
		Season:   "2024",
	})

	// INFO: Parse TeamsInfoResponse to models.TeamsInfo

	// fmt.Println(teamsResponse)

	// =========== INFO: For Current Season 2024 Home Away Metricss ============
	// leagueCollector.CollectLeagueData(con.PREMIER_LEAGUE, "2024", homeAwayMetricsGen)
	// leagueCollector.CollectLeagueData(con.SUPER_LEAGUE, "2024", homeAwayMetricsGen)
	// leagueCollector.CollectLeagueData(con.LA_LIGA, "2024", homeAwayMetricsGen)
	// leagueCollector.CollectLeagueData(con.SERIE_A, "2024", homeAwayMetricsGen)
	// leagueCollector.CollectLeagueData(con.BUNDESLIGA, "2024", homeAwayMetricsGen)
	// leagueCollector.CollectLeagueData(con.LEAGUE_ONE, "2024", homeAwayMetricsGen)

	//  ============ INFO: For Current Season 2024 Average Season Metrics ============
	// leagueCollector.CollectLeagueData(con.PREMIER_LEAGUE, "2024", avgSeasonMetricsGen)
	// leagueCollector.CollectLeagueData(con.SUPER_LEAGUE, "2024", avgSeasonMetricsGen)
	// leagueCollector.CollectLeagueData(con.LA_LIGA, "2024", avgSeasonMetricsGen)
	// leagueCollector.CollectLeagueData(con.SERIE_A, "2024", avgSeasonMetricsGen)
	// leagueCollector.CollectLeagueData(con.BUNDESLIGA, "2024", avgSeasonMetricsGen)
	// leagueCollector.CollectLeagueData(con.LEAGUE_ONE, "2024", avgSeasonMetricsGen)
}
