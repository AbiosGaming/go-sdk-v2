package structs

// PaginatedOrganisations holds a list of Team as well as information about pages.
type PaginatedOrganisations struct {
	LastPage    int64          `json:"last_page,omitempty"`
	CurrentPage int64          `json:"current_page,omitempty"`
	Data        []Organisation `json:"data,omitempty"`
}

// Organisation represents a logical grouping of Teams across different Games
type Organisation struct {
	Id    int64  `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Teams []Team `json:"teams"`
}
