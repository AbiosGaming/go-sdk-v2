package structs

// CasterStruct represents an individual shoutcaster, casting a Series.
type CasterStruct struct {
	Name    string        `json:"name,omitempty"`
	Type    *int64        `json:"type,omitempty"`
	Url     string        `json:"url,omitempty"`
	Stream  StreamStruct  `json:"stream,omitempty"`
	Country CountryStruct `json:"country,omitempty"`
}
