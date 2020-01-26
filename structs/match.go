package structs

import (
	"encoding/json"
)

// Match represents an actual map being played between two rosters.
type Match struct {
	Id           int64            `json:"id,omitempty"`
	SeriesId     int64            `json:"series_id,omitmepty"`
	Order        int64            `json:"order,omitempty"`
	Winner       *int64           `json:"winner"`
	Map          *Map             `json:"map,omitempty"`
	DeletedAt    *string          `json:"deleted_at"`
	Game         Game             `json:"game"`
	HasPbpStats  bool             `json:"has_pbpstats"`
	Scores       *Scores          `json:"scores"`
	Forfeit      Forfeit          `json:"forfeit,omitempty"`
	Seeding      Seeding          `json:"seeding,omitempty"`
	Rosters      []Roster         `json:"rosters"`
	Performance  MatchPerformance `json:"performance,omitempty"`
	MatchSummary MatchSummary     `json:"match_summary"` // Play by Play
}

// avoid recursion when unmarshaling
type match Match

// We need to unmarshal match_summary into the game-specific struct
func (m *Match) UnmarshalJSON(data []byte) error {
	// find the outer-most keys
	var partial map[string]json.RawMessage
	if err := json.Unmarshal(data, &partial); err != nil {
		return err
	}
	summary := partial["match_summary"]

	// This is not strictly necessary but it is faster
	delete(partial, "match_summary")
	data, _ = json.Marshal(partial)

	var mm match
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
		case 5:
			var tmp CsMatchSummary
			if err := json.Unmarshal(summary, &tmp); err != nil {
				return err
			}
			mm.MatchSummary = tmp
		default:
		}
	}

	*m = Match(mm)
	return nil
}
