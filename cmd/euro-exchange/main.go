package main

import (
	"context"
	"time"

	"euro-exchange/config"
	"euro-exchange/src"

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
		app.CheckLastHundredDays(context.Background())
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
