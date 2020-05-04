package alphavantage

import (
	"time"
)

func mapTimestamp(value string) (*time.Time, error) {
	t, err := time.Parse("2006-01-02 15:04:05", value)
	if err != nil {
		return nil, err
	}
	return &t, nil
}
