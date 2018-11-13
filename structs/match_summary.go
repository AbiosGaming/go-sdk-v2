package structs

// MatchSummaryStruct holds information about play by play statistics for a certain match.
type MatchSummaryStruct struct {
	CsMatchSummaryStruct
	DotaMatchSummaryStruct
	LolMatchSummaryStruct
}

// CsMatchSummaryStruct is the summarization of a CS:GO match.
type CsMatchSummaryStruct struct {
	Home        int64 `json:"home"`
	Away        int64 `json:"away"`
	MatchLength int64 `json:"match_length"`
	ScoreBoard  struct {
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
		Killer int64 `json:"killer"`
		Victim int64 `json:"victim"`
		AtTime int64 `json:"at_time"`
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
				ByMobs DotaDmg `json:"by_mobs"`
			} `json:"given"`
			Taken struct {
				FromHeroes DotaDmg `json:"from_heroes"`
				FromMobs   DotaDmg `json:"from_mobs"`
			} `json:"taken"`
		} `json:"hero_dmg"`
		HeroHealing struct {
			Given struct {
				ByHero int64 `json:"by_hero"`
				ByMobs int64 `json:"by_mobs"`
			} `json:"given"`
			Taken struct {
				FromHeroes int64 `json:"from_heroes"`
				FromMobs   int64 `json:"from_mobs"`
			} `json:"taken"`
		} `json:"hero_healing"`
		Items struct {
			Inventory struct {
				Slot1 DotaItemStruct `json:"slot_1"`
				Slot2 DotaItemStruct `json:"slot_2"`
				Slot3 DotaItemStruct `json:"slot_3"`
				Slot4 DotaItemStruct `json:"slot_4"`
				Slot5 DotaItemStruct `json:"slot_5"`
				Slot6 DotaItemStruct `json:"slot_6"`
			} `json:"inventory"`
			Backpack struct {
				Slot1 DotaItemStruct `json:"slot_1"`
				Slot2 DotaItemStruct `json:"slot_2"`
				Slot3 DotaItemStruct `json:"slot_3"`
			} `json:"backback"`
			Stash struct {
				Slot1 DotaItemStruct `json:"slot_1"`
				Slot2 DotaItemStruct `json:"slot_2"`
				Slot3 DotaItemStruct `json:"slot_3"`
				Slot4 DotaItemStruct `json:"slot_4"`
				Slot5 DotaItemStruct `json:"slot_5"`
				Slot6 DotaItemStruct `json:"slot_6"`
			} `json:"stash"`
		} `json:"items"`
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

type LolMatchSummaryStruct struct {
	MatchLength  int64             `json:"match_length"`
	BlueRoster   []LolPlayerStruct `json:"blue_roster"`
	PurpleRoster []LolPlayerStruct `json:"purple_roster"`
	Firsts       struct {
		FirstBlood struct {
			PlayerId  int64  `json:"player_id"`
			Timestamp int64  `json:"timestamp"`
			Team      string `json:"team"`
		} `json:"first_blood"`
		FirstTower struct {
			PlayerId  int64  `json:"player_id"`
			Timestamp int64  `json:"timestamp"`
			Team      string `json:"team"`
		} `json:"first_tower"`
		FirstInhibitor struct {
			PlayerId  int64  `json:"player_id"`
			Timestamp int64  `json:"timestamp"`
			Team      string `json:"team"`
		} `json:"first_inhibitor"`
		FirstBaron struct {
			PlayerId  int64  `json:"player_id"`
			Timestamp int64  `json:"timestamp"`
			Team      string `json:"team"`
		} `json:"first_baron"`
		FirstDragon struct {
			PlayerId  int64  `json:"player_id"`
			Timestamp int64  `json:"timestamp"`
			Team      string `json:"team"`
		} `json:"first_dragon"`
		FirstRiftHerald struct {
			PlayerId  int64  `json:"player_id"`
			Timestamp int64  `json:"timestamp"`
			Team      string `json:"team"`
		} `json:"first_rift_herald"`
	} `json:"firsts"`
	Wards []struct {
		EventType string `json:"event_type"`
		Type      string `json:"type"`
		PlayerId  int64  `json:"player_id"`
		Timestamp int64  `json:"timestamp"`
	} `json:"wards"`
	KillTimeline []struct {
		Timestamp int64     `json:"timestamp"`
		Position  PosStruct `json:"position"`
		KillerId  int64     `json:"killer_id"`
		VictimId  int64     `json:"victim_id"`
		Assists   []int64   `json:"assists"`
	} `json:"kill_timeline"`
	ObjectiveEvents struct {
		Towers []struct {
			Type    string  `json:"type"`
			Assists []int64 `json:"assists"`
			LolLaneEventStruct
		} `json:"towers"`
		Inihibitors []struct {
			LolLaneEventStruct
			Assists []int64 `json:"assists"`
		} `json:"inhibitors"`
		Barons  []LolEventStruct `json:"barons"`
		Dragons []struct {
			Type string `json:"type"`
			LolEventStruct
		} `json:"dragons"`
		RiftHeralds []LolEventStruct `json:"rift_heralds"`
	} `json:"objective_events"`
	Draft []struct {
		RosterId int64 `json:"roster_id"`
		Champion struct {
			Name       string `json:"name"`
			ExternalId int64  `json:"external_id"`
		} `json:"champion"`
		Order *int64 `json:"order"`
		Type  string `json:"type"`
	} `json:"draft"`
}

type LolPlayerStruct struct {
	PlayerId   int64   `json:"player_id"`
	Role       string  `json:"role"`
	Lane       string  `json:"lane"`
	Kills      int64   `json:"kills"`
	Deaths     int64   `json:"deaths"`
	Assists    int64   `json:"assists"`
	GoldEarned int64   `json:"gold_earned"`
	GoldSpent  int64   `json:"gold_spent"`
	Gpm        float64 `json:"gpm"`
	TotalXp    int64   `json:"total_xp"`
	Xpm        float64 `json:"xpm"`
	Champion   struct {
		Name       string `json:"name"`
		ExternalId int64  `json:"external_id"`
	} `json:"champion"`
	KillCombos struct {
		Double              int64 `json:"double"`
		Triple              int64 `json:"triple"`
		Quadra              int64 `json:"quadra"`
		Penta               int64 `json:"penta"`
		Unreal              int64 `json:"unreal"`
		LargestKillingSpree int64 `json:"largest_killing_spree"`
		LargestMultiKill    int64 `json:"largest_multi_kill"`
		KillingSprees       int64 `json:"killing_sprees"`
	} `json:"kill_combos"`
	Items struct {
		Slot1 LolItemStruct `json:"slot_1"`
		Slot2 LolItemStruct `json:"slot_2"`
		Slot3 LolItemStruct `json:"slot_3"`
		Slot4 LolItemStruct `json:"slot_4"`
		Slot5 LolItemStruct `json:"slot_5"`
		Slot6 LolItemStruct `json:"slot_6"`
		Slot7 LolItemStruct `json:"slot_7"`
	} `json:"items"`
	Damage struct {
		Total        LolDmgStruct `json:"total"`
		ToHeroes     LolDmgStruct `json:"to_heroes"`
		DamageTaken  LolDmgStruct `json:"damage_taken"`
		LargestCrit  LolDmgStruct `json:"largest_crit"`
		ToObjectives LolDmgStruct `json:"to_objectives"`
		ToTurrets    LolDmgStruct `json:"to_turrets"`
	} `json:"damage"`
	Support struct {
		AmountHealed     int64 `json:"amount_healed"`
		UnitsHealed      int64 `json:"units_healed"`
		CrowsControlTime int64 `json:"crowd_control_time"`
	} `json:"support"`
	MinionKills struct {
		Total              int64 `json:"total"`
		NeutralJungle      int64 `json:"neutral_jungle"`
		NeutralEnemyJungle int64 `json:"neutral_enemy_jungle"`
	} `json:"minion_kills"`
	Wards []struct {
		Type      string `json:"type"`
		Destroyed int64  `json:"destroyed"`
		Placed    int64  `json:"placed"`
	} `json:"wards"`
	ChampionSpells map[string]struct {
		Name string `json:"name"`
	} `json:"champion_spells"`
	Masteries []struct {
		Name string `json:"name"`
		Rank int64  `json:"rank"`
	} `json:"masteries"`
	Runes []struct {
		Name string `json:"name"`
		Rank int64  `json:"rank"`
	} `json:"runes"`
	RunesReforged struct {
		PrimaryPath struct {
			Path struct {
				Name       string `json:"name"`
				ExternalId int64  `json:"internal_id"`
			} `json:"path"`
			Keystone struct {
				Name       string `json:"name"`
				ExternalId int64  `json:"internal_id"`
			} `json:"keystone"`
			Rune1 struct {
				Name       string `json:"name"`
				ExternalId int64  `json:"internal_id"`
			} `json:"rune_1"`
			Rune2 struct {
				Name       string `json:"name"`
				ExternalId int64  `json:"internal_id"`
			} `json:"rune_2"`
			Rune3 struct {
				Name       string `json:"name"`
				ExternalId int64  `json:"internal_id"`
			} `json:"rune_3"`
		} `json:"primary_path"`
		SecondaryPath struct {
			Path struct {
				Name       string `json:"name"`
				ExternalId int64  `json:"internal_id"`
			} `json:"path"`
			Rune1 struct {
				Name       string `json:"name"`
				ExternalId int64  `json:"internal_id"`
			} `json:"rune_1"`
			Rune2 struct {
				Name       string `json:"name"`
				ExternalId int64  `json:"internal_id"`
			} `json:"rune_2"`
		} `json:"secondary_path"`
	} `json:"runes_reforged"`
	Skillups []struct {
		Time        int64  `json:"time"`
		AbilitySlot int64  `json:"ability_slot"`
		Type        string `json:"type"`
	} `json:"skillups"`
}

type LolDmgStruct struct {
	Magic    int64 `json:"magic"`
	Physical int64 `json:"physical"`
	True     int64 `json:"true"`
}

type LolEventStruct struct {
	Timestamp int64     `json:"timestamp"`
	Position  PosStruct `json:"position"`
	KillerId  int64     `json:"killer_id"`
}

type LolLaneEventStruct struct {
	Lane string `json:"lane"`
	LolEventStruct
}
