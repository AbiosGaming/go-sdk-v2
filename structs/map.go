package structs

// MapStruct represents the map being played in a Match
type MapStruct struct {
	Id   int64      `json:"id"`
	Name string     `json:"name,omitempty"`
	Game GameStruct `json:"game,omitempty"`
}
