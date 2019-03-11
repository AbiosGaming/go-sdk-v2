package structs

// Platform represents a third party streaming platform.
type Platform struct {
	Id     int64          `json:"id,omitempty"`
	Name   string         `json:"name,omitempty"`
	Color  string         `json:"color,omitempty"`
	Images PlatformImages `json:"images,omitempty"`
}
