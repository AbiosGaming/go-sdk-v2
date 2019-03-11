package structs

// BracketPosition represents a Series position in a Substages bracket.
type BracketPosition struct {
	Part    string  `json:"part,omitempty"`
	Col     int64   `json:"col"`
	Offset  int64   `json:"offset"`
	Seeding Seeding `json:"seeding,omitempty"`
}
