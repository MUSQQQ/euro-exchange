package main

import (
	"context"
	"euro-exchange/config"
	"euro-exchange/src"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.NewConfig()

	logger := src.NewLogger(cfg)

	logger.Log(logrus.InfoLevel, logrus.Fields{
		"message": "app start",
	})

	app := src.NewApp(cfg, logger)

	if cfg.CheckLastHundredDays {
		app.CheckLastHundredDays()
	}

	ticker := time.Tick(time.Duration(cfg.Timeout) * time.Second)

	ctx, cancFunc := context.WithCancel(context.Background())

	for i := 0; i < cfg.ChecksNumber; i++ {
		go app.Run(ctx)
	}

	<-ticker
	cancFunc()

	logger.Log(logrus.InfoLevel, logrus.Fields{
		"message": "graceful app shutdown",
	})
}
