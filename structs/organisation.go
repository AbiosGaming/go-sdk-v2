package structs

// OrganisationStructPaginated holds a list of TeamStruct as well as information about pages.
type OrganisationsStructPaginated struct {
	LastPage    int64                `json:"last_page,omitempty"`
	CurrentPage int64                `json:"current_page,omitempty"`
	Data        []OrganisationStruct `json:"data,omitempty"`
}

type OrganisationStruct struct {
	Id    int64        `json:"id,omitempty"`
	Name  string       `json:"name,omitempty"`
	Teams []TeamStruct `json:"teams,omitempty"`
}
