package alphavantage

import (
	"fmt"
)

func mapCsvToQuotes(csv [][]string) ([]*Quote, error) {
	if csv[0][0] == "timestamp" {
		// drop header
		csv = csv[1:]
	}
	result := make([]*Quote, len(csv))
	var err error
	for i := 0; i < len(csv); i++ {
		result[i], err = mapCsvRowToQuote(csv[i])
		if err != nil {
			return nil, fmt.Errorf("Error while reading CSV line %d. %s", i, err.Error())
		}
	}

	return result, nil
}
