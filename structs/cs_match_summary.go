package structs

// CsMatchSummary is the summarization of a CS:GO match.
type CsMatchSummary struct {
	Home        int64 `json:"home"`
	Away        int64 `json:"away"`
	MatchLength int64 `json:"match_length"`
	ScoreBoard  struct {
		Home []CsScoreBoardEntry `json:"home"`
		Away []CsScoreBoardEntry `json:"away"`
	} `json:"scoreboard"`
	Rounds []Round `json:"rounds"`
}

// Round is summaryization of a CS:GO round.
type Round struct {
	RoundNr    int64  `json:"round_nr"`
	TSide      int64  `json:"t_side"`
	CtSide     int64  `json:"ct_side"`
	Winner     int64  `json:"winner"`
	WinReason  string `json:"win_reason"`
	BombEvents []BombEvent `json:"bomb_events"`
	Kills []Kill `json:"kills"`
	PlayerStats struct {
		TSide  []RoundPlayerStats `json:"t_side"`
		CtSide []RoundPlayerStats `json:"ct_side"`
	} `json:"player_stats"`
}

// RoundPlayerStats reflects how well a player performed in a round.
type RoundPlayerStats struct {
	PlayerId int64   `json:"player_id"`
	DmgGiven float64 `json:"dmg_given"`
	DmgTaken float64 `json:"dmg_taken"`
	Kills    int64   `json:"kills"`
	Assists  int64   `json:"assists"`
	Died     bool    `json:"died"`
	Accuracy struct {
		General  float64 `json:"general"`
		Headshot float64 `json:"head_shot"`
	} `json:"accuracy"`
}

// BombEvent hold data about bomb event in CS:GO round
type BombEvent struct {
	Type       string `json:"type"`
	PlayerId   int64  `json:"player_id"`
	RoundClock int64  `json:"round_clock"`
	Pos        Pos    `json:"pos"`
}

// Pos hold x, y and z coordinates.
type Pos struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

// Kill holds CS:GO kill data
type Kill struct {
	RoundClock int64 `json:"round_clock"`
	Damage     int64 `json:"damage"`
	Attacker   struct {
		PlayerId int64 `json:"player_id"`
		Pos      Pos   `json:"pos"`
	} `json:"attacker"`
	Victim struct {
		PlayerId int64 `json:"player_id"`
		Pos      Pos   `json:"pos"`
	} `json:"victim"`
	Assists  *int64 `json:"assist"`
	Weapon   Weapon `json:"weapon"`
	HitGroup string `json:"hit_group"`
}

// CsScoreBoardEntry reflects a CS:GO scoreboard entry.
type CsScoreBoardEntry struct {
	PlayerId int64   `json:"player_id"`
	Kills    int64   `json:"kills"`
	Assists  int64   `json:"assists"`
	Deaths   int64   `json:"deaths"`
	Adr      float64 `json:"adr"`
}
