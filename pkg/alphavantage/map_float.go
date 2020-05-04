package alphavantage

import (
	"strconv"
)

func mapFloat(value string) (float64, error) {
	if len(value) > 0 && value != "-" {
		result, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return 0, err
		}
		return result, nil
	}
	return 0, nil
}
