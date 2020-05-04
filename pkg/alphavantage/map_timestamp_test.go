package alphavantage

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_mapTimestamp(t *testing.T) {
	validTime := time.Unix(1587511775, 0).UTC()
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    *time.Time
		wantErr error
	}{
		{
			name: "Should correctly parse valid timestamp",
			args: args{
				value: "2020-04-21 23:29:35",
			},
			want:    &validTime,
			wantErr: nil,
		},
		{
			name: "Should return error if timestamp is invalid",
			args: args{
				value: "invalid-timestamp",
			},
			want:    nil,
			wantErr: errors.New("parsing time \"invalid-timestamp\" as \"2006-01-02 15:04:05\": cannot parse \"invalid-timestamp\" as \"2006\""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := mapTimestamp(tt.args.value)
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
