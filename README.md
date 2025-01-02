# Current Crypto Prices
13-Nov-2024

`ccp` is a simple minded Go CLI command to fetch and display crypto prices.

- Fetches prices using the Binance HTTP ticker price API.

## TODO
- Finish tests.
- Better error handling e.g. on invalid symbol.
- Add --help option.
- Add --date option to prefix output with dd-mmm-yyyy hh:mm, NO: if you want the date write an alias e.g.

        btc='echo $(date +%H:%M): $(ccp BTC)'


