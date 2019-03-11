package structs

// RosterStats hold performance information about a particular Roster.
type RosterStats struct {
	Streak struct {
		Match struct {
			StreakScope // defined in stats.go
		} `json:"match,omitempty"`
	} `json:"streak,omitempty"`
	Winrate struct {
		Match struct {
			WinrateMatchScope // defined in stats.go
		} `json:"match,omitempty"`
	} `json:"winrate,omitempty"`
	Nemesis *struct {
		Match struct {
			Roster Roster `json:"roster"`
			Losses int64  `json:"losses"`
		} `json:"match,omitempty"`
	} `json:"nemesis"`
	Dominating *struct {
		Match struct {
			Roster Roster `json:"roster"`
			Wins   int64  `json:"wins"`
		} `json:"match,omitempty"`
	} `json:"dominating"`
}
