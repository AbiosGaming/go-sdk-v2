package structs

// PaginatedPlayers holds a list of Player as well as information about pages.
type PaginatedPlayers struct {
	LastPage    int64    `json:"last_page,omitempty"`
	CurrentPage int64    `json:"current_page,omitempty"`
	Data        []Player `json:"data,omitempty"`
}

// Player represents a player that competes in Series' and Matches.
type Player struct {
	Id                  int64                `json:"id,omitempty"`
	FirstName           string               `json:"first_name"`
	LastName            string               `json:"last_name"`
	Nickname            string               `json:"nick_name,omitempty"`
	DeletedAt           *string              `json:"deleted_at"`
	Images              PlayerImages         `json:"images,omitempty"`
	Country             *Country             `json:"country,omitempty"`
	Roles               []Role               `json:"roles"`
	Race                *Race                `json:"race,omitempty"`
	Team                *Team                `json:"team,omitempty"`
	PlayerStats         PlayerStats          `json:"player_stats,omitempty"`
	Rosters             []DefaultRoster      `json:"rosters,omitempty"`
	Game                Game                 `json:"game,omitempty"`
	SocialMediaAccounts []SocialMediaAccount `json:"social_media_accounts,omitempty"`
}
