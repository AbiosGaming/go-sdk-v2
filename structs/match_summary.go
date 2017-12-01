package structs

// MatchSummaryStruct holds information about play by play statistics for a certain match.
type MatchSummaryStruct struct {
	CsMatchSummaryStruct
	DotaMatchSummaryStruct
}

// CsMatchSummaryStruct is the summarization of a CS:GO match.
type CsMatchSummaryStruct struct {
	Home       int64 `json:"home"`
	Away       int64 `json:"away"`
	ScoreBoard struct {
		Home []CsScoreBoardEntry `json:"home"`
		Away []CsScoreBoardEntry `json:"away"`
	} `json:"scoreboard"`
	Rounds []struct {
		RoundNr    int64  `json:"round_nr"`
		TSide      int64  `json:"t_side"`
		CtSide     int64  `json:"ct_side"`
		Winner     int64  `json:"winner"`
		WinReason  string `json:"win_reason"`
		BombEvents []struct {
			Type       string    `json:"type"`
			PlayerId   int64     `json:"player_id"`
			RoundClock int64     `json:"round_clock"`
			Pos        PosStruct `json:"pos"`
		} `json:"bomb_events"`
		Kills []struct {
			RoundClock int64 `json:"round_clock"`
			Damage     int64 `json:"damage"`
			Attacker   struct {
				PlayerId int64     `json:"player_id"`
				Pos      PosStruct `json:"pos"`
			} `json:"attacker"`
			Victim struct {
				PlayerId int64     `json:"player_id"`
				Pos      PosStruct `json:"pos"`
			} `json:"victim"`
			Assists  *int64       `json:"assist"`
			Weapon   WeaponStruct `json:"weapon"`
			HitGroup string       `json:"hit_group"`
		} `json:"kills"`
		PlayerStats struct {
			TSide  []RoundPlayerStatsStruct `json:"t_side"`
			CtSide []RoundPlayerStatsStruct `json:"ct_side"`
		} `json:"player_stats"`
	} `json:"rounds"`
}

// RoundPlayerStatsStruct reflects how well a player performed in a round.
type RoundPlayerStatsStruct struct {
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

// PosStruct hold x, y and z coordinates.
type PosStruct struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

// CsScoreBoardEntry reflects a CS:GO scoreboard entry.
type CsScoreBoardEntry struct {
	PlayerId int64   `json:"player_id"`
	Kills    int64   `json:"kills"`
	Assists  int64   `json:"assists"`
	Deaths   int64   `json:"deaths"`
	Adr      float64 `json:"adr"`
}

// DotaMatchSummaryStruct is the summarization of a Dota match.
type DotaMatchSummaryStruct struct {
	RadiantRoster int64 `json:"radiant_roster"`
	DireRoster    int64 `json:"dire_roster"`
	MatchLength   int64 `json:"match_length"`
	DraftSeq      []struct {
		Order    int64      `json:"order"`
		Type     string     `json:"type"`
		RosterId int64      `json:"roster_id"`
		Hero     HeroStruct `json:"hero"`
	} `json:"draft_seq"`
	FirstBlood struct {
		Killer  int64   `json:"killer"`
		Victim  int64   `json:"victim"`
		AtTime  int64   `json:"at_time"`
		Assists []int64 `json:"assists"`
	} `json:"first_blood"`
	Kills []struct {
		Killer  int64   `json:"killer"`
		Victim  int64   `json:"victim"`
		AtTime  int64   `json:"at_time"`
		Assists []int64 `json:"assists"`
	} `json:"kills"`
	StructureDest []struct {
		Killer        int64  `json:"killer"`
		StructureType string `json:"structure_type"`
		StructurePos  string `json:"structure_pos"`
		AtTime        int64  `json:"at_time"`
	} `json:"structure_dest"`
	PlayerStats []struct {
		PlayerId    int64            `json:"player_id"`
		Hero        HeroStruct       `json:"hero"`
		Kills       int64            `json:"kills"`
		Deaths      int64            `json:"deaths"`
		Assists     int64            `json:"assists"`
		Gpm         float64          `json:"gpm"`
		Xpm         float64          `json:"xpm"`
		Levels      map[string]int64 `json:"levels"`
		CreepKills  int64            `json:"creep_kills"`
		CreepDenies int64            `json:"creep_denies"`
		HeroDmg     struct {
			Given struct {
				ByHero DotaDmg `json:"by_hero"`
				ByMob  DotaDmg `json:"by_mob"`
			} `json:"given"`
			Taken struct {
				FromHeroes DotaDmg `json:"from_heroes"`
				FromMobs   DotaDmg `json:"from_mobs"`
			} `json:"taken"`
		} `json:"hero_dmg"`
	} `json:"player_stats"`
	RoshanEvents []struct {
		Type   string `json:"type"`
		Killer int64  `json:"killer"`
		AtTime int64  `json:"at_time"`
	} `json:"roshan_events"`
}

// DotaDmg is a collection of different types of damages given/taken in dota.
type DotaDmg struct {
	HpRemoval   int64 `json:"hp_removal"`
	MagicalDmg  int64 `json:"magical_dmg"`
	PhysicalDmg int64 `json:"physical_dmg"`
	PureDmg     int64 `json:"pure_dmg"`
}
