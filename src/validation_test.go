package src_test

import (
	"testing"

	"euro-exchange/src"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name    string
		input   *src.ExchangeRates
		wantErr bool
	}{
		{
			name: "valid struct",
			input: &src.ExchangeRates{
				Table:    "a",
				Currency: "euro",
				Code:     "EUR",
				Rates: []src.Rate{
					{
						No:   "no-1",
						Date: "2024-01-01",
						Mid:  4.45,
					},
					{
						No:   "no-2",
						Date: "2024-01-02",
						Mid:  4.60,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "invalid currency",
			input: &src.ExchangeRates{
				Table:    "a",
				Currency: "",
				Code:     "EUR",
				Rates: []src.Rate{
					{
						No:   "no-1",
						Date: "2024-01-01",
						Mid:  4.45,
					},
					{
						No:   "no-2",
						Date: "2024-01-02",
						Mid:  4.60,
					},
				},
			},
			wantErr: true,
		},
		{
			name: "invalid rate",
			input: &src.ExchangeRates{
				Table: "a",
				Code:  "EUR",
				Rates: []src.Rate{
					{
						No: "no-1",
					},
					{
						No:   "no-2",
						Date: "2024-01-02",
						Mid:  4.60,
					},
				},
			},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		tc := tc

		err := tc.input.Validate()

		if tc.wantErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}
