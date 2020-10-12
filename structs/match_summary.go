package structs

// MatchSummary holds information about play by play statistics for a certain match.
type MatchSummary interface{}

// CsMatchSummary is the summarization of a CS:GO match.
type CsMatchSummary struct {
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
			Type       string `json:"type"`
			PlayerId   int64  `json:"player_id"`
			RoundClock int64  `json:"round_clock"`
			Pos        Pos    `json:"pos"`
		} `json:"bomb_events"`
		Kills []struct {
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
		} `json:"kills"`
		PlayerStats struct {
			TSide  []RoundPlayerStats `json:"t_side"`
			CtSide []RoundPlayerStats `json:"ct_side"`
		} `json:"player_stats"`
	} `json:"rounds"`
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

// Pos hold x, y and z coordinates.
type Pos struct {
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

type LolMatchSummary struct {
	MatchLength int64 `json:"match_length"`
	BlueRoster  struct {
		Id      int64       `json:"id"`
		Players []LolPlayer `json:"players"`
	} `json:"blue_roster"`
	PurpleRoster struct {
		Id      int64       `json:"id"`
		Players []LolPlayer `json:"players"`
	} `json:"purple_roster"`
	Firsts struct {
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
		Timestamp int64   `json:"timestamp"`
		Position  Pos     `json:"position"`
		KillerId  int64   `json:"killer_id"`
		VictimId  int64   `json:"victim_id"`
		Assists   []int64 `json:"assists"`
	} `json:"kill_timeline"`
	ObjectiveEvents struct {
		Towers []struct {
			Type    string  `json:"type"`
			Assists []int64 `json:"assists"`
			LolLaneEvent
		} `json:"towers"`
		Inihibitors []struct {
			LolLaneEvent
			Assists []int64 `json:"assists"`
		} `json:"inhibitors"`
		Barons  []LolEvent `json:"barons"`
		Dragons []struct {
			Type string `json:"type"`
			LolEvent
		} `json:"dragons"`
		RiftHeralds []LolEvent `json:"rift_heralds"`
	} `json:"objective_events"`
	Draft []struct {
		RosterId int64    `json:"roster_id"`
		Champion Champion `json:"champion"`
		Order    *int64   `json:"order"`
		Type     string   `json:"type"`
	} `json:"draft"`
}

type LolPlayer struct {
	PlayerId   int64    `json:"player_id"`
	Role       string   `json:"role"`
	Lane       string   `json:"lane"`
	Kills      int64    `json:"kills"`
	Deaths     int64    `json:"deaths"`
	Assists    int64    `json:"assists"`
	GoldEarned int64    `json:"gold_earned"`
	GoldSpent  int64    `json:"gold_spent"`
	Gpm        float64  `json:"gpm"`
	TotalXp    int64    `json:"total_xp"`
	Xpm        float64  `json:"xpm"`
	Champion   Champion `json:"champion"`
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
		Inventory struct {
			Slot1 LolItem `json:"slot_1"`
			Slot2 LolItem `json:"slot_2"`
			Slot3 LolItem `json:"slot_3"`
			Slot4 LolItem `json:"slot_4"`
			Slot5 LolItem `json:"slot_5"`
			Slot6 LolItem `json:"slot_6"`
			Slot7 LolItem `json:"slot_7"`
		} `json:"inventory"`
	} `json:"items"`
	Damage struct {
		Total        LolDmg `json:"total"`
		ToHeroes     LolDmg `json:"to_heroes"`
		DamageTaken  LolDmg `json:"damage_taken"`
		LargestCrit  int64  `json:"largest_crit"`
		ToObjectives int64  `json:"to_objectives"`
		ToTurrets    int64  `json:"to_turrets"`
	} `json:"damage"`
	Support struct {
		AmountHealed     int64 `json:"amount_healed"`
		UnitsHealed      int64 `json:"units_healed"`
		CrowdControlTime int64 `json:"crowd_control_time"`
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
				ExternalId int64  `json:"external_id"`
			} `json:"path"`
			Keystone struct {
				Name       string `json:"name"`
				ExternalId int64  `json:"external_id"`
			} `json:"keystone"`
			Rune1 struct {
				Name       string `json:"name"`
				ExternalId int64  `json:"external_id"`
			} `json:"rune_1"`
			Rune2 struct {
				Name       string `json:"name"`
				ExternalId int64  `json:"external_id"`
			} `json:"rune_2"`
			Rune3 struct {
				Name       string `json:"name"`
				ExternalId int64  `json:"external_id"`
			} `json:"rune_3"`
		} `json:"primary_path"`
		SecondaryPath struct {
			Path struct {
				Name       string `json:"name"`
				ExternalId int64  `json:"external_id"`
			} `json:"path"`
			Rune1 struct {
				Name       string `json:"name"`
				ExternalId int64  `json:"external_id"`
			} `json:"rune_1"`
			Rune2 struct {
				Name       string `json:"name"`
				ExternalId int64  `json:"external_id"`
			} `json:"rune_2"`
		} `json:"secondary_path"`
	} `json:"runes_reforged"`
	Skillups []struct {
		Time        int64  `json:"time"`
		AbilitySlot int64  `json:"ability_slot"`
		Type        string `json:"type"`
	} `json:"skillups"`
}

type LolDmg struct {
	Magic    int64 `json:"magic"`
	Physical int64 `json:"physical"`
	True     int64 `json:"true"`
}

type LolEvent struct {
	Timestamp int64 `json:"timestamp"`
	Position  Pos   `json:"position"`
	KillerId  int64 `json:"killer_id"`
}

type LolLaneEvent struct {
	Lane string `json:"lane"`
	LolEvent
}
