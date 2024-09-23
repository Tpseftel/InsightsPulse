package main

import (
	"insights-pulse/src/apiclients"
	"insights-pulse/src/collectors"
	"insights-pulse/src/config"
	con "insights-pulse/src/constants"
	"insights-pulse/src/dataclients"
	"insights-pulse/src/db"
	"insights-pulse/src/insightsgenerator/teamgenerator"
	"insights-pulse/src/logger"
	"insights-pulse/src/repositories/sqlrepo"
)

func main() {
	// Initialize Database Global Variable
	db.InitDb()

	config, err := config.GetConfig()
	if err != nil {
		logger.GetLogger().Error("Cannot load Configuration variables: " + err.Error())
		panic(err)
	}

	// INFO: Initialize clients
	apiClient := apiclients.NewApiFootballClientImp()
	teamClient := dataclients.NewTeamClient(apiClient)
	teamRepo := sqlrepo.NewTeamRepository(db.DB)

	var insigeneratorBase *teamgenerator.InsightGeneratorBase = teamgenerator.NewInsightGeneratorBase(teamClient, teamRepo)

	var homeAwayMetricsGen teamgenerator.InsightsGenerator = &teamgenerator.HomeAwayMetricsGenerator{
		InsightGeneratorBase: insigeneratorBase,
	}
	var leagueCollector *collectors.LeagueCollector = &collectors.LeagueCollector{
		TeamClient: teamClient,
		Config:     config,
	}

	leagueCollector.CollectLeagueData(con.PREMIER_LEAGUE, "2023", homeAwayMetricsGen)
}
