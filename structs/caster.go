package structs

// CasterStruct represents an individual shoutcaster, casting a Series.
type CasterStruct struct {
	Id      int64         `json:"id,omitempty"`
	Name    string        `json:"name,omitempty"`
	Type    *int64        `json:"type,omitempty"`
	Url     string        `json:"url,omitempty"`
	Primary bool          `json:"primary,omitmepty"`
	Stream  *StreamStruct `json:"stream,omitempty"`
	Country CountryStruct `json:"country,omitempty"`
}
