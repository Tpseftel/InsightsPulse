package dataclients

import (
	"errors"

	"insights-pulse/src/apiclients"

	"insights-pulse/src/logger"
	"insights-pulse/src/models/responses"
)

var log logger.Logger

type TeamClient struct {
	apiClient apiclients.ApiClient
}

func NewTeamClient(apiClient apiclients.ApiClient) *TeamClient {
	log = logger.GetLogger()
	return &TeamClient{apiClient}
}

func (c *TeamClient) GetTeamSeasonStats(teamId, leagueId, season string) *responses.TeamStatsResponse {
	endpoint := "/teams/statistics"
	queryParams := map[string]string{
		"team":   teamId,
		"season": season,
		"league": leagueId,
	}
	var dataResponse responses.TeamStatsResponse

	resp, err := c.apiClient.GetClient().
		R().
		SetResult(&dataResponse).
		SetQueryParams(queryParams).
		Get(endpoint)
	if err != nil {
		log.Warn(err.Error())
		return nil
	}
	if resp.IsError() {
		log.Warn(errors.New("api responde with error status code >=400").Error())
		return nil
	}

	return &dataResponse
}

func (c *TeamClient) GetFixtures(teamId, leagueId, season string) []int {
	endpoint := "/fixtures"
	queryParams := map[string]string{
		"team":   teamId,
		"season": season,
		"league": leagueId,
	}

	var dataResponse responses.TeamFixturesIResponse
	resp, err := c.apiClient.GetClient().
		R().
		SetResult(&dataResponse).
		SetQueryParams(queryParams).
		Get(endpoint)
	if err != nil {
		log.Warn(err.Error())
		return []int{}
	}

	if resp.IsError() {
		log.Warn("endpoint: " + endpoint + "\n api responde with error status code >=400")
		return []int{}
	}

	if dataResponse.Results == 0 {
		log.Warn("endpoint: " + endpoint + "\n no results")
		return []int{}
	}

	fixturesIds := make([]int, 0, dataResponse.Results)
	for _, v := range dataResponse.Response {
		fixturesIds = append(fixturesIds, v.Fixture.ID)
	}

	return fixturesIds
}

func (c *TeamClient) GetFixturebyId(fixtureId string) *responses.FixtureStatsResponse {
	endpoint := "/fixtures"
	queryParams := map[string]string{
		"id": fixtureId,
	}
	var dataResponse responses.FixtureStatsResponse
	resp, err := c.apiClient.GetClient().
		R().
		SetResult(&dataResponse).
		SetQueryParams(queryParams).
		Get(endpoint)
	if err != nil {
		log.Warn(err.Error())
		return nil
	}

	if resp.IsError() {
		log.Warn("endpoint: " + endpoint + "\n api responde with error status code >=400")
		return nil
	}

	if dataResponse.Results == 0 {
		log.Warn("endpoint: " + endpoint + "\n no results")
		return nil
	}

	return &dataResponse
}

// Ids must be in the form of "id1-id2-id3"
//
//	Max number of ids is 20
func (c *TeamClient) GetFixturebyIds(stringIds string) *responses.FixtureStatsResponse {
	endpoint := "/fixtures"
	queryParams := map[string]string{
		"ids": stringIds, // "3442-4124-43243"
	}
	var dataResponse responses.FixtureStatsResponse
	resp, err := c.apiClient.GetClient().
		R().
		SetResult(&dataResponse).
		SetQueryParams(queryParams).
		Get(endpoint)
	if err != nil {
		log.Warn(err.Error())
		return nil
	}

	if resp.IsError() {
		log.Warn("endpoint: " + endpoint + "\n api responde with error status code >=400")
		return nil
	}

	if dataResponse.Results == 0 {
		log.Warn("endpoint: " + endpoint + "\n no results")
		return nil
	}

	return &dataResponse
}

func (c *TeamClient) GetFixtureStats(teamId, fixtureId string) *responses.FixtureTeamStatsResponse {
	endpoint := "/fixtures/statistics"
	queryParams := map[string]string{
		"fixture": fixtureId,
		"team":    teamId,
	}
	var dataResponse responses.FixtureTeamStatsResponse

	resp, err := c.apiClient.GetClient().
		R().
		SetResult(&dataResponse).
		SetQueryParams(queryParams).
		Get(endpoint)
	if err != nil {
		log.Warn(err.Error())
		return nil
	}
	if resp.IsError() {
		log.Warn(errors.New("api responde with error status code >=400").Error())
		return nil
	}
	return &dataResponse
}
