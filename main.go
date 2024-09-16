package main

import (
	"fmt"
	"insights-pulse/src/apiclients"
	con "insights-pulse/src/constants"
	"insights-pulse/src/dataclients"
	"insights-pulse/src/db"
	"insights-pulse/src/insightsgenerator/teamgenerator"
	"insights-pulse/src/models/insights/teaminsights"
	"insights-pulse/src/repositories/sqlrepo"
)

func main() {
	// Initialize Database Global Variable
	db.InitDb()

	// INFO: Initialize clients
	apiClient := apiclients.NewApiFootballClientImp()
	teamClient := dataclients.NewTeamClient(apiClient)
	teamRepo := sqlrepo.NewTeamRepository(db.DB)

	var insigeneratorBase *teamgenerator.InsightGeneratorBase = teamgenerator.NewInsightGeneratorBase(teamClient, teamRepo)

	// INFO: Insights Metadata
	statMetadata := teaminsights.StatsMetaData{
		TeamId:   "42",
		Season:   "2023",
		LeagueId: con.PREMIER_LEAGUE,
	}

	// INFO: ====== Avg Match Metrics Generator ======
	fmt.Println("------------ Avg Match Metrics Generator ------------")
	// INFO: Initialize the Insights generator
	var avgMetricsGen teamgenerator.InsightsGenerator = &teamgenerator.AvgMatchMetricsGenerator{
		InsightGeneratorBase: insigeneratorBase,
	}

	// INFO: Check if the insights should be updated
	if avgMetricsGen.ShouldUpdate(avgMetricsGen.GetConfig()) {
		avgMetricsGen.GenerateAndSaveInsights(teaminsights.StatsMetaData(statMetadata))
	} else {
		fmt.Println("No time for update yet")
	}

	// INFO: ====== Home Away Metrics Generator ======
	fmt.Println("------------ Home Away Metrics Generator ------------")

	var homeAwayMetricsGen teamgenerator.InsightsGenerator = &teamgenerator.HomeAwayMetricsGenerator{
		InsightGeneratorBase: insigeneratorBase,
	}

	// INFO: Check if the insights should be updated
	if homeAwayMetricsGen.ShouldUpdate(homeAwayMetricsGen.GetConfig()) {
		homeAwayMetricsGen.GenerateAndSaveInsights(teaminsights.StatsMetaData(statMetadata))
	} else {
		fmt.Println("No time for update yet")
	}

}
