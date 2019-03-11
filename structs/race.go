package structs

// Race represents a race, faction, class etc. in a Game.
type Race struct {
	Id     int64      `json:"id,omitempty"`
	GameId int64      `json:"game_id,omitempty"`
	Name   string     `json:"name,omitempty"`
	Images RaceImages `json:"images,omitempty"`
}
