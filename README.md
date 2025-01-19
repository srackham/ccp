# Crypto Currency Prices

A simple CLI command to fetch crypto prices.
Fetches prices using the Binance HTTP ticker price API and displays the price in USD.

    Usage: ccp SYMBOL...

SYMBOL is a ticker symbol e.g. BTC:
```
$ ccp BTC ETH
BTC: $104486.83 USD
ETH: $3353.79 USD
```

## Installation
Clone the Github repo and compile and install using the `go` command:

```
$ go install

```
