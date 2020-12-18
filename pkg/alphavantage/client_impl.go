package alphavantage

import (
	"errors"
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
	headers := map[string][]string{
		"User-Agent": {"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.16; rv:84.0) Gecko/20100101 Firefox/84.0"},
	}
	err := restutil.GetJSON(c.baseURL.String(), headers, response)
	if err != nil {
		return nil, err
	}

	return mapResponseToRate(response)
}

func (c *clientImpl) Quotes(from CurrencyCode, to CurrencyCode, interval ForexInterval) ([]*Quote, error) {
	return c.getQuotes(from, to, interval, true)
}

func (c *clientImpl) QuotesCompact(from CurrencyCode, to CurrencyCode, interval ForexInterval) ([]*Quote, error) {
	return c.getQuotes(from, to, interval, false)
}

func (c *clientImpl) getQuotes(from CurrencyCode, to CurrencyCode, interval ForexInterval, full bool) ([]*Quote, error) {
	if !from.IsPhysical() || !to.IsPhysical() {
		return nil, errors.New("Quotes are not supported for digital currencies")
	}

	c.baseURL.RawQuery = intradayQuery(from.String(), to.String(), interval, !full, c.apiKey)
	rows, err := restutil.GetCSV(c.baseURL.String())
	if err != nil {
		return nil, errors.New("Unable to read CSV from the response. " + err.Error())
	}
	return mapCsvToQuotes(rows)
}
