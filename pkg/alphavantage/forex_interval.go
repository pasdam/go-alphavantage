package alphavantage

import "time"

// ForexInterval enumerates the supported interval values for
type ForexInterval time.Duration

const (

	// ForexInterval1Min indicates an interval of 1 minutes
	ForexInterval1Min = 60000000 // 1 m = 1*60*1000*1000 ns

	// ForexInterval5Min indicates an interval of 5 minutes
	ForexInterval5Min = 5 * 60000000 // 5m")

	// ForexInterval15Min indicates an interval of 15 minutes
	ForexInterval15Min = 15 * 60000000 // 15m")

	// ForexInterval30Min indicates an interval of 30 minutes
	ForexInterval30Min = 30 * 60000000 // 30m")

	// ForexInterval60Min indicates an interval of 60 minutes
	ForexInterval60Min = 60 * 60000000 // 60m")
)

func (i ForexInterval) String() string {
	switch i {
	case ForexInterval1Min:
		return "1min"
	case ForexInterval5Min:
		return "5min"
	case ForexInterval15Min:
		return "15min"
	case ForexInterval30Min:
		return "30min"
	case ForexInterval60Min:
		return "60min"
	default:
		return ""
	}
}
