package structs

import "encoding/json"

// MatchPerformance is associated with a Match and contains performance information
// about the Rosters with respect to the specific map.
type MatchPerformance struct {
	Winrate MatchWinrate `json:"winrate,omitempty"`
}

// MatchWinrate holds the top-level keys for winrate statistics.
type MatchWinrate struct {
	Overall *MatchWinrateOverall `json:"over_all"`
	PerMap  []MatchWinratePerMap `json:"per_map"`
}

// MatchWinrateOverall holds information about the summarized performance statistics.
type MatchWinrateOverall struct {
	History int64              `json:"history,omitempty"`
	Rosters map[string]float64 `json:"-"`
}

type _MatchWinrateOverall MatchWinrateOverall

func (m *MatchWinrateOverall) UnmarshalJSON(b []byte) (err error) {
	foo := _MatchWinrateOverall{}

	if err = json.Unmarshal(b, &foo); err == nil {
		*m = MatchWinrateOverall(foo)
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

// MatchWinratePerMap breaks down the winrate statistics per Map.
type MatchWinratePerMap struct {
	Map     Map                `json:"map,omitempty"`
	History int64              `json:"history,omitempty"`
	Rosters map[string]float64 `json:"-"`
}

type _MatchWinratePerMap MatchWinratePerMap

func (m *MatchWinratePerMap) UnmarshalJSON(b []byte) (err error) {
	foo := _MatchWinratePerMap{}

	if err = json.Unmarshal(b, &foo); err == nil {
		*m = MatchWinratePerMap(foo)
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
