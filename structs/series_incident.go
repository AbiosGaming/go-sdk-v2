package structs

// SeriesIncidents holds information about all incidents associated with a
// Series and all it's Matches.
type SeriesIncidents struct {
	SeriesIncidents []Incident `json:"series_incidents"`
	MatchIncidents  []Incident `json:"match_incidents"`
}
