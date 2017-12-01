package structs

// TeamStatsStruct holds performance statistics about a particular Team.
type TeamStatsStruct struct {
	Streak struct {
		Series struct {
			StreakScopeStruct // defined in stats.go
		} `json:"series,omitempty"`
		Match struct {
			StreakScopeStruct // defined in stats.go
		} `json:"match,omitempty"`
	} `json:"streak,omitempty"`
	Winrate struct {
		Series struct {
			WinrateSeriesScopeStruct // defined in stats.go
		} `json:"series,omitempty"`
		Match struct {
			WinrateMatchScopeStruct // defined in stats.go
		} `json:"match,omitempty"`
	} `json:"winrate,omitempty"`
	Nemesis *struct {
		Series struct {
			Competitor TeamStruct `json:"competitor,omitempty"` // defined in stats.go
			Losses     int64      `json:"losses"`
		} `json:"series,omitempty"`
		Match struct {
			Competitor TeamStruct `json:"competitor,omitempty"` // defined in stats.go
			Losses     int64      `json:"losses"`
		} `json:"match,omitempty"`
	} `json:"nemesis,omitempty"`
	Dominating *struct {
		Series struct {
			Competitor TeamStruct `json:"competitor,omitempty"` // defined in stats.go
			Wins       int64      `json:"wins"`
		} `json:"series,omitempty"`
		Match struct {
			Competitor TeamStruct `json:"competitor,omitempty"` // defined in stats.go
			Wins       int64      `json:"wins"`
		} `json:"match,omitempty"`
	} `json:"dominating,omitempty"`
	PlayByPlay TeamPlayByPlayStatsStruct `json:"play_by_play"`
}

// TeamPlayByPlayStatsStruct is a grouping of structure for TeamPlayByPlay stats for different
// games.
type TeamPlayByPlayStatsStruct struct {
	CsTeamPlayByPlayStatsStruct
	DotaTeamPlayByPlayStatsStruct
	SharedTeamPlayByPlayStatsStruct
}

// SharedTeamPlayByPlayStatsStruct holds a merge of the shared top-level json objects
type SharedTeamPlayByPlayStatsStruct struct {
	TopStats   SharedTopStatsStruct   `json:"top_stats"`
	TopMatches SharedTopMatchesStruct `json:"top_matches"`
}

// SharedTopStatsStruct holds all information that can be present in the top_stats key.
type SharedTopStatsStruct struct {
	// CS & Dota
	Kills struct {
		DotaPlayerAgainstStruct       // DotaPlayerAgainst extends PlayerAgainst (which is all that is needed for CS)
		Kills                   int64 `json:"kills"`
	} `json:"kills"`

	// CS
	Adr struct {
		PlayerAgainstStruct
		Adr float64 `json:"adr"`
	} `json:"adr"`
	Assists struct {
		PlayerAgainstStruct
		Assists int64 `json:"assists"`
	} `json:"assists"`
	Plants struct {
		PlayerAgainstStruct
		Plants int64 `json:"plants"`
	} `json:"plants"`
	Defuses struct {
		PlayerAgainstStruct
		Defuses int64 `json:"defuses"`
	} `json:"defuses"`

	// Dota
	Gpm struct {
		DotaPlayerAgainstStruct
		Gpm float64 `json:"gpm"`
	} `json:"gpm"`
	Xpm struct {
		DotaPlayerAgainstStruct
		Xpm float64 `json:"xpm"`
	} `json:"xpm"`
	DmgGiven struct {
		DotaPlayerAgainstStruct
		DmgGiven float64 `json:"dmg_given"`
	} `json:"dmg_given"`
	CreepKills struct {
		DotaPlayerAgainstStruct
		LastHits int64 `json:"last_hits"`
	} `json:"creep_kills"`
	CreepDenies struct {
		DotaPlayerAgainstStruct
		Denies int64 `json:"denies"`
	} `json:"creep_denies"`
}

// SharedTopMatchesStruct holds all information that can be present in the top_matches key.
type SharedTopMatchesStruct struct {
	// Cs & Dota

	// Cs
	BiggestLoss struct {
		TeamAgainstStruct
		Rounds int64 `json:"rounds"`
	} `json:"biggest_loss"`
	BiggestWin struct {
		TeamAgainstStruct
		Rounds int64 `json:"rounds"`
	} `json:"biggest_win"`
	MostRounds struct {
		TeamAgainstStruct
		Rounds int64 `json:"rounds"`
	} `json:"most_rounds"`

	// Dota
	AvgLength int64 `json:"avg_length"`
	Longest   struct {
		Won struct {
			TeamAgainstStruct
			Length int64 `json:"length"`
		} `json:"won"`
		Lost struct {
			TeamAgainstStruct
			Length int64 `json:"length"`
		} `json:"lost"`
	} `json:"longest"`
	Shortest struct {
		Won struct {
			TeamAgainstStruct
			Length int64 `json:"length"`
		} `json:"won"`
		Lost struct {
			TeamAgainstStruct
			Length int64 `json:"length"`
		} `json:"lost"`
	} `json:"shortest"`
	AvgKpm float64 `json:"avg_kpm"`
	Kpm    struct {
		Highest struct {
			Kpm float64 `json:"kpm"`
			TeamAgainstStruct
		} `json:"highest"`
		Lowest struct {
			Kpm float64 `json:"kpm"`
			TeamAgainstStruct
		} `json:"lowest"`
	} `json:"kpm"`
}

// CsTeamPlayByPlayStruct holds data about a team play by play stats for cs
type CsTeamPlayByPlayStatsStruct struct {
	Totals struct {
		Kills  int64 `json:"kills"`
		Deaths int64 `json:"deaths"`
		CsTeamCommonStatsStruct
	} `json:"totals"`
	Maps []struct {
		Map MapStruct `json:"map"`
		CsTeamCommonStatsStruct
	} `json:"maps"`
	Marksman []struct {
		PlayerId int64        `json:"player_id"`
		Adr      float64      `json:"adr"`
		Weapon   WeaponStruct `json:"weapon"`
	} `json:"marksman"`
}

// CsTeamCommonStatsStruct holds information common to multiple JSON objects.
type CsTeamCommonStatsStruct struct {
	NrMatches      int64   `json:"nr_matches"`
	CtRounds       int64   `json:"ct_rounds"`
	CtWins         int64   `json:"ct_wins"`
	TRounds        int64   `json:"t_rounds"`
	TWins          int64   `json:"t_wins"`
	PistolRounds   int64   `json:"pistol_rounds"`
	PistolWins     int64   `json:"pistol_wins"`
	FirstKillRate  float64 `json:"first_kill_rate"`
	FirstDeathRate float64 `json:"first_death_rate"`
}

// DotaTeamPlayByPlayStatsStruct holds data about a teams play by play stats for dota
type DotaTeamPlayByPlayStatsStruct struct {
	FactionStats struct {
		Radiant struct {
			Matches int64 `json:"matches"`
			Wins    int64 `json:"wins"`
		} `json:"radiant"`
		Dire struct {
			Matches int64 `json:"matches"`
			Wins    int64 `json:"wins"`
		} `json:"dire"`
	} `json:"faction_stats"`
	Drafts struct {
		Own struct {
			MostPicked []DotaHeroWithAmountStruct `json:"most_picked"`
			MostBanned []DotaHeroWithAmountStruct `json:"most_banned"`
		} `json:"own"`
		Opponenents struct {
			MostPicked []DotaHeroWithAmountStruct `json:"most_picked"`
			MostBanned []DotaHeroWithAmountStruct `json:"most_banned"`
		} `json:"opponents"`
	} `json:"drafts"`
}

// TeamAgainstStruct is a collection of common data when examining specific stats.
// It is grouped with the specific stat in another struct.
type TeamAgainstStruct struct {
	MatchId int64       `json:"match_id"`
	Against *TeamStruct `json:"against"` // Declared as pointer to avoid invalid recursive type
}

// DotaPlayerAgainstStruct is a grouping of PlayerAgainstStruct and a HeroStruct
type DotaPlayerAgainstStruct struct {
	PlayerAgainstStruct
	Hero HeroStruct `json:"hero"`
}

// PlayerAgainstStruct is a collection of common data when examining specific stats.
// It is grouped with the specific stat in another struct.
type PlayerAgainstStruct struct {
	PlayerId int64       `json:"player_id"`
	Against  *TeamStruct `json:"against"` // Declared as pointer to avoid invalid recursive type
	MatchId  int64       `json:"match_id"`
}

// DotaHeroWithAmountStruct holds information about a Dota Hero and an integer representing
// and amount (e.g amount of times picked).
type DotaHeroWithAmountStruct struct {
	Amount int64      `json:"amount"`
	Hero   HeroStruct `json:"hero"`
}
