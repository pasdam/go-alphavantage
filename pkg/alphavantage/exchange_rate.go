package alphavantage

import (
	"time"
)

// ExchangeRate contains info about the exchange rate between two currencies
type ExchangeRate struct {

	// From is the base currency
	From Currency

	// To is the quote currency
	To Currency

	// LastRefreshed is the timestamp of the last refresh for the currency pair
	LastRefreshed *time.Time

	// Rate current exchange rate
	Rate float64

	// Bid bid price
	Bid float64

	// Ask ask price
	Ask float64
}
