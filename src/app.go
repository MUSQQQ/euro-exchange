package src

import (
	"context"
	"euro-exchange/config"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

type App struct {
	client *client
	logger *Logger
	Sleep  time.Duration
}

func NewApp(cfg *config.Config, logger *Logger) *App {

	return &App{
		logger: logger,
		client: newClient(cfg, logger),
		Sleep:  time.Second * time.Duration(cfg.ChecksFrequency),
	}
}

func (a *App) Run(ctx context.Context, wg *sync.WaitGroup) {
	_, err := a.client.getExchangeRate()
	if err != nil {
		a.logger.Log(logrus.WarnLevel, logrus.Fields{
			"error": err,
		})
	}

	ticker := time.Tick(a.Sleep)
	select {
	case <-ticker:
		_, err := a.client.getExchangeRate()
		if err != nil {
			a.logger.Log(logrus.WarnLevel, logrus.Fields{
				"error":         "failed to get exchange rates",
				"error_message": err,
			})
		}
	case <-ctx.Done():
		a.logger.Log(logrus.InfoLevel, logrus.Fields{
			"message": "graceful runner shutdown",
		})
		wg.Done()

	}
}
