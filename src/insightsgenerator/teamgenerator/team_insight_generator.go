package teamgenerator

import "insights-pulse/src/models/insights/teaminsights"

type InsightsGenerator interface {
	GenerateAndSaveInsights(imeta teaminsights.StatsMetaData) error
	GetConfig() InsightConfig
}
