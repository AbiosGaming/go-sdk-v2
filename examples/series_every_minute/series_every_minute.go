package main

import (
	"fmt"

	"github.com/AbiosGaming/go-sdk-v2/v3"
)

func main() {
	a := abios.New("username", "password")
	a.SetRate(0, 1)

	parameters := make(abios.Parameters)
	parameters.Add("games[]", "1") // Only get series' for Dota
	for {
		series, err := a.Series(parameters)
		if err != nil {
			fmt.Println(err)
			return
		}
		// Do something with series
		fmt.Println(series.Data[0].Title, series.Data[0].BestOf)
	}
}
