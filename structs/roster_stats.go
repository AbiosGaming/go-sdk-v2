package structs

// RosterStatsStruct hold performance information about a particular Roster.
type RosterStatsStruct struct {
	Streak struct {
		Match struct {
			StreakScopeStruct // defined in stats.go
		} `json:"match,omitempty"`
	} `json:"streak,omitempty"`
	Winrate struct {
		Match struct {
			WinrateMatchScopeStruct // defined in stats.go
		} `json:"match,omitempty"`
	} `json:"winrate,omitempty"`
	Nemesis *struct {
		Match struct {
			Roster RosterStruct `json:"roster"`
			Losses int64        `json:"losses"`
		} `json:"match,omitempty"`
	} `json:"nemesis"`
	Dominating *struct {
		Match struct {
			Roster RosterStruct `json:"roster"`
			Wins   int64        `json:"wins"`
		} `json:"match,omitempty"`
	} `json:"dominating"`
}
