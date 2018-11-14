package structs

import (
	"encoding/json"
)

// MatchStruct represents an actual map being played between two rosters.
type MatchStruct struct {
	Id           int64                  `json:"id,omitempty"`
	Order        int64                  `json:"order,omitempty"`
	Winner       *int64                 `json:"winner"`
	Map          *MapStruct             `json:"map,omitempty"`
	DeletedAt    *string                `json:"deleted_at"`
	Game         GameStruct             `json:"game"`
	HasPbpStats  bool                   `json:"has_pbpstats"`
	Scores       *ScoresStruct          `json:"scores"`
	Forfeit      ForfeitStruct          `json:"forfeit,omitempty"`
	Seeding      SeedingStruct          `json:"seeding,omitempty"`
	Rosters      []RosterStruct         `json:"rosters"`
	Performance  MatchPerformanceStruct `json:"performance,omitempty"`
	MatchSummary MatchSummaryStruct     `json:"match_summary"` // Play by Play
}

// avoid recursion when unmarshaling
type matchStruct MatchStruct

// We need to unmarshal match_summary into the game-specific struct
func (m *MatchStruct) UnmarshalJSON(data []byte) error {
	// find the outer-most keys
	var partial map[string]json.RawMessage
	if err := json.Unmarshal(data, &partial); err != nil {
		return err
	}
	summary := partial["match_summary"]

	// This is not strictly necessary but it is faster
	delete(partial, "match_summary")
	data, _ = json.Marshal(partial)

	var mm matchStruct
	if err := json.Unmarshal(data, &mm); err != nil {
		return err
	}

	// "null" is 4 bytes
	if len(summary) > 4 {
		switch mm.Game.Id {
		// Dota
		case 1:
			var tmp DotaMatchSummary
			if err := json.Unmarshal(summary, &tmp); err != nil {
				return err
			}
			mm.MatchSummary = tmp
		// Lol
		case 2:
			var tmp LolMatchSummary
			if err := json.Unmarshal(summary, &tmp); err != nil {
				return err
			}
			mm.MatchSummary = tmp
		//Cs
		case 3:
			var tmp CsMatchSummary
			if err := json.Unmarshal(summary, &tmp); err != nil {
				return err
			}
			mm.MatchSummary = tmp
		default:
		}
	}

	*m = MatchStruct(mm)
	return nil
}
