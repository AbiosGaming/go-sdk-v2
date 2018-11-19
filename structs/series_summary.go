package structs

type SeriesSummary interface{}

type DotaSeriesSummary struct {
	Scoreboard map[string][]struct {
		PlayerID      int64   `json:"player_id"`
		MatchesPlayed int64   `json:"matches_played"`
		Kills         int64   `json:"kills"`
		Deaths        int64   `json:"deaths"`
		Assists       int64   `json:"assists"`
		Gpm           float64 `json:"gpm"`
		Xpm           float64 `json:"xpm"`
		CreepKills    int64   `json:"creep_kills"`
		CreepDenies   int64   `json:"creep_denies"`
		HeroDmg       struct {
			Given struct {
				ByHero struct {
					HpRemoval   int64 `json:"hp_removal"`
					MagicalDmg  int64 `json:"magical_dmg"`
					PhysicalDmg int64 `json:"physical_dmg"`
					PureDmg     int64 `json:"pure_dmg"`
				} `json:"by_hero"`
				ByMobs struct {
					HpRemoval   int64 `json:"hp_removal"`
					MagicalDmg  int64 `json:"magical_dmg"`
					PhysicalDmg int64 `json:"physical_dmg"`
					PureDmg     int64 `json:"pure_dmg"`
				} `json:"by_mobs"`
			} `json:"given"`
			Taken struct {
				FromHeroes struct {
					HpRemoval   int64 `json:"hp_removal"`
					MagicalDmg  int64 `json:"magical_dmg"`
					PhysicalDmg int64 `json:"physical_dmg"`
					PureDmg     int64 `json:"pure_dmg"`
				} `json:"from_heroes"`
				FromMobs struct {
					HpRemoval   int64 `json:"hp_removal"`
					MagicalDmg  int64 `json:"magical_dmg"`
					PhysicalDmg int64 `json:"physical_dmg"`
					PureDmg     int64 `json:"pure_dmg"`
				} `json:"from_mobs"`
			} `json:"taken"`
		} `json:"hero_dmg"`
	} `json:"scoreboard"`
}

type CsSeriesSummary struct {
	Scoreboard map[string][]struct {
		PlayerID      int64   `json:"player_id"`
		MatchesPlayed int64   `json:"matches_played"`
		Kills         int64   `json:"kills"`
		Deaths        int64   `json:"deaths"`
		Assists       int64   `json:"assists"`
		Adr           float64 `json:"adr"`
	}
}

type LolSeriesSummary struct {
	Scoreboard map[string][]struct {
		PlayerID      int64   `json:"player_id"`
		MatchesPlayed int64   `json:"matches_played"`
		Kills         int64   `json:"kills"`
		Deaths        int64   `json:"deaths"`
		Assists       int64   `json:"assists"`
		GoldEarned    int64   `json:"gold_earned"`
		GoldSpent     int64   `json:"gold_spent"`
		TotalXp       int64   `json:"total_xp"`
		Xpm           float64 `json:"xpm"`
		Gpm           float64 `json:"gpm"`
		KillCombos    struct {
			Double              int64 `json:"double"`
			Triple              int64 `json:"triple"`
			Quadra              int64 `json:"quadra"`
			Penta               int64 `json:"penta"`
			Unreal              int64 `json:"unreal"`
			LargestKillingSpree int64 `json:"largest_killing_spree"`
			LargestMultiKill    int64 `json:"largest_multi_kill"`
			KillingSprees       int64 `json:"killing_sprees"`
		} `json:"kill_combos"`
		MinionKills struct {
			Total              int64 `json:"total"`
			NeutralJungle      int64 `json:"neutral_jungle"`
			NeutralEnemyJungle int64 `json:"neutral_enemy_jungle"`
		} `json:"minion_kills"`
		Damage struct {
			Total struct {
				Magic    int64 `json:"magic"`
				Physical int64 `json:"physical"`
				True     int64 `json:"true"`
			} `json:"total"`
			ToHeroes struct {
				Magic    int64 `json:"magic"`
				Physical int64 `json:"physical"`
				True     int64 `json:"true"`
			} `json:"to_heroes"`
			DamageTaken struct {
				Magic    int64 `json:"magic"`
				Physical int64 `json:"physical"`
				True     int64 `json:"true"`
			} `json:"damage_taken"`
			LargestCrit  int64 `json:"largest_crit"`
			ToObjectives int64 `json:"to_objectives"`
			ToTurrets    int64 `json:"to_turrets"`
		} `json:"damage"`
		Support struct {
			AmountHealed     int64 `json:"amount_healed"`
			UnitsHealed      int64 `json:"units_healed"`
			CrowdControlTime int64 `json:"crowd_control_time"`
		} `json:"support"`
	}
}
