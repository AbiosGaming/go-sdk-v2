package structs

import "encoding/json"

// MatchPerformanceStruct is associated with a Match and contains performance information
// about the Rosters with respect to the specific map.
type MatchPerformanceStruct struct {
	Winrate MatchWinrateStruct `json:"winrate,omitempty"`
}

// MatchWinrateStruct holds the top-level keys for winrate statistics.
type MatchWinrateStruct struct {
	Overall *MatchWinrateOverallStruct `json:"over_all"`
	PerMap  []MatchWinratePerMapStruct `json:"per_map"`
}

// MatchWinrateOverallStruct holds information about the summarized performance statistics.
type MatchWinrateOverallStruct struct {
	History int64              `json:"history,omitempty"`
	Rosters map[string]float64 `json"-"`
}

type _MatchWinrateOverallStruct MatchWinrateOverallStruct

func (m *MatchWinrateOverallStruct) UnmarshalJSON(b []byte) (err error) {
	foo := _MatchWinrateOverallStruct{}

	if err = json.Unmarshal(b, &foo); err == nil {
		*m = MatchWinrateOverallStruct(foo)
	}

	stuff := make(map[string]interface{})

	m.Rosters = make(map[string]float64)
	if err = json.Unmarshal(b, &stuff); err == nil {
		delete(stuff, "history")
		for key, value := range stuff {
			m.Rosters[key] = value.(float64)
		}
	}

	return err
}

// MatchWinratePerMapStruct breaks down the winrate statistics per Map.
type MatchWinratePerMapStruct struct {
	Map     MapStruct          `json:"map,omitempty"`
	History int64              `json:"history,omitempty"`
	Rosters map[string]float64 `json:"-"`
}

type _MatchWinratePerMapStruct MatchWinratePerMapStruct

func (m *MatchWinratePerMapStruct) UnmarshalJSON(b []byte) (err error) {
	foo := _MatchWinratePerMapStruct{}

	if err = json.Unmarshal(b, &foo); err == nil {
		*m = MatchWinratePerMapStruct(foo)
	}

	stuff := make(map[string]interface{})

	m.Rosters = make(map[string]float64)
	if err = json.Unmarshal(b, &stuff); err == nil {
		delete(stuff, "history")
		delete(stuff, "map")
		for key, value := range stuff {
			m.Rosters[key] = value.(float64)
		}
	}

	return err
}
