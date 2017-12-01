package structs

// SeriesStructPaginated holds a list of SeriesStruct as well as information about pages.
type SeriesStructPaginated struct {
	LastPage    int64          `json:"last_page,omitempty"`
	CurrentPage int64          `json:"current_page,omitempty"`
	Data        []SeriesStruct `json:"data,omitempty"`
}

// SeriesStruct represents a Series of Matches.
type SeriesStruct struct {
	Id              int64                   `json:"id,omitempty"`
	Title           string                  `json:"title,omitempty"`
	BestOf          int64                   `json:"bestOf,omitempty"`
	Tier            *int64                  `json:"tier"`
	Start           *string                 `json:"start"`
	End             *string                 `json:"end"`
	PostponedFrom   *string                 `json:"postponed_from"`
	DeletedAt       *string                 `json:"deleted_at"`
	Scores          *ScoresStruct           `json:"scores"`
	Forfeit         ForfeitStruct           `json:"forfeit,omitempty"`
	Streamed        bool                    `json:"streamed"`
	Seeding         SeedingStruct           `json:"seeding,omitempty"`
	Rosters         []RosterStruct          `json:"rosters,omitempty"`
	Game            GameStruct              `json:"game,omitempty"`
	Matches         []MatchStruct           `json:"matches,omitempty"`
	Casters         []CasterStruct          `json:"casters,omitempty"`
	SubstageId      int64                   `json:"substage_id,omitempty"`
	BracketPosition *BracketPositionStruct  `json:"bracket_pos"`
	Tournament      TournamentStruct        `json:"tournament,omitempty"`
	Performance     SeriesPerformanceStruct `json:"performance,omitempty"`
	SportsbookOdds  []SportsbookOddsStruct  `json:"sportsbook_odds"`
	Chain           *[]int64                `json:"chain"`
}
