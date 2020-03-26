package abios

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/AbiosGaming/go-sdk-v2/v3/structs"
)

// Constant variables that represents endpoints
const (
	baseUrl           = "https://api.abiosgaming.com/v2/"
	access_token      = "oauth/access_token"
	games             = "games"
	series            = "series"
	seriesById        = series + "/"
	matches           = "matches/"
	tournaments       = "tournaments"
	tournamentsById   = tournaments + "/"
	substages         = "substages/"
	teams             = "teams"
	teamsById         = teams + "/"
	organisations     = "organisations"
	organisationsById = organisations + "/"
	players           = "players"
	playersById       = players + "/"
	rosters           = "rosters/"
	search            = "search"
	incidents         = "incidents"
	incidentsBySeries = incidents + "/"
)

// AbiosSdk defines the interface of an implementation of a SDK targeting the Abios endpoints.
type AbiosSdk interface {
	SetRate(second, minute uint)
	Games(params Parameters) (structs.PaginatedGames, error)
	Series(params Parameters) (structs.PaginatedSeries, error)
	SeriesById(id int64, params Parameters) (structs.Series, error)
	MatchesById(id int64, params Parameters) (structs.Match, error)
	Tournaments(params Parameters) (structs.PaginatedTournaments, error)
	TournamentsById(id int64, params Parameters) (structs.Tournament, error)
	SubstagesById(id int64, params Parameters) (structs.Substage, error)
	Teams(params Parameters) (structs.PaginatedTeams, error)
	TeamsById(id int64, params Parameters) (structs.Team, error)
	Organisations(params Parameters) (structs.PaginatedOrganisations, error)
	OrganisationsById(id int64) (structs.Organisation, error)
	Players(params Parameters) (structs.PaginatedPlayers, error)
	PlayersById(id int64, params Parameters) (structs.Player, error)
	RostersById(id int64, params Parameters) (structs.Roster, error)
	Search(query string, params Parameters) ([]structs.SearchResult, error)
	Incidents(params Parameters) (structs.PaginatedIncidents, error)
	IncidentsBySeriesId(id int64) (structs.SeriesIncidents, error)
}

// client holds the oauth string returned from Authenticate as well as this sessions
// requestHandler.
type client struct {
	username string
	password string
	oauth    structs.AccessToken
	handler  *requestHandler
	base_url string
}

// authenticator makes sure the oauth token doesn't expire.
func (a *client) authenticator() {
	for {
		// Wait until token is about to expire
		expires := time.Duration(a.oauth.ExpiresIn) * time.Second
		time.Sleep(expires - time.Minute*9) // Sleep until at most 9 minutes left.

		err := a.authenticate() // try once
		if err == nil {
			continue // It succeded.
		}

		// If we get an error we retry every 30 seconds for 7 minutes before we override
		// the responses.
		retry := time.NewTicker(30 * time.Second)
		fail := time.NewTimer(7 * time.Minute)

		select {
		case <-retry.C:
			err = a.authenticate()
			if err == nil {
				continue
			}
		case <-fail.C:
			a.handler.override = responseOverride{override: true, data: *err}
			return
		}
	}
}

// New returns a new endpoint-wrapper for api version 2 with given credentials using the default url.
func New(username, password string) *client {
	return NewWithUrl(username, password, baseUrl)
}

// NewWithUrl returns a new endpoint-wrapper for api version 2 with given credentials and base_url.
// Will append "/" to base_url if missing.
func NewWithUrl(username, password, base_url string) *client {
	if !strings.HasSuffix(base_url, "/") {
		base_url += "/"
	}

	r := newRequestHandler()
	c := &client{username, password, structs.AccessToken{}, r, base_url}
	err := c.authenticate()
	if err != nil {
		c.handler.override = responseOverride{override: true, data: *err}
	}
	go c.authenticator() // Launch authenticator
	return c
}

// SetRate sets the outgoing rate to "second" requests per second and "minute" requests
// per minte. A value less than or equal to 0 means previous
// value is kept. Default values are (5, 300)
func (a *client) SetRate(second, minute uint) {
	a.handler.setRate(second, minute)
}

// authenticate queries the /oauth/access_token endpoint with the given credentials and
// stores the returned oauth token. Return nil if the request was successful.
func (a *client) authenticate() *result {
	var payload = []byte(`grant_type=client_credentials&client_id=` + a.username + `&client_secret=` + a.password)

	req, err := http.NewRequest("POST", a.base_url+access_token, bytes.NewBuffer(payload))
	if err != nil {
		return &result{body: nil, err: err}
	}
	req.Header = http.Header{"Content-Type": {"application/x-www-form-urlencoded"}}

	statusCode, b, err := apiCall(req)
	if err != nil {
		return &result{body: nil, err: err}
	}
	dec := json.NewDecoder(bytes.NewBuffer(b))
	if 200 <= statusCode && statusCode < 300 {
		target := structs.AccessToken{}
		err := dec.Decode(&target)
		if err != nil {
			return &result{body: nil, err: err}
		}
		a.oauth = target
		return nil
	}

	return &result{body: b, err: nil}
}

// Games queries the /games endpoint and returns a structs.PaginatedGames.
func (a *client) Games(params Parameters) (structs.PaginatedGames, error) {
	if params == nil {
		params = make(Parameters)
	} else {
		params = copyParams(params)
	}
	params.Set("access_token", a.oauth.AccessToken)
	result := <-a.handler.addRequest(a.base_url+games, params)
	if result.err != nil {
		return structs.PaginatedGames{}, result.err
	}

	target := structs.PaginatedGames{}
	err := json.Unmarshal(result.body, &target)
	if err != nil {
		return structs.PaginatedGames{}, err
	}
	return target, nil
}

// Series queries the /series endpoint and returns a structs.PaginatedSeries.
func (a *client) Series(params Parameters) (structs.PaginatedSeries, error) {
	if params == nil {
		params = make(Parameters)
	} else {
		params = copyParams(params)
	}

	params.Set("access_token", a.oauth.AccessToken)
	result := <-a.handler.addRequest(a.base_url+series, params)
	if result.err != nil {
		return structs.PaginatedSeries{}, result.err
	}

	target := structs.PaginatedSeries{}
	err := json.Unmarshal(result.body, &target)
	return target, err
}

// SeriesById queries the /series/:id endpoint and returns a structs.Series.
func (a *client) SeriesById(id int64, params Parameters) (structs.Series, error) {
	sId := strconv.FormatInt(id, 10)
	if params == nil {
		params = make(Parameters)
	} else {
		params = copyParams(params)
	}

	params.Set("access_token", a.oauth.AccessToken)
	result := <-a.handler.addRequest(a.base_url+seriesById+sId, params)
	if result.err != nil {
		return structs.Series{}, result.err
	}

	target := structs.Series{}
	err := json.Unmarshal(result.body, &target)
	return target, err
}

// MatchesById queries the /matches/:id endpoint and returns a structs.Match.
func (a *client) MatchesById(id int64, params Parameters) (structs.Match, error) {
	sId := strconv.FormatInt(id, 10)
	if params == nil {
		params = make(Parameters)
	} else {
		params = copyParams(params)
	}

	params.Set("access_token", a.oauth.AccessToken)
	result := <-a.handler.addRequest(a.base_url+matches+sId, params)
	if result.err != nil {
		return structs.Match{}, result.err
	}

	target := structs.Match{}
	err := json.Unmarshal(result.body, &target)
	return target, err
}

// Tournaments queries the /tournaments endpoint and returns a list of structs.PaginatedTournaments.
func (a *client) Tournaments(params Parameters) (structs.PaginatedTournaments, error) {
	if params == nil {
		params = make(Parameters)
	} else {
		params = copyParams(params)
	}

	params.Set("access_token", a.oauth.AccessToken)
	result := <-a.handler.addRequest(a.base_url+tournaments, params)
	if result.err != nil {
		return structs.PaginatedTournaments{}, result.err
	}

	target := structs.PaginatedTournaments{}
	err := json.Unmarshal(result.body, &target)
	return target, err
}

// TournamentsById queries the /tournaments/:id endpoint and return a structs.Tournament.
func (a *client) TournamentsById(id int64, params Parameters) (structs.Tournament, error) {
	sId := strconv.FormatInt(id, 10)
	if params == nil {
		params = make(Parameters)
	} else {
		params = copyParams(params)
	}

	params.Set("access_token", a.oauth.AccessToken)
	result := <-a.handler.addRequest(a.base_url+tournamentsById+sId, params)
	if result.err != nil {
		return structs.Tournament{}, result.err
	}

	target := structs.Tournament{}
	err := json.Unmarshal(result.body, &target)
	return target, err
}

// SubstagesById queries the /substages/:id endpoint and returns a structs.Substage.
func (a *client) SubstagesById(id int64, params Parameters) (structs.Substage, error) {
	sId := strconv.FormatInt(id, 10)
	if params == nil {
		params = make(Parameters)
	} else {
		params = copyParams(params)
	}

	params.Set("access_token", a.oauth.AccessToken)
	result := <-a.handler.addRequest(a.base_url+substages+sId, params)
	if result.err != nil {
		return structs.Substage{}, result.err
	}

	target := structs.Substage{}
	err := json.Unmarshal(result.body, &target)
	return target, err
}

// Teams queries the /teams endpoint and returns a structs.PaginatedTeams.
func (a *client) Teams(params Parameters) (structs.PaginatedTeams, error) {
	if params == nil {
		params = make(Parameters)
	} else {
		params = copyParams(params)
	}

	params.Set("access_token", a.oauth.AccessToken)
	result := <-a.handler.addRequest(a.base_url+teams, params)
	if result.err != nil {
		return structs.PaginatedTeams{}, result.err
	}

	target := structs.PaginatedTeams{}
	err := json.Unmarshal(result.body, &target)
	return target, err
}

// TeamsById queries the /teams/:id endpoint and return a structs.Team.
func (a *client) TeamsById(id int64, params Parameters) (structs.Team, error) {
	sId := strconv.FormatInt(id, 10)
	if params == nil {
		params = make(Parameters)
	} else {
		params = copyParams(params)
	}

	params.Set("access_token", a.oauth.AccessToken)
	result := <-a.handler.addRequest(a.base_url+teamsById+sId, params)
	if result.err != nil {
		return structs.Team{}, result.err
	}

	target := structs.Team{}
	err := json.Unmarshal(result.body, &target)
	return target, err
}

// Organisations queries the /organisations endpoint
func (a *client) Organisations(params Parameters) (structs.PaginatedOrganisations, error) {
	if params == nil {
		params = make(Parameters)
	} else {
		params = copyParams(params)
	}

	params.Set("access_token", a.oauth.AccessToken)
	result := <-a.handler.addRequest(a.base_url+organisations, params)
	if result.err != nil {
		return structs.PaginatedOrganisations{}, result.err
	}

	target := structs.PaginatedOrganisations{}
	err := json.Unmarshal(result.body, &target)
	return target, err
}

// OrganisationsById queries the /organisations/:id endpoint
func (a *client) OrganisationsById(id int64) (structs.Organisation, error) {
	sId := strconv.FormatInt(id, 10)
	params := make(Parameters)

	params.Set("access_token", a.oauth.AccessToken)
	result := <-a.handler.addRequest(a.base_url+organisationsById+sId, params)
	if result.err != nil {
		return structs.Organisation{}, result.err
	}

	target := structs.Organisation{}
	err := json.Unmarshal(result.body, &target)
	return target, err
}

// Players queries the /players endpoint and returns structs.PaginatedPlayers.
func (a *client) Players(params Parameters) (structs.PaginatedPlayers, error) {
	if params == nil {
		params = make(Parameters)
	} else {
		params = copyParams(params)
	}

	params.Set("access_token", a.oauth.AccessToken)
	result := <-a.handler.addRequest(a.base_url+players, params)
	if result.err != nil {
		return structs.PaginatedPlayers{}, result.err
	}

	target := structs.PaginatedPlayers{}
	err := json.Unmarshal(result.body, &target)
	return target, err
}

// PlayersById queries the /players/:id endpoint and returns a structs.Player.
func (a *client) PlayersById(id int64, params Parameters) (structs.Player, error) {
	sId := strconv.FormatInt(id, 10)
	if params == nil {
		params = make(Parameters)
	} else {
		params = copyParams(params)
	}

	params.Set("access_token", a.oauth.AccessToken)
	result := <-a.handler.addRequest(a.base_url+playersById+sId, params)
	if result.err != nil {
		return structs.Player{}, result.err
	}

	target := structs.Player{}
	err := json.Unmarshal(result.body, &target)
	return target, err
}

// RostersById queries the /rosters/:id endpoint and returns a structs.Roster.
func (a *client) RostersById(id int64, params Parameters) (structs.Roster, error) {
	sId := strconv.FormatInt(id, 10)
	if params == nil {
		params = make(Parameters)
	} else {
		params = copyParams(params)
	}

	params.Set("access_token", a.oauth.AccessToken)
	result := <-a.handler.addRequest(a.base_url+rosters+sId, params)
	if result.err != nil {
		return structs.Roster{}, result.err
	}

	target := structs.Roster{}
	err := json.Unmarshal(result.body, &target)
	return target, err
}

// Search queries the /search endpoint with the given query and returns a list of
// structs.SearchResult.
func (a *client) Search(query string, params Parameters) ([]structs.SearchResult, error) {
	if params == nil {
		params = make(Parameters)
	} else {
		params = copyParams(params)
	}

	params.Set("access_token", a.oauth.AccessToken)
	params.Add("q", query)
	result := <-a.handler.addRequest(a.base_url+search, params)
	if result.err != nil {
		return nil, result.err
	}

	target := []structs.SearchResult{}
	err := json.Unmarshal(result.body, &target)
	return target, err
}

// Incidents queries the /incidents endpoint and returns an structs.PaginatedIncidents.
func (a *client) Incidents(params Parameters) (structs.PaginatedIncidents, error) {
	if params == nil {
		params = make(Parameters)
	} else {
		params = copyParams(params)
	}

	params.Set("access_token", a.oauth.AccessToken)
	result := <-a.handler.addRequest(a.base_url+incidents, params)
	if result.err != nil {
		return structs.PaginatedIncidents{}, result.err
	}

	target := structs.PaginatedIncidents{}
	err := json.Unmarshal(result.body, &target)
	return target, err
}

// IncidentBySeriesId queries the /incidents/:series_id endpoint and returns a
// structs.SeriesIncidents.
func (a *client) IncidentsBySeriesId(id int64) (structs.SeriesIncidents, error) {
	sId := strconv.FormatInt(id, 10)
	params := make(Parameters)
	params.Set("access_token", a.oauth.AccessToken)
	result := <-a.handler.addRequest(a.base_url+incidentsBySeries+sId, params)
	if result.err != nil {
		return structs.SeriesIncidents{}, result.err
	}

	target := structs.SeriesIncidents{}
	err := json.Unmarshal(result.body, &target)
	return target, err
}

// copyParams copies the parameters to a new map so different routines don't share a
// non-thread-safe map.
func copyParams(from Parameters) (to Parameters) {
	to = make(Parameters)
	for k, v := range from {
		to[k] = v
	}
	return
}
