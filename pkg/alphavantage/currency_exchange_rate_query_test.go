package alphavantage

import (
	"testing"
)

func Test_currencyExchangeRateQuery(t *testing.T) {
	type args struct {
		from   string
		to     string
		apiKey string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Should return expected query",
			args: args{
				from:   "EUR",
				to:     "USD",
				apiKey: "some-api-key",
			},
			want: "apikey=some-api-key&from_currency=EUR&function=CURRENCY_EXCHANGE_RATE&to_currency=USD",
		},
		{
			name: "Should return expected query with different parameters",
			args: args{
				from:   "USD",
				to:     "IDR",
				apiKey: "some-other-key",
			},
			want: "apikey=some-other-key&from_currency=USD&function=CURRENCY_EXCHANGE_RATE&to_currency=IDR",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := currencyExchangeRateQuery(tt.args.from, tt.args.to, tt.args.apiKey); got != tt.want {
				t.Errorf("currencyExchangeRateQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
