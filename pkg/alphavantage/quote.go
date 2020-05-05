package alphavantage

import (
	"time"
)

// Quote contains the details of a price quote at a specific time
type Quote struct {

	// Timestamp indicates the timestamp of the data point
	Timestamp *time.Time

	// Open indicates the open value
	Open float64

	// Close indicates the close value
	Close float64

	// High indicates the highest value in the interval
	High float64

	// Low indicates the lowesr value in the interval
	Low float64
}
