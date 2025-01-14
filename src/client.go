package src

import (
	"encoding/json"
	"net/http"
	"time"

	"euro-exchange/config"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type client struct {
	url        string
	httpClient http.Client
	logger     *Logger
	validator  *validator
}

func newClient(cfg *config.Config, logger *Logger) *client {
	return &client{
		url:        cfg.ExchangeURL,
		httpClient: *http.DefaultClient,
		logger:     logger,
		validator:  newValidator(),
	}
}

func (c *client) getExchangeRate() (any, error) {
	req, err := http.NewRequest("GET", c.url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", uuid.NewString())

	timeStart := time.Now()
	rsp, err := c.httpClient.Do(req)
	timeElapsed := time.Since(timeStart)
	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(rsp.Body)
	var body map[string]interface{}
	err = decoder.Decode(&body)
	if err != nil {
		return nil, err
	}

	c.logger.Log(logrus.InfoLevel, logrus.Fields{
		"time_elapsed": timeElapsed,
		"status_code":  rsp.StatusCode,
		"is_json":      isJSON(rsp.Header.Get("Content-Type")),
		"is_valid":     c.validator.validate(body),
	})

	return nil, nil
}
