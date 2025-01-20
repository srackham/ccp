# Crypto Currency Prices

A simple CLI command to fetch crypto prices.

    Usage: ccp SYMBOL...

`SYMBOL` is a ticker symbol e.g. `BTC`:

```
$ ccp BTC ETH
BTC: $104486.83 USD
ETH: $3353.79 USD
```

## Implementation

- Fetches prices using the [Binance HTTP ticker price API](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md#symbol-price-ticker) and displays the price in USD.
- The price is that of the `<symbol>USDT` spot pair e.g. the `BTCUSDT` spot pair is used to fetch the Bitcoin price.

## Installation

Clone the [ccp Github repo](https://github.com/srackham/ccp) and compile and install using the `go` command:

```
$ git clone https://github.com/srackham/ccp.git
$ cd ccp
$ go install
```
