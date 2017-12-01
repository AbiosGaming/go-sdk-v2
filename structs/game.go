/*
Package structs contains the structs AbiosGaming/sdk uses to unmarshal JSON.
*/
package structs

// GameStructPaginated holds a list of GameStruct as well as information about pages.
type GameStructPaginated struct {
	LastPage    int64        `json:"last_page,omitempty"`
	CurrentPage int64        `json:"current_page,omitempty"`
	Data        []GameStruct `json:"data,omitempty"`
}

// GameStruct represents the actual game being played in a Series
type GameStruct struct {
	Id        int64            `json:"id,omitempty"`
	Title     string           `json:"title,omitempty"`
	LongTitle string           `json:"long_title,omitempty"`
	DeletedAt *string          `json:"deleted_at"` // Datettime
	Images    GameImagesStruct `json:"images,omitempty"`
	Color     string           `json:"color,omitempty"`
}
