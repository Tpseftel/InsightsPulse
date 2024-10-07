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

type HomeAwayData struct {
	Home  interface{} `json:"home"`
	Away  interface{} `json:"away"`
	Total interface{} `json:"total"`
}

type MinuteData struct {
	Total      interface{} `json:"total"`
	Percentage interface{} `json:"percentage"`
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
			Played HomeAwayData `json:"played"`
			Wins   HomeAwayData `json:"wins"`
			Draws  HomeAwayData `json:"draws"`
			Loses  HomeAwayData `json:"loses"`
		} `json:"fixtures"`
		Goals struct {
			For struct {
				Total   HomeAwayData          `json:"total"`
				Average HomeAwayData          `json:"average"`
				Minute  map[string]MinuteData `json:"minute"`
			} `json:"for"`
			Against struct {
				Total   HomeAwayData          `json:"total"`
				Average HomeAwayData          `json:"average"`
				Minute  map[string]MinuteData `json:"minute"`
			} `json:"against"`
		} `json:"goals"`
		Biggest struct {
			Streak struct {
				Wins  int `json:"wins"`
				Draws int `json:"draws"`
				Loses int `json:"loses"`
			} `json:"streak"`
			Wins  HomeAwayData `json:"wins"`
			Loses HomeAwayData `json:"loses"`
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
		CleanSheets   HomeAwayData `json:"clean_sheet"`
		FailedToScore HomeAwayData `json:"failed_to_score"`
		Penalty       struct {
			Scored struct {
				Total      interface{} `json:"total"`
				Percentage interface{} `json:"percentage"`
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

type TeamInfo struct {
	Team struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Code     string `json:"code"`
		Country  string `json:"country"`
		Founded  int    `json:"founded"`
		National bool   `json:"national"`
		Logo     string `json:"logo"`
	} `json:"team"`
	Venue struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Address  string `json:"address"`
		City     string `json:"city"`
		Capacity int    `json:"capacity"`
		Surface  string `json:"surface"`
		Image    string `json:"image"`
	} `json:"venue"`
}

// INFO: For endpoint: v3.football.api-sports.io/teams
type TeamsInfoResponse struct {
	BasicInfoResponse
	Response []TeamInfo `json:"response"`
}
