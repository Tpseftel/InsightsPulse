package main

import (
	"insights-pulse/src/apiclients"
	con "insights-pulse/src/constants"
	"insights-pulse/src/dataclients"
	"insights-pulse/src/db"
	"insights-pulse/src/insightsgenerator/teaminsights"
	"insights-pulse/src/models/insights"
)

func main() {
	// Initialize Database Global Variable
	db.InitDb()

	// INFO: Initialize the clients
	apiClient := apiclients.NewApiFootballClientImp()
	teamClient := dataclients.NewTeamClient(apiClient)

	// INFO: Initialize the generator
	avgMetricsGen := &teaminsights.AvgMatchMetricsGenerator{
		TeamClient: teamClient,
	}
	// INFO: Initialize the team insights generator
	tg := teaminsights.TeamInsightGenerator{
		Ig: avgMetricsGen,
	}

	// INFO: Generate and save insights
	statMetadata := insights.StatsMetadata{
		TeamId:   "33",
		Season:   "2023",
		LeagueId: con.PREMIER_LEAGUE,
	}
	tg.GenerateAndSaveInsights(insights.StatsMetaData(statMetadata))

}
