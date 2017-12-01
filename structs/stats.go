package structs

// StreakScopeStruct is the second-level object holding streak statistics.
type StreakScopeStruct struct {
	Current int64 `json:"current"`
	Best    int64 `json:"best"`
	Worst   int64 `json:"worst"`
}

// WinrateScopeStruct is a second-level object holding winrate statistics for matches.
type WinrateMatchScopeStruct struct {
	Rate    float64               `json:"rate"`
	History int64                 `json:"history"`
	PerMap  []WinratePerMapStruct `json:"per_map"`
}

// WinratePerMapStruct holds the innermost JSON about winrates related to a single map.
type WinratePerMapStruct struct {
	Map     MapStruct `json:"map,omitempty"`
	Rate    float64   `json:"rate"`
	History int64     `json:"history"`
}

// WinrateScopeStruct is a second-level object holding winrate statistics for series'.
type WinrateSeriesScopeStruct struct {
	Rate      float64                  `json:"rate"`
	History   int64                    `json:"history"`
	PerFormat []WinratePerFormatStruct `json:"per_format"`
}

// WinratePerFormatStruct holds the innermost JSON about winrates related to a specific format.
type WinratePerFormatStruct struct {
	BestOf  int64   `json:"best_of"`
	Rate    float64 `json:"rate"`
	History int64   `json:"history"`
}
