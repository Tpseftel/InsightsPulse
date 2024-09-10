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

	// INFO: Initialize the Insights generator
	avgMetricsGen := &teamgenerator.AvgMatchMetricsGenerator{
		TeamClient: teamClient,
		TeamRepo:   teamRepo,
	}

	// INFO: Generate and save insights
	statMetadata := teaminsights.StatsMetaData{
		// TeamId: "33",
		// TeamId:   "50",
		TeamId: "42",
		// TeamId:   "55",
		Season:   "2023",
		LeagueId: con.PREMIER_LEAGUE,
	}

	// INFO: Check if the insights should be updated
	if avgMetricsGen.ShouldUpdate(avgMetricsGen.GetConfig()) {
		avgMetricsGen.GenerateAndSaveInsights(teaminsights.StatsMetaData(statMetadata))
	} else {
		fmt.Println("No time for update yet")
	}

	fmt.Println("------------Successfully generated and saved insights------------")

}
