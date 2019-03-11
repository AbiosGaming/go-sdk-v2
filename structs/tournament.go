package structs

// PaginatedTournaments holds a list of Tournament as well as information about pages
type PaginatedTournaments struct {
	LastPage    int64        `json:"last_page,omitempty"`
	CurrentPage int64        `json:"current_page,omitempty"`
	Data        []Tournament `json:"data,omitempty"`
}

// Tournament represents a tournament (i.e a structured group of Series').
type Tournament struct {
	Id               int64            `json:"id,omitempty"`
	Title            string           `json:"title,omitempty"`
	ShortTitle       string           `json:"short_title,omitempty"`
	Country          Country          `json:"country,omitempty"`
	City             string           `json:"city,omitempty"`
	Tier             int64            `json:"tier"`
	Description      string           `json:"description,omitempty"`
	ShortDescription string           `json:"short_description,omitempty"`
	Format           string           `json:"format,omitempty"`
	Start            *string          `json:"start"`      // Datettime
	End              *string          `json:"end"`        // Datettime
	DeletedAt        *string          `json:"deleted_at"` // Datettime
	Url              string           `json:"url,omitempty"`
	HasPbpStats      bool             `json:"has_pbpstats"`
	Images           TournamentImages `json:"images,omitempty"`
	PrizepoolString  Prizepool        `json:"prizepool_string,omitempty"`
	Links            Links            `json:"links,omitempty"`
	NextSeries       *Series          `json:"next_series"`
	Series           []Series         `json:"series,omitempty"`  // Optional
	Stages           []Stage          `json:"stages,omitempty"`  // Optional
	Rosters          []Roster         `json:"rosters,omitempty"` // Optional
	Game             Game             `json:"game,omitempty"`
	Casters          []Caster         `json:"casters"`
}

// Prizepool holds information about a Tournament's prizepool.
type Prizepool struct {
	Total  string `json:"total,omitempty"`
	First  string `json:"first,omitempty"`
	Second string `json:"second,omitempty"`
	Third  string `json:"third,omitempty"`
}

// Links holds information about links relevant to a Tournament.
type Links struct {
	Website string `json:"website"`
	Youtube string `json:"youtube"`
}
