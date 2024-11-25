package teamgenerator

import (
	"errors"
	"math"
	"strconv"
	"sync"
	"time"

	"insights-pulse/src/dataclients"
	"insights-pulse/src/models/insights/teaminsights"
	"insights-pulse/src/models/responses"
	"insights-pulse/src/utils"

	"insights-pulse/src/logger"
)

type AvgMatchMetricsGenerator struct {
	*InsightGeneratorBase
}

func (a *AvgMatchMetricsGenerator) GetConfig() InsightConfig {
	return InsightConfig{
		Type:            "AvgMatchMetricsGenerator",
		TableName:       "avg_insights_per_game_team",
		Api:             "https://v3.football.api-sports.io",
		Endpoints:       []string{"/fixtures?team=33&league=39&season=2020"},
		UpdateFrequency: 55 * time.Minute, //  Weekly update
	}
}

func (a *AvgMatchMetricsGenerator) GenerateAndSaveInsights(imeta teaminsights.StatsMetaData) error {
	// INFO: Step 1. Get fixture ids
	fixtureIds := a.getFixtureIds(imeta.TeamId, imeta.Season, imeta.LeagueId)
	if len(fixtureIds) == 0 {
		return errors.New("something went wrong while fetching fixture ids")
	}

	idsChunks := utils.StringfyIds(fixtureIds, 20)
	// INFO: Step 2. Get fixture stats
	fixtureStats := a.getFixtureStats(idsChunks)
	if len(fixtureStats) == 0 {
		return errors.New("something went wrong while fetching fixture stats")
	}

	// INFO: Step 3. Generate stats details
	statsDetails := a.calculateStatsDetails(fixtureStats, imeta.TeamId)

	// INFO: Step 4. Save the insights
	err := a.TeamRepo.SaveAvgInsightsPerGame(imeta, statsDetails)
	if err != nil {
		logger.GetLogger().Error("Error saving avg insights per game: " + err.Error())
		return err
	}

	a.LogInfo(a.GetConfig(), imeta)
	return nil
}

func (a *AvgMatchMetricsGenerator) getFixtureIds(teamId, season, league string) []int {
	params := dataclients.QueryParameters{
		TeamId:   teamId,
		Season:   season,
		LeagueId: league,
	}
	fixtureIds := a.TeamClient.GetFixtures(params)
	return fixtureIds
}

func (a *AvgMatchMetricsGenerator) getFixtureStats(idsChunks []string) []responses.FixtureStatsResponse {
	// Holds the api responses for each idschunk
	var seasonFixtures = make([]responses.FixtureStatsResponse, 0)
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, chunk := range idsChunks {
		wg.Add(1)
		go func(ids string) {
			defer wg.Done()
			mu.Lock()
			seasonFixtures = append(seasonFixtures, *a.TeamClient.GetFixturebyIds(ids))
			mu.Unlock()
		}(chunk)
	}
	wg.Wait()
	return seasonFixtures
}

func (a *AvgMatchMetricsGenerator) calculateStatsDetails(fixtureStats []responses.FixtureStatsResponse, teamId string) *teaminsights.MatchMetrics {
	stats := make(map[string]teaminsights.MatchStatsDetail)
	for _, response := range fixtureStats {
		for _, fixture := range response.Response {
			for _, stat := range fixture.Statistics {
				if utils.ConvToString(stat.Team.ID) == teamId {
					for _, v := range stat.Statistics {
						switch v.Type {
						case "Shots on Goal":
							if value, ok := v.Value.(float64); ok {
								tempStat := stats[v.Type]
								tempStat.Num++
								tempStat.Sum += value
								stats[v.Type] = tempStat
							}
						case "Shots off Goal":
							if value, ok := v.Value.(float64); ok {
								tempStat := stats[v.Type]
								tempStat.Num++
								tempStat.Sum += value
								stats[v.Type] = tempStat
							}
						case "Total Shots":
							if value, ok := v.Value.(float64); ok {
								tempStat := stats[v.Type]
								tempStat.Num++
								tempStat.Sum += value
								stats[v.Type] = tempStat
							}
						case "Blocked Shots":
							if value, ok := v.Value.(float64); ok {
								tempStat := stats[v.Type]
								tempStat.Num++
								tempStat.Sum += value
								stats[v.Type] = tempStat
							}
						case "Shots insidebox":
							if value, ok := v.Value.(float64); ok {
								tempStat := stats[v.Type]
								tempStat.Num++
								tempStat.Sum += value
								stats[v.Type] = tempStat
							}
						case "Shots outsidebox":
							if value, ok := v.Value.(float64); ok {
								tempStat := stats[v.Type]
								tempStat.Num++
								tempStat.Sum += value
								stats[v.Type] = tempStat
							}
						case "Fouls":
							if value, ok := v.Value.(float64); ok {
								tempStat := stats[v.Type]
								tempStat.Num++
								tempStat.Sum += value
								stats[v.Type] = tempStat
							}
						case "Corner Kicks":
							if value, ok := v.Value.(float64); ok {
								tempStat := stats[v.Type]
								tempStat.Num++
								tempStat.Sum += value
								stats[v.Type] = tempStat
							}
						case "Offsides":
							if value, ok := v.Value.(float64); ok {
								tempStat := stats[v.Type]
								tempStat.Num++
								tempStat.Sum += value
								stats[v.Type] = tempStat
							}
						case "Ball Possession":
							tempStat := stats[v.Type]
							if value, ok := v.Value.(float64); ok {
								tempStat.Num++
								tempStat.Sum += value
								stats[v.Type] = tempStat
							} else if value, ok := v.Value.(string); ok {
								pureFloat, _ := utils.GetFloatFromPercentage(value)
								tempStat.Num++
								tempStat.Sum += pureFloat
								stats[v.Type] = tempStat
							}
						case "Yellow Cards":
							if value, ok := v.Value.(float64); ok {
								tempStat := stats[v.Type]
								tempStat.Num++
								tempStat.Sum += value
								stats[v.Type] = tempStat
							}
						case "Red Cards":
							if value, ok := v.Value.(float64); ok {
								tempStat := stats[v.Type]
								tempStat.Num++
								tempStat.Sum += value
								stats[v.Type] = tempStat
							}
						case "Goalkeeper Saves":
							if value, ok := v.Value.(float64); ok {
								tempStat := stats[v.Type]
								tempStat.Num++
								tempStat.Sum += value
								stats[v.Type] = tempStat
							}
						case "Total passes":
							if value, ok := v.Value.(float64); ok {
								tempStat := stats[v.Type]
								tempStat.Num++
								tempStat.Sum += value
								stats[v.Type] = tempStat
							}
						case "Passes accurate":
							if value, ok := v.Value.(float64); ok {
								tempStat := stats[v.Type]
								tempStat.Num++
								tempStat.Sum += value
								stats[v.Type] = tempStat
							}
						case "Passes %":
							tempStat := stats[v.Type]
							if value, ok := v.Value.(float64); ok {
								tempStat.Num++
								tempStat.Sum += value
								stats[v.Type] = tempStat
							} else if value, ok := v.Value.(string); ok {
								pureFloat, _ := utils.GetFloatFromPercentage(value)
								tempStat.Num++
								tempStat.Sum += pureFloat
								stats[v.Type] = tempStat
							}
						case "expected_goals":
							tempStat := stats[v.Type]
							if value, ok := v.Value.(float64); ok {
								tempStat.Num++
								tempStat.Sum += value
								stats[v.Type] = tempStat
							} else if value, ok := v.Value.(string); ok {
								pureFloat, err := strconv.ParseFloat(value, 64)
								if err != nil {
									logger.GetLogger().Warn("Error parsing float: %v\n")
									continue
								}
								tempStat.Num++
								tempStat.Sum += pureFloat
								stats[v.Type] = tempStat
							}
						}
					}
				}
			}
		}
	}
	calculateAverageStats(stats)

	return mapStatsToInsights(stats)

}

func calculateAverageStats(stats map[string]teaminsights.MatchStatsDetail) {
	for key, v := range stats {
		if v.Num != 0 {
			tempVar := v.Sum / v.Num
			v.Avg = math.Round(tempVar*10) / 10
		} else {
			v.Avg = 0
		}
		stats[key] = v
	}
}

func mapStatsToInsights(stats map[string]teaminsights.MatchStatsDetail) *teaminsights.MatchMetrics {
	// Initialize the AverageInsightsPerGame with empty StatDetail pointers
	insights := teaminsights.NewMatchMetrics()
	// Map the data from stats to the fields in AverageInsightsPerGame
	for key, stat := range stats {
		switch key {
		case "Shots on Goal":
			insights.ShotsOnGoal = &stat
		case "Shots off Goal":
			insights.ShotsOffGoal = &stat
		case "Total Shots":
			insights.TotalShots = &stat
		case "Blocked Shots":
			insights.BlockedShots = &stat
		case "Shots insidebox":
			insights.ShotsInsideBox = &stat
		case "Shots outsidebox":
			insights.ShotsOutsideBox = &stat
		case "Fouls":
			insights.Fouls = &stat
		case "Corner Kicks":
			insights.CornerKicks = &stat
		case "Offsides":
			insights.Offsides = &stat
		case "Ball Possession":
			insights.BallPossession = &stat
		case "Yellow Cards":
			insights.YellowCards = &stat
		case "Red Cards":
			insights.RedCards = &stat
		case "Goalkeeper Saves":
			insights.GoalkeeperSaves = &stat
		case "Total passes":
			insights.TotalPasses = &stat
		case "Passes accurate":
			insights.PassesAccurate = &stat
		case "Passes %":
			insights.PassesPercentage = &stat
		case "expected_goals":
			insights.ExpectedGoals = &stat
		}
	}
	return insights
}
