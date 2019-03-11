/*
Package structs contains the structs AbiosGaming/sdk uses to unmarshal JSON.
*/
package structs

// PaginatedGames holds a list of Game as well as information about pages.
type PaginatedGames struct {
	LastPage    int64  `json:"last_page,omitempty"`
	CurrentPage int64  `json:"current_page,omitempty"`
	Data        []Game `json:"data,omitempty"`
}

// Game represents the actual game being played in a Series
type Game struct {
	Id        int64      `json:"id,omitempty"`
	Title     string     `json:"title,omitempty"`
	LongTitle string     `json:"long_title,omitempty"`
	DeletedAt *string    `json:"deleted_at"` // Datettime
	Images    GameImages `json:"images,omitempty"`
	Color     string     `json:"color,omitempty"`
}
