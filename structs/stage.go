package structs

// Stage represents a phase in a Tournament and a higher level grouping of Substages.
type Stage struct {
	Id        int64      `json:"id,omitempty"`
	Title     string     `json:"title,omitempty"`
	DeletedAt *string    `json:"deleted_at"`
	Substages []Substage `json:"substages"`
}
