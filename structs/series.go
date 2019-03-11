package structs

import "encoding/json"

// PaginatedSeries holds a list of Series as well as information about pages.
type PaginatedSeries struct {
	LastPage    int64    `json:"last_page,omitempty"`
	CurrentPage int64    `json:"current_page,omitempty"`
	Data        []Series `json:"data,omitempty"`
}

// Series represents a Series of Matches.
type Series struct {
	Id              int64             `json:"id,omitempty"`
	Title           string            `json:"title,omitempty"`
	BestOf          int64             `json:"bestOf,omitempty"`
	Tier            *int64            `json:"tier"`
	Start           *string           `json:"start"`
	End             *string           `json:"end"`
	PostponedFrom   *string           `json:"postponed_from"`
	DeletedAt       *string           `json:"deleted_at"`
	Scores          *Scores           `json:"scores"`
	Forfeit         Forfeit           `json:"forfeit,omitempty"`
	Streamed        bool              `json:"streamed"`
	Seeding         Seeding           `json:"seeding,omitempty"`
	Rosters         []Roster          `json:"rosters,omitempty"`
	Game            Game              `json:"game,omitempty"`
	Matches         []Match           `json:"matches,omitempty"`
	Casters         []Caster          `json:"casters,omitempty"`
	TournamentId    int64             `json:"tournament_id,omitempty"`
	SubstageId      int64             `json:"substage_id,omitempty"`
	BracketPosition *BracketPosition  `json:"bracket_pos"`
	Tournament      Tournament        `json:"tournament,omitempty"`
	Performance     SeriesPerformance `json:"performance,omitempty"`
	SportsbookOdds  []SportsbookOdds  `json:"sportsbook_odds"`
	Chain           *[]struct {
		RootId   int64 `json:"root_id"`
		SeriesId int64 `json:"series_id"`
		Order    int64 `json:"order"`
	} `json:"chain"`
	Summary SeriesSummary `json:"summary"`
}

// avoid recursion when unmarshaling
type series Series

// We need to unmarshal summary into the game-specific struct
func (s *Series) UnmarshalJSON(data []byte) error {
	// find the outer-most keys
	var partial map[string]json.RawMessage
	if err := json.Unmarshal(data, &partial); err != nil {
		return err
	}
	summary := partial["summary"]

	// This is not strictly necessary but it is faster
	delete(partial, "summary")
	data, _ = json.Marshal(partial)

	var ss series
	if err := json.Unmarshal(data, &ss); err != nil {
		return err
	}

	// "null" is 4 bytes
	if len(summary) > 4 {
		switch ss.Game.Id {
		// Dota
		case 1:
			var tmp DotaSeriesSummary
			if err := json.Unmarshal(summary, &tmp); err != nil {
				return err
			}
			ss.Summary = tmp
		// Lol
		case 2:
			var tmp LolSeriesSummary
			if err := json.Unmarshal(summary, &tmp); err != nil {
				return err
			}
			ss.Summary = tmp
		//Cs
		case 5:
			var tmp CsSeriesSummary
			if err := json.Unmarshal(summary, &tmp); err != nil {
				return err
			}
			ss.Summary = tmp
		default:
		}
	}

	*s = Series(ss)
	return nil

}
