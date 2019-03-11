package structs

// Map represents the map being played in a Match
type Map struct {
	Id       int64  `json:"id"`
	Name     string `json:"name,omitempty"`
	Official bool   `json:"official"`
	Game     Game   `json:"game,omitempty"`
}
