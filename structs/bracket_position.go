package structs

// BracketPosition represents a Series position in a Substages bracket.
type BracketPositionStruct struct {
	Part    string        `json:"part,omitempty"`
	Col     int64         `json:"col"`
	Offset  int64         `json:"offset"`
	Seeding SeedingStruct `json:"seeding,omitempty"`
}
