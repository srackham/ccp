package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPrice_Success(t *testing.T) {
	// Mock successful HTTP GET request
	expectedPrice := 123.45
	mockBody := []byte(fmt.Sprintf(`{"symbol": "BTCUSDT", "price": "%f"}`, expectedPrice))
	mockClient := func(url string) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewReader(mockBody)),
		}, nil
	}

	price, err := GetPrice("BTC", mockClient)
	assert.NoError(t, err, "GetPrice should not return error")
	assert.Equal(t, expectedPrice, price, "GetPrice should return the correct price")
}

func TestGetPrice_ErrorScenarios(t *testing.T) {
	testCases := []struct {
		name           string
		mockResponse   *http.Response
		mockError      error
		expectedErrMsg string
	}{
		{
			name:           "HTTP Error",
			mockResponse:   nil,
			mockError:      errors.New("mock http error"),
			expectedErrMsg: "error making request:",
		},
		{
			name: "Invalid JSON in response",
			mockResponse: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(`{"invalid": json}`)),
			},
			mockError:      nil,
			expectedErrMsg: "error parsing JSON:",
		},
		{
			name: "Missing Price Data",
			mockResponse: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(`{"data":{}}`)),
			},
			mockError:      nil,
			expectedErrMsg: "error converting price to float:",
		},
		{
			name: "Non-200 Status Code",
			mockResponse: &http.Response{
				StatusCode: http.StatusNotFound,
				Body:       io.NopCloser(strings.NewReader(`{"error": "Not Found"}`)),
			},
			mockError:      nil,
			expectedErrMsg: "unexpected HTTP response status code:",
		},
		{
			name: "Invalid trading pair",
			mockResponse: &http.Response{
				StatusCode: http.StatusBadRequest,
				Body:       io.NopCloser(strings.NewReader(``)),
			},
			mockError:      nil,
			expectedErrMsg: "invalid trading pair:",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockGet := func(url string) (*http.Response, error) {
				return tc.mockResponse, tc.mockError
			}
			_, err := GetPrice("ETH", mockGet)
			assert.Error(t, err, "GetPrice should return error for %s", tc.name)
			assert.Contains(t, err.Error(), tc.expectedErrMsg, "Error message should indicate %s", tc.name)
		})
	}
}
