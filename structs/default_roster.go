package structs

/* DefaultRosterStruct represents the time period(s) when a Roster has been a Team's
 * main roster or line-up.
 */
type DefaultRosterStruct struct {
	From   string       `json:"from,omitempty"`
	To     *string      `json:"to,omitempty"`
	Roster RosterStruct `json:"roster,omitempty"`
}
