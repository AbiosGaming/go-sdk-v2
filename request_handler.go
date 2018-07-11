package abios

import (
	"net/url"
	"sync/atomic"
	"time"
)

// Default values for the outgoing rate and size of request buffer.
const (
	default_requests_per_second int32 = 5
	default_requests_per_minute int32 = 300

	// Buffer one minutes worth of requests (this can not be changed at runtime)
	default_request_buffer_size = default_requests_per_minute
)

// Parameters maps a key (string) to a list of values ([]string).
type Parameters map[string][]string

// Add appends a value to the list associated with the key.
func (p Parameters) Add(key, value string) {
	p[key] = append(p[key], value)
}

// Del removes a key from the Parameters.
func (p Parameters) Del(key string) {
	p[key] = []string{}
}

// Set uses Del and Add to reset to list to length 1.
func (p Parameters) Set(key, value string) {
	p.Del(key)
	p.Add(key, value)
}

// encode formats the string according to url.Values.Encode.
func (p Parameters) encode() string {
	v := url.Values(p)
	return v.Encode()
}

// request is a logical container that groups which endpoint (as a complete url) to
// target with what parameters as well as a channel on which the result will be available
type request struct {
	url    string
	params Parameters
	ch     chan result
}

// result hold the returned data of an API request.
type result struct {
	statuscode int
	body       []byte
}

// requestHandler buffers requests and sends them out at a user-specified rate.
type requestHandler struct {
	requests_per_second int32            // How many requests can be performed per second.
	requests_per_minute int32            // How many requests can be performed per minute.
	queue               chan *request    // The queue of requests.
	override            responseOverride // Do we need to override the expected responses?
}

// responseOverride is a struct containing the logic of overriding responses.
// Used by e.g authenticator to indicate that something has gone wrong.
type responseOverride struct {
	override bool   // Should we override the reponse?
	data     result // The data we should return instead.
}

// addRequest creates and adds a Request to the requestHandler queue. It returns
// the channel on which the result will eventually be available.
func (r *requestHandler) addRequest(url string, params Parameters) chan result {
	returnCh := make(chan result)
	req := request{url, params, returnCh}
	r.queue <- &req
	return returnCh
}

// newRequestHandler creates a new requestHandler and starts the dispatcher
// goroutine.
func newRequestHandler() *requestHandler {
	h := &requestHandler{
		default_requests_per_second,
		default_requests_per_minute,
		make(chan *request, default_request_buffer_size),
		responseOverride{
			override: false,
			data:     result{},
		},
	}

	go h.dispatcher()
	return h
}

// SetRate sets the outgoing rate according to the give parameters. 0 or less means do nothing.
func (r *requestHandler) setRate(second, minute int32) {
	if 0 < second {
		r.requests_per_second = second
	}

	if 0 < minute {
		r.requests_per_minute = minute
	}

	// Make sure they are consistent
	if r.requests_per_second > r.requests_per_minute {
		r.requests_per_second = r.requests_per_minute
	}

}

// dispatcher does requestHandler.Rate api-calls every requestHandler.ResetInterval
func (r *requestHandler) dispatcher() {
	var requests_this_minute int32

	ticker_second := time.NewTicker(time.Second)
	ticker_minute := time.NewTicker(time.Minute)

	for {
		select {
		//case <-ticker_day.C: // Example of how to add more time-frames
		//	// Allow for more requests!
		//	requests_today = 0
		case <-ticker_minute.C:
			//if requests_today < r.requests_per_day // Also example
			// Allow for more requests this minute if we still have requests left today
			atomic.AddInt32(&requests_this_minute, -requests_this_minute)
		case <-ticker_second.C:
			// Allow for more requests this second if we still have requests left this minute
			if requests_this_minute < r.requests_per_minute {
				go func() {
					number_to_send := r.requests_per_second

					// If there are less requests left this minute than the specified rate per second
					// then send the lesser amount.
					left_this_minute := r.requests_per_minute - requests_this_minute // requests left this minute
					if left_this_minute < number_to_send {
						number_to_send = left_this_minute
					}

					// If there are less items in the queue than the current "number_to_send" then
					// send the lesser amount.
					if int32(len(r.queue)) < number_to_send {
						number_to_send = int32(len(r.queue))
					}

					// Consider the requests sent
					atomic.AddInt32(&requests_this_minute, number_to_send)
					for i := int32(0); i < number_to_send; i++ {
						// One go-routine per request
						go func() {
							currentRequest := <-r.queue
							re := result{}

							// Do we have to override the response?
							if r.override.override {
								currentRequest.ch <- r.override.data
							} else {
								re.statuscode, re.body = performRequest(currentRequest.url, currentRequest.params)
								currentRequest.ch <- re
							}
						}()
					}
				}()
			}
		}
	}
}
