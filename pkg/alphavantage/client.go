package alphavantage

// Client is the interface to interact with the Alphavantage API
type Client interface {

	// ExchangeRate will retrieve the current exchange rate for the specified
	// currency pair
	ExchangeRate(from CurrencyCode, to CurrencyCode) (*ExchangeRate, error)
}
