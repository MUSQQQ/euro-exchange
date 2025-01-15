package src

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"euro-exchange/config"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

const expectedContentType = "application/json; charset=utf-8"

// mockgen -source=client.go -destination=./mocks/client_mock.go
type Client interface {
	GetExchangeRate(ctx context.Context) (*ExchangeRates, error)
}

type client struct {
	url        string
	httpClient http.Client
	logger     *Logger
}

func newClient(cfg *config.Config, logger *Logger) *client {
	return &client{
		url:        cfg.ExchangeURL,
		httpClient: *http.DefaultClient,
		logger:     logger,
	}
}

func (c *client) GetExchangeRate(ctx context.Context) (*ExchangeRates, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", c.url, nil)
	if err != nil {
		return nil, err
	}

	// workaround to satisfy nbp API
	req.Header.Set("User-Agent", uuid.NewString())

	timeStart := time.Now()
	rsp, err := c.httpClient.Do(req)
	timeElapsed := time.Since(timeStart)
	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(rsp.Body)
	var rates *ExchangeRates
	err = decoder.Decode(&rates)
	if err != nil {
		return nil, err
	}

	c.logger.Log(logrus.InfoLevel, logrus.Fields{
		"time_elapsed": timeElapsed,
		"status_code":  rsp.StatusCode,
		"is_json":      rsp.Header.Get("Content-Type") == expectedContentType,
		"is_valid":     rates.Validate() == nil,
	})

	return rates, nil
}
