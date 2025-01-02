package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

type TickerPrice struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

// Define a custom type for the HTTP Get function
type httpGet func(url string) (*http.Response, error)

func GetPrice(symbol string, get httpGet) (float64, error) {
	url := "https://api.binance.com/api/v1/ticker/price?symbol=" + symbol + "USDT"

	resp, err := get(url)
	if err != nil {
		return 0, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("error reading response: %v", err)
	}

	var ticker TickerPrice
	err = json.Unmarshal(body, &ticker)
	if err != nil {
		return 0, fmt.Errorf("error parsing JSON: %v", err)
	}

	price, err := strconv.ParseFloat(ticker.Price, 64)
	if err != nil {
		return 0, fmt.Errorf("error converting price to float: %v", err)
	}

	return price, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ccp <symbol1> <symbol2> ...")
		os.Exit(1)
	}

	symbols := os.Args[1:]

	err_count := 0
	for _, symbol := range symbols {
		price, err := GetPrice(symbol, http.Get)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting price for %s: %v\n", symbol, err)
			err_count += 1
			continue
		}
		fmt.Printf("%s: $%.2f USD\n", symbol, price)
	}
	if err_count > 0 {
		os.Exit(1)
	}
}
