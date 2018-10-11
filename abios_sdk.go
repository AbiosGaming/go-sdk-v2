package abios

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	. "github.com/AbiosGaming/go-sdk-v2/structs"
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
	Games(params Parameters) (GameStructPaginated, *ErrorStruct)
	Series(params Parameters) (SeriesStructPaginated, *ErrorStruct)
	SeriesById(id int64, params Parameters) (SeriesStruct, *ErrorStruct)
	MatchesById(id int64, params Parameters) (MatchStruct, *ErrorStruct)
	Tournaments(params Parameters) (TournamentStructPaginated, *ErrorStruct)
	TournamentsById(id int64, params Parameters) (TournamentStruct, *ErrorStruct)
	SubstagesById(id int64, params Parameters) (SubstageStruct, *ErrorStruct)
	Teams(params Parameters) (TeamStructPaginated, *ErrorStruct)
	TeamsById(id int64, params Parameters) (TeamStruct, *ErrorStruct)
	Organisations(params Parameters) (OrganisationStructPaginated, *ErrorStruct)
	OrganisationsById(id int64) (OrganisationStruct, *ErrorStruct)
	Players(params Parameters) (PlayerStructPaginated, *ErrorStruct)
	PlayersById(id int64, params Parameters) (PlayerStruct, *ErrorStruct)
	RostersById(id int64, params Parameters) (RosterStruct, *ErrorStruct)
	Search(query string, params Parameters) ([]SearchResultStruct, *ErrorStruct)
	Incidents(params Parameters) (IncidentStructPaginated, *ErrorStruct)
	IncidentsBySeriesId(id int64) (SeriesIncidentsStruct, *ErrorStruct)
}

// client holds the oauth string returned from Authenticate as well as this sessions
// requestHandler.
type client struct {
	username string
	password string
	oauth    AccessTokenStruct
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

		// If we get an error we retry every 30 seconds for 5 minutes before we override
		// the responses.
		retry := time.NewTicker(30 * time.Second)
		fail := time.NewTimer(5 * time.Minute)

		select {
		case <-retry.C:
			err = a.authenticate()
			if err == nil {
				a.handler.override = responseOverride{override: false, data: result{}}
				return
			}
		case <-fail.C:
			a.handler.override = responseOverride{override: true, data: *err}
			return
		}
	}
}

// NewAbios returns a new endpoint-wrapper for api version 2 with given credentials.
func New(username, password string) *client {
	return NewWithUrl(username, password, baseUrl)
}

// NewAbios returns a new endpoint-wrapper for api version 2 with given credentials and baseUrl.
func NewWithUrl(username, password, base_url string) *client {
	r := newRequestHandler()
	c := &client{username, password, AccessTokenStruct{}, r, base_url}
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

	req, _ := http.NewRequest("POST", a.base_url+access_token, bytes.NewBuffer(payload))
	req.Header = http.Header{"Content-Type": {"application/x-www-form-urlencoded"}}

	statusCode, b := apiCall(req)
	dec := json.NewDecoder(bytes.NewBuffer(b))
	if 200 <= statusCode && statusCode < 300 {
		target := AccessTokenStruct{}
		dec.Decode(&target)
		a.oauth = target
		return nil
	}

	return &result{statuscode: statusCode, body: b}
}

// Games queries the /games endpoint and returns a GameStructPaginated.
func (a *client) Games(params Parameters) (GameStructPaginated, *ErrorStruct) {
	if params == nil {
		params = make(Parameters)
	} else {
		params = copyParams(params)
	}
	params.Set("access_token", a.oauth.AccessToken)
	result := <-a.handler.addRequest(a.base_url+games, params)

	dec := json.NewDecoder(bytes.NewBuffer(result.body))
	if 200 <= result.statuscode && result.statuscode < 300 {
		target := GameStructPaginated{}
		dec.Decode(&target)
		return target, nil
	}

	target := ErrorStruct{}
	dec.Decode(&target)
	return GameStructPaginated{}, &target
}

// Series queries the /series endpoint and returns a SeriesStructPaginated.
func (a *client) Series(params Parameters) (SeriesStructPaginated, *ErrorStruct) {
	if params == nil {
		params = make(Parameters)
	} else {
		params = copyParams(params)
	}

	params.Set("access_token", a.oauth.AccessToken)
	result := <-a.handler.addRequest(a.base_url+series, params)

	dec := json.NewDecoder(bytes.NewBuffer(result.body))
	if 200 <= result.statuscode && result.statuscode < 300 {
		target := SeriesStructPaginated{}
		dec.Decode(&target)
		return target, nil
	}

	target := ErrorStruct{}
	dec.Decode(&target)
	return SeriesStructPaginated{}, &target
}

// SeriesById queries the /series/:id endpoint and returns a SeriesStruct.
func (a *client) SeriesById(id int64, params Parameters) (SeriesStruct, *ErrorStruct) {
	sId := strconv.FormatInt(id, 10)
	if params == nil {
		params = make(Parameters)
	} else {
		params = copyParams(params)
	}

	params.Set("access_token", a.oauth.AccessToken)
	result := <-a.handler.addRequest(a.base_url+seriesById+sId, params)

	dec := json.NewDecoder(bytes.NewBuffer(result.body))
	if 200 <= result.statuscode && result.statuscode < 300 {
		target := SeriesStruct{}
		dec.Decode(&target)
		return target, nil
	}

	target := ErrorStruct{}
	dec.Decode(&target)
	return SeriesStruct{}, &target
}

// MatchesById queries the /matches/:id endpoint and returns a MatchStruct.
func (a *client) MatchesById(id int64, params Parameters) (MatchStruct, *ErrorStruct) {
	sId := strconv.FormatInt(id, 10)
	if params == nil {
		params = make(Parameters)
	} else {
		params = copyParams(params)
	}

	params.Set("access_token", a.oauth.AccessToken)
	result := <-a.handler.addRequest(a.base_url+matches+sId, params)

	dec := json.NewDecoder(bytes.NewBuffer(result.body))
	if 200 <= result.statuscode && result.statuscode < 300 {
		target := MatchStruct{}
		dec.Decode(&target)
		return target, nil
	}

	target := ErrorStruct{}
	dec.Decode(&target)
	return MatchStruct{}, &target
}

// Tournaments queries the /tournaments endpoint and returns a list of TournamentStructPaginated.
func (a *client) Tournaments(params Parameters) (TournamentStructPaginated, *ErrorStruct) {
	if params == nil {
		params = make(Parameters)
	} else {
		params = copyParams(params)
	}

	params.Set("access_token", a.oauth.AccessToken)
	result := <-a.handler.addRequest(a.base_url+tournaments, params)

	dec := json.NewDecoder(bytes.NewBuffer(result.body))
	if 200 <= result.statuscode && result.statuscode < 300 {
		target := TournamentStructPaginated{}
		dec.Decode(&target)
		return target, nil
	}

	target := ErrorStruct{}
	dec.Decode(&target)
	return TournamentStructPaginated{}, &target
}

// TournamentsById queries the /tournaments/:id endpoint and return a TournamentStruct.
func (a *client) TournamentsById(id int64, params Parameters) (TournamentStruct, *ErrorStruct) {
	sId := strconv.FormatInt(id, 10)
	if params == nil {
		params = make(Parameters)
	} else {
		params = copyParams(params)
	}

	params.Set("access_token", a.oauth.AccessToken)
	result := <-a.handler.addRequest(a.base_url+tournamentsById+sId, params)

	dec := json.NewDecoder(bytes.NewBuffer(result.body))
	if 200 <= result.statuscode && result.statuscode < 300 {
		target := TournamentStruct{}
		dec.Decode(&target)
		return target, nil
	}

	target := ErrorStruct{}
	dec.Decode(&target)
	return TournamentStruct{}, &target
}

// SubstagesById queries the /substages/:id endpoint and returns a SubstageStruct.
func (a *client) SubstagesById(id int64, params Parameters) (SubstageStruct, *ErrorStruct) {
	sId := strconv.FormatInt(id, 10)
	if params == nil {
		params = make(Parameters)
	} else {
		params = copyParams(params)
	}

	params.Set("access_token", a.oauth.AccessToken)
	result := <-a.handler.addRequest(a.base_url+substages+sId, params)

	dec := json.NewDecoder(bytes.NewBuffer(result.body))
	if 200 <= result.statuscode && result.statuscode < 300 {
		target := SubstageStruct{}
		dec.Decode(&target)
		return target, nil
	}

	target := ErrorStruct{}
	dec.Decode(&target)
	return SubstageStruct{}, &target
}

// Teams queries the /teams endpoint and returns a TeamsStructPaginated.
func (a *client) Teams(params Parameters) (TeamStructPaginated, *ErrorStruct) {
	if params == nil {
		params = make(Parameters)
	} else {
		params = copyParams(params)
	}

	params.Set("access_token", a.oauth.AccessToken)
	result := <-a.handler.addRequest(a.base_url+teams, params)

	dec := json.NewDecoder(bytes.NewBuffer(result.body))
	if 200 <= result.statuscode && result.statuscode < 300 {
		target := TeamStructPaginated{}
		dec.Decode(&target)
		return target, nil
	}

	target := ErrorStruct{}
	dec.Decode(&target)
	return TeamStructPaginated{}, &target
}

// TeamsById queries the /teams/:id endpoint and return a TeamStruct.
func (a *client) TeamsById(id int64, params Parameters) (TeamStruct, *ErrorStruct) {
	sId := strconv.FormatInt(id, 10)
	if params == nil {
		params = make(Parameters)
	} else {
		params = copyParams(params)
	}

	params.Set("access_token", a.oauth.AccessToken)
	result := <-a.handler.addRequest(a.base_url+teamsById+sId, params)

	dec := json.NewDecoder(bytes.NewBuffer(result.body))
	if 200 <= result.statuscode && result.statuscode < 300 {
		target := TeamStruct{}
		dec.Decode(&target)
		return target, nil
	}

	target := ErrorStruct{}
	dec.Decode(&target)
	return TeamStruct{}, &target
}

// Organisations queries the /organisations endpoint
func (a *client) Organisations(params Parameters) (OrganisationStructPaginated, *ErrorStruct) {
	if params == nil {
		params = make(Parameters)
	} else {
		params = copyParams(params)
	}

	params.Set("access_token", a.oauth.AccessToken)
	result := <-a.handler.addRequest(a.base_url+organisations, params)

	dec := json.NewDecoder(bytes.NewBuffer(result.body))
	if 200 <= result.statuscode && result.statuscode < 300 {
		target := OrganisationStructPaginated{}
		dec.Decode(&target)
		return target, nil
	}

	target := ErrorStruct{}
	dec.Decode(&target)
	return OrganisationStructPaginated{}, &target
}

// OrganisationsById queries the /organisations/:id endpoint
func (a *client) OrganisationsById(id int64) (OrganisationStruct, *ErrorStruct) {
	sId := strconv.FormatInt(id, 10)
	params := make(Parameters)

	params.Set("access_token", a.oauth.AccessToken)
	result := <-a.handler.addRequest(a.base_url+organisationsById+sId, params)

	dec := json.NewDecoder(bytes.NewBuffer(result.body))
	if 200 <= result.statuscode && result.statuscode < 300 {
		target := OrganisationStruct{}
		dec.Decode(&target)
		return target, nil
	}

	target := ErrorStruct{}
	dec.Decode(&target)
	return OrganisationStruct{}, &target
}

// Players queries the /players endpoint and returns PlayerStructPaginated.
func (a *client) Players(params Parameters) (PlayerStructPaginated, *ErrorStruct) {
	if params == nil {
		params = make(Parameters)
	} else {
		params = copyParams(params)
	}

	params.Set("access_token", a.oauth.AccessToken)
	result := <-a.handler.addRequest(a.base_url+players, params)

	dec := json.NewDecoder(bytes.NewBuffer(result.body))
	if 200 <= result.statuscode && result.statuscode < 300 {
		target := PlayerStructPaginated{}
		dec.Decode(&target)
		return target, nil
	}

	target := ErrorStruct{}
	dec.Decode(&target)
	return PlayerStructPaginated{}, &target
}

// PlayersById queries the /players/:id endpoint and returns a PlayerStruct.
func (a *client) PlayersById(id int64, params Parameters) (PlayerStruct, *ErrorStruct) {
	sId := strconv.FormatInt(id, 10)
	if params == nil {
		params = make(Parameters)
	} else {
		params = copyParams(params)
	}

	params.Set("access_token", a.oauth.AccessToken)
	result := <-a.handler.addRequest(a.base_url+playersById+sId, params)

	dec := json.NewDecoder(bytes.NewBuffer(result.body))
	if 200 <= result.statuscode && result.statuscode < 300 {
		target := PlayerStruct{}
		dec.Decode(&target)
		return target, nil
	}

	target := ErrorStruct{}
	dec.Decode(&target)
	return PlayerStruct{}, &target
}

// RostersById queries the /rosters/:id endpoint and returns a RosterStruct.
func (a *client) RostersById(id int64, params Parameters) (RosterStruct, *ErrorStruct) {
	sId := strconv.FormatInt(id, 10)
	if params == nil {
		params = make(Parameters)
	} else {
		params = copyParams(params)
	}

	params.Set("access_token", a.oauth.AccessToken)
	result := <-a.handler.addRequest(a.base_url+rosters+sId, params)

	dec := json.NewDecoder(bytes.NewBuffer(result.body))
	if 200 <= result.statuscode && result.statuscode < 300 {
		target := RosterStruct{}
		dec.Decode(&target)
		return target, nil
	}

	target := ErrorStruct{}
	dec.Decode(&target)
	return RosterStruct{}, &target
}

// Search queries the /search endpoint with the given query and returns a list of
// SearchResultStruct.
func (a *client) Search(query string, params Parameters) ([]SearchResultStruct, *ErrorStruct) {
	if params == nil {
		params = make(Parameters)
	} else {
		params = copyParams(params)
	}

	params.Set("access_token", a.oauth.AccessToken)
	params.Add("q", query)
	result := <-a.handler.addRequest(a.base_url+search, params)

	dec := json.NewDecoder(bytes.NewBuffer(result.body))
	if 200 <= result.statuscode && result.statuscode < 300 {
		target := []SearchResultStruct{}
		dec.Decode(&target)
		return target, nil
	}

	target := ErrorStruct{}
	dec.Decode(&target)
	return []SearchResultStruct{}, &target
}

// Incidents queries the /incidents endpoint and returns an IncidentStructPaginated.
func (a *client) Incidents(params Parameters) (IncidentStructPaginated, *ErrorStruct) {
	if params == nil {
		params = make(Parameters)
	} else {
		params = copyParams(params)
	}

	params.Set("access_token", a.oauth.AccessToken)
	result := <-a.handler.addRequest(a.base_url+incidents, params)

	dec := json.NewDecoder(bytes.NewBuffer(result.body))
	if 200 <= result.statuscode && result.statuscode < 300 {
		target := IncidentStructPaginated{}
		dec.Decode(&target)
		return target, nil
	}

	target := ErrorStruct{}
	dec.Decode(&target)
	return IncidentStructPaginated{}, &target
}

// IncidentBySeriesId queries the /incidents/:series_id endpoint and returns a
// SeriesIncidentsStruct.
func (a *client) IncidentsBySeriesId(id int64) (SeriesIncidentsStruct, *ErrorStruct) {
	sId := strconv.FormatInt(id, 10)
	params := make(Parameters)
	params.Set("access_token", a.oauth.AccessToken)
	result := <-a.handler.addRequest(a.base_url+incidentsBySeries+sId, params)

	dec := json.NewDecoder(bytes.NewBuffer(result.body))
	if 200 <= result.statuscode && result.statuscode < 300 {
		target := SeriesIncidentsStruct{}
		dec.Decode(&target)
		return target, nil
	}

	target := ErrorStruct{}
	dec.Decode(&target)
	return SeriesIncidentsStruct{}, &target
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
