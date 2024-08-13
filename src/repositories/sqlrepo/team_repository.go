package sqlrepo

import (
	"database/sql"
	"encoding/json"

	"insights-pulse/src/logger"
	"insights-pulse/src/models"
	"insights-pulse/src/models/insights"
)

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