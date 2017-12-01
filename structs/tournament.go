package structs

// TournamentStructPaginated holds a list of TournamentStruct as well as information about pages
type TournamentStructPaginated struct {
	LastPage    int64              `json:"last_page,omitempty"`
	CurrentPage int64              `json:"current_page,omitempty"`
	Data        []TournamentStruct `json:"data,omitempty"`
}

// TournamentStruct represents a tournament (i.e a structured group of Series').
type TournamentStruct struct {
	Id               int64                  `json:"id,omitempty"`
	Title            string                 `json:"title,omitempty"`
	ShortTitle       string                 `json:"short_title,omitempty"`
	Country          CountryStruct          `json:"country,omitempty"`
	City             string                 `json:"city,omitempty"`
	Tier             int64                  `json:"tier"`
	Description      string                 `json:"description,omitempty"`
	ShortDescription string                 `json:"short_description,omitempty"`
	Format           string                 `json:"format,omitempty"`
	Start            *string                `json:"start"`      // Datettime
	End              *string                `json:"end"`        // Datettime
	DeletedAt        *string                `json:"deleted_at"` // Datettime
	Url              string                 `json:"url,omitempty"`
	Images           TournamentImagesStruct `json:"images,omitempty"`
	PrizepoolString  PrizepoolStruct        `json:"prizepool_string,omitempty"`
	Links            LinksStruct            `json:"links,omitempty"`
	NextSeries       *SeriesStruct          `json:"next_series"`
	Series           []SeriesStruct         `json:"series,omitempty"`  // Optional
	Stages           []StageStruct          `json:"stages,omitempty"`  // Optional
	Rosters          []RosterStruct         `json:"rosters,omitempty"` // Optional
	Game             GameStruct             `json:"game,omitempty"`
	Casters          []CasterStruct         `json:"casters"`
}

// PrizepoolStruct holds information about a Tournament's prizepool.
type PrizepoolStruct struct {
	Total  string `json:"total,omitempty"`
	First  string `json:"first,omitempty"`
	Second string `json:"second,omitempty"`
	Third  string `json:"third,omitempty"`
}

// LinksStruct holds information about links relevant to a Tournament.
type LinksStruct struct {
	Website string `json:"website"`
	Youtube string `json:"youtube"`
}
