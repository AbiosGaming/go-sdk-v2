package structs

// DotaMatchSummary is the summarization of a Dota match.
type DotaMatchSummary struct {
	RadiantRoster int64                  `json:"radiant_roster"`
	DireRoster    int64                  `json:"dire_roster"`
	MatchLength   int64                  `json:"match_length"`
	DraftSeq      []DotaDraftEvent       `json:"draft_seq"`
	FirstBlood    DotaFirstBloodEvent    `json:"first_blood"`
	Kills         []DotaMatchKill        `json:"kills"`
	StructureDest []DotaStructureDest    `json:"structure_dest"`
	PlayerStats   []DotaMatchPlayerStats `json:"player_stats"`
	RoshanEvents  []RoshanEvent          `json:"roshan_events"`
}

// DotaMatchPlayerStats is the summarization of a Player in Dota match.
type DotaMatchPlayerStats struct {
	PlayerId     int64             `json:"player_id"`
	Hero         Hero              `json:"hero"`
	Kills        int64             `json:"kills"`
	Deaths       int64             `json:"deaths"`
	Assists      int64             `json:"assists"`
	Gpm          float64           `json:"gpm"`
	Networth     int64             `json:"networth"`
	Xpm          float64           `json:"xpm"`
	Levels       map[string]int64  `json:"levels"`
	CreepKills   int64             `json:"creep_kills"`
	CreepDenies  int64             `json:"creep_denies"`
	CampsStacked int64             `json:"camps_stacked"`
	HeroDmg      DotaPlayerDmg     `json:"hero_dmg"`
	HeroHealing  DotaPlayerHealing `json:"hero_healing"`
	Wards        DotaWardsStat     `json:"wards"`
	Runes        DotaRunesStat     `json:"runes"`
	Items        struct {
		Inventory struct {
			Slot1 DotaItem `json:"slot_1"`
			Slot2 DotaItem `json:"slot_2"`
			Slot3 DotaItem `json:"slot_3"`
			Slot4 DotaItem `json:"slot_4"`
			Slot5 DotaItem `json:"slot_5"`
			Slot6 DotaItem `json:"slot_6"`
		} `json:"inventory"`
		Backpack struct {
			Slot1 DotaItem `json:"slot_1"`
			Slot2 DotaItem `json:"slot_2"`
			Slot3 DotaItem `json:"slot_3"`
		} `json:"backpack"`
		Stash struct {
			Slot1 DotaItem `json:"slot_1"`
			Slot2 DotaItem `json:"slot_2"`
			Slot3 DotaItem `json:"slot_3"`
			Slot4 DotaItem `json:"slot_4"`
			Slot5 DotaItem `json:"slot_5"`
			Slot6 DotaItem `json:"slot_6"`
		} `json:"stash"`
	} `json:"items"`
}

// DotaPlayerDmg is a collection of damage stat given or taken by player in Dota match.
type DotaPlayerDmg struct {
	Given struct {
		ByHero DotaDmg `json:"by_hero"`
		ByMobs DotaDmg `json:"by_mobs"`
	} `json:"given"`
	Taken struct {
		FromHeroes DotaDmg `json:"from_heroes"`
		FromMobs   DotaDmg `json:"from_mobs"`
	} `json:"taken"`
}

// DotaPlayerDmg is a collection of damage stat given or taken by player in Dota match.
type DotaPlayerHealing struct {
	Given struct {
		ByHero int64 `json:"by_hero"`
		ByMobs int64 `json:"by_mobs"`
	} `json:"given"`
	Taken struct {
		FromHeroes int64 `json:"from_heroes"`
		FromMobs   int64 `json:"from_mobs"`
	} `json:"taken"`
}

// DotaFirstBloodEvent is a information about first blood in Dota match.
type DotaFirstBloodEvent struct {
	Killer int64 `json:"killer"`
	Victim int64 `json:"victim"`
	AtTime int64 `json:"at_time"`
}

// DotaDraftEvent is information about event of draft phase of Dota match.
type DotaDraftEvent struct {
	Order    int64  `json:"order"`
	Type     string `json:"type"`
	RosterId int64  `json:"roster_id"`
	Hero     Hero   `json:"hero"`
}

// DotaMatchKill is a summarization of data about kill event of Dota match.
type DotaMatchKill struct {
	Killer  int64   `json:"killer"`
	Victim  int64   `json:"victim"`
	AtTime  int64   `json:"at_time"`
	Assists []int64 `json:"assists"`
}

// DotaStructureDest is a summarization of data about structure destruction event of Dota match.
type DotaStructureDest struct {
	Killer        int64  `json:"killer"`
	StructureType string `json:"structure_type"`
	StructurePos  string `json:"structure_pos"`
	AtTime        int64  `json:"at_time"`
}

// RoshanEvent is a summarization of data about Roshan event of Dota match.
type RoshanEvent struct {
	Type   string `json:"type"`
	Killer int64  `json:"killer"`
	AtTime int64  `json:"at_time"`
}

// DotaDmg is a collection of different types of damages given/taken in dota.
type DotaDmg struct {
	HpRemoval   int64 `json:"hp_removal"`
	MagicalDmg  int64 `json:"magical_dmg"`
	PhysicalDmg int64 `json:"physical_dmg"`
	PureDmg     int64 `json:"pure_dmg"`
}

// DotaWardsStat is a collection of data about wards actions of a player
type DotaWardsStat struct {
	Observers struct {
		Killed int64 `json:"killed"`
		Placed int64 `json:"placed"`
	} `json:"observers"`
	Sentries struct {
		Killed int64 `json:"killed"`
		Placed int64 `json:"placed"`
	} `json:"sentries"`
}

// DotaRunesStat is a collection of data about picked up runes
type DotaRunesStat struct {
	DoubleDamageRunes int64 `json:"double_damage_runes"`
	HasteRunes        int64 `json:"haste_runes"`
	IllusionRunes     int64 `json:"illusion_runes"`
	InvisibilityRunes int64 `json:"invisibility_runes"`
	RegenerationRunes int64 `json:"regeneration_runes"`
	BountyRunes       int64 `json:"bounty_runes"`
	ArcaneRunes       int64 `json:"arcane_runes"`
}
