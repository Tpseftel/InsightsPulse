package teaminsights

import "insights-pulse/src/models/responses"

type HomeAwayData struct {
	Home  *float64 `json:"home"`
	Away  *float64 `json:"away"`
	Total *float64 `json:"total"`
}

type MinuteData = responses.MinuteData

// INFO: Main struct
type HomeAwayMetrics struct {
	Fixtures            *HomeAwayData         `json:"fixtures"`
	Wins                *HomeAwayData         `json:"wins"`
	Draws               *HomeAwayData         `json:"draws"`
	Loses               *HomeAwayData         `json:"loses"`
	GoalsForTotal       *HomeAwayData         `json:"goalsForTotal"`
	GoalsForAverage     *HomeAwayData         `json:"goalsForAverage"`
	GoalsForMinute      map[string]MinuteData `json:"goalsForMinute"`
	GoalsAgainstTotal   *HomeAwayData         `json:"goalsAgainstTotal"`
	GoalsAgainstAverage *HomeAwayData         `json:"goalsAgainstAverage"`
	GoalsAgainstMinute  map[string]MinuteData `json:"goalsAgainstMinute"`
	CleanSheets         *HomeAwayData         `json:"cleanSheets"`
	FailedToScore       *HomeAwayData         `json:"failedToScore"`
	PointsPerGame       *HomeAwayData         `json:"pointsPerGame"`
}

func NewHomeAwayMetrics() *HomeAwayMetrics {
	return &HomeAwayMetrics{
		Fixtures:            &HomeAwayData{},
		Wins:                &HomeAwayData{},
		Draws:               &HomeAwayData{},
		Loses:               &HomeAwayData{},
		GoalsForTotal:       &HomeAwayData{},
		GoalsForAverage:     &HomeAwayData{},
		GoalsForMinute:      make(map[string]MinuteData),
		GoalsAgainstTotal:   &HomeAwayData{},
		GoalsAgainstAverage: &HomeAwayData{},
		GoalsAgainstMinute:  make(map[string]MinuteData),
		CleanSheets:         &HomeAwayData{},
		FailedToScore:       &HomeAwayData{},
		PointsPerGame:       &HomeAwayData{},
	}
}

