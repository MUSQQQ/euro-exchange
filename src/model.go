package src

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ExchangeRates struct {
	Table    string `json:"table"`
	Currency string `json:"currency"`
	Code     string `json:"code"`
	Rates    []Rate `json:"rates"`
}

type Rate struct {
	No   string  `json:"no"`
	Date string  `json:"effectiveDate"`
	Mid  float64 `json:"mid"`
}

func (r Rate) Validate() error {
	return validation.ValidateStruct(
		&r,
		validation.Field(&r.Date, validation.Required),
		validation.Field(&r.Date, validation.Required, validation.Date("2006-01-02")),
		validation.Field(&r.Mid, validation.Required, validation.Min(0.01)),
	)
}

func (er ExchangeRates) Validate() error {
	return validation.ValidateStruct(
		&er,
		validation.Field(&er.Table, validation.Required),
		validation.Field(&er.Currency, validation.Required, validation.In("euro")),
		validation.Field(&er.Code, validation.Required, validation.In("EUR")),
		validation.Field(&er.Rates, validation.Required, validation.Length(1, 100)),
	)
}
