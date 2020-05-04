package alphavantage

func mapResponseToRate(resp *currencyExchangeRateResponse) (*ExchangeRate, error) {
	rate, err := mapFloat(resp.Rate.FiveExchangeRate)
	if err != nil {
		return nil, err
	}

	bid, err := mapFloat(resp.Rate.EightBidPrice)
	if err != nil {
		return nil, err
	}

	ask, err := mapFloat(resp.Rate.NineAskPrice)
	if err != nil {
		return nil, err
	}

	lastRefreshed, err := mapTimestamp(resp.Rate.SixLastRefreshed)
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
