package alphavantage

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_mapCsvRowToQuote(t *testing.T) {
	validTimestamps := []time.Time{
		time.Unix(1588370400, 0).UTC(),
		time.Unix(1588370410, 0).UTC(),
	}
	type args struct {
		row []string
	}
	tests := []struct {
		name    string
		args    args
		want    *Quote
		wantErr error
	}{
		{
			name: "Should correctly parse valid values",
			args: args{
				row: []string{"2020-05-01 22:00:00", "1.0981", "1.0982", "1.0983", "1.0984"},
			},
			want: &Quote{
				Timestamp: &validTimestamps[0],
				Open:      1.0981,
				High:      1.0982,
				Low:       1.0983,
				Close:     1.0984,
			},
			wantErr: nil,
		},
		{
			name: "Should return 0 if csv row contains '-'",
			args: args{
				row: []string{"2020-05-01 22:00:10", "-", "-", "-", "-"},
			},
			want: &Quote{
				Timestamp: &validTimestamps[1],
				Open:      0,
				High:      0,
				Low:       0,
				Close:     0,
			},
			wantErr: nil,
		},
		{
			name: "Should return error if timestamp is invalid",
			args: args{
				row: []string{"invalid-timestamp", "1.0981", "1.0982", "1.0983", "1.0984"},
			},
			want:    nil,
			wantErr: errors.New("parsing time \"invalid-timestamp\" as \"2006-01-02 15:04:05\": cannot parse \"invalid-timestamp\" as \"2006\""),
		},
		{
			name: "Should return error if open is invalid",
			args: args{
				row: []string{"2020-05-01 22:00:00", "invalid-open", "1.0982", "1.0983", "1.0984"},
			},
			want:    nil,
			wantErr: errors.New("strconv.ParseFloat: parsing \"invalid-open\": invalid syntax"),
		},
		{
			name: "Should return error if high is invalid",
			args: args{
				row: []string{"2020-05-01 22:00:00", "1.0981", "invalid-high", "1.0983", "1.0984"},
			},
			want:    nil,
			wantErr: errors.New("strconv.ParseFloat: parsing \"invalid-high\": invalid syntax"),
		},
		{
			name: "Should return error if low is invalid",
			args: args{
				row: []string{"2020-05-01 22:00:00", "1.0981", "1.0982", "invalid-low", "1.0984"},
			},
			want:    nil,
			wantErr: errors.New("strconv.ParseFloat: parsing \"invalid-low\": invalid syntax"),
		},
		{
			name: "Should return error if close is invalid",
			args: args{
				row: []string{"2020-05-01 22:00:00", "1.0981", "1.0982", "1.0983", "invalid-close"},
			},
			want:    nil,
			wantErr: errors.New("strconv.ParseFloat: parsing \"invalid-close\": invalid syntax"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := mapCsvRowToQuote(tt.args.row)

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
