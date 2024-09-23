package collectors

import (
	"fmt"
	"insights-pulse/src/config"
	"insights-pulse/src/dataclients"
	"insights-pulse/src/insightsgenerator/teamgenerator"
	"insights-pulse/src/logger"
	"insights-pulse/src/models/insights/teaminsights"
	"insights-pulse/src/utils"
	"time"
)

type LeagueCollector struct {
	TeamClient *dataclients.TeamClient
	Config     *config.Config
}

func (tc *LeagueCollector) CollectLeagueData(leagueId, season string, insight teamgenerator.InsightsGenerator) {

	// INFO: Batching related data for the API
	delay := time.Minute * 1
	batchSize := tc.Config.ApiRequestsPerMinuteLimit - 1
	fmt.Println("Batch Size: ", batchSize)

	plTeams := tc.TeamClient.GetTeams(leagueId, season)
	if plTeams == nil {
		logger.GetLogger().Warn(fmt.Sprintf("No teams found for league %s", leagueId))
		return
	}

	requesteCounter := 0
	for _, team := range plTeams.Response {
		statMetadata := teaminsights.StatsMetaData{
			TeamId:   fmt.Sprintf("%d", team.Team.ID),
			Season:   "2023",
			LeagueId: leagueId,
		}

		if insight.ShouldUpdate(insight.GetConfig()) {
			err := insight.GenerateAndSaveInsights(statMetadata)
			if err != nil {
				logger.GetLogger().Error(fmt.Sprintf("Error saving insights for team: %v \n League: %v  \n Season: %v",
					statMetadata.TeamId, statMetadata.LeagueId, statMetadata.Season))
			}
		} else {
			fmt.Println("No time for update yet")
			logger.GetLogger().Info(fmt.Sprintf("No time for update yet for team %v", statMetadata.TeamId))
		}

		requesteCounter++

		fmt.Println("Batch size: ", batchSize)
		fmt.Println("Data: ", utils.StructToString(statMetadata))
		fmt.Println("Request Counter: ", requesteCounter)

		// INFO: Check if the request counter is equal to the batch size and sleep for a minute
		if requesteCounter == batchSize {
			fmt.Println("============ Time for sleep:" + delay.Abs().String() + " ============")
			time.Sleep(delay)
			requesteCounter = 0
		}
	}

}
