package abios

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/AbiosGaming/go-sdk-v2/structs"
)

// performRequest creates the request, sends it and return the response's statuscode along
// with the response's body.
func performRequest(targetUrl string, params Parameters) (int, []byte, error) {
	u, err := url.Parse(targetUrl)
	if err != nil {
		return 0, nil, err
	}

	u.RawQuery = params.encode()

	httpReq := &http.Request{
		Method: "GET",
		URL:    u,
		Header: http.Header{
			"Content-Type": {"application/x-www-form-urlencoded"},
		},
	}

	return apiCall(httpReq)
}

// apiCall performs the actual http request and returns the resulting statuscode and body.
func apiCall(req *http.Request) (int, []byte, error) {
	client := &http.Client{Timeout: 20 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}

	// If it is an error try to unmarshal it into the ErrorStruct.
	// 410 still returns data in the expected format
	if resp.StatusCode != 410 && (resp.StatusCode < 200 || 300 <= resp.StatusCode) {
		target := structs.ErrorStruct{}
		err := json.Unmarshal(body, &target)
		if err != nil {
			return 0, nil, err
		}
		return resp.StatusCode, nil, target
	}

	return resp.StatusCode, body, nil
}
