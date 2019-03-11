package structs

// Country hold information about a country, nationality or language
// associated with a resource.
type Country struct {
	Name      string        `json:"name,omitempty"`
	ShortName string        `json:"short_name,omitempty"`
	Images    CountryImages `json:"images,omitempty"`
}
