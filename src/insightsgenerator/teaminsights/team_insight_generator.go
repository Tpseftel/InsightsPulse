package teaminsights

import (
	"fmt"
	"insights-pulse/src/models/insights"
	"insights-pulse/src/models/responses"

	"insights-pulse/src/utils"
)

type InsightsGenerator interface {
	GetFixtureIds(teamId, season, league string) []int
	GetFixtureStats(idsChunks []string) []responses.FixtureStatsResponse
	CalculateStatsDetails(fixtureStats []responses.FixtureStatsResponse) *insights.MatchMetrics
	SaveMetrics(meta insights.StatsMetaData, insights *insights.MatchMetrics) error
}

type TeamInsightGenerator struct {
	Ig InsightsGenerator
}

func (i *TeamInsightGenerator) GenerateAndSaveInsights(imeta insights.StatsMetaData) error {
	// INFO: Step 1. Get fixture ids
	fixtureIds := i.Ig.GetFixtureIds(imeta.TeamId, imeta.Season, imeta.LeagueId)
	fmt.Println("Fixture Ids: ", fixtureIds)

	idsChunks := utils.StringfyIds(fixtureIds, 20)
	fmt.Println(" idsChunks: ", idsChunks)

	// INFO: Step 2. Get fixture stats
	fixtureStats := i.Ig.GetFixtureStats(idsChunks)

	// INFO: Step 3. Generate stats details
	statsDetails := i.Ig.CalculateStatsDetails(fixtureStats)

	fmt.Println("Stats Details: ", statsDetails)

	// INFO: Step 4. Save the insights
	i.Ig.SaveMetrics(imeta, statsDetails)

	return nil
}
