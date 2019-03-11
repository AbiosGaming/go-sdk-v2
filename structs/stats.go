package structs

// StreakScope is the second-level object holding streak statistics.
type StreakScope struct {
	Current int64 `json:"current"`
	Best    int64 `json:"best"`
	Worst   int64 `json:"worst"`
}

// WinrateScope is a second-level object holding winrate statistics for matches.
type WinrateMatchScope struct {
	Rate    float64         `json:"rate"`
	History int64           `json:"history"`
	PerMap  []WinratePerMap `json:"per_map"`
}

// WinratePerMap holds the innermost JSON about winrates related to a single map.
type WinratePerMap struct {
	Map     Map     `json:"map,omitempty"`
	Rate    float64 `json:"rate"`
	History int64   `json:"history"`
}

// WinrateScope is a second-level object holding winrate statistics for series'.
type WinrateSeriesScope struct {
	Rate      float64            `json:"rate"`
	History   int64              `json:"history"`
	PerFormat []WinratePerFormat `json:"per_format"`
}

// WinratePerFormat holds the innermost JSON about winrates related to a specific format.
type WinratePerFormat struct {
	BestOf  int64   `json:"best_of"`
	Rate    float64 `json:"rate"`
	History int64   `json:"history"`
}
