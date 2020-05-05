package alphavantage

import "testing"

func Test_intradayQuery(t *testing.T) {
	type args struct {
		from     string
		to       string
		interval ForexInterval
		compact  bool
		apiKey   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Should add outputSize parameter if compact is false",
			args: args{
				from:     "USD",
				to:       "EUR",
				interval: ForexInterval1Min,
				compact:  false,
				apiKey:   "some-api-key",
			},
			want: "apikey=some-api-key&datatype=csv&from_symbol=USD&function=FX_INTRADAY&interval=1min&outputSize=full&to_symbol=EUR",
		},
		{
			name: "Should not add outputSize parameter if compact is true",
			args: args{
				from:     "AUD",
				to:       "JPY",
				interval: ForexInterval15Min,
				compact:  true,
				apiKey:   "some-other-api-key",
			},
			want: "apikey=some-other-api-key&datatype=csv&from_symbol=AUD&function=FX_INTRADAY&interval=15min&to_symbol=JPY",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intradayQuery(tt.args.from, tt.args.to, tt.args.interval, tt.args.compact, tt.args.apiKey); got != tt.want {
				t.Errorf("intradayQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
