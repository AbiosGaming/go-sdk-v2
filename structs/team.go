package structs

// PaginatedTeams holds a list of Team as well as information about pages.
type PaginatedTeams struct {
	LastPage    int64  `json:"last_page,omitempty"`
	CurrentPage int64  `json:"current_page,omitempty"`
	Data        []Team `json:"data,omitempty"`
}

// Team represents a team that competes in Series' and Matches.
type Team struct {
	Id                  int64                `json:"id,omitempty"`
	Name                string               `json:"name,omitempty"`
	ShortName           string               `json:"short_name,omitempty"`
	DeletedAt           *string              `json:"deleted_at"`
	Images              TeamImages           `json:"images,omitempty"`
	Country             *Country             `json:"country"`
	TeamStats           TeamStats            `json:"team_stats,omitempty"`
	Players             *[]Player            `json:"players,omitempty"`
	Rosters             []DefaultRoster      `json:"rosters,omitempty"`
	UpcomingSeries      []Series             `json:"upcoming_series"`
	RecentSeries        []Series             `json:"recent_series"`
	Game                Game                 `json:"game,omitempty"`
	SocialMediaAccounts []SocialMediaAccount `json:"social_media_accounts,omitempty"`
}
