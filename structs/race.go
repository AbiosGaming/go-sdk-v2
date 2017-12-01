package structs

// RaceStruct represents a race, faction, class etc. in a Game.
type RaceStruct struct {
	Id     int64            `json:"id,omitempty"`
	GameId int64            `json:"game_id,omitempty"`
	Name   string           `json:"name,omitempty"`
	Images RaceImagesStruct `json:"images,omitempty"`
}
