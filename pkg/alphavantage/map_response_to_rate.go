package alphavantage

import (
	"strconv"
	"time"
)

func mapResponseToRate(resp *currencyExchangeRateResponse) (*ExchangeRate, error) {
	var err error

	var rate float64 = 0
	if len(resp.Rate.FiveExchangeRate) > 0 && resp.Rate.FiveExchangeRate != "-" {
		rate, err = strconv.ParseFloat(resp.Rate.FiveExchangeRate, 64)
		if err != nil {
			return nil, err
		}
	}

	var bid float64 = 0
	if len(resp.Rate.EightBidPrice) > 0 && resp.Rate.EightBidPrice != "-" {
		bid, err = strconv.ParseFloat(resp.Rate.EightBidPrice, 64)
		if err != nil {
			return nil, err
		}
	}

	var ask float64 = 0
	if len(resp.Rate.NineAskPrice) > 0 && resp.Rate.NineAskPrice != "-" {
		ask, err = strconv.ParseFloat(resp.Rate.NineAskPrice, 64)
		if err != nil {
			return nil, err
		}
	}

	lastRefreshed, err := time.Parse("2006-01-02 15:04:05", resp.Rate.SixLastRefreshed)
	if err != nil {
		return nil, err
	}

	return &ExchangeRate{
		Rate:          rate,
		Bid:           bid,
		Ask:           ask,
		LastRefreshed: lastRefreshed,
	}, err
}
