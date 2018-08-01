package structs

// PlayerStructPaginated holds a list of PlayerStruct as well as information about pages.
type PlayerStructPaginated struct {
	LastPage    int64          `json:"last_page,omitempty"`
	CurrentPage int64          `json:"current_page,omitempty"`
	Data        []PlayerStruct `json"data,omitempty"`
}

// PlayerStruct represents a player that competes in Series' and Matches.
type PlayerStruct struct {
	Id                  int64                      `json:"id,omitempty"`
	FirstName           string                     `json:"first_name"`
	LastName            string                     `json:"last_name"`
	Nickname            string                     `json:"nick_name,omitempty"`
	DeletedAt           *string                    `json:"deleted_at"`
	Images              PlayerImagesStruct         `json:"images,omitempty"`
	Country             *CountryStruct             `json:"country,omitempty"`
	Race                *RaceStruct                `json:"race,omitempty"`
	Team                *TeamStruct                `json:"team,omitempty"`
	PlayerStats         PlayerStatsStruct          `json:"player_stats,omitempty"`
	Rosters             []DefaultRosterStruct      `json:"rosters,omitempty"`
	Game                GameStruct                 `json:"game,omitempty"`
	SocialMediaAccounts []SocialMediaAccountStruct `json:"social_media_accounts,omitempty"`
}
