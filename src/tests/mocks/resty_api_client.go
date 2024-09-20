package mocks

import (
	"io"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
)

// Define a mock client that implements the ApiClient interface
type MockApiClient struct {
	restyClient *resty.Client
	isClientOk  bool
}

// Custom RoundTripper to mock HTTP responses
type MockRoundTripper struct {
	response *http.Response
	err      error
}

// RoundTrip function to mock an HTTP request and response
func (m *MockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.response, m.err
}

// GetClient returns the underlying resty client
func (m *MockApiClient) GetClient() *resty.Client {
	return m.restyClient
}

// IsClientOk returns whether the client is in a valid state
func (m *MockApiClient) IsClientOk() bool {
	return m.isClientOk
}

// Constructor function to initialize the MockApiClient with mock behavior
func NewMockApiClient(responseBody string, err error, clientOk bool) *MockApiClient {
	restyClient := resty.New()

	// Create a mock HTTP response
	response := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(strings.NewReader(responseBody)),
		Header:     make(http.Header),
	}

	// Use a custom RoundTripper to simulate HTTP requests
	restyClient.SetTransport(&MockRoundTripper{
		response: response,
		err:      err,
	})

	return &MockApiClient{
		restyClient: restyClient,
		isClientOk:  clientOk,
	}
}
