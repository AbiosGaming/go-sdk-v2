package structs

// SeriesIncidentStruct holds information about all incidents associated with a
// Series and all it's Matches.
type SeriesIncidentsStruct struct {
	SeriesIncidents []IncidentStruct `json:"series_incidents"`
	MatchIncidents  []IncidentStruct `json:"match_incidents"`
}
