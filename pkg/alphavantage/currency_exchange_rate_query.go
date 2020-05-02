package alphavantage

import (
	"net/url"
)

func currencyExchangeRateQuery(from string, to string, apiKey string) string {
	values := url.Values{}
	values.Add("function", "CURRENCY_EXCHANGE_RATE")
	values.Add("from_currency", from)
	values.Add("to_currency", to)
	values.Add("apikey", apiKey)
	return values.Encode()
}
