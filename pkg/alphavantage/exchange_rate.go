package alphavantage

import (
	"time"
)

// ExchangeRate contains info about the exchange rate between two currencies
type ExchangeRate struct {
	From          Currency
	To            Currency
	LastRefreshed time.Time
	Rate          float64
	Bid           float64
	Ask           float64
}
