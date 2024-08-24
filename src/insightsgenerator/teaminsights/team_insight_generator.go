package teaminsights

import (
	"insights-pulse/src/models/insights"
)

type InsightsGenerator interface {
	GenerateAndSaveInsights(imeta insights.StatsMetaData) error
	GetConfig() InsightConfig
}
