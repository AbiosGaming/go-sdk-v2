package structs

// Substage is the lowest structure of a Tournament and is a grouping of Series.
type Substage struct {
	TournamentId int64         `json:"tournament_id,omitempty"`
	StageId      int64         `json:"stage_id,omitempty"`
	Id           int64         `json:"id,omitempty"`
	Title        string        `json:"title,omitempty"`
	Tier         int64         `json:"tier"`
	Type         int64         `json:"type"`
	Order        int64         `json:"order"`
	Rules        SubstageRules `json:"rules"`
	Standing     []Standings   `json:"standings"`
	Series       []Series      `json:"series,omitempty"`
	Rosters      []Roster      `json:"rosters,omitempty"`
	DeletedAt    *string       `json:"deleted_at"`
}

// SubstageRules hold information about the rules for a particular substage.
type SubstageRules struct {
	Advance struct {
		Number     *int64 `json:"number"`
		SubstageId *int64 `json:"substage_id"`
	} `json:"advance"`
	Descend struct {
		Number     *int64 `json:"number"`
		SubstageId *int64 `json:"substage_id"`
	} `json:"descend"`
	Points struct {
		Win   *int64 `json:"win"`
		Draw  *int64 `json:"draw"`
		Loss  *int64 `json:"loss"`
		Scope string `json:"scope"`
	}
}

// Stands represents the current standings in a substage.
type Standings struct {
	RosterId int64  `json:"roster_id"`
	Points   *int64 `json:"points"`
	Wins     int64  `json:"wins"`
	Draws    int64  `json:"draws"`
	Losses   int64  `json:"losses"`
}
