package alphavantage

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

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

func Test_clientImpl_ExchangeRate_RealService(t *testing.T) {
	t.Skip("This is used only for manual testing")

	c := NewClient("demo") // Replace the api key with a valid one to run the test
	r, err := c.ExchangeRate(CurrencyCodeUSD, CurrencyCodeEUR)
	assert.Nil(t, err)
	assert.NotNil(t, r)
	assert.True(t, r.Rate > 0)
}

func Test_clientImpl_Quotes(t *testing.T) {
	validTimestamps := []time.Time{
		time.Unix(1588370400, 0).UTC(),
		time.Unix(1588370405, 0).UTC(),
	}
	type mocks struct {
		csv string
	}
	type fields struct {
		apiKey  string
		baseURL string
	}
	type args struct {
		from     CurrencyCode
		to       CurrencyCode
		interval ForexInterval
	}
	tests := []struct {
		name             string
		mocks            mocks
		fields           fields
		args             args
		shouldCallServer bool
		want             []*Quote
		wantErr          error
	}{
		{
			name: "Should correctly parse successful response",
			mocks: mocks{
				csv: "timestamp,open,high,low,close\n2020-05-01 22:00:00,1.001,1.002,1.003,1.004\n2020-05-01 22:00:05,2.001,2.002,2.003,2.004",
			},
			fields: fields{
				apiKey:  "some-success-api-key",
				baseURL: "some-success-path",
			},
			args: args{
				from:     CurrencyCodeUSD,
				to:       CurrencyCodeEUR,
				interval: ForexInterval1Min,
			},
			shouldCallServer: true,
			want: []*Quote{
				{
					Timestamp: &validTimestamps[0],
					Open:      1.001,
					High:      1.002,
					Low:       1.003,
					Close:     1.004,
				},
				{
					Timestamp: &validTimestamps[1],
					Open:      2.001,
					High:      2.002,
					Low:       2.003,
					Close:     2.004,
				},
			},
			wantErr: nil,
		},
		{
			name: "Should return error if restutil.GetCSV raises it",
			mocks: mocks{
				csv: "\"invalid-csv",
			},
			fields: fields{
				apiKey:  "some-get-csv-error-api-key",
				baseURL: "some-get-csv-error-path",
			},
			args: args{
				from:     CurrencyCodeINR,
				to:       CurrencyCodeIDR,
				interval: ForexInterval30Min,
			},
			shouldCallServer: true,
			want:             nil,
			wantErr:          errors.New("Unable to read CSV from the response. record on line 1; parse error on line 2, column 0: extraneous or missing \" in quoted-field"),
		},
		{
			name: "Should not perform network call and return error if from is not a physical currency",
			mocks: mocks{
				csv: "",
			},
			fields: fields{
				apiKey:  "some-not-physical-from-api-key",
				baseURL: "some-not-physical-from-path",
			},
			args: args{
				from:     CurrencyCodeBTC,
				to:       CurrencyCodeJPY,
				interval: ForexInterval30Min,
			},
			shouldCallServer: false,
			want:             nil,
			wantErr:          errors.New("Quotes are not supported for digital currencies"),
		},
		{
			name: "Should not perform network call and return error if from is not a physical currency",
			mocks: mocks{
				csv: "",
			},
			fields: fields{
				apiKey:  "some-not-physical-from-api-key",
				baseURL: "some-not-physical-from-path",
			},
			args: args{
				from:     CurrencyCodeJPY,
				to:       CurrencyCodeBTC,
				interval: ForexInterval5Min,
			},
			shouldCallServer: false,
			want:             nil,
			wantErr:          errors.New("Quotes are not supported for digital currencies"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			called := false
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				called = true
				assert.Equal(t, "/"+tt.fields.baseURL, r.URL.Path)
				q := r.URL.Query()
				assert.Equal(t, q.Get("function"), "FX_INTRADAY")
				assert.Equal(t, q.Get("from_symbol"), tt.args.from.String())
				assert.Equal(t, q.Get("to_symbol"), tt.args.to.String())
				assert.Equal(t, q.Get("apikey"), tt.fields.apiKey)
				assert.Equal(t, q.Get("outputSize"), "full")
				assert.Equal(t, q.Get("datatype"), "csv")
				assert.Equal(t, q.Get("interval"), tt.args.interval.String())
				fmt.Fprintln(w, tt.mocks.csv)
			}))
			u, err := url.Parse(ts.URL + "/" + tt.fields.baseURL)
			assert.Nil(t, err)
			c := &clientImpl{
				apiKey:  tt.fields.apiKey,
				baseURL: u,
			}

			got, err := c.Quotes(tt.args.from, tt.args.to, tt.args.interval)

			assert.Equal(t, tt.shouldCallServer, called)
			if tt.wantErr != nil {
				assert.Nil(t, got)
				assert.Equal(t, tt.wantErr.Error(), err.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func Test_clientImpl_QuotesCompact(t *testing.T) {
	validTimestamps := []time.Time{
		time.Unix(1588370400, 0).UTC(),
		time.Unix(1588370405, 0).UTC(),
	}
	type mocks struct {
		csv string
	}
	type fields struct {
		apiKey  string
		baseURL string
	}
	type args struct {
		from     CurrencyCode
		to       CurrencyCode
		interval ForexInterval
	}
	tests := []struct {
		name             string
		mocks            mocks
		fields           fields
		args             args
		shouldCallServer bool
		want             []*Quote
		wantErr          error
	}{
		{
			name: "Should correctly parse successful response",
			mocks: mocks{
				csv: "timestamp,open,high,low,close\n2020-05-01 22:00:00,1.001,1.002,1.003,1.004\n2020-05-01 22:00:05,2.001,2.002,2.003,2.004",
			},
			fields: fields{
				apiKey:  "some-success-api-key",
				baseURL: "some-success-path",
			},
			args: args{
				from:     CurrencyCodeUSD,
				to:       CurrencyCodeEUR,
				interval: ForexInterval1Min,
			},
			shouldCallServer: true,
			want: []*Quote{
				{
					Timestamp: &validTimestamps[0],
					Open:      1.001,
					High:      1.002,
					Low:       1.003,
					Close:     1.004,
				},
				{
					Timestamp: &validTimestamps[1],
					Open:      2.001,
					High:      2.002,
					Low:       2.003,
					Close:     2.004,
				},
			},
			wantErr: nil,
		},
		{
			name: "Should return error if restutil.GetCSV raises it",
			mocks: mocks{
				csv: "\"invalid-csv",
			},
			fields: fields{
				apiKey:  "some-get-csv-error-api-key",
				baseURL: "some-get-csv-error-path",
			},
			args: args{
				from:     CurrencyCodeINR,
				to:       CurrencyCodeIDR,
				interval: ForexInterval30Min,
			},
			shouldCallServer: true,
			want:             nil,
			wantErr:          errors.New("Unable to read CSV from the response. record on line 1; parse error on line 2, column 0: extraneous or missing \" in quoted-field"),
		},
		{
			name: "Should not perform network call and return error if from is not a physical currency",
			mocks: mocks{
				csv: "",
			},
			fields: fields{
				apiKey:  "some-not-physical-from-api-key",
				baseURL: "some-not-physical-from-path",
			},
			args: args{
				from:     CurrencyCodeBTC,
				to:       CurrencyCodeJPY,
				interval: ForexInterval30Min,
			},
			shouldCallServer: false,
			want:             nil,
			wantErr:          errors.New("Quotes are not supported for digital currencies"),
		},
		{
			name: "Should not perform network call and return error if from is not a physical currency",
			mocks: mocks{
				csv: "",
			},
			fields: fields{
				apiKey:  "some-not-physical-from-api-key",
				baseURL: "some-not-physical-from-path",
			},
			args: args{
				from:     CurrencyCodeJPY,
				to:       CurrencyCodeBTC,
				interval: ForexInterval5Min,
			},
			shouldCallServer: false,
			want:             nil,
			wantErr:          errors.New("Quotes are not supported for digital currencies"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			called := false
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				called = true
				assert.Equal(t, "/"+tt.fields.baseURL, r.URL.Path)
				q := r.URL.Query()
				assert.Equal(t, q.Get("function"), "FX_INTRADAY")
				assert.Equal(t, q.Get("from_symbol"), tt.args.from.String())
				assert.Equal(t, q.Get("to_symbol"), tt.args.to.String())
				assert.Equal(t, q.Get("apikey"), tt.fields.apiKey)
				assert.Empty(t, q.Get("outputSize"))
				assert.Equal(t, q.Get("datatype"), "csv")
				assert.Equal(t, q.Get("interval"), tt.args.interval.String())
				fmt.Fprintln(w, tt.mocks.csv)
			}))
			u, err := url.Parse(ts.URL + "/" + tt.fields.baseURL)
			assert.Nil(t, err)
			c := &clientImpl{
				apiKey:  tt.fields.apiKey,
				baseURL: u,
			}

			got, err := c.QuotesCompact(tt.args.from, tt.args.to, tt.args.interval)

			assert.Equal(t, tt.shouldCallServer, called)
			if tt.wantErr != nil {
				assert.Nil(t, got)
				assert.Equal(t, tt.wantErr.Error(), err.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func Test_clientImpl_QuotestCompact_RealService(t *testing.T) {
	t.Skip("This is used only for manual testing")

	c := NewClient("demo") // Replace the api key with a valid one to run the test
	q, err := c.QuotesCompact(CurrencyCodeUSD, CurrencyCodeEUR, ForexInterval15Min)
	assert.Nil(t, err)
	assert.NotNil(t, q)
	assert.Equal(t, 100, len(q))
}
