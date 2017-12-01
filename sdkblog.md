# Introducing the first Abios SDK

While we are always busy with adding new features to our API and increasing out data coverage
we thought it was time to make it easier for you, our users, to build your applications.
Thus, we are happy to reveal our Golang SDK.

On our own backend side we are utilizing Golang to great lengths, and thus we decided to
start with an SDK for golang. However, the plan is to build more SDKs with similar features.
Do you want one in your favourite language? Make your voice heard [here](https://goo.gl/forms/hDLs0MOjksSbKKd73).

# What does it do?

The SDK allows you to easily make requests to the Abios API. All you need to provide is
your credentials and it will handle the rest. No more will you need to worry about creating
a proper URL with all the wanted parameters. No more will you have to worry about re-creating
that pesky token every hour. No more will you have to worry about not sending more requests
than allowed!

# That doesn't sound too bad. How do I use it?

Fortunately, using the SDK is very simple! All that is required is, of course, Golang. You
can find more information about how to setup Golang [here](https://golang.org/doc/install).

When you have Golang all set up you can simply run the command `go get -u github.com/AbiosGaming/go-sdk-v2`
and the Go toolchain will download the latest version and install it for you. Using it is
about as simple!

# A quick introduction

As a software developer I love examples. That's why I will provide a small example about
how one can use the SDK. Don't be discouraged though, just because there aren't many lines
of codes doesn't mean that it isn't powerful.

So what are we going to build? We want to build a backend for our simple calendar app. So
we want to be able to handle http requests and responds with a JSON payload that represents
what we want to display on our calendar.

So what do we want to display? Since we are building a calendar we want the start time of
every series, that's a given. We also want the teams playing and how many matches can be
played in the series, otherwise the start time doesn't tell us much. As a favour to our
users we also want to display the title of the tournament as well as short text what the
series is about. Putting it all together we want to respond to each request with something
like this:

```json
{
    "start_time": "The date and time when the series will be played",
    "home": "The name of the first team",
    "away": "The name of the second team",
    "best_of": "How many matches can be played?",
    "title": "Title of the series",
    "tournament_title": "Name of the tournament",
}
```

Simple enough, right? That should at least get us started with our calendar.

Now we know what we want to display to our client-side applications. But how to we get it
there? It is very easy to build a webserver in Golang so let's start with our client-facing
code. First we define the struct that corresponds to our JSON response. Adding the struct
tags will allow us to use the official `encoding/json` package with ease.

```Go
// calendarEntry holds all data neccesary to display a series in a calendar
type calendarEntry struct {
    StartTime       string `json:"start_time"`
    Roster1         string `json:"home"`
    Roster2         string `json:"away"`
    BestOf          int64  `json:"best_of"`
    Title           string `json:"title"`
    TournamentTitle string `json:"tournament_title"`
}
```

Since we have a limited request rate at the Abios API we want to store a list of these
entries in global memory so that we don't need to fetch new data every time someone asks
for our calendar.

```Go
// currentCalendar holds what we repond to our users
var currentCalendar []calendarEntry
```

Now let's get the actual webserver going! We need a main function that listens on a port
and responds with our `currentCalendar` to all requests:

```Go
package main

import (
    "encoding/json"
    "log"
    "net/http"
)

// calendarEntry holds all data neccesary to display a series in a calendar
type calendarEntry struct {
    StartTime       string `json:"start_time"`
    Roster1         string `json:"home"`
    Roster2         string `json:"away"`
    BestOf          int64  `json:"best_of"`
    Title           string `json:"title"`
    TournamentTitle string `json:"tournament_title"`
}

// currentCalendar holds what we repond to our users
var currentCalendar []calendarEntry

// handler marshals the currentCalendar as JSON and writes to the requester
func handler(w http.ResponseWriter, r *http.Request) {
    payload, err := json.MarshalIndent(currentCalendar, "", "\t")
    if err != nil {
        log.Println("Unable to marshal response")
        return // Since we couldn't marshal proper JSON we don't want to write anything
    }
    w.Write(payload)
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
```

Now we are getting somewhere. We have a webserver that responds to all requests with
our `currentCalendar` as a JSON-formatted string. Want to try it out? Copy paste the code
into a .go file and run with `go run`. With the application running you can simply visit
`localhost:8080` and you will get the calendar in it's current state.

If you actually did try it out you probably noticed that the response was merely "null".
That is unsurprising. We don't have any data in our `currentCalendar`!

Now we get to the juicy part. Let's intergrate the Abios API using the SDK! First we need
to import and initialize it. We want our instance of the SDK to persist through the entire
program so we will initalize it in the `main` function:

```Go
import (
    "encoding/json"
    "log"
    "net/http"

    "github.com/AbiosGaming/go-sdk-v2"
)

/*
    All that other code here
*/

func main() {
    // Add our credentials to the SDK
    a := abios.New("username", "password")

    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
```

Of course you have to substitute `"username"` and `"password"` with actual credentials.

Let's pick up the pace a little. We will define a function called `getSeries` that does
a request to the Abios API and picks out the data corresponding to the `calendarEntry`
struct. We will also launch a go-routine from `main` that uses `getSeries` to update
`currentCalendar` once every minute. To limit our rate to once every minute we will
use the SDK feature `SetRate(second, minute int)`. This will make sure our outgoing
requests never exceed this rate.

```Go
import (
    "encoding/json"
    "log"
    "net/http"

    "github.com/AbiosGaming/go-sdk-v2"
)

/*
    All that other code here
*/

func main() {
    // Add our credentials to the SDK
    a := abios.New("username", "password")

    // Set the outgoing rate to once very minute so we don't query too much
    a.SetRate(0, 1)


    // We want to update what we respond with once every minute. The SDK makes sure that
    // the rate we specified above isn't exceeded.
    go func(a abios.AbiosSdk) {
       for {
           currentCalendar = getSeries(a)
       }
    }(a)

    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
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

That is all the code we need for our application. This will update our internal `currentCalendar`
with data from the Abios API once every minute and responds to all client requests with that data.
Don't believe me? Try it out. Download this code and add it to the file you have from before,
run it with `go run` and visit `localhost:8080`.

All of the heavy lifting, and basically all of the code involving the SDK happends in this
`getSeries` function, so I will talk a little about it.

Since we want to display the tournament title we have to add the parameter `with[]=tournament`
to our request. Doing so is simple. We create a new instace of `abios.Parameters` and use
the `Add(key, value string)` function to add our parameter. We then pass this instance
to the API request, conveniently called `Series`. What we then have in our `series` variable
is a struct corresponding to the JSON returned. Neat!

We know we get a paginated result so we loop over the `Data` list. For each entry here
we have what we can expect from the `Series` resource in the Abios API. We do our necessary
checks and append a calendarEntry for this data to the list we will return. Easy-peasy,
right?

Behind the scenes of the SDK it will handle the authentication (and re-authentication). If
it fails we will get the error on the subsequent `Series` call, so we don't need any explicit
code to handle such a case.

Want to know more? The official SDK documentation is located in the file called README.md
in the Github repo found [here](https://github.com/AbiosGaming/go-sdk-v2). The repo
also contains the entire source code of this example and even a few more! If there is
something missing or you discover a bug somewhere just create an issue there and we will
get right to fixing it!

Want something similar in your favourite language? [Make your voice heard.](https://goo.gl/forms/hDLs0MOjksSbKKd73)
