package structs

// RosterStruct represents a roster (or line-up) in a team game.
type RosterStruct struct {
	Id          int64              `json:"id,omitempty"`
	Teams       []TeamStruct       `json:"teams,omitempty"`
	Players     []PlayerStruct     `json:"players,omitempty"`
	RosterStats *RosterStatsStruct `json:"roster_stats"`
	Game        GameStruct         `json:"game"`
}
