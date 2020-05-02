package alphavantage

import (
	"net/url"

	"github.com/pasdam/go-rest-util/pkg/restutil"
)

type clientImpl struct {
	apiKey  string
	baseURL *url.URL
}

// NewClient creates a new Alphavantage client that used the specified API key
func NewClient(apiKey string) Client {
	c, _ := NewClientForURL(apiKey, "https://www.alphavantage.co")
	return c
}

// NewClientForURL creates a new Alphavantage client that used the specified API
// key and base URL
func NewClientForURL(apiKey string, baseURL string) (Client, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	u.Path = "query"
	return &clientImpl{
		apiKey:  apiKey,
		baseURL: u,
	}, nil
}

func (c *clientImpl) ExchangeRate(from CurrencyCode, to CurrencyCode) (*ExchangeRate, error) {
	response := &currencyExchangeRateResponse{}
	c.baseURL.RawQuery = currencyExchangeRateQuery(from.String(), to.String(), c.apiKey)
	err := restutil.GetJSON(c.baseURL.String(), response)
	if err != nil {
		return nil, err
	}

	return mapResponseToRate(response)
}
