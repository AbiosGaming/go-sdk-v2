package main

import (
	"fmt"
	"time"

	"github.com/AbiosGaming/go-sdk-v2/v3"
)

func main() {
	a := abios.New("username", "password")

	thirtyMinutesAgo := time.Now().Add(-time.Minute * 30).UTC().Format("2006-01-02T15:04:05Z")

	parameters := make(abios.Parameters)
	parameters.Add("starts_after", thirtyMinutesAgo)

	series, err := a.Series(parameters)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, s := range series.Data {
		for _, roster := range s.Rosters {
			// roster.Teams is either of length 1 or empty.
			for _, team := range roster.Teams {
				// Get the team_stats from /teams:id endpoint.
				parameters.Del("starts_after")
				parameters.Add("with[]", "team_stats")
				teamWithStats, err := a.TeamsById(team.Id, parameters)
				if err != nil {
					fmt.Println(err)
					continue
				}
				seriesWinrate := teamWithStats.TeamStats.Winrate.Series
				fmt.Printf("%v has a winrate of %.2f%% over %v series in %v and their latest match started %v\n",
					teamWithStats.Name,
					seriesWinrate.Rate*100,
					seriesWinrate.History,
					s.Game.LongTitle,
					*s.Start)
			}
		}
	}
}
