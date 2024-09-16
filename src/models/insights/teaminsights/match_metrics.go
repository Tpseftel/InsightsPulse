package teaminsights

type MatchStatsDetail struct {
	Sum float64
	Num float64
	Avg float64
}

type StatsMetaData struct {
	TeamId   string `json:"teamId"`
	Season   string `json:"season"`
	LeagueId string `json:"league"`
}

type MatchMetrics struct {
	ShotsOnGoal      *MatchStatsDetail `json:"Shots on Goal"`
	ShotsOffGoal     *MatchStatsDetail `json:"Shots off Goal"`
	TotalShots       *MatchStatsDetail `json:"Total Shots"`
	BlockedShots     *MatchStatsDetail `json:"Blocked Shots"`
	ShotsInsideBox   *MatchStatsDetail `json:"Shots insidebox"`
	ShotsOutsideBox  *MatchStatsDetail `json:"Shots outsidebox"`
	Fouls            *MatchStatsDetail `json:"Fouls"`
	CornerKicks      *MatchStatsDetail `json:"Corner Kicks"`
	Offsides         *MatchStatsDetail `json:"Offsides"`
	BallPossession   *MatchStatsDetail `json:"Ball Possession"`
	YellowCards      *MatchStatsDetail `json:"Yellow Cards"`
	RedCards         *MatchStatsDetail `json:"Red Cards"`
	GoalkeeperSaves  *MatchStatsDetail `json:"Goalkeeper Saves"`
	TotalPasses      *MatchStatsDetail `json:"Total passes"`
	PassesAccurate   *MatchStatsDetail `json:"Passes accurate"`
	PassesPercentage *MatchStatsDetail `json:"Passes %"`
	ExpectedGoals    *MatchStatsDetail `json:"expected_goals"`
}

func NewMatchMetrics() *MatchMetrics {
	return &MatchMetrics{
		ShotsOnGoal:      &MatchStatsDetail{},
		ShotsOffGoal:     &MatchStatsDetail{},
		TotalShots:       &MatchStatsDetail{},
		BlockedShots:     &MatchStatsDetail{},
		ShotsInsideBox:   &MatchStatsDetail{},
		ShotsOutsideBox:  &MatchStatsDetail{},
		Fouls:            &MatchStatsDetail{},
		CornerKicks:      &MatchStatsDetail{},
		Offsides:         &MatchStatsDetail{},
		BallPossession:   &MatchStatsDetail{},
		YellowCards:      &MatchStatsDetail{},
		RedCards:         &MatchStatsDetail{},
		GoalkeeperSaves:  &MatchStatsDetail{},
		TotalPasses:      &MatchStatsDetail{},
		PassesAccurate:   &MatchStatsDetail{},
		PassesPercentage: &MatchStatsDetail{},
		ExpectedGoals:    &MatchStatsDetail{},
	}
}
