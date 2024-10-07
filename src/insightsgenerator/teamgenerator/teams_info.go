package teamgenerator

import (
	"insights-pulse/src/logger"
	"insights-pulse/src/models"
	"insights-pulse/src/models/insights/teaminsights"
	"time"
)

type TeamsInfoGenerator struct {
	*InsightGeneratorBase
}

func (h *TeamsInfoGenerator) GetConfig() InsightConfig {
	return InsightConfig{
		Type:            "TeamsInfoGenerator",
		TableName:       "teams",
		Api:             "https://v3.football.api-sports.io",
		Endpoints:       []string{"/teams"},
		UpdateFrequency: 24 * 365 * time.Hour, // Once a year
	}
}

func (h *TeamsInfoGenerator) GenerateAndSaveInsights(imeta teaminsights.StatsMetaData) error {
	var teamsInfo []models.TeamInfo

	teamsResponse := h.TeamClient.GetTeams(imeta.LeagueId, imeta.Season)
	for _, team := range teamsResponse.Response {
		teamI := models.NewTeamInfo()
		teamI.TeamId = team.Team.ID
		teamI.Name = team.Team.Name
		teamI.Code = team.Team.Code
		teamI.Country = team.Team.Country
		teamI.Founded = team.Team.Founded
		teamI.National = team.Team.National
		teamI.Logo = team.Team.Logo
		teamI.VenueId = team.Venue.ID
		teamI.VenueName = team.Venue.Name
		teamI.VenueAddress = team.Venue.Address
		teamI.VenueCity = team.Venue.City
		teamI.VenueCapacity = team.Venue.Capacity
		teamI.VenueSurface = team.Venue.Surface
		teamI.VenueImage = team.Venue.Image
		teamsInfo = append(teamsInfo, *teamI)
	}

	if len(teamsInfo) == 0 {
		logger.GetLogger().Error("No teams found")
		return nil
	}

	err := h.TeamRepo.SaveTeam(teamsInfo)
	if err != nil {
		logger.GetLogger().Error("Error saving team: " + err.Error())
		return err
	}

	return nil
}
