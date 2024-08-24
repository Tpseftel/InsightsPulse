package sqlrepo

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"insights-pulse/src/logger"
	"insights-pulse/src/models"
	"insights-pulse/src/models/insights"
)

const timeLayout = "2006-01-02 15:04:05"

type TeamRepository struct {
	Conn *sql.DB
}

func NewTeamRepository(db *sql.DB) *TeamRepository {
	return &TeamRepository{Conn: db}
}

func (repo *TeamRepository) SaveTeam(t models.Team) error {
	query := `
	INSERT INTO teams (id, name, code, country, founded, national, logo)
	VALUES (?, ?, ?, ?, ?, ?, ?)
	`
	stmt, err := repo.Conn.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(t.ID, t.Name, t.Code, t.Country, t.Founded, t.National, t.Logo)
	if err != nil {
		return err
	}
	logger.GetLogger().Info("Saved Team")
	return err
}

// INFO: Insights related methods
func (repo *TeamRepository) SaveAvgInsightsPerGame(meta insights.StatsMetaData, avgInsights *insights.MatchMetrics) error {
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

func (repo *TeamRepository) GetLastUpdatedTime(tableName string) (time.Time, error) {
	query := `
		SELECT updated_at
		FROM ` + tableName + `
		ORDER BY updated_at DESC
		LIMIT 1
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
