package models

type TeamInfo struct {
	TeamId        int    `json:"team_id"`
	Name          string `json:"name"`
	Code          string `json:"code"`
	Country       string `json:"country"`
	Founded       int    `json:"founded"`
	National      bool   `json:"national"`
	Logo          string `json:"logo"`
	VenueId       int    `json:"venue_id"`
	VenueName     string `json:"venue_name"`
	VenueAddress  string `json:"venue_address"`
	VenueCity     string `json:"venue_city"`
	VenueCapacity int    `json:"venue_capacity"`
	VenueSurface  string `json:"venue_surface"`
	VenueImage    string `json:"venue_image"`
	UpdatedAt     string `json:"updated_at"`
	CreatedAt     string `json:"created_at"`
}

func NewTeamInfo() *TeamInfo {
	return &TeamInfo{}
}
