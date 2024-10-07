package sqlrepo

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"insights-pulse/src/logger"
	"insights-pulse/src/models"
	"insights-pulse/src/models/insights/teaminsights"
)

const timeLayout = "2006-01-02 15:04:05"

type TeamRepository struct {
	Conn *sql.DB
}

func NewTeamRepository(db *sql.DB) *TeamRepository {
	return &TeamRepository{Conn: db}
}

func (repo *TeamRepository) SaveTeam(t []models.TeamInfo) error {
	query := `
		INSERT INTO teams (
			team_id, name, code, country, founded, national, logo, venue_id, venue_name, venue_address, venue_city, venue_capacity, venue_surface, venue_image 
		) VALUES `

	values := []interface{}{}
	for i, team := range t {
		query += "(?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
		if i < len(t)-1 {
			query += ","
		} else {
			query += ";"
		}
		values = append(values,
			team.TeamId,
			team.Name,
			team.Code,
			team.Country,
			team.Founded,
			team.National,
			team.Logo,
			team.VenueId,
			team.VenueName,
			team.VenueAddress,
			team.VenueCity,
			team.VenueCapacity,
			team.VenueSurface,
			team.VenueImage,
		)

	}

	_, err := repo.Conn.Exec(query, values...)
	if err != nil {
		logger.GetLogger().Error("Error saving team: " + err.Error())
		return err
	}
	logger.GetLogger().Info("Saved Team")
	return err
}

// INFO: Insights related methods
func (repo *TeamRepository) SaveAvgInsightsPerGame(meta teaminsights.StatsMetaData, avgInsights *teaminsights.MatchMetrics) error {
	query := `
		INSERT INTO avg_insights_per_game_team (
			team_id, season, league_id, shots_on_goal, shots_off_goal, total_shots, blocked_shots, shots_inside_box, shots_outside_box, fouls, corner_kicks, offsides, ball_possession, yellow_cards, red_cards, goalkeeper_saves, total_passes, passes_accurate, passes_percentage, expected_goals
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	 `
	stmt, err := repo.Conn.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	// INFO: Serialize the struct fields to JSON
	shotsOnGoal, _ := json.Marshal(avgInsights.ShotsOnGoal)
	shotsOffGoal, _ := json.Marshal(avgInsights.ShotsOffGoal)
	totalShots, _ := json.Marshal(avgInsights.TotalShots)
	blockedShots, _ := json.Marshal(avgInsights.BlockedShots)
	shotsInsideBox, _ := json.Marshal(avgInsights.ShotsInsideBox)
	shotsOutsideBox, _ := json.Marshal(avgInsights.ShotsOutsideBox)
	fouls, _ := json.Marshal(avgInsights.Fouls)
	cornerKicks, _ := json.Marshal(avgInsights.CornerKicks)
	offsides, _ := json.Marshal(avgInsights.Offsides)
	ballPossession, _ := json.Marshal(avgInsights.BallPossession)
	yellowCards, _ := json.Marshal(avgInsights.YellowCards)
	redCards, _ := json.Marshal(avgInsights.RedCards)
	goalkeeperSaves, _ := json.Marshal(avgInsights.GoalkeeperSaves)
	totalPasses, _ := json.Marshal(avgInsights.TotalPasses)
	passesAccurate, _ := json.Marshal(avgInsights.PassesAccurate)
	passesPercentage, _ := json.Marshal(avgInsights.PassesPercentage)
	expectedGoals, _ := json.Marshal(avgInsights.ExpectedGoals)

	_, err = stmt.Exec(meta.TeamId, meta.Season, meta.LeagueId, shotsOnGoal, shotsOffGoal, totalShots, blockedShots, shotsInsideBox, shotsOutsideBox, fouls, cornerKicks, offsides, ballPossession, yellowCards, redCards, goalkeeperSaves, totalPasses, passesAccurate, passesPercentage, expectedGoals)
	if err != nil {
		return err
	}
	logger.GetLogger().Info("Saved Average Insights Per Game")
	return err
}

func (repo *TeamRepository) SaveHomeAwayMetrics(meta teaminsights.StatsMetaData, homeAwayMetrics *teaminsights.HomeAwayMetrics) error {
	query := `
		INSERT INTO home_away_metrics (
			team_id, season, league_id, fixtures, wins, draws, loses, goals_for_total, goals_for_average, goals_for_minute, goals_against_total, goals_against_average, goals_against_minute, clean_sheets, failed_to_score, points_per_game
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	stm, err := repo.Conn.Prepare(query)
	if err != nil {
		return err
	}
	defer stm.Close()

	// INFO: Serialize the struct fields to JSON
	fixtures, _ := json.Marshal(homeAwayMetrics.Fixtures)
	wins, _ := json.Marshal(homeAwayMetrics.Wins)
	draws, _ := json.Marshal(homeAwayMetrics.Draws)
	loses, _ := json.Marshal(homeAwayMetrics.Loses)
	goalsForTotal, _ := json.Marshal(homeAwayMetrics.GoalsForTotal)
	goalsForAverage, _ := json.Marshal(homeAwayMetrics.GoalsForAverage)
	goalsForMinute, _ := json.Marshal(homeAwayMetrics.GoalsForMinute)
	goalsAgainstTotal, _ := json.Marshal(homeAwayMetrics.GoalsAgainstTotal)
	goalsAgainstAverage, _ := json.Marshal(homeAwayMetrics.GoalsAgainstAverage)
	goalsAgainstMinute, _ := json.Marshal(homeAwayMetrics.GoalsAgainstMinute)
	cleanSheets, _ := json.Marshal(homeAwayMetrics.CleanSheets)
	failedToScore, _ := json.Marshal(homeAwayMetrics.FailedToScore)
	pointsPerGame, _ := json.Marshal(homeAwayMetrics.PointsPerGame)

	_, err = stm.Exec(meta.TeamId, meta.Season, meta.LeagueId, fixtures, wins, draws, loses, goalsForTotal, goalsForAverage, goalsForMinute, goalsAgainstTotal, goalsAgainstAverage, goalsAgainstMinute, cleanSheets, failedToScore, pointsPerGame)
	if err != nil {
		return err
	}
	logger.GetLogger().Info("Saved Home Away Metrics")
	return err

}

func (repo *TeamRepository) GetLastUpdatedTime(tableName, leaugeId string) (time.Time, error) {
	query := `
		SELECT Max(updated_at) as last_updated
		FROM ` + tableName + `
		WHERE league_id = '` + leaugeId + `'
	`
	var updateTimeBytes []byte
	err := repo.Conn.QueryRow(query).Scan(&updateTimeBytes)
	if err != nil {
		if err == sql.ErrNoRows {
			return time.Time{}, nil
		}
		logger.GetLogger().Error("Error getting last updated time: " + err.Error())
		return time.Time{}, errors.New("error getting last updated time")
	}
	updatedAtStr := string(updateTimeBytes)
	updateTime, err := time.Parse(timeLayout, updatedAtStr)
	if err != nil {
		logger.GetLogger().Error("Error parsing time: " + err.Error())
		return time.Time{}, errors.New("error parsing time")
	}

	if updateTime.IsZero() {
		logger.GetLogger().Error("No data found")
		return time.Time{}, errors.New("no data found")
	}

	logger.GetLogger().Info(fmt.Sprintf("Table: %s, Last Updated Time: %s", tableName, updateTime))

	return updateTime, nil

}
