package structs

// SeriesPerformance is associated with a Series and contains performance information
// about the Teams or Players participating in the Series.
type SeriesPerformance struct {
	PastEncounters    []Series            `json:"past_encounters,omitempty"`
	RecentPerformance map[string][]Series `json:"recent_performance,omitempty"`
}
