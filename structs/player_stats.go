package structs

import (
	"encoding/json"
)

// PlayerStats hold performance statistics about a particular Player.
type PlayerStats struct {
	SinglePlayer *SinglePlayerStats    `json:"single_player,omitempty"` // Null when player does not play single player game. Otherwise format is equal to stats for a team
	PlayByPlay   PlayerPlayByPlayStats `json:"play_by_play"`
}

// SinglePlayerStats hold information for players playing single-player games.
// For team games see the players corresponding Team and TeamStats.
type SinglePlayerStats struct {
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
			Competitor Player `json:"competitor,omitempty"` // defined in stats.go
			Losses     int64  `json:"losses"`
		} `json:"series,omitempty"`
		Match struct {
			Competitor Player `json:"competitor,omitempty"` // defined in stats.go
			Losses     int64  `json:"losses"`
		} `json:"match,omitempty"`
	} `json:"nemesis,omitempty"`
	Dominating *struct {
		Series struct {
			Competitor Player `json:"competitor,omitempty"` // defined in stats.go
			Wins       int64  `json:"wins"`
		} `json:"series,omitempty"`
		Match struct {
			Competitor Player `json:"competitor,omitempty"` // defined in stats.go
			Wins       int64  `json:"wins"`
		} `json:"match,omitempty"`
	} `json:"dominating,omitempty"`
}

// PlayByPlayStats holds information about play by play statistics for a certain player.
type PlayerPlayByPlayStats interface{}

type playerPlayByPlayStats PlayerPlayByPlayStats

func (p *PlayerStats) UnmarshalJSON(data []byte) error {
	var partial map[string]json.RawMessage
	if err := json.Unmarshal(data, &partial); err != nil {
		return err
	}

	var single_player SinglePlayerStats
	if err := json.Unmarshal(partial["single_player"], &single_player); err != nil {
		return err
	}
	p.SinglePlayer = &single_player

	var pbp_map map[string]json.RawMessage
	if err := json.Unmarshal(partial["play_by_play"], &pbp_map); err != nil {
		return err
	}

	if _, ok := pbp_map["faction_stats"]; ok {
		var tmp DotaPlayerStats
		if err := json.Unmarshal(partial["play_by_play"], &tmp); err != nil {
			return err
		}
		p.PlayByPlay = tmp
	}

	if _, ok := pbp_map["side_stats"]; ok {
		var tmp LolPlayerStats
		if err := json.Unmarshal(partial["play_by_play"], &tmp); err != nil {
			return err
		}
		p.PlayByPlay = tmp
	}

	if _, ok := pbp_map["over_all"]; ok {
		var tmp CsPlayerStats
		if err := json.Unmarshal(partial["play_by_play"], &tmp); err != nil {
			return err
		}
		p.PlayByPlay = tmp
	}

	return nil
}

// CsPlayerByPlayStats holds play by play stats for cs players
type CsPlayerStats struct {
	Overall struct {
		CsPlayerPerformance
		Plants  float64 `json:"plants"`
		Defuses float64 `json:"defuses"`
	} `json:"over_all"`
	PerMap []struct {
		Map    Map `json:"map"`
		CtSide struct {
			CsPlayerPerformance
			Defuses float64 `json:"defuses"`
		} `json:"ct_side"`
		TSide struct {
			CsPlayerPerformance
			Plants float64 `json:"plants"`
		} `json:"t_side"`
		Overall struct {
			CsPlayerPerformance
		} `json:"over_all"`
	} `json:"per_map"`
	PerWeapon []struct {
		Weapon        Weapon `json:"weapon"`
		DmgGivenRound int64  `json:"dmg_given_round"`
		Accuracy      struct {
			General  float64 `json:"general"`
			Headshot float64 `json:"head_shot"`
		} `json:"accuracy"`
		History int64 `json:"history"`
	} `json:"per_weapon"`
}

// CsPlayerPerformance holds some general data about a players performance. This
// struct is re-used for different levels data (e.g per_map and over_all).
type CsPlayerPerformance struct {
	Kills       float64 `json:"kills"`
	Assists     float64 `json:"assists"`
	Deaths      float64 `json:"deaths"`
	DamageGiven float64 `json:"dmg_given"`
	DamageTaken float64 `json:"dmg_taken"`
	History     int64   `json:"history"`
	Accuracy    struct {
		General  float64 `json:"general"`
		Headshot float64 `json:"head_shot"`
	} `json:"accuracy"`
}

// DotaPlayerByPlayStats holds play by play stats for dota players
type DotaPlayerStats struct {
	Stats     DotaPlayerPerformance `json:"stats"`
	HeroStats struct {
		Attribute struct {
			Strength     DotaPlayerPerformance `json:"strength"`
			Agility      DotaPlayerPerformance `json:"agility"`
			Intelligence DotaPlayerPerformance `json:"intelligence"`
		} `json:"attribute"`
		TopHeroes []struct {
			Hero Hero `json:"hero"`
			DotaPlayerPerformance
		} `json:"top_heroes"`
	} `json:"hero_stats"`
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
}

// DotaPlayerPerformance holds some data about a a players performance. This
// struct is re-used for different levels of data (e.g hero_stats and top_hero)
type DotaPlayerPerformance struct {
	Matches        int64   `json:"matches"`
	Wins           int64   `json:"wins"`
	AvgKills       float64 `json:"avg_kills"`
	AvgDeaths      float64 `json:"avg_deaths"`
	AvgAssists     float64 `json:"avg_assists"`
	AvgCreepKills  float64 `json:"avg_creep_kills"`
	AvgCreepDenies float64 `json:"avg_creep_denies"`
	AvgGpm         float64 `json:"avg_gpm"`
	AvgXpm         float64 `json:"avg_xpm"`
}

type LolPlayerStats struct {
	NrMatches int64 `json:"nr_matches"`
	NrWins    int64 `json:"nr_wins"`
	AvgStats  struct {
		Kills       float64 `json:"kills"`
		Deaths      float64 `json:"deaths"`
		Assists     float64 `json:"assists"`
		Gpm         float64 `json:"gpm"`
		Xpm         float64 `json:"xpm"`
		MinionKills struct {
			Total              float64 `json:"total"`
			NeutralMinions     float64 `json:"neutral_minions"`
			NeutralJungle      float64 `json:"neutral_jungle"`
			NeutralEnemyJungle float64 `json:"neutral_enemy_jungle"`
		} `json:"minion_kills"`
		Wards []struct {
			Placed    float64 `json:"placed"`
			Destroyed float64 `json:"destroyed"`
			Type      string  `json:"type"`
		} `json:"wards"`
	} `json:"avg_stats"`
	LargestCombo struct {
		Double              int64 `json:"double"`
		Triple              int64 `json:"triple"`
		Quadra              int64 `json:"quadra"`
		Penta               int64 `json:"penta"`
		Unreal              int64 `json:"unreal"`
		LargestKillingSpree int64 `json:"largest_killing_spree"`
		LargestMultiKill    int64 `json:"largest_multi_kill"`
		KillingSprees       int64 `json:"killing_sprees"`
	} `json:"largest_combos"`
	MostPlayedChampion []struct {
		Champion   Champion `json:"champion"`
		NrMatches  int64    `json:"nr_matches"`
		NrWins     int64    `json:"nr_wins"`
		AvgKills   float64  `json:"avg_kills"`
		AvgDeaths  float64  `json:"avg_deaths"`
		AvgAssists float64  `json:"avg_assists"`
		AvgGpm     float64  `json:"avg_gpm"`
		AvgXpm     float64  `json:"avg_xpm"`
	} `json:"most_played_champions"`
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
}
