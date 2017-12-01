package abios

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// performRequest creates the request, sends it and return the response's statuscode along
// with the response's body.
func performRequest(targetUrl string, params Parameters) (int, []byte) {
	u, err := url.Parse(targetUrl)
	if err != nil {
		// Return something that looks similar to Abios API errors.
		errData := []byte(`
			{
				"error": "application error when parsing URL",
				"error_code": 0,
				"error_description": "` + err.Error() + `"
			}
		`)
		return 0, errData
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
func apiCall(req *http.Request) (int, []byte) {
	client := &http.Client{Timeout: 20 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		// Return something that looks similar to Abios API errors.
		errData := []byte(`
			{
				"error": "application error when attempting to perform HTTP request",
				"error_code": 0,
				"error_description": "` + err.Error() + `"
			}
		`)
		return 0, errData
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return resp.StatusCode, body
}
