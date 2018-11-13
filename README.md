Abios SDK
=========

Abios SDK is meant to provide a thin wrapper around http requests and unmarshal
the resulting JSON into an appropriate struct. You can think of the SDK as a collection
of structs (each of which corresponds to an Abios endpoint) combined with methods that
unmarshals and type checks.

The SDK also allows you to specify the rate of outgoing requests. I.e you can force it
to never send requests at a higher rate than you pay for.

It also automatically handles re-authentication when the token is about to expire.

# Documentation

Each method and struct is documented according to [Effective Go](https://golang.org/doc/effective_go.html).
However, the documentation in the source code merely describes what the function does,
not how the endpoint itself functions. To see documentation relating to the endpoints
themselves see the [official documentation](https://docs.abiosgaming.com/).

# Bugs, Issues and General Shortcomings

Find a bug? Missing a feature? Create an issue (or if you are feeling particularly ambitious,
create a pull request) and we'll get to it as soon as possible!

# Installation

```Bash
$ go get -u github.com/AbiosGaming/go-sdk-v2/
```

## Quick Start
Add the import line:

```Go
import "github.com/AbiosGaming/go-sdk-v2"
```

Use the function `abios.New(username, password string)` to create a new instance of the
abios struct and authenticate with the given credentials.

```Go
a := abios.New("username", "password")
```

To set the outgoing rate use the `abios.SetRate(second, minute int)` function like so:

```Go
a.SetRate(5, 300)
```

This will allow you to send 5 requests every second and up to 300 requests every minute.
See [Outgoing Rate](#rate) for more information.

To get all available games (from the /games endpoint) the following code will do:

```Go
games, err := a.Games(nil)
```

To only get information about a particular game one would just add the parameter:

```Go
parameters := make(abios.Parameters) // Remember to make.
parameters.Add("q", "counter")
games = a.Games(parameters)
```

For more information about parameters see [Parameters](#parameters).

# Authentication
Authentication is done in the constructor. The constructor will query the
/oauth/access\_token endpoint and store the resulting access token internally.
This will then be automatically added to all outgoing requests.

In addition, the credentials will be stored in memory and used
to query for a new token when there is 10 minutes or less until the current one expires.

If the initial authentication fails the error message returned will override all subsequent
responses. If an authentication request other than the intial one fails the SDK will re-try
every 30 seconds for 5 minutes, after which all responses will be overriden (and thus return an error).

# <a name="rate"></a>Outgoing Rate
Allowing the specification of an outgoing rate is to minimize the number of "429 (Too many
requests)" errors.

To specify the outgoing rate use the `SetRate(second, minute int)` function. The
values given will correspond to the maximum number of requests allowed within that timeframe.

Specifying 0 will leave the rate unchanged.

The default rate is 5 requests/second and 300requests/minute.

The requests are not spread out evenly within this interval. They are sent out as soon
as possible.

This does **_not_** *guarantee* no 429 errors, it simply aims to reduce them. It guarantees
that the outgoing rate from one specific instance of the SDK will not exceed the specified
outgoing rate. However, not every clock is synchronized with our server and not every
application uses the same instance of the SDK.

# <a name="errors"></a>Errors
Errors returned from the SDK is **_not_** of type `error` but instead a pointer to a struct
corresponding to the JSON returned from the endpoint when an error occurs. See [official documentation](https://docs.abiosgaming.com/v2/reference#errors).

If no errors are returned type will be `<nil>`.

The ErrorStruct implements the `Stringer` interface.

Errors of type `error` will be forwarded to your application in the form of an ErrorStruct.
The `ErrorCode` will then be equal to 0 and the `Error` will specify that is is an application
error (rather than a client or server error).

# Endpoints
For each endpoint in the /v2/ API you can expect to find a corresponding method implemented
on the struct returned by `abios.New`. The names of these method directly corresponds to
the name of the endpoint, except with a capital letter. For /:id endpoints the corresponding
method will be named `EndpointById`, or if the type of id is specified `EndpointByTypeId`.

E.g:

| Endpoint               | SDK                 |
|------------------------|---------------------|
| /games                 | Games               |
| /series                | Series              |
| /series/:id            | SeriesById          |
| /incidents/:series\_id | IncidentsBySeriesId |

For a full list of endpoints see the [official documentation](https://docs.abiosgaming.com/v2/reference).

# <a name="parameters"></a>Parameters
All SDK methods requires a parameter of the type `type Parameters map[string][]string` which simply
maps keys to values.

There are three methods implemented the type `Parameters`, which we recommend you use.

| Method Signature           | Description                                             |
|----------------------------|---------------------------------------------------------|
|`Add(key, value string)`    |Append `val` to the list at `key`                        |
|`Del(key string)`           |Reset list associated with `key` to an empty list        |
|`Set(key, value string`     |Reset list at `key` to only contain `value`              |

If you don't want to provide any parameters simply provide an empty or `<nil>` map.

# Default Values and Types
Since Go is statically typed all possible fields will be available as at least their default value.

[encoding/json](https://golang.org/pkg/encoding/json/) is used to unmarshal the response from the API.

That means that if a key is not present in the JSON the corresponding field will be it's default value.
Thus an optional field that is not present in the returned JSON will be represented in the SDK
as the default value of that type.

# Concurrency
The struct returned from `abios.New` returns a `type AbiosSdk interface` and supports
concurrent use. You can easily pass this around to different go routines while still
being sure that the token will be refreshed before expiration and that the specified rate
will not be exceeded. See [Concurrent Use](#concurrent_example) for an example.

The SDK will try to perform as many requests as possible concurrently. It will create one
go-routine per request per second. For example, if the specified rate limit is 10 per
second then every second up to 10 go-routines will be created, one for each request. 

# Example Applications

## Usage

```Bash
git clone git@github.com:AbiosGaming/go-sdk-v2.git
cd go-sdk-v2/examples
# Edit the example and add your credentials to the abios.New() function call
go run <file>
```

## Querying /series Once Every Minute
```Go
package main

import (
    "fmt"
    "github.com/AbiosGaming/go-sdk-v2"
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
```

## Getting the Winrates of Teams that are Playing or are About to Play

Note that this might cause a lot of requests.

```Go
package main

import (
    "fmt"
    "github.com/AbiosGaming/go-sdk-v2"
)

func main() {
    a := abios.New("username", "password")

    thirtyMinutesAgo := time.Now().Add(-time.Minute*30).UTC().Format("2006-01-02T15:04:05Z")

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
```
## <a name="concurrent_example"></a>Concurrent Use
Query the /series endpoint every minute and the /games endpoint every thirty seconds.

```Go
package main

import (
    "fmt"
    "github.com/AbiosGaming/go-sdk-v2"
    "time"
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
```
## <a name="calendar_example"></a>Calendar Backend
A small webserver that queries the /series endpoint once every minutes and and serves each
request with the teams playing, the starttime, how many matches can be played (i.e best of)
as well as the name of the stage and tournament in JSON format.

Serves on port 8080 so visit localhost:8080 in your favourite browser or by using cUrl/wget
to see the response.

```Go
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
```
