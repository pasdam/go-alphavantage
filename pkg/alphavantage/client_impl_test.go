package alphavantage

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/pasdam/mockit/mockit"
	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	type args struct {
		apiKey string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Should set baseURL and apiKey",
			args: args{
				apiKey: "some-api-key",
			},
		},
		{
			name: "Should set baseURL and different apiKey",
			args: args{
				apiKey: "some-other-api-key",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewClient(tt.args.apiKey).(*clientImpl)

			assert.Equal(t, tt.args.apiKey, got.apiKey)
			assert.Equal(t, "https://www.alphavantage.co/query", got.baseURL.String())
		})
	}
}

func TestNewClientForURL(t *testing.T) {
	type args struct {
		apiKey  string
		baseURL string
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "Should set baseURL and apiKey",
			args: args{
				apiKey:  "some-api-key",
				baseURL: "http://www.some-base.url",
			},
			wantErr: nil,
		},
		{
			name: "Should set baseURL and different apiKey",
			args: args{
				apiKey:  "some-other-api-key",
				baseURL: "http://www.some-other.url",
			},
			wantErr: nil,
		},
		{
			name: "Should return error if url is invalid",
			args: args{
				apiKey:  "",
				baseURL: "some\tinvalid\turl",
			},
			wantErr: errors.New("parse \"some\\tinvalid\\turl\": net/url: invalid control character in URL"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewClientForURL(tt.args.apiKey, tt.args.baseURL)
			if err == nil {
				c := got.(*clientImpl)
				assert.Equal(t, tt.args.apiKey, c.apiKey)
				assert.Equal(t, tt.args.baseURL+"/query", c.baseURL.String())
			} else {
				assert.Nil(t, got)
				assert.Equal(t, tt.wantErr.Error(), err.Error())
			}
		})
	}
}

func Test_clientImpl_ExchangeRate(t *testing.T) {
	type fields struct {
		apiKey  string
		baseURL string
	}
	type mocks struct {
		json     string
		response *currencyExchangeRateResponse
		rate     *ExchangeRate
		mapErr   error
	}
	type args struct {
		from CurrencyCode
		to   CurrencyCode
	}
	tests := []struct {
		name    string
		mocks   mocks
		fields  fields
		args    args
		wantErr error
	}{
		{
			name: "Should return expected exchange rate",
			mocks: mocks{
				json: `{
					"Realtime Currency Exchange Rate": {
						"1. From_Currency Code": "USD",
						"2. From_Currency Name": "United States Dollar",
						"3. To_Currency Code": "JPY",
						"4. To_Currency Name": "Japanese Yen",
						"5. Exchange Rate": "106.89700000",
						"6. Last Refreshed": "2020-05-01 14:34:45",
						"7. Time Zone": "UTC",
						"8. Bid Price": "0.123",
						"9. Ask Price": "1.234"
					}
				}`,
				response: &currencyExchangeRateResponse{
					Rate: realtimeCurrencyExchangeRate{
						OneFromCurrencyCode: "USD",
						TwoFromCurrencyName: "United States Dollar",
						ThreeToCurrencyCode: "JPY",
						FourToCurrencyName:  "Japanese Yen",
						FiveExchangeRate:    "106.89700000",
						SixLastRefreshed:    "2020-05-01 14:34:45",
						SevenTimeZone:       "UTC",
						EightBidPrice:       "0.123",
						NineAskPrice:        "1.234",
					},
				},
				rate:   &ExchangeRate{Rate: 106.897},
				mapErr: nil,
			},
			fields: fields{
				apiKey:  "some-success-api-key",
				baseURL: "some-success-path",
			},
			args: args{
				from: CurrencyCodeUSD,
				to:   CurrencyCodeJPY,
			},
			wantErr: nil,
		},
		{
			name: "Should return error raised by mapResponseToRate",
			mocks: mocks{
				json: `{
					"Realtime Currency Exchange Rate": {
						"1. From_Currency Code": "EUR",
						"2. From_Currency Name": "Euro",
						"3. To_Currency Code": "AUD",
						"4. To_Currency Name": "Australian Dollar",
						"5. Exchange Rate": "206.89700000",
						"6. Last Refreshed": "2021-05-01 14:34:45",
						"7. Time Zone": "CEST",
						"8. Bid Price": "1.123",
						"9. Ask Price": "2.234"
					}
				}`,
				response: &currencyExchangeRateResponse{
					Rate: realtimeCurrencyExchangeRate{
						OneFromCurrencyCode: "EUR",
						TwoFromCurrencyName: "Euro",
						ThreeToCurrencyCode: "AUD",
						FourToCurrencyName:  "Australian Dollar",
						FiveExchangeRate:    "206.89700000",
						SixLastRefreshed:    "2021-05-01 14:34:45",
						SevenTimeZone:       "CEST",
						EightBidPrice:       "1.123",
						NineAskPrice:        "2.234",
					},
				},
				rate:   nil,
				mapErr: errors.New("some-map-error"),
			},
			fields: fields{
				apiKey:  "some-map-error-api-key",
				baseURL: "some-map-error-path",
			},
			args: args{
				from: CurrencyCodeEUR,
				to:   CurrencyCodeAUD,
			},
			wantErr: errors.New("some-map-error"),
		},
		{
			name: "Should return error raised by json parsing",
			mocks: mocks{
				json:     "some-invalid-json",
				response: &currencyExchangeRateResponse{},
				rate:     nil,
				mapErr:   nil,
			},
			fields: fields{
				apiKey:  "some-map-error-api-key",
				baseURL: "some-map-error-path",
			},
			args: args{
				from: CurrencyCodeCHF,
				to:   CurrencyCodeIDR,
			},
			wantErr: errors.New("invalid character 's' looking for beginning of value"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockit.MockFunc(t, mapResponseToRate).With(tt.mocks.response).Return(tt.mocks.rate, tt.mocks.mapErr)
			called := false
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				called = true
				assert.Equal(t, "/"+tt.fields.baseURL, r.URL.Path)
				q := r.URL.Query()
				assert.Equal(t, q.Get("function"), "CURRENCY_EXCHANGE_RATE")
				assert.Equal(t, q.Get("from_currency"), tt.args.from.String())
				assert.Equal(t, q.Get("to_currency"), tt.args.to.String())
				assert.Equal(t, q.Get("apikey"), tt.fields.apiKey)
				fmt.Fprintln(w, tt.mocks.json)
			}))
			u, err := url.Parse(ts.URL + "/" + tt.fields.baseURL)
			assert.Nil(t, err)
			c := &clientImpl{
				apiKey:  tt.fields.apiKey,
				baseURL: u,
			}

			got, err := c.ExchangeRate(tt.args.from, tt.args.to)

			assert.True(t, called)
			if tt.wantErr != nil {
				assert.Nil(t, got)
				assert.Equal(t, tt.wantErr.Error(), err.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.mocks.rate, got)
			}
		})
	}
}
