package structs

// Roster represents a roster (or line-up) in a team game.
type Roster struct {
	Id          int64        `json:"id,omitempty"`
	Teams       []Team       `json:"teams,omitempty"`
	Players     []Player     `json:"players,omitempty"`
	RosterStats *RosterStats `json:"roster_stats"`
	Game        Game         `json:"game"`
}
