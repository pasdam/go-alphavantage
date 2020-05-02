package alphavantage

import "testing"

func TestCurrencyCode_String(t *testing.T) {
	tests := []struct {
		code CurrencyCode
		want string
	}{
		{
			code: CurrencyCodeEUR,
			want: "EUR",
		},
		{
			code: CurrencyCodeUSD,
			want: "USD",
		},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.code.String(); got != tt.want {
				t.Errorf("CurrencyCode.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
