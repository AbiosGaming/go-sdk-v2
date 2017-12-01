package structs

// SportsbookOddsStruct is a top-level struct holding information about sportsbook odds.
type SportsbookOddsStruct struct {
	Sportsbook string          `json:"sportsbook"`
	Link       string          `json:"link"`
	Moneyline  MoneylineStruct `json:"moneyline"`
}

// MoneylineStruct holds information about the mouneyline for a particular sportsbook.
type MoneylineStruct struct {
	Home        float64  `json:"home"`
	HomeBetSlip *string  `json:"home_bet_slip"`
	Away        float64  `json:"away"`
	AwayBetSlip *string  `json:"away_bet_slip"`
	Draw        *float64 `json:"draw"`
	DrawBetSlip *string  `json:"draw_bet_slip"`
}
