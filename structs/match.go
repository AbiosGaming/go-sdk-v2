package structs

// MatchStruct represents an actual map being played between two rosters.
type MatchStruct struct {
	Id           int64                  `json:"id,omitempty"`
	Order        int64                  `json:"order,omitempty"`
	Winner       *int64                 `json:"winner"`
	Map          *MapStruct             `json:"map,omitempty"`
	DeletedAt    *string                `json:"deleted_at"`
	Game         GameStruct             `json:"game"`
	HasPbpStats  bool                   `json:"has_pbpstats"`
	Scores       *ScoresStruct          `json:"scores"`
	Forfeit      ForfeitStruct          `json:"forfeit,omitempty"`
	Seeding      SeedingStruct          `json:"seeding,omitempty"`
	Rosters      []RosterStruct         `json:"rosters"`
	Performance  MatchPerformanceStruct `json:"performance,omitempty"`
	MatchSummary *MatchSummaryStruct    `json:"match_summary"` // Play by Play
}
