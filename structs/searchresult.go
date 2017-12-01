package structs

// SearchResultStruct represents a result from the /search endpoint.
type SearchResultStruct struct {
	Id       int64   `json:"id,omitempty"`
	Matched  string  `json:"matched,omitempty"`
	AltLabel string  `json:"alt_label,omitempty"`
	Type     string  `json:"type,omitempty"`
	Logo     string  `json:"logo,omitempty"`
	GameId   int64   `json:"game_id,omitempty"`
	GameLogo *string `json:"game_logo"`
}
