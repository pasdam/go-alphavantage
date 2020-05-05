package alphavantage

import (
	"net/url"
)

func intradayQuery(from string, to string, interval ForexInterval, compact bool, apiKey string) string {
	values := url.Values{}
	values.Add("function", "FX_INTRADAY")
	values.Add("from_symbol", from)
	values.Add("to_symbol", to)
	values.Add("apikey", apiKey)
	if !compact {
		values.Add("outputSize", "full")
	}
	values.Add("datatype", "csv")
	values.Add("interval", interval.String())
	return values.Encode()
}
