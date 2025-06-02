package test

import (
	"net/http"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
)

type HTTPHandlerRegistry struct {
	hc               *http.Client
	requestResponses []requestResponseInfo
}

type requestResponseInfo struct {
	requestMethod string
	requestURL    string

	responseCode int
	responseBody string

	handlerTimeout time.Duration
}

// NewMockHTTPHandlerRegistry is a registry function to register mocked http handlers.
// It will mock any HTTP call from std http client or custom http client like Resty.
//
//	r := NewMockHTTPHandlerRegistry(nil)
//	r.NewHandler("GET", "http://example.test").Return(200, `{"message": "Aloha"}`)
//	r.NewHandler("GET", "http://example.test/list").Return(200, `{"data": []}`)
//	r.Register(t)
func NewMockHTTPHandlerRegistry(hc *http.Client) *HTTPHandlerRegistry {
	return &HTTPHandlerRegistry{
		hc:               hc,
		requestResponses: make([]requestResponseInfo, 0),
	}
}

func (r *HTTPHandlerRegistry) NewHandler(requestMethod, requestURL string) *handlerBuilder {
	reqRes := requestResponseInfo{}
	reqRes.requestMethod = requestMethod
	reqRes.requestURL = requestURL

	return &handlerBuilder{
		registry:        r,
		requestResponse: reqRes,
	}
}

func (r *HTTPHandlerRegistry) Register(t *testing.T) {
	if r.hc == nil {
		httpmock.Activate()
	} else {
		httpmock.ActivateNonDefault(r.hc)
	}

	t.Cleanup(func() {
		httpmock.DeactivateAndReset()
	})

	for _, reqRes := range r.requestResponses {
		// responder := httpmock.NewStringResponder(reqRes.responseCode, reqRes.responseBody)
		httpmock.RegisterResponder(reqRes.requestMethod, reqRes.requestURL, func(*http.Request) (*http.Response, error) {
			time.Sleep(reqRes.handlerTimeout)
			return httpmock.NewStringResponse(reqRes.responseCode, reqRes.responseBody), nil
		})
	}
}

type handlerBuilder struct {
	registry        *HTTPHandlerRegistry
	requestResponse requestResponseInfo
}

func (eb *handlerBuilder) WithTimeout(timeout time.Duration) *handlerBuilder {
	eb.requestResponse.handlerTimeout = timeout
	return eb
}

func (eb *handlerBuilder) Return(responseCode int, responseBody string) *HTTPHandlerRegistry {
	eb.requestResponse.responseCode = responseCode
	eb.requestResponse.responseBody = responseBody

	eb.registry.requestResponses = append(eb.registry.requestResponses, eb.requestResponse)
	return eb.registry
}
