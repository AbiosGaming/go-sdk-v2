package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/AbiosGaming/go-sdk-v2"
)

// currentCalendar holds what we respond to our users
var currentCalendar []calendarEntry

func main() {
	// Add our credentials to the SDK
	a := abios.New("username", "password")

	// Set out outgoing rate to once every minute so we don't query too much.
	a.SetRate(0, 1)

	// We want to update what we respond with once every minute. The SDK makes sure that
	// the rate we specified above isn't exceeded.
	go func(a abios.AbiosSdk) {
		for {
			currentCalendar = getSeries(a)
		}
	}(a)

	// Setup webserver
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

// handler marshals the currentCalendar as JSON and writes to the requester
func handler(w http.ResponseWriter, r *http.Request) {
	payload, err := json.MarshalIndent(currentCalendar, "", "\t")
	if err != nil {
		log.Println("Unable to marshal response")
		return // Since we couldn't marshal proper JSON we don't want to write anything
	}
	w.Write(payload)
}

// calendarEntry holds all data neccesary to display a series in a calendar
type calendarEntry struct {
	StartTime       string `json:"start_time"`
	Roster1         string `json:"home"`
	Roster2         string `json:"away"`
	BestOf          int64  `json:"best_of"`
	Title           string `json:"title"`
	TournamentTitle string `json:"tournament_title"`
}

// getSeries queries the abios /series endpoint and returns the relevant data for each
// series.
func getSeries(a abios.AbiosSdk) (calendarEntries []calendarEntry) {
	p := make(abios.Parameters)
	p.Add("with[]", "tournament")
	series, err := a.Series(p) // The actual API request
	if err != nil {
		log.Println("Couldn't get series from Abios:", err)
		return currentCalendar // Return the last update.
	}

	// We are given a paginated result
	for _, data := range series.Data {
		// Determine rosters.
		rosters := [2]string{"TBD", "TBD"} // Holds the team name or "TBD"

		// Roster.Teams is usually (but not necessarily) one or empty
		for i, roster := range data.Rosters {
			if len(data.Rosters[i].Teams) > 0 {
				rosters[i] = roster.Teams[0].Name
			}
		}

		// Determine start time. Can be nil if it hasn't been announced yet.
		startTime := "TBD"
		if data.Start != nil {
			startTime = *data.Start
		}

		calendarEntries = append(calendarEntries, calendarEntry{
			StartTime:       startTime,
			Roster1:         rosters[0],
			Roster2:         rosters[1],
			BestOf:          data.BestOf,
			Title:           data.Title,
			TournamentTitle: data.Tournament.Title,
		})
	}
	return
}
