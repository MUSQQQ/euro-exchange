package src

import (
	"euro-exchange/config"

	"github.com/sirupsen/logrus"
)

type App struct {
	client *client
	logger *Logger
}

func NewApp(cfg *config.Config, logger *Logger) *App {

	return &App{
		logger: logger,
		client: newClient(cfg, logger),
	}
}

func (a *App) Run() {

	_, err := a.client.getExchangeRate()
	if err != nil {
		a.logger.Log(logrus.WarnLevel, logrus.Fields{
			"error": err,
		})
	}
}
