package alphavantage

import "testing"

func TestForexInterval_String(t *testing.T) {
	tests := []struct {
		name string
		i    ForexInterval
	}{
		{name: "1min", i: ForexInterval1Min},
		{name: "5min", i: ForexInterval5Min},
		{name: "15min", i: ForexInterval15Min},
		{name: "30min", i: ForexInterval30Min},
		{name: "60min", i: ForexInterval60Min},
		{name: "", i: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.String(); got != tt.name {
				t.Errorf("ForexInterval.String() = %v, want %v", got, tt.name)
			}
		})
	}
}
