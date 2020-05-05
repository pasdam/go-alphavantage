package alphavantage

// Client is the interface to interact with the Alphavantage API
type Client interface {

	// ExchangeRate will retrieve the current exchange rate for the specified
	// currency pair
	ExchangeRate(from CurrencyCode, to CurrencyCode) (*ExchangeRate, error)

	// Quotes retrieve the intra day time series of the specified currency
	// pair.
	// Supported intervals are: 1min, 5min, 15min, 30min, 60min; other values
	// will result in an error.
	Quotes(from CurrencyCode, to CurrencyCode, interval ForexInterval) ([]*Quote, error)

	// QuotesCompact retrieve the latest 100 data points of the intra day time
	// series of the specified currency pair.
	// Supported intervals are: 1min, 5min, 15min, 30min, 60min; other values
	// will result in an error.
	QuotesCompact(from CurrencyCode, to CurrencyCode, interval ForexInterval) ([]*Quote, error)
}
