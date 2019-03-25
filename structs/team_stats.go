package structs

import (
	"encoding/json"
)

// TeamStats holds performance statistics about a particular Team.
type TeamStats struct {
	Streak struct {
		Series struct {
			StreakScope // defined in stats.go
		} `json:"series,omitempty"`
		Match struct {
			StreakScope // defined in stats.go
		} `json:"match,omitempty"`
	} `json:"streak,omitempty"`
	Winrate struct {
		Series struct {
			WinrateSeriesScope // defined in stats.go
		} `json:"series,omitempty"`
		Match struct {
			WinrateMatchScope // defined in stats.go
		} `json:"match,omitempty"`
	} `json:"winrate,omitempty"`
	Nemesis *struct {
		Series struct {
			Competitor Team  `json:"competitor,omitempty"` // defined in stats.go
			Losses     int64 `json:"losses"`
		} `json:"series,omitempty"`
		Match struct {
			Competitor Team  `json:"competitor,omitempty"` // defined in stats.go
			Losses     int64 `json:"losses"`
		} `json:"match,omitempty"`
	} `json:"nemesis,omitempty"`
	Dominating *struct {
		Series struct {
			Competitor Team  `json:"competitor,omitempty"` // defined in stats.go
			Wins       int64 `json:"wins"`
		} `json:"series,omitempty"`
		Match struct {
			Competitor Team  `json:"competitor,omitempty"` // defined in stats.go
			Wins       int64 `json:"wins"`
		} `json:"match,omitempty"`
	} `json:"dominating,omitempty"`
	PlayByPlay TeamPlayByPlayStats `json:"play_by_play"`
}

type TeamPlayByPlayStats interface{}

type teamStats TeamStats

func (t *TeamStats) UnmarshalJSON(data []byte) error {
	var partial map[string]json.RawMessage
	if err := json.Unmarshal(data, &partial); err != nil {
		return err
	}
	pbp_data := partial["play_by_play"]

	// This seems to be faster than not doing it
	delete(partial, "play_by_play")
	data, _ = json.Marshal(partial)

	// Unmarshal every but the "play_by_play" key
	var tt teamStats
	if err := json.Unmarshal(data, &tt); err != nil {
		return err
	}

	// Find out the top keys of pbp_data
	var pbp_map map[string]json.RawMessage
	if err := json.Unmarshal(pbp_data, &pbp_map); err != nil {
		return err
	}

	// Dota
	if _, ok := pbp_map["faction_stats"]; ok {
		var tmp DotaTeamStats
		// Unmarshal the play_by_play data into tt.PlayByPlay
		if err := json.Unmarshal(pbp_data, &tmp); err != nil {
			return err
		}
		tt.PlayByPlay = tmp
	}
	// Lol
	if _, ok := pbp_map["side_stats"]; ok {
		var tmp LolTeamStats
		// Unmarshal the play_by_play data into tt.PlayByPlay
		if err := json.Unmarshal(pbp_data, &tmp); err != nil {
			return err
		}
		tt.PlayByPlay = tmp
	}
	// CS
	if _, ok := pbp_map["totals"]; ok {
		var tmp CsTeamStats
		// Unmarshal the play_by_play data into tt.PlayByPlay
		if err := json.Unmarshal(pbp_data, &tmp); err != nil {
			return err
		}
		tt.PlayByPlay = tmp
	}
	*t = TeamStats(tt)

	return nil
}

// CsTeamStats holds data about a team play by play stats for cs
type CsTeamStats struct {
	Totals struct {
		Kills  int64 `json:"kills"`
		Deaths int64 `json:"deaths"`
		CsTeamCommonStats
	} `json:"totals"`
	Maps []struct {
		Map Map `json:"map"`
		CsTeamCommonStats
	} `json:"maps"`
	Marksman []struct {
		PlayerId int64   `json:"player_id"`
		Adr      float64 `json:"adr"`
		Weapon   Weapon  `json:"weapon"`
	} `json:"marksman"`
	TopStats struct {
		Kills struct {
			PlayerAgainst
			Kills int64 `json:"kills"`
		} `json:"kills"`
		Adr struct {
			PlayerAgainst
			Adr float64 `json:"adr"`
		} `json:"adr"`
		Assists struct {
			PlayerAgainst
			Assists int64 `json:"assists"`
		} `json:"assists"`
		Plants struct {
			PlayerAgainst
			Plants int64 `json:"plants"`
		} `json:"plants"`
		Defuses struct {
			PlayerAgainst
			Defuses int64 `json:"defuses"`
		} `json:"defuses"`
	} `json:"top_stats"`
	TopMatches struct {
		BiggestLoss struct {
			TeamAgainst
			Rounds int64 `json:"rounds"`
		} `json:"biggest_loss"`
		BiggestWin struct {
			TeamAgainst
			Rounds int64 `json:"rounds"`
		} `json:"biggest_win"`
		MostRounds struct {
			TeamAgainst
			Rounds int64 `json:"rounds"`
		} `json:"most_rounds"`
	} `json:"top_matches"`
}

// CsTeamCommonStats holds information common to multiple JSON objects.
type CsTeamCommonStats struct {
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

// DotaTeamStats holds data about a teams play by play stats for dota
type DotaTeamStats struct {
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
			MostPicked []DotaHeroWithWins `json:"most_picked"`
			MostBanned []DotaHeroWithWins `json:"most_banned"`
		} `json:"own"`
		Opponents struct {
			MostPicked []DotaHeroWithWins `json:"most_picked"`
			MostBanned []DotaHeroWithWins `json:"most_banned"`
		} `json:"opponents"`
	} `json:"drafts"`
	TopStats struct {
		Kills struct {
			DotaPlayerAgainst
			Kills int64 `json:"kills"`
		} `json:"kills"`
		Gpm struct {
			DotaPlayerAgainst
			Gpm float64 `json:"gpm"`
		} `json:"gpm"`
		Xpm struct {
			DotaPlayerAgainst
			Xpm float64 `json:"xpm"`
		} `json:"xpm"`
		DmgGiven struct {
			DotaPlayerAgainst
			DmgGiven float64 `json:"dmg_given"`
		} `json:"dmg_given"`
		CreepKills struct {
			DotaPlayerAgainst
			LastHits int64 `json:"last_hits"`
		} `json:"creep_kills"`
		CreepDenies struct {
			DotaPlayerAgainst
			Denies int64 `json:"denies"`
		} `json:"creep_denies"`
	} `json:"top_stats"`
	TopMatches struct {
		AvgLength float64 `json:"avg_length"`
		Longest   struct {
			Won struct {
				TeamAgainst
				Length int64 `json:"length"`
			} `json:"won"`
			Lost struct {
				TeamAgainst
				Length int64 `json:"length"`
			} `json:"lost"`
		} `json:"longest"`
		Shortest struct {
			Won struct {
				TeamAgainst
				Length int64 `json:"length"`
			} `json:"won"`
			Lost struct {
				TeamAgainst
				Length int64 `json:"length"`
			} `json:"lost"`
		} `json:"shortest"`
		AvgKpm float64 `json:"avg_kpm"`
		Kpm    struct {
			Highest struct {
				Kpm float64 `json:"kpm"`
				TeamAgainst
			} `json:"highest"`
			Lowest struct {
				Kpm float64 `json:"kpm"`
				TeamAgainst
			} `json:"lowest"`
		} `json:"kpm"`
	} `json:"top_matches"`
}

type LolTeamStats struct {
	NrMatches int64 `json:"nr_matches"`
	NrWins    int64 `json:"nr_wins"`
	SideStats struct {
		Purple struct {
			NrMatches int64 `json:"nr_matches"`
			NrWins    int64 `json:"nr_wins"`
		} `json:"purple"`
		Blue struct {
			NrMatches int64 `json:"nr_matches"`
			NrWins    int64 `json:"nr_wins"`
		} `json:"blue"`
	} `json:"side_stats"`
	Champions []struct {
		NrMatches int64    `json:"nr_matches"`
		NrWins    int64    `json:"nr_wins"`
		Champion  Champion `json:"champion"`
	} `json:"champions"`
	TopStats struct {
		Kills               LolPlayerAgainst  `json:"kills"`
		Gpm                 LolPlayerAgainst  `json:"gpm"`
		Xpm                 LolPlayerAgainst  `json:"xpm"`
		DoubleKills         *LolPlayerAgainst `json:"double_kills"`
		TripleKills         *LolPlayerAgainst `json:"triple_kills"`
		QuadraKills         *LolPlayerAgainst `json:"quadra_kills"`
		PentaKills          *LolPlayerAgainst `json:"Penta_kills"`
		UnrealKills         *LolPlayerAgainst `json:"Unreal_kills"`
		LargestKillingSpree *LolPlayerAgainst `json:"largest_killing_spree"`
		LargestMultiKill    *LolPlayerAgainst `json:"largest_multi_kill"`
	} `json:"top_stats"`
	TopMatches struct {
		Kpm struct {
			Avg     float64 `json:"avg"` // Only lol
			Highest struct {
				LolTeamAgainst
			} `json:"highest"`
			Lowest struct {
				LolTeamAgainst
			} `json:"lowest"`
		} `json:"kpm"`
		Length struct {
			Avg     float64 `json:"avg"`
			Longest struct {
				Won  LolTeamAgainst `json:"won"`
				Lost LolTeamAgainst `json:"lost"`
			} `json:"longest"`
			Shortest struct {
				Won  LolTeamAgainst `json:"won"`
				Lost LolTeamAgainst `json:"lost"`
			} `json:"shortest"`
		} `json:"length"`
	} `json:"top_matches"`
}

// TeamAgainst is a collection of common data when examining specific stats.
// It is grouped with the specific stat in another struct.
type TeamAgainst struct {
	MatchId int64 `json:"match_id"`
	Against *Team `json:"against"` // Declared as pointer to avoid invalid recursive type
}

// DotaPlayerAgainst is a grouping of PlayerAgainst and a Hero
type DotaPlayerAgainst struct {
	PlayerAgainst
	Hero Hero `json:"hero"`
}

type LolTeamAgainst struct {
	Value   float64 `json:"value"`
	MatchId int64   `json:"match_id"`
	Against Roster  `json:"against"`
}

type LolPlayerAgainst struct {
	Value    float64  `json:"value"`
	PlayerId int64    `json:"player_id"`
	MatchId  int64    `json:"match_id"`
	Champion Champion `json:"champion"`
	Against  Roster   `json:"against"`
}

// PlayerAgainst is a collection of common data when examining specific stats.
// It is grouped with the specific stat in another struct.
type PlayerAgainst struct {
	PlayerId int64 `json:"player_id"`
	Against  *Team `json:"against"` // Declared as pointer to avoid invalid recursive type
	MatchId  int64 `json:"match_id"`
}

// DotaHeroWithAmount holds information about a Dota Hero and an integer representing
// and amount (e.g amount of times picked).
type DotaHeroWithWins struct {
	Amount int64 `json:"amount"`
	Wins   int64 `json:"wins"`
	Hero   Hero  `json:"hero"`
}
