package structs

// TeamStructPaginated holds a list of TeamStruct as well as information about pages.
type TeamStructPaginated struct {
	LastPage    int64        `json:"last_page,omitempty"`
	CurrentPage int64        `json:"current_page,omitempty"`
	Data        []TeamStruct `json:"data,omitempty"`
}

// TeamStruct represents a team that competes in Series' and Matches.
type TeamStruct struct {
	Id                  int64                      `json:"id,omitempty"`
	Name                string                     `json:"name,omitempty"`
	ShortName           string                     `json:"short_name,omitempty"`
	DeletedAt           *string                    `json:"deleted_at"`
	Images              TeamImagesStruct           `json:"images,omitempty"`
	Country             *CountryStruct             `json:"country"`
	TeamStats           TeamStatsStruct            `json:"team_stats,omitempty"`
	Players             *[]PlayerStruct            `json:"players,omitempty"`
	Rosters             []DefaultRosterStruct      `json:"rosters,omitempty"`
	UpcomingSeries      []SeriesStruct             `json:"upcoming_series"`
	RecentSeries        []SeriesStruct             `json:"recent_series"`
	Game                GameStruct                 `json:"game,omitempty"`
	SocialMediaAccounts []SocialMediaAccountStruct `json:"social_media_accounts,omitempty"`
}
