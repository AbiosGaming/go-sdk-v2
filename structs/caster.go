package structs

// Caster represents an individual shoutcaster, casting a Series.
type Caster struct {
	Id      int64   `json:"id,omitempty"`
	Name    string  `json:"name,omitempty"`
	Type    *int64  `json:"type,omitempty"`
	Url     string  `json:"url,omitempty"`
	Primary bool    `json:"primary,omitmepty"`
	Stream  *Stream `json:"stream,omitempty"`
	Country Country `json:"country,omitempty"`
}
