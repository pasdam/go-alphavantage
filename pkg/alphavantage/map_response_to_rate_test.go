package alphavantage

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_mapResponseToRate(t *testing.T) {
	validTimestamps := []time.Time{
		time.Unix(1587511775, 0).UTC(),
		time.Unix(1587511785, 0).UTC(),
		time.Unix(1587511795, 0).UTC(),
	}
	type args struct {
		resp *currencyExchangeRateResponse
	}
	tests := []struct {
		name    string
		args    args
		want    *ExchangeRate
		wantErr error
	}{
		{
			name: "Should correctly map values",
			args: args{
				resp: &currencyExchangeRateResponse{
					Rate: realtimeCurrencyExchangeRate{
						FiveExchangeRate: "106.897",
						SixLastRefreshed: "2020-04-21 23:29:35",
						EightBidPrice:    "10.11",
						NineAskPrice:     "20.22",
					},
				},
			},
			want: &ExchangeRate{
				Rate:          106.897,
				Bid:           10.11,
				Ask:           20.22,
				LastRefreshed: &validTimestamps[0],
			},
			wantErr: nil,
		},
		{
			name: "Should correctly have values set to 0 if the response contains empty strings",
			args: args{
				resp: &currencyExchangeRateResponse{
					Rate: realtimeCurrencyExchangeRate{
						FiveExchangeRate: "",
						SixLastRefreshed: "2020-04-21 23:29:45",
						EightBidPrice:    "",
						NineAskPrice:     "",
					},
				},
			},
			want: &ExchangeRate{
				Rate:          0,
				Bid:           0,
				Ask:           0,
				LastRefreshed: &validTimestamps[1],
			},
			wantErr: nil,
		},
		{
			name: "Should correctly have values set to 0 if the response contains \"-\"",
			args: args{
				resp: &currencyExchangeRateResponse{
					Rate: realtimeCurrencyExchangeRate{
						FiveExchangeRate: "-",
						SixLastRefreshed: "2020-04-21 23:29:55",
						EightBidPrice:    "-",
						NineAskPrice:     "-",
					},
				},
			},
			want: &ExchangeRate{
				Rate:          0,
				Bid:           0,
				Ask:           0,
				LastRefreshed: &validTimestamps[2],
			},
			wantErr: nil,
		},
		{
			name: "Should return error if rate value is invalid",
			args: args{
				resp: &currencyExchangeRateResponse{
					Rate: realtimeCurrencyExchangeRate{
						FiveExchangeRate: "invalid-rate",
						SixLastRefreshed: "2020-04-22 23:29:35",
						EightBidPrice:    "42",
						NineAskPrice:     "56",
					},
				},
			},
			want:    nil,
			wantErr: errors.New("strconv.ParseFloat: parsing \"invalid-rate\": invalid syntax"),
		},
		{
			name: "Should return error if buy value is invalid",
			args: args{
				resp: &currencyExchangeRateResponse{
					Rate: realtimeCurrencyExchangeRate{
						SixLastRefreshed: "2020-04-22 23:29:35",
						EightBidPrice:    "invalid-buy",
						NineAskPrice:     "56",
					},
				},
			},
			want:    nil,
			wantErr: errors.New("strconv.ParseFloat: parsing \"invalid-buy\": invalid syntax"),
		},
		{
			name: "Should return error if sel value is invalid",
			args: args{
				resp: &currencyExchangeRateResponse{
					Rate: realtimeCurrencyExchangeRate{
						SixLastRefreshed: "2020-04-23 23:29:35",
						EightBidPrice:    "12",
						NineAskPrice:     "invalid-sell",
					},
				},
			},
			want:    nil,
			wantErr: errors.New("strconv.ParseFloat: parsing \"invalid-sell\": invalid syntax"),
		},
		{
			name: "Should return error if sel value is invalid",
			args: args{
				resp: &currencyExchangeRateResponse{
					Rate: realtimeCurrencyExchangeRate{
						SixLastRefreshed: "invalid-timestamp",
						EightBidPrice:    "67",
						NineAskPrice:     "98",
					},
				},
			},
			want:    nil,
			wantErr: errors.New("parsing time \"invalid-timestamp\" as \"2006-01-02 15:04:05\": cannot parse \"invalid-timestamp\" as \"2006\""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := mapResponseToRate(tt.args.resp)

			if tt.wantErr != nil {
				assert.Equal(t, tt.wantErr.Error(), err.Error())
			} else {
				assert.Nil(t, err)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
