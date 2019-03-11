package structs

// DefaultRoster represents the time period(s) when a Roster has been a Team's
// main roster or line-up.
type DefaultRoster struct {
	From   string  `json:"from,omitempty"`
	To     *string `json:"to,omitempty"`
	Roster Roster  `json:"roster,omitempty"`
}
