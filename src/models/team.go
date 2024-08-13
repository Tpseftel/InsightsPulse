package models

// Team represents a sports team
type Team struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Code     string `json:"code"`
	Country  string `json:"country"`
	Founded  int    `json:"founded"`
	National bool   `json:"national"`
	Logo     string `json:"logo"`
}

type TeamWithTimestamps struct {
	Team
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
