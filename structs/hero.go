package structs

type HeroStruct struct {
	Name      string `json:"name"`
	Attribute string `json:"attribute"`
	Images    struct {
		Large string `json:"large"`
		Small string `json:"small"`
	} `json:"images"`
}
