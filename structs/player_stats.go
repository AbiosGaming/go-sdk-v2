package structs

// PlayerStatsStruct hold performance statistics about a particular Player.
type PlayerStatsStruct struct {
	SinglePlayer *SinglePlayerStatsStruct     `json:"single_player,omitempty"` // Null when player does not play single player game. Otherwise format is equal to stats for a team
	PlayByPlay   *PlayerPlayByPlayStatsStruct `json:"play_by_play"`
}

// SinglePlayerStatsStruct hold information for players playing single-player games.
// For team games see the players corresponding Team and TeamStats.
type SinglePlayerStatsStruct struct {
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
			Competitor PlayerStruct `json:"competitor,omitempty"` // defined in stats.go
			Losses     int64        `json:"losses"`
		} `json:"series,omitempty"`
		Match struct {
			Competitor PlayerStruct `json:"competitor,omitempty"` // defined in stats.go
			Losses     int64        `json:"losses"`
		} `json:"match,omitempty"`
	} `json:"nemesis,omitempty"`
	Dominating *struct {
		Series struct {
			Competitor PlayerStruct `json:"competitor,omitempty"` // defined in stats.go
			Wins       int64        `json:"wins"`
		} `json:"series,omitempty"`
		Match struct {
			Competitor PlayerStruct `json:"competitor,omitempty"` // defined in stats.go
			Wins       int64        `json:"wins"`
		} `json:"match,omitempty"`
	} `json:"dominating,omitempty"`
}

// PlayByPlayStatsStruct holds information about play by play statistics for a certain player.
type PlayerPlayByPlayStatsStruct struct {
	CsPlayerPlayByPlayStatsStruct
	DotaPlayerPlayByPlayStatsStruct
}

// CsPlayerByPlayStatsStruct holds play by play stats for cs players
type CsPlayerPlayByPlayStatsStruct struct {
	Overall struct {
		CsPlayerPerformanceStruct
		Plants  float64 `json:"plants"`
		Defuses float64 `json:"defuses"`
	} `json:"over_all"`
	PerMap []struct {
		Map    MapStruct `json:"map"`
		CtSide struct {
			CsPlayerPerformanceStruct
			Defuses float64 `json:"defuses"`
		} `json:"ct_side"`
		TSide struct {
			CsPlayerPerformanceStruct
			Plants float64 `json:"plants"`
		} `json:"t_side"`
		Overall struct {
			CsPlayerPerformanceStruct
		} `json:"over_all"`
	} `json:"per_map"`
	PerWeapon []struct {
		Weapon        WeaponStruct `json:"weapon"`
		DmgGivenRound int64        `json:"dmg_given_round"`
		Accuracy      struct {
			General  float64 `json:"general"`
			Headshot float64 `json:"head_shot"`
		} `json:"accuracy"`
		History int64 `json:"history"`
	} `json:"per_weapon"`
}

// CsPlayerPerformanceStruct holds some general data about a players performance. This
// struct is re-used for different levels data (e.g per_map and over_all).
type CsPlayerPerformanceStruct struct {
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

// DotaPlayerByPlayStatsStruct holds play by play stats for dota players
type DotaPlayerPlayByPlayStatsStruct struct {
	Stats     DotaPlayerPerformanceStruct `json:"stats"`
	HeroStats struct {
		Attribute struct {
			Strength     DotaPlayerPerformanceStruct `json:"strength"`
			Agility      DotaPlayerPerformanceStruct `json:"agility"`
			Intelligence DotaPlayerPerformanceStruct `json:"intelligence"`
		} `json:"attribute"`
		TopHeroes []struct {
			Hero HeroStruct `json:"hero"`
			DotaPlayerPerformanceStruct
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

// DotaPlayerPerformanceStruct holds some data about a a players performance. This
// struct is re-used for different levels of data (e.g hero_stats and top_hero)
type DotaPlayerPerformanceStruct struct {
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
