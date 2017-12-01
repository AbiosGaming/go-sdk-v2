package main

import (
	"fmt"
	"time"

	"github.com/AbiosGaming/go-sdk-v2"
)

func main() {
	a := abios.New("username", "password")

	go series(a)
	games(a)
}

// series queries the /series endpoint once every minute.
func series(a abios.AbiosSdk) {
	lastRequest := time.Now()
	for {
		series, err := a.Series(nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Series received:", len(series.Data), " --- time since last request:", time.Now().Sub(lastRequest))
		lastRequest = time.Now()
		time.Sleep(1 * time.Minute)
	}
}

// games queries the /games endpoint once every thirty seconds.
func games(a abios.AbiosSdk) {
	lastRequest := time.Now()
	for {
		games, err := a.Games(nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Games received:", len(games.Data), " --- time since last request:", time.Now().Sub(lastRequest))
		lastRequest = time.Now()
		time.Sleep(30 * time.Second)
	}
}
