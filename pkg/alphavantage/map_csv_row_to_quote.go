package alphavantage

func mapCsvRowToQuote(row []string) (*Quote, error) {
	timestamp, err := mapTimestamp(row[0])
	if err != nil {
		return nil, err
	}

	open, err := mapFloat(row[1])
	if err != nil {
		return nil, err
	}

	high, err := mapFloat(row[2])
	if err != nil {
		return nil, err
	}

	low, err := mapFloat(row[3])
	if err != nil {
		return nil, err
	}

	close, err := mapFloat(row[4])
	if err != nil {
		return nil, err
	}

	return &Quote{
		Timestamp: timestamp,
		Open:      open,
		Close:     close,
		High:      high,
		Low:       low,
	}, nil
}
