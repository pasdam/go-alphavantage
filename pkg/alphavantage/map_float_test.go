package alphavantage

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mapFloat(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr error
	}{
		{
			name: "Should return 0 if the value is empty",
			args: args{
				value: "",
			},
			want:    0,
			wantErr: nil,
		},
		{
			name: "Should return 0 if the value is '-'",
			args: args{
				value: "-",
			},
			want:    0,
			wantErr: nil,
		},
		{
			name: "Should return the parsed float if value is valid",
			args: args{
				value: "123.456",
			},
			want:    123.456,
			wantErr: nil,
		},
		{
			name: "Should return error if value is invalid",
			args: args{
				value: "invalid-value",
			},
			want:    0,
			wantErr: errors.New("strconv.ParseFloat: parsing \"invalid-value\": invalid syntax"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := mapFloat(tt.args.value)

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
