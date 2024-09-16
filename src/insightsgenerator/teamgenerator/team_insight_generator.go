package teamgenerator

import (
	"insights-pulse/src/dataclients"
	"insights-pulse/src/logger"
	"insights-pulse/src/models/insights/teaminsights"
	"insights-pulse/src/repositories/sqlrepo"
	"insights-pulse/src/utils"
	"time"
)

type InsightsGenerator interface {
	GenerateAndSaveInsights(imeta teaminsights.StatsMetaData) error
	GetConfig() InsightConfig
	ShouldUpdate(config InsightConfig) bool
}

// INFO: Config for the insights generator
type InsightConfig struct {
	Type            string
	Api             string
	Endpoints       []string
	TableName       string
	UpdateFrequency time.Duration
}

// INFO: Base struct for the insights generator
type InsightGeneratorBase struct {
	TeamClient *dataclients.TeamClient
	TeamRepo   *sqlrepo.TeamRepository
}

func NewInsightGeneratorBase(teamClient *dataclients.TeamClient, teamRepo *sqlrepo.TeamRepository) *InsightGeneratorBase {
	return &InsightGeneratorBase{teamClient, teamRepo}
}

func (i *InsightGeneratorBase) ShouldUpdate(config InsightConfig) bool {
	{
		lastUpdated, err := i.TeamRepo.GetLastUpdatedTime(config.TableName)
		if err != nil {
			logger.GetLogger().Error("Error getting last updated: " + err.Error())
			return true
		}
		if lastUpdated.IsZero() {
			return true

		}

		if time.Since(lastUpdated) > config.UpdateFrequency {
			return true
		}
		return false

	}
}

func (i *HomeAwayMetricsGenerator) LogInfo(config InsightConfig, imeta teaminsights.StatsMetaData) {
	logger.GetLogger().Info(config.Type + " : Successfully run")
	logger.GetLogger().Info(utils.ConvToString(imeta))
	logger.GetLogger().Info(utils.StructToString(config))
}
