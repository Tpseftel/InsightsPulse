package db

import (
	"database/sql"
	"fmt"

	"insights-pulse/src/config"
	"insights-pulse/src/logger"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

var log logger.Logger

func InitDb() {
	log = logger.GetLogger()

	config, err := config.GetConfig()
	if err != nil {
		logger.GetLogger().Error("Cannot load Configuration variables: " + err.Error())
		panic(err)
	}
	user := config.User
	password := config.Password
	dbName := config.DbName
	host := config.Host
	port := config.Port

	// Connection string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbName)

	// Open a database connection
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Error("Error opening database: " + err.Error())
		panic(err)
	}

	// Ping the database to verify the connection
	if err := DB.Ping(); err != nil {
		log.Error("Error pinging database: " + err.Error())
		panic(err)
	}

	log.Info("Successfuly connected to database!!")

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {

	// Teams Table
	teamsTable := `CREATE TABLE IF NOT EXISTS teams (
		id INT PRIMARY KEY,
		name VARCHAR(255),
		code VARCHAR(255),
		country VARCHAR(255),
		founded INT,
		national BOOLEAN,
		logo VARCHAR(255),
		venue_id INT,
		venue_name VARCHAR(255),
		venue_surface VARCHAR(255),
		venue_address VARCHAR(255),
		venue_city VARCHAR(255),
		venue_capacity INT,
		venue_image VARCHAR(255),
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	);`
	_, err := DB.Exec(teamsTable)
	if err != nil {
		log.Error("Create table Error: " + err.Error())
		panic(err)
	}

	// Insights Tables
	avgInsightsPerGameTeam := `CREATE TABLE IF NOT EXISTS avg_insights_per_game_team (
		team_id VARCHAR(255),
		season VARCHAR(255),
		league_id VARCHAR(255),
		shots_on_goal JSON,
		shots_off_goal JSON,
		total_shots JSON,
		blocked_shots JSON,
		shots_inside_box JSON,
		shots_outside_box JSON,
		fouls JSON,
		corner_kicks JSON,
		offsides JSON,
		ball_possession JSON,
		yellow_cards JSON,
		red_cards JSON,
		goalkeeper_saves JSON,
		total_passes JSON,
		passes_accurate JSON,
		passes_percentage JSON,
		expected_goals JSON,
		PRIMARY KEY (team_id, season, league_id),
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	);`
	_, err = DB.Exec(avgInsightsPerGameTeam)
	if err != nil {
		log.Error("Create table Error: " + err.Error())
		panic(err)
	}
	log.Info("Table avg_insights_per_game_team created successfully")

	homeAwayMetrics := `CREATE TABLE IF NOT EXISTS home_away_metrics (
		team_id VARCHAR(255),
		season VARCHAR(255),
		league_id VARCHAR(255),
		fixtures JSON,
		wins JSON,
		draws JSON,
		loses JSON,
		goals_for_total JSON,
		goals_for_average JSON,
		goals_for_minute JSON,
		goals_against_total JSON,
		goals_against_average JSON,
		goals_against_minute JSON,
		clean_sheets JSON,
		failed_to_score JSON,
		points_per_game JSON,
		PRIMARY KEY (team_id, season, league_id),
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	);`
	_, err = DB.Exec(homeAwayMetrics)
	if err != nil {
		log.Error("Create table Error: " + err.Error())
		panic(err)
	}
	log.Info("Table home_away_metrics created successfully")
}
