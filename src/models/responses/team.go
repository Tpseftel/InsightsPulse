package responses

// INFO:
// - This package contain responses from v3.football.api-sports.io/ api endpoints
// - Here are gather responses about teams

type BasicInfoResponse struct {
	Results int `json:"results"`
}

type Team struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Logo string `json:"logo"`
}

type Statistics struct {
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

type TeamFixturesIResponse struct {
	BasicInfoResponse
	Response []struct {
		Fixture struct {
			ID int `json:"id"`
		} `json:"fixture"`
	} `json:"response"`
}

// INFO: For endpoint: v3.football.api-sports.io/teams/statistics
type TeamStatsResponse struct {
	BasicInfoResponse
	Response struct {
		League struct {
			ID      int    `json:"id"`
			Name    string `json:"name"`
			Country string `json:"country"`
			Logo    string `json:"logo"`
			Flag    string `json:"flag"`
			Season  int    `json:"season"`
		} `json:"league"`
		Team     `json:"team"`
		Form     string `json:"form"`
		Fixtures struct {
			Played struct {
				Home  int `json:"home"`
				Away  int `json:"away"`
				Total int `json:"total"`
			} `json:"played"`
			Wins struct {
				Home  int `json:"home"`
				Away  int `json:"away"`
				Total int `json:"total"`
			} `json:"wins"`
			Draws struct {
				Home  int `json:"home"`
				Away  int `json:"away"`
				Total int `json:"total"`
			} `json:"draws"`
			Loses struct {
				Home  int `json:"home"`
				Away  int `json:"away"`
				Total int `json:"total"`
			} `json:"loses"`
		} `json:"fixtures"`
		Goals struct {
			For struct {
				Total struct {
					Home  int `json:"home"`
					Away  int `json:"away"`
					Total int `json:"total"`
				} `json:"total"`
				Average struct {
					Home  string `json:"home"`
					Away  string `json:"away"`
					Total string `json:"total"`
				} `json:"average"`
				Minute map[string]struct {
					Total      int    `json:"total"`
					Percentage string `json:"percentage"`
				} `json:"minute"`
			} `json:"for"`
			Against struct {
				Total struct {
					Home  int `json:"home"`
					Away  int `json:"away"`
					Total int `json:"total"`
				} `json:"total"`
				Average struct {
					Home  string `json:"home"`
					Away  string `json:"away"`
					Total string `json:"total"`
				} `json:"average"`
				Minute map[string]struct {
					Total      int    `json:"total"`
					Percentage string `json:"percentage"`
				} `json:"minute"`
			} `json:"against"`
		} `json:"goals"`
		Biggest struct {
			Streak struct {
				Wins  int `json:"wins"`
				Draws int `json:"draws"`
				Loses int `json:"loses"`
			} `json:"streak"`
			Wins struct {
				Home string `json:"home"`
				Away string `json:"away"`
			} `json:"wins"`
			Loses struct {
				Home string `json:"home"`
				Away string `json:"away"`
			} `json:"loses"`
			Goals struct {
				For struct {
					Home int `json:"home"`
					Away int `json:"away"`
				} `json:"for"`
				Against struct {
					Home int `json:"home"`
					Away int `json:"away"`
				} `json:"against"`
			} `json:"goals"`
		} `json:"biggest"`
		CleanSheet struct {
			Home  int `json:"home"`
			Away  int `json:"away"`
			Total int `json:"total"`
		} `json:"clean_sheet"`
		FailedToScore struct {
			Home  int `json:"home"`
			Away  int `json:"away"`
			Total int `json:"total"`
		} `json:"failed_to_score"`
		Penalty struct {
			Scored struct {
				Total      int    `json:"total"`
				Percentage string `json:"percentage"`
			} `json:"scored"`
			Missed struct {
				Total      int    `json:"total"`
				Percentage string `json:"percentage"`
			} `json:"missed"`
			Total int `json:"total"`
		} `json:"penalty"`
		Lineups []struct {
			Formation string `json:"formation"`
			Played    int    `json:"played"`
		} `json:"lineups"`
		Cards struct {
			Yellow map[string]struct {
				Total      int    `json:"total"`
				Percentage string `json:"percentage"`
			} `json:"yellow"`
			Red map[string]struct {
				Total      int    `json:"total"`
				Percentage string `json:"percentage"`
			} `json:"red"`
		} `json:"cards"`
	} `json:"response"`
}

// INFO: For endpoint: v3.football.api-sports.io/fixtures/statistics?fixture=1035054&team=463
type FixtureTeamStatsResponse struct {
	BasicInfoResponse
	Response []struct {
		Team       Team         `json:"team"`
		Statistics []Statistics `json:"statistics"`
	} `json:"response"`
}

type Fixture struct {
	ID        int    `json:"id"`
	Referee   string `json:"referee"`
	Date      string `json:"date"`
	Timestamp int64  `json:"timestamp"`
}

type ScoreDetail struct {
	Home int
	Away int
}

type Score struct {
	Halftime  ScoreDetail  `json:"halftime"`
	Fulltime  ScoreDetail  `json:"fulltime"`
	Extratime *ScoreDetail `json:"extratime"` // Use pointer to handle null values
	Penalty   *ScoreDetail `json:"penalty"`   // Use pointer to handle null values
}

// INFO: For endpoint: v3.football.api-sports.io/fixtures?id=1035054
type FixtureStatsResponse struct {
	BasicInfoResponse
	Response []struct {
		Fixture Fixture `json:"fixture"`
		Teams   struct {
			Home struct {
				ID     int    `json:"id"`
				Name   string `json:"name"`
				Logo   string `json:"logo"`
				Winner bool   `json:"winner"`
			} `json:"home"`
			Away struct {
				ID     int    `json:"id"`
				Name   string `json:"name"`
				Logo   string `json:"logo"`
				Winner bool   `json:"winner"`
			} `json:"away"`
		} `json:"teams"`
		Goals      ScoreDetail `json:"goals"`
		Score      Score       `json:"score"`
		Statistics []struct {
			Team       Team         `json:"team"`
			Statistics []Statistics `json:"statistics"`
		}
	} `json:"response"`
}
