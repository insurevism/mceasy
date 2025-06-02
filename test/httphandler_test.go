package test_test

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"mceasy/pkg/tokopediaclient/tests"

	"github.com/stretchr/testify/assert"
)

func TestNewMockHTTPHandlerRegistry(t *testing.T) {
	requestURL := "http://example.com/foo"

	// define a registry
	r := tests.NewMockHTTPHandlerRegistry(nil)

	// define a handler that we want to mock
	r.
		NewHandler(http.MethodGet, requestURL).
		Return(http.StatusOK, `{"message":"foo"}`).
		Register(t)

	// hit the handler with http client
	resp, err := http.Get(requestURL)
	assert.NoError(t, err)

	defer resp.Body.Close()

	// validate the response
	respBytes, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)

	var response struct {
		Message string `json:"message"`
	}
	err = json.Unmarshal(respBytes, &response)
	assert.NoError(t, err)
	assert.Equal(t, "foo", response.Message)
}
