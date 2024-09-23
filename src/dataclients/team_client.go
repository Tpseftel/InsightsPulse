package dataclients

import (
	"encoding/json"
	"errors"

	"insights-pulse/src/apiclients"

	"insights-pulse/src/logger"
	"insights-pulse/src/models/responses"
)

type TeamClient struct {
	apiClient apiclients.ApiClient
}

func NewTeamClient(apiClient apiclients.ApiClient) *TeamClient {
	return &TeamClient{apiClient}
}

func (c *TeamClient) GetTeamSeasonStats(teamId, leagueId, season string) *responses.TeamStatsResponse {
	endpoint := "/teams/statistics"
	queryParams := map[string]string{
		"team":   teamId,
		"season": season,
		"league": leagueId,
	}

	resp, err := c.apiClient.GetClient().
		R().
		SetQueryParams(queryParams).
		Get(endpoint)

	if err != nil {
		logger.GetLogger().Warn(err.Error())
		return nil
	}

	if resp.IsError() {
		logger.GetLogger().Warn(errors.New("api responde with error status code >=400").Error())
		return nil
	}

	// Unmarshal the response body to the response struct
	var dataResponse responses.TeamStatsResponse
	err = json.Unmarshal(resp.Body(), &dataResponse)
	if err != nil {
		logger.GetLogger().Warn(err.Error())
		logger.GetLogger().Warn(resp.String())

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

	resp, err := c.apiClient.GetClient().
		R().
		SetQueryParams(queryParams).
		Get(endpoint)
	if err != nil {
		logger.GetLogger().Warn(err.Error())
		return []int{}
	}

	if resp.IsError() {
		logger.GetLogger().Warn("endpoint: " + endpoint + "\n api responde with error status code >=400")
		return []int{}
	}

	var dataResponse responses.TeamFixturesIResponse
	err = json.Unmarshal(resp.Body(), &dataResponse)
	if err != nil {
		logger.GetLogger().Warn(err.Error())
		return []int{}
	}

	if dataResponse.Results == 0 {
		logger.GetLogger().Warn("endpoint: " + endpoint + "\n no results")
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
	resp, err := c.apiClient.GetClient().
		R().
		SetQueryParams(queryParams).
		Get(endpoint)
	if err != nil {
		logger.GetLogger().Warn(err.Error())
		return nil
	}

	if resp.IsError() {
		logger.GetLogger().Warn("endpoint: " + endpoint + "\n api responde with error status code >=400")
		return nil
	}

	var dataResponse responses.FixtureStatsResponse
	err = json.Unmarshal(resp.Body(), &dataResponse)
	if err != nil {
		logger.GetLogger().Warn(err.Error())
		return nil
	}

	if dataResponse.Results == 0 {
		logger.GetLogger().Warn("endpoint: " + endpoint + "\n no results")
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
	resp, err := c.apiClient.GetClient().
		R().
		SetQueryParams(queryParams).
		Get(endpoint)
	if err != nil {
		logger.GetLogger().Warn(err.Error())
		return nil
	}

	if resp.IsError() {
		logger.GetLogger().Warn("endpoint: " + endpoint + "\n api responde with error status code >=400")
		return nil
	}

	var dataResponse responses.FixtureStatsResponse
	err = json.Unmarshal(resp.Body(), &dataResponse)
	if err != nil {
		logger.GetLogger().Warn(err.Error())
		return nil
	}
	if dataResponse.Results == 0 {
		logger.GetLogger().Warn("endpoint: " + endpoint + "\n no results")
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

	resp, err := c.apiClient.GetClient().
		R().
		SetQueryParams(queryParams).
		Get(endpoint)
	if err != nil {
		logger.GetLogger().Warn(err.Error())
		return nil
	}
	if resp.IsError() {
		logger.GetLogger().Warn(errors.New("api responde with error status code >=400").Error())
		return nil
	}
	var dataResponse responses.FixtureTeamStatsResponse
	err = json.Unmarshal(resp.Body(), &dataResponse)
	if err != nil {
		logger.GetLogger().Warn(err.Error())
		return nil
	}

	return &dataResponse
}

func (c *TeamClient) GetTeams(leagueId, season string) *responses.TeamResponse {
	endpoint := "/teams"
	queryParams := map[string]string{
		"league": leagueId,
		"season": season,
	}

	resp, err := c.apiClient.GetClient().
		R().
		SetQueryParams(queryParams).
		Get(endpoint)
	if err != nil {
		logger.GetLogger().Warn(err.Error())
		return nil
	}
	if resp.IsError() {
		logger.GetLogger().Warn(errors.New("api responde with error status code >=400").Error())
		return nil
	}
	var dataResponse responses.TeamResponse
	err = json.Unmarshal(resp.Body(), &dataResponse)
	if err != nil {
		logger.GetLogger().Warn(err.Error())
		return nil
	}

	return &dataResponse

}
