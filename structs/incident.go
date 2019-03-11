package structs

// PaginatedIncidents holds a list of Incident as well as information about pages
type PaginatedIncidents struct {
	LastPage    int64      `json:"last_page,omitempty"`
	CurrentPage int64      `json:"current_page,omitempty"`
	Data        []Incident `json:"data"`
}

// Incident represents an incident.
type Incident struct {
	SeriesId   int64   `json:"series_id,omitempty"`
	MatchId    *int64  `json:"match_id,omitempty"`
	Comment    string  `json:"comment,omitempty"`
	CreatedAt  string  `json:"created_at,omitempty"`
	UpdatedAt  *string `json:"updated_at"`
	IncidentId int64   `json:"incident_id,omitempty"`
}
