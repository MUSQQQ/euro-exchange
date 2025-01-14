package main

import (
	"euro-exchange/config"
	"euro-exchange/src"

	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.NewConfig()

	logger := src.NewLogger(cfg)

	logger.Log(logrus.InfoLevel, logrus.Fields{
		"app_name": "euro-exchange",
		"message":  "app start",
	})

	app := src.NewApp(cfg, logger)

	app.Run()
}
