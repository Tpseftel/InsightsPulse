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
		log.Error("Cannot load Configuration variables: " + err.Error())
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
	teamsTable := `
	CREATE TABLE IF NOT EXISTS teams (
		id INTEGER PRIMARY KEY,
        name TEXT NOT NULL,
        code TEXT NOT NULL,
        country TEXT NOT NULL,
        founded INTEGER NOT NULL,
        national BOOLEAN NOT NULL,
        logo TEXT,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
    );`
	_, err := DB.Exec(teamsTable)
	if err != nil {
		log.Error("Create table Error: " + err.Error())
		panic(err)
	}
	log.Info("Table teams created successfully")

	usersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		username VARCHAR(50) NOT NULL UNIQUE,
		password_hash VARCHAR(255) NOT NULL,
		email VARCHAR(100) NOT NULL UNIQUE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	);`
	_, err = DB.Exec(usersTable)
	if err != nil {
		log.Error("Create table Error: " + err.Error())
		panic(err)
	}
	log.Info("Table users created successfully")

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
		PRIMARY KEY (team_id, season, league_id)
	);`
	_, err = DB.Exec(avgInsightsPerGameTeam)
	if err != nil {
		log.Error("Create table Error: " + err.Error())
		panic(err)
	}
	log.Info("Table avg_insights_per_game_team created successfully")

}
