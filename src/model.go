package src

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
