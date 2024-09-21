package unit_test

import (
	"fmt"
	"insights-pulse/src/dataclients"
	"insights-pulse/src/tests/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTeamSeasonStats(t *testing.T) {
	tests := []struct {
		name               string
		statsMetadata      map[string]string
		mockedResponse     string
		expectError        bool
		expectedTeamID     int
		expectedTeamName   string
		expectedSeason     int
		expectedLeagueName string
	}{
		{
			name: "Valid response",
			statsMetadata: map[string]string{
				"teamId":   "33",
				"leagueId": "39",
				"season":   "2023",
			},
			mockedResponse: `
            {
                "get": "teams/statistics",
                "parameters": {
                    "team": "33",
                    "season": "2023",
                    "league": "39"
                },
                "errors": [],
                "results": 11,
                "paging": {
                    "current": 1,
                    "total": 1
                },
                "response": {
                    "league": {
                        "id": 39,
                        "name": "Premier League",
                        "country": "England",
                        "logo": "https://media.api-sports.io/football/leagues/39.png",
                        "flag": "https://media.api-sports.io/flags/gb.svg",
                        "season": 2023
                    },
                    "team": {
                        "id": 33,
                        "name": "Manchester United",
                        "logo": "https://media.api-sports.io/football/teams/33.png"
                    },
                    "form": "WLWLLWLWWLWWWLWLDLWLDWWWWLLWDLDDWDLLWW",
                    "fixtures": {
                        "played": {
                            "home": 19,
                            "away": 19,
                            "total": 38
                        },
                        "wins": {
                            "home": 10,
                            "away": 8,
                            "total": 18
                        },
                        "draws": {
                            "home": 3,
                            "away": 3,
                            "total": 6
                        },
                        "loses": {
                            "home": 6,
                            "away": 8,
                            "total": 14
                        }
                    },
                    "goals": {
                        "for": {
                            "total": {
                                "home": 31,
                                "away": 26,
                                "total": 57
                            },
                            "average": {
                                "home": "1.6",
                                "away": "1.4",
                                "total": "1.5"
                            }
                        },
                        "against": {
                            "total": {
                                "home": 28,
                                "away": 30,
                                "total": 58
                            },
                            "average": {
                                "home": "1.5",
                                "away": "1.6",
                                "total": "1.5"
                            }
                        }
                    }
                }
            }`,
			expectError:        false,
			expectedTeamID:     33,
			expectedTeamName:   "Manchester United",
			expectedSeason:     2023,
			expectedLeagueName: "Premier League",
		},
		{
			name: "Empty response",
			statsMetadata: map[string]string{
				"teamId":   "33",
				"leagueId": "39",
				"season":   "2023",
			},
			mockedResponse: `
            {
                "get": "teams/statistics",
                "parameters": {
                    "team": "33",
                    "season": "2023",
                    "league": "39"
                },
                "errors": [],
                "results": 0,
                "paging": {
                    "current": 1,
                    "total": 1
                },
                "response": {}
            }`,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create the mock API client with the mocked response
			mockApiClient := mocks.NewMockApiClient(tt.mockedResponse, nil, true)
			// Create a new instance of the team client with the mock client
			teamClient := dataclients.NewTeamClient(mockApiClient)

			// INFO: Call the function being tested
			result := teamClient.GetTeamSeasonStats(tt.statsMetadata["teamId"], tt.statsMetadata["leagueId"], tt.statsMetadata["season"])

			if tt.expectError {
				assert.Empty(t, result.Response)
			} else {
				assert.NotNil(t, result)
				assert.NotEmpty(t, result.Response)
				assert.Equal(t, tt.expectedTeamID, result.Response.Team.ID)
				assert.Equal(t, tt.expectedTeamName, result.Response.Team.Name)
				assert.Equal(t, tt.expectedSeason, result.Response.League.Season)
				assert.Equal(t, tt.expectedLeagueName, result.Response.League.Name)
			}
		})
	}
}

func TestGetFixtures(t *testing.T) {
	tests := []struct {
		name             string
		statsMetadata    map[string]string
		mockedResponse   string
		expectError      bool
		expectedFixtures []int
	}{
		{
			name: "Valid response",
			statsMetadata: map[string]string{
				"teamId":   "33",
				"leagueId": "39",
				"season":   "2023",
			},
			expectError:      false,
			expectedFixtures: []int{1035046, 1035054},
			mockedResponse: `
            {
    "get": "fixtures",
    "parameters": {
        "team": "33",
        "league": "39",
        "season": "2023"
    },
    "errors": [],
    "results": 38,
    "paging": {
        "current": 1,
        "total": 1
    },
    "response": [
        {
            "fixture": {
                "id": 1035046,
                "referee": "S. Hooper",
                "timezone": "UTC",
                "date": "2023-08-14T19:00:00+00:00",
                "timestamp": 1692039600,
                "periods": {
                    "first": 1692039600,
                    "second": 1692043200
                },
                "venue": {
                    "id": 556,
                    "name": "Old Trafford",
                    "city": "Manchester"
                },
                "status": {
                    "long": "Match Finished",
                    "short": "FT",
                    "elapsed": 90
                }
            },
            "league": {
                "id": 39,
                "name": "Premier League",
                "country": "England",
                "logo": "https://media.api-sports.io/football/leagues/39.png",
                "flag": "https://media.api-sports.io/flags/gb.svg",
                "season": 2023,
                "round": "Regular Season - 1"
            },
            "teams": {
                "home": {
                    "id": 33,
                    "name": "Manchester United",
                    "logo": "https://media.api-sports.io/football/teams/33.png",
                    "winner": true
                },
                "away": {
                    "id": 39,
                    "name": "Wolves",
                    "logo": "https://media.api-sports.io/football/teams/39.png",
                    "winner": false
                }
            },
            "goals": {
                "home": 1,
                "away": 0
            },
            "score": {
                "halftime": {
                    "home": 0,
                    "away": 0
                },
                "fulltime": {
                    "home": 1,
                    "away": 0
                },
                "extratime": {
                    "home": null,
                    "away": null
                },
                "penalty": {
                    "home": null,
                    "away": null
                }
            }
        },
        {
            "fixture": {
                "id": 1035054,
                "referee": "M. Oliver",
                "timezone": "UTC",
                "date": "2023-08-19T16:30:00+00:00",
                "timestamp": 1692462600,
                "periods": {
                    "first": 1692462600,
                    "second": 1692466200
                },
                "venue": {
                    "id": 593,
                    "name": "Tottenham Hotspur Stadium",
                    "city": "London"
                },
                "status": {
                    "long": "Match Finished",
                    "short": "FT",
                    "elapsed": 90
                }
            },
            "league": {
                "id": 39,
                "name": "Premier League",
                "country": "England",
                "logo": "https://media.api-sports.io/football/leagues/39.png",
                "flag": "https://media.api-sports.io/flags/gb.svg",
                "season": 2023,
                "round": "Regular Season - 2"
            },
            "teams": {
                "home": {
                    "id": 47,
                    "name": "Tottenham",
                    "logo": "https://media.api-sports.io/football/teams/47.png",
                    "winner": true
                },
                "away": {
                    "id": 33,
                    "name": "Manchester United",
                    "logo": "https://media.api-sports.io/football/teams/33.png",
                    "winner": false
                }
            },
            "goals": {
                "home": 2,
                "away": 0
            },
            "score": {
                "halftime": {
                    "home": 0,
                    "away": 0
                },
                "fulltime": {
                    "home": 2,
                    "away": 0
                },
                "extratime": {
                    "home": null,
                    "away": null
                },
                "penalty": {
                    "home": null,
                    "away": null
                }
            }
        }
        ]
        }`,
		},
		{
			name: "Empty response",
			statsMetadata: map[string]string{
				"teamId":   "1234",
				"leagueId": "39",
				"season":   "2023",
			},
			expectError:      true,
			expectedFixtures: []int{},
			mockedResponse:   `[]`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create the mock API client with the mocked response
			mockApiClient := mocks.NewMockApiClient(tt.mockedResponse, nil, true)
			// Create a new instance of the team client with the mock client
			teamClient := dataclients.NewTeamClient(mockApiClient)

			// INFO: Call the function being tested
			result := teamClient.GetFixtures(tt.statsMetadata["teamId"], tt.statsMetadata["leagueId"], tt.statsMetadata["season"])

			if tt.expectError {
				assert.Empty(t, result)
			} else {
				assert.NotNil(t, result)
				assert.Equal(t, tt.expectedFixtures, result)
			}
		})
	}
}

// TODO: Implement the following tests
func TestGetFixturesbyId(t *testing.T) {
	tests := []struct {
		name           string
		mockedResponse string
		expectError    bool
		fixtureId      string
	}{
		{
			name:        "Valid response",
			fixtureId:   "1035054",
			expectError: false,
			mockedResponse: `{
                "get": "fixtures",
                "parameters": {
                    "id": "1035054"
                },
                "errors": [],
                "results": 1,
                "paging": {
                    "current": 1,
                    "total": 1
                },
                "response": [
                    {
                        "fixture": {
                            "id": 1035054,
                            "referee": "M. Oliver",
                            "timezone": "UTC",
                            "date": "2023-08-19T16:30:00+00:00",
                            "timestamp": 1692462600,
                            "periods": {
                                "first": 1692462600,
                                "second": 1692466200
                            },
                            "venue": {
                                "id": 593,
                                "name": "Tottenham Hotspur Stadium",
                                "city": "London"
                            },
                            "status": {
                                "long": "Match Finished",
                                "short": "FT",
                                "elapsed": 90
                            }
                        },
                        "league": {
                            "id": 39,
                            "name": "Premier League",
                            "country": "England",
                            "logo": "https://media.api-sports.io/football/leagues/39.png",
                            "flag": "https://media.api-sports.io/flags/gb.svg",
                            "season": 2023,
                            "round": "Regular Season - 2"
                        },
                        "teams": {
                            "home": {
                                "id": 47,
                                "name": "Tottenham",
                                "logo": "https://media.api-sports.io/football/teams/47.png",
                                "winner": true
                            },
                            "away": {
                                "id": 33,
                                "name": "Manchester United",
                                "logo": "https://media.api-sports.io/football/teams/33.png",
                                "winner": false
                            }
                        },
                        "goals": {
                            "home": 2,
                            "away": 0
                        },
                        "score": {
                            "halftime": {
                                "home": 0,
                                "away": 0
                            },
                            "fulltime": {
                                "home": 2,
                                "away": 0
                            },
                            "extratime": {
                                "home": null,
                                "away": null
                            },
                            "penalty": {
                                "home": null,
                                "away": null
                            }
                        },
                        "events": [],
                        "lineups": [],
                        "statistics": [],
                        "players": [
                            {
                                "team": {
                                    "id": 47,
                                    "name": "Tottenham",
                                    "logo": "https://media.api-sports.io/football/teams/47.png",
                                    "update": "2024-05-21T04:12:18+00:00"
                                },
                                "players": []
                            },
                            {
                                "team": {
                                    "id": 33,
                                    "name": "Manchester United",
                                    "logo": "https://media.api-sports.io/football/teams/33.png",
                                    "update": "2024-05-21T04:12:18+00:00"
                                },
                                "players": []
                            }
                        ]
                    }
                ]
            }`,
		},
		{
			name:           "Empty response",
			fixtureId:      "542362632cfa",
			mockedResponse: `[]`,
			expectError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create the mock API client with the mocked response
			mockApiClient := mocks.NewMockApiClient(tt.mockedResponse, nil, true)
			// Create a new instance of the team client with the mock client
			teamClient := dataclients.NewTeamClient(mockApiClient)

			// INFO: Call the GetTeamStats method
			result := teamClient.GetFixturebyId(tt.fixtureId)

			if tt.expectError {
				assert.Empty(t, result)
			} else {
				assert.NotNil(t, result)
				assert.NotEmpty(t, result.Response)
			}

		})
	}
}
func TestGetFixturebyIds(t *testing.T) {
	tests := []struct {
		name           string
		mockedResponse string
		expectError    bool
		fixtureIds     string
	}{
		{
			name:        "Valid response",
			expectError: false,

			fixtureIds: "1035090-103510",
			mockedResponse: `{
                "get": "fixtures",
                "parameters": {
                    "ids": "1035090-103510"
                },
                "errors": [],
                "results": 2,
                "paging": {
                    "current": 1,
                    "total": 1
                },
                "response": [
                    {
                        "fixture": {
                            "id": 103510,
                            "referee": null,
                            "timezone": "UTC",
                            "date": "2019-06-16T11:00:00+00:00",
                            "timestamp": 1560682800,
                            "periods": {
                                "first": 1560682800,
                                "second": 1560686400
                            },
                            "venue": {
                                "id": null,
                                "name": "Thammasat Stadium (Pathum Thani)",
                                "city": null
                            },
                            "status": {
                                "long": "Match Finished",
                                "short": "FT",
                                "elapsed": 90
                            }
                        },
                        "league": {
                            "id": 296,
                            "name": "Thai League 1",
                            "country": "Thailand",
                            "logo": "https://media.api-sports.io/football/leagues/296.png",
                            "flag": "https://media.api-sports.io/flags/th.svg",
                            "season": 2019,
                            "round": "Regular Season - 14"
                        },
                        "teams": {
                            "home": {
                                "id": 2770,
                                "name": "Bangkok United",
                                "logo": "https://media.api-sports.io/football/teams/2770.png",
                                "winner": true
                            },
                            "away": {
                                "id": 2776,
                                "name": "Ratchaburi",
                                "logo": "https://media.api-sports.io/football/teams/2776.png",
                                "winner": false
                            }
                        }
                    }
                ]
            }`,
		},
		{
			name:           "Empty response",
			fixtureIds:     "",
			expectError:    true,
			mockedResponse: `[]`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create the mock API client with the mocked response
			mockApiClient := mocks.NewMockApiClient(tt.mockedResponse, nil, true)
			// Create a new instance of the team client with the mock client
			teamClient := dataclients.NewTeamClient(mockApiClient)

			// INFO: Call the GetTeamStats method
			result := teamClient.GetFixturebyIds(tt.fixtureIds)

			if tt.expectError {
				assert.Empty(t, result)
			} else {
				fmt.Println("result", result)
				assert.NotNil(t, result)
				assert.Equal(t, 2, result.Results)
			}

		})
	}
}

func TestGetFixtureStats(t *testing.T) {
	tests := []struct {
		name           string
		mockedResponse string
		expectError    bool
		fixtureId      string
		teamId         string
	}{
		{
			name:        "Valid response",
			expectError: false,
			fixtureId:   "1035054",
			teamId:      "33",
			mockedResponse: `{
                "get": "fixtures/statistics",
                "parameters": {
                    "fixture": "1035054",
                    "team": "33"
                },
                "errors": [],
                "results": 1,
                "paging": {
                    "current": 1,
                    "total": 1
                },
                "response": [
                    {
                        "team": {
                            "id": 33,
                            "name": "Manchester United",
                            "logo": "https://media.api-sports.io/football/teams/33.png"
                        },
                        "statistics": [
                            {
                                "type": "Shots on Goal",
                                "value": 6
                            },
                            {
                                "type": "Shots off Goal",
                                "value": 7
                            },
                            {
                                "type": "Total Shots",
                                "value": 22
                            },
                            {
                                "type": "Blocked Shots",
                                "value": 9
                            },
                            {
                                "type": "Shots insidebox",
                                "value": 16
                            },
                            {
                                "type": "Shots outsidebox",
                                "value": 6
                            },
                            {
                                "type": "Fouls",
                                "value": 8
                            },
                            {
                                "type": "Corner Kicks",
                                "value": 6
                            },
                            {
                                "type": "Offsides",
                                "value": 4
                            },
                            {
                                "type": "Ball Possession",
                                "value": "44%"
                            },
                            {
                                "type": "Yellow Cards",
                                "value": 3
                            },
                            {
                                "type": "Red Cards",
                                "value": null
                            },
                            {
                                "type": "Goalkeeper Saves",
                                "value": 5
                            },
                            {
                                "type": "Total passes",
                                "value": 401
                            },
                            {
                                "type": "Passes accurate",
                                "value": 332
                            },
                            {
                                "type": "Passes %",
                                "value": "83%"
                            },
                            {
                                "type": "expected_goals",
                                "value": "2.07"
                            }
                        ]
                    }
                ]
            }`,
		},
		{
			name:           "Empty response",
			fixtureId:      "",
			teamId:         "",
			mockedResponse: `[]`,
			expectError:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create the mock API client with the mocked response
			mockApiClient := mocks.NewMockApiClient(tt.mockedResponse, nil, true)
			// Create a new instance of the team client with the mock client
			teamClient := dataclients.NewTeamClient(mockApiClient)

			// INFO: Call the GetTeamStats method
			result := teamClient.GetFixtureStats(tt.teamId, tt.fixtureId)

			if tt.expectError {
				assert.Empty(t, result)
			} else {
				assert.NotNil(t, result)
				assert.NotEmpty(t, result.Response[0].Statistics)
			}
		})
	}
}

func TestGetTeams(t *testing.T) {
	tests := []struct {
		name            string
		expectError     bool
		leagueId        string
		season          string
		expectedResults int
		expectedTeamId  int
		expectedVenueId int
		mockedResponse  string
	}{
		{
			name:            "Valid response",
			expectError:     false,
			leagueId:        "39", // Premien League
			season:          "2023",
			expectedResults: 20,
			expectedTeamId:  33,
			expectedVenueId: 556,
			mockedResponse: `{
                "get": "teams",
                "parameters": {
                    "league": "39",
                    "season": "2023"
                },
                "errors": [],
                "results": 20,
                "paging": {
                    "current": 1,
                    "total": 1
                },
                "response": [
                    {
                        "team": {
                            "id": 33,
                            "name": "Manchester United",
                            "code": "MUN",
                            "country": "England",
                            "founded": 1878,
                            "national": false,
                            "logo": "https://media.api-sports.io/football/teams/33.png"
                        },
                        "venue": {
                            "id": 556,
                            "name": "Old Trafford",
                            "address": "Sir Matt Busby Way",
                            "city": "Manchester",
                            "capacity": 76212,
                            "surface": "grass",
                            "image": "https://media.api-sports.io/football/venues/556.png"
                        }
                    },
                    {
                        "team": {
                            "id": 34,
                            "name": "Newcastle",
                            "code": "NEW",
                            "country": "England",
                            "founded": 1892,
                            "national": false,
                            "logo": "https://media.api-sports.io/football/teams/34.png"
                        },
                        "venue": {
                            "id": 562,
                            "name": "St. James' Park",
                            "address": "St. James&apos; Street",
                            "city": "Newcastle upon Tyne",
                            "capacity": 52758,
                            "surface": "grass",
                            "image": "https://media.api-sports.io/football/venues/562.png"
                        }
                    }
                ]
            }`,
		},
		{
			name:            "Empty response - Season with no data coverage",
			leagueId:        "39",
			expectError:     true,
			season:          "1990",
			expectedResults: 0,
			expectedTeamId:  0,
			expectedVenueId: 0,
			mockedResponse: `{
                "get": "teams",
                "parameters": {
                    "league": "39",
                    "season": "1990"
                },
                "errors": [],
                "results": 0,
                "paging": {
                    "current": 1,
                    "total": 1
                },
                "response": []
            }`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create the mock API client with the mocked response
			mockApiClient := mocks.NewMockApiClient(tt.mockedResponse, nil, true)
			// Create a new instance of the team client with the mock client
			teamClient := dataclients.NewTeamClient(mockApiClient)

			// INFO: Call the GetTeams method
			result := teamClient.GetTeams(tt.leagueId, tt.season)

			if tt.expectError {
				assert.Empty(t, result.Response)
				assert.Equal(t, tt.expectedResults, result.Results)
				assert.Empty(t, tt.expectedTeamId, result.Response)
			} else {
				assert.NotNil(t, result)
				assert.Equal(t, tt.expectedResults, result.Results)
				assert.Equal(t, tt.expectedTeamId, result.Response[0].Team.ID)
				assert.Equal(t, tt.expectedVenueId, result.Response[0].Venue.ID)
			}
		})
	}
}
