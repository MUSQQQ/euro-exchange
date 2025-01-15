package src

import (
	"context"
	"euro-exchange/config"
	"time"

	"github.com/sirupsen/logrus"
)

type App struct {
	Client Client
	logger *Logger
	Sleep  time.Duration
}

func NewApp(cfg *config.Config, logger *Logger) *App {

	return &App{
		logger: logger,
		Client: newClient(cfg, logger),
		Sleep:  time.Second * time.Duration(cfg.ChecksFrequency),
	}
}

func (a *App) Run(ctx context.Context) {
	_, err := a.Client.GetExchangeRate()
	if err != nil {
		a.logger.Log(logrus.WarnLevel, logrus.Fields{
			"error": err,
		})
	}

	ticker := time.Tick(a.Sleep)
	for {
		select {
		case <-ticker:
			_, err := a.Client.GetExchangeRate()
			if err != nil {
				a.logger.Log(logrus.WarnLevel, logrus.Fields{
					"error":         "failed to get exchange rates",
					"error_message": err,
				})
			}
		case <-ctx.Done():
			return
		}
	}
}

func (a *App) CheckLastHundredDays() {
	rates, err := a.Client.GetExchangeRate()
	if err != nil {
		a.logger.Log(logrus.ErrorLevel, logrus.Fields{
			"error":         "failed to check last 100 days",
			"error_message": err,
		})
	}
	var validDates []string
	for _, rate := range rates.Rates {
		if rate.Mid < 4.2 || rate.Mid > 4.3 {
			validDates = append(validDates, rate.Date)
		}
	}

	a.logger.Log(logrus.InfoLevel, logrus.Fields{
		"message": "last 100 dates with rates outside of 4.2-4.3 range",
		"dates":   validDates,
	})
}
