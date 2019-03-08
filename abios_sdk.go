package abios

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	. "github.com/AbiosGaming/go-sdk-v2/v3/structs"
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
	Games(params Parameters) (GameStructPaginated, error)
	Series(params Parameters) (SeriesStructPaginated, error)
	SeriesById(id int64, params Parameters) (SeriesStruct, error)
	MatchesById(id int64, params Parameters) (MatchStruct, error)
	Tournaments(params Parameters) (TournamentStructPaginated, error)
	TournamentsById(id int64, params Parameters) (TournamentStruct, error)
	SubstagesById(id int64, params Parameters) (SubstageStruct, error)
	Teams(params Parameters) (TeamStructPaginated, error)
	TeamsById(id int64, params Parameters) (TeamStruct, error)
	Organisations(params Parameters) (OrganisationStructPaginated, error)
	OrganisationsById(id int64) (OrganisationStruct, error)
	Players(params Parameters) (PlayerStructPaginated, error)
	PlayersById(id int64, params Parameters) (PlayerStruct, error)
	RostersById(id int64, params Parameters) (RosterStruct, error)
	Search(query string, params Parameters) ([]SearchResultStruct, error)
	Incidents(params Parameters) (IncidentStructPaginated, error)
	IncidentsBySeriesId(id int64) (SeriesIncidentsStruct, error)
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

	statusCode, b, err := apiCall(req)
	if err != nil {
		return &result{body: nil, err: err}
	}
	dec := json.NewDecoder(bytes.NewBuffer(b))
	if 200 <= statusCode && statusCode < 300 {
		target := AccessTokenStruct{}
		err := dec.Decode(&target)
		if err != nil {
			return &result{body: nil, err: err}
		}
		a.oauth = target
		return nil
	}

	return &result{body: b, err: nil}
}

// Games queries the /games endpoint and returns a GameStructPaginated.
func (a *client) Games(params Parameters) (GameStructPaginated, error) {
	if params == nil {
		params = make(Parameters)
	} else {
		params = copyParams(params)
	}
	params.Set("access_token", a.oauth.AccessToken)
	result := <-a.handler.addRequest(a.base_url+games, params)
	if result.err != nil {
		return GameStructPaginated{}, result.err
	}

	target := GameStructPaginated{}
	err := json.Unmarshal(result.body, &target)
	if err != nil {
		return GameStructPaginated{}, err
	}
	return target, nil
}

// Series queries the /series endpoint and returns a SeriesStructPaginated.
func (a *client) Series(params Parameters) (SeriesStructPaginated, error) {
	if params == nil {
		params = make(Parameters)
	} else {
		params = copyParams(params)
	}

	params.Set("access_token", a.oauth.AccessToken)
	result := <-a.handler.addRequest(a.base_url+series, params)
	if result.err != nil {
		return SeriesStructPaginated{}, result.err
	}

	target := SeriesStructPaginated{}
	err := json.Unmarshal(result.body, &target)
	return target, err
}

// SeriesById queries the /series/:id endpoint and returns a SeriesStruct.
func (a *client) SeriesById(id int64, params Parameters) (SeriesStruct, error) {
	sId := strconv.FormatInt(id, 10)
	if params == nil {
		params = make(Parameters)
	} else {
		params = copyParams(params)
	}

	params.Set("access_token", a.oauth.AccessToken)
	result := <-a.handler.addRequest(a.base_url+seriesById+sId, params)
	if result.err != nil {
		return SeriesStruct{}, result.err
	}

	target := SeriesStruct{}
	err := json.Unmarshal(result.body, &target)
	return target, err
}

// MatchesById queries the /matches/:id endpoint and returns a MatchStruct.
func (a *client) MatchesById(id int64, params Parameters) (MatchStruct, error) {
	sId := strconv.FormatInt(id, 10)
	if params == nil {
		params = make(Parameters)
	} else {
		params = copyParams(params)
	}

	params.Set("access_token", a.oauth.AccessToken)
	result := <-a.handler.addRequest(a.base_url+matches+sId, params)
	if result.err != nil {
		return MatchStruct{}, result.err
	}

	target := MatchStruct{}
	err := json.Unmarshal(result.body, &target)
	return target, err
}

// Tournaments queries the /tournaments endpoint and returns a list of TournamentStructPaginated.
func (a *client) Tournaments(params Parameters) (TournamentStructPaginated, error) {
	if params == nil {
		params = make(Parameters)
	} else {
		params = copyParams(params)
	}

	params.Set("access_token", a.oauth.AccessToken)
	result := <-a.handler.addRequest(a.base_url+tournaments, params)
	if result.err != nil {
		return TournamentStructPaginated{}, result.err
	}

	target := TournamentStructPaginated{}
	err := json.Unmarshal(result.body, &target)
	return target, err
}

// TournamentsById queries the /tournaments/:id endpoint and return a TournamentStruct.
func (a *client) TournamentsById(id int64, params Parameters) (TournamentStruct, error) {
	sId := strconv.FormatInt(id, 10)
	if params == nil {
		params = make(Parameters)
	} else {
		params = copyParams(params)
	}

	params.Set("access_token", a.oauth.AccessToken)
	result := <-a.handler.addRequest(a.base_url+tournamentsById+sId, params)
	if result.err != nil {
		return TournamentStruct{}, result.err
	}

	target := TournamentStruct{}
	err := json.Unmarshal(result.body, &target)
	return target, err
}

// SubstagesById queries the /substages/:id endpoint and returns a SubstageStruct.
func (a *client) SubstagesById(id int64, params Parameters) (SubstageStruct, error) {
	sId := strconv.FormatInt(id, 10)
	if params == nil {
		params = make(Parameters)
	} else {
		params = copyParams(params)
	}

	params.Set("access_token", a.oauth.AccessToken)
	result := <-a.handler.addRequest(a.base_url+substages+sId, params)
	if result.err != nil {
		return SubstageStruct{}, result.err
	}

	target := SubstageStruct{}
	err := json.Unmarshal(result.body, &target)
	return target, err
}

// Teams queries the /teams endpoint and returns a TeamsStructPaginated.
func (a *client) Teams(params Parameters) (TeamStructPaginated, error) {
	if params == nil {
		params = make(Parameters)
	} else {
		params = copyParams(params)
	}

	params.Set("access_token", a.oauth.AccessToken)
	result := <-a.handler.addRequest(a.base_url+teams, params)
	if result.err != nil {
		return TeamStructPaginated{}, result.err
	}

	target := TeamStructPaginated{}
	err := json.Unmarshal(result.body, &target)
	return target, err
}

// TeamsById queries the /teams/:id endpoint and return a TeamStruct.
func (a *client) TeamsById(id int64, params Parameters) (TeamStruct, error) {
	sId := strconv.FormatInt(id, 10)
	if params == nil {
		params = make(Parameters)
	} else {
		params = copyParams(params)
	}

	params.Set("access_token", a.oauth.AccessToken)
	result := <-a.handler.addRequest(a.base_url+teamsById+sId, params)
	if result.err != nil {
		return TeamStruct{}, result.err
	}

	target := TeamStruct{}
	err := json.Unmarshal(result.body, &target)
	return target, err
}

// Organisations queries the /organisations endpoint
func (a *client) Organisations(params Parameters) (OrganisationStructPaginated, error) {
	if params == nil {
		params = make(Parameters)
	} else {
		params = copyParams(params)
	}

	params.Set("access_token", a.oauth.AccessToken)
	result := <-a.handler.addRequest(a.base_url+organisations, params)
	if result.err != nil {
		return OrganisationStructPaginated{}, result.err
	}

	target := OrganisationStructPaginated{}
	err := json.Unmarshal(result.body, &target)
	return target, err
}

// OrganisationsById queries the /organisations/:id endpoint
func (a *client) OrganisationsById(id int64) (OrganisationStruct, error) {
	sId := strconv.FormatInt(id, 10)
	params := make(Parameters)

	params.Set("access_token", a.oauth.AccessToken)
	result := <-a.handler.addRequest(a.base_url+organisationsById+sId, params)
	if result.err != nil {
		return OrganisationStruct{}, result.err
	}

	target := OrganisationStruct{}
	err := json.Unmarshal(result.body, &target)
	return target, err
}

// Players queries the /players endpoint and returns PlayerStructPaginated.
func (a *client) Players(params Parameters) (PlayerStructPaginated, error) {
	if params == nil {
		params = make(Parameters)
	} else {
		params = copyParams(params)
	}

	params.Set("access_token", a.oauth.AccessToken)
	result := <-a.handler.addRequest(a.base_url+players, params)
	if result.err != nil {
		return PlayerStructPaginated{}, result.err
	}

	target := PlayerStructPaginated{}
	err := json.Unmarshal(result.body, &target)
	return target, err
}

// PlayersById queries the /players/:id endpoint and returns a PlayerStruct.
func (a *client) PlayersById(id int64, params Parameters) (PlayerStruct, error) {
	sId := strconv.FormatInt(id, 10)
	if params == nil {
		params = make(Parameters)
	} else {
		params = copyParams(params)
	}

	params.Set("access_token", a.oauth.AccessToken)
	result := <-a.handler.addRequest(a.base_url+playersById+sId, params)
	if result.err != nil {
		return PlayerStruct{}, result.err
	}

	target := PlayerStruct{}
	err := json.Unmarshal(result.body, &target)
	return target, err
}

// RostersById queries the /rosters/:id endpoint and returns a RosterStruct.
func (a *client) RostersById(id int64, params Parameters) (RosterStruct, error) {
	sId := strconv.FormatInt(id, 10)
	if params == nil {
		params = make(Parameters)
	} else {
		params = copyParams(params)
	}

	params.Set("access_token", a.oauth.AccessToken)
	result := <-a.handler.addRequest(a.base_url+rosters+sId, params)
	if result.err != nil {
		return RosterStruct{}, result.err
	}

	target := RosterStruct{}
	err := json.Unmarshal(result.body, &target)
	return target, err
}

// Search queries the /search endpoint with the given query and returns a list of
// SearchResultStruct.
func (a *client) Search(query string, params Parameters) ([]SearchResultStruct, error) {
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

	target := []SearchResultStruct{}
	err := json.Unmarshal(result.body, &target)
	return target, err
}

// Incidents queries the /incidents endpoint and returns an IncidentStructPaginated.
func (a *client) Incidents(params Parameters) (IncidentStructPaginated, error) {
	if params == nil {
		params = make(Parameters)
	} else {
		params = copyParams(params)
	}

	params.Set("access_token", a.oauth.AccessToken)
	result := <-a.handler.addRequest(a.base_url+incidents, params)
	if result.err != nil {
		return IncidentStructPaginated{}, result.err
	}

	target := IncidentStructPaginated{}
	err := json.Unmarshal(result.body, &target)
	return target, err
}

// IncidentBySeriesId queries the /incidents/:series_id endpoint and returns a
// SeriesIncidentsStruct.
func (a *client) IncidentsBySeriesId(id int64) (SeriesIncidentsStruct, error) {
	sId := strconv.FormatInt(id, 10)
	params := make(Parameters)
	params.Set("access_token", a.oauth.AccessToken)
	result := <-a.handler.addRequest(a.base_url+incidentsBySeries+sId, params)
	if result.err != nil {
		return SeriesIncidentsStruct{}, result.err
	}

	target := SeriesIncidentsStruct{}
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
