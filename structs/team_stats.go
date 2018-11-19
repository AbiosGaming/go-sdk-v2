package structs

import (
	"encoding/json"
)

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
	PlayByPlay TeamPlayByPlayStats `json:"play_by_play"`
}

type TeamPlayByPlayStats interface{}

type teamStatsStruct TeamStatsStruct

func (t *TeamStatsStruct) UnmarshalJSON(data []byte) error {
	var partial map[string]json.RawMessage
	if err := json.Unmarshal(data, &partial); err != nil {
		return err
	}
	pbp_data := partial["play_by_play"]

	// This seems to be faster than not doing it
	delete(partial, "play_by_play")
	data, _ = json.Marshal(partial)

	// Unmarshal every but the "play_by_play" key
	var tt teamStatsStruct
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
	*t = TeamStatsStruct(tt)

	return nil
}

// CsTeamStats holds data about a team play by play stats for cs
type CsTeamStats struct {
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
	TopStats struct {
		Kills struct {
			PlayerAgainstStruct
			Kills int64 `json:"kills"`
		} `json:"kills"`
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
	} `json:"top_stats"`
	TopMatches struct {
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
	} `json:"top_matches"`
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
			MostPicked []DotaHeroWithWinsStruct `json:"most_picked"`
			MostBanned []DotaHeroWithWinsStruct `json:"most_banned"`
		} `json:"own"`
		Opponenents struct {
			MostPicked []DotaHeroWithWinsStruct `json:"most_picked"`
			MostBanned []DotaHeroWithWinsStruct `json:"most_banned"`
		} `json:"opponents"`
	} `json:"drafts"`
	TopStats struct {
		Kills struct {
			DotaPlayerAgainstStruct
			Kills int64 `json:"kills"`
		} `json:"kills"`
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
	} `json:"top_stats"`
	TopMatches struct {
		AvgLength float64 `json:"avg_length"`
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
		NrMatches int64 `json:"nr_matches"`
		NrWins    int64 `json:"nr_wins"`
		Champions struct {
			Name string `json:"name"`
		} `json:"champion"`
	} `json:"champions"`
	TopStats struct {
		Kills               LolPlayerAgainstStruct  `json:"kills"`
		Gpm                 LolPlayerAgainstStruct  `json:"gpm"`
		Xpm                 LolPlayerAgainstStruct  `json:"xpm"`
		DoubleKills         *LolPlayerAgainstStruct `json:"double_kills"`
		TripleKills         *LolPlayerAgainstStruct `json:"triple_kills"`
		QuadraKills         *LolPlayerAgainstStruct `json:"quadra_kills"`
		PentaKills          *LolPlayerAgainstStruct `json:"Penta_kills"`
		UnrealKills         *LolPlayerAgainstStruct `json:"Unreal_kills"`
		LargestKillingSpree *LolPlayerAgainstStruct `json:"largest_killing_spree"`
		LargestMultiKill    *LolPlayerAgainstStruct `json:"largest_multi_kill"`
	} `json:"top_stats"`
	TopMatches struct {
		Kpm struct {
			Avg     float64 `json:"avg"` // Only lol
			Highest struct {
				LolTeamAgainstStruct
			} `json:"highest"`
			Lowest struct {
				LolTeamAgainstStruct
			} `json:"lowest"`
		} `json:"kpm"`
		Length struct {
			Avg     float64 `json:"avg"`
			Longest struct {
				Won  LolTeamAgainstStruct `json:"won"`
				Lost LolTeamAgainstStruct `json:"lost"`
			} `json:"longest"`
			Shortest struct {
				Won  LolTeamAgainstStruct `json:"won"`
				Lost LolTeamAgainstStruct `json:"lost"`
			} `json:"shortest"`
		} `json:"length"`
	} `json:"top_matches"`
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

type LolTeamAgainstStruct struct {
	Value   float64      `json:"value"`
	MatchId int64        `json:"match_id"`
	Against RosterStruct `json:"against"`
}

type LolPlayerAgainstStruct struct {
	Value    float64 `json:"value"`
	PlayerId int64   `json:"player_id"`
	MatchId  int64   `json:"match_id"`
	Champion struct {
		Name string `json:"name"`
	} `json:"champion"`
	Against RosterStruct `json:"against"`
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
type DotaHeroWithWinsStruct struct {
	Amount int64      `json:"amount"`
	Wins   int64      `json:"wins"`
	Hero   HeroStruct `json:"hero"`
}
