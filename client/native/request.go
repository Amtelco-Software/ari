package native

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

// MaxIdleConnections is the maximum number of idle web client
// connections to maintain.
var MaxIdleConnections = 20

// RequestTimeout describes the maximum amount of time to wait
// for a response to any request.
var RequestTimeout = 2 * time.Second

// RequestError describes an error with an error Code.
type RequestError interface {
	error
	Code() int
}

type requestError struct {
	statusCode int
	text       string
}

// Error returns the request error as a string.
func (e *requestError) Error() string {
	return e.text
}

// Code returns the status code from the request.
func (e *requestError) Code() int {
	return e.statusCode
}

// CodeFromError extracts and returns the code from an error, or
// 0 if not found.
func CodeFromError(err error) int {
	if reqerr, ok := err.(RequestError); ok {
		return reqerr.Code()
	}
	return 0
}

func maybeRequestError(resp *http.Response) RequestError {
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		// 2xx response: All good.
		return nil
	}
	return &requestError{
		text:       "Non-2XX response: " + resp.Status,
		statusCode: resp.StatusCode,
	}
}

// MissingParams is an error message response emitted when a request
// does not contain required parameters
type MissingParams struct {
	//Message
	Type   string   `json:"type"`
	Params []string `json:"params"` // List of missing parameters which are required
}

// Get calls the ARI server with a GET request
func (c *Conn) Get(url string, resp interface{}) error {

	url = c.Options.URL + url

	return c.makeRequest("GET", url, resp, nil)
}

// Post calls the ARI server with a POST request.
func (c *Conn) Post(requestURL string, resp interface{}, req interface{}) error {

	url = c.Options.URL + requestURL

	return c.makeRequest("POST", url, resp, req)
}

// Put calls the ARI server with a PUT request.
func (c *Conn) Put(url string, resp interface{}, req interface{}) error {

	url = c.Options.URL + url

	return c.makeRequest("PUT", url, resp, req)
}

// Delete calls the ARI server with a DELETE request
func (c *Conn) Delete(url string, resp interface{}, req string) error {

	url = c.Options.URL + url
	if req != "" {
		url = url + "?" + req
	}

	return c.makeRequest("DELETE", url, resp, req)
}

func (c *Conn) makeRequest(method, url, resp interface{}, req interface{}) (err error) {
	var reqBody io.Reader
	if req != nil {
		reqBody, err = structToRequestBody(req)
		if err != nil {
			return errors.Wrap(err, "failed to marshal request")
		}
	}

	r, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, err
	}
	r.Header.Set("Content-Type", "application/json")

	if c.Options.Username != "" {
		req.SetBasicAuth(c.Options.Username, c.Options.Password)
	}

	ret, err := c.httpClient.Do(r)
	if err != nil {
		return errors.Wrap(err, "failed to make request")
	}
	defer ret.Body.Close()

	if resp != nil {
		err = json.NewDecoder(ret.Body).Decode(resp)
		if err != nil {
			return errors.Wrap(err, "failed to decode response")
		}
	}

	return maybeRequestError(ret)
}

func structToRequestBody(req interface{}) (io.Reader, error) {
	buf := new(bytes.Buffer)
	if req != nil {
		if err := json.NewEncoder(buf).Encode(req); err != nil {
			return nil, err
		}
	}

	return buf, nil
}
