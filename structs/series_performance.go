package structs

// SeriesPerformanceStruct is associated with a Series and contains performance information
// about the Teams or Players participating in the Series.
type SeriesPerformanceStruct struct {
	PastEncounters    []SeriesStruct            `json:"past_encounters,omitempty"`
	RecentPerformance map[string][]SeriesStruct `json:"recent_performance,omitempty"`
}
