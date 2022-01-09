package httpmock

import (
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
)

// Response Wraps the response of the RoundTrip.
type Response struct {
	Response *http.Response
	Error    error
}

// MockRoundTripper implements http.RoundTripper and returns a pair of an http.Response and error
// when the RoundTrip method is invoked.
type MockRoundTripper struct {
	Responses []*Response

	iteration int32
	mu        sync.RWMutex
}

// NewMockClient returns an *http.Client that uses a mocked transport to avoid real http calls.
// The mocked transport is able to handle multiple responses.
func NewMockClient(r ...*Response) *http.Client {
	return &http.Client{
		Transport: NewMockRoundTripper(r...),
	}
}

// NewMockRoundTripper initializes a new RoundTripper with the given list of Response.
func NewMockRoundTripper(r ...*Response) *MockRoundTripper {
	return &MockRoundTripper{
		Responses: r,
	}
}

// NewMockResponse returns a Response reference initialized with the given http.Response and the error.
func NewMockResponse(resp *http.Response, err error) *Response {
	return &Response{
		Response: resp,
		Error:    err,
	}
}

// RoundTrip returns a pre-defined response and error based on the mocks provided by the user
// during the initialisation of MockRoundTripper.
func (mock *MockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	mock.mu.RLock()
	defer func() {
		atomic.AddInt32(&mock.iteration, 1)
		mock.mu.RUnlock()
	}()

	iteration := atomic.LoadInt32(&mock.iteration)
	if int(iteration) > len(mock.Responses)-1 {
		return nil, fmt.Errorf(
			"no mocked response found for iteration %d: %s %s",
			iteration+1, req.Method, req.URL,
		)
	}

	r := mock.Responses[iteration]

	return r.Response, r.Error
}
