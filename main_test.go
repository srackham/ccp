package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
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

func TestGetPrice_HttpError(t *testing.T) {
	// Mock failing HTTP GET request
	mockClient := func(url string) (*http.Response, error) {
		return nil, errors.New("mock http error")
	}

	_, err := GetPrice("ETH", mockClient)
	assert.Error(t, err, "GetPrice should return error for failed request")
	assert.Contains(t, err.Error(), "error making request:", "Error message should indicate HTTP error")
}

// ... other test cases (ReadBodyError, JsonUnmarshalError, PriceConversionError)
