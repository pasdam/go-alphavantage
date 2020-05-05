package alphavantage

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_mapCsvToQuotes(t *testing.T) {
	validTimestamps := []time.Time{
		time.Unix(1588370400, 0).UTC(),
		time.Unix(1588370405, 0).UTC(),
		time.Unix(1588370410, 0).UTC(),
		time.Unix(1588370415, 0).UTC(),
	}
	type args struct {
		csv [][]string
	}
	tests := []struct {
		name    string
		args    args
		want    []*Quote
		wantErr error
	}{
		{
			name: "Should parse all rows",
			args: args{
				csv: [][]string{
					{"2020-05-01 22:00:00", "1.1", "1.2", "1.3", "1.4"},
					{"2020-05-01 22:00:05", "2.1", "2.2", "2.3", "2.4"},
				},
			},
			want: []*Quote{
				{
					Timestamp: &validTimestamps[0],
					Open:      1.1,
					High:      1.2,
					Low:       1.3,
					Close:     1.4,
				},
				{
					Timestamp: &validTimestamps[1],
					Open:      2.1,
					High:      2.2,
					Low:       2.3,
					Close:     2.4,
				},
			},
			wantErr: nil,
		},
		{
			name: "Should ignore header row",
			args: args{
				csv: [][]string{
					{"timestamp", "open", "high", "low", "close"},
					{"2020-05-01 22:00:10", "3.1", "3.2", "3.3", "3.4"},
					{"2020-05-01 22:00:15", "4.1", "4.2", "4.3", "4.4"},
				},
			},
			want: []*Quote{
				{
					Timestamp: &validTimestamps[2],
					Open:      3.1,
					High:      3.2,
					Low:       3.3,
					Close:     3.4,
				},
				{
					Timestamp: &validTimestamps[3],
					Open:      4.1,
					High:      4.2,
					Low:       4.3,
					Close:     4.4,
				},
			},
			wantErr: nil,
		},
		{
			name: "Should return error if a row contains an invalid value",
			args: args{
				csv: [][]string{
					{"2020-05-01 22:00:10", "3.1", "3.2", "3.3", "3.4"},
					{"2020-05-01 22:00:15", "4.1", "4.2", "4.3", "invalid-value"},
				},
			},
			want:    nil,
			wantErr: errors.New("Error while reading CSV line 1. strconv.ParseFloat: parsing \"invalid-value\": invalid syntax"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := mapCsvToQuotes(tt.args.csv)

			if tt.wantErr != nil {
				assert.NotNil(t, err)
				assert.Equal(t, tt.wantErr.Error(), err.Error())
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
